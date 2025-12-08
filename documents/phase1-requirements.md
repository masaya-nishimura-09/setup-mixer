# Phase 1: MVP 要件定義書

**作成日**: 2024 年  
**対象フェーズ**: Phase 1（MVP - 最小限の機能）  
**データソース**: 楽天市場 API

---

## 1. Phase 1 の目標

MVP（Minimum Viable Product）として、以下の 4 つのコア機能を実装する:

1. キーワード検索機能
2. 商品カード一覧表示（14 カテゴリ）（人気順、価格順）
3. 基本的なセット作成機能（状態管理、合計金額を表示、写真を並べる）
4. それぞれのアフィリエイトリンク表示

---

## 2. 機能詳細要件

### 2.1 キーワード検索機能

#### 2.1.1 機能概要

- ユーザーが検索バーにキーワードを入力
- スペース区切りで複数キーワードに対応
- 商品名、説明、タグなどから検索

#### 2.1.2 実装要件

**入力:**

- 検索バー（テキスト入力フィールド）
- プレースホルダー: "例: 白 おしゃれな 狭い"
- 検索ボタンまたは Enter キーで実行

**検索ロジック:**

- スペースで区切られたキーワードを配列に分割
- 各キーワードを商品データの以下フィールドとマッチング:
  - `itemName`（商品名）
  - `itemCaption`（商品説明）
  - `tags`（タグ配列、カスタムフィールド）
- AND 検索（すべてのキーワードに一致する商品を表示）

**出力:**

- 検索結果一覧ページに遷移
- 該当商品をグリッド形式で表示
- 検索結果件数を表示

#### 2.1.3 技術実装

- **フロントエンド**: Next.js の検索ページ
- **検索処理**: クライアントサイドまたはサーバーサイド（要検討）
- **データソース**: 楽天 API から取得した JSON データ

### 2.2 商品カード一覧表示

#### 2.2.1 機能概要

- 14 カテゴリの商品を一覧表示
- ソート機能（人気順、価格順）
- グリッド形式で表示

#### 2.2.2 カテゴリ一覧

1. 机
2. 椅子
3. キーボード
4. モニター
5. モニター台
6. マウス
7. マウスパッド
8. マット
9. ライト
10. ケーブル類
11. スピーカー
12. カメラ
13. マイク
14. ヘッドホン

#### 2.2.3 商品カードの表示項目

**必須項目:**

- **商品画像**: `mediumImageUrls[0]`（メイン画像）
- **商品名**: `itemName`
- **価格**: `itemPrice`（税込表示、3 桁区切り）
- **短い説明**: `itemCaption` の最初の 100 文字程度
- **レビュー情報**:
  - 平均評価: `reviewAverage`（★ 表示、5 段階）
  - レビュー件数: `reviewCount`
- **アフィリエイトリンク**: `affiliateUrl` または `itemUrl`

**表示形式:**

- カード形式（画像 + 情報）
- ホバー時に詳細情報を表示（オプション）
- クリックで商品詳細ページまたは購入ページへ遷移

#### 2.2.4 ソート機能

**ソートオプション:**

1. **人気順**（デフォルト）:

   - `reviewCount` の降順
   - 同数の場合は `reviewAverage` の降順

2. **価格順（安い順）**:

   - `itemPrice` の昇順

3. **価格順（高い順）**:
   - `itemPrice` の降順

**UI:**

- ドロップダウンメニューまたはタブで選択
- 選択状態を視覚的に表示

#### 2.2.5 技術実装

- **データ取得**: 楽天 API から取得した JSON ファイルを読み込み
- **状態管理**: Zustand または React Context
- **表示**: Next.js のページコンポーネント
- **スタイリング**: Tailwind CSS

### 2.3 セット作成機能

#### 2.3.1 機能概要

- ユーザーが各カテゴリから商品を選択
- 選択した商品をセットとして管理
- セットの合計金額を表示
- 選択した商品の写真を並べて表示

#### 2.3.2 実装要件

**商品選択:**

