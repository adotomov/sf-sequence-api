SELECT
  st.ID,
  st.EMAIL_SUBJECT,
  st.EMAIL_CONTENT
FROM sequence s
  LEFT JOIN step st ON st.SEQUENCE_ID = s.ID
WHERE s.ID = $1
