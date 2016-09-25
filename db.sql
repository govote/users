CREATE DATABASE IF NOT EXISTS deputados_em_foco;

DROP TABLE IF EXISTS `user`;

CREATE TABLE user (
    id          VARCHAR(255) NOT NULL,
    name        VARCHAR(255) NOT NULL,
    email       VARCHAR(150) NOT NULL,
    photoUrl    VARCHAR(255) NULL,
    facebookId  VARCHAR(255) NULL,
    googleId    VARCHAR(255) NULL,
    createdAt   TIMESTAMP    NOT NULL,
    updatedAt   TIMESTAMP    NOT NULL,

    PRIMARY KEY ( id ),

    INDEX idx_user_email ( email ), 

    UNIQUE INDEX udx_user_email ( email ),
    UNIQUE INDEX udx_user_facebookid ( facebookId)
);