- 各カテゴリごとに商品一覧を表示
- 商品カードに「セットに追加」ボタンを表示
- クリックでセットに追加
- 既に追加済みの商品は「追加済み」または「削除」ボタンを表示

**セット管理:**

- セットに追加された商品を状態管理で保持
- 各カテゴリから最大 1 商品まで選択可能（要確認）
- または、複数商品を選択可能（要確認）

**セット表示エリア:**

- 選択した商品の写真を横並びで表示
- 各商品の名前と価格を表示
- 合計金額を大きく表示
- 各商品の削除ボタン

**合計金額計算:**

- セット内の全商品の価格を合計
- リアルタイムで更新
- 3 桁区切りで表示（例: ¥123,456）

#### 2.3.3 UI/UX 要件

**レイアウト:**

- 左側: カテゴリ選択タブ + 商品一覧
- 右側: セット表示エリア（固定またはスティッキー）

**状態表示:**

- セットに追加された商品は視覚的に区別（ボーダー、背景色など）
- セットが空の場合は「商品を選択してください」と表示

#### 2.3.4 技術実装

- **状態管理**: Zustand（推奨）または React Context
- **データ構造**:

  ```typescript
  interface SetItem {
    category: string;
    product: Product;
  }

  interface SetState {
    items: SetItem[];
    totalPrice: number;
    addItem: (category: string, product: Product) => void;
    removeItem: (category: string) => void;
    clearSet: () => void;
  }
  ```

### 2.4 アフィリエイトリンク表示

#### 2.4.1 機能概要

- セットに含まれる各商品のアフィリエイトリンクを表示
- まとめて購入できるように整理

#### 2.4.2 実装要件

**リンク表示:**

- セット表示エリアに「購入する」セクションを追加
- 各商品のアフィリエイトリンクをリスト表示
- 新しいタブで開く（`target="_blank"`）
- 外部リンクであることを明示（アイコン表示）

**UI:**

- 「購入する」ボタンまたはセクション
- 各商品のリンクをクリック可能なボタンまたはリンクとして表示
- 商品名と価格を併記

#### 2.4.3 技術実装

- 楽天アフィリエイトリンクを使用
- `affiliateUrl` または `itemUrl` を使用
- 楽天アフィリエイトプログラムに参加後、適切なタグを付与

---

## 3. データ構造

### 3.1 商品データ（楽天 API レスポンス）

```typescript
interface RakutenProduct {
  itemName: string; // 商品名
  itemPrice: number; // 価格
  itemUrl: string; // 商品URL
  affiliateUrl: string; // アフィリエイトURL
  mediumImageUrls: string[]; // 画像URL配列
  itemCaption: string; // 商品説明
  reviewAverage: number; // 平均評価（0-5）
  reviewCount: number; // レビュー件数
  itemCode: string; // 商品コード
  // カスタムフィールド（追加）
  category: string; // カテゴリ名
  tags?: string[]; // 検索用タグ
  color?: string; // 色
  style?: string; // スタイル
}
```

### 3.2 セットデータ

```typescript
interface SetItem {
  category: string;
  product: RakutenProduct;
}

interface Set {
  items: SetItem[];
  totalPrice: number;
}
```

### 3.3 データファイル構造

**推奨構造:**

```
/data
  /products
    desk.json          # 机カテゴリの商品
    chair.json         # 椅子カテゴリの商品
    keyboard.json      # キーボードカテゴリの商品
    ...
  categories.json      # カテゴリ一覧
```

**JSON ファイル例（desk.json）:**

```json
{
  "category": "机",
  "products": [
    {
      "itemName": "デスク 白",
      "itemPrice": 15000,
      "itemUrl": "https://...",
      "affiliateUrl": "https://...",
      "mediumImageUrls": ["https://..."],
      "itemCaption": "シンプルでおしゃれな白いデスク...",
      "reviewAverage": 4.5,
      "reviewCount": 123,
      "itemCode": "R123456",
      "category": "机",
      "tags": ["白", "おしゃれ", "ミニマル"],
      "color": "白",
      "style": "ミニマル"
    }
  ]
}
```

