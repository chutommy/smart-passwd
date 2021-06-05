async function randomWord(length) {
  const [SQL, buf] = await Promise.all([sqlPromise, dataPromise]);
  const db = new SQL.Database(new Uint8Array(buf));

  const res = db.exec(
    `SELECT word
     FROM words
     WHERE length(word) = $length
     ORDER BY RANDOM()
     LIMIT 1`,
    {"$length": length}
  );

  return res[0].values[0][0];
}
