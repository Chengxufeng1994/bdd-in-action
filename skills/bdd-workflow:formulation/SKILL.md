---
name: bdd-workflow:formulation
description: |
  從 BDD 工作流程發現階段產生的使用者故事和 Example Mapping 結果，
  轉換成 Gherkin 格式的可執行規格（Given-When-Then 場景）。
allowed-tools: Read, Write, Glob, Bash
argument-hint: [discoveryFile]
version: 1.0.0
---

# BDD 工作流程制定技能

## 概述

此技能將 **BDD 工作流程發現階段**的輸出（存放在 `specs/[功能名稱]-discovery.md`）轉換成 **Gherkin 格式的可執行規格**。

### 輸入

- `specs/[功能名稱]-discovery.md` - 包含：
  - 使用者故事
  - 故事拆分結果
  - Example Mapping 結果（規則、範例、問題）

### 輸出

- `features/[功能名稱].feature` - Gherkin 場景文件

---

## 使用方式

```bash
# 方式一：指定 discovery 文件
/bdd-workflow-formulation specs/user-login-discovery.md

# 方式二：從 specs 目錄列出可用文件
/bdd-workflow-formulation
```

---

## 工作流程

### 階段一：讀取 Discovery 文檔

1. **自動掃描**
   - 如果未指定文件，掃描 `specs/` 目錄
   - 列出所有 `*-discovery.md` 文件
   - 讓用戶選擇要處理的文件

2. **解析內容**
   - 提取使用者故事
   - 提取拆分後的小故事（如果有）
   - 提取 Example Mapping 結果
     - 業務規則（📏）
     - 具體範例（📝）
     - 待解決問題（❓）

### 階段二：Gherkin 場景生成

**遵循 DFS 原則**：一次處理一個小故事，生成完整的 Feature 文件後再繼續下一個。

#### 2.1 創建 Feature 文件頭部

```gherkin
Feature: [功能名稱]

  [使用者故事描述]
  As a [角色]
  I want [需求]
  So that [價值]
```

#### 2.2 轉換 Example Mapping 為 Scenarios

**對每個業務規則：**

1. **規則名稱** → `Scenario` 或 `Rule` + `Scenario`
2. **範例** → `Given-When-Then` 步驟
3. **多個範例** → `Scenario Outline` + `Examples`

**轉換規則：**

```
📏 業務規則 1：有效憑證驗證
  └─ 📝 範例 1.1：成功登入
      Given 使用者 "john" 的密碼是 "password123"
      When 使用者以 "john" 和 "password123" 登入
      Then 登入應該成功
  └─ 📝 範例 1.2：密碼錯誤
      Given 使用者 "john" 的密碼是 "password123"
      When 使用者以 "john" 和 "wrongpassword" 登入
      Then 登入應該失敗
      And 顯示錯誤訊息 "帳號或密碼錯誤"

↓ 轉換為 ↓

Scenario: 成功登入
  Given 使用者 "john" 的密碼是 "password123"
  When 使用者以 "john" 和 "password123" 登入
  Then 登入應該成功

Scenario: 密碼錯誤導致登入失敗
  Given 使用者 "john" 的密碼是 "password123"
  When 使用者以 "john" 和 "wrongpassword" 登入
  Then 登入應該失敗
  And 顯示錯誤訊息 "帳號或密碼錯誤"
```

#### 2.3 使用 Scenario Outline（如適用）

當多個範例具有相同結構但不同數據時：

```gherkin
Scenario Outline: 驗證登入憑證
  Given 使用者 "<username>" 的密碼是 "<correct_password>"
  When 使用者以 "<username>" 和 "<input_password>" 登入
  Then 登入應該<result>

  Examples:
    | username | correct_password | input_password  | result |
    | john     | password123      | password123     | 成功   |
    | john     | password123      | wrongpassword   | 失敗   |
    | alice    | secret456        | secret456       | 成功   |
```

#### 2.4 使用 Rule（如適用）

當有多個相關規則時，使用 `Rule` 關鍵字組織：

```gherkin
Feature: 使用者登入

  As a 一般使用者
  I want 能夠登入系統
  So that 我可以存取受保護的內容

  Rule: 有效憑證驗證

    Scenario: 成功登入
      Given ...

    Scenario: 密碼錯誤
      Given ...

  Rule: 帳號狀態檢查

    Scenario: 帳號被鎖定
      Given ...
```

### 階段三：待解決問題處理

對於 Example Mapping 中的 ❓ 待解決問題：

1. **添加註解**
   ```gherkin
   # TODO: 確認 VIP 客戶的限額是否不同（負責人：產品經理）
   # TODO: 確認轉帳是即時還是延遲處理（負責人：技術架構師）
   ```

2. **添加 @pending 標籤**
   ```gherkin
   @pending @needs-clarification
   Scenario: VIP 客戶轉帳限額
     # 待產品經理確認 VIP 限額規則
     Given ...
   ```

