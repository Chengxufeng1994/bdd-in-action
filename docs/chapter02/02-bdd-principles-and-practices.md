# Chapter 2.3: BDD 原則與實踐 - 從業務目標到程式碼

## 概述

Behaviour-Driven Development (BDD) 不僅是一種測試方法，更是一種將業務價值、協作溝通與技術實現緊密結合的開發方式。BDD 提供了一個清晰的層次結構，讓業務需求能夠有條理地流動到最終的程式碼實現。

### BDD 的層次結構

```
┌─────────────────────────────────────────┐
│     1. Business Goal (業務目標)          │  ← 為什麼要做？ (WHY)
│        商業價值與策略目標                  │     利益相關者
└─────────────────────────────────────────┘
              ↓
┌─────────────────────────────────────────┐
│     2. Features (功能特性)               │  ← 做什麼？ (WHAT)
│        系統提供的能力                      │     業務分析師
└─────────────────────────────────────────┘
              ↓
┌─────────────────────────────────────────┐
│     3. Examples (具體範例)               │  ← 具體表現如何？
│        具體的使用場景                      │     Three Amigos
└─────────────────────────────────────────┘
              ↓
┌─────────────────────────────────────────┐
│     4. Executable (可執行規格)           │  ← 形式化的場景
│        Gherkin 規格                       │     業務 + 測試
└─────────────────────────────────────────┘
              ↓
┌─────────────────────────────────────────┐
│     5. Specifications (驗收規格)         │  ← 自動化驗收測試
│        Step Definitions                  │     測試 + 開發
└─────────────────────────────────────────┘
              ↓
┌─────────────────────────────────────────┐
│     6. Low-Level Specs (底層規格)        │  ← 單元測試
│        TDD 測試                           │     開發人員
└─────────────────────────────────────────┘
              ↓
┌─────────────────────────────────────────┐
│     7. Application Code (應用程式碼)      │  ← 怎麼做？ (HOW)
│        實際實現                           │     開發人員
└─────────────────────────────────────────┘
```

### 每個層次的作用

| 層次 | 主要參與者 | 產出物 | 目的 |
|------|-----------|--------|------|
| Business Goal | 利益相關者、產品負責人 | 業務目標說明 | 定義價值和方向 |
| Features | 業務分析師、產品負責人 | User Stories, Feature 描述 | 定義系統能力 |
| Examples | Three Amigos | 具體場景、Example Map | 建立共同理解 |
| Executable | 業務分析師、測試人員 | .feature 文件 (Gherkin) | 可執行的規格 |
| Specifications | 測試人員、開發人員 | Step Definitions | 自動化驗收測試 |
| Low-Level Specs | 開發人員 | 單元測試 | 驅動設計和實現 |
| Application Code | 開發人員 | 產品代碼 | 交付業務價值 |

### 核心原則

1. **Outside-In Development** - 從業務價值開始，由外而內開發
2. **Traceability** - 每一行代碼都能追溯到業務價值
3. **Living Documentation** - 規格即文檔，始終保持最新
4. **Collaboration** - 不同角色在不同層次協作
5. **Executable Specifications** - 所有規格都可執行和驗證

接下來，我們將使用**帳戶轉賬功能**作為貫穿所有層次的實例，展示每個層次如何運作以及它們之間如何連接。

---

## 1. Business Goal (業務目標)

### 什麼是業務目標？

業務目標是組織希望達成的商業成果，它回答了「**為什麼**」要建立這個功能的問題。良好的業務目標應該：

- 清楚說明業務價值
- 可衡量成功與否
- 與公司戰略對齊
- 使用利益相關者的語言

### 業務目標的特徵

#### ✅ 好的業務目標

```
增加客戶自助服務能力，減少分行櫃檯交易量，
降低營運成本並提升客戶滿意度。

目標：
- 在 6 個月內減少 30% 的櫃檯轉賬交易
- 提升客戶滿意度分數 15 個百分點
- 降低每筆轉賬的處理成本 50%
```

#### ❌ 不好的業務目標

```
實現帳戶轉賬功能
```

### 實例：帳戶轉賬的業務目標

**商業背景**：
銀行發現大量客戶到分行僅是為了進行簡單的帳戶間轉賬，這造成：

- 分行櫃檯排隊時間過長
- 人力成本居高不下
- 客戶體驗不佳

**業務目標**：

```
【策略目標】數位轉型 - 提升自助服務比例

【業務目標】提供線上帳戶轉賬功能
讓客戶能夠自主完成帳戶間資金調撥，無需到分行辦理，
從而提升客戶便利性並降低營運成本。

【成功指標】
1. 功能上線後 3 個月內，30% 的轉賬交易透過線上完成
2. 客戶滿意度調查中，「轉賬便利性」項目提升至 4.5/5.0
3. 每筆轉賬的平均處理成本從 $5 降至 $0.50
4. 減少 20% 的分行櫃檯人力需求

【目標客群】
- 現有數位銀行用戶（約 100,000 人）
- 每月至少進行一次轉賬的活躍用戶（約 30,000 人）

【風險考量】
- 必須符合金融監管要求
- 確保交易安全性和可追蹤性
- 防範詐欺和洗錢風險
```

### 業務目標驅動決策

業務目標會影響下游所有決策：

| 業務目標 | 影響的決策 |
|---------|----------|
| 降低營運成本 | → 優先實現最常見的轉賬場景，避免過度設計 |
| 提升便利性 | → 介面要簡單直觀，減少輸入步驟 |
| 確保安全性 | → 實施交易限額、二次驗證等安全機制 |
| 符合監管 | → 保留完整交易記錄，實施反洗錢檢查 |

---

## 2. Features (功能特性)

### 什麼是 Feature？

Feature 是系統提供的能力，它回答「**做什麼**」的問題。Feature 將業務目標轉化為具體的系統功能，並以用戶故事的形式表達。

### Feature 的結構

在 BDD 中，Feature 通常使用以下格式：

```gherkin
Feature: [功能名稱]
  As a [角色/用戶類型]
  I want [想要的能力]
  So that [獲得的價值]
```

這個格式確保每個 Feature 都：

- **有明確的用戶** (As a...)
- **有具體的能力** (I want...)
- **有清晰的價值** (So that...)

### 實例：從業務目標到 Feature

**業務目標** → **Features 分解**

```
業務目標：提供線上帳戶轉賬功能
    ↓
衍生出的 Features：
├─ Feature 1: 帳戶間轉賬
├─ Feature 2: 轉賬限額控制
├─ Feature 3: 轉賬歷史記錄
├─ Feature 4: 轉賬通知
└─ Feature 5: 預約轉賬
```

### 核心 Feature: 帳戶間轉賬

#### User Story 結構

一個完整的 User Story 應該包含以下元素：

**基本格式**:

```
As a [角色]
I want [功能]
So that [價值]
```

**完整 User Story 範例**:

```
Story ID: US-001
Title: 帳戶間轉賬

As a 銀行客戶
I want to 在我的帳戶之間轉移資金
So that 我可以靈活管理我的財務而無需到分行辦理

Story Points: 8
Priority: High
Sprint: Sprint 3

Description:
客戶需要能夠在自己名下的不同帳戶之間轉移資金。
這是最常見的銀行操作之一，目前必須到分行櫃檯辦理，
造成客戶不便和銀行營運成本高。

Business Value:
- 減少分行櫃檯交易量 30%
- 提升客戶滿意度
- 降低每筆交易處理成本從 $5 到 $0.50

Personas:
1. Chris - 年輕專業人士，頻繁使用網路銀行
2. Maria - 小企業主，需要在多個帳戶間管理資金
3. David - 退休人士，偶爾需要調整帳戶資金配置

Assumptions:
- 用戶已經完成身份驗證
- 用戶擁有至少兩個帳戶
- 系統可以即時處理轉賬

Dependencies:
- 用戶認證系統 (已完成)
- 帳戶管理系統 (已完成)
- 交易日誌系統 (需同步開發)

Risks:
- 並發轉賬可能導致餘額不一致
- 需要確保交易原子性
- 需符合金融監管法規

Definition of Ready:
✓ User Story 符合 INVEST 原則
✓ 業務規則已明確定義
✓ 驗收標準已與 PO 確認
✓ 技術可行性已評估
✓ UI/UX 設計草圖已準備

Definition of Done:
✓ 所有驗收測試通過
✓ 代碼審查完成
✓ 單元測試覆蓋率 > 80%
✓ 整合測試通過
✓ 文檔已更新
✓ 已部署到測試環境
✓ PO 已驗收
```

