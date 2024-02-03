INSERT INTO discord_message (event_clan_id, bot_instance_id)
SELECT ec.id,
    bi.id
FROM event_clan ec
    JOIN subscription_clan sc ON sc.clan_id = ec.clan_id
    AND sc.region = ec.region
    AND sc.is_disabled = 0
    JOIN bot_instance bi ON bi.id = sc.instance_id
    AND bi.type = 'clan'
    LEFT JOIN discord_message dm ON dm.event_clan_id = ec.id
    AND dm.bot_instance_id = bi.id
WHERE ec.is_processed = 0
    AND dm.is_processed IS NULL
