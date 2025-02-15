CREATE TABLE special_offers (
    id                  UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title               TEXT NOT NULL,
    description         TEXT NOT NULL,
    type                TEXT NOT NULL,
    expiry_date         TIMESTAMP WITH TIME ZONE NOT NULL,
    commission_agent    NUMERIC(10,2), 
    commission_agency   NUMERIC(10,2),
    discount            NUMERIC(5,2), 
    price_installment   NUMERIC(15,2),
    price_cash          NUMERIC(15,2),
    status              TEXT NOT NULL DEFAULT 'pending',
    media               TEXT[],
    created_at          TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE offer_properties (
    offer_id     UUID REFERENCES special_offers(id) ON DELETE CASCADE,
    property_id  UUID NOT NULL,
    PRIMARY KEY (offer_id, property_id)
);