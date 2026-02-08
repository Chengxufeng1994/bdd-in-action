# BDD 流程快速參考

## 六步驟速查表

| 步驟 | 目標 | 主要活動 | 產出 | 時間 |
|------|------|----------|------|------|
| **1. 推測** 🤔 | 探索價值 | 識別假設、定義成功標準 | 功能概述、成功指標 | 1-2 小時 |
| **2. 說明** 📖 | 具體範例 | Example Mapping 工作坊 | 範例集合、業務規則 | 30-60 分鐘 |
| **3. 制定** ✍️ | 結構化場景 | 編寫 Gherkin | .feature 文件 | 1-2 小時 |
| **4. 自動化** 🤖 | 可執行測試 | 實現步驟定義 | 測試代碼 | 4-8 小時 |
| **5. 展示** 📊 | 運行測試 | 執行和報告 | 測試報告 | 30 分鐘 |
| **6. 驗證** ✅ | 確認價值 | 業務驗收、度量 | 驗收報告 | 1-2 小時 |

---

## Example Mapping 速查

### 四色便利貼系統

```
📋 黃色 = 使用者故事
📏 藍色 = 業務規則
📝 綠色 = 具體範例
❓ 紅色 = 待解決問題
```

### 時間盒

- **設定時限**：25-30 分鐘
- **如果超時**：故事太大，需要拆分
- **如果問題太多**：需要更多研究

### 完成標準

- ✅ 所有規則都有至少 2 個範例
- ✅ 包含正常流程和異常案例
- ✅ 紅色便利貼 < 3 張
- ✅ 所有參與者理解一致

---

## Gherkin 語法速查

### 基本結構

```gherkin
Feature: [功能名稱]
  [功能描述]

  Background:
    [所有場景的共同前置條件]

  Scenario: [場景名稱]
    Given [前置條件]
    And [更多前置條件]
    When [執行的動作]
    And [更多動作]
    Then [預期結果]
    And [更多預期結果]
    But [不應該發生的事]
```

### Scenario Outline

```gherkin
Scenario Outline: [場景範本名稱]
  Given [使用 <參數> 的步驟]
  When [使用 <參數> 的步驟]
  Then [使用 <參數> 的步驟]

  Examples:
    | 參數1 | 參數2 | 預期結果 |
    | 值1   | 值2   | 結果1    |
    | 值3   | 值4   | 結果2    |
```

### 資料表

```gherkin
Given 以下帳戶：
  | 帳戶類型 | 餘額  |
  | 支票     | 1000 |
  | 儲蓄     | 2000 |
```

### 標籤

```gherkin
@smoke @high-priority @wip
Scenario: 重要的冒煙測試場景
```

常用標籤：
- `@smoke` - 冒煙測試
- `@wip` - 開發中
- `@bug` - 缺陷修復
- `@slow` - 執行較慢
- `@manual` - 手動測試

---

## 步驟定義模式速查

### 基本模式

```go
// Given 步驟 - 設置前置條件
func (s *Steps) given前置條件(參數 string) error {
    // 準備測試數據
    return nil
}

// When 步驟 - 執行動作
func (s *Steps) when執行動作(參數 string) error {
    // 執行業務邏輯
    // 捕獲結果或錯誤
    return nil
}

// Then 步驟 - 驗證結果
func (s *Steps) then驗證結果(預期 string) error {
    // 斷言
    if 實際 != 預期 {
        return fmt.Errorf("預期 %s，得到 %s", 預期, 實際)
    }
    return nil
}
```

### 正則表達式模式

```go
// 匹配數字
sc.Step(`^餘額是 \$(\d+(?:\.\d+)?)$`, s.balance)

// 匹配文字
sc.Step(`^(\w+) 有一個 (\w+) 帳戶$`, s.hasAccount)

// 匹配表格
sc.Step(`^以下帳戶：$`, s.accountsTable)

// 可選文字
sc.Step(`^轉帳(應該)?失敗$`, s.transferFails)
```

### 常用參數類型

