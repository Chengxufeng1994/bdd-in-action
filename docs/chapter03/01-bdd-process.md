# BDD 流程：從推測到驗證

## 概述

行為驅動開發 (BDD) 是一個結構化的開發流程，透過六個關鍵步驟將業務需求轉化為可執行的測試和可運行的軟體。這個流程強調協作、溝通和持續驗證。

## BDD 六步驟流程

```
推測 → 說明 → 制定 → 自動化 → 展示 → 驗證
  ↑                                      ↓
  └──────────────── 回饋循環 ─────────────┘
```

---

## 1. 推測 (Speculate) 🤔

### 目標

識別和理解業務價值，探索可能的解決方案

### 關鍵活動

- **探索業務需求**：了解要解決的問題
- **識別利益相關者**：確定誰會從這個功能中受益
- **提出假設**：「我們相信如果實現 X，將會帶來 Y 的價值」
- **定義成功標準**：如何衡量這個功能是否成功

### 參與者

- 產品負責人 (Product Owner)
- 業務分析師 (Business Analyst)
- 開發團隊
- 利益相關者 (Stakeholders)

### 產出

- 功能概述
- 業務假設
- 初步的成功指標

### 實例

```markdown
## 功能：網路銀行轉帳

### 業務假設

我們相信提供網路銀行間轉帳功能，將會：

- 減少客戶到分行的次數 30%
- 提升客戶滿意度至 85%
- 降低人工處理成本

### 利益相關者

- 銀行客戶（主要受益者）
- 客服人員（減少工作量）
- 銀行（降低營運成本）

### 初步成功標準

- 轉帳成功率 > 99%
- 平均轉帳時間 < 3 秒
- 使用率達到 60% 的活躍用戶
```

### 關鍵問題

- 這個功能為誰創造價值？
- 我們如何知道它成功了？
- 有哪些風險和假設？
- 最小可行產品 (MVP) 是什麼？

---

## 2. 說明 (Illustrate) 📖

### 目標

使用具體範例來闡明功能的預期行為

### 關鍵活動

- **範例對映工作坊 (Example Mapping)**：使用便利貼組織範例
- **發現邊界案例**：探索正常流程和異常情況
- **使用真實數據**：基於實際業務場景
- **協作討論**：三方協作（業務、開發、測試）

### 範例對映結構

```
📋 使用者故事（黃色便利貼）
  └── 📏 規則（藍色便利貼）
       └── 📝 範例（綠色便利貼）
            └── ❓ 問題（紅色便利貼）
```

### 參與者

- 業務分析師（主導）
- 開發人員
- 測試人員
- 產品負責人

### 產出

- 具體範例集合
- 業務規則清單
- 待解決的問題清單

### 實例

```markdown
## 使用者故事：轉帳到儲蓄帳戶

### 規則 1：基本轉帳

範例 1：成功轉帳
給定 Tess 的支票帳戶有 $1000
而且 她的儲蓄帳戶有 $2000
當 她從支票帳戶轉 $100 到儲蓄帳戶
那麼 她的支票帳戶應該有 $900
而且 她的儲蓄帳戶應該有 $2100

### 規則 2：餘額不足

範例 2：餘額不足時轉帳失敗
給定 Tess 的支票帳戶有 $100
當 她嘗試轉 $200 到儲蓄帳戶
那麼 轉帳應該失敗
而且 顯示「餘額不足」錯誤訊息

### 規則 3：轉帳限額

範例 3：超過每日限額
給定 Tess 今天已經轉帳 $9,000
而且 每日轉帳限額是 $10,000
當 她嘗試轉 $2,000
那麼 轉帳應該失敗
而且 顯示「超過每日限額」錯誤

### 問題

- ❓ 轉帳限額是否因帳戶類型而異？
- ❓ 轉帳是即時到帳還是有延遲？
- ❓ 是否需要雙重認證？
```

### 範例對映工作坊流程

1. **準備** (5 分鐘)
   - 選擇一個使用者故事
   - 準備便利貼和白板

