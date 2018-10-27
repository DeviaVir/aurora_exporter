package main

import (
	"regexp"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

var counters = map[string]*prometheus.Desc{
	"async_tasks_completed": newDesc(
		"async_tasks", "completed",
		"Number of completed async tasks.",
	),
	"assigner_evaluated_offers": newDesc(
		"assigner_evaluated", "offers",
		"Number of offers evaluated before a task was assigned.",
	),
	"assigner_launch_failures": newDesc(
		"assigner_launch", "failures",
		"Number of task launch failures.",
	),
	"cron_job_launch_failures": newDesc(
		"cron", "job_launch_failures",
		"Scheduled job failures total.",
	),
	"cron_jobs_loaded": newDesc(
		"cron", "jobs_loaded",
		"Cron scheduler loaded.",
	),
	"dropped_update_events": newDesc(
		"", "dropped_update_events",
		"Number of dropped update events.",
	),
	"framework_registered": newDesc(
		"", "framework_registered",
		"Framework registered total.",
	),
	"gc_executor_tasks_lost": newDesc(
		"", "gc_executor_tasks_lost",
		"Lost garbage collection task total.",
	),
	"http_200_responses_events": newDesc(
		"http_200", "responses_events",
		"Number of HTTP 200 status response events.",
	),
	"http_200_responses_nanos_total": newDesc(
		"http_200", "responses_nanos_total",
		"Timed HTTP 200 status response total.",
	),
	"http_500_responses_events": newDesc(
		"http_500", "responses_events",
		"Number of HTTP 500 status responses sent by the scheduler total.",
	),
	"job_update_delete_errors": newDesc(
		"job_update", "delete_errors",
		"Failed delete job total.",
	),
	"job_update_recovery_errors": newDesc(
		"job_update", "recovery_errors",
		"Failed resume job updates total.",
	),
	"job_update_state_change_errors": newDesc(
		"job_update", "state_change_errors",
		"State change errors total.",
	),
	"jvm_class_loaded_count": newDesc(
		"jvm", "class_loaded_count",
		"Number of classes loaded.",
	),
	"jvm_class_total_loaded_count": newDesc(
		"jvm", "class_total_loaded_count",
		"Total number of classes loaded.",
	),
	"jvm_class_unloaded_count": newDesc(
		"jvm", "class_unloaded_count",
		"Number of classes unloaded.",
	),
	"jvm_gc_PS_MarkSweep_collection_count": newDesc(
		"jvm", "gc_ps_marksweep_collection_count",
		"Parallel mark and sweep collection run total.",
	),
	"jvm_gc_PS_Scavenge_collection_count": newDesc(
		"jvm", "gc_ps_scavenge_collection_count",
		"Parallel scavenge collector runs total.",
	),
	"jvm_gc_collection_count": newDesc(
		"jvm", "gc_collection_count",
		"Garbage collection total.",
	),
	"jvm_memory_heap_mb_max": newDesc(
		"jvm", "memory_heap_mb_max",
		"Maximum heap memory.",
	),
	"jvm_memory_max_mb": newDesc(
		"jvm", "memory_max_mb",
		"Maximum amount of memory that the Java virtual machine will attempt to use.",
	),
	"jvm_memory_mb_total": newDesc(
		"jvm", "memory_mb_total",
		"Total amount of memory in the Java virtual machine.",
	),
	"jvm_memory_non_heap_mb_max": newDesc(
		"jvm", "memory_non_heap_mb_max",
		"Max non heap memory used in MB.",
	),
	"jvm_threads_peak": newDesc(
		"jvm", "threads_peak",
		"Peak thread count.",
	),
	"jvm_threads_started": newDesc(
		"jvm", "threads_started",
		"Total threads started.",
	),
	"jvm_uptime_secs": newDesc(
		"jvm", "uptime_secs",
		"Number of seconds the JVM process has been running.",
	),
	"log_storage_write_lock_wait_events": newDesc(
		"log_storage", "write_lock_wait_events", "",
	),
	"log_storage_write_lock_wait_ns_total": newDesc(
		"log_storage", "write_lock_wait_ns_total", "",
	),
	"offer_accept_races": newDesc(
		"", "offer_accept_races",
		"Accepted offer no longer exists in offer queue.",
	),
	"preemptor_missing_attributes": newDesc(
		"preemptor", "missing_attributes",
		"Preemptor missing attributes.",
	),
	"preemptor_slot_search_attempts_for_non_prod": newDesc(
		"preemptor", "slot_search_attempts_for_non_prod",
		"Preemptor slot search attempts for non prod.",
	),
	"preemptor_slot_search_attempts_for_prod": newDesc(
		"preemptor", "slot_search_attempts_for_prod",
		"Preemptor slot search attempts for prod.",
	),
	"preemptor_slot_search_failed_for_non_prod": newDesc(
		"preemptor", "slot_search_failed_for_non_prod",
		"Preemptor slot search failed for non prod.",
	),
	"preemptor_slot_search_failed_for_prod": newDesc(
		"preemptor", "slot_search_failed_for_prod",
		"Preemptor slot search failed for prod.",
	),
	"preemptor_slot_search_successful_for_non_prod": newDesc(
		"preemptor", "slot_search_successful_for_non_prod",
		"Preemptor slot search successful for non prod.",
	),
	"preemptor_slot_search_successful_for_prod": newDesc(
		"preemptor", "slot_search_successful_for_prod",
		"Preemptor slot search successful for prod.",
	),
	"preemptor_slot_validation_failed": newDesc(
		"preemptor", "slot_validation_failed",
		"Preemptor slot validation failed.",
	),
	"preemptor_slot_validation_successful": newDesc(
		"preemptor", "slot_validation_successful",
		"Preemptor slot validation successful.",
	),
	"preemptor_task_processor_runs": newDesc(
		"preemptor", "task_processor_runs",
		"Preemptor task processor runs.",
	),
	"preemptor_tasks_preempted_non_prod": newDesc(
		"preemptor", "tasks_preempted_non_prod",
		"Preemptor tasks preempted non prod.",
	),
	"preemptor_tasks_preempted_prod": newDesc(
		"preemptor", "tasks_preempted_prod",
		"Preemptor tasks preempted prod.",
	),
	"process_max_fd_count": newDesc(
		"process", "max_fd_count",
		"Max open file descriptors.",
	),
	"process_open_fd_count": newDesc(
		"process", "open_fd_count",
		"Open file descriptors in use.",
	),
	"schedule_attempts_failed": newDesc(
		"schedule", "attempts_failed",
		"Number of failed attempts to schedule tasks.",
	),
	"schedule_attempts_fired": newDesc(
		"schedule", "attempts_fired",
		"Number of attempts to schedule tasks.",
	),
	"schedule_attempts_no_match": newDesc(
		"schedule", "attempts_no_match",
		"Number of task which could not be scheduled.",
	),
	"scheduled_task_penalty_events": newDesc(
		"scheduled_task", "penalty_events",
		"Number of tasks with penalty events.",
	),
	"scheduled_task_penalty_ms_total": newDesc(
		"scheduled_task", "penalty_ms_total",
		"Timed penalty events total.",
	),
	"scheduler_backup_failed": newDesc(
		"scheduler", "backup_failed",
		"Number of failed storage backup.",
	),
	"scheduler_backup_success": newDesc(
		"scheduler", "backup_success",
		"Number successful storage backup.",
	),
	"scheduler_driver_kill_failures": newDesc(
		"scheduler", "driver_kill_failures",
		"Number of driver kill failures.",
	),
	"scheduler_gc_insufficient_offers": newDesc(
		"scheduler_gc", "insufficient_offers",
		"Number off resource offer that was too small for a garbage collection task.",
	),
	"scheduler_gc_offers_consumed": newDesc(
		"scheduler_gc", "offers_consumed",
		"Number of resource offers consumed for garbage collection tasks.",
	),
	"scheduler_gc_tasks_created": newDesc(
		"scheduler_gc", "tasks_created",
		"Number of garbage collection tasks created.",
	),
	"scheduler_log_bad_frames_read": newDesc(
		"scheduler_log", "bad_frames_read",
		"Number of bad frames read.",
	),
	"scheduler_log_bytes_read": newDesc(
		"scheduler_log", "bytes_read",
		"Number of resource bytes read.",
	),
	"scheduler_log_deflated_entries_read": newDesc(
		"scheduler_log", "deflated_entries_read",
		"Number of deflated entries.",
	),
	"scheduler_log_entries_read": newDesc(
		"scheduler_log", "entries_read",
		"Number of entries read.",
	),
	"scheduler_log_entries_written": newDesc(
		"scheduler_log", "entries_written",
		"Number of entries written.",
	),
	"scheduler_log_native_append_events": newDesc(
		"scheduler_log", "native_append_events",
		"Number of append operations total.",
	),
	"scheduler_log_native_append_failures": newDesc(
		"scheduler_log", "native_append_failures",
		"Number of append failures total.",
	),
	"scheduler_log_native_append_nanos_total": newDesc(
		"scheduler_log", "native_append_nanos_total",
		"Timed append operations total.",
	),
	"scheduler_log_native_append_nanos_total_per_sec": newDesc(
		"scheduler_log", "native_append_nanos_total_per_sec",
		"Timed append operations total per second.",
	),
	"scheduler_log_native_append_timeouts": newDesc(
		"scheduler_log", "native_append_timeouts",
		"Number of append timeouts.",
	),
	"scheduler_log_native_native_entries_skipped": newDesc(
		"scheduler_log", "native_native_entries_skipped",
		"Number of native entries skipped.",
	),
	"scheduler_log_native_read_events": newDesc(
		"scheduler_log", "native_read_events",
		"Number of native read events.",
	),
	"scheduler_log_native_read_failures": newDesc(
		"scheduler_log", "native_read_failures",
		"Number of read failures.",
	),
	"scheduler_log_native_read_nanos_total": newDesc(
		"scheduler_log", "native_read_nanos_total",
		"Timed native read total.",
	),
	"scheduler_log_native_read_timeouts": newDesc(
		"scheduler_log", "native_read_timeouts",
		"Number of read timeouts.",
	),
	"scheduler_log_native_truncate_events": newDesc(
		"scheduler_log", "native_truncate_events",
		"Number of truncate events.",
	),
	"scheduler_log_native_truncate_failures": newDesc(
		"scheduler_log", "native_truncate_failures",
		"Number of truncate failures.",
	),
	"scheduler_log_native_truncate_nanos_total": newDesc(
		"scheduler_log", "native_truncate_nanos_total",
		"Timed truncate total.",
	),
	"scheduler_log_native_truncate_timeouts": newDesc(
		"scheduler_log", "native_truncate_timeouts",
		"Number of truncate timeouts.",
	),
	"scheduler_log_snapshots": newDesc(
		"scheduler_log", "snapshots",
		"Number of snapshots.",
	),
	"scheduler_log_un_snapshotted_transactions": newDesc(
		"scheduler_log", "un_snapshotted_transactions",
		"Number of unsnapshotted transactions.",
	),
	"scheduler_resource_offers": newDesc(
		"scheduler", "resource_offers",
		"Number of resource offers that the scheduler has received.",
	),
	"scheduler_thrift_getJobSummary_events": newDesc(
		"scheduler_thrift", "getJobSummary_events",
		"Number of getJobSummary events.",
	),
	"scheduler_thrift_getJobSummary_nanos_total": newDesc(
		"scheduler_thrift", "getJobSummary_nanos_total",
		"Timed getJobSummary events.",
	),
	"scheduler_thrift_getQuota_events": newDesc(
		"scheduler_thrift", "getQuota_events",
		"Number of getQuota events.",
	),
	"scheduler_thrift_getQuota_nanos_per_event": newDesc(
		"scheduler_thrift", "getQuota_nanos_per_event",
		"Timed getQuota per event.",
	),
	"scheduler_thrift_getQuota_nanos_total": newDesc(
		"scheduler_thrift", "getQuota_nanos_total",
		"Timed getQuota total.",
	),
	"task_kill_retries": newDesc(
		"task", "kill_retries",
		"Number of times the scheduler has retried to kill a Task.",
	),
	"task_queries_all": newDesc(
		"task", "queries_all",
		"Number of queries.",
	),
	"task_queries_by_host": newDesc(
		"task", "queries_by_host",
		"Number of queries by host.",
	),
	"task_queries_by_id": newDesc(
		"task", "queries_by_id",
		"Number of queries by id.",
	),
	"task_queries_by_job": newDesc(
		"task", "queries_by_job",
		"Number of queries by job.",
	),
	"task_throttle_events": newDesc(
		"task", "throttle_events",
		"Number of throttle events.",
	),
	"task_throttle_ms_total": newDesc(
		"task", "throttle_ms_total",
		"Timed throttle total.",
	),
	"timed_out_tasks": newDesc(
		"", "timed_out_tasks",
		"Number of times the scheduler has given up waiting to hear back about a task in a transient state.",
	),
	"uncaught_exceptions": newDesc(
		"", "uncaught_exceptions",
		"Uncaught java exceptions.",
	),
}

var gauges = map[string]*prometheus.Desc{
	"http_200_responses_events_per_sec": newDesc(
		"http_200", "responses_events_per_sec",
		"Number of HTTP 200 status response events per second.",
	),
	"http_200_responses_nanos_per_event": newDesc(
		"http_200", "responses_nanos_per_event",
		"Timed HTTP 200 status responses per event.",
	),
	"http_200_responses_nanos_total_per_sec": newDesc(
		"http_200", "responses_nanos_total_per_sec",
		"Number of HTTP 200 status response events total per second.",
	),
	"jvm_gc_PS_MarkSweep_collection_time_ms": newDesc(
		"jvm", "gc_ps_marksweep_collection_time_ms",
		"Parallel mark and sweep collection time.",
	),
	"jvm_gc_PS_Scavenge_collection_time_ms": newDesc(
		"jvm", "gc_ps_scavenge_collection_time_ms",
		"Parallel scavenge collector time.",
	),
	"jvm_gc_collection_time_ms": newDesc(
		"jvm", "gc_collection_time_ms",
		"Garbage collection time.",
	),
	"jvm_memory_free_mb": newDesc(
		"jvm", "memory_free_mb",
		"Amount of free memory in the Java Virtual Machine.",
	),
	"jvm_memory_heap_mb_committed": newDesc(
		"jvm", "memory_heap_mb_committed",
		"Commited heap memory.",
	),
	"jvm_memory_heap_mb_used": newDesc(
		"jvm", "memory_heap_mb_used",
		"Current memory usage of the heap.",
	),
	"jvm_memory_non_heap_mb_committed": newDesc(
		"jvm", "memory_non_heap_mb_committed",
		"Commited non heap memory used.",
	),
	"jvm_memory_non_heap_mb_used": newDesc(
		"jvm", "memory_non_heap_mb_used",
		"Non heap memory used in MB.",
	),
	"jvm_threads_active": newDesc(
		"jvm", "threads_active",
		"Current number of live threads both daemon and non-daemon threads.",
	),
	"jvm_threads_daemon": newDesc(
		"jvm", "threads_daemon",
		"Current number of live daemon threads.",
	),
	"log_storage_write_lock_wait_events_per_sec": newDesc(
		"log_storage", "write_lock_wait_events_per_sec",
		"Number of lock wait events per second.",
	),
	"log_storage_write_lock_wait_ns_per_event": newDesc(
		"log_storage", "write_lock_wait_ns_per_event",
		"Number of write lock wait nanoseconds per event.",
	),
	"log_storage_write_lock_wait_ns_total_per_sec": newDesc(
		"log_storage", "write_lock_wait_ns_total_per_sec",
		"Number of lock wait nanoseconds total per second.",
	),
	"outstanding_offers": newDesc(
		"", "outstanding_offers",
		"Outstanding offers waiting to be returned.",
	),
	"process_cpu_cores_utilized": newDesc(
		"", "process_cpu_cores_utilized",
		"CPU time used by the process.",
	),
	"pubsub_executor_queue_size": newDesc(
		"", "pubsub_executor_queue_size",
		"Current pubsub executor queue size.",
	),
	"schedule_queue_size": newDesc(
		"schedule", "queue_size",
		"Task scheduler queue size.",
	),
	"scheduled_task_penalty_events_per_sec": newDesc(
		"scheduled_task", "penalty_events_per_sec",
		"Number of penalty events per second.",
	),
	"scheduled_task_penalty_ms_per_event": newDesc(
		"scheduled_task", "penalty_ms_per_event",
		"Number of scheduled task penalty in millisecond per event.",
	),
	"scheduled_task_penalty_ms_total_per_sec": newDesc(
		"scheduled_task", "penalty_ms_total_per_sec",
		"Number of scheduled task penalty in millisecond total per event.",
	),
	"scheduler_log_bytes_written": newDesc(
		"scheduler_log", "bytes_written",
		"Total bytes written.",
	),
	"scheduler_log_native_append_events_per_sec": newDesc(
		"scheduler_log", "native_append_events_per_sec",
		"Number of append events per second.",
	),
	"scheduler_log_native_append_nanos_per_event": newDesc(
		"scheduler_log", "native_append_nanos_per_event",
		"Number of appends per event.",
	),
	"scheduler_log_native_read_events_per_sec": newDesc(
		"scheduler_log", "native_read_events_per_sec",
		"Number of read events per second.",
	),
	"scheduler_log_native_read_nanos_per_event": newDesc(
		"scheduler_log", "native_read_nanos_per_event",
		"Number of reads per event.",
	),
	"scheduler_log_native_read_nanos_total_per_sec": newDesc(
		"scheduler_log", "native_read_nanos_total_per_sec",
		"Timed reads per second.",
	),
	"scheduler_log_native_truncate_events_per_sec": newDesc(
		"scheduler_log", "native_truncate_events_per_sec",
		"NUmber of truncate events per second.",
	),
	"scheduler_log_native_truncate_nanos_per_event": newDesc(
		"scheduler_log", "native_truncate_nanos_per_event",
		"Number of truncate per event.",
	),
	"scheduler_log_native_truncate_nanos_total_per_sec": newDesc(
		"scheduler_log", "native_truncate_nanos_total_per_sec",
		"Timed total truncate per second.",
	),
	"scheduler_thrift_getJobSummary_events_per_sec": newDesc(
		"scheduler_thrift", "getJobSummary_events_per_sec",
		"Timed getJobSummary per second.",
	),
	"scheduler_thrift_getJobSummary_nanos_per_event": newDesc(
		"scheduler_thrift", "getJobSummary_nanos_per_event",
		"Timed getJobSummary per event.",
	),
	"scheduler_thrift_getJobSummary_nanos_total_per_sec": newDesc(
		"scheduler_thrift", "getJobSummary_nanos_total_per_sec",
		"Timed getJobSummary total per second.",
	),
	"scheduler_thrift_getQuota_events_per_sec": newDesc(
		"scheduler_thrift", "getQuota_events_per_sec",
		"Timed getQuota events per second.",
	),
	"scheduler_thrift_getQuota_nanos_total_per_sec": newDesc(
		"scheduler_thrift", "getQuota_nanos_total_per_sec",
		"Timed getQuota per second.",
	),
	"sla_cluster_mtta_ms": newDesc(
		"sla", "cluster_mtta_ms",
		"Median time to assigned.",
	),
	"sla_cluster_mttr_ms": newDesc(
		"sla", "cluster_mttr_ms",
		"Median time to running.",
	),
	"sla_cluster_platform_uptime_percent": newDesc(
		"sla", "cluster_platform_uptime_percent",
		"Aggregate amount of time a job spends in a non-runnable state.",
	),
	"sla_cluster_mtta_ms_nonprod": newDesc(
		"sla", "cluster_mtta_ms_nonprod",
		"Median time to assigned nonprod.",
	),
	"sla_cluster_mttr_ms_nonprod": newDesc(
		"sla", "cluster_mttr_ms_nonprod",
		"Median time to running nonprod.",
	),
	"sla_cluster_platform_uptime_percent_nonprod": newDesc(
		"sla", "cluster_platform_uptime_percent_nonprod",
		"Aggregate amount of time a job spends in a non-runnable state nonprod.",
	),
	"system_free_physical_memory_mb": newDesc(
		"system", "free_physical_memory_mb",
		"Free physical memory in MB.",
	),
	"system_load_avg": newDesc(
		"system", "load_avg",
		"1 minute load average.",
	),
	"task_throttle_events_per_sec": newDesc(
		"task", "throttle_events_per_sec",
		"Timed throttle events per second.",
	),
	"task_throttle_ms_total_per_sec": newDesc(
		"task", "throttle_ms_total_per_sec",
		"Timed throttle per second.",
	),
	"timeout_queue_size": newDesc(
		"", "timeout_queue_size",
		"Total timeout queue size.",
	),
}

// Quota values containing _mb in their name are converted to bytes
var quotas = map[string]*prometheus.Desc{
	"cpu_cores": newDescWithLabels(
		"quota", "cpus", "Maximum number of cpu cores.",
		[]string{"role", },
	),
	"ram_mb": newDescWithLabels(
		"quota", "ram_bytes", "Maximum amount of RAM available in bytes.",
		[]string{"role", },
	),
	"disk_mb": newDescWithLabels(
		"quota", "disk_bytes", "Maximum amount of disk space available in bytes.",
		[]string{"role", },
	),
}

type parser struct {
	match int
	vt    prometheus.ValueType
	desc  *prometheus.Desc
	regex *regexp.Regexp
}

func (p *parser) parse(name string, value float64, ch chan<- prometheus.Metric) {
	match := p.regex.FindStringSubmatch(name)
	if len(match) == p.match {
		ch <- prometheus.MustNewConstMetric(
			p.desc,
			p.vt,
			value,
			match[1:]...,
		)
	}
}

var prefixParser = map[string]*parser{
	"tasks_": &parser{
		match: 5,
		vt:    prometheus.CounterValue,
		desc: newDescWithLabels(
			"", "tasks",
			"Task state per job.",
			[]string{"state", "role", "env", "job"},
		),
		regex: regexp.MustCompile("tasks_(?P<state>.*)_(?P<role>.*)/(?P<env>.*)/(?P<job>.*)"),
	},
	"tasks_lost_rack_": &parser{
		match: 2,
		vt:    prometheus.CounterValue,
		desc: newDescWithLabels(
			"", "tasks_lost_rack",
			"Task lost per rack total.",
			[]string{"rack"},
		),
		regex: regexp.MustCompile("tasks_lost_rack_(?P<rack>.*)"),
	},
	"task_store_": &parser{
		match: 2,
		vt:    prometheus.GaugeValue,
		desc: newDescWithLabels(
			"", "task_store",
			"Task store state.",
			[]string{"state"},
		),
		regex: regexp.MustCompile("task_store_(?P<state>[A-Z]+)"),
	},
	"thrift_workload_": &parser{
		match: 2,
		vt:    prometheus.CounterValue,
		desc: newDescWithLabels(
			"", "thrift_workload",
			"Number of thrift calls per method.",
			[]string{"method"},
		),
		regex: regexp.MustCompile("thrift_workload_(?P<method>.*)"),
	},
	"update_transition_": &parser{
		match: 2,
		vt:    prometheus.CounterValue,
		desc: newDescWithLabels(
			"", "update_transition",
			"Update transition.",
			[]string{"state"},
		),
		regex: regexp.MustCompile("update_transition_(?P<state>.*)"),
	},
	"scheduler_lifecycle_": &parser{
		match: 2,
		vt:    prometheus.GaugeValue,
		desc: newDescWithLabels(
			"", "scheduler_lifecycle",
			"Scheduler lifecycle.",
			[]string{"state"},
		),
		regex: regexp.MustCompile("scheduler_lifecycle_(?P<state>[A-Z]+)"),
	},
	"zk_connection_state_": &parser{
		match: 2,
		vt:    prometheus.GaugeValue,
		desc: newDescWithLabels(
			"", "zk_connection_state",
			"Current ZooKeeper connection state.",
			[]string{"state"},
		),
		regex: regexp.MustCompile("zk_connection_state_(?P<state>[A-Z]+)$"),
	},
}

var suffixParser = map[string]*parser{
	"_counter": &parser{
		match: 2,
		vt:    prometheus.GaugeValue,
		desc: newDescWithLabels(
			"", "zk_connection_state_total",
			"Number of ZooKeeper connection state changes.",
			[]string{"state"},
		),
		regex: regexp.MustCompile("zk_connection_state_(?P<state>[A-Z]+)_counter$"),
	},
	"_mtta_ms": &parser{
		match: 4,
		vt:    prometheus.GaugeValue,
		desc: newDescWithLabels(
			"", "sla_mtta_ms",
			"Median time to assigned.",
			[]string{"role", "env", "job"},
		),
		regex: regexp.MustCompile("sla_(?P<role>.*)/(?P<env>.*)/(?P<job>.*)_mtta_ms$"),
	},
	"_mttr_ms": &parser{
		match: 4,
		vt:    prometheus.GaugeValue,
		desc: newDescWithLabels(
			"", "sla_mttr_ms",
			"Median time to running.",
			[]string{"role", "env", "job"},
		),
		regex: regexp.MustCompile("sla_(?P<role>.*)/(?P<env>.*)/(?P<job>.*)_mttr_ms$"),
	},
	"_mtta_ms_nonprod": &parser{
		match: 4,
		vt:    prometheus.GaugeValue,
		desc: newDescWithLabels(
			"", "sla_mtta_ms_nonprod",
			"Median time to assigned nonprod.",
			[]string{"role", "env", "job"},
		),
		regex: regexp.MustCompile("sla_(?P<role>.*)/(?P<env>.*)/(?P<job>.*)_mtta_ms_nonprod$"),
	},
	"_mttr_ms_nonprod": &parser{
		match: 4,
		vt:    prometheus.GaugeValue,
		desc: newDescWithLabels(
			"", "sla_mttr_ms_nonprod",
			"Median time to running nonprod.",
			[]string{"role", "env", "job"},
		),
		regex: regexp.MustCompile("sla_(?P<role>.*)/(?P<env>.*)/(?P<job>.*)_mttr_ms_nonprod$"),
	},
	"_platform_uptime_percent": &parser{
		match: 4,
		vt:    prometheus.GaugeValue,
		desc: newDescWithLabels(
			"", "sla_platform_uptime_percent",
			"Aggregate platform uptime.",
			[]string{"role", "env", "job"},
		),
		regex: regexp.MustCompile("sla_(?P<role>.*)/(?P<env>.*)/(?P<job>.*)_platform_uptime_percent$"),
	},
	"_platform_uptime_percent_nonprod": &parser{
		match: 4,
		vt:    prometheus.GaugeValue,
		desc: newDescWithLabels(
			"", "sla_platform_uptime_percent_nonprod",
			"Aggregate platform uptime nonprod.",
			[]string{"role", "env", "job"},
		),
		regex: regexp.MustCompile("sla_(?P<role>.*)/(?P<env>.*)/(?P<job>.*)_platform_uptime_percent_nonprod$"),
	},
}

func labelVars(ch chan<- prometheus.Metric, name string, value float64) {
	for prefix, parser := range prefixParser {
		if strings.HasPrefix(name, prefix) {
			parser.parse(name, value, ch)
		}
	}

	for suffix, parser := range suffixParser {
		if strings.HasSuffix(name, suffix) {
			parser.parse(name, value, ch)
		}
	}
}

func newDesc(subsys, name, descr string) *prometheus.Desc {
	fqn := prometheus.BuildFQName(namespace, subsys, name)
	return prometheus.NewDesc(fqn, descr, nil, nil)
}

func newDescWithLabels(subsys, name, descr string, labels []string) *prometheus.Desc {
	fqn := prometheus.BuildFQName(namespace, subsys, name)
	return prometheus.NewDesc(fqn, descr, labels, nil)
}