#### INVEST 原則檢查

好的 User Story 應該符合 INVEST 原則：

| 原則 | 說明 | 本 Story 如何符合 |
|------|------|------------------|
| **I**ndependent | 獨立的 | 不依賴其他未完成的 Story |
| **N**egotiable | 可協商的 | 實現細節可以討論（如限額設定） |
| **V**aluable | 有價值的 | 明確的業務價值：降低成本、提升滿意度 |
| **E**stimable | 可估算的 | 團隊可以估算為 8 點 |
| **S**mall | 小的 | 可在一個 Sprint 內完成 |
| **T**estable | 可測試的 | 有明確的驗收標準 |

#### Feature 文件

**文件**: `features/money_transfer.feature`

```gherkin
Feature: 帳戶間轉賬
  As a 銀行客戶
  I want to 在我的帳戶之間轉移資金
  So that 我可以靈活管理我的財務而無需到分行辦理

  Background:
    這個功能是數位銀行策略的核心部分，
    目標是讓客戶能夠自主完成最常見的銀行操作。

  Business Rules:
  - 轉賬金額必須大於零
  - 來源帳戶必須有足夠餘額
  - 不能轉賬到相同帳戶
  - 所有轉賬必須記錄以供審計
  - 轉賬必須具有原子性（ACID）

  Acceptance Criteria:
  - 成功轉賬後，兩個帳戶餘額正確更新
  - 失敗時提供清楚的錯誤訊息
  - 交易具有原子性（全部成功或全部失敗）
  - 交易完成時間 < 2 秒
  - 系統記錄完整的審計日誌

  Technical Constraints:
  - 使用資料庫事務確保一致性
  - 支援每秒 100 筆並發轉賬
  - 符合 PCI-DSS 安全標準
```

### Feature 與業務目標的對應

```
業務目標：降低 30% 櫃檯交易
    ↓
Feature: 帳戶間轉賬
    ↓
預期影響：
- 客戶可自行完成轉賬 → 減少到分行次數
- 24/7 可用 → 提升便利性
- 即時完成 → 改善體驗
    ↓
可衡量指標：
- 線上轉賬交易量
- 櫃檯轉賬交易減少量
- 客戶滿意度提升
```

### 多個相關 Features

```gherkin
Feature: 轉賬限額控制
  As a 風險管理人員
  I want to 設定不同客戶等級的轉賬限額
  So that 可以降低詐欺風險並符合監管要求

Feature: 轉賬通知
  As a 銀行客戶
  I want to 轉賬完成後收到通知
  So that 我可以確認交易成功並保持帳務透明

Feature: 轉賬歷史記錄
  As a 銀行客戶
  I want to 查看我的轉賬歷史
  So that 我可以追蹤我的財務活動
```

---

## 3. Examples (具體範例)

### 為什麼需要具體範例？

抽象的需求容易產生誤解，具體的範例建立共同理解。這就是 **Specification by Example** 的核心思想。

### 抽象 vs 具體

| 抽象描述 | 具體範例 |
|---------|---------|
| "帳戶餘額必須足夠" | "Chris 的帳戶有 $1,000，轉賬 $300 應該成功" |
| "系統應驗證金額" | "轉賬 $0 應該失敗，顯示 '金額必須大於零'" |
| "錯誤訊息要清楚" | "當 Chris 轉賬 $1,500 但只有 $1,000 時，<br>應顯示 '餘額不足，當前餘額 $1,000'" |

### Example Mapping 工作坊

**參與者**: Three Amigos（業務、開發、測試）

**工具**: 四種顏色的卡片

```
┌─────────────────────────────┐
│      Feature Card           │  黃色 - 要實現的 Feature
│   帳戶間轉賬功能               │
└─────────────────────────────┘
          ↓
┌─────────────────────────────┐
│      Rule Cards             │  藍色 - 業務規則
│ • 餘額必須足夠                 │
│ • 金額必須大於零               │
│ • 不能轉到相同帳戶             │
└─────────────────────────────┘
          ↓
┌─────────────────────────────┐
│    Example Cards            │  綠色 - 具體範例
│ • 餘額 $1000 轉 $300 → 成功   │
│ • 餘額 $100 轉 $150 → 失敗    │
│ • 轉 $0 → 失敗                │
└─────────────────────────────┘
          ↓
┌─────────────────────────────┐
│    Question Cards           │  紅色 - 未解問題
│ • 並發轉賬如何處理？           │
│ • 最大金額限制？               │
└─────────────────────────────┘
```

### 實例：轉賬功能的範例探索

#### Rule 1: 餘額驗證

**業務規則**: 來源帳戶必須有足夠餘額

**正面範例** (應該成功):

```
範例 1.1: 餘額充足的轉賬
Given:
  - Chris 的支票帳戶 (ACC001) 有 $1,000
  - Lisa 的儲蓄帳戶 (ACC002) 有 $500
When:
  - Chris 轉賬 $300 從 ACC001 到 ACC002
Then:
  - 轉賬成功
  - ACC001 餘額變為 $700
  - ACC002 餘額變為 $800
  - Chris 收到確認訊息 "轉賬成功"

範例 1.2: 全額轉賬（邊界情況）
Given:
  - Chris 的帳戶有 $500
When:
  - Chris 轉賬 $500
Then:
  - 轉賬成功
  - Chris 的帳戶餘額變為 $0
```

**反面範例** (應該失敗):

```
範例 1.3: 餘額不足
Given:
  - Chris 的帳戶有 $100
When:
  - Chris 嘗試轉賬 $150
Then:
  - 轉賬失敗
  - 錯誤訊息: "餘額不足，當前餘額 $100"
  - 兩個帳戶餘額都不改變
  - 不記錄任何交易

範例 1.4: 零餘額嘗試轉賬
Given:
  - Chris 的帳戶餘額為 $0
When:
  - Chris 嘗試轉賬 $50
Then:
  - 轉賬失敗
  - 錯誤訊息: "帳戶餘額為零，無法轉賬"
```

#### Rule 2: 金額驗證

**業務規則**: 轉賬金額必須大於零

**反面範例**:

```
範例 2.1: 零金額轉賬
Given:
  - Chris 的帳戶有 $1,000
When:
  - Chris 嘗試轉賬 $0
Then:
  - 轉賬失敗
  - 錯誤訊息: "轉賬金額必須大於零"

範例 2.2: 負金額轉賬
Given:
  - Chris 的帳戶有 $1,000
When:
  - Chris 嘗試轉賬 $-50
Then:
  - 轉賬失敗
  - 錯誤訊息: "轉賬金額不能為負數"

範例 2.3: 最小有效金額（邊界情況）
Given:
  - Chris 的帳戶有 $100
When:
  - Chris 轉賬 $0.01
Then:
  - 轉賬成功
  - 扣款 $0.01
```

#### Rule 3: 帳戶驗證

**業務規則**: 不能轉賬到相同的帳戶

**反面範例**:

```
範例 3.1: 轉到相同帳戶
Given:
  - Chris 的帳戶編號為 ACC001
When:
  - Chris 嘗試從 ACC001 轉賬到 ACC001
Then:
  - 轉賬失敗
  - 錯誤訊息: "不能轉賬到相同的帳戶"

範例 3.2: 目標帳戶不存在
Given:
  - Chris 的帳戶為 ACC001
When:
  - Chris 嘗試轉賬到不存在的帳戶 ACC999
Then:
  - 轉賬失敗
  - 錯誤訊息: "目標帳戶不存在"
```

### 範例的價值

**建立共同理解**:

- 業務：驗證這些場景符合業務規則
- 開發：理解邊界條件和錯誤處理
- 測試：識別測試場景和邊緣案例

**發現隱藏需求**:
通過討論範例，團隊發現：

- 需要處理並發轉賬的競爭條件
- 需要明確最小和最大轉賬金額
- 需要考慮帳戶凍結狀態
- 需要審計日誌記錄

---

## 4. Executable (可執行規格)

### 從範例到可執行規格

