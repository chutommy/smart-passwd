SELECT words.word
FROM words
WHERE LENGTH(words.word) = $1
AND words.category_id = $2
OFFSET $3
LIMIT 1;
