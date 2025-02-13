// Code generated by "stringer -type=APIErrorCode -trimprefix=Err api-errors.go"; DO NOT EDIT.

package cmd

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ErrNone-0]
	_ = x[ErrAccessDenied-1]
	_ = x[ErrBadDigest-2]
	_ = x[ErrEntityTooSmall-3]
	_ = x[ErrEntityTooLarge-4]
	_ = x[ErrPolicyTooLarge-5]
	_ = x[ErrIncompleteBody-6]
	_ = x[ErrInternalError-7]
	_ = x[ErrInvalidAccessKeyID-8]
	_ = x[ErrInvalidBucketName-9]
	_ = x[ErrInvalidDigest-10]
	_ = x[ErrInvalidRange-11]
	_ = x[ErrInvalidRangePartNumber-12]
	_ = x[ErrInvalidCopyPartRange-13]
	_ = x[ErrInvalidCopyPartRangeSource-14]
	_ = x[ErrInvalidMaxKeys-15]
	_ = x[ErrInvalidEncodingMethod-16]
	_ = x[ErrInvalidMaxUploads-17]
	_ = x[ErrInvalidMaxParts-18]
	_ = x[ErrInvalidPartNumberMarker-19]
	_ = x[ErrInvalidPartNumber-20]
	_ = x[ErrInvalidRequestBody-21]
	_ = x[ErrInvalidCopySource-22]
	_ = x[ErrInvalidMetadataDirective-23]
	_ = x[ErrInvalidCopyDest-24]
	_ = x[ErrInvalidPolicyDocument-25]
	_ = x[ErrInvalidObjectState-26]
	_ = x[ErrMalformedXML-27]
	_ = x[ErrMissingContentLength-28]
	_ = x[ErrMissingContentMD5-29]
	_ = x[ErrMissingRequestBodyError-30]
	_ = x[ErrMissingSecurityHeader-31]
	_ = x[ErrNoSuchBucket-32]
	_ = x[ErrNoSuchBucketPolicy-33]
	_ = x[ErrNoSuchBucketLifecycle-34]
	_ = x[ErrNoSuchLifecycleConfiguration-35]
	_ = x[ErrNoSuchBucketSSEConfig-36]
	_ = x[ErrNoSuchCORSConfiguration-37]
	_ = x[ErrNoSuchWebsiteConfiguration-38]
	_ = x[ErrReplicationConfigurationNotFoundError-39]
	_ = x[ErrRemoteDestinationNotFoundError-40]
	_ = x[ErrReplicationDestinationMissingLock-41]
	_ = x[ErrRemoteTargetNotFoundError-42]
	_ = x[ErrReplicationRemoteConnectionError-43]
	_ = x[ErrReplicationBandwidthLimitError-44]
	_ = x[ErrBucketRemoteIdenticalToSource-45]
	_ = x[ErrBucketRemoteAlreadyExists-46]
	_ = x[ErrBucketRemoteLabelInUse-47]
	_ = x[ErrBucketRemoteArnTypeInvalid-48]
	_ = x[ErrBucketRemoteArnInvalid-49]
	_ = x[ErrBucketRemoteRemoveDisallowed-50]
	_ = x[ErrRemoteTargetNotVersionedError-51]
	_ = x[ErrReplicationSourceNotVersionedError-52]
	_ = x[ErrReplicationNeedsVersioningError-53]
	_ = x[ErrReplicationBucketNeedsVersioningError-54]
	_ = x[ErrReplicationNoMatchingRuleError-55]
	_ = x[ErrObjectRestoreAlreadyInProgress-56]
	_ = x[ErrNoSuchKey-57]
	_ = x[ErrNoSuchUpload-58]
	_ = x[ErrInvalidVersionID-59]
	_ = x[ErrNoSuchVersion-60]
	_ = x[ErrNotImplemented-61]
	_ = x[ErrPreconditionFailed-62]
	_ = x[ErrRequestTimeTooSkewed-63]
	_ = x[ErrSignatureDoesNotMatch-64]
	_ = x[ErrMethodNotAllowed-65]
	_ = x[ErrInvalidPart-66]
	_ = x[ErrInvalidPartOrder-67]
	_ = x[ErrAuthorizationHeaderMalformed-68]
	_ = x[ErrMalformedPOSTRequest-69]
	_ = x[ErrPOSTFileRequired-70]
	_ = x[ErrSignatureVersionNotSupported-71]
	_ = x[ErrBucketNotEmpty-72]
	_ = x[ErrAllAccessDisabled-73]
	_ = x[ErrMalformedPolicy-74]
	_ = x[ErrMissingFields-75]
	_ = x[ErrMissingCredTag-76]
	_ = x[ErrCredMalformed-77]
	_ = x[ErrInvalidRegion-78]
	_ = x[ErrInvalidServiceS3-79]
	_ = x[ErrInvalidServiceSTS-80]
	_ = x[ErrInvalidRequestVersion-81]
	_ = x[ErrMissingSignTag-82]
	_ = x[ErrMissingSignHeadersTag-83]
	_ = x[ErrMalformedDate-84]
	_ = x[ErrMalformedPresignedDate-85]
	_ = x[ErrMalformedCredentialDate-86]
	_ = x[ErrMalformedCredentialRegion-87]
	_ = x[ErrMalformedExpires-88]
	_ = x[ErrNegativeExpires-89]
	_ = x[ErrAuthHeaderEmpty-90]
	_ = x[ErrExpiredPresignRequest-91]
	_ = x[ErrRequestNotReadyYet-92]
	_ = x[ErrUnsignedHeaders-93]
	_ = x[ErrMissingDateHeader-94]
	_ = x[ErrInvalidQuerySignatureAlgo-95]
	_ = x[ErrInvalidQueryParams-96]
	_ = x[ErrBucketAlreadyOwnedByYou-97]
	_ = x[ErrInvalidDuration-98]
	_ = x[ErrBucketAlreadyExists-99]
	_ = x[ErrMetadataTooLarge-100]
	_ = x[ErrUnsupportedMetadata-101]
	_ = x[ErrMaximumExpires-102]
	_ = x[ErrSlowDown-103]
	_ = x[ErrInvalidPrefixMarker-104]
	_ = x[ErrBadRequest-105]
	_ = x[ErrKeyTooLongError-106]
	_ = x[ErrInvalidBucketObjectLockConfiguration-107]
	_ = x[ErrObjectLockConfigurationNotFound-108]
	_ = x[ErrObjectLockConfigurationNotAllowed-109]
	_ = x[ErrNoSuchObjectLockConfiguration-110]
	_ = x[ErrObjectLocked-111]
	_ = x[ErrInvalidRetentionDate-112]
	_ = x[ErrPastObjectLockRetainDate-113]
	_ = x[ErrUnknownWORMModeDirective-114]
	_ = x[ErrBucketTaggingNotFound-115]
	_ = x[ErrObjectLockInvalidHeaders-116]
	_ = x[ErrInvalidTagDirective-117]
	_ = x[ErrInvalidEncryptionMethod-118]
	_ = x[ErrInsecureSSECustomerRequest-119]
	_ = x[ErrSSEMultipartEncrypted-120]
	_ = x[ErrSSEEncryptedObject-121]
	_ = x[ErrInvalidEncryptionParameters-122]
	_ = x[ErrInvalidSSECustomerAlgorithm-123]
	_ = x[ErrInvalidSSECustomerKey-124]
	_ = x[ErrMissingSSECustomerKey-125]
	_ = x[ErrMissingSSECustomerKeyMD5-126]
	_ = x[ErrSSECustomerKeyMD5Mismatch-127]
	_ = x[ErrInvalidSSECustomerParameters-128]
	_ = x[ErrIncompatibleEncryptionMethod-129]
	_ = x[ErrKMSNotConfigured-130]
	_ = x[ErrNoAccessKey-131]
	_ = x[ErrInvalidToken-132]
	_ = x[ErrEventNotification-133]
	_ = x[ErrARNNotification-134]
	_ = x[ErrRegionNotification-135]
	_ = x[ErrOverlappingFilterNotification-136]
	_ = x[ErrFilterNameInvalid-137]
	_ = x[ErrFilterNamePrefix-138]
	_ = x[ErrFilterNameSuffix-139]
	_ = x[ErrFilterValueInvalid-140]
	_ = x[ErrOverlappingConfigs-141]
	_ = x[ErrUnsupportedNotification-142]
	_ = x[ErrContentSHA256Mismatch-143]
	_ = x[ErrReadQuorum-144]
	_ = x[ErrWriteQuorum-145]
	_ = x[ErrStorageFull-146]
	_ = x[ErrRequestBodyParse-147]
	_ = x[ErrObjectExistsAsDirectory-148]
	_ = x[ErrInvalidObjectName-149]
	_ = x[ErrInvalidObjectNamePrefixSlash-150]
	_ = x[ErrInvalidResourceName-151]
	_ = x[ErrServerNotInitialized-152]
	_ = x[ErrOperationTimedOut-153]
	_ = x[ErrClientDisconnected-154]
	_ = x[ErrOperationMaxedOut-155]
	_ = x[ErrInvalidRequest-156]
	_ = x[ErrTransitionStorageClassNotFoundError-157]
	_ = x[ErrInvalidStorageClass-158]
	_ = x[ErrBackendDown-159]
	_ = x[ErrMalformedJSON-160]
	_ = x[ErrAdminNoSuchUser-161]
	_ = x[ErrAdminNoSuchGroup-162]
	_ = x[ErrAdminGroupNotEmpty-163]
	_ = x[ErrAdminNoSuchPolicy-164]
	_ = x[ErrAdminInvalidArgument-165]
	_ = x[ErrAdminInvalidAccessKey-166]
	_ = x[ErrAdminInvalidSecretKey-167]
	_ = x[ErrAdminConfigNoQuorum-168]
	_ = x[ErrAdminConfigTooLarge-169]
	_ = x[ErrAdminConfigBadJSON-170]
	_ = x[ErrAdminConfigDuplicateKeys-171]
	_ = x[ErrAdminCredentialsMismatch-172]
	_ = x[ErrInsecureClientRequest-173]
	_ = x[ErrObjectTampered-174]
	_ = x[ErrSiteReplicationInvalidRequest-175]
	_ = x[ErrSiteReplicationPeerResp-176]
	_ = x[ErrSiteReplicationBackendIssue-177]
	_ = x[ErrSiteReplicationServiceAccountError-178]
	_ = x[ErrSiteReplicationBucketConfigError-179]
	_ = x[ErrSiteReplicationBucketMetaError-180]
	_ = x[ErrSiteReplicationIAMError-181]
	_ = x[ErrAdminBucketQuotaExceeded-182]
	_ = x[ErrAdminNoSuchQuotaConfiguration-183]
	_ = x[ErrHealNotImplemented-184]
	_ = x[ErrHealNoSuchProcess-185]
	_ = x[ErrHealInvalidClientToken-186]
	_ = x[ErrHealMissingBucket-187]
	_ = x[ErrHealAlreadyRunning-188]
	_ = x[ErrHealOverlappingPaths-189]
	_ = x[ErrIncorrectContinuationToken-190]
	_ = x[ErrEmptyRequestBody-191]
	_ = x[ErrUnsupportedFunction-192]
	_ = x[ErrInvalidExpressionType-193]
	_ = x[ErrBusy-194]
	_ = x[ErrUnauthorizedAccess-195]
	_ = x[ErrExpressionTooLong-196]
	_ = x[ErrIllegalSQLFunctionArgument-197]
	_ = x[ErrInvalidKeyPath-198]
	_ = x[ErrInvalidCompressionFormat-199]
	_ = x[ErrInvalidFileHeaderInfo-200]
	_ = x[ErrInvalidJSONType-201]
	_ = x[ErrInvalidQuoteFields-202]
	_ = x[ErrInvalidRequestParameter-203]
	_ = x[ErrInvalidDataType-204]
	_ = x[ErrInvalidTextEncoding-205]
	_ = x[ErrInvalidDataSource-206]
	_ = x[ErrInvalidTableAlias-207]
	_ = x[ErrMissingRequiredParameter-208]
	_ = x[ErrObjectSerializationConflict-209]
	_ = x[ErrUnsupportedSQLOperation-210]
	_ = x[ErrUnsupportedSQLStructure-211]
	_ = x[ErrUnsupportedSyntax-212]
	_ = x[ErrUnsupportedRangeHeader-213]
	_ = x[ErrLexerInvalidChar-214]
	_ = x[ErrLexerInvalidOperator-215]
	_ = x[ErrLexerInvalidLiteral-216]
	_ = x[ErrLexerInvalidIONLiteral-217]
	_ = x[ErrParseExpectedDatePart-218]
	_ = x[ErrParseExpectedKeyword-219]
	_ = x[ErrParseExpectedTokenType-220]
	_ = x[ErrParseExpected2TokenTypes-221]
	_ = x[ErrParseExpectedNumber-222]
	_ = x[ErrParseExpectedRightParenBuiltinFunctionCall-223]
	_ = x[ErrParseExpectedTypeName-224]
	_ = x[ErrParseExpectedWhenClause-225]
	_ = x[ErrParseUnsupportedToken-226]
	_ = x[ErrParseUnsupportedLiteralsGroupBy-227]
	_ = x[ErrParseExpectedMember-228]
	_ = x[ErrParseUnsupportedSelect-229]
	_ = x[ErrParseUnsupportedCase-230]
	_ = x[ErrParseUnsupportedCaseClause-231]
	_ = x[ErrParseUnsupportedAlias-232]
	_ = x[ErrParseUnsupportedSyntax-233]
	_ = x[ErrParseUnknownOperator-234]
	_ = x[ErrParseMissingIdentAfterAt-235]
	_ = x[ErrParseUnexpectedOperator-236]
	_ = x[ErrParseUnexpectedTerm-237]
	_ = x[ErrParseUnexpectedToken-238]
	_ = x[ErrParseUnexpectedKeyword-239]
	_ = x[ErrParseExpectedExpression-240]
	_ = x[ErrParseExpectedLeftParenAfterCast-241]
	_ = x[ErrParseExpectedLeftParenValueConstructor-242]
	_ = x[ErrParseExpectedLeftParenBuiltinFunctionCall-243]
	_ = x[ErrParseExpectedArgumentDelimiter-244]
	_ = x[ErrParseCastArity-245]
	_ = x[ErrParseInvalidTypeParam-246]
	_ = x[ErrParseEmptySelect-247]
	_ = x[ErrParseSelectMissingFrom-248]
	_ = x[ErrParseExpectedIdentForGroupName-249]
	_ = x[ErrParseExpectedIdentForAlias-250]
	_ = x[ErrParseUnsupportedCallWithStar-251]
	_ = x[ErrParseNonUnaryAgregateFunctionCall-252]
	_ = x[ErrParseMalformedJoin-253]
	_ = x[ErrParseExpectedIdentForAt-254]
	_ = x[ErrParseAsteriskIsNotAloneInSelectList-255]
	_ = x[ErrParseCannotMixSqbAndWildcardInSelectList-256]
	_ = x[ErrParseInvalidContextForWildcardInSelectList-257]
	_ = x[ErrIncorrectSQLFunctionArgumentType-258]
	_ = x[ErrValueParseFailure-259]
	_ = x[ErrEvaluatorInvalidArguments-260]
	_ = x[ErrIntegerOverflow-261]
	_ = x[ErrLikeInvalidInputs-262]
	_ = x[ErrCastFailed-263]
	_ = x[ErrInvalidCast-264]
	_ = x[ErrEvaluatorInvalidTimestampFormatPattern-265]
	_ = x[ErrEvaluatorInvalidTimestampFormatPatternSymbolForParsing-266]
	_ = x[ErrEvaluatorTimestampFormatPatternDuplicateFields-267]
	_ = x[ErrEvaluatorTimestampFormatPatternHourClockAmPmMismatch-268]
	_ = x[ErrEvaluatorUnterminatedTimestampFormatPatternToken-269]
	_ = x[ErrEvaluatorInvalidTimestampFormatPatternToken-270]
	_ = x[ErrEvaluatorInvalidTimestampFormatPatternSymbol-271]
	_ = x[ErrEvaluatorBindingDoesNotExist-272]
	_ = x[ErrMissingHeaders-273]
	_ = x[ErrInvalidColumnIndex-274]
	_ = x[ErrAdminConfigNotificationTargetsFailed-275]
	_ = x[ErrAdminProfilerNotEnabled-276]
	_ = x[ErrInvalidDecompressedSize-277]
	_ = x[ErrAddUserInvalidArgument-278]
	_ = x[ErrAdminAccountNotEligible-279]
	_ = x[ErrAccountNotEligible-280]
	_ = x[ErrAdminServiceAccountNotFound-281]
	_ = x[ErrPostPolicyConditionInvalidFormat-282]
}

