-- created by Olga Sinepalnikova
-- creating tables
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS groups (
                                      id uuid UNIQUE NOT NULL PRIMARY KEY ,
                                      name character(255)
);

CREATE TABLE  IF NOT EXISTS songs (
                                      id uuid PRIMARY KEY ,
                                      name character(255) NOT NULL ,
                                      group_id  uuid NOT NULL,
                                      FOREIGN KEY (group_id) REFERENCES groups(id)
);

CREATE TABLE IF NOT EXISTS song_details (
                                            song_id uuid NOT NULL ,
                                            FOREIGN KEY (song_id) REFERENCES songs(id),
                                            release_date date NOT NULL ,
                                            link character(511)
);

CREATE TABLE IF NOT EXISTS verses (
    song_id uuid NOT NULL ,
    FOREIGN KEY (song_id) REFERENCES  songs(id),
    count integer NOT NULL,
    text text NOT NULL
);

-- filling tables
INSERT INTO groups (id, name) VALUES
                                  (uuid_generate_v4(), 'Twenty One Pilot'),
                                  (uuid_generate_v4(), 'Linkin Park'),
                                  (uuid_generate_v4(), 'Сны саламандры');

INSERT INTO songs (id, name, group_id) VALUES
                                           (uuid_generate_v4(), 'Heathens', (SELECT id FROM groups WHERE name='Twenty One Pilot')),
                                           (uuid_generate_v4(), 'Heavy is the Crown', (SELECT id FROM groups WHERE name='Linkin Park')),
                                           (uuid_generate_v4(), 'Медвежья невеста', (SELECT id FROM groups WHERE name='Сны саламандры'));

SET datestyle TO 'DMY';
INSERT INTO song_details (song_id, release_date, link) VALUES
                                                                 ((SELECT id from songs where name='Heathens'), '16-06-2016', 'https://youtu.be/UprcpdwuwCg'),
                                                                 ((SELECT id FROM songs WHERE name='Heavy is the Crown'), '24-09-2024', 'https://www.youtube.com/watch?v=5FrhtahQiRc'),
                                                                 ((SELECT id FROM songs WHERE name='Медвежья невеста'), '13-09-2024', 'https://music.yandex.ru/album/33131669/track/130732838');

