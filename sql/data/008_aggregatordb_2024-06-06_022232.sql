-- +goose Up

INSERT INTO
    public.users (
        id,
        created_at,
        updated_at,
        username,
        email,
        password,
        api_key
    )
VALUES (
        1,
        '2024-06-05 19:40:34.669633',
        '2024-06-05 19:40:34.669633',
        'John',
        'nsltharaka@gmail.com',
        '$2a$10$mmV/yBkDu.zezikwR7VE8.3Cp14Yp7q7oEDncxA1F/hhJsL/wwdgy',
        'dc83b49e9f9c6fb2e9d37541963205edfc638a22ad7603d4b2b883bd94b1ad9a'
    );

INSERT INTO
    public.topics (
        id,
        name,
        img_url,
        updated_at,
        created_by
    )
VALUES (
        '8332b040-d20d-4862-a7a6-9f94c036f07e',
        'History',
        'https://history.vcu.edu/media/history/hero/HistHeroCollage21x9.jpg',
        '2024-06-05 19:53:02.549908',
        1
    ),
    (
        'a66a2ed5-c1e0-456d-a83c-c519efb54946',
        'Go Programming Language',
        'https://miro.medium.com/v2/resize:fit:1400/1*yWS1bEU6aoaXYHCUm2L2Rg.png',
        '2024-06-05 20:00:51.938621',
        1
    ),
    (
        '991483ee-5e57-425d-893f-569a339e2f7a',
        'Central bank of sri lanka feeds',
        'https://ichef.bbci.co.uk/news/976/cpsprodpb/2B90/production/_124025111_whatsappimage2022-04-05at10.42.24am.jpg',
        '2024-06-05 20:09:11.167173',
        1
    );

INSERT INTO
    public.feeds (
        id,
        created_at,
        updated_at,
        url
    )
VALUES (
        '43a4c38a-9855-4c62-a9c2-77953a44c3d6',
        '2024-06-05 20:00:51.941765',
        '2024-06-06 01:39:23.897926',
        'https://changelog.com/gotime/feed'
    ),
    (
        '64d174c7-ab5d-49fe-83bc-633024fcec25',
        '2024-06-05 19:53:02.556207',
        '2024-06-06 01:39:23.898949',
        'https://historytoday.com/feed/rss.xml'
    ),
    (
        '302cb9e8-b327-4bff-9731-a3da2195c06c',
        '2024-06-05 20:09:11.170534',
        '2024-06-06 01:39:24.008293',
        'https://www.cbsl.gov.lk/en/press/notices'
    ),
    (
        '19bfd476-4b05-4709-9e7d-89aed71cc711',
        '2024-06-05 20:09:11.173698',
        '2024-06-06 01:39:24.02963',
        'https://www.cbsl.gov.lk/en/rss.xml'
    ),
    (
        '99d07f44-a481-4966-86bd-619cd4d518f1',
        '2024-06-05 20:00:51.943676',
        '2024-06-06 01:39:24.030523',
        'https://go.dev/blog/feed.atom'
    ),
    (
        '6e363ac8-51f0-4f6f-9fea-d1af7fd955b3',
        '2024-06-05 20:09:11.177225',
        '2024-06-06 01:39:24.041482',
        'https://www.cbsl.gov.lk/en/statistics/economic-indicators/meirss'
    ),
    (
        '4205dbb6-dd9d-4c6e-9557-fcd6102b15e1',
        '2024-06-05 20:00:51.944988',
        '2024-06-06 01:39:24.041575',
        'https://appliedgo.net/index.xml'
    );

INSERT INTO
    public.topic_contains_feed (topic_id, feed_id, user_id)
VALUES (
        '8332b040-d20d-4862-a7a6-9f94c036f07e',
        '64d174c7-ab5d-49fe-83bc-633024fcec25',
        1
    ),
    (
        'a66a2ed5-c1e0-456d-a83c-c519efb54946',
        '43a4c38a-9855-4c62-a9c2-77953a44c3d6',
        1
    ),
    (
        'a66a2ed5-c1e0-456d-a83c-c519efb54946',
        '99d07f44-a481-4966-86bd-619cd4d518f1',
        1
    ),
    (
        'a66a2ed5-c1e0-456d-a83c-c519efb54946',
        '4205dbb6-dd9d-4c6e-9557-fcd6102b15e1',
        1
    ),
    (
        '991483ee-5e57-425d-893f-569a339e2f7a',
        '302cb9e8-b327-4bff-9731-a3da2195c06c',
        1
    ),
    (
        '991483ee-5e57-425d-893f-569a339e2f7a',
        '19bfd476-4b05-4709-9e7d-89aed71cc711',
        1
    ),
    (
        '991483ee-5e57-425d-893f-569a339e2f7a',
        '6e363ac8-51f0-4f6f-9fea-d1af7fd955b3',
        1
    );

INSERT INTO
    public.user_follows_topic (user_id, topic_id)
VALUES (
        1,
        '8332b040-d20d-4862-a7a6-9f94c036f07e'
    ),
    (
        1,
        'a66a2ed5-c1e0-456d-a83c-c519efb54946'
    ),
    (
        1,
        '991483ee-5e57-425d-893f-569a339e2f7a'
    );