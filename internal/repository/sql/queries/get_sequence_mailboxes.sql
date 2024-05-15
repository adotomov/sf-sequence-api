SELECT
  m.ID,
  m.EMAIL,
  m.MAX_LIMIT
FROM sequence s
  LEFT JOIN sequence_mailbox sm ON sm.SEQUENCE_ID = s.ID
  LEFT JOIN mailbox m ON m.ID = sm.MAILBOX_ID
WHERE s.ID = $1
