package persistences

import (
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/xorm/migrate"
	"github.com/jinzhu/inflection"
	"xorm.io/core"
)

var engine *xorm.Engine

// InitEngine は xorm engine の初期化を行う.
func InitEngine(driver, dataSource string, showSQL bool, migrationPath string) error {
	var err error
	engine, err = xorm.NewEngine(driver, dataSource)
	if err != nil {
		return err
	}

	// NEWS テーブルは entity も News なので複数 <-> 単数変換は無効.
	inflection.AddUncountable("news")

	engine.SetTableMapper(core.NewCacheMapper(&LintGonicLowerCasePluralizeMapper))
	engine.SetColumnMapper(core.NewCacheMapper(&LintGonicLowerCaseMapper))
	//engine.SetLogger(utils.NewXormLogger(log.StandardLogger()))

	if showSQL {
		// Debug: SQL Output する場合有効にする.
		engine.ShowSQL(true)
	}

	// Migration 定義をファイルから読み込み.
	if migrationPath != "" {
		migrations, err := ReadMigrationDir(migrationPath)
		if err != nil {
			return err
		}

		// migration テーブルがなければ作成.
		err = CreateMigrationTableIfNotExists(engine, migrate.DefaultOptions)
		if err != nil {
			return err
		}

		// Migration.
		m := migrate.New(engine, migrate.DefaultOptions, migrations)
		err = m.Migrate()
		if err != nil {
			return err
		}
	}

	return nil
}

// GetEngine は xorm engine の取得を行う.
func GetEngine() *xorm.Engine {
	return engine
}

// NewSession は xorm session を開始する.
func NewSession() *xorm.Session {
	return engine.NewSession()
}
