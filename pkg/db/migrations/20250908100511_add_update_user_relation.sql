-- +goose Up
-- +goose StatementBegin
ALTER TABLE "patients" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "doctors" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "patients" DROP CONSTRAINT "patients_user_id_fkey";
ALTER TABLE "doctors" DROP CONSTRAINT "doctors_user_id_fkey";
-- +goose StatementEnd
