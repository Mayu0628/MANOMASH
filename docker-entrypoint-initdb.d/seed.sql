INSERT INTO users (
  user_name, email, password, introduce, created_at, updated_at
) VALUES (
  '木下真菜', 'hoge@hoge.com', '私は木下真菜です。', 'password', now(), now()
);

INSERT INTO users (
  user_name, email, password, introduce, created_at, updated_at
) VALUES (
  '二瓶俊介', 'nihei@hoge.com', '私は二瓶俊介です。', 'password', now(), now()
);

INSERT INTO oshis (
  oshi_name, birthday, oshi_meet, free_space, interest, reaction_num, created_at, updated_at
) VALUES (
  '中村ゆきの', '20190907', 'もし会えたら一緒にサッカーをしたいです！！', '自由に記入しちゃうよ！', '推しについて気になったらここをみてね！！', 5, now(), now()
);

INSERT INTO tag (
  tag_name
) VALUES (
  'サッカー'
);

INSERT INTO oshi_tag (
  oshi_id, tag_id
) VALUES (
  1, 1
);

INSERT INTO oshi_like (
  oshi_id,like_point
) VALUES (
  1,'面白いところ'
);

INSERT INTO comments (
  oshi_id, nice_num, comment, user_id, created_at, updated_at
) VALUES (
  1, 10, 'すごくセンスありますね！！', 2, now(), now()
);
