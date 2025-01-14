CREATE TABLE quotes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,  -- Unique identifier for each record
    code TEXT NOT NULL,                    -- Currency code (e.g., USD)
    codein TEXT NOT NULL,                  -- Target currency code (e.g., BRL)
    name TEXT NOT NULL,                    -- Full name of the currency pair
    high TEXT NOT NULL,                    -- Highest value in the current period
    low TEXT NOT NULL,                     -- Lowest value in the current period
    var_bid TEXT NOT NULL,                 -- Variation in bid value
    pct_change TEXT NOT NULL,              -- Percentage change in value
    bid TEXT NOT NULL,                     -- Current bid value
    ask TEXT NOT NULL,                     -- Current ask value
    timestamp TEXT NOT NULL,               -- Unix timestamp of the data
    create_date TEXT NOT NULL              -- ISO 8601 date and time when the data was created
);
