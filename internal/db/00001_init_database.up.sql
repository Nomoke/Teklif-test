CREATE TYPE offer_type AS ENUM ('gift', 'extra_commission', 'discount', 'special_price');
CREATE TYPE offer_status AS ENUM ('pending', 'approved', 'expired');

CREATE TABLE special_offers (
    id                  UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title               TEXT NOT NULL,
    description         TEXT NOT NULL,
    type                offer_type NOT NULL,
    expiry_date         TIMESTAMP WITH TIME ZONE NOT NULL,
    commission_agent    NUMERIC(10,2) CHECK (commission_agent >= 0), 
    commission_agency   NUMERIC(10,2) CHECK (commission_agency >= 0),
    discount            NUMERIC(5,2)  CHECK (discount BETWEEN 0 AND 100), 
    price_installment   NUMERIC(15,2) CHECK (price_installment >= 0),
    price_cash          NUMERIC(15,2) CHECK (price_cash >= 0),
    status              offer_status NOT NULL DEFAULT 'pending',
    media               TEXT[],
    created_at          TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
);

CREATE TABLE offer_properties (
    offer_id     UUID REFERENCES special_offers(id) ON DELETE CASCADE,
    property_id  UUID NOT NULL,
    PRIMARY KEY (offer_id, property_id)
);