2. **探索規則** (10-15 分鐘)
   - 識別業務規則
   - 為每個規則寫藍色便利貼

3. **創建範例** (15-20 分鐘)
   - 為每個規則創建具體範例
   - 使用綠色便利貼

4. **識別問題** (5-10 分鐘)
   - 標記不確定的地方
   - 紅色便利貼表示需要澄清的問題

5. **總結** (5 分鐘)
   - 確認理解一致
   - 分配問題解決責任

---

## 3. 制定 (Formulate) ✍️

### 目標

將範例轉化為結構化的 Gherkin 場景

### 關鍵活動

- **編寫 Gherkin 語法**：Given-When-Then 格式
- **組織場景**：按功能和業務能力分類
- **重構場景**：消除重複，提高可讀性
- **審查場景**：確保業務和技術都能理解

### Gherkin 最佳實踐

#### ✅ 好的範例

```gherkin
Scenario: 成功轉帳到儲蓄帳戶
  Given Tess 的支票帳戶有 $1000
  And 她的儲蓄帳戶有 $2000
  When 她從支票帳戶轉 $100 到儲蓄帳戶
  Then 她的支票帳戶餘額應該是 $900
  And 她的儲蓄帳戶餘額應該是 $2100
```

#### ❌ 避免的寫法

```gherkin
# 太過技術性
Scenario: 轉帳測試
  Given 資料庫中有 user_id=123 的使用者
  When 呼叫 API POST /api/transfer 帶參數 {"from": "acc1", "to": "acc2", "amount": 100}
  Then 回應狀態碼應該是 200

# 太過含糊
Scenario: 轉帳功能
  Given 使用者有帳戶
  When 進行轉帳
  Then 轉帳成功
```

### 場景組織結構

```
features/
├── transfers/
│   ├── basic_transfer.feature           # 基本轉帳
│   ├── transfer_limits.feature          # 轉帳限額
│   └── cross_account_transfer.feature   # 跨帳戶轉帳
├── interest/
│   └── earning_interest.feature         # 賺取利息
└── accounts/
    ├── opening_account.feature          # 開戶
    └── closing_account.feature          # 銷戶
```

### 參與者

- 業務分析師（編寫場景）
- 開發人員（技術審查）
- 測試人員（測試覆蓋率審查）
- 產品負責人（業務驗收）

### 產出

- `.feature` 文件
- 場景目錄結構
- 場景審查報告

### 完整 Feature 文件範例

```gherkin
# language: zh-TW
@transfers @high-priority
Feature: 銀行帳戶間轉帳
  作為一個銀行客戶
  我想要在我的帳戶間轉帳
  以便我可以管理我的資金

  Background:
    Given 系統中有以下帳戶類型：
      | 帳戶類型 | 最低餘額 | 每日轉帳限額 |
      | 支票帳戶 | $0       | $10,000     |
      | 儲蓄帳戶 | $100     | $5,000      |

  @smoke
  Scenario: 成功的基本轉帳
    Given Tess 有以下帳戶：
      | 帳戶類型 | 餘額   |
      | 支票帳戶 | $1,000 |
      | 儲蓄帳戶 | $2,000 |
    When 她從支票帳戶轉 $100 到儲蓄帳戶
    Then 轉帳應該成功
    And 她應該看到確認訊息
    And 她的帳戶餘額應該是：
      | 帳戶類型 | 餘額   |
      | 支票帳戶 | $900   |
      | 儲蓄帳戶 | $2,100 |

  @error-handling
  Scenario: 餘額不足時拒絕轉帳
    Given Tess 的支票帳戶有 $100
    And 她的儲蓄帳戶有 $0
    When 她嘗試從支票帳戶轉 $200 到儲蓄帳戶
    Then 轉帳應該失敗
    And 她應該看到錯誤訊息「餘額不足」
    And 她的支票帳戶餘額應該仍然是 $100

  @business-rules
  Scenario Outline: 不同金額的轉帳
    Given Tess 的支票帳戶有 $<初始餘額>
    When 她從支票帳戶轉 $<轉帳金額> 到儲蓄帳戶
    Then 轉帳應該<結果>
    And 她的支票帳戶餘額應該是 $<最終餘額>

    Examples:
      | 初始餘額 | 轉帳金額 | 結果 | 最終餘額 |
      | 1000     | 100      | 成功 | 900      |
      | 1000     | 1000     | 成功 | 0        |
      | 100      | 200      | 失敗 | 100      |
      | 500      | 500      | 成功 | 0        |
```

