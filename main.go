package main

import "github.com/posener/complete"

func main() {

	add_blob := complete.Command{
		Flags: complete.Flags{
			"--dir=": complete.PredictAnything,
		},
	}

	alias_env := complete.Command{
		Flags: complete.Flags{},
	}

	attach_disk := complete.Command{
		Flags: complete.Flags{
			"--disk-properties=": complete.PredictAnything,
		},
	}

	blobs := complete.Command{
		Flags: complete.Flags{
			"--dir=": complete.PredictAnything,
		},
	}

	cancel_task := complete.Command{
		Flags: complete.Flags{},
	}

	clean_up := complete.Command{
		Flags: complete.Flags{
			"--all": complete.PredictAnything,
		},
	}

	cloud_check := complete.Command{
		Flags: complete.Flags{
			"-a":            complete.PredictAnything,
			"--auto":        complete.PredictAnything,
			"--resolution=": complete.PredictAnything,
			"-r":            complete.PredictAnything,
			"--report":      complete.PredictAnything,
		},
	}

	cloud_config := complete.Command{
		Flags: complete.Flags{},
	}

	config := complete.Command{
		Flags: complete.Flags{
			"--name=": complete.PredictAnything,
			"--type=": complete.PredictAnything,
		},
	}

	configs := complete.Command{
		Flags: complete.Flags{
			"--name=":   complete.PredictAnything,
			"--type=":   complete.PredictAnything,
			"-r":        complete.PredictAnything,
			"--recent=": complete.PredictAnything,
		},
	}

	cpi_config := complete.Command{
		Flags: complete.Flags{},
	}

	create_env := complete.Command{
		Flags: complete.Flags{
			"-v":                          complete.PredictAnything,
			"--var=VAR=VALUE":             complete.PredictAnything,
			"--var-file=VAR=PATH":         complete.PredictAnything,
			"-l":                          complete.PredictAnything,
			"--vars-file=PATH":            complete.PredictAnything,
			"--vars-env=PREFIX":           complete.PredictAnything,
			"--vars-store=PATH":           complete.PredictAnything,
			"-o":                          complete.PredictAnything,
			"--ops-file=PATH":             complete.PredictAnything,
			"--skip-drain":                complete.PredictAnything,
			"--state=PATH":                complete.PredictAnything,
			"--recreate":                  complete.PredictAnything,
			"--recreate-persistent-disks": complete.PredictAnything,
		},
	}

	create_release := complete.Command{
		Flags: complete.Flags{
			"--dir=":              complete.PredictAnything,
			"--name=":             complete.PredictAnything,
			"--version=":          complete.PredictAnything,
			"--timestamp-version": complete.PredictAnything,
			"--final":             complete.PredictAnything,
			"--tarball=":          complete.PredictAnything,
			"--force":             complete.PredictAnything,
		},
	}

	delete_config := complete.Command{
		Flags: complete.Flags{
			"--type=": complete.PredictAnything,
			"--name=": complete.PredictAnything,
		},
	}

	delete_deployment := complete.Command{
		Flags: complete.Flags{
			"--force": complete.PredictAnything,
		},
	}

	delete_disk := complete.Command{
		Flags: complete.Flags{},
	}

	delete_env := complete.Command{
		Flags: complete.Flags{
			"-v":                  complete.PredictAnything,
			"--var=VAR=VALUE":     complete.PredictAnything,
			"--var-file=VAR=PATH": complete.PredictAnything,
			"-l":                  complete.PredictAnything,
			"--vars-file=PATH":    complete.PredictAnything,
			"--vars-env=PREFIX":   complete.PredictAnything,
			"--vars-store=PATH":   complete.PredictAnything,
			"-o":                  complete.PredictAnything,
			"--ops-file=PATH":     complete.PredictAnything,
			"--skip-drain":        complete.PredictAnything,
			"--state=PATH":        complete.PredictAnything,
		},
	}

	delete_network := complete.Command{
		Flags: complete.Flags{},
	}

	delete_release := complete.Command{
		Flags: complete.Flags{
			"--force": complete.PredictAnything,
		},
	}

	delete_snapshot := complete.Command{
		Flags: complete.Flags{},
	}

	delete_snapshots := complete.Command{
		Flags: complete.Flags{},
	}

	delete_stemcell := complete.Command{
		Flags: complete.Flags{
			"--force": complete.PredictAnything,
		},
	}

	delete_vm := complete.Command{
		Flags: complete.Flags{},
	}

	deploy := complete.Command{
		Flags: complete.Flags{
			"-v":                          complete.PredictAnything,
			"--var=VAR=VALUE":             complete.PredictAnything,
			"--var-file=VAR=PATH":         complete.PredictAnything,
			"-l":                          complete.PredictAnything,
			"--vars-file=PATH":            complete.PredictAnything,
			"--vars-env=PREFIX":           complete.PredictAnything,
			"--vars-store=PATH":           complete.PredictAnything,
			"-o":                          complete.PredictAnything,
			"--ops-file=PATH":             complete.PredictAnything,
			"--no-redact":                 complete.PredictAnything,
			"--recreate":                  complete.PredictAnything,
			"--recreate-persistent-disks": complete.PredictAnything,
			"--fix":                       complete.PredictAnything,
			"--skip-drain=INSTANCE-GROUP": complete.PredictAnything,
			"--canaries=":                 complete.PredictAnything,
			"--max-in-flight=":            complete.PredictAnything,
			"--dry-run":                   complete.PredictAnything,
		},
	}

	deployment := complete.Command{
		Flags: complete.Flags{},
	}

	deployments := complete.Command{
		Flags: complete.Flags{},
	}

	diff_config := complete.Command{
		Flags: complete.Flags{
			"--from-id=":      complete.PredictAnything,
			"--to-id=":        complete.PredictAnything,
			"--from-content=": complete.PredictAnything,
			"--to-content=":   complete.PredictAnything,
		},
	}

	disks := complete.Command{
		Flags: complete.Flags{
			"-o":         complete.PredictAnything,
			"--orphaned": complete.PredictAnything,
		},
	}

	environment := complete.Command{
		Flags: complete.Flags{},
	}

	environments := complete.Command{
		Flags: complete.Flags{},
	}

	errands := complete.Command{
		Flags: complete.Flags{},
	}

	event := complete.Command{
		Flags: complete.Flags{},
	}

	events := complete.Command{
		Flags: complete.Flags{
			"--before-id=":   complete.PredictAnything,
			"--before=":      complete.PredictAnything,
			"--after=":       complete.PredictAnything,
			"--task=":        complete.PredictAnything,
			"--instance=":    complete.PredictAnything,
			"--event-user=":  complete.PredictAnything,
			"--action=":      complete.PredictAnything,
			"--object-type=": complete.PredictAnything,
			"--object-name=": complete.PredictAnything,
		},
	}

	export_release := complete.Command{
		Flags: complete.Flags{
			"--dir=": complete.PredictAnything,
			"--job=": complete.PredictAnything,
		},
	}

	finalize_release := complete.Command{
		Flags: complete.Flags{
			"--dir=":     complete.PredictAnything,
			"--name=":    complete.PredictAnything,
			"--version=": complete.PredictAnything,
			"--force":    complete.PredictAnything,
		},
	}

	generate_job := complete.Command{
		Flags: complete.Flags{
			"--dir=": complete.PredictAnything,
		},
	}

	generate_package := complete.Command{
		Flags: complete.Flags{
			"--dir=": complete.PredictAnything,
		},
	}

	help := complete.Command{
		Flags: complete.Flags{},
	}

	ignore := complete.Command{
		Flags: complete.Flags{},
	}

	init_release := complete.Command{
		Flags: complete.Flags{
			"--dir=": complete.PredictAnything,
			"--git":  complete.PredictAnything,
		},
	}

	inspect_local_stemcell := complete.Command{
		Flags: complete.Flags{},
	}

	inspect_release := complete.Command{
		Flags: complete.Flags{},
	}

	instances := complete.Command{
		Flags: complete.Flags{
			"-i":        complete.PredictAnything,
			"--details": complete.PredictAnything,
			"--dns":     complete.PredictAnything,
			"--vitals":  complete.PredictAnything,
			"-p":        complete.PredictAnything,
			"--ps":      complete.PredictAnything,
			"-f":        complete.PredictAnything,
			"--failing": complete.PredictAnything,
		},
	}

	interpolate := complete.Command{
		Flags: complete.Flags{
			"-v":                  complete.PredictAnything,
			"--var=VAR=VALUE":     complete.PredictAnything,
			"--var-file=VAR=PATH": complete.PredictAnything,
			"-l":                  complete.PredictAnything,
			"--vars-file=PATH":    complete.PredictAnything,
			"--vars-env=PREFIX":   complete.PredictAnything,
			"--vars-store=PATH":   complete.PredictAnything,
			"-o":                  complete.PredictAnything,
			"--ops-file=PATH":     complete.PredictAnything,
			"--path=OP-PATH":      complete.PredictAnything,
			"--var-errs":          complete.PredictAnything,
			"--var-errs-unused":   complete.PredictAnything,
		},
	}

	locks := complete.Command{
		Flags: complete.Flags{},
	}

	log_in := complete.Command{
		Flags: complete.Flags{},
	}

	log_out := complete.Command{
		Flags: complete.Flags{},
	}

	logs := complete.Command{
		Flags: complete.Flags{
			"--dir=":            complete.PredictAnything,
			"-f":                complete.PredictAnything,
			"--follow":          complete.PredictAnything,
			"--num=":            complete.PredictAnything,
			"-q":                complete.PredictAnything,
			"--quiet":           complete.PredictAnything,
			"--job=":            complete.PredictAnything,
			"--only=":           complete.PredictAnything,
			"--agent":           complete.PredictAnything,
			"--gw-disable":      complete.PredictAnything,
			"--gw-user=":        complete.PredictAnything,
			"--gw-host=":        complete.PredictAnything,
			"--gw-private-key=": complete.PredictAnything,
			"--gw-socks5=":      complete.PredictAnything,
		},
	}

	manifest := complete.Command{
		Flags: complete.Flags{},
	}

	networks := complete.Command{
		Flags: complete.Flags{
			"-o":         complete.PredictAnything,
			"--orphaned": complete.PredictAnything,
		},
	}

	orphan_disk := complete.Command{
		Flags: complete.Flags{},
	}

	orphaned_vms := complete.Command{
		Flags: complete.Flags{},
	}

	recreate := complete.Command{
		Flags: complete.Flags{
			"--skip-drain":     complete.PredictAnything,
			"--force":          complete.PredictAnything,
			"--fix":            complete.PredictAnything,
			"--canaries=":      complete.PredictAnything,
			"--max-in-flight=": complete.PredictAnything,
			"--dry-run":        complete.PredictAnything,
		},
	}

	releases := complete.Command{
		Flags: complete.Flags{},
	}

	remove_blob := complete.Command{
		Flags: complete.Flags{
			"--dir=": complete.PredictAnything,
		},
	}

	repack_stemcell := complete.Command{
		Flags: complete.Flags{
			"--name=":             complete.PredictAnything,
			"--cloud-properties=": complete.PredictAnything,
			"--empty-image":       complete.PredictAnything,
			"--format=":           complete.PredictAnything,
			"--version=":          complete.PredictAnything,
		},
	}

	reset_release := complete.Command{
		Flags: complete.Flags{
			"--dir=": complete.PredictAnything,
		},
	}

	restart := complete.Command{
		Flags: complete.Flags{
			"--skip-drain":     complete.PredictAnything,
			"--force":          complete.PredictAnything,
			"--canaries=":      complete.PredictAnything,
			"--max-in-flight=": complete.PredictAnything,
		},
	}

	run_errand := complete.Command{
		Flags: complete.Flags{
			"--instance=INSTANCE-GROUP[/INSTANCE-ID]": complete.PredictAnything,
			"--keep-alive":    complete.PredictAnything,
			"--when-changed":  complete.PredictAnything,
			"--download-logs": complete.PredictAnything,
			"--logs-dir=":     complete.PredictAnything,
		},
	}

	runtime_config := complete.Command{
		Flags: complete.Flags{
			"--name=": complete.PredictAnything,
		},
	}

	scp := complete.Command{
		Flags: complete.Flags{
			"-r":                complete.PredictAnything,
			"--recursive":       complete.PredictAnything,
			"--gw-disable":      complete.PredictAnything,
			"--gw-user=":        complete.PredictAnything,
			"--gw-host=":        complete.PredictAnything,
			"--gw-private-key=": complete.PredictAnything,
			"--gw-socks5=":      complete.PredictAnything,
		},
	}

	snapshots := complete.Command{
		Flags: complete.Flags{},
	}

	ssh := complete.Command{
		Flags: complete.Flags{
			"-c":                complete.PredictAnything,
			"--command=":        complete.PredictAnything,
			"--opts=":           complete.PredictAnything,
			"-r":                complete.PredictAnything,
			"--results":         complete.PredictAnything,
			"--gw-disable":      complete.PredictAnything,
			"--gw-user=":        complete.PredictAnything,
			"--gw-host=":        complete.PredictAnything,
			"--gw-private-key=": complete.PredictAnything,
			"--gw-socks5=":      complete.PredictAnything,
		},
	}

	start := complete.Command{
		Flags: complete.Flags{
			"--force":          complete.PredictAnything,
			"--canaries=":      complete.PredictAnything,
			"--max-in-flight=": complete.PredictAnything,
		},
	}

	stemcells := complete.Command{
		Flags: complete.Flags{},
	}

	stop := complete.Command{
		Flags: complete.Flags{
			"--soft":           complete.PredictAnything,
			"--hard":           complete.PredictAnything,
			"--skip-drain":     complete.PredictAnything,
			"--force":          complete.PredictAnything,
			"--canaries=":      complete.PredictAnything,
			"--max-in-flight=": complete.PredictAnything,
		},
	}

	sync_blobs := complete.Command{
		Flags: complete.Flags{
			"--dir=": complete.PredictAnything,
		},
	}

	take_snapshot := complete.Command{
		Flags: complete.Flags{},
	}

	task := complete.Command{
		Flags: complete.Flags{
			"--event":  complete.PredictAnything,
			"--cpi":    complete.PredictAnything,
			"--debug":  complete.PredictAnything,
			"--result": complete.PredictAnything,
			"-a":       complete.PredictAnything,
			"--all":    complete.PredictAnything,
		},
	}

	tasks := complete.Command{
		Flags: complete.Flags{
			"-r":        complete.PredictAnything,
			"--recent=": complete.PredictAnything,
			"-a":        complete.PredictAnything,
			"--all":     complete.PredictAnything,
		},
	}

	unignore := complete.Command{
		Flags: complete.Flags{},
	}

	update_cloud_config := complete.Command{
		Flags: complete.Flags{
			"-v":                  complete.PredictAnything,
			"--var=VAR=VALUE":     complete.PredictAnything,
			"--var-file=VAR=PATH": complete.PredictAnything,
			"-l":                  complete.PredictAnything,
			"--vars-file=PATH":    complete.PredictAnything,
			"--vars-env=PREFIX":   complete.PredictAnything,
			"--vars-store=PATH":   complete.PredictAnything,
			"-o":                  complete.PredictAnything,
			"--ops-file=PATH":     complete.PredictAnything,
		},
	}

	update_config := complete.Command{
		Flags: complete.Flags{
			"--type=":               complete.PredictAnything,
			"--name=":               complete.PredictAnything,
			"--expected-latest-id=": complete.PredictAnything,
			"-v":                    complete.PredictAnything,
			"--var=VAR=VALUE":       complete.PredictAnything,
			"--var-file=VAR=PATH":   complete.PredictAnything,
			"-l":                    complete.PredictAnything,
			"--vars-file=PATH":      complete.PredictAnything,
			"--vars-env=PREFIX":     complete.PredictAnything,
			"--vars-store=PATH":     complete.PredictAnything,
			"-o":                    complete.PredictAnything,
			"--ops-file=PATH":       complete.PredictAnything,
		},
	}

	update_cpi_config := complete.Command{
		Flags: complete.Flags{
			"-v":                  complete.PredictAnything,
			"--var=VAR=VALUE":     complete.PredictAnything,
			"--var-file=VAR=PATH": complete.PredictAnything,
			"-l":                  complete.PredictAnything,
			"--vars-file=PATH":    complete.PredictAnything,
			"--vars-env=PREFIX":   complete.PredictAnything,
			"--vars-store=PATH":   complete.PredictAnything,
			"-o":                  complete.PredictAnything,
			"--ops-file=PATH":     complete.PredictAnything,
			"--no-redact":         complete.PredictAnything,
		},
	}

	update_resurrection := complete.Command{
		Flags: complete.Flags{},
	}

	update_runtime_config := complete.Command{
		Flags: complete.Flags{
			"-v":                  complete.PredictAnything,
			"--var=VAR=VALUE":     complete.PredictAnything,
			"--var-file=VAR=PATH": complete.PredictAnything,
			"-l":                  complete.PredictAnything,
			"--vars-file=PATH":    complete.PredictAnything,
			"--vars-env=PREFIX":   complete.PredictAnything,
			"--vars-store=PATH":   complete.PredictAnything,
			"-o":                  complete.PredictAnything,
			"--ops-file=PATH":     complete.PredictAnything,
			"--no-redact":         complete.PredictAnything,
			"--name=":             complete.PredictAnything,
		},
	}

	upload_blobs := complete.Command{
		Flags: complete.Flags{
			"--dir=": complete.PredictAnything,
		},
	}

	upload_release := complete.Command{
		Flags: complete.Flags{
			"--dir=":                complete.PredictAnything,
			"--rebase":              complete.PredictAnything,
			"--fix":                 complete.PredictAnything,
			"--name=":               complete.PredictAnything,
			"--version=":            complete.PredictAnything,
			"--sha1=":               complete.PredictAnything,
			"--stemcell=OS/VERSION": complete.PredictAnything,
		},
	}

	upload_stemcell := complete.Command{
		Flags: complete.Flags{
			"--fix":      complete.PredictAnything,
			"--name=":    complete.PredictAnything,
			"--version=": complete.PredictAnything,
			"--sha1=":    complete.PredictAnything,
		},
	}

	variables := complete.Command{
		Flags: complete.Flags{},
	}

	vendor_package := complete.Command{
		Flags: complete.Flags{
			"--dir=": complete.PredictAnything,
		},
	}

	vms := complete.Command{
		Flags: complete.Flags{
			"--dns":              complete.PredictAnything,
			"--vitals":           complete.PredictAnything,
			"--cloud-properties": complete.PredictAnything,
		},
	}

	bosh := complete.Command{
		Sub: complete.Commands{
			"add-blob":               add_blob,
			"alias-env":              alias_env,
			"attach-disk":            attach_disk,
			"blobs":                  blobs,
			"cancel-task":            cancel_task,
			"clean-up":               clean_up,
			"cloud-check":            cloud_check,
			"cloud-config":           cloud_config,
			"config":                 config,
			"configs":                configs,
			"cpi-config":             cpi_config,
			"create-env":             create_env,
			"create-release":         create_release,
			"delete-config":          delete_config,
			"delete-deployment":      delete_deployment,
			"delete-disk":            delete_disk,
			"delete-env":             delete_env,
			"delete-network":         delete_network,
			"delete-release":         delete_release,
			"delete-snapshot":        delete_snapshot,
			"delete-snapshots":       delete_snapshots,
			"delete-stemcell":        delete_stemcell,
			"delete-vm":              delete_vm,
			"deploy":                 deploy,
			"deployment":             deployment,
			"deployments":            deployments,
			"diff-config":            diff_config,
			"disks":                  disks,
			"environment":            environment,
			"environments":           environments,
			"errands":                errands,
			"event":                  event,
			"events":                 events,
			"export-release":         export_release,
			"finalize-release":       finalize_release,
			"generate-job":           generate_job,
			"generate-package":       generate_package,
			"help":                   help,
			"ignore":                 ignore,
			"init-release":           init_release,
			"inspect-local-stemcell": inspect_local_stemcell,
			"inspect-release":        inspect_release,
			"instances":              instances,
			"interpolate":            interpolate,
			"locks":                  locks,
			"log-in":                 log_in,
			"log-out":                log_out,
			"logs":                   logs,
			"manifest":               manifest,
			"networks":               networks,
			"orphan-disk":            orphan_disk,
			"orphaned-vms":           orphaned_vms,
			"recreate":               recreate,
			"releases":               releases,
			"remove-blob":            remove_blob,
			"repack-stemcell":        repack_stemcell,
			"reset-release":          reset_release,
			"restart":                restart,
			"run-errand":             run_errand,
			"runtime-config":         runtime_config,
			"scp":                    scp,
			"snapshots":              snapshots,
			"ssh":                    ssh,
			"start":                  start,
			"stemcells":              stemcells,
			"stop":                   stop,
			"sync-blobs":             sync_blobs,
			"take-snapshot":          take_snapshot,
			"task":                   task,
			"tasks":                  tasks,
			"unignore":               unignore,
			"update-cloud-config":    update_cloud_config,
			"update-config":          update_config,
			"update-cpi-config":      update_cpi_config,
			"update-resurrection":    update_resurrection,
			"update-runtime-config":  update_runtime_config,
			"upload-blobs":           upload_blobs,
			"upload-release":         upload_release,
			"upload-stemcell":        upload_stemcell,
			"variables":              variables,
			"vendor-package":         vendor_package,
			"vms":                    vms,
		},

		GlobalFlags: complete.Flags{
			"-v":                complete.PredictAnything,
			"--version":         complete.PredictAnything,
			"--config=":         complete.PredictAnything,
			"-e":                complete.PredictAnything,
			"--environment=":    complete.PredictAnything,
			"--ca-cert=":        complete.PredictAnything,
			"--sha2":            complete.PredictAnything,
			"--parallel=":       complete.PredictAnything,
			"--client=":         complete.PredictAnything,
			"--client-secret=":  complete.PredictAnything,
			"-d":                complete.PredictAnything,
			"--deployment=":     complete.PredictAnything,
			"--column=":         complete.PredictAnything,
			"--json":            complete.PredictAnything,
			"--tty":             complete.PredictAnything,
			"--no-color":        complete.PredictAnything,
			"-n":                complete.PredictAnything,
			"--non-interactive": complete.PredictAnything,
			"-h":                complete.PredictAnything,
			"--help":            complete.PredictAnything,
		},
	}

	complete.New("bosh", bosh).Run()
}
