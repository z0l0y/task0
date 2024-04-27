-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS "user";
CREATE TABLE "user" (
                        id SERIAL PRIMARY KEY,
                        name varchar(50),
                        nick_name varchar(150),
                        password varchar(100)
);

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO "user" (name, nick_name, password)
VALUES ('admin', '', '123456');
INSERT INTO "user" (name, nick_name, password)
VALUES ('admin1', 'admin', '123456');