具體範例雖然清楚，但仍是非結構化的文字。**Gherkin** 提供了一種標準化的語法，讓範例變得可執行。

### Gherkin 的優勢

1. **結構化** - 明確的 Given-When-Then 結構
2. **可讀** - 業務人員可以閱讀和驗證
3. **可執行** - 可以直接轉換為自動化測試
4. **工具支援** - 有豐富的工具生態系統

### Gherkin 語法概覽

```gherkin
Feature: 功能名稱
  功能的業務描述

  Background:
    Given 所有場景共享的前置條件

  Scenario: 場景名稱
    Given 前置條件 (設置上下文)
    And 更多前置條件
    When 觸發的動作 (執行行為)
    And 更多動作
    Then 預期的結果 (驗證輸出)
    And 更多結果驗證

  Scenario Outline: 參數化場景
    Given 某個 <參數>
    When 執行 <動作>
    Then 得到 <結果>

    Examples:
      | 參數 | 動作 | 結果 |
      | 值1  | 值2  | 值3  |
```

### 實例：完整的可執行規格

**文件**: `features/money_transfer.feature`

```gherkin
Feature: 帳戶間轉賬
  As a 銀行客戶
  I want to 在我的帳戶之間轉移資金
  So that 我可以靈活管理我的財務而無需到分行辦理

  Background:
    Given 系統中存在以下帳戶:
      | 帳戶編號 | 戶名   | 類型   | 餘額     |
      | ACC001   | Chris  | 支票   | 1000.00  |
      | ACC002   | Lisa   | 儲蓄   | 500.00   |
      | ACC003   | Jordan | 支票   | 0.00     |

  # 主要成功場景
  Scenario: 成功轉賬 - 餘額充足
    Given Chris 已經登入系統
    When Chris 從帳戶 "ACC001" 轉賬 $300.00 到帳戶 "ACC002"
    Then 轉賬應該成功
    And 帳戶 "ACC001" 的餘額應該是 $700.00
    And 帳戶 "ACC002" 的餘額應該是 $800.00
    And Chris 應該看到確認訊息 "轉賬成功"
    And 系統應該記錄一筆轉賬交易

  # 餘額不足場景
  Scenario: 轉賬失敗 - 餘額不足
    Given Chris 已經登入系統
    When Chris 從帳戶 "ACC001" 轉賬 $1500.00 到帳戶 "ACC002"
    Then 轉賬應該失敗
    And Chris 應該看到錯誤訊息 "餘額不足，當前餘額 $1000.00"
    And 帳戶 "ACC001" 的餘額應該保持 $1000.00
    And 帳戶 "ACC002" 的餘額應該保持 $500.00
    And 系統不應該記錄轉賬交易

  # 邊界情況
  Scenario: 全額轉賬
    Given Chris 已經登入系統
    When Chris 從帳戶 "ACC001" 轉賬 $1000.00 到帳戶 "ACC003"
    Then 轉賬應該成功
    And 帳戶 "ACC001" 的餘額應該是 $0.00
    And 帳戶 "ACC003" 的餘額應該是 $1000.00
    And Jordan 應該收到轉賬通知

  # 參數化測試 - 金額驗證
  Scenario Outline: 金額驗證規則
    Given Chris 已經登入系統
    When Chris 從帳戶 "ACC001" 轉賬 <金額> 到帳戶 "ACC002"
    Then 轉賬應該<結果>
    And Chris 應該看到<訊息類型> "<訊息內容>"

    Examples: 無效金額
      | 金額     | 結果 | 訊息類型 | 訊息內容             |
      | $0.00    | 失敗 | 錯誤訊息 | 轉賬金額必須大於零   |
      | $-50.00  | 失敗 | 錯誤訊息 | 轉賬金額不能為負數   |

    Examples: 有效金額
      | 金額     | 結果 | 訊息類型 | 訊息內容       |
      | $0.01    | 成功 | 確認訊息 | 轉賬成功       |
      | $500.00  | 成功 | 確認訊息 | 轉賬成功       |

  # 帳戶驗證場景
  Scenario: 轉賬失敗 - 相同帳戶
    Given Chris 已經登入系統
    When Chris 從帳戶 "ACC001" 轉賬 $100.00 到帳戶 "ACC001"
    Then 轉賬應該失敗
    And Chris 應該看到錯誤訊息 "不能轉賬到相同的帳戶"

  Scenario: 轉賬失敗 - 目標帳戶不存在
    Given Chris 已經登入系統
    When Chris 從帳戶 "ACC001" 轉賬 $100.00 到帳戶 "ACC999"
    Then 轉賬應該失敗
    And Chris 應該看到錯誤訊息 "目標帳戶不存在"

  # 並發處理
  @concurrent
  Scenario: 並發轉賬處理
    Given Chris 已經登入系統
    And Chris 的帳戶 "ACC001" 有 $1000.00
    When Chris 同時發起兩筆轉賬:
      | 目標帳戶 | 金額    |
      | ACC002   | $600.00 |
      | ACC003   | $600.00 |
    Then 只有一筆轉賬應該成功
    And 另一筆應該失敗並顯示 "餘額不足"
    And 帳戶 "ACC001" 的最終餘額應該是 $400.00
```

### Gherkin 最佳實踐

#### ✅ DO

1. **使用業務語言**

   ```gherkin
   # 好
   When Chris 轉賬 $100 到 Lisa

   # 不好
   When the user calls TransferService.execute(100, "ACC001", "ACC002")
   ```

2. **描述行為而非實現**

   ```gherkin
   # 好
   Then Chris 應該看到確認訊息

   # 不好
   Then the HTTP response code should be 200
   And the JSON should contain "success": true
   ```

3. **保持場景獨立**

   ```gherkin
   # 好 - 每個場景完整獨立
   Scenario: 第一次轉賬
     Given Chris 的帳戶有 $1000
     When Chris 轉賬 $100
     Then 餘額應該是 $900

   Scenario: 第二次轉賬
     Given Chris 的帳戶有 $1000
     When Chris 轉賬 $200
     Then 餘額應該是 $800
   ```

#### ❌ DON'T

1. **不要使用技術術語**
2. **不要描述 UI 細節**
3. **不要讓場景相互依賴**
4. **不要在一個場景中測試多個規則**

---

## 5. Specifications (驗收規格)

### Step Definitions - 將 Gherkin 變為可執行代碼

Step Definitions 是連接 Gherkin 場景和實際代碼的橋樑。每個 Gherkin 步驟都會映射到一個 Go 函數。

### 測試架構

```
features/
├── money_transfer.feature          # Gherkin 規格
└── steps/
    ├── transfer_test.go            # Test runner
    ├── transfer_steps.go           # Step definitions
    └── transfer_context.go         # Test context
```

### 實例：Go + Godog 實現

#### 5.1 測試入口

**文件**: `features/steps/transfer_test.go`

```go
package steps

import (
 "context"
 "testing"

 "github.com/cucumber/godog"
)

func TestFeatures(t *testing.T) {
 suite := godog.TestSuite{
  ScenarioInitializer: InitializeScenario,
  Options: &godog.Options{
   Format:   "pretty",        // 輸出格式
   Paths:    []string{"../"},  // feature 文件路徑
   TestingT: t,
  },
 }

 if suite.Run() != 0 {
  t.Fatal("non-zero status returned, failed to run feature tests")
 }
}
```

#### 5.2 測試上下文

**文件**: `features/steps/transfer_context.go`

```go
package steps

import (
 "your-project/internal/domain"
 "your-project/internal/service"

 "github.com/shopspring/decimal"
)

// transferContext 保存場景執行期間的狀態
type transferContext struct {
 // 測試資料
 accounts map[string]*domain.Account

 // 當前用戶
 currentUser string

 // 轉賬結果
 transferResult *service.TransferResult
 lastError      error

 // 服務
 transferService *service.TransferService
 accountRepo     *MockAccountRepository
 transferRepo    *MockTransferRepository
}

func newTransferContext() *transferContext {
 accountRepo := NewMockAccountRepository()
 transferRepo := NewMockTransferRepository()

 return &transferContext{
  accounts:        make(map[string]*domain.Account),
  accountRepo:     accountRepo,
  transferRepo:    transferRepo,
  transferService: service.NewTransferService(accountRepo, transferRepo),
 }
}

func (tc *transferContext) cleanup() {
 tc.accounts = make(map[string]*domain.Account)
 tc.currentUser = ""
 tc.transferResult = nil
 tc.lastError = nil
}
```