### 階段四：儲存 Feature 文件

**目錄結構：**

```
項目根目錄/
├── specs/
│   └── user-login-discovery.md
└── features/
    └── user-login.feature
```

**命名規則：**
- Feature 文件名與 discovery 文件對應
- 使用 kebab-case
- 範例：`user-login-discovery.md` → `user-login.feature`

**儲存邏輯：**

1. 確保 `features/` 目錄存在
2. 生成完整的 Gherkin 內容
3. 使用 Write 工具儲存到 `features/[功能名稱].feature`
4. 向用戶確認文件路徑

---

## 實現邏輯

### 步驟 1：選擇 Discovery 文件

```
如果用戶提供了文件路徑：
  使用該文件

否則：
  1. 使用 Glob 掃描 specs/*-discovery.md
  2. 列出所有找到的文件（顯示 feature 名稱和狀態）
  3. 讓用戶選擇要處理的文件
```

### 步驟 2：解析 Discovery 文檔

```
使用 Read 工具讀取文件內容

2.1 讀取頂部元數據區塊（<!-- discovery ... -->）：
   - feature: 功能名稱（kebab-case）→ 推導輸出檔案名
   - file: 來源文件路徑（確認）
   - status: 確認為 complete

2.2 從 feature 欄位推導輸出路徑：
   output_file = "features/[feature].feature"

2.3 解析正文結構：
   - 提取 "## 功能描述"
   - 提取每個 "### Story X.X：[名稱]"
     - 使用者故事（As a / I want / So that）
     - Acceptance Criteria
     - Example Mapping 結果（業務規則、範例表格）
     - 準備度評估（只處理 🟢 綠燈或 🟡 黃燈的故事）
```

### 步驟 3：選擇要生成的 Story（DFS 原則）

```
如果有多個拆分後的故事：
  顯示故事列表（按優先級排序）
  讓用戶選擇要生成的 Story
  （建議從 MVP 的第一個開始）

否則：
  處理單一故事
```

### 步驟 4：生成 Gherkin 場景

```
4.1 創建 Feature 頭部
--------------------
Feature: [功能名稱]

  [使用者故事]
  As a [角色]
  I want [需求]
  So that [價值]

4.2 轉換業務規則和範例
----------------------
對每個業務規則：

  判斷使用 Scenario 還是 Scenario Outline：
    - 如果範例結構相似 → Scenario Outline
    - 否則 → 多個獨立 Scenario

  對每個範例：
    提取 Given-When-Then 步驟
    轉換成 Gherkin 格式

4.3 處理待解決問題
------------------
對每個問題：
  添加 TODO 註解
  或添加 @pending 標籤的 Scenario

4.4 添加標籤
------------
根據場景類型添加適當標籤：
  @smoke - 冒煙測試
  @regression - 回歸測試
  @pending - 待實現
  @needs-clarification - 需要釐清
```

### 步驟 5：儲存並確認

```
5.1 準備輸出目錄
-----------------
mkdir -p features

5.2 從元數據推導檔案名稱
--------------------------
從 discovery 文件頂部的 <!-- discovery ... --> 讀取 feature 欄位：
  feature: account-management → features/account-management.feature

若無元數據，則從檔案名移除 "-discovery" 後綴：
  account-management-discovery.md → account-management.feature

5.3 儲存文件（含元數據區塊）
------------------------------
Write(
  file_path: "features/[feature].feature",
  content:
    # formulation 元數據（Gherkin 註解格式）
    # formulation
    # feature: [feature]
    # source: specs/[feature]-discovery.md
    # file: features/[feature].feature
    # status: complete
    #

    Feature: ...
    [Gherkin 內容]
)

5.4 向用戶顯示完成摘要
------------------------
---
✅ Formulation 完成！

📄 文件位置：features/[feature].feature
📊 Scenarios：X 個（含 Y 個 @pending）
🔗 來源：specs/[feature]-discovery.md

▶️ 下一步：實現步驟定義（Automation 階段）
make test  ← 執行測試（預期 RED）
---
```

---

## Gherkin 最佳實踐

### ✅ Do（應該做）

1. **使用業務語言**
   - 避免技術術語
   - 使用領域專家能理解的詞彙
   - 保持場景可讀性

2. **保持 Scenario 獨立**
   - 每個 Scenario 應該可以獨立執行
   - 不依賴其他 Scenario 的執行順序
   - 使用 Background 處理共同前置條件

3. **明確的 Given-When-Then**
   - Given：設定初始狀態
   - When：執行動作（通常只有一個）
   - Then：驗證結果

4. **使用 Scenario Outline 避免重複**
   - 當多個場景結構相同時
   - 使用 Examples 表格提供不同數據

5. **適當使用標籤**
   ```gherkin
   @smoke @login
   Scenario: 成功登入
   ```

### ❌ Don't（避免做）

1. **不要包含實現細節**
   ```gherkin
   # ❌ 錯誤
   When 我點擊 id 為 "login-button" 的按鈕

   # ✅ 正確
   When 我點擊登入按鈕
   ```

