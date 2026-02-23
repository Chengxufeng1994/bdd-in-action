# Chapter 06 Research：三種技術的比較分析

## 概述

本文件比較分析 Chapter 06 的三個核心技術，釐清各自的定位、差異與協作關係。

---

## 1. 定位比較

| 維度 | Feature Mapping | Example Mapping | Working with Tables |
|------|----------------|-----------------|---------------------|
| **抽象層次** | 高（全局規劃） | 中（單一故事） | 低（表達工具） |
| **時機** | 專案/迭代初期 | 每個 Story 開發前 | Discovery + Formulation 兩階段皆用 |
| **輸出** | Feature Map 文件 | `specs/*-discovery.md` | 表格（自然語言或 Gherkin） |
| **核心問題** | 「我們要做什麼？」 | 「這個故事的規則是什麼？」 | 「如何清楚表達多個範例？」 |
| **參與者** | PO、業務代表、技術主管 | Three Amigos | 開發 + 測試 |
| **持續時間** | 數小時（規劃會議） | 25 分鐘（時間盒） | 隨寫隨用（無固定時間） |

---

## 2. 三者的本質差異

### Feature Mapping：「選擇要做什麼」

Feature Mapping 解決的是**優先順序與範圍**的問題。它回答：

- 這個業務領域有哪些能力（Capabilities）需要支援？
- 每個能力對應哪些功能（Features）？
- 哪些功能是 MVP 必須有的？
- 各功能之間的依賴關係是什麼？

**核心輸出：** 一份有優先順序的功能清單，明確定義 Must / Should / Could / Won't。

---

### Example Mapping：「搞清楚怎麼做」

Example Mapping 解決的是**單一故事的需求理解**問題。它回答：

- 這個故事有哪些業務規則？
- 每條規則在不同情境下的行為是什麼？
- 還有哪些不確定的問題需要釐清？
- 這個故事夠小、可以開始開發了嗎？

**核心輸出：** 一份包含規則、範例（表格）、待解決問題的 discovery 文件。

---

### Working with Tables：「清楚地表達出來」

Working with Tables 解決的是**範例的表達方式**問題。它回答：

- 如何用最簡潔的方式表示多個相似範例？
- Discovery 階段用什麼格式的表格？
- Formulation 階段用 Scenario Outline 還是 Data Table？
- 如何在 Go/Godog 中解析表格資料？

**核心輸出：** 可讀性高、可維護的表格格式（Discovery 或 Gherkin）。

---

## 3. 核心理念的對應

三種技術共同體現 BDD 的兩個核心理念：

### 刻意探索（Deliberate Discovery）

| 技術 | 如何體現刻意探索 |
|------|----------------|
| **Feature Mapping** | 主動識別所有能力和功能，避免開發後期才發現遺漏的功能域 |
| **Example Mapping** | 25 分鐘工作坊主動揭露業務規則盲點和待解決問題（紅色便利貼） |
| **Working with Tables** | 邊界值測試（正常/異常/邊界三行）確保沒有遺漏的情境 |

### 實際選擇（Real Options）

| 技術 | 如何體現實際選擇 |
|------|----------------|
| **Feature Mapping** | Could/Won't have 的功能先標記，等有更多資訊再決定是否開發 |
| **Example Mapping** | 紅色便利貼記錄問題而不立即決定，Clarify 循環在有答案時才更新 |
| **Working with Tables** | Discovery 階段用自然語言表格，不過早鎖定 Gherkin 格式（保留 Formulation 的彈性） |

---

## 4. 流程中的位置

```
業務目標
    ↓
Feature Mapping ──────────────────────────────────────────────────┐
  • 識別 Capabilities / Features / Stories                        │
  • 排定 MoSCoW 優先順序                                          │
  • 定義 MVP 範圍                                                  │ 全局視圖
    ↓ 選出最高優先 Story（DFS 原則）                               │
    ↓ ────────────────────────────────────────────────────────────┘

Example Mapping ──────────────────────────────────────────────────┐
  • 探索業務規則（藍色便利貼）                                     │
  • 用表格描述範例（綠色便利貼）←── Working with Tables（Discovery） │ 單一故事
  • 識別問題（紅色便利貼）                                         │
  • Clarify 循環（消除黃燈）                                        │
    ↓ 準備度 🟢 綠燈                                              │
    ↓ ────────────────────────────────────────────────────────────┘

Formulation
  • Example Mapping 表格 → Gherkin Scenario Outline ←── Working with Tables
  • specs/*-discovery.md → features/*.feature

Automation
  • Step Definitions 解析 Data Table ←── Working with Tables（Go 實作）
  • 執行 Godog 測試
```

---

## 5. 常見混淆與釐清

### 混淆 1：Feature Mapping 和 Example Mapping 的邊界

| 常見錯誤 | 正確做法 |
|---------|---------|
| 在 Feature Mapping 就討論每個 Story 的業務規則細節 | Feature Mapping 只識別 Story 存在、排優先，細節留給 Example Mapping |
| 對所有 Story 同時做 Example Mapping | 遵循 DFS：一次只深入一個 Story |
| 跳過 Example Mapping 直接從 Feature Map 開始寫 Gherkin | Example Mapping 是必要步驟，確保規則被正確理解 |