#### 5.3 Step Definitions

**文件**: `features/steps/transfer_steps.go`

```go
package steps

import (
 "context"
 "fmt"

 "github.com/cucumber/godog"
 "github.com/shopspring/decimal"
)

func InitializeScenario(ctx *godog.ScenarioContext) {
 tc := newTransferContext()

 // Background steps
 ctx.Step(`^系統中存在以下帳戶:$`, tc.thereAreAccountsInSystem)

 // Given steps - 設置前置條件
 ctx.Step(`^(\w+) 已經登入系統$`, tc.userIsLoggedIn)
 ctx.Step(`^(\w+) 的帳戶 "([^"]*)" 有 \$(\d+(?:\.\d+)?)$`, tc.accountHasBalance)

 // When steps - 執行動作
 ctx.Step(`^(\w+) 從帳戶 "([^"]*)" 轉賬 \$(\d+(?:\.\d+)?) 到帳戶 "([^"]*)"$`,
  tc.userTransfersMoney)

 ctx.Step(`^(\w+) 同時發起兩筆轉賬:$`, tc.userMakesConcurrentTransfers)

 // Then steps - 驗證結果
 ctx.Step(`^轉賬應該成功$`, tc.transferShouldSucceed)
 ctx.Step(`^轉賬應該失敗$`, tc.transferShouldFail)
 ctx.Step(`^帳戶 "([^"]*)" 的餘額應該是 \$(\d+(?:\.\d+)?)$`,
  tc.accountBalanceShouldBe)
 ctx.Step(`^帳戶 "([^"]*)" 的餘額應該保持 \$(\d+(?:\.\d+)?)$`,
  tc.accountBalanceShouldRemain)
 ctx.Step(`^(\w+) 應該看到確認訊息 "([^"]*)"$`,
  tc.userShouldSeeConfirmation)
 ctx.Step(`^(\w+) 應該看到錯誤訊息 "([^"]*)"$`,
  tc.userShouldSeeError)
 ctx.Step(`^系統應該記錄一筆轉賬交易$`, tc.transferShouldBeRecorded)
 ctx.Step(`^系統不應該記錄轉賬交易$`, tc.noTransferShouldBeRecorded)
 ctx.Step(`^(\w+) 應該收到轉賬通知$`, tc.userShouldReceiveNotification)

 // 場景後清理
 ctx.After(func(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
  tc.cleanup()
  return ctx, nil
 })
}

// Background step implementation
func (tc *transferContext) thereAreAccountsInSystem(table *godog.Table) error {
 // 跳過表頭
 for _, row := range table.Rows[1:] {
  accountID := row.Cells[0].Value
  owner := row.Cells[1].Value
  accountType := row.Cells[2].Value
  balance := row.Cells[3].Value

  account := &domain.Account{
   ID:      accountID,
   Owner:   owner,
   Type:    accountType,
   Balance: decimal.RequireFromString(balance),
  }

  tc.accounts[accountID] = account
  tc.accountRepo.Save(account)
 }
 return nil
}

// Given step implementations
func (tc *transferContext) userIsLoggedIn(username string) error {
 tc.currentUser = username
 return nil
}

func (tc *transferContext) accountHasBalance(username, accountID string, balance float64) error {
 account, exists := tc.accounts[accountID]
 if !exists {
  return fmt.Errorf("account %s not found", accountID)
 }
 account.Balance = decimal.NewFromFloat(balance)
 tc.accountRepo.Update(account)
 return nil
}

// When step implementations
func (tc *transferContext) userTransfersMoney(
 username, fromID string,
 amount float64,
 toID string,
) error {
 ctx := context.Background()
 transferAmount := decimal.NewFromFloat(amount)

 result, err := tc.transferService.ExecuteTransfer(
  ctx,
  fromID,
  toID,
  transferAmount,
 )

 tc.transferResult = result
 tc.lastError = err

 // 更新本地帳戶狀態以供驗證
 if err == nil {
  fromAccount, _ := tc.accountRepo.FindByID(ctx, fromID)
  toAccount, _ := tc.accountRepo.FindByID(ctx, toID)
  tc.accounts[fromID] = fromAccount
  tc.accounts[toID] = toAccount
 }

 return nil
}

// Then step implementations
func (tc *transferContext) transferShouldSucceed() error {
 if tc.lastError != nil {
  return fmt.Errorf("expected transfer to succeed but got error: %v", tc.lastError)
 }
 if tc.transferResult == nil {
  return fmt.Errorf("transfer result is nil")
 }
 if tc.transferResult.Status != domain.TransferStatusCompleted {
  return fmt.Errorf("transfer status is not completed: %s", tc.transferResult.Status)
 }
 return nil
}

func (tc *transferContext) transferShouldFail() error {
 if tc.lastError == nil {
  return fmt.Errorf("expected transfer to fail but it succeeded")
 }
 return nil
}

func (tc *transferContext) accountBalanceShouldBe(accountID string, expectedBalance float64) error {
 account, exists := tc.accounts[accountID]
 if !exists {
  return fmt.Errorf("account %s not found", accountID)
 }

 expected := decimal.NewFromFloat(expectedBalance)
 if !account.Balance.Equal(expected) {
  return fmt.Errorf(
   "account %s: expected balance $%.2f but got $%.2f",
   accountID,
   expectedBalance,
   account.Balance.InexactFloat64(),
  )
 }

 return nil
}

func (tc *transferContext) accountBalanceShouldRemain(accountID string, expectedBalance float64) error {
 // 與 accountBalanceShouldBe 相同，但語義上表示"保持不變"
 return tc.accountBalanceShouldBe(accountID, expectedBalance)
}

func (tc *transferContext) userShouldSeeConfirmation(username, expectedMessage string) error {
 if tc.transferResult == nil {
  return fmt.Errorf("no transfer result available")
 }
 if tc.transferResult.Message != expectedMessage {
  return fmt.Errorf(
   "expected message %q but got %q",
   expectedMessage,
   tc.transferResult.Message,
  )
 }
 return nil
}

func (tc *transferContext) userShouldSeeError(username, expectedError string) error {
 if tc.lastError == nil {
  return fmt.Errorf("expected error but got none")
 }

 actualError := tc.lastError.Error()
 if actualError != expectedError {
  return fmt.Errorf(
   "expected error %q but got %q",
   expectedError,
   actualError,
  )
 }
 return nil
}

func (tc *transferContext) transferShouldBeRecorded() error {
 if tc.transferResult == nil || tc.transferResult.ID == "" {
  return fmt.Errorf("transfer was not recorded")
 }

 // 驗證資料庫中存在該交易
 transfer, err := tc.transferRepo.FindByID(tc.transferResult.ID)
 if err != nil {
  return fmt.Errorf("transfer not found in repository: %w", err)
 }
 if transfer.Status != domain.TransferStatusCompleted {
  return fmt.Errorf("transfer status is not completed")
 }

 return nil
}

func (tc *transferContext) noTransferShouldBeRecorded() error {
 // 驗證沒有新的交易記錄
 transfers := tc.transferRepo.GetAll()
 if len(transfers) > 0 {
  return fmt.Errorf("expected no transfers but found %d", len(transfers))
 }
 return nil
}

func (tc *transferContext) userShouldReceiveNotification(username string) error {
 // 在實際系統中，這裡會檢查通知系統
 // 在測試中，我們可以檢查 mock 通知服務
 return nil
}
```

### 運行測試

```bash
# 運行所有 BDD 場景
go test -v ./features/...

# 運行特定標籤的場景
go test -v ./features/... -godog.tags="@transfer"

