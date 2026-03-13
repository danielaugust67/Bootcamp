--a
select p.*
from t_policy p
join t_client c on p.client_number = c.client_number
where p.policy_submit_date > '2018-01-15' 
  and extract(month from c.birth_date) = 9;

--b
select p.*
from t_policy p
join t_agent a on p.agent_code = a.agent_code
where p.policy_status = 'INFORCE' 
  and a.agent_office = 'JAKARTA';

--c
update t_agent
set basic_commission = subquery.hasil_hitung
from (
    select agent_code, max((commission / premium) * 100) as hasil_hitung
    from t_policy
    group by agent_code
) as subquery
where t_agent.agent_code = subquery.agent_code;

select 
    agent_code, 
    agent_name, 
    basic_commission 
from t_agent
order by agent_code;

--d
update t_policy
set policy_due_date = cast(
    date_trunc('month', policy_submit_date + 30) + interval '1 month' - interval '1 day' 
    as date
);

--e
select 
    a.agent_code, 
    a.agent_name, 
    p.policy_number, 
    p.premium, 
    p.discount,
    (p.premium - (p.premium * (p.discount / 100.0))) as premium_setelah_diskon
from t_policy p
join t_agent a on p.agent_code = a.agent_code
where (p.premium - (p.premium * (p.discount / 100.0))) < 1000000
order by premium_setelah_diskon asc;