2. **不要寫太長的 Scenario**
   - 一個 Scenario 應該測試一個行為
   - 超過 10 個步驟可能太複雜

3. **不要在 When 中驗證**
   ```gherkin
   # ❌ 錯誤
   When 我登入並且成功

   # ✅ 正確
   When 我登入
   Then 登入應該成功
   ```

4. **不要過度使用 And**
   - 考慮是否應該拆分成多個 Scenario

---

## 範例輸出

### 輸入：Discovery 文檔

```markdown
# 功能：使用者登入登出

## 原始使用者故事

In order to 存取受保護的內容或資源，並確保操作安全性
As a 一般使用者
I want 能夠登入和登出系統

## Example Mapping 結果

### 📏 業務規則

#### 規則 1：有效憑證驗證

使用者必須提供存在於系統中的帳號，並且密碼正確匹配。

📝 **範例 1.1**：成功登入
```
Given 使用者 "john" 的密碼是 "password123"
When 使用者以 "john" 和 "password123" 登入
Then 登入應該成功
```

📝 **範例 1.2**：密碼錯誤
```
Given 使用者 "john" 的密碼是 "password123"
When 使用者以 "john" 和 "wrongpassword" 登入
Then 登入應該失敗
And 顯示錯誤訊息 "帳號或密碼錯誤"
```

#### 規則 2：登入失敗處理

登入在以下情況會失敗：
- 帳號不存在於系統中
- 密碼與帳號不匹配

📝 **範例 2.1**：帳號不存在
```
Given 系統中不存在使用者 "unknown"
When 使用者以 "unknown" 和 "anypassword" 登入
Then 登入應該失敗
And 顯示錯誤訊息 "帳號或密碼錯誤"
```

### ❓ 待解決問題

1. 是否需要雙重認證？ - 負責人：安全團隊
2. 連續失敗多次後是否鎖定帳號？ - 負責人：產品經理
```

### 輸出：Feature 文件

```gherkin
Feature: 使用者登入登出

  In order to 存取受保護的內容或資源，並確保操作安全性
  As a 一般使用者
  I want 能夠登入和登出系統

  Rule: 有效憑證驗證

    使用者必須提供存在於系統中的帳號，並且密碼正確匹配。

    Scenario: 成功登入
      Given 使用者 "john" 的密碼是 "password123"
      When 使用者以 "john" 和 "password123" 登入
      Then 登入應該成功

    Scenario: 密碼錯誤導致登入失敗
      Given 使用者 "john" 的密碼是 "password123"
      When 使用者以 "john" 和 "wrongpassword" 登入
      Then 登入應該失敗
      And 顯示錯誤訊息 "帳號或密碼錯誤"

  Rule: 登入失敗處理

    登入在以下情況會失敗：
    - 帳號不存在於系統中
    - 密碼與帳號不匹配

    Scenario: 帳號不存在導致登入失敗
      Given 系統中不存在使用者 "unknown"
      When 使用者以 "unknown" 和 "anypassword" 登入
      Then 登入應該失敗
      And 顯示錯誤訊息 "帳號或密碼錯誤"

  # TODO: 確認是否需要雙重認證（負責人：安全團隊）
  # TODO: 確認連續失敗多次後是否鎖定帳號（負責人：產品經理）

  @pending @needs-clarification
  Scenario: 連續登入失敗後鎖定帳號
    # 待產品經理確認鎖定規則
    Given 使用者 "john" 的密碼是 "password123"
    When 使用者連續 3 次以錯誤密碼登入
    Then 帳號應該被鎖定
```

---

## 參考文檔

- [Gherkin 編寫指南](./gherkin-guide.md) - **完整的 Gherkin 語法和最佳實踐指南**
- [Gherkin 語法參考](https://cucumber.io/docs/gherkin/reference/)
- [BDD 工作流程發現](../bdd-workflow-discovery/SKILL.md)
- [Godog 文檔](https://github.com/cucumber/godog)

---

## 版本歷史

- **1.0.0** - 初始版本
  - 從 Discovery 文檔生成 Gherkin 場景
  - 支援 Scenario 和 Scenario Outline
  - 支援 Rule 組織
  - 處理待解決問題
  - 遵循 DFS 原則

---

## 與其他 Skills 的關係

```
bdd-workflow-discovery
  ↓
  產生 specs/[功能名稱]-discovery.md
  ↓
bdd-workflow-formulation ← 此 Skill
  ↓
  產生 features/[功能名稱].feature
  ↓
（下一步：實現步驟定義和執行測試）
```

---

**注意**：此技能專注於 BDD 流程的「制定」階段（Formulation）。完成後，開發者應該：

1. 檢視生成的 Feature 文件
2. 實現步驟定義（Step Definitions）
3. 執行 Godog 測試
4. 實現業務邏輯
5. 重複 Red-Green-Refactor 循環
