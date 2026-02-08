# Train Timetables - BDD 實作範例

這是一個使用 BDD (Behavior-Driven Development) 方法開發的火車時刻表系統。

## 專案結構

```
train-timetables/
├── domain/                 # 領域模型和業務邏輯
├── tests/
│   ├── acceptancetests/   # BDD 驗收測試
│   │   ├── stepdefinitions/  # Godog 步驟定義
│   │   ├── testcontext/      # 測試上下文
│   │   └── acceptance_suite_test.go
│   ├── features/          # Gherkin 場景檔案
│   └── unitests/          # 單元測試
├── go.mod
├── Makefile
└── README.md
```

## 快速開始

### 安裝依賴

```bash
make deps
```

### 執行測試

```bash
# 執行所有驗收測試
make test

# 詳細輸出
make test-verbose

# 生成 Cucumber JSON 報告
make test-cucumber-report
```

### 開發工作流程

1. **撰寫 Feature 檔案**
   - 在 `tests/features/` 目錄下使用 Gherkin 語法描述業務行為

2. **實現步驟定義**
   - 在 `tests/acceptancetests/stepdefinitions/` 實現 Godog 步驟

3. **開發領域邏輯**
   - 在 `domain/` 目錄實現業務邏輯

4. **執行測試**
   - 使用 `make test` 驗證實現

## BDD 工作流程

本專案遵循 BDD 六步驟流程：

1. **推測 (Speculate)** - 識別業務價值
2. **說明 (Illustrate)** - Example Mapping 發現範例
3. **制定 (Formulate)** - 撰寫 Gherkin 場景
4. **自動化 (Automate)** - 實現步驟定義
5. **展示 (Demonstrate)** - 執行測試展示結果
6. **驗證 (Validate)** - 確認業務價值

詳細指南請參考：[docs/chapter03/](../../docs/chapter03/)

## Makefile 指令

```bash
make help                   # 顯示所有可用指令
make build                  # 編譯專案
make test                   # 執行驗收測試
make test-verbose           # 詳細輸出測試結果
make test-cucumber-report   # 生成 Cucumber JSON 報告
make fmt                    # 格式化程式碼
make vet                    # 執行 go vet
make lint                   # 執行 lint（包含 fmt 和 vet）
make tidy                   # 整理依賴
make clean                  # 清理測試產物
make deps                   # 安裝/更新依賴
make all                    # 執行 lint、build 和 test
```

## 技術棧

- **語言**: Go 1.25.6
- **BDD 框架**: Godog v0.15.1
- **測試架構**: 分層測試（驗收測試 + 單元測試）

## 範例場景

```gherkin
# language: zh-TW
功能: 查詢火車時刻表
  作為一個乘客
  我想要查詢火車時刻表
  以便我可以規劃我的行程

  場景: 查詢兩個車站之間的火車
    假如 a train schedule exists
    當 I search for trains from "台北" to "台中"
    那麼 I should see 5 available trains
```

## 下一步

1. ✅ 已完成 BDD 基礎架構設置
2. ⏭️ 使用 Example Mapping 發現更多業務規則
3. ⏭️ 撰寫完整的 Gherkin 場景
4. ⏭️ 實現領域模型
5. ⏭️ 實現步驟定義
6. ⏭️ 執行並驗證測試

## 參考資源

- [BDD in Action](../../docs/chapter03/)
- [Godog Documentation](https://github.com/cucumber/godog)
- [Gherkin Reference](https://cucumber.io/docs/gherkin/reference/)
