# Chapter 03: BDD 流程詳解

## 概述

本章節詳細介紹 BDD (Behavior-Driven Development) 的完整流程，從業務假設到價值驗證的六個關鍵步驟。

## 📚 章節內容

### [01-bdd-process.md](./01-bdd-process.md)
完整的 BDD 六步驟流程指南

### [02-quick-reference.md](./02-quick-reference.md)
BDD 流程快速參考手冊

### [03-workflow-diagrams.md](./03-workflow-diagrams.md)
視覺化流程圖和架構圖

### [04-feature-description.md](./04-feature-description.md)
使用者故事撰寫指南

#### 使用者故事格式
- **In order to... As a... I want...** (目標導向)
- **As a... I want... So that...** (角色導向)
- INVEST 原則
- 驗收標準編寫

### [05-example-mapping.md](./05-example-mapping.md)
Example Mapping 協作技術

#### 四色便利貼系統
- 📋 黃色：使用者故事
- 📏 藍色：業務規則
- 📝 綠色：具體範例
- ❓ 紅色：待解決問題

#### 核心流程
1. **推測 (Speculate)** 🤔
   - 探索業務價值
   - 識別利益相關者
   - 定義成功標準

2. **說明 (Illustrate)** 📖
   - Example Mapping 工作坊
   - 具體範例創建
   - 邊界案例發現

3. **制定 (Formulate)** ✍️
   - Gherkin 場景編寫
   - 場景組織結構
   - 最佳實踐

4. **自動化 (Automate)** 🤖
   - 步驟定義實現
   - 測試架構設計
   - Fluent API 設計

5. **展示 (Demonstrate)** 📊
   - 執行自動化測試
   - 生成測試報告
   - 利益相關者展示

6. **驗證 (Validate)** ✅
   - 業務驗收
   - 指標檢查
   - 持續改進

## 🎯 學習目標

完成本章後，你將能夠：

- ✅ 理解 BDD 的完整生命週期
- ✅ 組織和主持 Example Mapping 工作坊
- ✅ 編寫高質量的 Gherkin 場景
- ✅ 實現可維護的步驟定義
- ✅ 建立有效的測試架構
- ✅ 向利益相關者展示活文檔
- ✅ 驗證業務價值實現

## 🛠️ 實踐練習

### 練習 1：Example Mapping 工作坊
使用便利貼組織一個簡單的功能：
- 📋 使用者故事
- 📏 業務規則
- 📝 具體範例
- ❓ 待解決問題

### 練習 2：Gherkin 場景編寫
將範例轉化為結構化場景：
```gherkin
Feature: [功能名稱]
  Scenario: [場景描述]
    Given [前置條件]
    When [執行動作]
    Then [預期結果]
```

### 練習 3：步驟定義實現
為場景實現可執行的步驟定義

### 練習 4：活文檔展示
運行測試並展示給團隊

## 📊 流程圖

```
推測 (Speculate)
    ↓
說明 (Illustrate)
    ↓
制定 (Formulate)
    ↓
自動化 (Automate)
    ↓
展示 (Demonstrate)
    ↓
驗證 (Validate)
    ↓
[反饋循環回到推測]
```

## 🎭 角色與職責

### Three Amigos 協作模式

| 角色 | 職責 | 參與階段 |
|------|------|----------|
| **業務** | 定義價值和規則 | 推測、說明、驗證 |
| **開發** | 實現功能 | 制定、自動化 |
| **測試** | 確保質量 | 所有階段 |

## 📈 成功指標

### 流程健康度
- Example Mapping 參與率 > 80%
- 場景審查週期 < 1 天
- 測試通過率 > 95%
- 文檔更新及時性 > 90%

### 業務價值
- 缺陷提前發現率 > 70%
- 需求變更適應速度 < 2 天
- 利益相關者滿意度 > 80%
- 交付週期縮短 > 30%

## 🔧 工具與技術

### BDD 框架
- **Godog** - Go 語言的 Cucumber 實現
- **Cucumber** - BDD 測試框架
- **Gherkin** - 場景描述語言

### 協作工具
- 便利貼和白板（Example Mapping）
- Miro / Mural（線上協作）
- JIRA / Trello（需求追蹤）

### 報告工具
- Cucumber JSON Reporter
- Allure Report
- Custom Dashboard

## 📖 相關章節

- [Chapter 01: BDD 快速入門](../chapter01/)
- [Chapter 02: BDD 原則與實踐](../chapter02/)
- [Chapter 02 實作: 銀行應用](../../chapter02/)

## 💡 最佳實踐

### Do ✅
- 從小範圍開始實踐
- 保持業務語言
- 定期重構場景
- 持續收集反饋
- 度量和改進

### Don't ❌
- 過度技術化場景
- 跳過 Example Mapping
- 忽視邊界案例
- 讓場景過時
- 孤立開發

## 🎓 學習路徑

```
1. 閱讀完整流程文檔
   ↓
2. 觀察 Example Mapping 範例
   ↓
3. 編寫第一個 Gherkin 場景
   ↓
4. 實現簡單的步驟定義
   ↓
5. 運行並展示測試
   ↓
6. 收集反饋並改進
```

## 🔗 延伸資源

### 書籍
- "BDD in Action" by John Ferguson Smart
- "The Cucumber Book" by Matt Wynne & Aslak Hellesøy
- "Specification by Example" by Gojko Adzic

### 線上資源
- [Cucumber Documentation](https://cucumber.io/docs/)
- [Godog GitHub](https://github.com/cucumber/godog)
- [BDD Wiki](https://en.wikipedia.org/wiki/Behavior-driven_development)

### 社群
- Cucumber Slack
- BDD Practitioners Group
- Local BDD Meetups

## 📝 快速參考

### Example Mapping 模板
```
📋 使用者故事
  └── 📏 規則 1
       ├── 📝 範例 1.1
       ├── 📝 範例 1.2
       └── ❓ 問題
  └── 📏 規則 2
       └── 📝 範例 2.1
```

### Gherkin 模板
```gherkin
Feature: [業務能力]
  作為 [角色]
  我想要 [功能]
  以便 [業務價值]

  Background:
    Given [共同前置條件]

  Scenario: [具體場景]
    Given [前置條件]
    When [執行動作]
    Then [預期結果]
```

### 步驟定義模板
```go
func (s *Steps) RegisterSteps(sc *godog.ScenarioContext) {
    sc.Step(`^給定(.+)$`, s.given)
    sc.Step(`^當(.+)$`, s.when)
    sc.Step(`^那麼(.+)$`, s.then)
}
```

---

**開始實踐 BDD，讓測試驅動你的開發！** 🚀
