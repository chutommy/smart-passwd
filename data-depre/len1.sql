SELECT COUNT(*)
FROM words
WHERE LENGTH(words.word) = $1
AND words.category_id = $2
LIMIT 1;