# 生成測試報告
go test -v ./features/... -godog.format=cucumber:report.json
```

### 測試輸出

```
Feature: 帳戶間轉賬

  Scenario: 成功轉賬 - 餘額充足                                # features/money_transfer.feature:10
    Given 系統中存在以下帳戶:                                    # transfer_steps.go:23
      | 帳戶編號 | 戶名   | 類型   | 餘額     |
      | ACC001   | Chris  | 支票   | 1000.00  |
      | ACC002   | Lisa   | 儲蓄   | 500.00   |
      | ACC003   | Jordan | 支票   | 0.00     |
    And Chris 已經登入系統                                      # transfer_steps.go:45
    When Chris 從帳戶 "ACC001" 轉賬 $300.00 到帳戶 "ACC002"      # transfer_steps.go:52
    Then 轉賬應該成功                                          # transfer_steps.go:78
    And 帳戶 "ACC001" 的餘額應該是 $700.00                       # transfer_steps.go:91
    And 帳戶 "ACC002" 的餘額應該是 $800.00                       # transfer_steps.go:91
    And Chris 應該看到確認訊息 "轉賬成功"                         # transfer_steps.go:106

7 scenarios (7 passed)
34 steps (34 passed)
256.78ms
```

---

## 6. Low-Level Specifications (底層規格)

### 單元測試與驗收測試的關係

```
┌─────────────────────────────────────────┐
│    Acceptance Tests (驗收測試)           │  ← 驗證業務行為
│    Features + Step Definitions          │     從外部視角
└─────────────────────────────────────────┘
              ↓ 引導實現
┌─────────────────────────────────────────┐
│    Unit Tests (單元測試)                 │  ← 驅動設計
│    Domain + Service Layer               │     從內部視角
└─────────────────────────────────────────┘
```

### TDD - 測試驅動開發

在實現 Step Definitions 所需的代碼時，使用 TDD 方法：

**Red → Green → Refactor**

1. **Red** - 寫一個失敗的測試
2. **Green** - 寫最小代碼讓測試通過
3. **Refactor** - 改善代碼質量

### 實例：Domain 層單元測試

#### 6.1 Account 實體測試

**文件**: `internal/domain/account_test.go`

```go
package domain_test

import (
 "testing"

 "your-project/internal/domain"

 "github.com/shopspring/decimal"
 "github.com/stretchr/testify/assert"
 "github.com/stretchr/testify/require"
)

func TestAccount_CanTransfer(t *testing.T) {
 tests := []struct {
  name           string
  accountBalance string
  transferAmount string
  expected       bool
 }{
  {
   name:           "sufficient balance",
   accountBalance: "1000.00",
   transferAmount: "500.00",
   expected:       true,
  },
  {
   name:           "exact balance",
   accountBalance: "500.00",
   transferAmount: "500.00",
   expected:       true,
  },
  {
   name:           "insufficient balance",
   accountBalance: "100.00",
   transferAmount: "150.00",
   expected:       false,
  },
  {
   name:           "zero balance",
   accountBalance: "0.00",
   transferAmount: "10.00",
   expected:       false,
  },
 }

 for _, tt := range tests {
  t.Run(tt.name, func(t *testing.T) {
   // Arrange
   account := &domain.Account{
    ID:      "ACC001",
    Owner:   "Chris",
    Balance: decimal.RequireFromString(tt.accountBalance),
   }
   amount := decimal.RequireFromString(tt.transferAmount)

   // Act
   result := account.CanTransfer(amount)

   // Assert
   assert.Equal(t, tt.expected, result)
  })
 }
}

func TestAccount_Debit(t *testing.T) {
 tests := []struct {
  name            string
  initialBalance  string
  debitAmount     string
  expectedBalance string
  expectError     bool
  expectedError   error
 }{
  {
   name:            "successful debit",
   initialBalance:  "1000.00",
   debitAmount:     "300.00",
   expectedBalance: "700.00",
   expectError:     false,
  },
  {
   name:            "debit full amount",
   initialBalance:  "500.00",
   debitAmount:     "500.00",
   expectedBalance: "0.00",
   expectError:     false,
  },
  {
   name:           "insufficient balance",
   initialBalance: "100.00",
   debitAmount:    "150.00",
   expectError:    true,
   expectedError:  domain.ErrInsufficientBalance,
  },
 }

 for _, tt := range tests {
  t.Run(tt.name, func(t *testing.T) {
   // Arrange
   account := &domain.Account{
    ID:      "ACC001",
    Balance: decimal.RequireFromString(tt.initialBalance),
   }
   amount := decimal.RequireFromString(tt.debitAmount)

   // Act
   err := account.Debit(amount)

   // Assert
   if tt.expectError {
    require.Error(t, err)
    assert.ErrorIs(t, err, tt.expectedError)
   } else {
    require.NoError(t, err)
    expected := decimal.RequireFromString(tt.expectedBalance)
    assert.True(t, account.Balance.Equal(expected),
     "expected balance %s but got %s", expected, account.Balance)
   }
  })
 }
}

func TestAccount_Credit(t *testing.T) {
 // Arrange
 account := &domain.Account{
  ID:      "ACC001",
  Balance: decimal.NewFromInt(500),
 }
 creditAmount := decimal.NewFromInt(300)

 // Act
 account.Credit(creditAmount)

 // Assert
 expectedBalance := decimal.NewFromInt(800)
 assert.True(t, account.Balance.Equal(expectedBalance))
}
```

#### 6.2 Transfer 實體測試

**文件**: `internal/domain/transfer_test.go`

```go
package domain_test

import (
 "testing"

 "your-project/internal/domain"

 "github.com/shopspring/decimal"
 "github.com/stretchr/testify/assert"
 "github.com/stretchr/testify/require"
)

func TestTransfer_Validate(t *testing.T) {
 tests := []struct {
  name          string
  transfer      *domain.Transfer
  expectedError error
 }{
  {
   name: "valid transfer",
   transfer: &domain.Transfer{
    FromAccountID: "ACC001",
    ToAccountID:   "ACC002",
    Amount:        decimal.NewFromInt(100),
   },
   expectedError: nil,
  },
  {
   name: "negative amount",
   transfer: &domain.Transfer{
    FromAccountID: "ACC001",
    ToAccountID:   "ACC002",
    Amount:        decimal.NewFromInt(-50),
   },
   expectedError: domain.ErrNegativeAmount,
  },
  {
   name: "zero amount",
   transfer: &domain.Transfer{
    FromAccountID: "ACC001",
    ToAccountID:   "ACC002",
    Amount:        decimal.Zero,
   },
   expectedError: domain.ErrInvalidAmount,
  },
  {
   name: "same account",
   transfer: &domain.Transfer{
    FromAccountID: "ACC001",
    ToAccountID:   "ACC001",
    Amount:        decimal.NewFromInt(100),
   },
   expectedError: domain.ErrSameAccount,
  },
  {
   name: "empty from account",
   transfer: &domain.Transfer{
    FromAccountID: "",
    ToAccountID:   "ACC002",
    Amount:        decimal.NewFromInt(100),
   },
   expectedError: domain.ErrAccountNotFound,
  },
  {
   name: "empty to account",
   transfer: &domain.Transfer{
    FromAccountID: "ACC001",
    ToAccountID:   "",
    Amount:        decimal.NewFromInt(100),
   },
   expectedError: domain.ErrAccountNotFound,
  },
 }

 for _, tt := range tests {
  t.Run(tt.name, func(t *testing.T) {
   // Act
   err := tt.transfer.Validate()

   // Assert
   if tt.expectedError != nil {
    require.Error(t, err)
    assert.ErrorIs(t, err, tt.expectedError)
   } else {
    assert.NoError(t, err)
   }
  })
 }
}
```

#### 6.3 TransferService 測試

**文件**: `internal/service/transfer_service_test.go`

```go
package service_test

import (
 "context"
 "testing"

 "your-project/internal/domain"
 "your-project/internal/service"
 "your-project/internal/service/mocks"

 "github.com/shopspring/decimal"
 "github.com/stretchr/testify/assert"
 "github.com/stretchr/testify/mock"
 "github.com/stretchr/testify/require"
)

