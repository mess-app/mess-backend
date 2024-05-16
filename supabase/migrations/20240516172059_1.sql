create type "public"."ConnectionStatus" as enum ('connected', 'declined', 'pending');

create table "public"."connections" (
    "id" uuid not null default gen_random_uuid(),
    "created_at" timestamp with time zone not null default now(),
    "pioneer_id" uuid not null default gen_random_uuid(),
    "recipient_id" uuid not null default gen_random_uuid(),
    "status" "ConnectionStatus" not null default 'pending'::"ConnectionStatus"
);


alter table "public"."connections" enable row level security;

create table "public"."group_connections" (
    "created_at" timestamp with time zone not null default now(),
    "user_id" uuid not null,
    "inviter" uuid not null,
    "group_id" uuid not null,
    "id" uuid not null default gen_random_uuid()
);


alter table "public"."group_connections" enable row level security;

create table "public"."groups" (
    "id" uuid not null default gen_random_uuid(),
    "created_at" timestamp with time zone not null default now(),
    "name" text not null,
    "creator" uuid not null,
    "cover_url" text
);


alter table "public"."groups" enable row level security;

alter table "public"."profile" add column "username" text not null;

alter table "public"."profile" alter column "user_id" set not null;

CREATE UNIQUE INDEX connections_pkey ON public.connections USING btree (id);

CREATE UNIQUE INDEX group_connections_pkey ON public.group_connections USING btree (id);

CREATE UNIQUE INDEX groups_pkey ON public.groups USING btree (id);

CREATE UNIQUE INDEX profile_username_key ON public.profile USING btree (username);

alter table "public"."connections" add constraint "connections_pkey" PRIMARY KEY using index "connections_pkey";

alter table "public"."group_connections" add constraint "group_connections_pkey" PRIMARY KEY using index "group_connections_pkey";

alter table "public"."groups" add constraint "groups_pkey" PRIMARY KEY using index "groups_pkey";

alter table "public"."connections" add constraint "public_connections_pioneer_id_fkey" FOREIGN KEY (pioneer_id) REFERENCES profile(id) ON DELETE CASCADE not valid;

alter table "public"."connections" validate constraint "public_connections_pioneer_id_fkey";

alter table "public"."connections" add constraint "public_connections_recipient_id_fkey" FOREIGN KEY (recipient_id) REFERENCES profile(id) ON DELETE CASCADE not valid;

alter table "public"."connections" validate constraint "public_connections_recipient_id_fkey";

alter table "public"."group_connections" add constraint "public_group_connections_group_id_fkey" FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE not valid;

alter table "public"."group_connections" validate constraint "public_group_connections_group_id_fkey";

alter table "public"."group_connections" add constraint "public_group_connections_inviter_fkey" FOREIGN KEY (inviter) REFERENCES auth.users(id) not valid;

alter table "public"."group_connections" validate constraint "public_group_connections_inviter_fkey";

alter table "public"."group_connections" add constraint "public_group_connections_user_id_fkey" FOREIGN KEY (user_id) REFERENCES auth.users(id) ON DELETE CASCADE not valid;

alter table "public"."group_connections" validate constraint "public_group_connections_user_id_fkey";

alter table "public"."groups" add constraint "public_groups_creator_fkey" FOREIGN KEY (creator) REFERENCES auth.users(id) not valid;

alter table "public"."groups" validate constraint "public_groups_creator_fkey";

alter table "public"."profile" add constraint "profile_username_check" CHECK ((username ~ '^[a-z][a-z0-9_]*[a-z0-9]$'::text)) not valid;

alter table "public"."profile" validate constraint "profile_username_check";

alter table "public"."profile" add constraint "profile_username_key" UNIQUE using index "profile_username_key";

set check_function_bodies = off;

CREATE OR REPLACE FUNCTION public.current_profile_id()
 RETURNS uuid
 LANGUAGE sql
AS $function$SELECT id FROM profile WHERE user_id = auth.uid();$function$
;

grant delete on table "public"."connections" to "anon";

grant insert on table "public"."connections" to "anon";

grant references on table "public"."connections" to "anon";

grant select on table "public"."connections" to "anon";

grant trigger on table "public"."connections" to "anon";

grant truncate on table "public"."connections" to "anon";

grant update on table "public"."connections" to "anon";

grant delete on table "public"."connections" to "authenticated";

grant insert on table "public"."connections" to "authenticated";

grant references on table "public"."connections" to "authenticated";

grant select on table "public"."connections" to "authenticated";

grant trigger on table "public"."connections" to "authenticated";

grant truncate on table "public"."connections" to "authenticated";

grant update on table "public"."connections" to "authenticated";

grant delete on table "public"."connections" to "service_role";

grant insert on table "public"."connections" to "service_role";

grant references on table "public"."connections" to "service_role";

grant select on table "public"."connections" to "service_role";

grant trigger on table "public"."connections" to "service_role";

grant truncate on table "public"."connections" to "service_role";

grant update on table "public"."connections" to "service_role";

grant delete on table "public"."group_connections" to "anon";

grant insert on table "public"."group_connections" to "anon";

grant references on table "public"."group_connections" to "anon";

grant select on table "public"."group_connections" to "anon";

grant trigger on table "public"."group_connections" to "anon";

grant truncate on table "public"."group_connections" to "anon";

grant update on table "public"."group_connections" to "anon";

grant delete on table "public"."group_connections" to "authenticated";

grant insert on table "public"."group_connections" to "authenticated";

grant references on table "public"."group_connections" to "authenticated";

grant select on table "public"."group_connections" to "authenticated";

grant trigger on table "public"."group_connections" to "authenticated";

grant truncate on table "public"."group_connections" to "authenticated";

grant update on table "public"."group_connections" to "authenticated";

grant delete on table "public"."group_connections" to "service_role";

grant insert on table "public"."group_connections" to "service_role";

grant references on table "public"."group_connections" to "service_role";

grant select on table "public"."group_connections" to "service_role";

grant trigger on table "public"."group_connections" to "service_role";

grant truncate on table "public"."group_connections" to "service_role";

grant update on table "public"."group_connections" to "service_role";

grant delete on table "public"."groups" to "anon";

grant insert on table "public"."groups" to "anon";

grant references on table "public"."groups" to "anon";

grant select on table "public"."groups" to "anon";

grant trigger on table "public"."groups" to "anon";

grant truncate on table "public"."groups" to "anon";

grant update on table "public"."groups" to "anon";

grant delete on table "public"."groups" to "authenticated";

grant insert on table "public"."groups" to "authenticated";

grant references on table "public"."groups" to "authenticated";

grant select on table "public"."groups" to "authenticated";

grant trigger on table "public"."groups" to "authenticated";

grant truncate on table "public"."groups" to "authenticated";

grant update on table "public"."groups" to "authenticated";

grant delete on table "public"."groups" to "service_role";

grant insert on table "public"."groups" to "service_role";

grant references on table "public"."groups" to "service_role";

grant select on table "public"."groups" to "service_role";

grant trigger on table "public"."groups" to "service_role";

grant truncate on table "public"."groups" to "service_role";

grant update on table "public"."groups" to "service_role";

create policy "Allow the ones who are part of it"
on "public"."connections"
as permissive
for all
to authenticated
using (((pioneer_id = current_profile_id()) OR (recipient_id = current_profile_id())))
with check ((pioneer_id <> recipient_id));