INSERT INTO verses (song_id, count, text) VALUES
    ((SELECT id FROM songs WHERE name='Heathens'), 1, 'All my friends are heathens, take it slow
Wait for them to ask you who you know
Please don''t make any sudden moves
You don''t know the half of the abuse'),
    ((SELECT id FROM songs WHERE name='Heathens'), 2, 'All my friends are heathens, take it slow
Wait for them to ask you who you know
Please don''t make any sudden moves
You don''t know the half of the abuse'),
    ((SELECT id FROM songs WHERE name='Heathens'), 3, 'Welcome to the room of people
Who have rooms of people that they loved one day
Docked away
Just because we check the guns at the door
Doesn''t mean our brains will change from hand grenades'),
    ((SELECT id FROM songs WHERE name='Heathens'), 4, 'You''ll never know the psychopath sitting next to you
You''ll never know the murderer sitting next to you
You''ll think, "How''d I get here, sitting next to you?"
But after all I''ve said, please don''t forget'),
    ((SELECT id FROM songs WHERE name='Heathens'), 5, 'All my friends are heathens, take it slow
Wait for them to ask you who you know
Please don''t make any sudden moves
You don''t know the half of the abuse'),
    ((SELECT id FROM songs WHERE name='Heathens'), 6, 'We don''t deal with outsiders very well
They say newcomers have a certain smell
You have trust issues, not to mention
They say they can smell your intentions'),
    ((SELECT id FROM songs WHERE name='Heathens'), 7, 'You''ll never know the freak show sitting next to you
You''ll have some weird people sitting next to you
You''ll think "How did I get here, sitting next to you?"
But after all I''ve said, please don''t forget'),
    ((SELECT id FROM songs WHERE name='Heathens'), 8, '(Watch it)
(Watch it)
All my friends are heathens, take it slow
Wait for them to ask you who you know
Please don''t make any sudden moves
You don''t know the half of the abuse'),
    ((SELECT id FROM songs WHERE name='Heathens'), 9, 'All my friends are heathens, take it slow (watch it)
Wait for them to ask you who you know (watch it)
Please all my friends are heathens, take it slow (watch it)
Wait for them to ask you who you know'),
    ((SELECT id FROM songs WHERE name='Heathens'), 10, ' Why''d you come? You knew you should have stayed
I tried to warn you just to stay away
And now they''re outside ready to bust
It looks like you might be one of us');

INSERT INTO verses (song_id, count, text) VALUES
    ((SELECT id FROM songs WHERE name='Heavy is the Crown'), 1, 'It''s pourin'' in, you''re laid on the floor again
One knock at the door and then
We both know how the story ends
You can''t win if your white flag''s out when the war begins
Aimin'' so high but swingin'' so low
Tryin'' to catch fire but feelin'' so cold
Hold it inside and hope it won''t show
I''m sayin'' it''s not, but inside, I know'),
    ((SELECT id FROM songs WHERE name='Heavy is the Crown'), 2, 'Today''s gonna be the day you notice
''Cause I''m tired of explainin'' what the joke is
This is what you asked for, heavy is the crown
Fire in the sunrise, ashes rainin'' down
Try to hold it in, but it keeps bleedin'' out
This is what you asked for, heavy is the
Heavy is the crown'),
    ((SELECT id FROM songs WHERE name='Heavy is the Crown'), 3, 'Turn to run, now look what it''s become
Outnumbered, ten to one
Back thеn should''ve bit your tongue
''Cause thеre''s no turnin'' back this path once it''s begun
You''re already on that list
Say you don''t want what you can''t resist
Wavin'' that sword when the pen won''t miss
Watch it all fallin'' apart like this'),
    ((SELECT id FROM songs WHERE name='Heavy is the Crown'), 4, 'This is what you asked for, heavy is the crown
Fire in the sunrise, ashes rainin'' down
Try to hold it in, but it keeps bleedin'' out
This is what you asked for, heavy is the
Heavy is the crown
Today''s gonna be the day you notice
''Cause I''m tired of explainin'' what the joke is
This is what you asked for'),
    ((SELECT id FROM songs WHERE name='Heavy is the Crown'), 5, 'This is what you asked for, heavy is the crown
Fire in the sunrise, ashes rainin'' down
Try to hold it in but it keeps bleedin'' out
This is what you asked for, heavy is the
Heavy is the crown
Heavy is the crown
Heavy is the, heavy is the crown');

INSERT INTO  verses (song_id, count, text) VALUES
    ((SELECT id FROM songs WHERE name='Медвежья невеста'), 1, 'Слышишь, девочка, слышишь, милая
Лютый зверь как ревёт с тоски
Жизнь моя без тебя постылая
Разлетелись то лепестки'),
    ((SELECT id FROM songs WHERE name='Медвежья невеста'), 2, 'Как пускала венок по реченьке
С той поры я ночей не сплю
Забывай своё, человечье
Да иди со мной к алтарю'),
    ((SELECT id FROM songs WHERE name='Медвежья невеста'), 3, 'Выйдешь из дому мне навстречу
Нежно к шерсти прильнёшь щекой
Забывай своё, человечье
Будь со мной!'),
    ((SELECT id FROM songs WHERE name='Медвежья невеста'), 4, 'Мне одежда - медвежья шкура
Наше царство - дремучий лес
Плачет мать: «Ну куда ж ты, дура?!»
Не иначе, попутал бес'),
    ((SELECT id FROM songs WHERE name='Медвежья невеста'), 5, 'Выйдешь из дому мне навстречу
Нежно к шерсти прильнёшь щекой
Забывай своё, человечье
Будь со мной!'),
    ((SELECT id FROM songs WHERE name='Медвежья невеста'), 6, 'Выйдешь из дому мне навстречу
Нежно к шерсти прильнёшь щекой
Забывай своё, человечье
Будь со мной!'),
    ((SELECT id FROM songs WHERE name='Медвежья невеста'), 7, 'Выйдешь из дому мне навстречу
Нежно к шерсти прильнёшь щекой
Забывай своё, человечье
Будь со мной!');