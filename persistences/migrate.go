package persistences

import (
"bufio"
"bytes"
"fmt"
"io/ioutil"
"path/filepath"
"strconv"
"strings"

"github.com/go-xorm/xorm"
"github.com/go-xorm/xorm/migrate"
)

const (
	migrateFileSuffix = ".sql"
	upMarker          = "# --- !Ups"
	downMarker        = "# --- !Downs"
	migrationIDLen    = 12
)

// CreateMigrationTableIfNotExists は migration テーブルを作成する.
// xorm.migrate.createMigrationTableIfNotExists は ROW_FORMAT を指定できないので使わない.
func CreateMigrationTableIfNotExists(db *xorm.Engine, options *migrate.Options) error {
	exists, err := db.IsTableExist(options.TableName)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}

	sql := fmt.Sprintf("CREATE TABLE %s (%s VARCHAR(255) PRIMARY KEY) ROW_FORMAT = DYNAMIC ENGINE = INNODB",
		options.TableName,
		options.IDColumnName)
	if _, err := db.Exec(sql); err != nil {
		return err
	}
	return nil
}

// ReadMigrationDir は dir から *.sql ファイルを読んで migration スライスを返す.
func ReadMigrationDir(dir string) ([]*migrate.Migration, error) {
	var migrations = make([]*migrate.Migration, 0, 100)

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("%s cannot open as a migration directory", dir)
	}
	for _, file := range files {
		fileName := file.Name()

		if file.IsDir() || !strings.HasSuffix(fileName, migrateFileSuffix) {
			continue
		}

		migrationID := strings.TrimSuffix(fileName, migrateFileSuffix)
		usIndex := strings.Index(migrationID, "_")
		if usIndex != -1 {
			migrationID = migrationID[:usIndex]
		}
		filePath := filepath.Join(dir, fileName)
		fileData, err := ioutil.ReadFile(filePath)
		if err != nil {
			return nil, err
		}
		scanner := bufio.NewScanner(strings.NewReader(string(fileData)))
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if line == upMarker {
				break
			}
		}
		// up statements の読み込み.
		upStatements := bytes.NewBuffer(make([]byte, 0, 1024))
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())

			if line == downMarker {
				break
			} else if strings.HasPrefix(line, "#") { // skip comment
				continue
			} else {
				upStatements.WriteString(line)
				upStatements.WriteString("\n")
			}
		}

		// down statements の読み込み.
		downStatements := bytes.NewBuffer(make([]byte, 0, 1024))
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())

			if line == downMarker {
				break
			} else if strings.HasPrefix(line, "#") { // skip comment
				continue
			} else {
				downStatements.WriteString(line)
				downStatements.WriteString("\n")
			}
		}

		upStatementsStr := upStatements.String()
		downStatementsStr := downStatements.String()

		if len(migrationID) != migrationIDLen {
			return nil, fmt.Errorf("%s: file name prefix can only contains timestamp", fileName)
		}
		if _, err := strconv.Atoi(migrationID); err != nil {
			return nil, fmt.Errorf("%s: file name prefix can only contains number characters", fileName)
		}
		if len(upStatementsStr) == 0 {
			return nil, fmt.Errorf("%s: up statement not found", fileName)
		}
		if len(downStatementsStr) == 0 {
			return nil, fmt.Errorf("%s: down statement not found", fileName)
		}

		migrations = append(migrations, &migrate.Migration{
			ID: migrationID,
			Migrate: func(tx *xorm.Engine) error {
				_, err := tx.Exec(upStatementsStr)
				return err
			},
			Rollback: func(tx *xorm.Engine) error {
				_, err := tx.Exec(downStatementsStr)
				return err
			},
		})
	}

	return migrations, nil
}