func TestTransferService_ExecuteTransfer_Success(t *testing.T) {
 // Arrange
 ctx := context.Background()

 fromAccount := &domain.Account{
  ID:      "ACC001",
  Owner:   "Chris",
  Balance: decimal.NewFromInt(1000),
 }
 toAccount := &domain.Account{
  ID:      "ACC002",
  Owner:   "Lisa",
  Balance: decimal.NewFromInt(500),
 }
 transferAmount := decimal.NewFromInt(300)

 // Mock repositories
 accountRepo := new(mocks.MockAccountRepository)
 transferRepo := new(mocks.MockTransferRepository)

 // 設置期望的調用
 accountRepo.On("FindByID", ctx, "ACC001").Return(fromAccount, nil)
 accountRepo.On("FindByID", ctx, "ACC002").Return(toAccount, nil)
 accountRepo.On("BeginTx", ctx).Return(&mocks.MockTransaction{}, nil)
 accountRepo.On("UpdateInTx", ctx, mock.Anything, mock.Anything).Return(nil)
 transferRepo.On("SaveInTx", ctx, mock.Anything, mock.Anything).Return(nil)

 // 創建 service
 svc := service.NewTransferService(accountRepo, transferRepo)

 // Act
 result, err := svc.ExecuteTransfer(ctx, "ACC001", "ACC002", transferAmount)

 // Assert
 require.NoError(t, err)
 assert.NotNil(t, result)
 assert.Equal(t, domain.TransferStatusCompleted, result.Status)
 assert.Equal(t, "轉賬成功", result.Message)

 // 驗證帳戶餘額
 assert.True(t, fromAccount.Balance.Equal(decimal.NewFromInt(700)))
 assert.True(t, toAccount.Balance.Equal(decimal.NewFromInt(800)))

 // 驗證 mock 期望
 accountRepo.AssertExpectations(t)
 transferRepo.AssertExpectations(t)
}

func TestTransferService_ExecuteTransfer_InsufficientBalance(t *testing.T) {
 // Arrange
 ctx := context.Background()

 fromAccount := &domain.Account{
  ID:      "ACC001",
  Balance: decimal.NewFromInt(100),
 }
 toAccount := &domain.Account{
  ID:      "ACC002",
  Balance: decimal.NewFromInt(500),
 }
 transferAmount := decimal.NewFromInt(150)

 accountRepo := new(mocks.MockAccountRepository)
 transferRepo := new(mocks.MockTransferRepository)

 accountRepo.On("FindByID", ctx, "ACC001").Return(fromAccount, nil)
 accountRepo.On("FindByID", ctx, "ACC002").Return(toAccount, nil)
 accountRepo.On("BeginTx", ctx).Return(&mocks.MockTransaction{}, nil)

 svc := service.NewTransferService(accountRepo, transferRepo)

 // Act
 result, err := svc.ExecuteTransfer(ctx, "ACC001", "ACC002", transferAmount)

 // Assert
 require.Error(t, err)
 assert.ErrorIs(t, err, domain.ErrInsufficientBalance)
 assert.NotNil(t, result)
 assert.Equal(t, domain.TransferStatusFailed, result.Status)

 // 餘額應該保持不變
 assert.True(t, fromAccount.Balance.Equal(decimal.NewFromInt(100)))
 assert.True(t, toAccount.Balance.Equal(decimal.NewFromInt(500)))
}
```

### Table-Driven Tests

Go 的慣例是使用 table-driven tests，這與 BDD 的 Scenario Outline 概念一致：

```go
func TestTransferValidation(t *testing.T) {
 tests := []struct {
  name    string
  amount  decimal.Decimal
  fromID  string
  toID    string
  wantErr error
 }{
  {"valid", decimal.NewFromInt(100), "ACC001", "ACC002", nil},
  {"zero amount", decimal.Zero, "ACC001", "ACC002", domain.ErrInvalidAmount},
  {"negative", decimal.NewFromInt(-50), "ACC001", "ACC002", domain.ErrNegativeAmount},
  {"same account", decimal.NewFromInt(100), "ACC001", "ACC001", domain.ErrSameAccount},
 }

 for _, tt := range tests {
  t.Run(tt.name, func(t *testing.T) {
   transfer := &domain.Transfer{
    Amount:        tt.amount,
    FromAccountID: tt.fromID,
    ToAccountID:   tt.toID,
   }

   err := transfer.Validate()

   if tt.wantErr != nil {
    assert.ErrorIs(t, err, tt.wantErr)
   } else {
    assert.NoError(t, err)
   }
  })
 }
}
```

---

## 7. Application Code (應用程式碼)

### 從規格到實現

有了驗收測試和單元測試的引導，我們可以實現實際的應用程式碼。這是 **Outside-In TDD** 的典型流程。

### 架構層次

```
┌─────────────────────────────────────────┐
│         API/HTTP Layer                  │  ← 對外接口
└─────────────────────────────────────────┘
              ↓
┌─────────────────────────────────────────┐
│         Service Layer                   │  ← 業務邏輯編排
└─────────────────────────────────────────┘
              ↓
┌─────────────────────────────────────────┐
│         Domain Layer                    │  ← 核心業務規則
└─────────────────────────────────────────┘
              ↓
┌─────────────────────────────────────────┐
│         Repository Layer                │  ← 數據持久化
└─────────────────────────────────────────┘
```

### 7.1 Domain Layer - 核心業務實體

**文件**: `internal/domain/account.go`

```go
package domain

import (
 "github.com/shopspring/decimal"
)

// Account 表示銀行帳戶
type Account struct {
 ID      string
 Owner   string
 Type    string // "checking", "savings", etc.
 Balance decimal.Decimal
}

// CanTransfer 檢查帳戶是否有足夠餘額進行轉賬
func (a *Account) CanTransfer(amount decimal.Decimal) bool {
 return a.Balance.GreaterThanOrEqual(amount)
}

// Debit 從帳戶扣款
func (a *Account) Debit(amount decimal.Decimal) error {
 if !a.CanTransfer(amount) {
  return ErrInsufficientBalance
 }
 a.Balance = a.Balance.Sub(amount)
 return nil
}

// Credit 向帳戶存款
func (a *Account) Credit(amount decimal.Decimal) {
 a.Balance = a.Balance.Add(amount)
}
```

**文件**: `internal/domain/transfer.go`

```go
package domain

import (
 "errors"
 "time"

 "github.com/shopspring/decimal"
)

type TransferStatus string

const (
 TransferStatusPending   TransferStatus = "pending"
 TransferStatusCompleted TransferStatus = "completed"
 TransferStatusFailed    TransferStatus = "failed"
)

// 業務錯誤定義
var (
 ErrInsufficientBalance = errors.New("餘額不足")
 ErrInvalidAmount       = errors.New("轉賬金額必須大於零")
 ErrNegativeAmount      = errors.New("轉賬金額不能為負數")
 ErrSameAccount         = errors.New("不能轉賬到相同的帳戶")
 ErrAccountNotFound     = errors.New("帳戶不存在")
)

// Transfer 表示一筆轉賬交易
type Transfer struct {
 ID            string
 FromAccountID string
 ToAccountID   string
 Amount        decimal.Decimal
 Status        TransferStatus
 Message       string
 CreatedAt     time.Time
 CompletedAt   *time.Time
}

// Validate 驗證轉賬請求的業務規則
func (t *Transfer) Validate() error {
 // 規則 1: 金額驗證
 if t.Amount.IsNegative() {
  return ErrNegativeAmount
 }
 if t.Amount.IsZero() {
  return ErrInvalidAmount
 }

 // 規則 2: 帳戶驗證
 if t.FromAccountID == t.ToAccountID {
  return ErrSameAccount
 }
 if t.FromAccountID == "" || t.ToAccountID == "" {
  return ErrAccountNotFound
 }

 return nil
}

// Complete 標記轉賬為完成
func (t *Transfer) Complete(message string) {
 now := time.Now()
 t.Status = TransferStatusCompleted
 t.Message = message
 t.CompletedAt = &now
}

// Fail 標記轉賬為失敗
func (t *Transfer) Fail(message string) {
 t.Status = TransferStatusFailed
 t.Message = message
}
```

### 7.2 Repository Interfaces

**文件**: `internal/domain/repository.go`

```go
package domain

import "context"

// AccountRepository 定義帳戶資料存取接口
type AccountRepository interface {
 FindByID(ctx context.Context, id string) (*Account, error)
 Save(ctx context.Context, account *Account) error
 Update(ctx context.Context, account *Account) error
 BeginTx(ctx context.Context) (Transaction, error)
 UpdateInTx(ctx context.Context, tx Transaction, account *Account) error
}

// TransferRepository 定義轉賬記錄存取接口
type TransferRepository interface {
 Save(ctx context.Context, transfer *Transfer) error
 FindByID(ctx context.Context, id string) (*Transfer, error)
 SaveInTx(ctx context.Context, tx Transaction, transfer *Transfer) error
}