---

## 4. 技術スタック

### 4.1 フロントエンド

- **フレームワーク**: Next.js 14+（App Router）
- **言語**: TypeScript
- **スタイリング**: Tailwind CSS
- **状態管理**: Zustand（推奨）または React Context
- **UI コンポーネント**: shadcn/ui（推奨）または自作

### 4.2 データ取得

- **API**: 楽天市場 API（IchibaItemSearchAPI）
- **データ保存**: JSON ファイル（`/data` ディレクトリ）
- **データ取得スクリプト**: Node.js スクリプト（`/scripts` ディレクトリ）

### 4.3 開発環境

- **パッケージマネージャー**: npm または yarn
- **バージョン管理**: Git
- **コードフォーマット**: Prettier
- **リンター**: ESLint

---

## 5. ページ構成

### 5.1 ページ一覧

1. **トップページ** (`/`)

   - 検索バー
   - カテゴリ一覧
   - 人気商品の表示（オプション）

2. **検索結果ページ** (`/search?q=キーワード`)

   - 検索結果一覧
   - ソート機能
   - ページネーション（オプション）

3. **カテゴリページ** (`/category/[categoryName]`)

   - カテゴリ別商品一覧
   - ソート機能

4. **セット作成ページ** (`/set`)
   - カテゴリ選択タブ
   - 商品一覧
   - セット表示エリア
   - 購入リンク

### 5.2 レイアウト

- **共通レイアウト**: ヘッダー、フッター
- **レスポンシブデザイン**: モバイル、タブレット、PC 対応

---

## 6. UI/UX 要件

### 6.1 デザイン方針

- **ミニマルで洗練されたデザイン**
- **白を基調とした配色**
- **商品画像を際立たせるレイアウト**
- **直感的な操作**

### 6.2 カラーパレット（案）

- **背景**: 白（#FFFFFF）
- **テキスト**: ダークグレー（#333333）
- **アクセント**: ブルーまたはグレー（要検討）
- **ボーダー**: ライトグレー（#E5E5E5）

### 6.3 コンポーネント

**必須コンポーネント:**

- `SearchBar`: 検索バー
- `ProductCard`: 商品カード
- `ProductGrid`: 商品グリッド
- `CategoryTabs`: カテゴリタブ
- `SetDisplay`: セット表示エリア
- `SortDropdown`: ソートドロップダウン

---

## 7. 楽天 API の使用方法

### 7.1 登録手順

1. 楽天ウェブサービスにアクセス: https://webservice.rakuten.co.jp/
2. 無料アカウントを作成
3. アプリケーション ID を取得

### 7.2 API エンドポイント

**商品検索 API:**

```
GET https://app.rakuten.co.jp/services/api/IchibaItem/Search/20170706
```

**パラメータ:**

- `applicationId`: アプリケーション ID（必須）
- `keyword`: 検索キーワード
- `genreId`: ジャンル ID（カテゴリ指定）
- `sort`: ソート順（`+reviewCount` など）
- `hits`: 取得件数（デフォルト: 30、最大: 30）
- `page`: ページ番号

### 7.3 データ取得スクリプト

**実装例（Node.js）:**

```javascript
// scripts/fetch-products.js
const axios = require("axios");
const fs = require("fs");
const path = require("path");

const APPLICATION_ID = "YOUR_APPLICATION_ID";
const BASE_URL =
  "https://app.rakuten.co.jp/services/api/IchibaItem/Search/20170706";

// カテゴリごとのキーワードマッピング
const categoryKeywords = {
  机: "デスク 机",
  椅子: "チェア 椅子",
  キーボード: "キーボード",
  // ...
};

async function fetchProducts(category, keyword) {
  try {
    const response = await axios.get(BASE_URL, {
      params: {
        applicationId: APPLICATION_ID,
        keyword: keyword,
        hits: 20,
        sort: "+reviewCount",
      },
    });

    return response.data.Items.map((item) => ({
      ...item.Item,
      category: category,
      tags: extractTags(item.Item),
    }));
  } catch (error) {
    console.error(`Error fetching ${category}:`, error);
    return [];
  }
}

// 各カテゴリの商品を取得してJSONファイルに保存
async function main() {
  for (const [category, keyword] of Object.entries(categoryKeywords)) {
    const products = await fetchProducts(category, keyword);
    const filePath = path.join(
      __dirname,
      "../data/products",
      `${category}.json`
    );

    fs.writeFileSync(
      filePath,
      JSON.stringify({ category, products }, null, 2),
      "utf-8"
    );

    console.log(`Saved ${products.length} products for ${category}`);
    // API制限を考慮して待機
    await new Promise((resolve) => setTimeout(resolve, 1000));
  }
}

main();
```