const _APIErrorCode_name = "NoneAccessDeniedBadDigestEntityTooSmallEntityTooLargePolicyTooLargeIncompleteBodyInternalErrorInvalidAccessKeyIDInvalidBucketNameInvalidDigestInvalidRangeInvalidRangePartNumberInvalidCopyPartRangeInvalidCopyPartRangeSourceInvalidMaxKeysInvalidEncodingMethodInvalidMaxUploadsInvalidMaxPartsInvalidPartNumberMarkerInvalidPartNumberInvalidRequestBodyInvalidCopySourceInvalidMetadataDirectiveInvalidCopyDestInvalidPolicyDocumentInvalidObjectStateMalformedXMLMissingContentLengthMissingContentMD5MissingRequestBodyErrorMissingSecurityHeaderNoSuchBucketNoSuchBucketPolicyNoSuchBucketLifecycleNoSuchLifecycleConfigurationNoSuchBucketSSEConfigNoSuchCORSConfigurationNoSuchWebsiteConfigurationReplicationConfigurationNotFoundErrorRemoteDestinationNotFoundErrorReplicationDestinationMissingLockRemoteTargetNotFoundErrorReplicationRemoteConnectionErrorReplicationBandwidthLimitErrorBucketRemoteIdenticalToSourceBucketRemoteAlreadyExistsBucketRemoteLabelInUseBucketRemoteArnTypeInvalidBucketRemoteArnInvalidBucketRemoteRemoveDisallowedRemoteTargetNotVersionedErrorReplicationSourceNotVersionedErrorReplicationNeedsVersioningErrorReplicationBucketNeedsVersioningErrorReplicationNoMatchingRuleErrorObjectRestoreAlreadyInProgressNoSuchKeyNoSuchUploadInvalidVersionIDNoSuchVersionNotImplementedPreconditionFailedRequestTimeTooSkewedSignatureDoesNotMatchMethodNotAllowedInvalidPartInvalidPartOrderAuthorizationHeaderMalformedMalformedPOSTRequestPOSTFileRequiredSignatureVersionNotSupportedBucketNotEmptyAllAccessDisabledMalformedPolicyMissingFieldsMissingCredTagCredMalformedInvalidRegionInvalidServiceS3InvalidServiceSTSInvalidRequestVersionMissingSignTagMissingSignHeadersTagMalformedDateMalformedPresignedDateMalformedCredentialDateMalformedCredentialRegionMalformedExpiresNegativeExpiresAuthHeaderEmptyExpiredPresignRequestRequestNotReadyYetUnsignedHeadersMissingDateHeaderInvalidQuerySignatureAlgoInvalidQueryParamsBucketAlreadyOwnedByYouInvalidDurationBucketAlreadyExistsMetadataTooLargeUnsupportedMetadataMaximumExpiresSlowDownInvalidPrefixMarkerBadRequestKeyTooLongErrorInvalidBucketObjectLockConfigurationObjectLockConfigurationNotFoundObjectLockConfigurationNotAllowedNoSuchObjectLockConfigurationObjectLockedInvalidRetentionDatePastObjectLockRetainDateUnknownWORMModeDirectiveBucketTaggingNotFoundObjectLockInvalidHeadersInvalidTagDirectiveInvalidEncryptionMethodInsecureSSECustomerRequestSSEMultipartEncryptedSSEEncryptedObjectInvalidEncryptionParametersInvalidSSECustomerAlgorithmInvalidSSECustomerKeyMissingSSECustomerKeyMissingSSECustomerKeyMD5SSECustomerKeyMD5MismatchInvalidSSECustomerParametersIncompatibleEncryptionMethodKMSNotConfiguredNoAccessKeyInvalidTokenEventNotificationARNNotificationRegionNotificationOverlappingFilterNotificationFilterNameInvalidFilterNamePrefixFilterNameSuffixFilterValueInvalidOverlappingConfigsUnsupportedNotificationContentSHA256MismatchReadQuorumWriteQuorumStorageFullRequestBodyParseObjectExistsAsDirectoryInvalidObjectNameInvalidObjectNamePrefixSlashInvalidResourceNameServerNotInitializedOperationTimedOutClientDisconnectedOperationMaxedOutInvalidRequestTransitionStorageClassNotFoundErrorInvalidStorageClassBackendDownMalformedJSONAdminNoSuchUserAdminNoSuchGroupAdminGroupNotEmptyAdminNoSuchPolicyAdminInvalidArgumentAdminInvalidAccessKeyAdminInvalidSecretKeyAdminConfigNoQuorumAdminConfigTooLargeAdminConfigBadJSONAdminConfigDuplicateKeysAdminCredentialsMismatchInsecureClientRequestObjectTamperedSiteReplicationInvalidRequestSiteReplicationPeerRespSiteReplicationBackendIssueSiteReplicationServiceAccountErrorSiteReplicationBucketConfigErrorSiteReplicationBucketMetaErrorSiteReplicationIAMErrorAdminBucketQuotaExceededAdminNoSuchQuotaConfigurationHealNotImplementedHealNoSuchProcessHealInvalidClientTokenHealMissingBucketHealAlreadyRunningHealOverlappingPathsIncorrectContinuationTokenEmptyRequestBodyUnsupportedFunctionInvalidExpressionTypeBusyUnauthorizedAccessExpressionTooLongIllegalSQLFunctionArgumentInvalidKeyPathInvalidCompressionFormatInvalidFileHeaderInfoInvalidJSONTypeInvalidQuoteFieldsInvalidRequestParameterInvalidDataTypeInvalidTextEncodingInvalidDataSourceInvalidTableAliasMissingRequiredParameterObjectSerializationConflictUnsupportedSQLOperationUnsupportedSQLStructureUnsupportedSyntaxUnsupportedRangeHeaderLexerInvalidCharLexerInvalidOperatorLexerInvalidLiteralLexerInvalidIONLiteralParseExpectedDatePartParseExpectedKeywordParseExpectedTokenTypeParseExpected2TokenTypesParseExpectedNumberParseExpectedRightParenBuiltinFunctionCallParseExpectedTypeNameParseExpectedWhenClauseParseUnsupportedTokenParseUnsupportedLiteralsGroupByParseExpectedMemberParseUnsupportedSelectParseUnsupportedCaseParseUnsupportedCaseClauseParseUnsupportedAliasParseUnsupportedSyntaxParseUnknownOperatorParseMissingIdentAfterAtParseUnexpectedOperatorParseUnexpectedTermParseUnexpectedTokenParseUnexpectedKeywordParseExpectedExpressionParseExpectedLeftParenAfterCastParseExpectedLeftParenValueConstructorParseExpectedLeftParenBuiltinFunctionCallParseExpectedArgumentDelimiterParseCastArityParseInvalidTypeParamParseEmptySelectParseSelectMissingFromParseExpectedIdentForGroupNameParseExpectedIdentForAliasParseUnsupportedCallWithStarParseNonUnaryAgregateFunctionCallParseMalformedJoinParseExpectedIdentForAtParseAsteriskIsNotAloneInSelectListParseCannotMixSqbAndWildcardInSelectListParseInvalidContextForWildcardInSelectListIncorrectSQLFunctionArgumentTypeValueParseFailureEvaluatorInvalidArgumentsIntegerOverflowLikeInvalidInputsCastFailedInvalidCastEvaluatorInvalidTimestampFormatPatternEvaluatorInvalidTimestampFormatPatternSymbolForParsingEvaluatorTimestampFormatPatternDuplicateFieldsEvaluatorTimestampFormatPatternHourClockAmPmMismatchEvaluatorUnterminatedTimestampFormatPatternTokenEvaluatorInvalidTimestampFormatPatternTokenEvaluatorInvalidTimestampFormatPatternSymbolEvaluatorBindingDoesNotExistMissingHeadersInvalidColumnIndexAdminConfigNotificationTargetsFailedAdminProfilerNotEnabledInvalidDecompressedSizeAddUserInvalidArgumentAdminAccountNotEligibleAccountNotEligibleAdminServiceAccountNotFoundPostPolicyConditionInvalidFormat"