// Transaction 表示資料庫事務
type Transaction interface {
 Commit() error
 Rollback() error
}
```

### 7.3 Service Layer - 業務邏輯編排

**文件**: `internal/service/transfer_service.go`

```go
package service

import (
 "context"
 "fmt"
 "time"

 "your-project/internal/domain"

 "github.com/google/uuid"
 "github.com/shopspring/decimal"
)

// TransferService 處理轉賬業務邏輯
type TransferService struct {
 accountRepo  domain.AccountRepository
 transferRepo domain.TransferRepository
}

// NewTransferService 創建轉賬服務實例
func NewTransferService(
 accountRepo domain.AccountRepository,
 transferRepo domain.TransferRepository,
) *TransferService {
 return &TransferService{
  accountRepo:  accountRepo,
  transferRepo: transferRepo,
 }
}

// TransferResult 表示轉賬操作的結果
type TransferResult struct {
 ID      string
 Status  domain.TransferStatus
 Message string
}

// ExecuteTransfer 執行轉賬操作
// 這是 BDD 場景中 "When Chris 從帳戶 ACC001 轉賬 $300 到帳戶 ACC002" 的實現
func (s *TransferService) ExecuteTransfer(
 ctx context.Context,
 fromAccountID, toAccountID string,
 amount decimal.Decimal,
) (*TransferResult, error) {
 // 1. 創建轉賬對象
 transfer := &domain.Transfer{
  ID:            uuid.New().String(),
  FromAccountID: fromAccountID,
  ToAccountID:   toAccountID,
  Amount:        amount,
  Status:        domain.TransferStatusPending,
  CreatedAt:     time.Now(),
 }

 // 2. 驗證業務規則
 if err := transfer.Validate(); err != nil {
  transfer.Fail(err.Error())
  return s.buildResult(transfer), err
 }

 // 3. 加載帳戶
 fromAccount, err := s.accountRepo.FindByID(ctx, fromAccountID)
 if err != nil {
  transfer.Fail("來源帳戶不存在")
  return s.buildResult(transfer), domain.ErrAccountNotFound
 }

 toAccount, err := s.accountRepo.FindByID(ctx, toAccountID)
 if err != nil {
  transfer.Fail("目標帳戶不存在")
  return s.buildResult(transfer), fmt.Errorf("目標%w", domain.ErrAccountNotFound)
 }

 // 4. 在事務中執行轉賬
 if err := s.executeInTransaction(ctx, transfer, fromAccount, toAccount); err != nil {
  transfer.Fail(err.Error())
  return s.buildResult(transfer), err
 }

 // 5. 轉賬成功
 transfer.Complete("轉賬成功")
 return s.buildResult(transfer), nil
}

// executeInTransaction 在數據庫事務中執行轉賬
func (s *TransferService) executeInTransaction(
 ctx context.Context,
 transfer *domain.Transfer,
 from, to *domain.Account,
) error {
 // 開始事務
 tx, err := s.accountRepo.BeginTx(ctx)
 if err != nil {
  return fmt.Errorf("無法開始事務: %w", err)
 }
 defer tx.Rollback() // 確保異常時回滾

 // 扣款 - 實現 "Then 帳戶 ACC001 的餘額應該是 $700"
 if err := from.Debit(transfer.Amount); err != nil {
  return err
 }

 // 入款 - 實現 "And 帳戶 ACC002 的餘額應該是 $800"
 to.Credit(transfer.Amount)

 // 更新來源帳戶
 if err := s.accountRepo.UpdateInTx(ctx, tx, from); err != nil {
  return fmt.Errorf("無法更新來源帳戶: %w", err)
 }

 // 更新目標帳戶
 if err := s.accountRepo.UpdateInTx(ctx, tx, to); err != nil {
  return fmt.Errorf("無法更新目標帳戶: %w", err)
 }

 // 保存轉賬記錄 - 實現 "And 系統應該記錄一筆轉賬交易"
 if err := s.transferRepo.SaveInTx(ctx, tx, transfer); err != nil {
  return fmt.Errorf("無法保存轉賬記錄: %w", err)
 }

 // 提交事務
 if err := tx.Commit(); err != nil {
  return fmt.Errorf("無法提交事務: %w", err)
 }

 return nil
}

func (s *TransferService) buildResult(transfer *domain.Transfer) *TransferResult {
 return &TransferResult{
  ID:      transfer.ID,
  Status:  transfer.Status,
  Message: transfer.Message,
 }
}

// GetTransferHistory 查詢轉賬歷史（支持另一個 Feature）
func (s *TransferService) GetTransferHistory(
 ctx context.Context,
 accountID string,
 limit int,
) ([]*domain.Transfer, error) {
 // 實現查詢邏輯...
 return nil, nil
}
```

### 7.4 HTTP Handler (API Layer)

**文件**: `internal/api/transfer_handler.go`

```go
package api

import (
 "encoding/json"
 "net/http"

 "your-project/internal/service"

 "github.com/shopspring/decimal"
)

type TransferHandler struct {
 transferService *service.TransferService
}

func NewTransferHandler(svc *service.TransferService) *TransferHandler {
 return &TransferHandler{
  transferService: svc,
 }
}

// TransferRequest 轉賬請求
type TransferRequest struct {
 FromAccountID string  `json:"from_account_id"`
 ToAccountID   string  `json:"to_account_id"`
 Amount        float64 `json:"amount"`
}

// TransferResponse 轉賬響應
type TransferResponse struct {
 Success bool   `json:"success"`
 Message string `json:"message"`
 Data    *struct {
  TransferID string  `json:"transfer_id"`
  Status     string  `json:"status"`
  Amount     float64 `json:"amount"`
 } `json:"data,omitempty"`
}

// HandleTransfer 處理轉賬 HTTP 請求
func (h *TransferHandler) HandleTransfer(w http.ResponseWriter, r *http.Request) {
 ctx := r.Context()

 // 解析請求
 var req TransferRequest
 if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
  h.respondError(w, http.StatusBadRequest, "無效的請求格式")
  return
 }

 // 執行轉賬
 amount := decimal.NewFromFloat(req.Amount)
 result, err := h.transferService.ExecuteTransfer(
  ctx,
  req.FromAccountID,
  req.ToAccountID,
  amount,
 )

 // 處理響應
 if err != nil {
  h.respondError(w, http.StatusBadRequest, err.Error())
  return
 }

 resp := &TransferResponse{
  Success: true,
  Message: result.Message,
  Data: &struct {
   TransferID string  `json:"transfer_id"`
   Status     string  `json:"status"`
   Amount     float64 `json:"amount"`
  }{
   TransferID: result.ID,
   Status:     string(result.Status),
   Amount:     req.Amount,
  },
 }

 w.Header().Set("Content-Type", "application/json")
 w.WriteHeader(http.StatusOK)
 json.NewEncoder(w).Encode(resp)
}

func (h *TransferHandler) respondError(w http.ResponseWriter, code int, message string) {
 resp := &TransferResponse{
  Success: false,
  Message: message,
 }
 w.Header().Set("Content-Type", "application/json")
 w.WriteHeader(code)
 json.NewEncoder(w).Encode(resp)
}
```

### 代碼追溯到規格

每段代碼都可以追溯到上層規格：

```
業務目標: 提升客戶便利性
    ↓
Feature: 帳戶間轉賬
    ↓
Scenario: 成功轉賬 - 餘額充足
    Given Chris 的帳戶有 $1,000
    When Chris 轉賬 $300 到 Lisa
    Then 轉賬成功
    And Chris 的帳戶餘額是 $700
    And Lisa 的帳戶餘額是 $800
    ↓
Step Definition: userTransfersMoney()
    ↓
Service: TransferService.ExecuteTransfer()
    ↓
Domain: Account.Debit(), Account.Credit()
    ↓
實現代碼:
    func (a *Account) Debit(amount decimal.Decimal) error {
        if !a.CanTransfer(amount) {
            return ErrInsufficientBalance
        }
        a.Balance = a.Balance.Sub(amount)
        return nil
    }
```

---

## 層次間的關係

### 向下流動：需求到實現

```
1. Business Goal (業務目標)
   "降低營運成本，提升客戶滿意度"
        ↓ 定義價值