### Gherkin 撰寫檢查清單

- [ ] 使用業務語言，避免技術細節
- [ ] Given-When-Then 結構清晰
- [ ] 每個場景只測試一個行為
- [ ] 場景名稱描述性強
- [ ] 使用 Scenario Outline 處理多個類似案例
- [ ] 添加適當的標籤 (@tag) 分類
- [ ] 包含 Background 減少重複
- [ ] 所有利益相關者都能理解

---

## 4. 自動化 (Automate) 🤖

### 目標

實現步驟定義，讓場景可執行

### 關鍵活動

- **編寫步驟定義 (Step Definitions)**
- **實現業務邏輯**
- **設計測試架構**
- **建立測試數據管理**

### 步驟定義實現

#### 1. 註冊步驟

```go
// transfer_steps.go
package stepdefinitions

import (
    "github.com/cucumber/godog"
)

type TransferSteps struct {
    ctx         *testcontext.TestContext
    transferApi *actions.TransferApi
}

func NewTransferSteps(ctx *testcontext.TestContext) *TransferSteps {
    return &TransferSteps{
        ctx:         ctx,
        transferApi: actions.NewTransferApi(ctx),
    }
}

func (ts *TransferSteps) RegisterSteps(sc *godog.ScenarioContext) {
    // Given 步驟
    sc.Step(`^(\w+)的(\w+)帳戶有 \$(\d+)$`,
        ts.clientHasAccountWithBalance)

    // When 步驟
    sc.Step(`^她從(\w+)帳戶轉 \$(\d+(?:\.\d+)?) 到(\w+)帳戶$`,
        ts.transfersBetweenAccounts)

    // Then 步驟
    sc.Step(`^她的(\w+)帳戶餘額應該是 \$(\d+(?:\.\d+)?)$`,
        ts.hasBalanceInAccount)

    sc.Step(`^轉帳應該失敗$`,
        ts.transferShouldFail)

    sc.Step(`^她應該看到錯誤訊息「(.+)」$`,
        ts.shouldSeeErrorMessage)
}
```

#### 2. 實現步驟邏輯

```go
func (ts *TransferSteps) clientHasAccountWithBalance(
    clientName, accountTypeStr string,
    balance int,
) error {
    accountType, err := actions.ParseAccountType(accountTypeStr)
    if err != nil {
        return err
    }

    account := banking.BankAccountOfType(accountType).
        WithBalance(float64(balance))

    ts.ctx.RegisterAccount(accountType, account)
    ts.ctx.Client.Opens(account)

    return nil
}

func (ts *TransferSteps) transfersBetweenAccounts(
    fromAccountStr string,
    amount float64,
    toAccountStr string,
) error {
    fromAccountType, err := actions.ParseAccountType(fromAccountStr)
    if err != nil {
        return err
    }

    toAccountType, err := actions.ParseAccountType(toAccountStr)
    if err != nil {
        return err
    }

    fromAccount := ts.ctx.GetAccount(fromAccountType)
    toAccount := ts.ctx.GetAccount(toAccountType)

    err = ts.transferApi.
        TheAmount(amount).
        From(fromAccount).
        To(toAccount)

    // 儲存錯誤供後續驗證
    ts.ctx.LastError = err

    return nil
}

func (ts *TransferSteps) hasBalanceInAccount(
    accountTypeStr string,
    expectedBalance float64,
) error {
    accountType, err := actions.ParseAccountType(accountTypeStr)
    if err != nil {
        return err
    }

    account := ts.ctx.GetAccount(accountType)
    if account == nil {
        return fmt.Errorf("找不到帳戶：%s", accountType)
    }

    actualBalance := account.Balance()
    if math.Abs(actualBalance-expectedBalance) > 0.01 {
        return fmt.Errorf(
            "預期餘額 %.2f，實際餘額 %.2f",
            expectedBalance,
            actualBalance,
        )
    }

    return nil
}
```

