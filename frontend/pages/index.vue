<template>
  <section>
    <section style="height: 100vh;">
      <v-row align="center" justify="center" style="height: 100%;">
        <v-col cols="12" sm="8" md="4">
          <p class="text-h4 text-center">ブイトーーク！</p>
          <p class="text-subtitle-1 text-center">VTuberファンでつくる掲示板</p>
          <div class="text-center">
            <v-btn>はじめる</v-btn>
          </div>
        </v-col>
      </v-row>
    </section>

    <v-row justify="center">
      <v-col cols="12" sm="6">
        <v-card>
          <v-card-title>ブイトーーク！とは？</v-card-title>
          <v-card-text>
            <p>
              <a
                href="https://ja.wikipedia.org/wiki/バーチャルYouTuber"
                target="_blank"
              >
                VTuber</a
              >
              のファンが集い、推しについて語り合う掲示板です。
            </p>
            <p>
              動画公開後や配信後に感想を書き込んだり、おすすめの歌ってみた動画を紹介したり...。
            </p>
            <p>様々な楽しみ方をして頂ければ幸いです。</p>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>

    <v-row justify="center">
      <v-col cols="12" sm="6">
        <v-card>
          <v-card-title>開発者</v-card-title>
          <v-card-text>
            <p><a href="https://lem0n.cc/" target="_blank">もぺねこ</a></p>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>

    <v-row justify="center">
      <v-col cols="12" sm="6">
        <v-card>
          <v-card-title>ソースコード</v-card-title>
          <v-card-text>
            <p>
              このサイトのソースコードは
              <a href="https://github.com/mopeneko/vtalk" target="_blank">
                こちら</a
              >
              にて公開されています。
            </p>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </section>
</template>
