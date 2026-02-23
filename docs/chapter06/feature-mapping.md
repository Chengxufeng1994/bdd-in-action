# Feature Mapping：從業務目標到功能組織的全局視圖

## 概述

Feature Mapping 是一種在 BDD 中用來組織和規劃功能的高層次技術。它幫助團隊從業務目標出發，識別系統需要支援的**能力（Capabilities）**，再將能力拆解為**功能（Features）**，最終用**使用者故事（User Stories）**驅動交付。

---

## 核心概念層次

```
業務目標（Business Goal）
    ↓
能力（Capability）          支持業務目標的抽象能力，與實作無關
    ↓
功能（Feature）             可交付的軟體功能，為使用者提供某種能力
    ↓
使用者故事（User Story）    敏捷專案規劃與交付的最小單位
    ↓
範例（Examples）            具體的使用場景（透過 Example Mapping 發現）
```

---

## 三個層次的定義

### 能力（Capability）

**定義：** 支持某個業務目標的能力，**與具體實作方式無關**。

| 特性 | 說明 |
|------|------|
| 以業務結果為導向 | 「讓客戶能自主管理資金」而非「提供轉帳 API」 |
| 抽象、穩定 | 業務目標很少改變，能力層次也相對穩定 |
| 多個 Feature 共同支持一個 Capability | Capability 是 Feature 的父層次 |

**範例：**
```
能力：客戶資金管理
  → 讓客戶能夠在不同帳戶間調配資金
  → 讓客戶能夠追蹤財務活動
  → 讓客戶能夠保護帳戶安全
```

---

### 功能（Feature）

**定義：** 一個可交付的軟體功能，能夠為使用者提供某種能力。

| 特性 | 說明 |
|------|------|
| 可交付 | 可以獨立部署、展示給業務人員 |
| 對應一個能力的某個面向 | 一個 Feature 支持一個 Capability 的部分需求 |
| 可拆分 | 較大的 Feature 可以拆解為較小的 Feature |
| 對應 `.feature` 檔案 | Formulation 後輸出為 Gherkin feature 檔案 |

**範例（對應「客戶資金管理」能力）：**
```
功能：帳戶間轉帳        → 讓客戶在自己的帳戶間轉移資金
功能：交易記錄查詢      → 讓客戶檢視過去的交易歷史
功能：轉帳限額設定      → 讓客戶自訂每日轉帳上限
```

**拆分原則（大 Feature → 小 Feature）：**

```
大功能：帳戶管理
    ├─ 小功能：帳戶間轉帳（Must have）
    ├─ 小功能：交易記錄查詢（Should have）
    └─ 小功能：轉帳限額設定（Could have）
```

拆分時遵循**垂直切分**：每個小功能包含完整的業務價值，而非技術層次的水平切分。

---

### 使用者故事（User Story）

**定義：** 敏捷專案中規劃與交付功能的單位，代表一個使用者能夠完成的具體行為。

**格式：**
```
As a <角色>
I want <需求>
So that <業務價值>
```

**INVEST 原則：**

| 字母 | 原則 | 說明 |
|------|------|------|
| **I** | Independent（獨立） | 故事間無強依賴 |
| **N** | Negotiable（可協商） | 範圍可討論調整 |
| **V** | Valuable（有價值） | 對業務有明確價值 |
| **E** | Estimable（可估算） | 能評估開發工作量 |
| **S** | Small（小的） | 可在 1-2 天內完成 |
| **T** | Testable（可測試） | 有清楚的驗收標準 |

---

## Feature Mapping 的輸出格式

Feature Map 是一份從業務目標到使用者故事的完整視圖：

```
業務目標：提升客戶金融自主性
│
├─ 能力：客戶資金管理
│   ├─ 功能：帳戶間轉帳
│   │   ├─ Story 1.1：基本帳戶間轉帳（Must）
│   │   ├─ Story 1.2：轉帳限額驗證（Must）
│   │   └─ Story 1.3：VIP 客戶高額轉帳（Should）
│   │
│   ├─ 功能：交易記錄查詢
│   │   ├─ Story 2.1：查看最近 30 天交易（Should）
│   │   └─ Story 2.2：匯出交易記錄（Could）
│   │
│   └─ 功能：轉帳限額設定
│       └─ Story 3.1：設定每日轉帳限額（Could）
│
└─ 能力：帳戶安全
    ├─ 功能：雙重驗證
    │   └─ Story 4.1：轉帳前 OTP 驗證（Must）
    └─ 功能：異常偵測
        └─ Story 5.1：大額轉帳警示通知（Should）
```

---

## Feature Mapping 與 Example Mapping 的關係

