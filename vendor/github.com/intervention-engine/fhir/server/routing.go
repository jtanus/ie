package server

import "github.com/gin-gonic/gin"

// RegisterController registers the CRUD routes (and middleware) for a FHIR resource
func RegisterController(name string, e *gin.Engine, m []gin.HandlerFunc, dal DataAccessLayer, config Config) {
	rc := NewResourceController(name, dal)
	rcBase := e.Group("/" + name)

	if len(m) > 0 {
		rcBase.Use(m...)
	}

	if config.UseSmartAuth {
		rcBase.Use(SmartAuthHandler(name))
	}

	rcBase.GET("", rc.IndexHandler)
	rcBase.POST("", rc.CreateHandler)
	rcBase.PUT("", rc.ConditionalUpdateHandler)
	rcBase.DELETE("", rc.ConditionalDeleteHandler)

	rcItem := rcBase.Group("/:id")
	rcItem.GET("", rc.ShowHandler)
	rcItem.PUT("", rc.UpdateHandler)
	rcItem.DELETE("", rc.DeleteHandler)
}

// RegisterRoutes registers the routes for each of the FHIR resources
func RegisterRoutes(e *gin.Engine, config map[string][]gin.HandlerFunc, dal DataAccessLayer, serverConfig Config) {

	// Batch Support
	batch := NewBatchController(dal)
	e.POST("/", batch.Post)

	// Resources

	RegisterController("Appointment", e, config["Appointment"], dal, serverConfig)
	RegisterController("ReferralRequest", e, config["ReferralRequest"], dal, serverConfig)
	RegisterController("Account", e, config["Account"], dal, serverConfig)
	RegisterController("Provenance", e, config["Provenance"], dal, serverConfig)
	RegisterController("Questionnaire", e, config["Questionnaire"], dal, serverConfig)
	RegisterController("ExplanationOfBenefit", e, config["ExplanationOfBenefit"], dal, serverConfig)
	RegisterController("DocumentManifest", e, config["DocumentManifest"], dal, serverConfig)
	RegisterController("Specimen", e, config["Specimen"], dal, serverConfig)
	RegisterController("AllergyIntolerance", e, config["AllergyIntolerance"], dal, serverConfig)
	RegisterController("CarePlan", e, config["CarePlan"], dal, serverConfig)
	RegisterController("Goal", e, config["Goal"], dal, serverConfig)
	RegisterController("StructureDefinition", e, config["StructureDefinition"], dal, serverConfig)
	RegisterController("EnrollmentRequest", e, config["EnrollmentRequest"], dal, serverConfig)
	RegisterController("EpisodeOfCare", e, config["EpisodeOfCare"], dal, serverConfig)
	RegisterController("OperationOutcome", e, config["OperationOutcome"], dal, serverConfig)
	RegisterController("Medication", e, config["Medication"], dal, serverConfig)
	RegisterController("Procedure", e, config["Procedure"], dal, serverConfig)
	RegisterController("List", e, config["List"], dal, serverConfig)
	RegisterController("ConceptMap", e, config["ConceptMap"], dal, serverConfig)
	RegisterController("Subscription", e, config["Subscription"], dal, serverConfig)
	RegisterController("ValueSet", e, config["ValueSet"], dal, serverConfig)
	RegisterController("OperationDefinition", e, config["OperationDefinition"], dal, serverConfig)
	RegisterController("DocumentReference", e, config["DocumentReference"], dal, serverConfig)
	RegisterController("Order", e, config["Order"], dal, serverConfig)
	RegisterController("Immunization", e, config["Immunization"], dal, serverConfig)
	RegisterController("Device", e, config["Device"], dal, serverConfig)
	RegisterController("VisionPrescription", e, config["VisionPrescription"], dal, serverConfig)
	RegisterController("Media", e, config["Media"], dal, serverConfig)
	RegisterController("Conformance", e, config["Conformance"], dal, serverConfig)
	RegisterController("ProcedureRequest", e, config["ProcedureRequest"], dal, serverConfig)
	RegisterController("EligibilityResponse", e, config["EligibilityResponse"], dal, serverConfig)
	RegisterController("DeviceUseRequest", e, config["DeviceUseRequest"], dal, serverConfig)
	RegisterController("DeviceMetric", e, config["DeviceMetric"], dal, serverConfig)
	RegisterController("Flag", e, config["Flag"], dal, serverConfig)
	RegisterController("RelatedPerson", e, config["RelatedPerson"], dal, serverConfig)
	RegisterController("SupplyRequest", e, config["SupplyRequest"], dal, serverConfig)
	RegisterController("Practitioner", e, config["Practitioner"], dal, serverConfig)
	RegisterController("AppointmentResponse", e, config["AppointmentResponse"], dal, serverConfig)
	RegisterController("Observation", e, config["Observation"], dal, serverConfig)
	RegisterController("MedicationAdministration", e, config["MedicationAdministration"], dal, serverConfig)
	RegisterController("Slot", e, config["Slot"], dal, serverConfig)
	RegisterController("EnrollmentResponse", e, config["EnrollmentResponse"], dal, serverConfig)
	RegisterController("Binary", e, config["Binary"], dal, serverConfig)
	RegisterController("MedicationStatement", e, config["MedicationStatement"], dal, serverConfig)
	RegisterController("Person", e, config["Person"], dal, serverConfig)
	RegisterController("Contract", e, config["Contract"], dal, serverConfig)
	RegisterController("CommunicationRequest", e, config["CommunicationRequest"], dal, serverConfig)
	RegisterController("RiskAssessment", e, config["RiskAssessment"], dal, serverConfig)
	RegisterController("TestScript", e, config["TestScript"], dal, serverConfig)
	RegisterController("Basic", e, config["Basic"], dal, serverConfig)
	RegisterController("Group", e, config["Group"], dal, serverConfig)
	RegisterController("PaymentNotice", e, config["PaymentNotice"], dal, serverConfig)
	RegisterController("Organization", e, config["Organization"], dal, serverConfig)
	RegisterController("ImplementationGuide", e, config["ImplementationGuide"], dal, serverConfig)
	RegisterController("ClaimResponse", e, config["ClaimResponse"], dal, serverConfig)
	RegisterController("EligibilityRequest", e, config["EligibilityRequest"], dal, serverConfig)
	RegisterController("ProcessRequest", e, config["ProcessRequest"], dal, serverConfig)
	RegisterController("MedicationDispense", e, config["MedicationDispense"], dal, serverConfig)
	RegisterController("DiagnosticReport", e, config["DiagnosticReport"], dal, serverConfig)
	RegisterController("ImagingStudy", e, config["ImagingStudy"], dal, serverConfig)
	RegisterController("ImagingObjectSelection", e, config["ImagingObjectSelection"], dal, serverConfig)
	RegisterController("HealthcareService", e, config["HealthcareService"], dal, serverConfig)
	RegisterController("DataElement", e, config["DataElement"], dal, serverConfig)
	RegisterController("DeviceComponent", e, config["DeviceComponent"], dal, serverConfig)
	RegisterController("FamilyMemberHistory", e, config["FamilyMemberHistory"], dal, serverConfig)
	RegisterController("NutritionOrder", e, config["NutritionOrder"], dal, serverConfig)
	RegisterController("Encounter", e, config["Encounter"], dal, serverConfig)
	RegisterController("Substance", e, config["Substance"], dal, serverConfig)
	RegisterController("AuditEvent", e, config["AuditEvent"], dal, serverConfig)
	RegisterController("MedicationOrder", e, config["MedicationOrder"], dal, serverConfig)
	RegisterController("SearchParameter", e, config["SearchParameter"], dal, serverConfig)
	RegisterController("PaymentReconciliation", e, config["PaymentReconciliation"], dal, serverConfig)
	RegisterController("Communication", e, config["Communication"], dal, serverConfig)
	RegisterController("Condition", e, config["Condition"], dal, serverConfig)
	RegisterController("Composition", e, config["Composition"], dal, serverConfig)
	RegisterController("DetectedIssue", e, config["DetectedIssue"], dal, serverConfig)
	RegisterController("Bundle", e, config["Bundle"], dal, serverConfig)
	RegisterController("DiagnosticOrder", e, config["DiagnosticOrder"], dal, serverConfig)
	RegisterController("Patient", e, config["Patient"], dal, serverConfig)
	RegisterController("OrderResponse", e, config["OrderResponse"], dal, serverConfig)
	RegisterController("Coverage", e, config["Coverage"], dal, serverConfig)
	RegisterController("QuestionnaireResponse", e, config["QuestionnaireResponse"], dal, serverConfig)
	RegisterController("DeviceUseStatement", e, config["DeviceUseStatement"], dal, serverConfig)
	RegisterController("ProcessResponse", e, config["ProcessResponse"], dal, serverConfig)
	RegisterController("NamingSystem", e, config["NamingSystem"], dal, serverConfig)
	RegisterController("Schedule", e, config["Schedule"], dal, serverConfig)
	RegisterController("SupplyDelivery", e, config["SupplyDelivery"], dal, serverConfig)
	RegisterController("ClinicalImpression", e, config["ClinicalImpression"], dal, serverConfig)
	RegisterController("MessageHeader", e, config["MessageHeader"], dal, serverConfig)
	RegisterController("Claim", e, config["Claim"], dal, serverConfig)
	RegisterController("ImmunizationRecommendation", e, config["ImmunizationRecommendation"], dal, serverConfig)
	RegisterController("Location", e, config["Location"], dal, serverConfig)
	RegisterController("BodySite", e, config["BodySite"], dal, serverConfig)
}
