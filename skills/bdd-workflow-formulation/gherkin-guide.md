# Gherkin 編寫指南

## 概述

Gherkin 是一種業務可讀的領域特定語言（DSL），用於描述軟體行為而不描述實作細節。它使用結構化的自然語言來定義測試案例。

**重要：** Gherkin 編寫屬於 BDD 流程的**制定階段（Formulate）**，應該在**探索階段（Discovery）**完成後才進行。

## BDD 流程中的位置

```
1. 推測（Speculate）     → 識別業務價值
2. 說明（Illustrate）    → Example Mapping 發現範例  ← Discovery 階段
3. 制定（Formulate）     → 撰寫 Gherkin 場景        ← 在這裡寫 Gherkin！
4. 自動化（Automate）    → 實現步驟定義
5. 展示（Demonstrate）   → 執行測試展示結果
6. 驗證（Validate）      → 確認業務價值
```

## Gherkin 基本語法

### Feature（功能）

描述整體功能及其業務價值：

```gherkin
Feature: 銀行帳戶間轉帳
  As a 銀行客戶
  I want 在我的帳戶間轉移資金
  So that 我可以靈活管理我的財務
```

### Scenario（場景）

描述具體的業務案例：

```gherkin
Scenario: 成功的轉帳
  Given 支票帳戶有 $1,000
  And 儲蓄帳戶有 $500
  When 我從支票帳戶轉 $200 到儲蓄帳戶
  Then 支票帳戶應該有 $800
  And 儲蓄帳戶應該有 $700
```

### 關鍵字

- **Feature**: 功能描述
- **Scenario**: 場景描述
- **Given**: 初始條件（前置狀態）
- **When**: 執行動作（觸發事件）
- **Then**: 預期結果（驗證結果）
- **And**: 連接詞（連接多個步驟）
- **But**: 對比連接詞
- **Background**: 所有場景共用的前置條件
- **Scenario Outline**: 參數化場景
- **Examples**: 提供測試資料

## 從 Example Mapping 轉換到 Gherkin

### Example Mapping 輸出範例

```
📏 規則 1：餘額驗證
來源帳戶必須有足夠的餘額來完成轉帳

📝 範例 1.1：餘額充足
Given 支票帳戶有 $1,000
And 儲蓄帳戶有 $500
When 從支票轉 $200 到儲蓄
Then 支票帳戶剩 $800
And 儲蓄帳戶變成 $700

📝 範例 1.2：餘額不足
Given 支票帳戶有 $100
When 嘗試轉 $200 到儲蓄
Then 轉帳失敗
And 顯示「餘額不足」錯誤
```

### 轉換為 Gherkin

```gherkin
Feature: 銀行帳戶間轉帳
  As a 銀行客戶
  I want 在我的帳戶間轉移資金
  So that 我可以靈活管理我的財務

  # 規則 1：餘額驗證
  Rule: 來源帳戶必須有足夠的餘額

    Scenario: 餘額充足時成功轉帳
      Given 支票帳戶有 $1,000
      And 儲蓄帳戶有 $500
      When 我從支票帳戶轉 $200 到儲蓄帳戶
      Then 支票帳戶應該有 $800
      And 儲蓄帳戶應該有 $700

    Scenario: 餘額不足時轉帳失敗
      Given 支票帳戶有 $100
      And 儲蓄帳戶有 $500
      When 我嘗試從支票帳戶轉 $200 到儲蓄帳戶
      Then 轉帳應該失敗
      And 我應該看到錯誤訊息「餘額不足」
```

## 使用 Rule 組織場景

從 Cucumber 6.0 開始，支援 `Rule` 關鍵字來組織相關場景：

```gherkin
Feature: 帳戶利息計算

  Rule: 儲蓄帳戶每月計息

    Scenario: 標準儲蓄帳戶
      Given 我有一個標準儲蓄帳戶
      When 計算月利息
      Then 利率應該是 1% 年利率

    Scenario: 高級儲蓄帳戶
      Given 我有一個高級儲蓄帳戶
      When 計算月利息
      Then 利率應該是 2% 年利率

  Rule: 活期帳戶不計息

    Scenario: 活期帳戶
      Given 我有一個活期帳戶
      When 計算月利息
      Then 利息應該是 $0
```

## Scenario Outline（場景大綱）

用於參數化測試，減少重複：

