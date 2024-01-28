SELECT dm.id,
    dm.is_processed,
    dm.event_clan_id,
    dm.bot_instance_id,
    ec.`time`,
    ec.`type`,
    ec.region,
    ec.clan_id,
    ec.account_id,
    bi.channel_id
FROM discord_message dm
    JOIN event_clan ec ON dm.event_clan_id = ec.id
    JOIN bot_instance bi ON dm.bot_instance_id = bi.id
WHERE dm.is_processed = 0
LIMIT 1
