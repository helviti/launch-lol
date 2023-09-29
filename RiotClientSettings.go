package main

type RiotClientSettings struct {
	Install struct {
		ArcaneOutroHidden bool `yaml:"arcane_outro_hidden"`
		AutoUpdate        struct {
			LeagueOfLegendsLive bool `yaml:"league_of_legends.live"`
		} `yaml:"auto-update"`
		AutoUpdateProductDismissCount struct {
			LeagueOfLegends int `yaml:"league_of_legends"`
		} `yaml:"auto-update-product-dismiss-count"`
		BackgroundRunningNotificationDisplayed bool `yaml:"background_running_notification_displayed"`
		Cohorts                                struct {
			RC15NewLifecycle string `yaml:"RC_15.new_lifecycle"`
		} `yaml:"cohorts"`
		Globals struct {
			Locale string `yaml:"locale"`
			Region string `yaml:"region"`
		} `yaml:"globals"`
		LastSelectedPatchlines struct {
			LeagueOfLegends string `yaml:"league_of_legends"`
		} `yaml:"last-selected-patchlines"`
		LastSessionTimestamp struct {
			LeagueOfLegendsLive int `yaml:"league_of_legends.live"`
		} `yaml:"last-session-timestamp"`
		Lifecycle struct {
			EnableLaunchOnComputerStartSetByPlayer bool        `yaml:"enable_launch_on_computer_start_set_by_player"`
			EnableRunInBackground                  bool        `yaml:"enable_run_in_background"`
			EnableRunInBackgroundSetByPlayer       interface{} `yaml:"enable_run_in_background_set_by_player"`
		} `yaml:"lifecycle"`
		MfaNotificationDismissed bool `yaml:"mfa_notification_dismissed"`
		MultigameClient          struct {
			ShortcutCreated bool `yaml:"shortcut_created"`
		} `yaml:"multigame-client"`
		PatchTags struct {
			LeagueOfLegendsLive []string `yaml:"league_of_legends.live"`
		} `yaml:"patch-tags"`
		PlayerAffinity struct {
			Product struct {
				Bacon struct {
					Live string `yaml:"live"`
				} `yaml:"bacon"`
				Valorant struct {
					Live string `yaml:"live"`
					Pbe  string `yaml:"pbe"`
				} `yaml:"valorant"`
			} `yaml:"product"`
			Service struct {
				Chat            string `yaml:"chat"`
				Mailbox         string `yaml:"mailbox"`
				ReportCollector string `yaml:"report-collector"`
				Rms             string `yaml:"rms"`
			} `yaml:"service"`
		} `yaml:"player-affinity"`
		ProductReleaseIds struct {
			LeagueOfLegendsLive string `yaml:"league_of_legends.live"`
		} `yaml:"product-release-ids"`
		RiotClientAppCommand struct {
			UpgradeAttempts int  `yaml:"upgrade_attempts"`
			Upgraded        bool `yaml:"upgraded"`
		} `yaml:"riot-client-app-command"`
		RiotClientAuth struct {
			HomeBaseCountry string `yaml:"home-base-country"`
		} `yaml:"riot-client-auth"`
		Riotgamesapi struct {
			Telemetry struct {
				AnalyticsPlatformAllowedEvents struct {
					ChatCommandTelemetry                       string `yaml:"ChatCommandTelemetry"`
					ChatDisconnectEvent                        bool   `yaml:"ChatDisconnectEvent"`
					ChatLoginState                             bool   `yaml:"ChatLoginState"`
					ChatStanzaResponse                         bool   `yaml:"ChatStanzaResponse"`
					ConsoleAccountSLI                          string `yaml:"ConsoleAccountSLI"`
					EndpointResult                             string `yaml:"EndpointResult"`
					ErrorLog                                   string `yaml:"ErrorLog"`
					ExternalWebRequest                         bool   `yaml:"ExternalWebRequest"`
					FirstPartyFulfillmentSLI                   string `yaml:"FirstPartyFulfillmentSLI"`
					LeagueRegionElection                       string `yaml:"LeagueRegionElection"`
					PayMobile                                  string `yaml:"PayMobile"`
					PlatformUIOperationalEvent                 string `yaml:"PlatformUIOperationalEvent"`
					PlatformUIRenderStats                      string `yaml:"PlatformUIRenderStats"`
					PlatformUIStatusEvent                      string `yaml:"PlatformUIStatusEvent"`
					PlayerAccountAliasSanitizedEvent           string `yaml:"PlayerAccountAliasSanitizedEvent"`
					PlayerSessionLifecycleSLI                  string `yaml:"PlayerSessionLifecycleSLI"`
					PlayerSessionLifecycleStatus               string `yaml:"PlayerSessionLifecycleStatus"`
					ProductIntegrationAppRepairEvent           string `yaml:"ProductIntegrationAppRepairEvent"`
					ProductIntegrationEvent                    string `yaml:"ProductIntegrationEvent"`
					RSOAuthLoginSteps                          string `yaml:"RSOAuthLoginSteps"`
					RSOMobileAccountDeletionUI                 string `yaml:"RSOMobileAccountDeletionUI"`
					RSOMobileUIEvent                           string `yaml:"RSOMobileUIEvent"`
					RSOSignOnAttempt                           string `yaml:"RSOSignOnAttempt"`
					RpgPciPenaltyNotificationPlayerActionEvent string `yaml:"RpgPciPenaltyNotificationPlayerActionEvent"`
					RpgPciReporterFeedbackAcknowledgedEvent    string `yaml:"RpgPciReporterFeedbackAcknowledgedEvent"`
					RpgPciReporterFeedbackReceivedEvent        string `yaml:"RpgPciReporterFeedbackReceivedEvent"`
					RpgPciWarningAcknowledgementEvent          string `yaml:"RpgPciWarningAcknowledgementEvent"`
					RpgPciWarningFoundUnacknowledgedEvent      string `yaml:"RpgPciWarningFoundUnacknowledgedEvent"`
					RpgServiceAvailability                     bool   `yaml:"RpgServiceAvailability"`
					RpgServiceLatency                          bool   `yaml:"RpgServiceLatency"`
					SessionStart                               string `yaml:"SessionStart"`
					SocialAuth                                 string `yaml:"SocialAuth"`
					StartDesktopPurchase                       string `yaml:"StartDesktopPurchase"`
					TermsOfServiceEvent                        string `yaml:"TermsOfServiceEvent"`
					VivoxCallStats                             string `yaml:"VivoxCallStats"`
					VoiceChatSLI                               string `yaml:"VoiceChatSLI"`
				} `yaml:"analytics_platform_allowed_events"`
				Endpoint struct {
					SendDeprecated bool `yaml:"send_deprecated"`
					SendFailure    bool `yaml:"send_failure"`
					SendSuccess    bool `yaml:"send_success"`
				} `yaml:"endpoint"`
				HeartbeatInterval         int  `yaml:"heartbeat_interval"`
				HeartbeatProducts         bool `yaml:"heartbeat_products"`
				HeartbeatVoiceChatMetrics bool `yaml:"heartbeat_voice_chat_metrics"`
				MetricsEnabled            bool `yaml:"metrics_enabled"`
				NewrelicAllowedEvents     struct {
					ChatCommandTelemetry                       string `yaml:"ChatCommandTelemetry"`
					ChatDisconnectEvent                        bool   `yaml:"ChatDisconnectEvent"`
					ChatLoginState                             bool   `yaml:"ChatLoginState"`
					ChatStanzaResponse                         bool   `yaml:"ChatStanzaResponse"`
					ConsoleAccountSLI                          bool   `yaml:"ConsoleAccountSLI"`
					EndpointResult                             string `yaml:"EndpointResult"`
					ErrorLog                                   string `yaml:"ErrorLog"`
					ExternalWebRequest                         bool   `yaml:"ExternalWebRequest"`
					FirstPartyFulfillmentSLI                   string `yaml:"FirstPartyFulfillmentSLI"`
					LeagueRegionElection                       bool   `yaml:"LeagueRegionElection"`
					PayMobile                                  string `yaml:"PayMobile"`
					PlatformUIOperationalEvent                 string `yaml:"PlatformUIOperationalEvent"`
					PlatformUIRenderStats                      string `yaml:"PlatformUIRenderStats"`
					PlatformUIStatusEvent                      string `yaml:"PlatformUIStatusEvent"`
					PlayerAccountAliasSanitizedEvent           string `yaml:"PlayerAccountAliasSanitizedEvent"`
					PlayerSessionLifecycleSLI                  string `yaml:"PlayerSessionLifecycleSLI"`
					PlayerSessionLifecycleStatus               string `yaml:"PlayerSessionLifecycleStatus"`
					ProductIntegrationAppRepairEvent           bool   `yaml:"ProductIntegrationAppRepairEvent"`
					ProductIntegrationEvent                    bool   `yaml:"ProductIntegrationEvent"`
					RSOAuthLoginSteps                          bool   `yaml:"RSOAuthLoginSteps"`
					RSOMobileAccountDeletionUI                 bool   `yaml:"RSOMobileAccountDeletionUI"`
					RSOMobileUIEvent                           bool   `yaml:"RSOMobileUIEvent"`
					RSOSignOnAttempt                           bool   `yaml:"RSOSignOnAttempt"`
					RpgPciPenaltyNotificationPlayerActionEvent bool   `yaml:"RpgPciPenaltyNotificationPlayerActionEvent"`
					RpgPciReporterFeedbackAcknowledgedEvent    bool   `yaml:"RpgPciReporterFeedbackAcknowledgedEvent"`
					RpgPciReporterFeedbackReceivedEvent        bool   `yaml:"RpgPciReporterFeedbackReceivedEvent"`
					RpgPciWarningAcknowledgementEvent          bool   `yaml:"RpgPciWarningAcknowledgementEvent"`
					RpgPciWarningFoundUnacknowledgedEvent      bool   `yaml:"RpgPciWarningFoundUnacknowledgedEvent"`
					RpgServiceAvailability                     bool   `yaml:"RpgServiceAvailability"`
					RpgServiceLatency                          bool   `yaml:"RpgServiceLatency"`
					SessionStart                               string `yaml:"SessionStart"`
					SocialAuth                                 bool   `yaml:"SocialAuth"`
					StartDesktopPurchase                       string `yaml:"StartDesktopPurchase"`
					TermsOfServiceEvent                        bool   `yaml:"TermsOfServiceEvent"`
					VivoxCallStats                             string `yaml:"VivoxCallStats"`
					VoiceChatSLI                               string `yaml:"VoiceChatSLI"`
				} `yaml:"newrelic_allowed_events"`
				NewrelicAllowedMetrics struct {
					OnlineHeartbeat  bool   `yaml:"OnlineHeartbeat"`
					ProductLaunches  bool   `yaml:"ProductLaunches"`
					ProductSessions  bool   `yaml:"ProductSessions"`
					VoiceChatMetrics string `yaml:"VoiceChatMetrics"`
				} `yaml:"newrelic_allowed_metrics"`
				NewrelicEventsV2Enabled           bool `yaml:"newrelic_events_v2_enabled"`
				NewrelicMetricsV1Enabled          bool `yaml:"newrelic_metrics_v1_enabled"`
				NewrelicSchemalessEventsV2Enabled bool `yaml:"newrelic_schemaless_events_v2_enabled"`
				SingularV1Enabled                 bool `yaml:"singular_v1_enabled"`
			} `yaml:"telemetry"`
		} `yaml:"riotgamesapi"`
		RsoAuth struct {
			InstallIdentifier string `yaml:"install-identifier"`
		} `yaml:"rso-auth"`
		RunElectron   bool `yaml:"run-electron"`
		StartupConfig struct {
			LaunchOnComputerSetByDefault bool `yaml:"launch_on_computer_set_by_default"`
		} `yaml:"startup-config"`
		StaySignedInModalShown bool `yaml:"stay-signed-in-modal-shown"`
		Telemetry              struct {
			Ap struct {
				DisabledEvents struct {
					RiotclientPlatformUIStatusEventV1 bool `yaml:"riotclient__PlatformUIStatusEvent__v1"`
					RiotclientPluginStatsV7           bool `yaml:"riotclient__PluginStats__v7"`
					RiotclientTimeSampleV4            bool `yaml:"riotclient__TimeSample__v4"`
					RpgServiceAvailabilityV2          bool `yaml:"rpg__ServiceAvailability__v2"`
					RpgServiceLatencyV2               bool `yaml:"rpg__ServiceLatency__v2"`
					SchemalessExternalWebRequestV1    bool `yaml:"schemaless__ExternalWebRequest__v1"`
				} `yaml:"disabled_events"`
			} `yaml:"ap"`
			Endpoint struct {
				SendDeprecated bool `yaml:"send_deprecated"`
				SendFailure    bool `yaml:"send_failure"`
				SendSuccess    bool `yaml:"send_success"`
			} `yaml:"endpoint"`
			HeartbeatInterval                 int    `yaml:"heartbeat_interval"`
			HeartbeatProducts                 bool   `yaml:"heartbeat_products"`
			HeartbeatVoiceChatMetrics         bool   `yaml:"heartbeat_voice_chat_metrics"`
			InstallationID                    string `yaml:"installation_id"`
			MetricsEnabled                    bool   `yaml:"metrics_enabled"`
			NewrelicAPIKey                    string `yaml:"newrelic_api_key"`
			NewrelicEventsV2Enabled           bool   `yaml:"newrelic_events_v2_enabled"`
			NewrelicMetricsV1Enabled          bool   `yaml:"newrelic_metrics_v1_enabled"`
			NewrelicSchemalessEventsV2Enabled bool   `yaml:"newrelic_schemaless_events_v2_enabled"`
			SingularAPIKey                    string `yaml:"singular_api_key"`
			SingularCustomuseridURL           string `yaml:"singular_customuserid_url"`
			SingularEventURL                  string `yaml:"singular_event_url"`
			SingularLaunchURL                 string `yaml:"singular_launch_url"`
			SingularV1Enabled                 bool   `yaml:"singular_v1_enabled"`
		} `yaml:"telemetry"`
	} `yaml:"install"`
}
