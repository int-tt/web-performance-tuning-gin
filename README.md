#やったこと
## 11/22

 + messages textにindexを貼った
 + buffer poolを大きくした
 + "SQLのcount(*)をcount(id)に変更した"
 + innodb_buffer_pool_size = 768Mにした
 + max_connections=256 にした
 + thread_cache=256にした

### コメント

**Transaction rateは6くらい**

## 11/23

 + followsのエンジンをMyISAMに変更
 	+ 30秒から20秒くらいになった
 + messagesの create_at にindexを貼った
 + part1にvarnishを置いた
 + select * from messages where title = ? order by created_at desc limit 10を
 + select title,message,created_at from messages where title = ? order by created_at desc limit 10 に変更
 + max_connections=512に変更
 + thread_cache=512に変更
 + follows の follow_user_idにindexを貼り直した(?)(要検証)(user_idが早いのにfollow_user_idが遅いのが気になった)
 + yum update -y
### コメント

まだまだMysqlがネックになってて言語を置き換えるよりも先にやることが多い感じ
MYSQLはusersだけがめっちゃデカくてメモリに載り切らない(しかもどうでも良いテーブルが邪魔で)なので、他のはmemcachedに載せるのもあり?
↑解消された
多分SQLで叩いてるテーブルには全部indexを貼った。ボトルネックがSQLからapacheに移りつつあるので、nignxにするかGoに変更したい。

**Transaction rateは127くらい**

## 11/24
 + apacheからngix+php-fpmに乗り換え
  + あんまり変わらず。そもそもマルチプロセスからマルチスレッドにさせたかったのにphp-fpmがマルチプロセスで動いてるので意味なし
 + hhvmのインストールが鬼門

### コメント
　hhvmのせいでmemcachedが死んだ

 **Transaction rateは150くらい**

## 11/25

 + PHPを全部Golangに書き換えた。4倍位早くなった
 + Varnish->Golang構成
 + Part5にもvarnishのキャッシュを作成
 + part2のcount SQLを統合->あんま効果なし。
 + やっぱりPart2,4がネック
  + Template処理が遅い?
 + 回数増やすとファイルディクリプタ問題が出てきた
 + frameworkをginに置き換えた

### コメント

Golang最強、frameworkがネックになると思ったけどそんなことなかった
RDB側はIndexとキャッシュ増やした以外にあんまり効果ない感じだった
part5が静的コンテンツに何故気付かなかったし。
だいぶネックのレイヤーが低くなってきた
varnishあればnginxいらない気がしてきた

**Transaction rateは900くらい**

##11/26
 + ulimit -n 10000 をした
 + go側でmysqlのconnection数を1024に
 + mysql側でmax_conections=2560に
 + thread_cacheを1024に
 +  **重要** wait_timeout = 2に
 + go側のlog表示をoffにした

###コメント

 + カーネル側がネックになってきた
 + seigeもかなりCPUを食うようになってきた->seigeもrootじゃないとファイルディスクリプタで死ぬ
 + log表示offは強い 500くらい伸びた
 + この辺が潮時な気がしてた。まだ少し残ってるからゆっくり試していく

**Transaction rateは1430くらい**

##11/26,27

 + ulimit -n 655360 にした
 + slow_qurey_logを吐かないようにした
 + varnishをnginxに置き換えた
 + nginxとgoをunix domain socketでつなぐようにした
 + innodb_buffer_pool_size 1248Mに変更
 + max_connectionsを4096
 + thread_cacheを4096に変更した
 + acsess_log,error_logを吐かないようにした
 + goやsiegeをroot権限で動かすようにした
 + 後他にもやった気がしたけど忘れた

##コメント

siegeが1500辺りからsocket errで死ぬのでおしまい
一応やれることはやったはず。やり残してるとしたらDB周りが改善できそう
現状のネックはINSERT文
1週間疲れたけど楽しかったです　まる

**最高Transaction rateは2000くらい** 
