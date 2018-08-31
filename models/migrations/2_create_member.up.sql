CREATE TABLE member (
    person_id INT(10) UNSIGNED NOT NULL,
    member_id INT(10) UNSIGNED NOT NULL,
    created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (person_id, member_id),
    KEY fk_person (person_id),
    CONSTRAINT fk_person FOREIGN KEY (person_id) REFERENCES person (id) ON UPDATE CASCADE ON DELETE CASCADE,
    KEY fk_member (member_id),
    CONSTRAINT fk_member FOREIGN KEY (member_id) REFERENCES person (id) ON UPDATE CASCADE ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE utf8_unicode_ci;