```gherkin
Scenario Outline: 不同帳戶類型的利息計算
  Given 我有一個 <帳戶類型> 帳戶，餘額 $<餘額>
  When 計算月利息
  Then 利息應該是 $<利息>

  Examples:
    | 帳戶類型 | 餘額    | 利息  |
    | 活期     | 1000.00 | 0.00  |
    | 儲蓄     | 1000.00 | 0.83  |
    | 高級儲蓄 | 1000.00 | 1.67  |
```

## Background（背景）

所有場景共用的前置條件：

```gherkin
Feature: 帳戶管理

  Background:
    Given 我已經登入系統
    And 我有以下帳戶：
      | 類型 | 帳號   | 餘額    |
      | 活期 | 12345  | 500.00  |
      | 儲蓄 | 67890  | 1000.00 |

  Scenario: 查看帳戶餘額
    When 我查看帳戶清單
    Then 我應該看到 2 個帳戶

  Scenario: 查看特定帳戶
    When 我查看帳號 12345
    Then 餘額應該顯示 $500.00
```

## 資料表（Data Tables）

用於傳遞結構化資料：

```gherkin
Scenario: 建立多個帳戶
  Given 以下客戶帳戶：
    | 姓名   | 帳戶類型 | 初始餘額 |
    | 張三   | 儲蓄     | 1000     |
    | 李四   | 活期     | 500      |
    | 王五   | 高級儲蓄 | 5000     |
  When 系統處理帳戶建立
  Then 應該成功建立 3 個帳戶
```

## 最佳實踐

### ✅ 應該做的（Do）

1. **使用業務語言**
   ```gherkin
   # ✅ 好
   When 客戶從 ATM 提款 $100

   # ❌ 壞
   When 調用 withdrawCash() 方法，參數為 100
   ```

2. **保持場景獨立**
   - 每個場景應該可以獨立執行
   - 不依賴其他場景的狀態

3. **使用描述性名稱**
   ```gherkin
   # ✅ 好
   Scenario: 提款金額超過每日限額時應該失敗

   # ❌ 壞
   Scenario: 測試提款
   ```

4. **一個場景一個行為**
   - 每個場景只測試一個業務規則
   - 避免場景過長

5. **使用具體範例**
   ```gherkin
   # ✅ 好
   Given 帳戶餘額為 $1,000

   # ❌ 壞
   Given 帳戶有足夠的錢
   ```

### ❌ 避免做的（Don't）

1. **不要混合抽象層次**
   ```gherkin
   # ❌ 壞 - 混合業務語言和技術細節
   Given 使用者在登入頁面
   When 點擊 id="submit" 的按鈕
   Then 資料庫應該有一筆 session 記錄

   # ✅ 好 - 保持業務層次
   Given 使用者在登入頁面
   When 使用者提交登入表單
   Then 使用者應該成功登入
   ```

2. **不要在 When 中驗證**
   ```gherkin
   # ❌ 壞
   When 我轉帳 $100 且餘額變成 $900

   # ✅ 好
   When 我轉帳 $100
   Then 餘額應該是 $900
   ```

3. **不要過度使用 Scenario Outline**
   - 只在真正需要參數化時使用
   - 不要為了減少行數而犧牲可讀性

4. **不要寫過長的場景**
   - 如果場景超過 10 個步驟，考慮拆分
   - 使用 Background 提取共同步驟

## 組織結構建議

### 目錄結構

```
tests/
└── features/
    ├── authentication/
    │   ├── login.feature
    │   └── logout.feature
    ├── transfers/
    │   ├── internal_transfer.feature
    │   └── external_transfer.feature
    └── interest/
        └── interest_calculation.feature
```

### 檔案命名

- 使用小寫和底線：`internal_transfer.feature`
- 一個檔案一個功能
- 按業務領域組織

## 標籤（Tags）

用於組織和過濾測試：

```gherkin
@smoke @critical
Feature: 使用者登入

  @happy_path
  Scenario: 有效憑證登入
    Given 使用者存在於系統
    When 使用者輸入正確的帳號密碼
    Then 使用者應該成功登入

  @error_handling
  Scenario: 無效密碼
    Given 使用者存在於系統
    When 使用者輸入錯誤的密碼
    Then 登入應該失敗
    And 顯示「密碼錯誤」訊息

  @wip
  Scenario: 忘記密碼功能
    # 開發中...
```

執行特定標籤的測試：
```bash
# 只執行標記為 @smoke 的場景
godog --tags=@smoke

# 排除標記為 @wip 的場景
godog --tags='not @wip'
```

## 註解

