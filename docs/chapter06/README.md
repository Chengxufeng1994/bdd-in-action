# Chapter 06：組織與發現需求的核心技術

本章涵蓋 BDD 工作流程中三個相互配合的核心技術，從全局規劃到細節發現。

## 文件索引

| 文件 | 主題 | 說明 |
|------|------|------|
| [feature-mapping.md](./feature-mapping.md) | Feature Mapping | 從業務目標到功能組織的全局視圖 |
| [example-mapping.md](./example-mapping.md) | Example Mapping | 單一故事的協作發現技術 |
| [working-with-tables.md](./working-with-tables.md) | Working with Tables | 用表格組織範例與測試資料 |

## 三者的關係

```
Feature Mapping（規劃全局）
    → 識別 Capabilities、Features、User Stories
    → 排定 MoSCoW 優先順序
    → 定義 MVP 範圍
         ↓
    選擇優先級最高的 Story
         ↓
Example Mapping（深入一個 Story）
    → 發現業務規則（藍色便利貼）
    → 用表格描述具體範例（綠色便利貼）
    → 識別待解決問題（紅色便利貼）
    → 評估準備度（🟢🟡🔴）
         ↓
Working with Tables（組織範例）
    → Discovery 階段：自然語言表格
    → Formulation 階段：Scenario Outline + Examples
    → Automation 階段：Data Table 傳入結構化資料
```
