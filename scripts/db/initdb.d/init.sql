-- Setup database and user
CREATE DATABASE go_project_layout;
CREATE USER go_project_layout WITH ENCRYPTED PASSWORD '1qazxsw23edc';

-- Setup tables
\c go_project_layout

CREATE OR REPLACE FUNCTION public.updated_time()
 RETURNS trigger
 LANGUAGE plpgsql
AS $function$
BEGIN
    NEW.updated_at = now();
RETURN NEW;
END;
$function$
;

CREATE TABLE public.customers (
    id bigserial NOT NULL,
    name varchar(255) NOT NULL,
    created_at timestamptz NULL DEFAULT now(),
    updated_at timestamptz NULL,
    PRIMARY KEY (id)
);

CREATE TABLE public.loans (
    id bigserial NOT NULL,
    customer_id bigint NOT NULL,
    repayment_type varchar(255) NOT NULL,
    amount decimal NOT NULL,
    term int  NOT NULL,
    state varchar(100) NOT NULL,
    created_at timestamptz NULL DEFAULT now(),
    updated_at timestamptz NULL,
    PRIMARY KEY(id),
    CONSTRAINT fk_customers FOREIGN KEY(customer_id) REFERENCES customers(id) ON DELETE SET NULL
);

CREATE TABLE public.repayments (
    id bigserial NOT NULL,
    loan_id bigint NOT NULL,
    amount decimal NOT NULL,
    due_date timestamptz NOT NULL,
    state varchar(100) NOT NULL,
    created_at timestamptz NULL DEFAULT now(),
    updated_at timestamptz NULL,
    PRIMARY KEY(id),
    CONSTRAINT fk_loans FOREIGN KEY(loan_id) REFERENCES loans(id) ON DELETE SET NULL
);

CREATE TABLE public.repayment_logs (
    id bigserial NOT NULL,
    loan_id bigint NOT NULL,
    amount decimal NOT NULL,
    created_at timestamptz NULL DEFAULT now(),
    PRIMARY KEY(id),
    CONSTRAINT fk_loans FOREIGN KEY(loan_id) REFERENCES loans(id) ON DELETE SET NULL
);

CREATE TABLE public.repayment_prepaid (
    id bigserial NOT NULL,
    loan_id bigint NOT NULL,
    amount decimal NOT NULL,
    created_at timestamptz NULL DEFAULT now(),
    deleted_at timestamptz NULL,
    PRIMARY KEY(id),
    CONSTRAINT fk_loans FOREIGN KEY(loan_id) REFERENCES loans(id) ON DELETE SET NULL
);

--  Setup seed data

INSERT INTO public.customers ("name") VALUES('aspire-code-challenge-customer');

-- Grant permission
GRANT CONNECT ON DATABASE go_project_layout TO go_project_layout;
GRANT USAGE ON SCHEMA public TO go_project_layout;
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA public TO go_project_layout;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO go_project_layout;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO go_project_layout;
