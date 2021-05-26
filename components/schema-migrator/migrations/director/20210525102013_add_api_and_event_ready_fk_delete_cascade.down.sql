BEGIN;

ALTER TABLE api_definitions DROP CONSTRAINT api_definitions_bundles_ready_fk;
ALTER TABLE api_definitions
    ADD CONSTRAINT api_definitions_bundles_ready_fk
        FOREIGN KEY (app_id, ready) REFERENCES applications (id, ready) ON UPDATE CASCADE;

ALTER TABLE event_api_definitions DROP CONSTRAINT event_api_definitions_bundles_ready_fk;
ALTER TABLE event_api_definitions
    ADD CONSTRAINT event_api_definitions_bundles_ready_fk
        FOREIGN KEY (app_id, ready) REFERENCES applications (id, ready) ON UPDATE CASCADE;

COMMIT;
