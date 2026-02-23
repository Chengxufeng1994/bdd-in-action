# Working with Tables：用表格組織範例與測試資料

## 概述

在 BDD 中，表格（Tables）是表達「多個相似範例」最有效的方式。當同一個業務規則需要用多組資料驗證時，表格能消除重複、提升可讀性，並讓非技術人員也能一目瞭然。

---

## 為什麼需要表格？

當你發現多個場景只差在輸入資料和預期結果時，就是使用表格的時機：

**❌ 重複的場景（難以維護）：**
```gherkin
Scenario: 餘額充足時轉帳成功
  Given Tess 的支票帳戶有 $1,000
  When 她轉 $200 到儲蓄帳戶
  Then 支票帳戶餘額為 $800

Scenario: 剛好等於餘額時轉帳成功
  Given Tess 的支票帳戶有 $500
  When 她轉 $500 到儲蓄帳戶
  Then 支票帳戶餘額為 $0
```

**✅ 用表格整合（清晰簡潔）：**
```gherkin
Scenario Outline: 餘額足夠時轉帳成功
  Given Tess 的支票帳戶有 $<初始餘額>
  When 她轉 $<轉帳金額> 到儲蓄帳戶
  Then 支票帳戶餘額為 $<最終餘額>

  Examples:
    | 初始餘額 | 轉帳金額 | 最終餘額 |
    | 1000     | 200      | 800      |
    | 500      | 500      | 0        |
    | 200      | 1        | 199      |
```

---

## 表格的兩種主要用途

### 1. Scenario Outline + Examples：資料驅動測試

用於同一個行為流程，只有資料不同的情況。

**結構：**
```gherkin
Scenario Outline: <場景名稱>
  Given <條件，使用 <參數>>
  When <動作，使用 <參數>>
  Then <結果，使用 <參數>>

  Examples:
    | 參數1 | 參數2 | 預期結果 |
    | 值1   | 值2   | 結果1    |
    | 值3   | 值4   | 結果2    |
```

**命名建議：**

| 做法 | 說明 |
|------|------|
| 欄位名稱用業務語言 | `初始餘額` 而非 `balance` |
| 每行加描述欄（可選） | `| 情境描述 | 值1 | 值2 |` |
| 涵蓋正常、異常、邊界 | 至少 3 行代表不同情境 |

---

### 2. Data Tables：在單一步驟傳入結構化資料

用於「一次設定多筆資料」，例如初始化多個帳戶、批次交易等。

**結構：**
```gherkin
Scenario: 客戶有多個帳戶
  Given Tess 有以下帳戶：
    | 帳戶類型 | 餘額   | 狀態 |
    | 支票     | $1,000 | 啟用 |
    | 儲蓄     | $2,500 | 啟用 |
    | 投資     | $5,000 | 停用 |
  When 她嘗試從投資帳戶轉帳
  Then 轉帳應該失敗
```

**與 Scenario Outline 的差異：**

| | Scenario Outline | Data Table |
|---|---|---|
| **用途** | 同行為、不同資料 | 單步驟傳入多筆資料 |
| **位置** | 場景末尾的 `Examples:` | 步驟定義之後，縮排對齊 |
| **產生場景數** | 每行一個獨立場景 | 仍是一個場景 |
| **適合情境** | 邊界值測試、多條件驗證 | 批次設定初始狀態 |

---

## Discovery 階段的表格（Example Mapping）

在 Discovery 階段，表格用來描述範例，**不使用 Gherkin 語法**：

```
| 範例 | 初始狀態 | 執行動作 | 預期結果 |
|------|---------|---------|----------|
| 1.1 - 餘額充足 | 支票 $1,000 | 轉 $200 | 支票 $800、成功 |
| 1.2 - 餘額不足 | 支票 $100  | 轉 $200 | 失敗「餘額不足」 |
| 1.3 - 剛好等於 | 支票 $200  | 轉 $200 | 支票 $0、成功   |
```

**Discovery 表格欄位說明：**

