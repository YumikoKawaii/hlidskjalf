create table pawn.synthetic_student_profiles
    on cluster hlidskjalf
(
    id String,
    name String,
    age UInt32,
    sex String,
    major String,
    year UInt32,
    gpa Float32,
    hobbies Array(String),
    country String,
    province String
    )
    engine = ReplicatedReplacingMergeTree
    order by id
    settings index_granularity = 8192;

create table pawn.synthetic_student_profiles_distributed
    on cluster hlidskjalf
as pawn.synthetic_student_profiles
    engine = Distributed('hlidskjalf','pawn','synthetic_student_profiles',murmurHash3_64(id));