| 模式 | 說明 | 範例 |
|------|------|------|
| `(\d+)` | 整數 | `100` |
| `(\d+\.\d+)` | 小數 | `99.95` |
| `(\d+(?:\.\d+)?)` | 整數或小數 | `100` 或 `100.50` |
| `(\w+)` | 單字 | `Current` |
| `(.+)` | 任意文字 | `error message` |
| `"([^"]*)"` | 引號內文字 | `"Hello World"` |

---

## 測試架構速查

### 目錄結構

```
tests/acceptancetests/
├── acceptance_suite_test.go    # 套件配置
├── testcontext/                # 測試上下文
│   └── test_context.go
├── actions/                    # Fluent API
│   └── *_api.go
├── stepdefinitions/            # 步驟定義
│   └── *_steps.go
└── domain/                     # 測試領域模型
    └── *.go

features/
├── 功能1/
│   └── *.feature
└── 功能2/
    └── *.feature
```

### 測試上下文模式

```go
type TestContext struct {
    // 系統對象
    Client *domain.Client

    // 測試狀態
    LastError   error
    LastResult  interface{}

    // 測試數據
    TestData map[string]interface{}
}

func (tc *TestContext) Reset() {
    // 重置所有狀態
}
```

### Fluent API 模式

```go
type API struct {
    ctx    *TestContext
    params map[string]interface{}
}

func (a *API) With參數(value string) *API {
    a.params["key"] = value
    return a
}

func (a *API) Execute() error {
    // 執行並儲存結果到 ctx
    return nil
}
```

---

## 常見場景範本

### 1. CRUD 操作

```gherkin
Scenario: 創建新實體
  Given 我是已登入的使用者
  When 我創建一個新的 [實體] 帶著：
    | 欄位1 | 值1 |
    | 欄位2 | 值2 |
  Then [實體] 應該被成功創建
  And 我應該看到 [實體] 詳情

Scenario: 讀取實體
  Given 系統中存在一個 [實體]
  When 我查詢該 [實體]
  Then 我應該看到正確的 [實體] 資訊

Scenario: 更新實體
  Given 系統中存在一個 [實體]
  When 我更新 [實體] 的 [欄位] 為 [新值]
  Then [實體] 應該被成功更新
  And [欄位] 應該是 [新值]

Scenario: 刪除實體
  Given 系統中存在一個 [實體]
  When 我刪除該 [實體]
  Then [實體] 應該被成功刪除
  And 我不應該能找到該 [實體]
```

### 2. 驗證場景

```gherkin
Scenario: 輸入驗證
  Given 我在 [表單] 頁面
  When 我輸入無效的 [欄位]
  Then 我應該看到錯誤訊息 "[錯誤訊息]"
  And [欄位] 應該被標記為錯誤
```

### 3. 權限場景

```gherkin
Scenario: 未授權訪問
  Given 我是未登入的使用者
  When 我嘗試訪問 [受保護資源]
  Then 我應該被重定向到登入頁面

Scenario: 權限不足
  Given 我是 [角色] 使用者
  When 我嘗試 [需要更高權限的操作]
  Then 我應該看到「權限不足」錯誤
```

### 4. 工作流程場景

```gherkin
Scenario: 多步驟流程
  Given 我在流程的 [起始狀態]
  When 我執行 [步驟1]
  And 我執行 [步驟2]
  And 我執行 [步驟3]
  Then 流程應該到達 [最終狀態]
  And 系統應該記錄所有步驟
```

---

## 執行與報告速查

### Make 命令

```bash
# 執行所有測試
make test

# 詳細輸出
make test-verbose

# 生成報告
make test-cucumber-report

# 特定功能
make test-transfers
make test-interest
```

### Godog 命令列

```bash
# 基本執行
go test -v

# 指定格式
go test -v --godog.format=pretty
go test -v --godog.format=progress
go test -v --godog.format=cucumber

# 標籤過濾
go test -v --godog.tags="@smoke"
go test -v --godog.tags="@smoke && @high-priority"
go test -v --godog.tags="~@wip"  # 排除 @wip

# 指定場景
go test -v --godog.paths=features/transfers/
```

