sets all values to defaults

The {{aka}} command sets all cached variables to their initial values. Any variable name from {{cmd "conf"}} will be used to initialize if defined.  Otherwise, the following hard-coded package globals will be used instead:

    duration    - {{dduration}}
    interval    - {{dinterval}}
    warn        - {{dwarn}}
    prefix      - {{dprefix}}
    prefixwarn  - {{dprefixwarn}}

