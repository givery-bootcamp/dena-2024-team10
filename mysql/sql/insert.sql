INSERT INTO hello_worlds (lang, message) VALUES ('en', 'Hello World');
INSERT INTO hello_worlds (lang, message) VALUES ('ja', 'こんにちは 世界');

INSERT INTO users (name, password) VALUES ('taro', '$2a$10$8N/2SHn57ZKQiijE8Cwj4uPXiV5ZV.4yU.sHGibMLuqr92Tl.HCHC'); -- password
INSERT INTO users (name, password) VALUES ('hanako', '$2a$10$oFfc4tR6VAC.Q9HXFz5e7uwe4Pbg0plRbC60JSg8DGDm45ZJIg3h2'); -- PASSWORD

INSERT INTO posts (user_id, title, body) VALUES (1, 'test1', '質問1\n改行');
INSERT INTO posts (user_id, title, body) VALUES (1, 'test2', '質問2\n改行');
INSERT INTO posts (user_id, title, body) VALUES (2, 'test3', '質問3\n改行');
INSERT INTO posts (user_id, title, body) VALUES (2, 'test4', '質問4\n改行');
INSERT INTO posts (user_id, title, body) VALUES (2, 'test5', '質問5\n改行');
INSERT INTO posts (user_id, title, body) VALUES (2, 'test6', '質問6\n改行');
INSERT INTO posts (user_id, title, body) VALUES (2, 'test7', '質問7\n改行');
INSERT INTO posts (user_id, title, body) VALUES (2, 'test8', '質問8\n改行');
INSERT INTO posts (user_id, title, body) VALUES (2, 'test9', '質問9\n改行');
INSERT INTO posts (user_id, title, body) VALUES (2, 'test10', '質問10\n改行');
INSERT INTO posts (user_id, title, body) VALUES (2, 'test11', '質問11\n改行');
INSERT INTO posts (user_id, title, body) VALUES (2, 'test12', '質問12\n改行');
INSERT INTO posts (user_id, title, body) VALUES (2, 'test13', '質問13\n改行');
INSERT INTO posts (user_id, title, body) VALUES (2, 'test14', '質問14\n改行');
INSERT INTO posts (user_id, title, body) VALUES (2, 'test15', '質問15\n改行');
INSERT INTO posts (user_id, title, body) VALUES (2, 'test16', '質問16\n改行');
INSERT INTO posts (user_id, title, body) VALUES (2, 'test17', '質問17\n改行');
INSERT INTO posts (user_id, title, body) VALUES (2, 'test18', '質問18\n改行');
INSERT INTO posts (user_id, title, body) VALUES (2, 'test19', '質問19\n改行');
INSERT INTO posts (user_id, title, body) VALUES (2, 'test20', '質問20\n改行');
INSERT INTO posts (user_id, title, body) VALUES (2, 'test21', '質問21\n改行');
INSERT INTO posts (user_id, title, body) VALUES (2, 'test22', '質問22\n改行');
INSERT INTO posts (user_id, title, body) VALUES (2, 'test23', '質問23\n改行');
INSERT INTO posts (user_id, title, body) VALUES (2, 'test24', '質問24\n改行');
INSERT INTO posts (user_id, title, body) VALUES (2, 'test25', '質問25\n改行');

INSERT INTO comments (post_id, user_id, body) VALUES (1, 1, 'comment1 on test1');
INSERT INTO comments (post_id, user_id, body) VALUES (1, 2, 'comment2 on test1');
INSERT INTO comments (post_id, user_id, body) VALUES (1, 1, 'comment3 on test1');
INSERT INTO comments (post_id, user_id, body) VALUES (1, 2, 'comment4 on test1');
INSERT INTO comments (post_id, user_id, body) VALUES (1, 1, 'comment5 on test1');
INSERT INTO comments (post_id, user_id, body) VALUES (1, 2, 'comment6 on test1');
INSERT INTO comments (post_id, user_id, body) VALUES (1, 1, 'comment7 on test1');
INSERT INTO comments (post_id, user_id, body) VALUES (1, 2, 'comment8 on test1');
INSERT INTO comments (post_id, user_id, body) VALUES (1, 1, 'comment9 on test1');
INSERT INTO comments (post_id, user_id, body) VALUES (1, 2, 'comment10 on test1');
INSERT INTO comments (post_id, user_id, body) VALUES (1, 1, 'comment11 on test1');
INSERT INTO comments (post_id, user_id, body) VALUES (1, 2, 'comment12 on test1');
INSERT INTO comments (post_id, user_id, body) VALUES (1, 1, 'comment13 on test1');
INSERT INTO comments (post_id, user_id, body) VALUES (1, 2, 'comment14 on test1');
INSERT INTO comments (post_id, user_id, body) VALUES (1, 1, 'comment15 on test1');
INSERT INTO comments (post_id, user_id, body) VALUES (1, 2, 'comment16 on test1');
INSERT INTO comments (post_id, user_id, body) VALUES (1, 1, 'comment17 on test1');
INSERT INTO comments (post_id, user_id, body) VALUES (1, 2, 'comment18 on test1');
INSERT INTO comments (post_id, user_id, body) VALUES (1, 1, 'comment19 on test1');
INSERT INTO comments (post_id, user_id, body) VALUES (1, 2, 'comment20 on test1');
INSERT INTO comments (post_id, user_id, body) VALUES (1, 1, 'comment21 on test1');
INSERT INTO comments (post_id, user_id, body) VALUES (1, 2, 'comment22 on test1');
INSERT INTO comments (post_id, user_id, body) VALUES (2, 1, 'comment1 on test2');
INSERT INTO comments (post_id, user_id, body) VALUES (2, 2, 'comment2 on test2');
INSERT INTO comments (post_id, user_id, body) VALUES (2, 1, 'comment3 on test2');
INSERT INTO comments (post_id, user_id, body) VALUES (2, 2, 'comment4 on test2');