使用 `#` 加入註解：

```gherkin
Feature: 帳戶管理
  # 這是功能的詳細說明
  # 可以有多行註解

  Scenario: 建立新帳戶
    # TODO: 需要確認最低開戶金額
    Given 客戶提供基本資料
    When 提交開戶申請
    Then 帳戶應該成功建立
```

## 多語言支援

Gherkin 支援多種語言關鍵字：

```gherkin
# language: zh-TW
功能: 銀行帳戶間轉帳
  作為 銀行客戶
  我想要 在我的帳戶間轉移資金
  所以 我可以靈活管理我的財務

  場景: 成功的轉帳
    假設 支票帳戶有 $1,000
    而且 儲蓄帳戶有 $500
    當 我從支票帳戶轉 $200 到儲蓄帳戶
    那麼 支票帳戶應該有 $800
    而且 儲蓄帳戶應該有 $700
```

## 從 Discovery 到 Formulation 的工作流程

### 步驟 1：完成 Example Mapping（Discovery 階段）

輸出：`specs/bank-transfer-discovery.md`

```markdown
### 📏 業務規則

#### 規則 1：餘額驗證

📝 **範例 1.1**：餘額充足
Given 支票帳戶有 $1,000
When 從支票轉 $200 到儲蓄
Then 支票帳戶剩 $800

📝 **範例 1.2**：餘額不足
Given 支票帳戶有 $100
When 嘗試轉 $200 到儲蓄
Then 轉帳失敗
```

### 步驟 2：轉換為 Gherkin（Formulation 階段）

創建：`features/transfers/internal_transfer.feature`

```gherkin
Feature: 銀行帳戶間轉帳
  As a 銀行客戶
  I want 在我的帳戶間轉移資金
  So that 我可以靈活管理我的財務

  Rule: 來源帳戶必須有足夠的餘額

    Scenario: 餘額充足時成功轉帳
      Given 支票帳戶有 $1,000
      And 儲蓄帳戶有 $500
      When 我從支票帳戶轉 $200 到儲蓄帳戶
      Then 支票帳戶應該有 $800
      And 儲蓄帳戶應該有 $700

    Scenario: 餘額不足時轉帳失敗
      Given 支票帳戶有 $100
      And 儲蓄帳戶有 $500
      When 我嘗試從支票帳戶轉 $200 到儲蓄帳戶
      Then 轉帳應該失敗
      And 我應該看到錯誤訊息「餘額不足」
```

## 常見錯誤

### 1. 技術細節洩漏

```gherkin
# ❌ 錯誤
Scenario: 儲存使用者
  Given 資料庫連線已建立
  When POST /api/users with JSON payload
  Then HTTP status code 應該是 201
  And 資料應該寫入 users 表

# ✅ 正確
Scenario: 註冊新使用者
  Given 我在註冊頁面
  When 我填寫註冊表單並提交
  Then 我應該看到「註冊成功」訊息
  And 我應該收到確認郵件
```

### 2. 場景過於抽象

```gherkin
# ❌ 錯誤
Scenario: 處理付款
  Given 系統運作正常
  When 執行付款流程
  Then 應該成功

# ✅ 正確
Scenario: 使用信用卡付款
  Given 購物車有總價 $50 的商品
  And 我選擇信用卡付款
  When 我輸入有效的信用卡資訊並確認
  Then 付款應該成功
  And 我應該收到訂單確認郵件
```

### 3. Given 中包含動作

```gherkin
# ❌ 錯誤
Given 我點擊登入按鈕

# ✅ 正確
Given 我在登入頁面
When 我點擊登入按鈕
```

## 工具支援

### Go - Godog

```bash
# 執行所有 feature
godog

# 執行特定 feature
godog features/transfers/

# 使用標籤過濾
godog --tags=@smoke

# 生成步驟定義模板
godog --format=steps
```

### 編輯器支援

- **VS Code**: Cucumber (Gherkin) Full Support
- **IntelliJ IDEA**: 內建 Gherkin 支援
- **Vim**: vim-cucumber

## 參考資源

- [Cucumber Gherkin Reference](https://cucumber.io/docs/gherkin/reference/)
- [Godog Documentation](https://github.com/cucumber/godog)
- [BDD in Action (Book)](https://www.manning.com/books/bdd-in-action-second-edition)

---

**記住：** Gherkin 是業務和技術之間的橋樑。寫 Gherkin 時，想像你是在向不懂程式的業務人員解釋系統行為！
