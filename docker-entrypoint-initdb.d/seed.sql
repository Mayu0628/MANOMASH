INSERT INTO users (
  user_id, user_name, email, password, introduce, created_at, updated_at
) VALUES (
  1,'木下真菜', 'hoge@hoge.com', 'password', '私は木下真菜です。', now(), now()
);

INSERT INTO users (
  user_id, user_name, email, password, introduce, created_at, updated_at
) VALUES (
  2,'二瓶俊介', 'nihei@hoge.com', 'password。', '私は二瓶俊介です。', now(), now()
);

INSERT INTO oshis (
  user_id,oshi_id, oshi_name, birthday, oshi_meet, oshi_like1, oshi_like2, oshi_like3 ,free_space, interest, reaction_num, created_at, updated_at
) VALUES (
  1,1,'中村ゆきの', '20190907', 'もし会えたら一緒にサッカーをしたいです！！', 'A', 'B', 'C',  '自由に記入しちゃうよ！', '推しについて気になったらここをみてね！！', 5, now(), now()
);
INSERT INTO oshis (
  user_id,oshi_id, oshi_name, birthday, oshi_meet, oshi_like1, oshi_like2, oshi_like3 ,free_space, interest, reaction_num, created_at, updated_at
) VALUES (
  1,2,'中村ゆきの2', '20190907', 'もし会えたら一緒にサッカーをしたいです！！', 'A', 'B', 'C',  '自由に記入しちゃうよ！', '推しについて気になったらここをみてね！！', 5, now(), now()
);
INSERT INTO oshis (
  user_id,oshi_id, oshi_name, birthday, oshi_meet, oshi_like1, oshi_like2, oshi_like3 ,free_space, interest, reaction_num, created_at, updated_at
) VALUES (
  2,3,'中村ゆきの3', '20190907', 'もし会えたら一緒にサッカーをしたいです！！', 'A', 'B', 'C',  '自由に記入しちゃうよ！', '推しについて気になったらここをみてね！！', 5, now(), now()
);
INSERT INTO oshis (
  user_id,oshi_id, oshi_name, birthday, oshi_meet, oshi_like1, oshi_like2, oshi_like3 ,free_space, interest, reaction_num, created_at, updated_at
) VALUES (
  2,4,'中村ゆきの4', '20190907', 'もし会えたら一緒にサッカーをしたいです！！', 'A', 'B', 'C',  '自由に記入しちゃうよ！', '推しについて気になったらここをみてね！！', 5, now(), now()
);

INSERT INTO comments (
  oshi_id, nice_num, comment, user_id, created_at, updated_at
) VALUES (
  1, 10, 'すごくセンスありますね！！', 2, now(), now()
);