var _APIErrorCode_index = [...]uint16{0, 4, 16, 25, 39, 53, 67, 81, 94, 112, 129, 142, 154, 176, 196, 222, 236, 257, 274, 289, 312, 329, 347, 364, 388, 403, 424, 442, 454, 474, 491, 514, 535, 547, 565, 586, 614, 635, 658, 684, 721, 751, 784, 809, 841, 871, 900, 925, 947, 973, 995, 1023, 1052, 1086, 1117, 1154, 1184, 1214, 1223, 1235, 1251, 1264, 1278, 1296, 1316, 1337, 1353, 1364, 1380, 1408, 1428, 1444, 1472, 1486, 1503, 1518, 1531, 1545, 1558, 1571, 1587, 1604, 1625, 1639, 1660, 1673, 1695, 1718, 1743, 1759, 1774, 1789, 1810, 1828, 1843, 1860, 1885, 1903, 1926, 1941, 1960, 1976, 1995, 2009, 2017, 2036, 2046, 2061, 2097, 2128, 2161, 2190, 2202, 2222, 2246, 2270, 2291, 2315, 2334, 2357, 2383, 2404, 2422, 2449, 2476, 2497, 2518, 2542, 2567, 2595, 2623, 2639, 2650, 2662, 2679, 2694, 2712, 2741, 2758, 2774, 2790, 2808, 2826, 2849, 2870, 2880, 2891, 2902, 2918, 2941, 2958, 2986, 3005, 3025, 3042, 3060, 3077, 3091, 3126, 3145, 3156, 3169, 3184, 3200, 3218, 3235, 3255, 3276, 3297, 3316, 3335, 3353, 3377, 3401, 3422, 3436, 3465, 3488, 3515, 3549, 3581, 3611, 3634, 3658, 3687, 3705, 3722, 3744, 3761, 3779, 3799, 3825, 3841, 3860, 3881, 3885, 3903, 3920, 3946, 3960, 3984, 4005, 4020, 4038, 4061, 4076, 4095, 4112, 4129, 4153, 4180, 4203, 4226, 4243, 4265, 4281, 4301, 4320, 4342, 4363, 4383, 4405, 4429, 4448, 4490, 4511, 4534, 4555, 4586, 4605, 4627, 4647, 4673, 4694, 4716, 4736, 4760, 4783, 4802, 4822, 4844, 4867, 4898, 4936, 4977, 5007, 5021, 5042, 5058, 5080, 5110, 5136, 5164, 5197, 5215, 5238, 5273, 5313, 5355, 5387, 5404, 5429, 5444, 5461, 5471, 5482, 5520, 5574, 5620, 5672, 5720, 5763, 5807, 5835, 5849, 5867, 5903, 5926, 5949, 5971, 5994, 6012, 6039, 6071}

func (i APIErrorCode) String() string {
	if i < 0 || i >= APIErrorCode(len(_APIErrorCode_index)-1) {
		return "APIErrorCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _APIErrorCode_name[_APIErrorCode_index[i]:_APIErrorCode_index[i+1]]
}