### 測試輸出格式

| 格式 | 說明 | 用途 |
|------|------|------|
| `pretty` | 彩色詳細輸出 | 開發階段 |
| `progress` | 簡潔進度條 | CI/CD |
| `cucumber` | JSON 格式 | 報告生成 |
| `junit` | JUnit XML | CI 整合 |

---

## 問題排查速查

### 常見問題

| 問題 | 原因 | 解決方案 |
|------|------|----------|
| 步驟未定義 | 缺少步驟定義 | 實現對應的步驟函數 |
| 步驟匹配多個定義 | 正則表達式衝突 | 使更精確的模式 |
| 測試超時 | 執行時間過長 | 優化測試數據，分層測試 |
| 隨機失敗 | 測試數據污染 | 確保 Reset() 清理乾淨 |
| 場景難以理解 | 過於技術化 | 改用業務語言重寫 |

### 除錯技巧

```go
// 1. 添加日誌
func (s *Steps) whenAction() error {
    log.Printf("執行動作，參數：%v", s.params)
    // ...
}

// 2. 斷言前檢查
func (s *Steps) thenResult(expected string) error {
    log.Printf("預期：%s，實際：%s", expected, s.actual)
    if s.actual != expected {
        return fmt.Errorf("不匹配")
    }
    return nil
}

// 3. 使用表格輸出
func printTable(data [][]string) {
    for _, row := range data {
        log.Printf("| %s |", strings.Join(row, " | "))
    }
}
```

---

## 最佳實踐檢查清單

### Gherkin 場景 ✅

- [ ] 使用業務語言，不是技術術語
- [ ] 每個場景測試一個行為
- [ ] Given-When-Then 結構清晰
- [ ] 場景名稱描述性強
- [ ] 避免 UI 細節（按鈕、連結）
- [ ] 使用 Background 消除重複
- [ ] Scenario Outline 處理變化
- [ ] 標籤分類合理

### 步驟定義 ✅

- [ ] 步驟定義可重用
- [ ] 正則表達式精確
- [ ] 錯誤訊息清晰
- [ ] 避免步驟間依賴
- [ ] 使用 Helper 函數
- [ ] 測試數據隔離
- [ ] 每次測試後清理

### 測試架構 ✅

- [ ] 清晰的目錄結構
- [ ] 測試上下文管理良好
- [ ] Fluent API 提高可讀性
- [ ] 測試執行快速（< 5 分鐘）
- [ ] 測試穩定不碎片
- [ ] 文檔與代碼同步

---

## 度量指標速查

### 測試健康度

```markdown
✅ 測試通過率：> 95%
✅ 測試執行時間：< 5 分鐘
✅ 場景覆蓋率：> 80%
✅ 步驟重用率：> 60%
✅ 文檔新鮮度：< 7 天
```

### 協作效率

```markdown
✅ Example Mapping 參與率：> 80%
✅ 場景審查週期：< 1 天
✅ Three Amigos 頻率：每週 ≥ 2 次
✅ 利益相關者滿意度：> 80%
```

### 業務價值

```markdown
✅ 缺陷提前發現率：> 70%
✅ 需求變更適應速度：< 2 天
✅ 交付週期縮短：> 30%
✅ 返工率降低：> 40%
```

---

## 快速鍵與提示

### VS Code 擴展

- **Cucumber (Gherkin) Full Support**
  - 語法高亮
  - 自動補全
  - 步驟跳轉

### 命令別名

```bash
# .bashrc / .zshrc
alias bdd-test="go test -v ./tests/acceptancetests"
alias bdd-smoke="go test -v --godog.tags='@smoke'"
alias bdd-report="go test -v --godog.format=cucumber"
```

### Git Hooks

```bash
# .git/hooks/pre-commit
#!/bin/sh
make test || exit 1
```

---

**此速查表提供 BDD 流程的快速參考，建議列印或加入書籤！** 📑