### 測試架構分層

```
tests/acceptancetests/
├── acceptance_suite_test.go    # 測試套件入口
├── testcontext/                # 測試上下文
│   └── test_context.go
├── actions/                    # 業務動作層（Fluent API）
│   ├── transfer_api.go
│   └── account_type_parser.go
├── stepdefinitions/            # 步驟定義
│   ├── transfer_steps.go
│   ├── interest_steps.go
│   └── helpers.go
└── domain/                     # 測試領域模型
    └── initial_account.go
```

### 測試上下文管理

```go
// test_context.go
package testcontext

type TestContext struct {
    Client              *banking.Client
    InterestCalculator  *banking.InterestCalculator
    LastError           error
    LastTransaction     banking.Transaction
    Accounts            map[banking.AccountType]*banking.BankAccount
}

func NewTestContext() *TestContext {
    return &TestContext{
        Client:             banking.NewClient("Tess"),
        InterestCalculator: banking.NewInterestCalculator(),
        Accounts:           make(map[banking.AccountType]*banking.BankAccount),
    }
}

func (tc *TestContext) Reset() {
    tc.Client = banking.NewClient("Tess")
    tc.InterestCalculator = banking.NewInterestCalculator()
    tc.LastError = nil
    tc.Accounts = make(map[banking.AccountType]*banking.BankAccount)
}
```

### Fluent API 設計

```go
// transfer_api.go
package actions

type TransferApi struct {
    ctx         *testcontext.TestContext
    amount      float64
    fromAccount *banking.BankAccount
}

func NewTransferApi(ctx *testcontext.TestContext) *TransferApi {
    return &TransferApi{ctx: ctx}
}

// 鏈式調用設計
func (ta *TransferApi) TheAmount(amount float64) *TransferApi {
    ta.amount = amount
    return ta
}

func (ta *TransferApi) From(account *banking.BankAccount) *TransferApi {
    ta.fromAccount = account
    return ta
}

func (ta *TransferApi) To(account *banking.BankAccount) error {
    if err := ta.fromAccount.Withdraw(ta.amount); err != nil {
        ta.ctx.LastError = err
        return err
    }
    account.Deposit(ta.amount)
    return nil
}

// 使用範例：
// transferApi.TheAmount(100).From(checkingAccount).To(savingsAccount)
```

### 參與者

- 開發人員（主導）
- 測試自動化工程師
- 業務分析師（驗證步驟定義的可讀性）

### 產出

- 步驟定義代碼
- 測試基礎設施
- 可執行的測試套件

---

## 5. 展示 (Demonstrate) 📊

### 目標

執行自動化測試，展示功能行為

### 關鍵活動

- **運行 BDD 測試**
- **生成測試報告**
- **展示給利益相關者**
- **收集反饋**

### 執行測試

```bash
# 運行所有測試
make test

# 運行特定功能
make test-transfers

# 生成覆蓋率報告
make test-cucumber-report
```

### 測試輸出範例

```
Feature: 銀行帳戶間轉帳

  Scenario: 成功的基本轉帳                                    # features/transfers/basic_transfer.feature:8
    Given Tess 的支票帳戶有 $1000                              # transfer_steps.go:31
    And 她的儲蓄帳戶有 $2000                                   # transfer_steps.go:31
    When 她從支票帳戶轉 $100 到儲蓄帳戶                         # transfer_steps.go:43
    Then 她的支票帳戶餘額應該是 $900                            # transfer_steps.go:58
    And 她的儲蓄帳戶餘額應該是 $2100                            # transfer_steps.go:58

  Scenario: 餘額不足時拒絕轉帳                                 # features/transfers/basic_transfer.feature:15
    Given Tess 的支票帳戶有 $100                               # transfer_steps.go:31
    And 她的儲蓄帳戶有 $0                                      # transfer_steps.go:31
    When 她嘗試從支票帳戶轉 $200 到儲蓄帳戶                     # transfer_steps.go:43
    Then 轉帳應該失敗                                          # transfer_steps.go:72
    And 她應該看到錯誤訊息「餘額不足」                          # transfer_steps.go:78

2 scenarios (2 passed)
10 steps (10 passed)
0m0.123s
```