2. Features (功能特性)
   "帳戶間轉賬功能"
        ↓ 描述能力
3. Examples (具體範例)
   "Chris 轉 $300: 成功"
   "Chris 轉 $1500: 失敗 - 餘額不足"
        ↓ 形式化
4. Executable (可執行規格)
   Given-When-Then Gherkin 場景
        ↓ 自動化
5. Specifications (驗收規格)
   Step Definitions (Go + godog)
        ↓ 驅動實現
6. Low-Level Specs (底層規格)
   單元測試 (domain + service)
        ↓ 引導設計
7. Application Code (應用程式碼)
   實際實現代碼
```

### 向上追溯：代碼到價值

當有人問「為什麼要這樣實現」時，可以向上追溯：

```
代碼: Account.CanTransfer()
    ↑ 被...測試
單元測試: TestAccount_CanTransfer
    ↑ 支持...
Step Definition: transferShouldFail()
    ↑ 實現...
Gherkin Scenario: "轉賬失敗 - 餘額不足"
    ↑ 驗證...
Example: "餘額 $100 轉 $150 → 失敗"
    ↑ 說明...
Business Rule: "轉賬金額不能超過餘額"
    ↑ 支持...
Feature: "帳戶間轉賬"
    ↑ 實現...
Business Goal: "降低營運成本 30%"
```

### 協作點：不同角色的參與

| 層次 | 誰參與 | 協作活動 |
|------|--------|---------|
| Business Goal | 利益相關者、產品負責人 | 戰略規劃會議 |
| Features | 產品負責人、業務分析師 | 需求分析、用戶故事撰寫 |
| Examples | Three Amigos | Example Mapping 工作坊 |
| Executable | 業務分析師、測試人員 | Gherkin 場景撰寫和審查 |
| Specifications | 測試人員、開發人員 | Step Definition 實現和配對 |
| Low-Level Specs | 開發人員 | TDD 結對編程 |
| Application Code | 開發人員 | 編碼和代碼審查 |

---

## 實踐指南

### 如何在團隊中應用

#### 1. 建立 BDD 工作流程

**週期開始**:

```
Sprint Planning
    ↓
選擇 User Stories
    ↓
Three Amigos Meeting
    ↓
Example Mapping
    ↓
編寫 Gherkin 場景
```

**開發期間**:

```
實現 Step Definitions
    ↓
Run BDD Tests (Red)
    ↓
TDD Implementation
    ↓
Run BDD Tests (Green)
    ↓
Refactor
```

**週期結束**:

```
Demo BDD Scenarios
    ↓
Review Living Documentation
    ↓
Retrospective
```

#### 2. 團隊協作模式

**Three Amigos 會議** (每個 Story 30-60 分鐘):

- **參與者**: 業務代表、開發人員、測試人員
- **產出**: Example Map、初步 Gherkin 場景
- **時機**: Story 開始開發前

**Example Mapping 工作坊**:

```
準備階段 (5分鐘):
- 寫下 Feature 在黃色卡片上

探索階段 (20-30分鐘):
- 討論業務規則（藍色卡片）
- 為每個規則提供範例（綠色卡片）
- 記錄未解問題（紅色卡片）

決策階段 (10分鐘):
- 範例是否足夠？
- 問題能否快速解決？
- Story 是否太大需要拆分？
```

#### 3. 工具鏈

**BDD 框架**:

- **Go**: godog
- **Ruby**: cucumber, rspec
- **Java**: cucumber-jvm, JBehave
- **JavaScript**: cucumber-js, jest
- **.NET**: SpecFlow

**報告工具**:

- cucumber-html-reporter
- allure
- serenity

**整合**:

```yaml
# CI/CD Pipeline
stages:
  - lint
  - unit-test
  - bdd-test
  - integration-test
  - deploy

bdd-test:
  script:
    - go test -v ./features/...
    - generate-report.sh
  artifacts:
    reports:
      junit: report.xml
    paths:
      - report.html
```

### 常見陷阱與解決方案

#### ❌ 陷阱 1: 過於技術化的場景

**問題**:

```gherkin
When I POST to /api/transfer with JSON payload
Then the HTTP status code should be 200
And the response Content-Type should be application/json
```

**解決**:

```gherkin
When Chris 轉賬 $300 到 Lisa
Then 轉賬應該成功
And Chris 應該看到確認訊息
```

#### ❌ 陷阱 2: 場景相互依賴

**問題**:

```gherkin
Scenario: Create account
  When I create account ACC001

Scenario: Transfer money
  When I transfer $100 from ACC001  # 依賴前一個場景
```

**解決**:

```gherkin
Background:
  Given account ACC001 exists with balance $1000

Scenario: Transfer money
  When I transfer $100 from ACC001 to ACC002
  Then transfer succeeds
```

#### ❌ 陷阱 3: 過度使用 Scenario Outline

**問題**:

```gherkin
Scenario Outline: Complex validation
  Given <precondition1> and <precondition2>
  When <action1> and <action2>
  Then <result1> and <result2> and <result3>

  Examples:
    # 20 rows with 10 columns...
```

**解決**:

- 拆分成多個簡單的 Scenario
- 每個 Scenario 測試一個業務規則
- 只在明顯重複時使用 Outline

#### ❌ 陷阱 4: 測試實現細節

**問題**:

```gherkin
Then the TransferService.execute method should be called
And the database transaction should be committed
```

**解決**:

```gherkin
Then the transfer should be completed
And the account balances should be updated
```

---

## 總結

### 核心原則回顧

BDD 的七個層次提供了一個完整的框架，確保：

1. **業務對齊** - 每個功能都追溯到業務價值
2. **協作溝通** - 不同角色在適當層次參與
3. **可追溯性** - 從目標到代碼，從代碼到目標
4. **可執行性** - 規格即測試，測試即文檔
5. **品質保證** - 多層次的驗證機制

### 關鍵記憶點

```
Business Goal  → WHY    → 利益相關者的語言
Features       → WHAT   → 用戶故事
Examples       → HOW    → 具體場景
Executable     → SPEC   → Gherkin
Specifications → TEST   → Step Definitions
Low-Level Spec → UNIT   → TDD
Application    → CODE   → Implementation
```

### 與 Chapter 1 工作流程的關係

**Chapter 1** 著重於**過程** (Process):

- 5個階段的工作流程
- 時間順序的活動
- 團隊如何協作

**Chapter 2.3** 著重於**結構** (Structure):

- 7個層次的規格
- 價值到實現的層次
- 不同層次的關注點

**兩者結合**:

```
Chapter 1: 何時做什麼 (When to do what)
Chapter 2.3: 在哪個層次做 (At which level to work)

實踐中：
在 Three Amigos 會議階段 (Chapter 1, 階段1)
↓
工作於 Examples 層次 (Chapter 2.3, 層次3)
↓
產出 Example Map 和初步場景
↓
接著進入 Executable 層次 (Chapter 2.3, 層次4)
編寫正式的 Gherkin 規格
```

### 下一步

繼續學習：

- **Chapter 3**: 高級 Gherkin 寫作技巧
- **Chapter 4**: 複雜領域的 BDD 建模
- **Chapter 5**: BDD 與微服務架構
- **Chapter 6**: 測試策略與維護

---

## 參考資源

### 書籍

- "BDD in Action" - John Ferguson Smart
- "Specification by Example" - Gojko Adzic
- "The Cucumber Book" - Matt Wynne & Aslak Hellesøy
- "Growing Object-Oriented Software, Guided by Tests" - Steve Freeman & Nat Pryce

### 線上資源

- [Cucumber BDD Guide](https://cucumber.io/docs/bdd/)
- [Example Mapping](https://cucumber.io/blog/bdd/example-mapping-introduction/)
- [Gherkin Reference](https://cucumber.io/docs/gherkin/reference/)
- [Godog Framework](https://github.com/cucumber/godog)
- [Dan North's Blog](https://dannorth.net/)

### 相關章節

- [Chapter 2.1: Introducing BDD](./01-introducing-bdd.md)
- [Chapter 1: BDD 核心工作流程](../chapter01/03-bdd-core-workflow.md)

---

**準備好實踐了嗎？** 選擇一個小功能，從業務目標開始，逐層向下實現，體驗完整的 BDD 流程！