| 欄位 | 說明 |
|------|------|
| **範例** | 編號 + 描述，命名要能一眼看懂情境 |
| **初始狀態** | 執行動作前的條件（Given） |
| **執行動作** | 使用者或系統做了什麼（When），通常只有一個 |
| **預期結果** | 應該發生什麼（Then），必須明確可驗證 |

---

## Formulation 階段的轉換

Discovery 表格 → Gherkin 表格的對應關係：

| Discovery 欄位 | Gherkin 關鍵字 |
|---------------|---------------|
| 初始狀態       | `Given` / `Background` |
| 執行動作       | `When` |
| 預期結果       | `Then` |
| 多行範例       | `Scenario Outline` + `Examples:` |

**轉換範例：**

Discovery 表格：
```
| 範例 | 初始狀態 | 動作 | 預期結果 |
| 1.1 - 在限額內 | 支票 $20,000 | 轉 $9,999 | 成功 |
| 1.2 - 剛好達限 | 支票 $20,000 | 轉 $10,000 | 成功 |
| 1.3 - 超過限額 | 支票 $20,000 | 轉 $10,001 | 失敗「超過限額」 |
```

Gherkin 場景：
```gherkin
Scenario Outline: 轉帳限額驗證
  Given Tess 的支票帳戶有 $20,000
  When 她轉 $<金額> 到儲蓄帳戶
  Then 轉帳應該 <結果>

  Examples:
    | 情境       | 金額   | 結果              |
    | 在限額內   | 9999   | 成功              |
    | 剛好達限   | 10000  | 成功              |
    | 超過限額   | 10001  | 失敗並顯示超過限額 |
```

---

## Go 實作：解析 Data Table

在 Godog 中，接收 Data Table 的步驟定義：

```go
// 接收 *godog.Table 型別
func (ts *TransferSteps) clientHasFollowingAccounts(table *godog.Table) error {
    for _, row := range table.Rows[1:] { // 跳過標頭行
        accountType := row.Cells[0].Value
        balance, err := parseAmount(row.Cells[1].Value)
        if err != nil {
            return fmt.Errorf("無效金額 %s: %w", row.Cells[1].Value, err)
        }
        status := row.Cells[2].Value

        account := banking.NewBankAccount(accountType, balance, status == "啟用")
        ts.ctx.RegisterAccount(accountType, account)
    }
    return nil
}
```

**標頭行處理：**
```go
// 方法一：跳過第一行
for _, row := range table.Rows[1:] { ... }

// 方法二：用欄位名稱對應（更安全）
headers := make(map[string]int)
for i, cell := range table.Rows[0].Cells {
    headers[cell.Value] = i
}
for _, row := range table.Rows[1:] {
    accountType := row.Cells[headers["帳戶類型"]].Value
}
```

---

## 最佳實踐

### ✅ 應該做

| 實踐 | 理由 |
|------|------|
| 表格欄位名稱用業務語言 | 業務人員能直接閱讀 |
| 每個 Examples 表格加「情境」欄 | 測試報告中能一眼識別失敗的情境 |
| 涵蓋正常、異常、邊界三種情境 | 確保完整的規則驗證 |
| 表格行數控制在 3-7 行 | 太多行應考慮拆分場景 |

### ❌ 避免做

| 反模式 | 問題 |
|--------|------|
| 欄位名稱用技術術語（`balance_before`）| 業務人員無法閱讀 |
| 一張表格混合多個規則的情境 | 失敗時不知道是哪個規則出問題 |
| 表格超過 10 行 | 應拆分為多個 Scenario Outline |
| Discovery 階段使用 Given/When/Then | 混淆 Discovery 和 Formulation 階段 |

---

## 總結

| 場景 | 使用方式 |
|------|---------|
| Discovery 階段描述範例 | 自然語言表格（初始狀態/動作/預期結果） |
| 同行為、多組資料 | `Scenario Outline` + `Examples:` |
| 步驟需要多筆初始資料 | `Data Table`（步驟後縮排表格） |
| 多個獨立場景 | 各自獨立的 `Scenario`（不強制用表格） |