### Cucumber JSON 報告

```json
{
  "uri": "features/transfers/basic_transfer.feature",
  "id": "bank-transfer",
  "keyword": "Feature",
  "name": "銀行帳戶間轉帳",
  "line": 1,
  "elements": [
    {
      "id": "bank-transfer;successful-transfer",
      "keyword": "Scenario",
      "name": "成功的基本轉帳",
      "line": 8,
      "type": "scenario",
      "steps": [
        {
          "keyword": "Given ",
          "name": "Tess 的支票帳戶有 $1000",
          "line": 9,
          "match": {
            "location": "transfer_steps.go:31"
          },
          "result": {
            "status": "passed",
            "duration": 123456
          }
        }
      ]
    }
  ]
}
```

### 展示會議

#### 準備

1. 確保所有測試都通過
2. 準備展示環境
3. 準備測試報告
4. 邀請所有利益相關者

#### 展示流程

1. **回顧業務目標** (5 分鐘)
   - 重申功能目的
   - 回顧成功標準

2. **運行活文檔** (10 分鐘)
   - 即時執行 BDD 測試
   - 展示測試結果
   - 說明場景覆蓋範圍

3. **展示功能** (15 分鐘)
   - 實際操作功能
   - 展示關鍵場景
   - 處理邊界案例

4. **測試報告** (10 分鐘)
   - 顯示 Cucumber 報告
   - 說明測試覆蓋率
   - 展示性能指標

5. **收集反饋** (15 分鐘)
   - 回答問題
   - 記錄建議
   - 識別改進點

### 參與者

- 開發團隊（展示）
- 產品負責人（驗收）
- 利益相關者（反饋）
- 測試團隊（質量保證）

### 產出

- 測試執行報告
- 反饋清單
- 改進建議

---

## 6. 驗證 (Validate) ✅

### 目標

確認功能符合業務預期並創造價值

### 關鍵活動

- **業務驗收**
- **度量指標檢查**
- **使用者反饋收集**
- **持續改進**

### 驗證清單

#### ✅ 功能驗證

- [ ] 所有場景都通過
- [ ] 邊界案例都被覆蓋
- [ ] 錯誤處理正確
- [ ] 性能符合要求

#### ✅ 業務驗證

- [ ] 符合業務規則
- [ ] 滿足使用者需求
- [ ] 創造預期價值
- [ ] 利益相關者接受

#### ✅ 技術驗證

- [ ] 代碼質量良好
- [ ] 測試覆蓋率達標
- [ ] 文檔完整
- [ ] 安全性檢查通過

### 成功指標檢查

```markdown
## 初始假設驗證

### 假設 1：減少客戶到分行次數

- 目標：減少 30%
- 實際：減少 35% ✅
- 證據：分行流量統計

### 假設 2：提升客戶滿意度

- 目標：達到 85%
- 實際：達到 82% ⚠️
- 行動：收集反饋改進 UI

### 假設 3：降低人工處理成本

- 目標：降低營運成本
- 實際：節省 40% 人工時間 ✅
- 證據：客服工單數據
```

### A/B 測試驗證

```markdown
## 轉帳功能 A/B 測試結果

### 測試組

- 樣本：1000 位使用者
- 期間：2 週
- 轉帳成功率：99.2%
- 平均完成時間：2.8 秒
- 使用率：65%

### 對照組（原有流程）

- 樣本：1000 位使用者
- 期間：2 週
- 轉帳成功率：95.1%
- 平均完成時間：45 秒（人工處理）
- 使用率：N/A

### 結論

新功能顯著改善使用者體驗 ✅
```

### 持續改進循環

```
驗證結果 → 識別改進點 → 新的推測 → 開始新循環
```