### 混淆 2：Discovery 表格 vs. Gherkin 表格

| | Discovery 表格 | Gherkin Scenario Outline | Gherkin Data Table |
|---|---|---|---|
| **格式** | 自然語言（初始狀態/動作/預期結果） | `Given/When/Then` + `Examples:` | 步驟後縮排表格 |
| **使用時機** | Example Mapping 會議中 | Formulation 階段 | Formulation 階段 |
| **用途** | 促進討論、對齊理解 | 資料驅動、多組值驗證 | 單步驟傳入多筆資料 |
| **產生測試數** | 不產生測試 | 每行一個獨立場景 | 仍是一個場景 |

### 混淆 3：Capability vs. Feature vs. User Story

| | Capability | Feature | User Story |
|---|---|---|---|
| **抽象程度** | 最高 | 中 | 最低 |
| **與實作的關係** | 完全無關 | 對應 `.feature` 檔案 | 對應開發任務 |
| **穩定性** | 最穩定（業務目標少改變） | 中等 | 最容易變動 |
| **範例** | 「客戶資金管理」 | 「帳戶間轉帳」 | 「基本帳戶間轉帳」 |

---

## 6. 各技術的關鍵成功因素

### Feature Mapping 成功關鍵

- ✅ 以業務目標為起點，而非技術功能
- ✅ 使用 MoSCoW 明確定義 MVP 邊界
- ✅ 大 Feature 垂直切分（每塊有完整業務價值），避免水平切分（技術分層）
- ❌ 避免在此階段討論技術實作細節
- ❌ 避免跳過 Capability 層次，直接從 Feature 開始

### Example Mapping 成功關鍵

- ✅ 嚴守 25 分鐘時間盒（超時代表故事太大）
- ✅ Three Amigos 三方參與（業務、開發、測試視角缺一不可）
- ✅ 用表格（自然語言）描述範例，禁用 Given/When/Then
- ✅ 問題（紅色便利貼）記錄但不當場爭論，進入 Clarify 循環解決
- ❌ 避免跳過 Clarify 直接在黃燈狀態進入 Formulation

### Working with Tables 成功關鍵

- ✅ 欄位名稱用業務語言（讓業務人員能讀懂）
- ✅ Examples 表格加「情境」欄（測試報告可讀）
- ✅ 每條規則的表格只涵蓋該規則的範例（不混用）
- ✅ 行數控制在 3-7 行（太多考慮拆分）
- ❌ 避免 Discovery 階段使用 Gherkin 表格（破壞探索自由度）

---

## 7. 三種技術的整合使用建議

### 推薦工作流程

```
1. [Feature Mapping]
   開專案/迭代時，先做 Feature Map
   → 識別所有 Capabilities 和 Features
   → 用 MoSCoW 排序，定義 MVP
   → 產出：specs/[domain]-feature-map.md

2. [Example Mapping] ← 依 DFS 順序，一次一個 Story
   針對 Must have 的第一個 Story
   → 25 分鐘工作坊，Three Amigos 參與
   → 用自然語言表格（Working with Tables：Discovery 格式）描述範例
   → 評估準備度，視情況進入 Clarify 循環
   → 產出：specs/[feature]-discovery.md

3. [Formulation]
   Discovery 達到 🟢 綠燈後
   → 將自然語言表格（Working with Tables）轉換為 Gherkin
   → 多個相似範例 → Scenario Outline + Examples
   → 多筆初始資料 → Data Table
   → 產出：features/[feature].feature

4. [Automation]
   → 實作 Step Definitions，解析 Data Table（Working with Tables：Go 實作）
   → 遵循 Red-Green-Refactor
```

### 何時使用哪種格式的表格

```
你在哪個階段？
    ├─ Discovery（Example Mapping 會議）
    │   └─ 用自然語言表格
    │       | 範例 | 初始狀態 | 執行動作 | 預期結果 |
    │
    └─ Formulation/Automation（寫 Gherkin）
        ├─ 同一個流程、只有資料不同？
        │   └─ 用 Scenario Outline + Examples:
        │
        └─ 某個步驟需要傳入多筆資料？
            └─ 用 Data Table（步驟後縮排）
```

---

## 8. 與書中其他章節的連結

| 本章技術 | 相關章節 |
|---------|---------|
| Feature Mapping | Chapter 03：BDD 六步驟流程（推測/說明階段）|
| Example Mapping | Chapter 03：Example Mapping 工作坊；Chapter 06 深化 |
| Working with Tables | Chapter 03：Formulation 階段；Chapter 06 深化 |
| Clarify 循環 | 本章新增（消除 🟡 黃燈的結構化方法）|
| DFS 原則 | [bdd-workflow-discovery skill](../../.claude/skills/bdd-workflow-discovery/dfs-principle.md)|

---

## 參考文件

- [feature-mapping.md](./feature-mapping.md)
- [example-mapping.md](./example-mapping.md)
- [working-with-tables.md](./working-with-tables.md)
- [Chapter 03：BDD 流程](../chapter03/01-bdd-process.md)
- [故事拆分指南](../chapter03/06-story-splitting.md)