| | Feature Mapping | Example Mapping |
|---|---|---|
| **層次** | 高層次（全局視圖） | 單一故事層次（深入細節） |
| **目的** | 組織所有功能、識別 MVP | 發現業務規則、釐清範例 |
| **時機** | 專案/迭代規劃初期 | 每個故事開發前 |
| **輸出** | Feature Map 文件 | `specs/*-discovery.md` |
| **參與者** | 產品負責人、業務代表、技術主管 | Three Amigos |

**互補關係：**
```
Feature Mapping（規劃）
    → 識別所有 Feature 和 Story
    → 排定優先順序（MoSCoW）
    → 定義 MVP 範圍

Example Mapping（深入）
    → 針對每個 Story 進行
    → 發現業務規則和範例
    → 識別待解決問題
```

---

## 優先順序：MoSCoW 方法

Feature Mapping 結合 MoSCoW 排定交付優先順序：

| 優先級 | 說明 | 納入時機 |
|--------|------|---------|
| **Must have** | 沒有這個功能，產品無法運作 | MVP（第一版）必須包含 |
| **Should have** | 重要但非緊急，可稍後添加 | 第二版 |
| **Could have** | 錦上添花，時間允許才做 | 第三版或以後 |
| **Won't have** | 這次不做，可能未來考慮 | Backlog |

**MVP 識別範例：**
```
MVP 範圍（Must have）：
  ✅ Story 1.1：基本帳戶間轉帳
  ✅ Story 1.2：轉帳限額驗證
  ✅ Story 4.1：轉帳前 OTP 驗證

第二版（Should have）：
  ⏭️ Story 1.3：VIP 客戶高額轉帳
  ⏭️ Story 2.1：查看最近 30 天交易

第三版（Could have）：
  ⏭️ Story 2.2：匯出交易記錄
  ⏭️ Story 3.1：設定每日轉帳限額
```

---

## 與 DFS 原則的配合

Feature Mapping 提供全局視圖（廣度），但執行時仍遵循 DFS（深度優先）：

```
Feature Mapping 產出優先順序清單
    ↓
選擇優先級最高的 Story（Must have 第一）
    ↓
針對該 Story 執行完整 Example Mapping（DFS 深入）
    ↓
完成 Clarify → Formulation → Automation
    ↓
回到 Feature Map，選擇下一個 Story
```

**不要做的事：**
- ❌ 對所有 Story 同時進行 Example Mapping（BFS 廣度優先）
- ❌ Feature Mapping 後直接跳入開發，跳過 Example Mapping
- ❌ 在 Feature Mapping 階段就規劃所有細節（留到 Example Mapping）

---

## 實際選擇（Real Options）在 Feature Mapping 的應用

Feature Mapping 是「保留選擇」的最佳時機：

| 選擇 | 建議做法 |
|------|---------|
| 功能邊界未確定 | 標記為 Could have，等有更多資訊再決定 |
| 技術方案未確定 | 只定義業務行為，不討論實作細節 |
| 某個故事依賴未完成的研究 | 標記為 Won't have（此迭代），記錄原因 |
| 功能範圍有爭議 | 記錄為問題，指定負責人確認，不立即決定 |

---

## 文件輸出格式

Feature Map 文件建議儲存為 `specs/[business-domain]-feature-map.md`：

```markdown
# Feature Map：[業務領域名稱]

## 業務目標

[描述這個功能域的業務目標]

## 能力與功能

### 能力：[能力名稱]

> [能力描述：支持什麼業務目標，與實作無關]

**功能 1：[功能名稱]**

| Story | 描述 | 優先級 | 估算 | 依賴 |
|-------|------|--------|------|------|
| 1.1 [名稱] | [描述] | Must | 2天 | 無 |
| 1.2 [名稱] | [描述] | Should | 1天 | 1.1 |

**功能 2：[功能名稱]**
...

## MVP 範圍

- ✅ Story 1.1：[名稱]（Must）
- ✅ Story 1.2：[名稱]（Must）

## 交付時間線

- **Sprint 1**：Story 1.1, 1.2（Must have）
- **Sprint 2**：Story 2.1（Should have）
- **Backlog**：Story 3.1（Could have）
```

---

## 參考

- [Example Mapping](./example-mapping.md) - 單一故事的深入發現技術
- [Working with Tables](./working-with-tables.md) - 範例的表格組織方式
- [故事拆分指南](../chapter03/06-story-splitting.md) - 大 Feature 拆分模式
- [DFS 原則](../../.claude/skills/bdd-workflow-discovery/dfs-principle.md) - 深度優先探索原則