#### 改進案例

```markdown
## 從驗證中學習

### 發現問題

用戶在轉帳限額部分感到困惑

### 根本原因

錯誤訊息不夠明確

### 改進行動

1. 推測：改善錯誤訊息會提高理解度
2. 說明：定義更清楚的錯誤訊息範例
3. 制定：更新 Gherkin 場景
4. 自動化：實現新的錯誤訊息
5. 展示：向使用者展示改進
6. 驗證：測量理解度提升

### 結果

客戶滿意度從 82% 提升到 87% ✅
```

### 參與者

- 產品負責人（業務驗收）
- 使用者（實際使用反饋）
- 開發團隊（技術驗收）
- 資料分析師（指標分析）

### 產出

- 驗收報告
- 指標儀表板
- 改進建議清單
- 學習筆記

---

## 流程整合與最佳實踐

### 迭代週期

```
Sprint 1 週期
├── 週一：推測 + 說明
│   └── Example Mapping 工作坊
├── 週二：制定
│   └── 編寫 Gherkin 場景
├── 週三-四：自動化
│   └── 實現步驟定義和功能
├── 週五：展示 + 驗證
│   └── Sprint Review 和回顧
└── 持續：測試執行
    └── CI/CD 自動化
```

### 協作模式

```
三方協作 (Three Amigos)
├── 業務角色
│   ├── 產品負責人
│   └── 業務分析師
├── 開發角色
│   ├── 前端開發
│   └── 後端開發
└── 測試角色
    ├── 測試工程師
    └── QA 專家
```

### 活文檔

BDD 場景作為活文檔：

- **可執行**：自動測試驗證
- **可讀**：業務人員能理解
- **可維護**：隨需求演進更新
- **可追溯**：連結需求與實現

### 常見陷阱與解決方案

| 陷阱           | 症狀             | 解決方案                     |
| -------------- | ---------------- | ---------------------------- |
| **過度技術化** | 場景充滿技術術語 | 回到業務語言，重寫場景       |
| **範例不足**   | 只有快樂路徑     | Example Mapping 發現邊界案例 |
| **測試太慢**   | CI/CD 時間過長   | 分層測試，優化測試數據       |
| **維護困難**   | 場景過時失效     | 納入 DoD，持續重構           |
| **缺乏協作**   | 各自為政         | 定期 Three Amigos 會議       |

### 度量指標

```markdown
## BDD 健康度指標

### 測試覆蓋

- 場景數量：42
- 步驟數量：186
- 功能覆蓋率：95%
- 業務規則覆蓋：100%

### 測試質量

- 通過率：98.5%
- 執行時間：3 分 45 秒
- 失敗場景：2（已知問題）
- 碎片化測試：0

### 協作效率

- Example Mapping 參與率：90%
- 場景審查週期：< 1 天
- 三方協作頻率：每週 2 次
- 文檔更新及時性：95%

### 業務價值

- 缺陷提前發現率：75%
- 需求變更適應速度：< 2 天
- 利益相關者滿意度：85%
- 交付週期縮短：40%
```

---

## 總結

BDD 六步驟流程提供了一個完整的框架，從業務假設到價值驗證：

1. **推測**：探索業務價值和假設
2. **說明**：用具體範例闡明需求
3. **制定**：轉化為結構化場景
4. **自動化**：實現可執行測試
5. **展示**：運行活文檔展示行為
6. **驗證**：確認業務價值實現

### 關鍵成功因素

- ✅ **協作優先**：三方持續溝通
- ✅ **業務導向**：始終關注價值
- ✅ **具體範例**：避免抽象描述
- ✅ **快速反饋**：短週期迭代
- ✅ **持續改進**：學習與適應

### 下一步

1. 選擇一個小功能開始實踐
2. 組織第一次 Example Mapping 工作坊
3. 建立 BDD 測試基礎設施
4. 定期回顧和改進流程

---

**記住**：BDD 不只是測試工具，而是一種協作和溝通的方式。透過這六個步驟，我們確保團隊始終聚焦於交付真正的業務價值。
