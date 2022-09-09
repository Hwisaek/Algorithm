SELECT firstName, lastName, city, state
FROM person p
         LEFT JOIN address a ON p.personID = a.personId