---

## 8. 開発手順

### 8.1 プロジェクトセットアップ

1. **Next.js プロジェクトの作成**

   ```bash
   npx create-next-app@latest setup-mixer --typescript --tailwind --app
   cd setup-mixer
   ```

2. **必要なパッケージのインストール**

   ```bash
   npm install zustand axios
   # または
   npm install @tanstack/react-query axios
   ```

3. **ディレクトリ構造の作成**
   ```
   setup-mixer/
   ├── app/
   ├── components/
   ├── data/
   │   └── products/
   ├── lib/
   ├── scripts/
   └── types/
   ```

### 8.2 データ取得

1. **楽天 API の登録**
2. **データ取得スクリプトの作成**（上記参照）
3. **サンプルデータの取得**（各カテゴリ 10-20 件）
4. **JSON ファイルに保存**

### 8.3 機能実装の順序

1. **データ読み込み機能**

   - JSON ファイルから商品データを読み込む
   - TypeScript の型定義を作成

2. **商品カードコンポーネント**

   - 基本的な商品カードを作成
   - スタイリング

3. **商品一覧ページ**

   - カテゴリ別の商品一覧を表示
   - ソート機能を実装

4. **検索機能**

   - 検索バーコンポーネント
   - 検索ロジックの実装

5. **セット作成機能**

   - 状態管理のセットアップ（Zustand）
   - セット表示コンポーネント
   - 合計金額の計算

6. **アフィリエイトリンク表示**
   - リンク表示コンポーネント
   - セットからの購入機能

### 8.4 テスト

- 各機能の動作確認
- レスポンシブデザインの確認
- ブラウザ互換性の確認

---

## 9. チェックリスト

### 9.1 準備段階

- [ ] Next.js プロジェクトの作成
- [ ] 楽天ウェブサービスへの登録
- [ ] アプリケーション ID の取得
- [ ] データ取得スクリプトの作成
- [ ] サンプルデータの取得（各カテゴリ 10-20 件）
- [ ] JSON ファイルへの保存

### 9.2 開発段階

- [ ] TypeScript 型定義の作成
- [ ] 商品カードコンポーネントの実装
- [ ] 商品一覧ページの実装
- [ ] ソート機能の実装
- [ ] 検索機能の実装
- [ ] セット作成機能の実装
- [ ] 状態管理の実装
- [ ] アフィリエイトリンク表示の実装
- [ ] レスポンシブデザインの実装

### 9.3 完成段階

- [ ] 全機能の動作確認
- [ ] デザインの調整
- [ ] パフォーマンスの最適化
- [ ] エラーハンドリングの実装
- [ ] ブラウザテスト

---

## 10. 注意事項

### 10.1 楽天 API の利用規約

- 利用規約を確認し、遵守すること
- リクエスト制限を考慮した実装（1 秒あたり 10 リクエスト程度）
- 取得したデータの取り扱いに注意

### 10.2 アフィリエイトリンク

- 楽天アフィリエイトプログラムに参加後、適切なリンクを使用
- 外部リンクであることを明示
- `target="_blank"` と `rel="noopener noreferrer"` を設定

### 10.3 パフォーマンス

- 大量の商品データを扱う場合は、ページネーションを検討
- 画像の遅延読み込み（lazy loading）を実装
- データのキャッシュを検討

---
