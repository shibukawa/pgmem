package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_GetAdditionalPinLimit(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_GetAdditionalPinLimit[0]))
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_GetAdditionalPinLimit[1]))
	v9 = v4 - v6 - int32(8)
	if base.Ui32(v9) <= base.Ui32(v4) {
		v12 = v9
	} else {
		v12 = int32(0)
	}
	return v12
}
func F_GetBackendTypeDesc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	if base.Ui32(l0) <= base.Ui32(int32(17)) {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_c_F_GetBackendTypeDesc[0])))
		v11 = v10
	} else {
		v11 = int32(_a_F_GetBackendTypeDesc_0)
	}
	return v11
}
func F_GetBackgroundWorkerPid(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_GetBackgroundWorkerPid[0]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_GetBackgroundWorkerPid[1]))
	v13 = F_LWLockAcquire(m, v9+int32(_a_F_GetBackgroundWorkerPid_0), int32(1))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
		v20 = v6 + v7*int32(1480)
		v21 = *(*int64)(unsafe.Add(mBase, uint32(v20)+24))
		if v17 == v21 {
			v24 = v20 + int32(16)
			v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
			if v25 != 0 {
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
				v37 = *(*int32)(unsafe.Add(mBase, _c_F_GetBackgroundWorkerPid[1]))
				F_LWLockRelease(m, v37+int32(_a_F_GetBackgroundWorkerPid_0))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					switch v35 + int32(1) {
					case 0:
						return int32(1)
					case 1:
						v49 = int32(2)
						return v49
					default:
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v35
						v49 = int32(0)
						return v49
					}
				}
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, _c_F_GetBackgroundWorkerPid[1]))
				F_LWLockRelease(m, v28+int32(_a_F_GetBackgroundWorkerPid_0))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					return int32(2)
				}
			}
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, _c_F_GetBackgroundWorkerPid[1]))
			F_LWLockRelease(m, v28+int32(_a_F_GetBackgroundWorkerPid_0))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				return int32(2)
			}
		}
	}
}
func F_GetComment(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	v8 = m.G0
	v10 = v8 - int32(160)
	m.G0 = v10
	F_ScanKeyInit(m, v10+int32(16), int32(1), int32(3), int32(184), l0)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		F_ScanKeyInit(m, v10-int32(-64), int32(2), int32(3), int32(184), l1)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			v30 = int32(3)
			F_ScanKeyInit(m, v10+int32(112), v30, v30, int32(65), l2)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				v37 = F_table_open(m, int32(2609), int32(1))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v37)+52))
					v40 = int32(0)
					v47 = F_systable_beginscan(m, v37, int32(2675), int32(1), v40, int32(3), v10+int32(16))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						v49 = F_systable_getnext(m, v47)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							if v49 == int32(0) {
								v117 = v40
								F_systable_endscan(m, v47)
								mBase = m.M
								v121 = m.ExcPending
								if v121 != 0 {
									return int32(0)
								} else {
									F_sequence_close(m, v37, int32(1))
									mBase = m.M
									v124 = m.ExcPending
									if v124 != 0 {
										return int32(0)
									} else {
										m.G0 = v10 + int32(160)
										return v117
									}
								}
							} else {
								v53 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
								v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53)+18)))
								if v54&int32(2044) != 0 {
									v57 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v57)
									v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+20)))
									if v59&int32(1) == v57 {
										v64 = *(*int32)(unsafe.Add(mBase, uint32(v39)+68))
										if int32(0) <= v64 {
											v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+22)))
											v69 = v53 + v67 + v64
											v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+74)))
											if v70 != int32(1) {
												v114 = v69
												v115 = F_text_to_cstring(m, v114)
												mBase = m.M
												v116 = m.ExcPending
												if v116 != 0 {
													return int32(0)
												} else {
													v117 = v115
													F_systable_endscan(m, v47)
													mBase = m.M
													v121 = m.ExcPending
													if v121 != 0 {
														return int32(0)
													} else {
														F_sequence_close(m, v37, int32(1))
														mBase = m.M
														v124 = m.ExcPending
														if v124 != 0 {
															return int32(0)
														} else {
															m.G0 = v10 + int32(160)
															return v117
														}
													}
												}
											} else {
												v73 = int32(*(*int16)(unsafe.Add(mBase, uint32(v39)+72)))
												switch v73&int32(_a_F_GetComment_0) - int32(1) {
												case 0:
													v78 = int32(*(*int8)(unsafe.Add(mBase, uint32(v69))))
													v114 = v78
													v115 = F_text_to_cstring(m, v114)
													mBase = m.M
													v116 = m.ExcPending
													if v116 != 0 {
														return int32(0)
													} else {
														v117 = v115
														F_systable_endscan(m, v47)
														mBase = m.M
														v121 = m.ExcPending
														if v121 != 0 {
															return int32(0)
														} else {
															F_sequence_close(m, v37, int32(1))
															mBase = m.M
															v124 = m.ExcPending
															if v124 != 0 {
																return int32(0)
															} else {
																m.G0 = v10 + int32(160)
																return v117
															}
														}
													}
												case 1:
													v79 = int32(*(*int16)(unsafe.Add(mBase, uint32(v69))))
													v114 = v79
													v115 = F_text_to_cstring(m, v114)
													mBase = m.M
													v116 = m.ExcPending
													if v116 != 0 {
														return int32(0)
													} else {
														v117 = v115
														F_systable_endscan(m, v47)
														mBase = m.M
														v121 = m.ExcPending
														if v121 != 0 {
															return int32(0)
														} else {
															F_sequence_close(m, v37, int32(1))
															mBase = m.M
															v124 = m.ExcPending
															if v124 != 0 {
																return int32(0)
															} else {
																m.G0 = v10 + int32(160)
																return v117
															}
														}
													}
												default:
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v84 = m.ExcPending
													if v84 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v10))) = v73
														F_errmsg_internal(m, int32(_a_F_GetComment_1), v10)
														mBase = m.M
														v88 = m.ExcPending
														if v88 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_GetComment_2), int32(70), int32(_a_F_GetComment_3))
															mBase = m.M
															v93 = m.ExcPending
															if v93 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												case 3:
													v80 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
													v114 = v80
													v115 = F_text_to_cstring(m, v114)
													mBase = m.M
													v116 = m.ExcPending
													if v116 != 0 {
														return int32(0)
													} else {
														v117 = v115
														F_systable_endscan(m, v47)
														mBase = m.M
														v121 = m.ExcPending
														if v121 != 0 {
															return int32(0)
														} else {
															F_sequence_close(m, v37, int32(1))
															mBase = m.M
															v124 = m.ExcPending
															if v124 != 0 {
																return int32(0)
															} else {
																m.G0 = v10 + int32(160)
																return v117
															}
														}
													}
												}
											}
										} else {
											v95 = F_nocachegetattr(m, v49, int32(4), v39)
											mBase = m.M
											v96 = m.ExcPending
											if v96 != 0 {
												return int32(0)
											} else {
												v114 = v95
												v115 = F_text_to_cstring(m, v114)
												mBase = m.M
												v116 = m.ExcPending
												if v116 != 0 {
													return int32(0)
												} else {
													v117 = v115
													F_systable_endscan(m, v47)
													mBase = m.M
													v121 = m.ExcPending
													if v121 != 0 {
														return int32(0)
													} else {
														F_sequence_close(m, v37, int32(1))
														mBase = m.M
														v124 = m.ExcPending
														if v124 != 0 {
															return int32(0)
														} else {
															m.G0 = v10 + int32(160)
															return v117
														}
													}
												}
											}
										}
									} else {
										v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+23)))
										if v97&int32(8) == int32(0) {
											v117 = v40
											F_systable_endscan(m, v47)
											mBase = m.M
											v121 = m.ExcPending
											if v121 != 0 {
												return int32(0)
											} else {
												F_sequence_close(m, v37, int32(1))
												mBase = m.M
												v124 = m.ExcPending
												if v124 != 0 {
													return int32(0)
												} else {
													m.G0 = v10 + int32(160)
													return v117
												}
											}
										} else {
											v103 = F_nocachegetattr(m, v49, int32(4), v39)
											mBase = m.M
											v104 = m.ExcPending
											if v104 != 0 {
												return int32(0)
											} else {
												v114 = v103
												v115 = F_text_to_cstring(m, v114)
												mBase = m.M
												v116 = m.ExcPending
												if v116 != 0 {
													return int32(0)
												} else {
													v117 = v115
													F_systable_endscan(m, v47)
													mBase = m.M
													v121 = m.ExcPending
													if v121 != 0 {
														return int32(0)
													} else {
														F_sequence_close(m, v37, int32(1))
														mBase = m.M
														v124 = m.ExcPending
														if v124 != 0 {
															return int32(0)
														} else {
															m.G0 = v10 + int32(160)
															return v117
														}
													}
												}
											}
										}
									}
								} else {
									v108 = F_getmissingattr(m, v39, int32(4), v10+int32(15))
									mBase = m.M
									v109 = m.ExcPending
									if v109 != 0 {
										return int32(0)
									} else {
										v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
										if v110&int32(1) != 0 {
											v117 = v40
											F_systable_endscan(m, v47)
											mBase = m.M
											v121 = m.ExcPending
											if v121 != 0 {
												return int32(0)
											} else {
												F_sequence_close(m, v37, int32(1))
												mBase = m.M
												v124 = m.ExcPending
												if v124 != 0 {
													return int32(0)
												} else {
													m.G0 = v10 + int32(160)
													return v117
												}
											}
										} else {
											v114 = v108
											v115 = F_text_to_cstring(m, v114)
											mBase = m.M
											v116 = m.ExcPending
											if v116 != 0 {
												return int32(0)
											} else {
												v117 = v115
												F_systable_endscan(m, v47)
												mBase = m.M
												v121 = m.ExcPending
												if v121 != 0 {
													return int32(0)
												} else {
													F_sequence_close(m, v37, int32(1))
													mBase = m.M
													v124 = m.ExcPending
													if v124 != 0 {
														return int32(0)
													} else {
														m.G0 = v10 + int32(160)
														return v117
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_GetExtensibleNodeMethods(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_GetExtensibleNodeMethods[0]))
	if v9 != 0 {
		v10 = int32(0)
		v12 = F_hash_search(m, v9, l0, v10, v10)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			if v12 != 0 {
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v12)+64))
				m.G0 = v6 + int32(16)
				return v33
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(67137668))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
						F_errmsg(m, int32(_a_F_GetExtensibleNodeMethods_0), v6)
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_GetExtensibleNodeMethods_1), int32(115), int32(_a_F_GetExtensibleNodeMethods_2))
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(67137668))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
				F_errmsg(m, int32(_a_F_GetExtensibleNodeMethods_0), v6)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_GetExtensibleNodeMethods_1), int32(115), int32(_a_F_GetExtensibleNodeMethods_2))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	}
}
func F_GetNonHistoricCatalogSnapshot(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v120 int32
	_ = v120
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	v2 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_GetNonHistoricCatalogSnapshot[0]))
	if v8 == v2 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v139
L2:
	;
	v129 = F_GetSnapshotData(m, int32(_a_F_GetNonHistoricCatalogSnapshot_0))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L31
	} else {
		goto L43
	}
L3:
	;
	v12 = int32(1)
	if l0 <= int32(2963) {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	v120 = *(*int32)(unsafe.Add(mBase, _c_F_GetNonHistoricCatalogSnapshot[0]))
	if v120 != 0 {
		v139 = v120
		goto L1
	} else {
		goto L42
	}
L5:
	;
	if v28 != 0 {
		goto L4
	} else {
		goto L16
	}
L6:
	;
	goto L5
L7:
	;
	v28 = int32(0)
	goto L6
L8:
	;
	if base.Ui32(l0-int32(2608)) < base.Ui32(int32(2)) {
		v28 = v12
		goto L6
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	switch l0 - int32(3592) {
	case 0, 4:
		v28 = v12
		goto L6
	case 1, 2, 3:
		goto L7
	default:
		goto L14
	}
L11:
	;
	if l0 == int32(1214) {
		v28 = v12
		goto L6
	} else {
		goto L12
	}
L12:
	;
	if l0 != int32(2396) {
		goto L7
	} else {
		goto L13
	}
L13:
	;
	v28 = v12
	goto L6
L14:
	;
	if l0 == int32(2964) {
		v28 = v12
		goto L6
	} else {
		goto L15
	}
L15:
	;
	goto L7
L16:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_GetNonHistoricCatalogSnapshot[1]))
	v32 = v30 - int32(1)
	if v32 < int32(0) {
		v63 = v2
		goto L17
	} else {
		goto L18
	}
L17:
	;
	if v63 != 0 {
		goto L4
	} else {
		goto L29
	}
L18:
	;
	v36 = v2
	v37 = v32
	goto L19
L19:
	;
	v42 = int32(2)
	v43 = base.I32_div_s(v37-v36, v42)
	v44 = v43 + v36
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v44<<(uint(v42)%32))+uint32(_c_F_GetNonHistoricCatalogSnapshot[2])))
	v50 = base.B2i32(v49 == l0)
	if v49 == l0 {
		v63 = v50
		goto L17
	} else {
		goto L21
	}
L20:
	;
	v63 = v50
	goto L17
L21:
	;
	v53 = base.B2i32(base.Ui32(v49) < base.Ui32(l0))
	if base.Ui32(v49) < base.Ui32(l0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v54 = v44 + int32(1)
	goto L24
L23:
	;
	v54 = v36
	goto L24
L24:
	;
	if base.Ui32(v49) < base.Ui32(l0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v57 = v37
	goto L27
L26:
	;
	v57 = v44 - int32(1)
	goto L27
L27:
	;
	if v54 <= v57 {
		v36 = v54
		v37 = v57
		goto L19
	} else {
		goto L28
	}
L28:
	;
	goto L20
L29:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_GetNonHistoricCatalogSnapshot[0]))
	if v66 == int32(0) {
		goto L2
	} else {
		goto L30
	}
L30:
	;
	F_pairingheap_remove(m, int32(_a_F_GetNonHistoricCatalogSnapshot_1), v66+int32(52))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	return int32(0)
L32:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_GetNonHistoricCatalogSnapshot[0])) = int32(0)
	v80 = *(*int32)(unsafe.Add(mBase, _c_F_GetNonHistoricCatalogSnapshot[3]))
	if v80 != 0 {
		goto L2
	} else {
		goto L33
	}
L33:
	;
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_GetNonHistoricCatalogSnapshot[4]))
	if v83 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v85 = *(*int32)(unsafe.Add(mBase, _c_F_GetNonHistoricCatalogSnapshot[5]))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+40))
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_GetNonHistoricCatalogSnapshot[4]))
	v90 = v88 - int32(48)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v91))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v86)) == int32(0) {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v107 = int32(0)
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_GetNonHistoricCatalogSnapshot[6])) = v107
	v111 = *(*int32)(unsafe.Add(mBase, _c_F_GetNonHistoricCatalogSnapshot[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v111)+40)) = v107
	goto L4
L37:
	;
	if v103 == int32(0) {
		goto L4
	} else {
		goto L41
	}
L38:
	;
	v103 = base.B2i32(base.Ui32(v86) < base.Ui32(v91))
	goto L37
L39:
	;
	goto L40
L40:
	;
	v103 = int32(base.Ui32(v86-v91) >> (uint(int32(31)) % 32))
	goto L37
L41:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	v107 = v106
	goto L36
L42:
	;
	goto L2
L43:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_GetNonHistoricCatalogSnapshot[0])) = v129
	F_pairingheap_add(m, int32(_a_F_GetNonHistoricCatalogSnapshot_1), v129+int32(52))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L31
	} else {
		goto L44
	}
L44:
	;
	v138 = *(*int32)(unsafe.Add(mBase, _c_F_GetNonHistoricCatalogSnapshot[0]))
	v139 = v138
	goto L1
}
func F_GetPrivateRefCountEntry(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = l0
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_GetPrivateRefCountEntry[0]))
	if l0 == v10 {
		v84 = int32(_a_F_GetPrivateRefCountEntry_0)
		m.G0 = v6 + int32(16)
		return v84
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_GetPrivateRefCountEntry[1]))
		if l0 == v14 {
			v84 = int32(_a_F_GetPrivateRefCountEntry_1)
			m.G0 = v6 + int32(16)
			return v84
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, _c_F_GetPrivateRefCountEntry[2]))
			if l0 == v18 {
				v84 = int32(_a_F_GetPrivateRefCountEntry_2)
				m.G0 = v6 + int32(16)
				return v84
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, _c_F_GetPrivateRefCountEntry[3]))
				if l0 == v22 {
					v84 = int32(_a_F_GetPrivateRefCountEntry_3)
					m.G0 = v6 + int32(16)
					return v84
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, _c_F_GetPrivateRefCountEntry[4]))
					if l0 == v26 {
						v84 = int32(_a_F_GetPrivateRefCountEntry_4)
						m.G0 = v6 + int32(16)
						return v84
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, _c_F_GetPrivateRefCountEntry[5]))
						if l0 == v30 {
							v84 = int32(_a_F_GetPrivateRefCountEntry_5)
							m.G0 = v6 + int32(16)
							return v84
						} else {
							v34 = *(*int32)(unsafe.Add(mBase, _c_F_GetPrivateRefCountEntry[6]))
							if l0 == v34 {
								v84 = int32(_a_F_GetPrivateRefCountEntry_6)
								m.G0 = v6 + int32(16)
								return v84
							} else {
								v38 = *(*int32)(unsafe.Add(mBase, _c_F_GetPrivateRefCountEntry[7]))
								if l0 == v38 {
									v84 = int32(_a_F_GetPrivateRefCountEntry_7)
									m.G0 = v6 + int32(16)
									return v84
								} else {
									v41 = int32(0)
									v43 = *(*int32)(unsafe.Add(mBase, _c_F_GetPrivateRefCountEntry[8]))
									if v43 == v41 {
										v84 = v41
										m.G0 = v6 + int32(16)
										return v84
									} else {
										v47 = *(*int32)(unsafe.Add(mBase, _c_F_GetPrivateRefCountEntry[9]))
										v50 = int32(0)
										v52 = F_hash_search(m, v47, v6+int32(12), v50, v50)
										mBase = m.M
										v55 = m.ExcPending
										if v55 != 0 {
											return int32(0)
										} else {
											if v52 == int32(0) {
												v84 = v52
												m.G0 = v6 + int32(16)
												return v84
											} else {
												F_ReservePrivateRefCountEntry(m)
												mBase = m.M
												v59 = m.ExcPending
												if v59 != 0 {
													return int32(0)
												} else {
													v60 = int32(_a_F_GetPrivateRefCountEntry_8)
													v61 = *(*int32)(unsafe.Add(mBase, _c_F_GetPrivateRefCountEntry[10]))
													*(*int32)(unsafe.Add(mBase, _c_F_GetPrivateRefCountEntry[10])) = int32(0)
													v65 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
													*(*int32)(unsafe.Add(mBase, uint32(v61))) = v65
													v67 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
													*(*int32)(unsafe.Add(mBase, uint32(v61)+4)) = v67
													v70 = *(*int32)(unsafe.Add(mBase, _c_F_GetPrivateRefCountEntry[9]))
													v76 = F_hash_search(m, v70, v6+int32(12), int32(2), v6+int32(11))
													mBase = m.M
													v77 = m.ExcPending
													if v77 != 0 {
														return int32(0)
													} else {
														v78 = int32(_a_F_GetPrivateRefCountEntry_9)
														v80 = *(*int32)(unsafe.Add(mBase, _c_F_GetPrivateRefCountEntry[8]))
														*(*int32)(unsafe.Add(mBase, _c_F_GetPrivateRefCountEntry[8])) = v80 - int32(1)
														v84 = v61
														m.G0 = v6 + int32(16)
														return v84
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_GetUserNameFromId(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = F_SearchSysCache1(m, int32(11), l0)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if v11 == int32(0) {
			if l1 != 0 {
				v43 = int32(0)
				m.G0 = v8 + int32(16)
				return v43
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(67137668))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
						F_errmsg(m, int32(_a_F_GetUserNameFromId_0), v8)
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_GetUserNameFromId_1), int32(1050), int32(_a_F_GetUserNameFromId_2))
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		} else {
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
			v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+22)))
			v38 = F_pstrdup(m, v33+v34+int32(4))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				F_ReleaseCatCache(m, v11)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					v43 = v38
					m.G0 = v8 + int32(16)
					return v43
				}
			}
		}
	}
}
func F_GlobalVisCheckRemovableXid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v52 int32
	_ = v52
	var v53 int64
	_ = v53
	var v57 int64
	_ = v57
	var v58 int64
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v74 int64
	_ = v74
	var v76 int32
	_ = v76
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_GlobalVisCheckRemovableXid[0]))
	v53 = *(*int64)(unsafe.Add(mBase, uint32(v52)))
	v57 = v53 + base.I64_extend_i32_s(l0-base.I32_wrap_i64(v53))
	v58 = *(*int64)(unsafe.Add(mBase, uint32(v52)+8))
	if base.Ui64(v57) < base.Ui64(v58) {
		v76 = int32(1)
		m.G0 = v9 + int32(48)
		return v76
	} else {
		v61 = int32(0)
		if base.Ui64(v53) <= base.Ui64(v57) {
			v76 = v61
			m.G0 = v9 + int32(48)
			return v76
		} else {
			v64 = *(*int32)(unsafe.Add(mBase, _c_F_GlobalVisCheckRemovableXid[1]))
			if v64 != 0 {
				v66 = *(*int32)(unsafe.Add(mBase, _c_F_GlobalVisCheckRemovableXid[2]))
				if v66 == v64 {
					v76 = v61
					m.G0 = v9 + int32(48)
					return v76
				} else {
					F_ComputeXidHorizons(m, v9+int32(8))
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int32(0)
					} else {
						v74 = *(*int64)(unsafe.Add(mBase, uint32(v52)+8))
						v76 = base.B2i32(base.Ui64(v57) < base.Ui64(v74))
						m.G0 = v9 + int32(48)
						return v76
					}
				}
			} else {
				F_ComputeXidHorizons(m, v9+int32(8))
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
					return int32(0)
				} else {
					v74 = *(*int64)(unsafe.Add(mBase, uint32(v52)+8))
					v76 = base.B2i32(base.Ui64(v57) < base.Ui64(v74))
					m.G0 = v9 + int32(48)
					return v76
				}
			}
		}
	}
}
func F_generate_dependencies_recurse(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v8-int32(1) <= l1 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v12 <= int32(0) {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v87 <= l2 {
		goto L1
	} else {
		goto L23
	}
L5:
	;
	v22 = int32(0)
	goto L6
L6:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l3+l1<<(uint(int32(1))%32)))) = uint16(v22)
	v26 = int32(0)
	if v26 < l1 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L1
L8:
	;
	v84 = v22 + int32(1)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v84 < v85 {
		v22 = v84
		goto L6
	} else {
		goto L22
	}
L9:
	;
	v31 = v26
	goto L12
L10:
	;
	goto L11
L11:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v53 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+12)))
	v54 = int32(1)
	v59 = F_repalloc(m, v51, v52*(v53+v54)<<(uint(v54)%32))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	v39 = int32(*(*int16)(unsafe.Add(mBase, uint32(l3+v31<<(uint(int32(1))%32)))))
	if v22 == v39 {
		goto L8
	} else {
		goto L14
	}
L13:
	;
	goto L11
L14:
	;
	v42 = v31 + int32(1)
	if v42 != l1 {
		v31 = v42
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	return
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v59
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v63 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+12)))
	v65 = int32(1)
	v69 = v62 << (uint(v65) % 32)
	if v69 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v72 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	v74 = v72 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v74)
	goto L8
L19:
	;
	v70 = F__emscripten_memcpy_bulkmem(m, v59+v62*v63<<(uint(v65)%32), l3, v69)
	mBase = m.M
	goto L21
L20:
	;
	goto L21
L21:
	;
	goto L18
L22:
	;
	goto L7
L23:
	;
	v89 = int32(1)
	v96 = l2
	goto L24
L24:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l3+l1<<(uint(v89)%32)))) = uint16(v96)
	v104 = base.I32_extend16_s(v96 + int32(1))
	F_generate_dependencies_recurse(m, l0, l1+v89, v104, l3)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L16
	} else {
		goto L26
	}
L25:
	;
	goto L1
L26:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v104 < v107 {
		v96 = v104
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
}
func F_generate_useful_gather_paths(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 float64
	_ = v33
	var v34 int32
	_ = v34
	var v35 float64
	_ = v35
	var v37 int32
	_ = v37
	var v43 float64
	_ = v43
	var v47 float64
	_ = v47
	var v49 float64
	_ = v49
	var v51 float64
	_ = v51
	var v52 float64
	_ = v52
	var v60 float64
	_ = v60
	var v64 float64
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 float64
	_ = v104
	var v105 int32
	_ = v105
	var v106 float64
	_ = v106
	var v108 int32
	_ = v108
	var v114 float64
	_ = v114
	var v118 float64
	_ = v118
	var v120 float64
	_ = v120
	var v122 float64
	_ = v122
	var v123 float64
	_ = v123
	var v131 float64
	_ = v131
	var v135 float64
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v174 int32
	_ = v174
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v278 int32
	_ = v278
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v303 int32
	_ = v303
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v437 int32
	_ = v437
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v455 int32
	_ = v455
	var v463 int32
	_ = v463
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v479 float64
	_ = v479
	var v480 int32
	_ = v480
	var v481 float64
	_ = v481
	var v483 int32
	_ = v483
	var v489 float64
	_ = v489
	var v493 float64
	_ = v493
	var v495 float64
	_ = v495
	var v497 float64
	_ = v497
	var v498 float64
	_ = v498
	var v506 float64
	_ = v506
	var v510 float64
	_ = v510
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	v4 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v18 == v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v16 + int32(16)
	return
L2:
	;
	v21 = m.G0
	v23 = v21 - int32(16)
	m.G0 = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v25 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v23 + int32(16)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if v163 == int32(0) {
		goto L1
	} else {
		goto L43
	}
L4:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v33 = *(*float64)(unsafe.Add(mBase, uint32(v29)+32))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v29)+24))
	v35 = base.F64_convert_i32_s(v34)
	v37 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_generate_useful_gather_paths[0])))
	if v37 == int32(1) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v23)+8)) = v64
	v66 = int32(0)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if l2 != 0 {
		goto L16
	} else {
		goto L17
	}
L6:
	;
	v43 = base.F64_add(base.F64_mul(v35, float64(-0.3)), float64(1))
	if base.F64_gt(v43, float64(0)) != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v49 = v35
	goto L8
L8:
	;
	v51 = float64(1e+100)
	v52 = base.F64_mul(v33, v49)
	if base.F64_gt(v52, v51) != 0 {
		v64 = v51
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v47 = v43
	goto L11
L10:
	;
	v47 = math.Float64frombits(uint64(0x8000000000000000))
	goto L11
L11:
	;
	v49 = base.F64_add(v47, v35)
	goto L8
L12:
	;
	goto L5
L13:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v52)&int64(9223372036854775807)) {
		v64 = v51
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v60 = float64(1)
	if base.F64_le(v52, v60) != 0 {
		v64 = v60
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v64 = base.F64_nearest(v52)
	goto L12
L16:
	;
	v71 = v23 + int32(8)
	goto L18
L17:
	;
	v71 = v66
	goto L18
L18:
	;
	v72 = F_create_gather_path(m, l0, l1, v29, v67, v71)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return
L20:
	;
	F_add_path(m, l1, v72)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v76 == int32(0) {
		goto L3
	} else {
		goto L22
	}
L22:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	if v79 <= int32(0) {
		goto L3
	} else {
		goto L23
	}
L23:
	;
	v85 = v66
	goto L24
L24:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v95+v85<<(uint(int32(2))%32))))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)+64))
	if v100 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	goto L3
L26:
	;
	v104 = *(*float64)(unsafe.Add(mBase, uint32(v99)+32))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v99)+24))
	v106 = base.F64_convert_i32_s(v105)
	v108 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_generate_useful_gather_paths[0])))
	if v108 == int32(1) {
		goto L30
	} else {
		goto L31
	}
L27:
	;
	goto L28
L28:
	;
	v144 = v85 + int32(1)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	if v144 < v145 {
		v85 = v144
		goto L24
	} else {
		goto L42
	}
L29:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v23)+8)) = v135
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v99)+64))
	v139 = F_create_gather_merge_path(m, l0, l1, v99, v137, v138, v71)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L19
	} else {
		goto L40
	}
L30:
	;
	v114 = base.F64_add(base.F64_mul(v106, float64(-0.3)), float64(1))
	if base.F64_gt(v114, float64(0)) != 0 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	v120 = v106
	goto L32
L32:
	;
	v122 = float64(1e+100)
	v123 = base.F64_mul(v104, v120)
	if base.F64_gt(v123, v122) != 0 {
		v135 = v122
		goto L36
	} else {
		goto L37
	}
L33:
	;
	v118 = v114
	goto L35
L34:
	;
	v118 = math.Float64frombits(uint64(0x8000000000000000))
	goto L35
L35:
	;
	v120 = base.F64_add(v118, v106)
	goto L32
L36:
	;
	goto L29
L37:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v123)&int64(9223372036854775807)) {
		v135 = v122
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v131 = float64(1)
	if base.F64_le(v123, v131) != 0 {
		v135 = v131
		goto L36
	} else {
		goto L39
	}
L39:
	;
	v135 = base.F64_nearest(v123)
	goto L36
L40:
	;
	F_add_path(m, l1, v139)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L19
	} else {
		goto L41
	}
L41:
	;
	goto L28
L42:
	;
	goto L25
L43:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v163)+4))
	if int32(0) < v166 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v312 = F_lappend(m, int32(0), v303)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L19
	} else {
		goto L82
	}
L45:
	;
	v174 = int32(0)
	goto L48
L46:
	;
	goto L47
L47:
	;
	if v166 != 0 {
		goto L1
	} else {
		goto L81
	}
L48:
	;
	v183 = int32(0)
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v163)+12))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v185+v174<<(uint(int32(2))%32))))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)+4))
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+41)))
	if v191 != 0 {
		v278 = v183
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if v289 != 0 {
		goto L75
	} else {
		goto L76
	}
L50:
	;
	if v278 != 0 {
		goto L71
	} else {
		goto L72
	}
L51:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)+4))
	if v193 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
	if int32(0) < v194 {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v249 = v183
	goto L54
L54:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v262 = F_find_computable_ec_member(m, l0, v190, v249, v260, int32(1))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L19
	} else {
		goto L68
	}
L55:
	;
	v200 = v183
	goto L58
L56:
	;
	goto L57
L57:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v192)+4))
	v249 = v245
	goto L54
L58:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v193)+12))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v210+v200<<(uint(int32(2))%32))))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v216 = F_find_ec_member_matching_expr(m, v190, v214, v215)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L19
	} else {
		goto L61
	}
L59:
	;
	goto L57
L60:
	;
	v229 = v200 + int32(1)
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
	if v229 < v230 {
		v200 = v229
		goto L58
	} else {
		goto L67
	}
L61:
	;
	if v216 == int32(0) {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v216)+4))
	v221 = F_expression_returns_set(m, v220)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L19
	} else {
		goto L63
	}
L63:
	;
	if v221 != 0 {
		goto L60
	} else {
		goto L64
	}
L64:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v216)+4))
	v225 = F_is_parallel_safe(m, l0, v224)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L19
	} else {
		goto L65
	}
L65:
	;
	if v225 != 0 {
		v278 = int32(1)
		goto L50
	} else {
		goto L66
	}
L66:
	;
	goto L60
L67:
	;
	goto L59
L68:
	;
	if v262 == int32(0) {
		v278 = int32(0)
		goto L50
	} else {
		goto L69
	}
L69:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v262)+4))
	v267 = F_expression_returns_set(m, v266)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L19
	} else {
		goto L70
	}
L70:
	;
	v278 = v267 ^ int32(1)
	goto L50
L71:
	;
	v285 = v174 + int32(1)
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v163)+4))
	if v285 < v286 {
		v174 = v285
		goto L48
	} else {
		goto L74
	}
L72:
	;
	v288 = v174
	goto L73
L73:
	;
	goto L49
L74:
	;
	v288 = v285
	goto L73
L75:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v289)+4))
	v292 = v290
	goto L77
L76:
	;
	v292 = int32(0)
	goto L77
L77:
	;
	if v292 == v288 {
		v303 = v289
		goto L44
	} else {
		goto L78
	}
L78:
	;
	if v288 <= int32(0) {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v296 = F_list_copy_head(m, v289, v288)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L19
	} else {
		goto L80
	}
L80:
	;
	v303 = v296
	goto L44
L81:
	;
	v303 = v163
	goto L44
L82:
	;
	if v312 == int32(0) {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v312)+4))
	if v316 <= int32(0) {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	if l2 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v322 = v16 + int32(8)
	goto L87
L86:
	;
	v322 = int32(0)
	goto L87
L87:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v323)+12))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v324)))
	v338 = v4
	goto L88
L88:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v339 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	goto L1
L90:
	;
	v539 = v338 + int32(1)
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v312)+4))
	if v539 < v540 {
		v338 = v539
		goto L88
	} else {
		goto L156
	}
L91:
	;
	v342 = int32(0)
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v339)+4))
	if v343 <= v342 {
		goto L90
	} else {
		goto L92
	}
L92:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v312)+12))
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v346+v338<<(uint(int32(2))%32))))
	v355 = v342
	goto L93
L93:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v339)+12))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v364+v355<<(uint(int32(2))%32))))
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v368)+64))
	v371 = v16 + int32(4)
	if v350 == v369 {
		goto L98
	} else {
		goto L99
	}
L94:
	;
	goto L90
L95:
	;
	v522 = v355 + int32(1)
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v339)+4))
	if v522 < v523 {
		v355 = v522
		goto L93
	} else {
		goto L155
	}
L96:
	;
	if v449 != 0 {
		goto L95
	} else {
		goto L128
	}
L97:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v350)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v371))) = v437
	v449 = int32(1)
	goto L96
L98:
	;
	if v350 != 0 {
		goto L97
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	if v350 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v371))) = int32(0)
	v449 = int32(1)
	goto L96
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v371))) = int32(0)
	v449 = int32(1)
	goto L96
L103:
	;
	goto L104
L104:
	;
	if v369 == int32(0) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v389 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v371))) = v389
	v449 = v389
	goto L96
L106:
	;
	goto L107
L107:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v369)+4))
	v393 = int32(0)
	if v393 < v392 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v396 = v392
	goto L110
L109:
	;
	v396 = v393
	goto L110
L110:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v350)+4))
	v401 = int32(0)
	goto L111
L111:
	;
	if v401 < v397 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v350)+12))
	v413 = v409 + v401<<(uint(int32(2))%32)
	goto L115
L114:
	;
	v413 = int32(0)
	goto L115
L115:
	;
	if v401 == v396 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v371))) = v396
	v449 = base.B2i32(v413 == int32(0))
	goto L96
L117:
	;
	goto L118
L118:
	;
	v419 = base.B2i32(v413 == int32(0))
	if v413 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v371))) = v401
	v449 = v419
	goto L96
L120:
	;
	goto L121
L121:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v369)+12))
	v426 = v423 + v401<<(uint(int32(2))%32)
	if v426 == int32(0) {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v371))) = v401
	v449 = v419
	goto L96
L123:
	;
	goto L124
L124:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v413)))
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v426)))
	if v430 != v431 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v371))) = v401
	v449 = int32(0)
	goto L96
L126:
	;
	v401 = v401 + int32(1)
	goto L111
L128:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v368 != v325 {
		goto L132
	} else {
		goto L133
	}
L129:
	;
	v479 = *(*float64)(unsafe.Add(mBase, uint32(v475)+32))
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v475)+24))
	v481 = base.F64_convert_i32_s(v480)
	v483 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_generate_useful_gather_paths[0])))
	if v483 == int32(1) {
		goto L143
	} else {
		goto L144
	}
L130:
	;
	v472 = F_create_incremental_sort_path(m, l0, l1, v368, v350, v450, float64(-1))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L19
	} else {
		goto L141
	}
L131:
	;
	v468 = F_create_sort_path(m, l1, v368, v350, float64(-1))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L19
	} else {
		goto L140
	}
L132:
	;
	if v450 == int32(0) {
		goto L95
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	if v450 == int32(0) {
		goto L131
	} else {
		goto L138
	}
L135:
	;
	v455 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_generate_useful_gather_paths[1])))
	if v455 == int32(0) {
		goto L95
	} else {
		goto L136
	}
L136:
	;
	if v455 == int32(0) {
		goto L131
	} else {
		goto L137
	}
L137:
	;
	goto L130
L138:
	;
	v463 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_generate_useful_gather_paths[1])))
	if v463&int32(1) != 0 {
		goto L130
	} else {
		goto L139
	}
L139:
	;
	goto L131
L140:
	;
	v475 = v468
	goto L129
L141:
	;
	v475 = v472
	goto L129
L142:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v16)+8)) = v510
	v512 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v475)+64))
	v514 = F_create_gather_merge_path(m, l0, l1, v475, v512, v513, v322)
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L19
	} else {
		goto L153
	}
L143:
	;
	v489 = base.F64_add(base.F64_mul(v481, float64(-0.3)), float64(1))
	if base.F64_gt(v489, float64(0)) != 0 {
		goto L146
	} else {
		goto L147
	}
L144:
	;
	v495 = v481
	goto L145
L145:
	;
	v497 = float64(1e+100)
	v498 = base.F64_mul(v479, v495)
	if base.F64_gt(v498, v497) != 0 {
		v510 = v497
		goto L149
	} else {
		goto L150
	}
L146:
	;
	v493 = v489
	goto L148
L147:
	;
	v493 = math.Float64frombits(uint64(0x8000000000000000))
	goto L148
L148:
	;
	v495 = base.F64_add(v493, v481)
	goto L145
L149:
	;
	goto L142
L150:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v498)&int64(9223372036854775807)) {
		v510 = v497
		goto L149
	} else {
		goto L151
	}
L151:
	;
	v506 = float64(1)
	if base.F64_le(v498, v506) != 0 {
		v510 = v506
		goto L149
	} else {
		goto L152
	}
L152:
	;
	v510 = base.F64_nearest(v498)
	goto L149
L153:
	;
	F_add_path(m, l1, v514)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L19
	} else {
		goto L154
	}
L154:
	;
	goto L95
L155:
	;
	goto L94
L156:
	;
	goto L89
}
func F_geqo_randint(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	var v8 int64
	_ = v8
	var v15 int64
	_ = v15
	var v17 int64
	_ = v17
	var v18 int64
	_ = v18
	var v21 int64
	_ = v21
	var v23 int64
	_ = v23
	var v27 int64
	_ = v27
	var v29 int64
	_ = v29
	var v37 int64
	_ = v37
	var v42 int64
	_ = v42
	var v55 int64
	_ = v55
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+368))
	v6 = v4 + int32(8)
	v7 = base.I64_extend_i32_s(l2)
	v8 = base.I64_extend_i32_s(l1)
	if base.Ui64(v8) <= base.Ui64(v7) {
		v55 = v7
	} else {
		v15 = v8 - v7
		v17 = *(*int64)(unsafe.Add(mBase, uint32(v6)+8))
		v18 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
		v21 = v18
		v23 = v17
		for {
			v27 = v21 ^ v23
			v29 = base.I64_rotl(v27, int64(37))
			v37 = v27 ^ (v27<<(uint(int64(16))%64) ^ base.I64_rotl(v21, int64(24)))
			v42 = int64(base.Ui64(base.I64_rotl(v21*int64(5), int64(7))*int64(9)) >> (uint(base.I64_clz(v15)) % 64))
			if base.Ui64(v15) < base.Ui64(v42) {
				v21 = v37
				v23 = v29
				continue
			} else {
				break
			}
			break
		}
		*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = v29
		*(*int64)(unsafe.Add(mBase, uint32(v6))) = v37
		v55 = v7 + v42
	}
	return base.I32_wrap_i64(v55)
}
func F_german_UTF_8_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v84 int32
	_ = v84
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v215 int32
	_ = v215
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
	var v253 int32
	_ = v253
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v348 int32
	_ = v348
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v374 int32
	_ = v374
	var v386 int32
	_ = v386
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v452 int32
	_ = v452
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v466 int32
	_ = v466
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v556 int32
	_ = v556
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v599 int32
	_ = v599
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v648 int32
	_ = v648
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v670 int32
	_ = v670
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v686 int32
	_ = v686
	var v699 int32
	_ = v699
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v719 int32
	_ = v719
	var v725 int32
	_ = v725
	var v732 int32
	_ = v732
	var v743 int32
	_ = v743
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v770 int32
	_ = v770
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v783 int32
	_ = v783
	var v786 int32
	_ = v786
	var v788 int32
	_ = v788
	var v792 int32
	_ = v792
	var v802 int32
	_ = v802
	var v804 int32
	_ = v804
	var v808 int32
	_ = v808
	var v821 int32
	_ = v821
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v841 int32
	_ = v841
	var v847 int32
	_ = v847
	var v855 int32
	_ = v855
	var v866 int32
	_ = v866
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v875 int32
	_ = v875
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v898 int32
	_ = v898
	var v905 int32
	_ = v905
	var v907 int32
	_ = v907
	var v911 int32
	_ = v911
	var v914 int32
	_ = v914
	var v916 int32
	_ = v916
	var v920 int32
	_ = v920
	var v930 int32
	_ = v930
	var v932 int32
	_ = v932
	var v936 int32
	_ = v936
	var v949 int32
	_ = v949
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v969 int32
	_ = v969
	var v975 int32
	_ = v975
	var v982 int32
	_ = v982
	var v993 int32
	_ = v993
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1020 int32
	_ = v1020
	var v1027 int32
	_ = v1027
	var v1029 int32
	_ = v1029
	var v1033 int32
	_ = v1033
	var v1036 int32
	_ = v1036
	var v1038 int32
	_ = v1038
	var v1042 int32
	_ = v1042
	var v1052 int32
	_ = v1052
	var v1054 int32
	_ = v1054
	var v1058 int32
	_ = v1058
	var v1071 int32
	_ = v1071
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1091 int32
	_ = v1091
	var v1097 int32
	_ = v1097
	var v1105 int32
	_ = v1105
	var v1116 int32
	_ = v1116
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1126 int32
	_ = v1126
	var v1130 int32
	_ = v1130
	var v1134 int32
	_ = v1134
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1151 int32
	_ = v1151
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1161 int32
	_ = v1161
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1167 int32
	_ = v1167
	var v1170 int32
	_ = v1170
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1179 int32
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1191 int32
	_ = v1191
	var v1193 int32
	_ = v1193
	var v1195 int32
	_ = v1195
	var v1199 int32
	_ = v1199
	var v1203 int32
	_ = v1203
	var v1206 int32
	_ = v1206
	var v1208 int32
	_ = v1208
	var v1211 int32
	_ = v1211
	var v1214 int32
	_ = v1214
	var v1217 int32
	_ = v1217
	var v1221 int32
	_ = v1221
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1241 int32
	_ = v1241
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1262 int32
	_ = v1262
	var v1263 int32
	_ = v1263
	var v1265 int32
	_ = v1265
	var v1267 int32
	_ = v1267
	var v1273 int32
	_ = v1273
	var v1275 int32
	_ = v1275
	var v1277 int32
	_ = v1277
	var v1279 int32
	_ = v1279
	var v1292 int32
	_ = v1292
	var v1294 int32
	_ = v1294
	var v1296 int32
	_ = v1296
	var v1314 int32
	_ = v1314
	var v1322 int32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1327 int32
	_ = v1327
	var v1333 int32
	_ = v1333
	var v1349 int32
	_ = v1349
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1370 int32
	_ = v1370
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1377 int32
	_ = v1377
	var v1379 int32
	_ = v1379
	var v1392 int32
	_ = v1392
	var v1393 int32
	_ = v1393
	var v1396 int32
	_ = v1396
	var v1398 int32
	_ = v1398
	var v1399 int32
	_ = v1399
	var v1403 int32
	_ = v1403
	var v1404 int32
	_ = v1404
	var v1420 int32
	_ = v1420
	var v1423 int32
	_ = v1423
	var v1424 int32
	_ = v1424
	var v1441 int32
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1444 int32
	_ = v1444
	var v1446 int32
	_ = v1446
	var v1452 int32
	_ = v1452
	var v1454 int32
	_ = v1454
	var v1456 int32
	_ = v1456
	var v1458 int32
	_ = v1458
	var v1471 int32
	_ = v1471
	var v1473 int32
	_ = v1473
	var v1475 int32
	_ = v1475
	var v1493 int32
	_ = v1493
	var v1501 int32
	_ = v1501
	var v1502 int32
	_ = v1502
	var v1506 int32
	_ = v1506
	var v1512 int32
	_ = v1512
	var v1528 int32
	_ = v1528
	var v1535 int32
	_ = v1535
	var v1536 int32
	_ = v1536
	var v1537 int32
	_ = v1537
	var v1538 int32
	_ = v1538
	var v1545 int32
	_ = v1545
	var v1547 int32
	_ = v1547
	var v1552 int32
	_ = v1552
	var v1554 int32
	_ = v1554
	var v1559 int32
	_ = v1559
	var v1564 int32
	_ = v1564
	var v1568 int32
	_ = v1568
	var v1571 int32
	_ = v1571
	var v1575 int32
	_ = v1575
	var v1589 int32
	_ = v1589
	var v1593 int32
	_ = v1593
	var v1594 int32
	_ = v1594
	var v1599 int32
	_ = v1599
	var v1603 int32
	_ = v1603
	var v1604 int32
	_ = v1604
	var v1606 int32
	_ = v1606
	var v1608 int32
	_ = v1608
	var v1621 int32
	_ = v1621
	var v1622 int32
	_ = v1622
	var v1625 int32
	_ = v1625
	var v1627 int32
	_ = v1627
	var v1628 int32
	_ = v1628
	var v1632 int32
	_ = v1632
	var v1633 int32
	_ = v1633
	var v1636 int32
	_ = v1636
	var v1638 int32
	_ = v1638
	var v1640 int32
	_ = v1640
	var v1643 int32
	_ = v1643
	var v1646 int32
	_ = v1646
	var v1649 int32
	_ = v1649
	var v1653 int32
	_ = v1653
	var v1656 int32
	_ = v1656
	var v1658 int32
	_ = v1658
	var v1660 int32
	_ = v1660
	var v1664 int32
	_ = v1664
	var v1667 int32
	_ = v1667
	var v1668 int32
	_ = v1668
	var v1670 int32
	_ = v1670
	var v1671 int32
	_ = v1671
	var v1674 int32
	_ = v1674
	var v1676 int32
	_ = v1676
	var v1680 int32
	_ = v1680
	var v1683 int32
	_ = v1683
	var v1684 int32
	_ = v1684
	var v1687 int32
	_ = v1687
	var v1688 int32
	_ = v1688
	var v1691 int32
	_ = v1691
	var v1693 int32
	_ = v1693
	var v1694 int32
	_ = v1694
	var v1696 int32
	_ = v1696
	var v1699 int32
	_ = v1699
	var v1702 int32
	_ = v1702
	var v1705 int32
	_ = v1705
	var v1709 int32
	_ = v1709
	var v1712 int32
	_ = v1712
	var v1714 int32
	_ = v1714
	var v1716 int32
	_ = v1716
	var v1718 int32
	_ = v1718
	var v1721 int32
	_ = v1721
	var v1724 int32
	_ = v1724
	var v1727 int32
	_ = v1727
	var v1731 int32
	_ = v1731
	var v1734 int32
	_ = v1734
	var v1736 int32
	_ = v1736
	var v1737 int32
	_ = v1737
	var v1739 int32
	_ = v1739
	var v1740 int32
	_ = v1740
	var v1743 int32
	_ = v1743
	var v1744 int32
	_ = v1744
	var v1747 int32
	_ = v1747
	var v1750 int32
	_ = v1750
	var v1751 int32
	_ = v1751
	var v1753 int32
	_ = v1753
	var v1755 int32
	_ = v1755
	var v1762 int32
	_ = v1762
	var v1765 int32
	_ = v1765
	var v1766 int32
	_ = v1766
	var v1769 int32
	_ = v1769
	var v1771 int32
	_ = v1771
	var v1772 int32
	_ = v1772
	var v1774 int32
	_ = v1774
	var v1778 int32
	_ = v1778
	var v1779 int32
	_ = v1779
	var v1785 int32
	_ = v1785
	var v1788 int32
	_ = v1788
	var v1795 int32
	_ = v1795
	var v1796 int32
	_ = v1796
	var v1797 int32
	_ = v1797
	var v1803 int32
	_ = v1803
	var v1804 int32
	_ = v1804
	var v1809 int32
	_ = v1809
	var v1810 int32
	_ = v1810
	var v1815 int32
	_ = v1815
	var v1816 int32
	_ = v1816
	var v1821 int32
	_ = v1821
	var v1822 int32
	_ = v1822
	var v1825 int32
	_ = v1825
	var v1826 int32
	_ = v1826
	var v1833 int32
	_ = v1833
	var v1835 int32
	_ = v1835
	var v1840 int32
	_ = v1840
	var v1842 int32
	_ = v1842
	var v1849 int32
	_ = v1849
	var v1852 int32
	_ = v1852
	var v1856 int32
	_ = v1856
	var v1863 int32
	_ = v1863
	var v1864 int32
	_ = v1864
	var v1878 int32
	_ = v1878
	var v1883 int32
	_ = v1883
	var v1888 int32
	_ = v1888
	var v1889 int32
	_ = v1889
	var v1893 int32
	_ = v1893
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v10 = v6
	goto L1
L1:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L5
L3:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v129 != 0 {
		v398 = v130
		goto L32
	} else {
		goto L33
	}
L4:
	;
	v129 = v122
	goto L3
L5:
	;
	if v24 <= v23 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v122 = int32(0)
	goto L4
L7:
	;
	v129 = int32(-1)
	goto L3
L8:
	;
	goto L9
L9:
	;
	v40 = int32(1)
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+v25))))
	if base.Ui32(v42) < base.Ui32(int32(192)) {
		v99 = v42
		v100 = v40
		goto L10
	} else {
		goto L11
	}
L10:
	;
	if int32(252) < v99 {
		v122 = v100
		goto L4
	} else {
		goto L23
	}
L11:
	;
	v46 = v23 + int32(1)
	if v46 == v24 {
		v99 = v42
		v100 = v40
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46+v25))))
	v51 = v49 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v42) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55+v25))))
	v67 = v65 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v42) {
		goto L19
	} else {
		goto L20
	}
L14:
	;
	v55 = v23 + int32(2)
	if v55 != v24 {
		goto L13
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v99 = v42<<(uint(int32(6))%32)&int32(1984) | v51
	v100 = int32(2)
	goto L10
L17:
	;
	goto L16
L18:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25+v71))))
	v99 = v84&int32(63) | (v42<<(uint(int32(18))%32)&int32(_a_F_german_UTF_8_stem_0) | v51<<(uint(int32(12))%32) | v67<<(uint(int32(6))%32))
	v100 = int32(4)
	goto L10
L19:
	;
	v71 = v23 + int32(3)
	if v71 != v24 {
		goto L18
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v99 = v42<<(uint(int32(12))%32)&int32(_a_F_german_UTF_8_stem_1) | v51<<(uint(int32(6))%32) | v67
	v100 = int32(3)
	goto L10
L22:
	;
	goto L21
L23:
	;
	v104 = v99 - int32(97)
	if v104 < int32(0) {
		v122 = v100
		goto L4
	} else {
		goto L24
	}
L24:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v104)>>(uint(int32(3))%32)))+uint32(_c_F_german_UTF_8_stem[0]))))
	if int32(base.Ui32(v110)>>(uint(v104&int32(7))%32))&int32(1) == int32(0) {
		v122 = v100
		goto L4
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v100 + v23
	goto L26
L26:
	;
	goto L6
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v10
	goto L1
L28:
	;
	return v1893
L29:
	;
	v1888 = F_slice_from_s(m, l0, int32(1), int32(_a_F_german_UTF_8_stem_2))
	mBase = m.M
	v1889 = m.ExcPending
	if v1889 != 0 {
		goto L112
	} else {
		goto L499
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6
	v466 = v6
	goto L115
L31:
	;
	v458 = F_slice_from_s(m, l0, int32(1), int32(_a_F_german_UTF_8_stem_3))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L112
	} else {
		goto L113
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v10
	v400 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L93
L33:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v131
	if v131 == v130 {
		v264 = v130
		goto L34
	} else {
		goto L35
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v131
	if v131 == v264 {
		goto L62
	} else {
		goto L63
	}
L35:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134+v131))))
	if v136 != int32(117) {
		v264 = v130
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v140 = v131 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v140
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L39
L37:
	;
	if v260 == int32(0) {
		goto L31
	} else {
		goto L61
	}
L38:
	;
	v260 = v253
	goto L37
L39:
	;
	if v155 <= v140 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v253 = int32(0)
	goto L38
L41:
	;
	v260 = int32(-1)
	goto L37
L42:
	;
	goto L43
L43:
	;
	v171 = int32(1)
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140+v156))))
	if base.Ui32(v173) < base.Ui32(int32(192)) {
		v230 = v173
		v231 = v171
		goto L44
	} else {
		goto L45
	}
L44:
	;
	if int32(252) < v230 {
		v253 = v231
		goto L38
	} else {
		goto L57
	}
L45:
	;
	v177 = v131 + int32(2)
	if v177 == v155 {
		v230 = v173
		v231 = v171
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177+v156))))
	v182 = v180 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v173) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186+v156))))
	v198 = v196 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v173) {
		goto L53
	} else {
		goto L54
	}
L48:
	;
	v186 = v131 + int32(3)
	if v186 != v155 {
		goto L47
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v230 = v173<<(uint(int32(6))%32)&int32(1984) | v182
	v231 = int32(2)
	goto L44
L51:
	;
	goto L50
L52:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156+v202))))
	v230 = v215&int32(63) | (v173<<(uint(int32(18))%32)&int32(_a_F_german_UTF_8_stem_0) | v182<<(uint(int32(12))%32) | v198<<(uint(int32(6))%32))
	v231 = int32(4)
	goto L44
L53:
	;
	v202 = v131 + int32(4)
	if v202 != v155 {
		goto L52
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v230 = v173<<(uint(int32(12))%32)&int32(_a_F_german_UTF_8_stem_1) | v182<<(uint(int32(6))%32) | v198
	v231 = int32(3)
	goto L44
L56:
	;
	goto L55
L57:
	;
	v235 = v230 - int32(97)
	if v235 < int32(0) {
		v253 = v231
		goto L38
	} else {
		goto L58
	}
L58:
	;
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v235)>>(uint(int32(3))%32)))+uint32(_c_F_german_UTF_8_stem[0]))))
	if int32(base.Ui32(v241)>>(uint(v235&int32(7))%32))&int32(1) == int32(0) {
		v253 = v231
		goto L38
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v231 + v140
	goto L60
L60:
	;
	goto L40
L61:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v264 = v263
	goto L34
L62:
	;
	v398 = v131
	goto L32
L63:
	;
	goto L64
L64:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v267+v131))))
	if v269 != int32(121) {
		v398 = v264
		goto L32
	} else {
		goto L65
	}
L65:
	;
	v273 = v131 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v273
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v273
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L68
L66:
	;
	if v393 == int32(0) {
		goto L29
	} else {
		goto L90
	}
L67:
	;
	v393 = v386
	goto L66
L68:
	;
	if v288 <= v273 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v386 = int32(0)
	goto L67
L70:
	;
	v393 = int32(-1)
	goto L66
L71:
	;
	goto L72
L72:
	;
	v304 = int32(1)
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273+v289))))
	if base.Ui32(v306) < base.Ui32(int32(192)) {
		v363 = v306
		v364 = v304
		goto L73
	} else {
		goto L74
	}
L73:
	;
	if int32(252) < v363 {
		v386 = v364
		goto L67
	} else {
		goto L86
	}
L74:
	;
	v310 = v131 + int32(2)
	if v310 == v288 {
		v363 = v306
		v364 = v304
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310+v289))))
	v315 = v313 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v306) {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v319+v289))))
	v331 = v329 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v306) {
		goto L82
	} else {
		goto L83
	}
L77:
	;
	v319 = v131 + int32(3)
	if v319 != v288 {
		goto L76
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	v363 = v306<<(uint(int32(6))%32)&int32(1984) | v315
	v364 = int32(2)
	goto L73
L80:
	;
	goto L79
L81:
	;
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v289+v335))))
	v363 = v348&int32(63) | (v306<<(uint(int32(18))%32)&int32(_a_F_german_UTF_8_stem_0) | v315<<(uint(int32(12))%32) | v331<<(uint(int32(6))%32))
	v364 = int32(4)
	goto L73
L82:
	;
	v335 = v131 + int32(4)
	if v335 != v288 {
		goto L81
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v363 = v306<<(uint(int32(12))%32)&int32(_a_F_german_UTF_8_stem_1) | v315<<(uint(int32(6))%32) | v331
	v364 = int32(3)
	goto L73
L85:
	;
	goto L84
L86:
	;
	v368 = v363 - int32(97)
	if v368 < int32(0) {
		v386 = v364
		goto L67
	} else {
		goto L87
	}
L87:
	;
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v368)>>(uint(int32(3))%32)))+uint32(_c_F_german_UTF_8_stem[0]))))
	if int32(base.Ui32(v374)>>(uint(v368&int32(7))%32))&int32(1) == int32(0) {
		v386 = v364
		goto L67
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v364 + v273
	goto L89
L89:
	;
	goto L69
L90:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v398 = v396
	goto L32
L91:
	;
	if v452 < int32(0) {
		goto L30
	} else {
		goto L111
	}
L93:
	;
	goto L94
L94:
	;
	goto L95
L95:
	;
	v407 = v10
	v409 = int32(1)
	goto L98
L97:
	;
	v452 = v437
	goto L91
L98:
	;
	if v398 <= v407 {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	goto L97
L100:
	;
	v452 = int32(-1)
	goto L91
L101:
	;
	goto L102
L102:
	;
	v414 = v407 + int32(1)
	v416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v400+v407))))
	if base.Ui32(v416) < base.Ui32(int32(192)) {
		v437 = v414
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v438 = int32(1)
	if v438 < v409 {
		v407 = v437
		v409 = v409 - v438
		goto L98
	} else {
		goto L110
	}
L104:
	;
	if v398 <= v414 {
		v437 = v414
		goto L103
	} else {
		goto L105
	}
L105:
	;
	v423 = v414
	goto L106
L106:
	;
	v426 = int32(*(*int8)(unsafe.Add(mBase, uint32(v400+v423))))
	if int32(-65) < v426 {
		v437 = v423
		goto L103
	} else {
		goto L108
	}
L107:
	;
	v437 = v398
	goto L103
L108:
	;
	v430 = v423 + int32(1)
	if v430 != v398 {
		v423 = v430
		goto L106
	} else {
		goto L109
	}
L109:
	;
	goto L107
L110:
	;
	goto L99
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v452
	v10 = v452
	goto L1
L112:
	;
	return int32(0)
L113:
	;
	if v458 < int32(0) {
		v1893 = v458
		goto L28
	} else {
		goto L114
	}
L114:
	;
	goto L27
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v466
	v473 = F_find_among(m, l0, int32(_a_F_german_UTF_8_stem_4), int32(6))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L112
	} else {
		goto L118
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6
	v563 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v564 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v563)+4)) = v564
	*(*int32)(unsafe.Add(mBase, uint32(v563)+8)) = v564
	v567 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v568 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v569 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L157
L117:
	;
	goto L116
L118:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v475
	switch v473 - int32(1) {
	case 0:
		goto L124
	case 1:
		goto L123
	case 2:
		goto L122
	case 3:
		goto L121
	case 4:
		goto L120
	default:
		goto L119
	}
L119:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v466 = v561
	goto L115
L120:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v504 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L135
L121:
	;
	v499 = F_slice_from_s(m, l0, int32(2), int32(_a_F_german_UTF_8_stem_5))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L112
	} else {
		goto L131
	}
L122:
	;
	v493 = F_slice_from_s(m, l0, int32(2), int32(_a_F_german_UTF_8_stem_6))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L112
	} else {
		goto L129
	}
L123:
	;
	v487 = F_slice_from_s(m, l0, int32(2), int32(_a_F_german_UTF_8_stem_7))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L112
	} else {
		goto L127
	}
L124:
	;
	v481 = F_slice_from_s(m, l0, int32(2), int32(_a_F_german_UTF_8_stem_8))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L112
	} else {
		goto L125
	}
L125:
	;
	if int32(0) <= v481 {
		goto L119
	} else {
		goto L126
	}
L126:
	;
	v1893 = v481
	goto L28
L127:
	;
	if int32(0) <= v487 {
		goto L119
	} else {
		goto L128
	}
L128:
	;
	v1893 = v487
	goto L28
L129:
	;
	if int32(0) <= v493 {
		goto L119
	} else {
		goto L130
	}
L130:
	;
	v1893 = v493
	goto L28
L131:
	;
	if int32(0) <= v499 {
		goto L119
	} else {
		goto L132
	}
L132:
	;
	v1893 = v499
	goto L28
L133:
	;
	if v556 < int32(0) {
		goto L117
	} else {
		goto L153
	}
L135:
	;
	goto L136
L136:
	;
	goto L137
L137:
	;
	v511 = v475
	v513 = int32(1)
	goto L140
L139:
	;
	v556 = v541
	goto L133
L140:
	;
	if v504 <= v511 {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	goto L139
L142:
	;
	v556 = int32(-1)
	goto L133
L143:
	;
	goto L144
L144:
	;
	v518 = v511 + int32(1)
	v520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v503+v511))))
	if base.Ui32(v520) < base.Ui32(int32(192)) {
		v541 = v518
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v542 = int32(1)
	if v542 < v513 {
		v511 = v541
		v513 = v513 - v542
		goto L140
	} else {
		goto L152
	}
L146:
	;
	if v504 <= v518 {
		v541 = v518
		goto L145
	} else {
		goto L147
	}
L147:
	;
	v527 = v518
	goto L148
L148:
	;
	v530 = int32(*(*int8)(unsafe.Add(mBase, uint32(v503+v527))))
	if int32(-65) < v530 {
		v541 = v527
		goto L145
	} else {
		goto L150
	}
L149:
	;
	v541 = v504
	goto L145
L150:
	;
	v534 = v527 + int32(1)
	if v534 != v504 {
		v527 = v534
		goto L148
	} else {
		goto L151
	}
L151:
	;
	goto L149
L152:
	;
	goto L141
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v556
	goto L119
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v6
	v1126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1126
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1126
	if v1126 <= v6 {
		goto L281
	} else {
		goto L282
	}
L155:
	;
	if v621 < int32(0) {
		goto L154
	} else {
		goto L175
	}
L157:
	;
	goto L158
L158:
	;
	goto L159
L159:
	;
	v576 = v568
	v578 = int32(3)
	goto L162
L161:
	;
	v621 = v606
	goto L155
L162:
	;
	if v569 <= v576 {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	goto L161
L164:
	;
	v621 = int32(-1)
	goto L155
L165:
	;
	goto L166
L166:
	;
	v583 = v576 + int32(1)
	v585 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v567+v576))))
	if base.Ui32(v585) < base.Ui32(int32(192)) {
		v606 = v583
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v607 = int32(1)
	if v607 < v578 {
		v576 = v606
		v578 = v578 - v607
		goto L162
	} else {
		goto L174
	}
L168:
	;
	if v569 <= v583 {
		v606 = v583
		goto L167
	} else {
		goto L169
	}
L169:
	;
	v592 = v583
	goto L170
L170:
	;
	v595 = int32(*(*int8)(unsafe.Add(mBase, uint32(v567+v592))))
	if int32(-65) < v595 {
		v606 = v592
		goto L167
	} else {
		goto L172
	}
L171:
	;
	v606 = v569
	goto L167
L172:
	;
	v599 = v592 + int32(1)
	if v599 != v569 {
		v592 = v599
		goto L170
	} else {
		goto L173
	}
L173:
	;
	goto L171
L174:
	;
	goto L163
L175:
	;
	v624 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v624))) = v621
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v568
	v639 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v640 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v648 = v568
	goto L178
L176:
	;
	if v743 < int32(0) {
		goto L154
	} else {
		goto L201
	}
L177:
	;
	v743 = v715
	goto L176
L178:
	;
	if v639 <= v648 {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v743 = int32(-1)
	goto L176
L181:
	;
	goto L182
L182:
	;
	v655 = int32(1)
	v657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v648+v640))))
	if base.Ui32(v657) < base.Ui32(int32(192)) {
		v714 = v657
		v715 = v655
		goto L183
	} else {
		goto L184
	}
L183:
	;
	if int32(252) < v714 {
		goto L196
	} else {
		goto L197
	}
L184:
	;
	v661 = v648 + int32(1)
	if v661 == v639 {
		v714 = v657
		v715 = v655
		goto L183
	} else {
		goto L185
	}
L185:
	;
	v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v661+v640))))
	v666 = v664 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v657) {
		goto L187
	} else {
		goto L188
	}
L186:
	;
	v680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v670+v640))))
	v682 = v680 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v657) {
		goto L192
	} else {
		goto L193
	}
L187:
	;
	v670 = v648 + int32(2)
	if v670 != v639 {
		goto L186
	} else {
		goto L190
	}
L188:
	;
	goto L189
L189:
	;
	v714 = v657<<(uint(int32(6))%32)&int32(1984) | v666
	v715 = int32(2)
	goto L183
L190:
	;
	goto L189
L191:
	;
	v699 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v640+v686))))
	v714 = v699&int32(63) | (v657<<(uint(int32(18))%32)&int32(_a_F_german_UTF_8_stem_0) | v666<<(uint(int32(12))%32) | v682<<(uint(int32(6))%32))
	v715 = int32(4)
	goto L183
L192:
	;
	v686 = v648 + int32(3)
	if v686 != v639 {
		goto L191
	} else {
		goto L195
	}
L193:
	;
	goto L194
L194:
	;
	v714 = v657<<(uint(int32(12))%32)&int32(_a_F_german_UTF_8_stem_1) | v666<<(uint(int32(6))%32) | v682
	v715 = int32(3)
	goto L183
L195:
	;
	goto L194
L196:
	;
	v732 = v715 + v648
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v732
	v648 = v732
	goto L178
L197:
	;
	v719 = v714 - int32(97)
	if v719 < int32(0) {
		goto L196
	} else {
		goto L198
	}
L198:
	;
	v725 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v719)>>(uint(int32(3))%32)))+uint32(_c_F_german_UTF_8_stem[0]))))
	if int32(base.Ui32(v725)>>(uint(v719&int32(7))%32))&int32(1) != 0 {
		goto L177
	} else {
		goto L199
	}
L199:
	;
	goto L196
L201:
	;
	v746 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v747 = v746 + v743
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v747
	v761 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v762 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v770 = v747
	goto L204
L202:
	;
	if v866 < int32(0) {
		goto L154
	} else {
		goto L226
	}
L203:
	;
	v866 = v837
	goto L202
L204:
	;
	if v761 <= v770 {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	v866 = int32(-1)
	goto L202
L207:
	;
	goto L208
L208:
	;
	v777 = int32(1)
	v779 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v770+v762))))
	if base.Ui32(v779) < base.Ui32(int32(192)) {
		v836 = v779
		v837 = v777
		goto L209
	} else {
		goto L210
	}
L209:
	;
	if int32(252) < v836 {
		goto L203
	} else {
		goto L222
	}
L210:
	;
	v783 = v770 + int32(1)
	if v783 == v761 {
		v836 = v779
		v837 = v777
		goto L209
	} else {
		goto L211
	}
L211:
	;
	v786 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v783+v762))))
	v788 = v786 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v779) {
		goto L213
	} else {
		goto L214
	}
L212:
	;
	v802 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v792+v762))))
	v804 = v802 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v779) {
		goto L218
	} else {
		goto L219
	}
L213:
	;
	v792 = v770 + int32(2)
	if v792 != v761 {
		goto L212
	} else {
		goto L216
	}
L214:
	;
	goto L215
L215:
	;
	v836 = v779<<(uint(int32(6))%32)&int32(1984) | v788
	v837 = int32(2)
	goto L209
L216:
	;
	goto L215
L217:
	;
	v821 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v762+v808))))
	v836 = v821&int32(63) | (v779<<(uint(int32(18))%32)&int32(_a_F_german_UTF_8_stem_0) | v788<<(uint(int32(12))%32) | v804<<(uint(int32(6))%32))
	v837 = int32(4)
	goto L209
L218:
	;
	v808 = v770 + int32(3)
	if v808 != v761 {
		goto L217
	} else {
		goto L221
	}
L219:
	;
	goto L220
L220:
	;
	v836 = v779<<(uint(int32(12))%32)&int32(_a_F_german_UTF_8_stem_1) | v788<<(uint(int32(6))%32) | v804
	v837 = int32(3)
	goto L209
L221:
	;
	goto L220
L222:
	;
	v841 = v836 - int32(97)
	if v841 < int32(0) {
		goto L203
	} else {
		goto L223
	}
L223:
	;
	v847 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v841)>>(uint(int32(3))%32)))+uint32(_c_F_german_UTF_8_stem[0]))))
	if int32(base.Ui32(v847)>>(uint(v841&int32(7))%32))&int32(1) == int32(0) {
		goto L203
	} else {
		goto L224
	}
L224:
	;
	v855 = v837 + v770
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v855
	v770 = v855
	goto L204
L226:
	;
	v869 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v870 = v869 + v866
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v870
	v872 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v872)))
	if v873 < v870 {
		goto L227
	} else {
		goto L228
	}
L227:
	;
	v875 = v870
	goto L229
L228:
	;
	v875 = v873
	goto L229
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v872)+8)) = v875
	v888 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v889 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v890 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v898 = v888
	goto L232
L230:
	;
	if v993 < int32(0) {
		goto L154
	} else {
		goto L255
	}
L231:
	;
	v993 = v965
	goto L230
L232:
	;
	if v889 <= v898 {
		goto L234
	} else {
		goto L235
	}
L234:
	;
	v993 = int32(-1)
	goto L230
L235:
	;
	goto L236
L236:
	;
	v905 = int32(1)
	v907 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v898+v890))))
	if base.Ui32(v907) < base.Ui32(int32(192)) {
		v964 = v907
		v965 = v905
		goto L237
	} else {
		goto L238
	}
L237:
	;
	if int32(252) < v964 {
		goto L250
	} else {
		goto L251
	}
L238:
	;
	v911 = v898 + int32(1)
	if v911 == v889 {
		v964 = v907
		v965 = v905
		goto L237
	} else {
		goto L239
	}
L239:
	;
	v914 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v911+v890))))
	v916 = v914 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v907) {
		goto L241
	} else {
		goto L242
	}
L240:
	;
	v930 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v920+v890))))
	v932 = v930 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v907) {
		goto L246
	} else {
		goto L247
	}
L241:
	;
	v920 = v898 + int32(2)
	if v920 != v889 {
		goto L240
	} else {
		goto L244
	}
L242:
	;
	goto L243
L243:
	;
	v964 = v907<<(uint(int32(6))%32)&int32(1984) | v916
	v965 = int32(2)
	goto L237
L244:
	;
	goto L243
L245:
	;
	v949 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v890+v936))))
	v964 = v949&int32(63) | (v907<<(uint(int32(18))%32)&int32(_a_F_german_UTF_8_stem_0) | v916<<(uint(int32(12))%32) | v932<<(uint(int32(6))%32))
	v965 = int32(4)
	goto L237
L246:
	;
	v936 = v898 + int32(3)
	if v936 != v889 {
		goto L245
	} else {
		goto L249
	}
L247:
	;
	goto L248
L248:
	;
	v964 = v907<<(uint(int32(12))%32)&int32(_a_F_german_UTF_8_stem_1) | v916<<(uint(int32(6))%32) | v932
	v965 = int32(3)
	goto L237
L249:
	;
	goto L248
L250:
	;
	v982 = v965 + v898
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v982
	v898 = v982
	goto L232
L251:
	;
	v969 = v964 - int32(97)
	if v969 < int32(0) {
		goto L250
	} else {
		goto L252
	}
L252:
	;
	v975 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v969)>>(uint(int32(3))%32)))+uint32(_c_F_german_UTF_8_stem[0]))))
	if int32(base.Ui32(v975)>>(uint(v969&int32(7))%32))&int32(1) != 0 {
		goto L231
	} else {
		goto L253
	}
L253:
	;
	goto L250
L255:
	;
	v996 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v997 = v996 + v993
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v997
	v1011 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1012 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1020 = v997
	goto L258
L256:
	;
	if v1116 < int32(0) {
		goto L154
	} else {
		goto L280
	}
L257:
	;
	v1116 = v1087
	goto L256
L258:
	;
	if v1011 <= v1020 {
		goto L260
	} else {
		goto L261
	}
L260:
	;
	v1116 = int32(-1)
	goto L256
L261:
	;
	goto L262
L262:
	;
	v1027 = int32(1)
	v1029 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1020+v1012))))
	if base.Ui32(v1029) < base.Ui32(int32(192)) {
		v1086 = v1029
		v1087 = v1027
		goto L263
	} else {
		goto L264
	}
L263:
	;
	if int32(252) < v1086 {
		goto L257
	} else {
		goto L276
	}
L264:
	;
	v1033 = v1020 + int32(1)
	if v1033 == v1011 {
		v1086 = v1029
		v1087 = v1027
		goto L263
	} else {
		goto L265
	}
L265:
	;
	v1036 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1033+v1012))))
	v1038 = v1036 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1029) {
		goto L267
	} else {
		goto L268
	}
L266:
	;
	v1052 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1042+v1012))))
	v1054 = v1052 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1029) {
		goto L272
	} else {
		goto L273
	}
L267:
	;
	v1042 = v1020 + int32(2)
	if v1042 != v1011 {
		goto L266
	} else {
		goto L270
	}
L268:
	;
	goto L269
L269:
	;
	v1086 = v1029<<(uint(int32(6))%32)&int32(1984) | v1038
	v1087 = int32(2)
	goto L263
L270:
	;
	goto L269
L271:
	;
	v1071 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1012+v1058))))
	v1086 = v1071&int32(63) | (v1029<<(uint(int32(18))%32)&int32(_a_F_german_UTF_8_stem_0) | v1038<<(uint(int32(12))%32) | v1054<<(uint(int32(6))%32))
	v1087 = int32(4)
	goto L263
L272:
	;
	v1058 = v1020 + int32(3)
	if v1058 != v1011 {
		goto L271
	} else {
		goto L275
	}
L273:
	;
	goto L274
L274:
	;
	v1086 = v1029<<(uint(int32(12))%32)&int32(_a_F_german_UTF_8_stem_1) | v1038<<(uint(int32(6))%32) | v1054
	v1087 = int32(3)
	goto L263
L275:
	;
	goto L274
L276:
	;
	v1091 = v1086 - int32(97)
	if v1091 < int32(0) {
		goto L257
	} else {
		goto L277
	}
L277:
	;
	v1097 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1091)>>(uint(int32(3))%32)))+uint32(_c_F_german_UTF_8_stem[0]))))
	if int32(base.Ui32(v1097)>>(uint(v1091&int32(7))%32))&int32(1) == int32(0) {
		goto L257
	} else {
		goto L278
	}
L278:
	;
	v1105 = v1087 + v1020
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1105
	v1020 = v1105
	goto L258
L280:
	;
	v1119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1119)+4)) = v1120 + v1116
	goto L154
L281:
	;
	v1370 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1370
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1370
	v1374 = v1370 - int32(1)
	v1375 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1374 <= v1375 {
		goto L342
	} else {
		goto L343
	}
L282:
	;
	v1130 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1130+v1126-int32(1)))))
	if v1134&int32(224) != int32(96) {
		goto L281
	} else {
		goto L283
	}
L283:
	;
	if int32(1)<<(uint(v1134)%32)&int32(_a_F_german_UTF_8_stem_9) == int32(0) {
		goto L281
	} else {
		goto L284
	}
L284:
	;
	v1147 = F_find_among_b(m, l0, int32(_a_F_german_UTF_8_stem_10), int32(11))
	mBase = m.M
	v1148 = m.ExcPending
	if v1148 != 0 {
		goto L112
	} else {
		goto L285
	}
L285:
	;
	if v1147 == int32(0) {
		goto L281
	} else {
		goto L286
	}
L286:
	;
	v1151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1151
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1154 = *(*int32)(unsafe.Add(mBase, uint32(v1153)+8))
	if v1151 < v1154 {
		goto L281
	} else {
		goto L287
	}
L287:
	;
	switch v1147 - int32(1) {
	case 0:
		goto L292
	case 1:
		goto L291
	case 2:
		goto L290
	case 3:
		goto L289
	case 4:
		goto L288
	default:
		goto L281
	}
L288:
	;
	v1363 = F_slice_from_s(m, l0, int32(1), int32(_a_F_german_UTF_8_stem_11))
	mBase = m.M
	v1364 = m.ExcPending
	if v1364 != 0 {
		goto L112
	} else {
		goto L340
	}
L289:
	;
	v1241 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L315
L290:
	;
	v1187 = F_slice_del(m, l0)
	mBase = m.M
	v1188 = m.ExcPending
	if v1188 != 0 {
		goto L112
	} else {
		goto L302
	}
L291:
	;
	v1183 = F_slice_del(m, l0)
	mBase = m.M
	v1184 = m.ExcPending
	if v1184 != 0 {
		goto L112
	} else {
		goto L300
	}
L292:
	;
	v1158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1159 = int32(4)
	v1161 = int32(0)
	v1163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1163-v1164 < v1159 {
		v1174 = v1161
		goto L294
	} else {
		goto L295
	}
L293:
	;
	if v1174 != 0 {
		goto L281
	} else {
		goto L297
	}
L294:
	;
	goto L293
L295:
	;
	v1167 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1170 = F_memcmp(m, v1167+v1163-v1159, int32(_a_F_german_UTF_8_stem_12), v1159)
	mBase = m.M
	if v1170 != 0 {
		v1174 = v1161
		goto L294
	} else {
		goto L296
	}
L296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1163 - v1159
	v1174 = int32(1)
	goto L294
L297:
	;
	v1175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1175 + (v1151 - v1158)
	v1179 = F_slice_del(m, l0)
	mBase = m.M
	v1180 = m.ExcPending
	if v1180 != 0 {
		goto L112
	} else {
		goto L298
	}
L298:
	;
	if int32(0) <= v1179 {
		goto L281
	} else {
		goto L299
	}
L299:
	;
	v1893 = v1179
	goto L28
L300:
	;
	if int32(0) <= v1183 {
		goto L281
	} else {
		goto L301
	}
L301:
	;
	v1893 = v1183
	goto L28
L302:
	;
	if v1187 < int32(0) {
		v1893 = v1187
		goto L28
	} else {
		goto L303
	}
L303:
	;
	v1191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1191
	v1193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1191 <= v1193 {
		goto L281
	} else {
		goto L304
	}
L304:
	;
	v1195 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1195+v1191-int32(1)))))
	if v1199 != int32(115) {
		goto L281
	} else {
		goto L305
	}
L305:
	;
	v1203 = v1191 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1203
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1203
	v1206 = int32(3)
	v1208 = int32(0)
	v1211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1203-v1211 < v1206 {
		v1221 = v1208
		goto L307
	} else {
		goto L308
	}
L306:
	;
	if v1221 == int32(0) {
		goto L281
	} else {
		goto L310
	}
L307:
	;
	goto L306
L308:
	;
	v1214 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1217 = F_memcmp(m, v1214+v1203-v1206, int32(_a_F_german_UTF_8_stem_13), v1206)
	mBase = m.M
	if v1217 != 0 {
		v1221 = v1208
		goto L307
	} else {
		goto L309
	}
L309:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1203 - v1206
	v1221 = int32(1)
	goto L307
L310:
	;
	v1224 = F_slice_del(m, l0)
	mBase = m.M
	v1225 = m.ExcPending
	if v1225 != 0 {
		goto L112
	} else {
		goto L311
	}
L311:
	;
	if int32(0) <= v1224 {
		goto L281
	} else {
		goto L312
	}
L312:
	;
	v1893 = v1224
	goto L28
L313:
	;
	if v1356 != 0 {
		goto L281
	} else {
		goto L337
	}
L314:
	;
	v1356 = v1349
	goto L313
L315:
	;
	if v1244 <= v1245 {
		v1349 = int32(-1)
		goto L314
	} else {
		goto L317
	}
L316:
	;
	v1349 = int32(0)
	goto L314
L317:
	;
	v1262 = int32(1)
	v1263 = v1244 - v1262
	v1265 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1241+v1263))))
	v1267 = v1265 & int32(255)
	if v1263 == v1245 {
		v1322 = v1267
		v1323 = v1262
		goto L318
	} else {
		goto L319
	}
L318:
	;
	if int32(116) < v1322 {
		goto L327
	} else {
		goto L328
	}
L319:
	;
	if int32(0) <= v1265 {
		v1322 = v1267
		v1323 = v1262
		goto L318
	} else {
		goto L320
	}
L320:
	;
	v1273 = v1267 & int32(63)
	v1275 = v1244 - int32(2)
	v1277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1241+v1275))))
	v1279 = v1277 << (uint(int32(6)) % 32)
	if base.B2i32(v1275 != v1245)&base.B2i32(base.Ui32(v1277) < base.Ui32(int32(192))) == int32(0) {
		goto L321
	} else {
		goto L322
	}
L321:
	;
	v1322 = v1279&int32(1984) | v1273
	v1323 = int32(2)
	goto L318
L322:
	;
	goto L323
L323:
	;
	v1292 = v1279&int32(4032) | v1273
	v1294 = v1244 - int32(3)
	v1296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1241+v1294))))
	if base.B2i32(v1294 != v1245)&base.B2i32(base.Ui32(v1296) < base.Ui32(int32(224))) == int32(0) {
		goto L324
	} else {
		goto L325
	}
L324:
	;
	v1322 = v1296<<(uint(int32(12))%32)&int32(_a_F_german_UTF_8_stem_1) | v1292
	v1323 = int32(3)
	goto L318
L325:
	;
	goto L326
L326:
	;
	v1314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1244+(v1241-int32(4))))))
	v1322 = v1296<<(uint(int32(12))%32)&int32(_a_F_german_UTF_8_stem_14) | v1314&int32(7)<<(uint(int32(18))%32) | v1292
	v1323 = int32(4)
	goto L318
L327:
	;
	v1356 = v1323
	goto L313
L328:
	;
	goto L329
L329:
	;
	v1327 = v1322 - int32(98)
	if v1327 < int32(0) {
		goto L330
	} else {
		goto L331
	}
L330:
	;
	v1356 = v1323
	goto L313
L331:
	;
	goto L332
L332:
	;
	v1333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1327)>>(uint(int32(3))%32)))+uint32(_c_F_german_UTF_8_stem[1]))))
	if int32(base.Ui32(v1333)>>(uint(v1327&int32(7))%32))&int32(1) == int32(0) {
		goto L333
	} else {
		goto L334
	}
L333:
	;
	v1356 = v1323
	goto L313
L334:
	;
	goto L335
L335:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1244 - v1323
	goto L336
L336:
	;
	goto L316
L337:
	;
	v1357 = F_slice_del(m, l0)
	mBase = m.M
	v1358 = m.ExcPending
	if v1358 != 0 {
		goto L112
	} else {
		goto L338
	}
L338:
	;
	if int32(0) <= v1357 {
		goto L281
	} else {
		goto L339
	}
L339:
	;
	v1893 = v1357
	goto L28
L340:
	;
	if v1363 < int32(0) {
		v1893 = v1363
		goto L28
	} else {
		goto L341
	}
L341:
	;
	goto L281
L342:
	;
	v1599 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1599
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1599
	v1603 = v1599 - int32(1)
	v1604 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1603 <= v1604 {
		goto L401
	} else {
		goto L402
	}
L343:
	;
	v1377 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1377+v1374))))
	if v1379&int32(224) != int32(96) {
		goto L342
	} else {
		goto L344
	}
L344:
	;
	if int32(1)<<(uint(v1379)%32)&int32(_a_F_german_UTF_8_stem_15) == int32(0) {
		goto L342
	} else {
		goto L345
	}
L345:
	;
	v1392 = F_find_among_b(m, l0, int32(_a_F_german_UTF_8_stem_16), int32(4))
	mBase = m.M
	v1393 = m.ExcPending
	if v1393 != 0 {
		goto L112
	} else {
		goto L346
	}
L346:
	;
	if v1392 == int32(0) {
		goto L342
	} else {
		goto L347
	}
L347:
	;
	v1396 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1396
	v1398 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1399 = *(*int32)(unsafe.Add(mBase, uint32(v1398)+8))
	if v1396 < v1399 {
		goto L342
	} else {
		goto L348
	}
L348:
	;
	switch v1392 - int32(1) {
	case 0:
		goto L350
	case 1:
		goto L349
	default:
		goto L342
	}
L349:
	;
	v1420 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1423 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1424 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L355
L350:
	;
	v1403 = F_slice_del(m, l0)
	mBase = m.M
	v1404 = m.ExcPending
	if v1404 != 0 {
		goto L112
	} else {
		goto L351
	}
L351:
	;
	if int32(0) <= v1403 {
		goto L342
	} else {
		goto L352
	}
L352:
	;
	v1893 = v1403
	goto L28
L353:
	;
	if v1535 != 0 {
		goto L342
	} else {
		goto L377
	}
L354:
	;
	v1535 = v1528
	goto L353
L355:
	;
	if v1423 <= v1424 {
		v1528 = int32(-1)
		goto L354
	} else {
		goto L357
	}
L356:
	;
	v1528 = int32(0)
	goto L354
L357:
	;
	v1441 = int32(1)
	v1442 = v1423 - v1441
	v1444 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1420+v1442))))
	v1446 = v1444 & int32(255)
	if v1442 == v1424 {
		v1501 = v1446
		v1502 = v1441
		goto L358
	} else {
		goto L359
	}
L358:
	;
	if int32(116) < v1501 {
		goto L367
	} else {
		goto L368
	}
L359:
	;
	if int32(0) <= v1444 {
		v1501 = v1446
		v1502 = v1441
		goto L358
	} else {
		goto L360
	}
L360:
	;
	v1452 = v1446 & int32(63)
	v1454 = v1423 - int32(2)
	v1456 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1420+v1454))))
	v1458 = v1456 << (uint(int32(6)) % 32)
	if base.B2i32(v1454 != v1424)&base.B2i32(base.Ui32(v1456) < base.Ui32(int32(192))) == int32(0) {
		goto L361
	} else {
		goto L362
	}
L361:
	;
	v1501 = v1458&int32(1984) | v1452
	v1502 = int32(2)
	goto L358
L362:
	;
	goto L363
L363:
	;
	v1471 = v1458&int32(4032) | v1452
	v1473 = v1423 - int32(3)
	v1475 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1420+v1473))))
	if base.B2i32(v1473 != v1424)&base.B2i32(base.Ui32(v1475) < base.Ui32(int32(224))) == int32(0) {
		goto L364
	} else {
		goto L365
	}
L364:
	;
	v1501 = v1475<<(uint(int32(12))%32)&int32(_a_F_german_UTF_8_stem_1) | v1471
	v1502 = int32(3)
	goto L358
L365:
	;
	goto L366
L366:
	;
	v1493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1423+(v1420-int32(4))))))
	v1501 = v1475<<(uint(int32(12))%32)&int32(_a_F_german_UTF_8_stem_14) | v1493&int32(7)<<(uint(int32(18))%32) | v1471
	v1502 = int32(4)
	goto L358
L367:
	;
	v1535 = v1502
	goto L353
L368:
	;
	goto L369
L369:
	;
	v1506 = v1501 - int32(98)
	if v1506 < int32(0) {
		goto L370
	} else {
		goto L371
	}
L370:
	;
	v1535 = v1502
	goto L353
L371:
	;
	goto L372
L372:
	;
	v1512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1506)>>(uint(int32(3))%32)))+uint32(_c_F_german_UTF_8_stem[2]))))
	if int32(base.Ui32(v1512)>>(uint(v1506&int32(7))%32))&int32(1) == int32(0) {
		goto L373
	} else {
		goto L374
	}
L373:
	;
	v1535 = v1502
	goto L353
L374:
	;
	goto L375
L375:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1423 - v1502
	goto L376
L376:
	;
	goto L356
L377:
	;
	v1536 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1537 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1538 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L380
L378:
	;
	if v1589 < int32(0) {
		goto L342
	} else {
		goto L398
	}
L380:
	;
	goto L381
L381:
	;
	goto L382
L382:
	;
	v1545 = v1537
	v1547 = int32(3)
	goto L385
L384:
	;
	v1589 = v1571
	goto L378
L385:
	;
	if v1545 <= v1538 {
		goto L387
	} else {
		goto L388
	}
L386:
	;
	goto L384
L387:
	;
	v1589 = int32(-1)
	goto L378
L388:
	;
	goto L389
L389:
	;
	v1552 = v1545 - int32(1)
	v1554 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1536+v1552))))
	if int32(0) <= v1554 {
		v1571 = v1552
		goto L390
	} else {
		goto L391
	}
L390:
	;
	v1575 = int32(1)
	if v1575 < v1547 {
		v1545 = v1571
		v1547 = v1547 - v1575
		goto L385
	} else {
		goto L397
	}
L391:
	;
	if v1552 <= v1538 {
		v1571 = v1552
		goto L390
	} else {
		goto L392
	}
L392:
	;
	v1559 = v1552
	goto L393
L393:
	;
	v1564 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1536+v1559))))
	if base.Ui32(int32(191)) < base.Ui32(v1564) {
		v1571 = v1559
		goto L390
	} else {
		goto L395
	}
L394:
	;
	v1571 = v1538
	goto L390
L395:
	;
	v1568 = v1559 - int32(1)
	if v1538 < v1568 {
		v1559 = v1568
		goto L393
	} else {
		goto L396
	}
L396:
	;
	goto L394
L397:
	;
	goto L386
L398:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1589
	v1593 = F_slice_del(m, l0)
	mBase = m.M
	v1594 = m.ExcPending
	if v1594 != 0 {
		goto L112
	} else {
		goto L399
	}
L399:
	;
	if v1593 < int32(0) {
		v1893 = v1593
		goto L28
	} else {
		goto L400
	}
L400:
	;
	goto L342
L401:
	;
	v1785 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1785
	v1788 = v1785
	goto L460
L402:
	;
	v1606 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1606+v1603))))
	if v1608&int32(224) != int32(96) {
		goto L401
	} else {
		goto L403
	}
L403:
	;
	if int32(1)<<(uint(v1608)%32)&int32(_a_F_german_UTF_8_stem_17) == int32(0) {
		goto L401
	} else {
		goto L404
	}
L404:
	;
	v1621 = F_find_among_b(m, l0, int32(_a_F_german_UTF_8_stem_18), int32(8))
	mBase = m.M
	v1622 = m.ExcPending
	if v1622 != 0 {
		goto L112
	} else {
		goto L405
	}
L405:
	;
	if v1621 == int32(0) {
		goto L401
	} else {
		goto L406
	}
L406:
	;
	v1625 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1625
	v1627 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1628 = *(*int32)(unsafe.Add(mBase, uint32(v1627)+4))
	if v1625 < v1628 {
		goto L401
	} else {
		goto L407
	}
L407:
	;
	switch v1621 - int32(1) {
	case 0:
		goto L411
	case 1:
		goto L410
	case 2:
		goto L409
	case 3:
		goto L408
	default:
		goto L401
	}
L408:
	;
	v1743 = F_slice_del(m, l0)
	mBase = m.M
	v1744 = m.ExcPending
	if v1744 != 0 {
		goto L112
	} else {
		goto L449
	}
L409:
	;
	v1687 = F_slice_del(m, l0)
	mBase = m.M
	v1688 = m.ExcPending
	if v1688 != 0 {
		goto L112
	} else {
		goto L432
	}
L410:
	;
	v1674 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1674 < v1625 {
		goto L426
	} else {
		goto L427
	}
L411:
	;
	v1632 = F_slice_del(m, l0)
	mBase = m.M
	v1633 = m.ExcPending
	if v1633 != 0 {
		goto L112
	} else {
		goto L412
	}
L412:
	;
	if v1632 < int32(0) {
		v1893 = v1632
		goto L28
	} else {
		goto L413
	}
L413:
	;
	v1636 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1636
	v1638 = int32(2)
	v1640 = int32(0)
	v1643 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1636-v1643 < v1638 {
		v1653 = v1640
		goto L415
	} else {
		goto L416
	}
L414:
	;
	if v1653 == int32(0) {
		goto L401
	} else {
		goto L418
	}
L415:
	;
	goto L414
L416:
	;
	v1646 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1649 = F_memcmp(m, v1646+v1636-v1638, int32(_a_F_german_UTF_8_stem_19), v1638)
	mBase = m.M
	if v1649 != 0 {
		v1653 = v1640
		goto L415
	} else {
		goto L417
	}
L417:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1636 - v1638
	v1653 = int32(1)
	goto L415
L418:
	;
	v1656 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1656
	v1658 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1658 < v1656 {
		goto L419
	} else {
		goto L420
	}
L419:
	;
	v1660 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1660+v1656-int32(1)))))
	if v1664 == int32(101) {
		goto L401
	} else {
		goto L422
	}
L420:
	;
	goto L421
L421:
	;
	v1667 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1668 = *(*int32)(unsafe.Add(mBase, uint32(v1667)+4))
	if v1656 < v1668 {
		goto L401
	} else {
		goto L423
	}
L422:
	;
	goto L421
L423:
	;
	v1670 = F_slice_del(m, l0)
	mBase = m.M
	v1671 = m.ExcPending
	if v1671 != 0 {
		goto L112
	} else {
		goto L424
	}
L424:
	;
	if int32(0) <= v1670 {
		goto L401
	} else {
		goto L425
	}
L425:
	;
	v1893 = v1670
	goto L28
L426:
	;
	v1676 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1676+v1625-int32(1)))))
	if v1680 == int32(101) {
		goto L401
	} else {
		goto L429
	}
L427:
	;
	goto L428
L428:
	;
	v1683 = F_slice_del(m, l0)
	mBase = m.M
	v1684 = m.ExcPending
	if v1684 != 0 {
		goto L112
	} else {
		goto L430
	}
L429:
	;
	goto L428
L430:
	;
	if int32(0) <= v1683 {
		goto L401
	} else {
		goto L431
	}
L431:
	;
	v1893 = v1683
	goto L28
L432:
	;
	if v1687 < int32(0) {
		v1893 = v1687
		goto L28
	} else {
		goto L433
	}
L433:
	;
	v1691 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1691
	v1693 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1694 = int32(2)
	v1696 = int32(0)
	v1699 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1691-v1699 < v1694 {
		v1709 = v1696
		goto L435
	} else {
		goto L436
	}
L434:
	;
	if v1709 == int32(0) {
		goto L438
	} else {
		goto L439
	}
L435:
	;
	goto L434
L436:
	;
	v1702 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1705 = F_memcmp(m, v1702+v1691-v1694, int32(_a_F_german_UTF_8_stem_20), v1694)
	mBase = m.M
	if v1705 != 0 {
		v1709 = v1696
		goto L435
	} else {
		goto L437
	}
L437:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1691 - v1694
	v1709 = int32(1)
	goto L435
L438:
	;
	v1712 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1714 = v1712 + (v1691 - v1693)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1714
	v1716 = int32(2)
	v1718 = int32(0)
	v1721 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1714-v1721 < v1716 {
		v1731 = v1718
		goto L442
	} else {
		goto L443
	}
L439:
	;
	goto L440
L440:
	;
	v1734 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1734
	v1736 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1737 = *(*int32)(unsafe.Add(mBase, uint32(v1736)+8))
	if v1734 < v1737 {
		goto L401
	} else {
		goto L446
	}
L441:
	;
	if v1731 == int32(0) {
		goto L401
	} else {
		goto L445
	}
L442:
	;
	goto L441
L443:
	;
	v1724 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1727 = F_memcmp(m, v1724+v1714-v1716, int32(_a_F_german_UTF_8_stem_21), v1716)
	mBase = m.M
	if v1727 != 0 {
		v1731 = v1718
		goto L442
	} else {
		goto L444
	}
L444:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1714 - v1716
	v1731 = int32(1)
	goto L442
L445:
	;
	goto L440
L446:
	;
	v1739 = F_slice_del(m, l0)
	mBase = m.M
	v1740 = m.ExcPending
	if v1740 != 0 {
		goto L112
	} else {
		goto L447
	}
L447:
	;
	if int32(0) <= v1739 {
		goto L401
	} else {
		goto L448
	}
L448:
	;
	v1893 = v1739
	goto L28
L449:
	;
	if v1743 < int32(0) {
		v1893 = v1743
		goto L28
	} else {
		goto L450
	}
L450:
	;
	v1747 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1747
	v1750 = v1747 - int32(1)
	v1751 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1750 <= v1751 {
		goto L401
	} else {
		goto L451
	}
L451:
	;
	v1753 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1755 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1753+v1750))))
	if base.Ui32(int32(1)) < base.Ui32((v1755-int32(103))&int32(255)) {
		goto L401
	} else {
		goto L452
	}
L452:
	;
	v1762 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1765 = F_find_among_b(m, l0, int32(_a_F_german_UTF_8_stem_22), int32(2))
	mBase = m.M
	v1766 = m.ExcPending
	if v1766 != 0 {
		goto L112
	} else {
		goto L453
	}
L453:
	;
	if v1765 == int32(0) {
		goto L401
	} else {
		goto L454
	}
L454:
	;
	v1769 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1769
	v1771 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1772 = *(*int32)(unsafe.Add(mBase, uint32(v1771)+4))
	if v1769 < v1772 {
		goto L455
	} else {
		goto L456
	}
L455:
	;
	v1774 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1774 + (v1747 - v1762)
	goto L401
L456:
	;
	goto L457
L457:
	;
	v1778 = F_slice_del(m, l0)
	mBase = m.M
	v1779 = m.ExcPending
	if v1779 != 0 {
		goto L112
	} else {
		goto L458
	}
L458:
	;
	if v1778 < int32(0) {
		v1893 = v1778
		goto L28
	} else {
		goto L459
	}
L459:
	;
	goto L401
L460:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1788
	v1795 = F_find_among(m, l0, int32(_a_F_german_UTF_8_stem_23), int32(6))
	mBase = m.M
	v1796 = m.ExcPending
	if v1796 != 0 {
		goto L112
	} else {
		goto L463
	}
L461:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1785
	v1893 = int32(1)
	goto L28
L462:
	;
	goto L461
L463:
	;
	v1797 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1797
	switch v1795 - int32(1) {
	case 0:
		goto L469
	case 1:
		goto L468
	case 2:
		goto L467
	case 3:
		goto L466
	case 4:
		goto L465
	default:
		goto L464
	}
L464:
	;
	v1883 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1788 = v1883
	goto L460
L465:
	;
	v1825 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1826 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L480
L466:
	;
	v1821 = F_slice_from_s(m, l0, int32(1), int32(_a_F_german_UTF_8_stem_24))
	mBase = m.M
	v1822 = m.ExcPending
	if v1822 != 0 {
		goto L112
	} else {
		goto L476
	}
L467:
	;
	v1815 = F_slice_from_s(m, l0, int32(1), int32(_a_F_german_UTF_8_stem_25))
	mBase = m.M
	v1816 = m.ExcPending
	if v1816 != 0 {
		goto L112
	} else {
		goto L474
	}
L468:
	;
	v1809 = F_slice_from_s(m, l0, int32(1), int32(_a_F_german_UTF_8_stem_26))
	mBase = m.M
	v1810 = m.ExcPending
	if v1810 != 0 {
		goto L112
	} else {
		goto L472
	}
L469:
	;
	v1803 = F_slice_from_s(m, l0, int32(1), int32(_a_F_german_UTF_8_stem_27))
	mBase = m.M
	v1804 = m.ExcPending
	if v1804 != 0 {
		goto L112
	} else {
		goto L470
	}
L470:
	;
	if int32(0) <= v1803 {
		goto L464
	} else {
		goto L471
	}
L471:
	;
	v1893 = v1803
	goto L28
L472:
	;
	if int32(0) <= v1809 {
		goto L464
	} else {
		goto L473
	}
L473:
	;
	v1893 = v1809
	goto L28
L474:
	;
	if int32(0) <= v1815 {
		goto L464
	} else {
		goto L475
	}
L475:
	;
	v1893 = v1815
	goto L28
L476:
	;
	if int32(0) <= v1821 {
		goto L464
	} else {
		goto L477
	}
L477:
	;
	v1893 = v1821
	goto L28
L478:
	;
	if v1878 < int32(0) {
		goto L462
	} else {
		goto L498
	}
L480:
	;
	goto L481
L481:
	;
	goto L482
L482:
	;
	v1833 = v1797
	v1835 = int32(1)
	goto L485
L484:
	;
	v1878 = v1863
	goto L478
L485:
	;
	if v1826 <= v1833 {
		goto L487
	} else {
		goto L488
	}
L486:
	;
	goto L484
L487:
	;
	v1878 = int32(-1)
	goto L478
L488:
	;
	goto L489
L489:
	;
	v1840 = v1833 + int32(1)
	v1842 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1825+v1833))))
	if base.Ui32(v1842) < base.Ui32(int32(192)) {
		v1863 = v1840
		goto L490
	} else {
		goto L491
	}
L490:
	;
	v1864 = int32(1)
	if v1864 < v1835 {
		v1833 = v1863
		v1835 = v1835 - v1864
		goto L485
	} else {
		goto L497
	}
L491:
	;
	if v1826 <= v1840 {
		v1863 = v1840
		goto L490
	} else {
		goto L492
	}
L492:
	;
	v1849 = v1840
	goto L493
L493:
	;
	v1852 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1825+v1849))))
	if int32(-65) < v1852 {
		v1863 = v1849
		goto L490
	} else {
		goto L495
	}
L494:
	;
	v1863 = v1826
	goto L490
L495:
	;
	v1856 = v1849 + int32(1)
	if v1856 != v1826 {
		v1849 = v1856
		goto L493
	} else {
		goto L496
	}
L496:
	;
	goto L494
L497:
	;
	goto L486
L498:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1878
	goto L464
L499:
	;
	if int32(0) <= v1888 {
		goto L27
	} else {
		goto L500
	}
L500:
	;
	v1893 = v1888
	goto L28
}
func F_getIdentitySequence(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+131)))
	if v13 == int32(1) {
		v16 = F_get_partition_ancestors(m, v11)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v21 = F_get_attname(m, v11, l1, int32(0))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v23+v24<<(uint(int32(2))%32)-int32(4))))
				v31 = F_get_attnum(m, v30, v21)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					if v31 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v84 = m.ExcPending
						if v84 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v30
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v21
							F_errmsg_internal(m, int32(_a_F_getIdentitySequence_0), v9)
							mBase = m.M
							v89 = m.ExcPending
							if v89 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_getIdentitySequence_1), int32(1036), int32(_a_F_getIdentitySequence_2))
								mBase = m.M
								v94 = m.ExcPending
								if v94 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						F_list_free(m, v16)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							v38 = v31
							v39 = v30
							v42 = F_getOwnedSequences_internal(m, v39, v38, int32(105))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								if v42 != 0 {
									v44 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
									if v44 < int32(2) {
										v74 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
										v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
										v76 = v75
										m.G0 = v9 + int32(16)
										return v76
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v50 = m.ExcPending
										if v50 != 0 {
											return int32(0)
										} else {
											F_errmsg_internal(m, int32(_a_F_getIdentitySequence_3), int32(0))
											mBase = m.M
											v54 = m.ExcPending
											if v54 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_getIdentitySequence_1), int32(1042), int32(_a_F_getIdentitySequence_2))
												mBase = m.M
												v59 = m.ExcPending
												if v59 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								} else {
									if l2 != 0 {
										v76 = int32(0)
										m.G0 = v9 + int32(16)
										return v76
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v64 = m.ExcPending
										if v64 != 0 {
											return int32(0)
										} else {
											F_errmsg_internal(m, int32(_a_F_getIdentitySequence_4), int32(0))
											mBase = m.M
											v68 = m.ExcPending
											if v68 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_getIdentitySequence_1), int32(1048), int32(_a_F_getIdentitySequence_2))
												mBase = m.M
												v73 = m.ExcPending
												if v73 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v38 = l1
		v39 = v11
		v42 = F_getOwnedSequences_internal(m, v39, v38, int32(105))
		mBase = m.M
		v43 = m.ExcPending
		if v43 != 0 {
			return int32(0)
		} else {
			if v42 != 0 {
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
				if v44 < int32(2) {
					v74 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
					v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
					v76 = v75
					m.G0 = v9 + int32(16)
					return v76
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(_a_F_getIdentitySequence_3), int32(0))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_getIdentitySequence_1), int32(1042), int32(_a_F_getIdentitySequence_2))
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				if l2 != 0 {
					v76 = int32(0)
					m.G0 = v9 + int32(16)
					return v76
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(_a_F_getIdentitySequence_4), int32(0))
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_getIdentitySequence_1), int32(1048), int32(_a_F_getIdentitySequence_2))
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_getOwnedSequences_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	v8 = m.G0
	v10 = v8 - int32(144)
	m.G0 = v10
	v15 = F_table_open(m, int32(2608), int32(1))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	F_ScanKeyInit(m, v10, int32(4), int32(3), int32(184), int32(1259))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_ScanKeyInit(m, v10+int32(48), int32(5), int32(3), int32(184), l0)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if l1 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v43 = int32(0)
	v47 = F_systable_beginscan(m, v15, int32(2674), int32(1), v43, v42, v10)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L10
	}
L6:
	;
	v42 = int32(2)
	goto L5
L7:
	;
	goto L8
L8:
	;
	F_ScanKeyInit(m, v10+int32(96), int32(6), int32(3), int32(65), l1)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v42 = int32(3)
	goto L5
L10:
	;
	v49 = F_systable_getnext(m, v47)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	if v49 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v53 = v43
	v54 = v49
	goto L15
L13:
	;
	v86 = v43
	goto L14
L14:
	;
	F_systable_endscan(m, v47)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L31
	}
L15:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+22)))
	v62 = v60 + v61
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	if v63 != int32(1259) {
		v83 = v53
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v86 = v83
	goto L14
L17:
	;
	v84 = F_systable_getnext(m, v47)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L29
	}
L18:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
	if v66 != 0 {
		v83 = v53
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v62)+20))
	if v67 == int32(0) {
		v83 = v53
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+24)))
	switch v70 - int32(97) {
	case 0, 8:
		goto L21
	default:
		v83 = v53
		goto L17
	}
L21:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	v74 = F_get_rel_relkind(m, v73)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	if v74 != int32(83) {
		v83 = v53
		goto L17
	} else {
		goto L23
	}
L23:
	;
	if l2 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+24)))
	if v78 != l2&int32(255) {
		v83 = v53
		goto L17
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	v81 = F_lappend_oid(m, v53, v80)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L28
	}
L27:
	;
	goto L26
L28:
	;
	v83 = v81
	goto L17
L29:
	;
	if v84 != 0 {
		v53 = v83
		v54 = v84
		goto L15
	} else {
		goto L30
	}
L30:
	;
	goto L16
L31:
	;
	F_sequence_close(m, v15, int32(1))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	m.G0 = v10 + int32(144)
	return v86
}
func F_get_am_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	v4 = F_SearchSysCache1(m, int32(2), l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			return int32(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
			v17 = F_pstrdup(m, v12+v13+int32(4))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				F_ReleaseCatCache(m, v4)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					return v17
				}
			}
		}
	}
}
func F_get_am_type_oid(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v12 = F_SearchSysCache1(m, int32(1), l0)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if v12 == int32(0) {
			v29 = int32(0)
			if v29|l2 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v87 = m.ExcPending
				if v87 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(67137668))
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
						F_errmsg(m, int32(_a_F_get_am_type_oid_0), v9)
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_get_am_type_oid_1), int32(154), int32(_a_F_get_am_type_oid_2))
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				m.G0 = v9 + int32(32)
				return v29
			}
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+22)))
			v21 = v19 + v20
			if l1 != 0 {
				v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+72)))
				if v22 != l1&int32(255) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(325))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							v45 = m.G0
							v47 = v45 - int32(16)
							m.G0 = v47
							switch l1 - int32(105) {
							case 0:
								v66 = int32(_a_F_get_am_type_oid_3)
								v67 = int32(16)
								m.G0 = v47 + v67
								*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v66
								*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v21 + int32(4)
								F_errmsg(m, int32(_a_F_get_am_type_oid_4), v9+v67)
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_get_am_type_oid_1), int32(145), int32(_a_F_get_am_type_oid_2))
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							default:
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v47))) = l1
									F_errmsg_internal(m, int32(_a_F_get_am_type_oid_5), v47)
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_get_am_type_oid_1), int32(222), int32(_a_F_get_am_type_oid_6))
										mBase = m.M
										v64 = m.ExcPending
										if v64 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							case 11:
								v66 = int32(_a_F_get_am_type_oid_7)
								v67 = int32(16)
								m.G0 = v47 + v67
								*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v66
								*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v21 + int32(4)
								F_errmsg(m, int32(_a_F_get_am_type_oid_4), v9+v67)
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_get_am_type_oid_1), int32(145), int32(_a_F_get_am_type_oid_2))
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					}
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
					F_ReleaseCatCache(m, v12)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						v29 = v26
						if v29|l2 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v87 = m.ExcPending
							if v87 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(67137668))
								mBase = m.M
								v90 = m.ExcPending
								if v90 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
									F_errmsg(m, int32(_a_F_get_am_type_oid_0), v9)
									mBase = m.M
									v94 = m.ExcPending
									if v94 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_get_am_type_oid_1), int32(154), int32(_a_F_get_am_type_oid_2))
										mBase = m.M
										v99 = m.ExcPending
										if v99 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							m.G0 = v9 + int32(32)
							return v29
						}
					}
				}
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
				F_ReleaseCatCache(m, v12)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					v29 = v26
					if v29|l2 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(67137668))
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
								F_errmsg(m, int32(_a_F_get_am_type_oid_0), v9)
								mBase = m.M
								v94 = m.ExcPending
								if v94 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_get_am_type_oid_1), int32(154), int32(_a_F_get_am_type_oid_2))
									mBase = m.M
									v99 = m.ExcPending
									if v99 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						m.G0 = v9 + int32(32)
						return v29
					}
				}
			}
		}
	}
}
func F_get_attgenerated(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = F_SearchSysCache2(m, int32(7), l0, l1)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l0
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
				F_errmsg_internal(m, int32(_a_F_get_attgenerated_0), v7)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_get_attgenerated_1), int32(991), int32(_a_F_get_attgenerated_2))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+22)))
			v33 = int32(*(*int8)(unsafe.Add(mBase, uint32(v30+v31)+90)))
			F_ReleaseCatCache(m, v10)
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				m.G0 = v7 + int32(16)
				return v33
			}
		}
	}
}
func F_get_backend_type_for_log(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_get_backend_type_for_log[0]))
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_get_backend_type_for_log[1]))
	if v4 != v6 {
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_get_backend_type_for_log[2]))
		if v9 == int32(5) {
			v13 = *(*int32)(unsafe.Add(mBase, _c_F_get_backend_type_for_log[3]))
			return v13 + int32(96)
		} else {
			if base.Ui32(v9) <= base.Ui32(int32(17)) {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v9<<(uint(int32(2))%32))+uint32(_c_F_get_backend_type_for_log[4])))
				v26 = v25
			} else {
				v26 = int32(_a_F_get_backend_type_for_log_0)
			}
			v28 = v26
			return v28
		}
	} else {
		v28 = int32(_a_F_get_backend_type_for_log_1)
		return v28
	}
}
func F_get_eclass_for_sort_expr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v104 int32
	_ = v104
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int64
	_ = v335
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v475 int32
	_ = v475
	var v486 int32
	_ = v486
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v566 int32
	_ = v566
	var v590 int32
	_ = v590
	var v611 int32
	_ = v611
	var v615 int32
	_ = v615
	var v620 int32
	_ = v620
	v19 = F_canonicalize_ec_expression(m, l1, l3, l4)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v26 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L1
	} else {
		goto L135
	}
L4:
	;
	return v590
L5:
	;
	v320 = int32(0)
	if l7 == v320 {
		v590 = v320
		goto L4
	} else {
		goto L73
	}
L6:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v29 <= int32(0) {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v42 = int32(0)
	goto L8
L8:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v50+v42<<(uint(int32(2))%32))))
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+41)))
	if v55 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L5
L10:
	;
	v299 = v42 + int32(1)
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v299 < v300 {
		v42 = v299
		goto L8
	} else {
		goto L72
	}
L11:
	;
	if l5 == int32(0) {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	if l4 != v62 {
		goto L10
	} else {
		goto L16
	}
L14:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v54)+44))
	if l5 != v60 {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	v65 = F_equal(m, l2, v64)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	if v65 == int32(0) {
		goto L10
	} else {
		goto L18
	}
L18:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
	if v69 == int32(0) {
		goto L10
	} else {
		goto L19
	}
L19:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v54)+20))
	if v73 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v74 = l6
	goto L22
L21:
	;
	v74 = int32(0)
	goto L22
L22:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	v85 = int32(-1)
	v86 = v69
	v89 = v75
	goto L23
L23:
	;
	if v89 != 0 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v590 = v54
	goto L4
L25:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v194)))
	if v200 == int32(0) {
		goto L10
	} else {
		goto L45
	}
L26:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v190 = v85
	v191 = v86
	v194 = v89
	v198 = v95
	goto L25
L27:
	;
	goto L28
L28:
	;
	v104 = v85
	goto L29
L29:
	;
	if v74 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v178)+12))
	v190 = v169
	v191 = v178
	v194 = v181
	v198 = v181
	goto L25
L31:
	;
	if v169 <= int32(0) {
		goto L10
	} else {
		goto L42
	}
L32:
	;
	v169 = base.I32_ctz(v155) | v156<<(uint(int32(5))%32)
	goto L31
L33:
	;
	v169 = int32(-2)
	goto L31
L34:
	;
	v120 = v104 + int32(1)
	v122 = base.I32_div_s(v120, int32(32))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	if v123 <= v122 {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v126 = v74 + int32(8)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v126+v122<<(uint(int32(2))%32))))
	v133 = v130 & (int32(-1) << (uint(v120) % 32))
	if v133 != 0 {
		v155 = v133
		v156 = v122
		goto L32
	} else {
		goto L36
	}
L36:
	;
	v135 = v122 + int32(1)
	if v135 == v123 {
		goto L33
	} else {
		goto L37
	}
L37:
	;
	v138 = v135
	goto L38
L38:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v126+v138<<(uint(int32(2))%32))))
	if v145 != 0 {
		v155 = v145
		v156 = v138
		goto L32
	} else {
		goto L40
	}
L39:
	;
	goto L33
L40:
	;
	v147 = v138 + int32(1)
	if v147 != v123 {
		v138 = v147
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v54)+12))
	if v172 <= v169 {
		goto L10
	} else {
		goto L43
	}
L43:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v54)+20))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v174+v169<<(uint(int32(2))%32))))
	if v178 == int32(0) {
		v104 = v169
		goto L29
	} else {
		goto L44
	}
L44:
	;
	goto L30
L45:
	;
	v204 = v194 + int32(4)
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v191)+4))
	if base.Ui32(v204) < base.Ui32(v198+v206<<(uint(int32(2))%32)) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v211 = v204
	goto L48
L47:
	;
	v211 = int32(0)
	goto L48
L48:
	;
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200)+13)))
	if v212 == int32(1) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v200)+8))
	v216 = int32(0)
	v223 = base.B2i32(v215|l6 == v216)
	if v215 == v216 {
		v262 = v223
		goto L53
	} else {
		goto L54
	}
L50:
	;
	goto L51
L51:
	;
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200)+12)))
	if v268 == int32(1) {
		goto L65
	} else {
		goto L66
	}
L52:
	;
	if v262 == int32(0) {
		v85 = v190
		v86 = v191
		v89 = v211
		goto L23
	} else {
		goto L64
	}
L53:
	;
	goto L52
L54:
	;
	if l6 == int32(0) {
		v262 = v223
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v215)+4))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v229 != v230 {
		v262 = int32(0)
		goto L53
	} else {
		goto L56
	}
L56:
	;
	v232 = int32(1)
	if v229 <= v232 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v235 = v232
	goto L59
L58:
	;
	v235 = v229
	goto L59
L59:
	;
	v236 = int32(8)
	v241 = int32(0)
	goto L60
L60:
	;
	v249 = v241 << (uint(int32(2)) % 32)
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v215+v236+v249)))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v249+(l6+v236))))
	v254 = base.B2i32(v251 == v253)
	if v253 != v251 {
		v262 = v254
		goto L53
	} else {
		goto L62
	}
L61:
	;
	v262 = v254
	goto L53
L62:
	;
	v257 = v241 + int32(1)
	if v257 != v235 {
		v241 = v257
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	goto L51
L65:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v200)+20))
	if v271 != v25 {
		v85 = v190
		v86 = v191
		v89 = v211
		goto L23
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v200)+16))
	if l3 != v273 {
		v85 = v190
		v86 = v191
		v89 = v211
		goto L23
	} else {
		goto L69
	}
L68:
	;
	goto L67
L69:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v200)+4))
	v276 = F_equal(m, v19, v275)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	if v276 == int32(0) {
		v85 = v190
		v86 = v191
		v89 = v211
		goto L23
	} else {
		goto L71
	}
L71:
	;
	goto L24
L72:
	;
	goto L9
L73:
	;
	v323 = int32(_a_F_get_eclass_for_sort_expr_0)
	v324 = *(*int32)(unsafe.Add(mBase, _c_F_get_eclass_for_sort_expr[0]))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	*(*int32)(unsafe.Add(mBase, _c_F_get_eclass_for_sort_expr[0])) = v326
	v329 = F_palloc0(m, int32(60))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v329))) = int32(273)
	v333 = F_list_copy(m, l2)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	v335 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v329)+12)) = v335
	*(*int32)(unsafe.Add(mBase, uint32(v329)+8)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v329)+4)) = v333
	*(*int64)(unsafe.Add(mBase, uint32(v329)+20)) = v335
	*(*int64)(unsafe.Add(mBase, uint32(v329)+28)) = v335
	*(*int64)(unsafe.Add(mBase, uint32(v329)+33)) = v335
	v345 = F_contain_volatile_functions(m, v19)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	v347 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v329)+56)) = v347
	*(*int64)(unsafe.Add(mBase, uint32(v329)+48)) = int64(4294967295)
	*(*int32)(unsafe.Add(mBase, uint32(v329)+44)) = l5
	*(*uint8)(unsafe.Add(mBase, uint32(v329)+42)) = uint8(v347)
	*(*uint8)(unsafe.Add(mBase, uint32(v329)+41)) = uint8(v345)
	if v345 != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v356 = l5
	goto L79
L78:
	;
	v356 = int32(1)
	goto L79
L79:
	;
	if v356 == int32(0) {
		goto L3
	} else {
		goto L80
	}
L80:
	;
	v359 = F_pull_varnos(m, l0, v19)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v361 = F_copyObjectImpl(m, v19)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	v364 = F_palloc0(m, int32(28))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v366 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v364)+24)) = v366
	*(*int32)(unsafe.Add(mBase, uint32(v364)+20)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v364)+16)) = l3
	*(*uint16)(unsafe.Add(mBase, uint32(v364)+12)) = uint16(v366)
	*(*int32)(unsafe.Add(mBase, uint32(v364)+8)) = v359
	*(*int32)(unsafe.Add(mBase, uint32(v364)+4)) = v361
	*(*int32)(unsafe.Add(mBase, uint32(v364))) = int32(274)
	if v359 == v366 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v378 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v364)+12)) = uint8(v378)
	*(*uint8)(unsafe.Add(mBase, uint32(v329)+40)) = uint8(v378)
	goto L86
L85:
	;
	goto L86
L86:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v329)+16))
	v383 = F_lappend(m, v382, v364)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v329)+16)) = v383
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v329)+36))
	v387 = F_bms_add_members(m, v386, v359)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v329)+36)) = v387
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v329)+40)))
	if v390 != int32(1) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v407 = F_lappend(m, v406, v329)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L1
	} else {
		goto L99
	}
L90:
	;
	v393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v329)+41)))
	if v393 != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v402 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v329)+40)) = uint8(v402)
	*(*uint8)(unsafe.Add(mBase, uint32(v364)+12)) = uint8(v402)
	goto L89
L92:
	;
	v394 = F_expression_returns_set(m, v19)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	if v394 != 0 {
		goto L91
	} else {
		goto L94
	}
L94:
	;
	v396 = F_contain_agg_clause(m, v19)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	if v396 != 0 {
		goto L91
	} else {
		goto L96
	}
L96:
	;
	v398 = F_contain_windowfuncs(m, v19)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	if v398 == int32(0) {
		goto L89
	} else {
		goto L98
	}
L98:
	;
	goto L91
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v407
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)))
	if v410 != int32(1) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_get_eclass_for_sort_expr[0])) = v324
	v590 = v329
	goto L4
L101:
	;
	if v407 != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v407)+4))
	v417 = v414 - int32(1)
	goto L104
L103:
	;
	v417 = int32(-1)
	goto L104
L104:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v329)+36))
	if v418 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L105:
	;
	if v475 <= int32(0) {
		goto L100
	} else {
		goto L116
	}
L106:
	;
	v475 = base.I32_ctz(v461) | v462<<(uint(int32(5))%32)
	goto L105
L107:
	;
	v475 = int32(-2)
	goto L105
L108:
	;
	v428 = base.I32_div_s(int32(0), int32(32))
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v418)+4))
	if v429 <= v428 {
		goto L107
	} else {
		goto L109
	}
L109:
	;
	v432 = v418 + int32(8)
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v432+v428<<(uint(int32(2))%32))))
	v439 = v436 & int32(-1)
	if v439 != 0 {
		v461 = v439
		v462 = v428
		goto L106
	} else {
		goto L110
	}
L110:
	;
	v441 = v428 + int32(1)
	if v441 == v429 {
		goto L107
	} else {
		goto L111
	}
L111:
	;
	v444 = v441
	goto L112
L112:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v432+v444<<(uint(int32(2))%32))))
	if v451 != 0 {
		v461 = v451
		v462 = v444
		goto L106
	} else {
		goto L114
	}
L113:
	;
	goto L107
L114:
	;
	v453 = v444 + int32(1)
	if v453 != v429 {
		v444 = v453
		goto L112
	} else {
		goto L115
	}
L115:
	;
	goto L113
L116:
	;
	v486 = v475
	goto L117
L117:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	if v486 == v496 {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	goto L100
L119:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v329)+36))
	if v510 == int32(0) {
		goto L125
	} else {
		goto L126
	}
L120:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v498+v486<<(uint(int32(2))%32))))
	if v502 == int32(0) {
		goto L119
	} else {
		goto L121
	}
L121:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v502)+136))
	v506 = F_bms_add_member(m, v505, v417)
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v502)+136)) = v506
	goto L119
L123:
	;
	if int32(0) < v566 {
		v486 = v566
		goto L117
	} else {
		goto L134
	}
L124:
	;
	v566 = base.I32_ctz(v552) | v553<<(uint(int32(5))%32)
	goto L123
L125:
	;
	v566 = int32(-2)
	goto L123
L126:
	;
	v517 = v486 + int32(1)
	v519 = base.I32_div_s(v517, int32(32))
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v510)+4))
	if v520 <= v519 {
		goto L125
	} else {
		goto L127
	}
L127:
	;
	v523 = v510 + int32(8)
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v523+v519<<(uint(int32(2))%32))))
	v530 = v527 & (int32(-1) << (uint(v517) % 32))
	if v530 != 0 {
		v552 = v530
		v553 = v519
		goto L124
	} else {
		goto L128
	}
L128:
	;
	v532 = v519 + int32(1)
	if v532 == v520 {
		goto L125
	} else {
		goto L129
	}
L129:
	;
	v535 = v532
	goto L130
L130:
	;
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v523+v535<<(uint(int32(2))%32))))
	if v542 != 0 {
		v552 = v542
		v553 = v535
		goto L124
	} else {
		goto L132
	}
L131:
	;
	goto L125
L132:
	;
	v544 = v535 + int32(1)
	if v544 != v520 {
		v535 = v544
		goto L130
	} else {
		goto L133
	}
L133:
	;
	goto L131
L134:
	;
	goto L118
L135:
	;
	F_errmsg_internal(m, int32(_a_F_get_eclass_for_sort_expr_1), int32(0))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(_a_F_get_eclass_for_sort_expr_2), int32(838), int32(_a_F_get_eclass_for_sort_expr_3))
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_ext_ver_list(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v472 int32
	_ = v472
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v526 int32
	_ = v526
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v591 int32
	_ = v591
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v607 int32
	_ = v607
	var v619 int32
	_ = v619
	v2 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v19&int32(3) == v2 {
		v43 = v19
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v77 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L2:
	;
	v76 = v68 - v19
	goto L1
L3:
	;
	v47 = v43
	goto L12
L4:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v27 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v76 = int32(0)
	goto L1
L6:
	;
	goto L7
L7:
	;
	v32 = v19
	goto L8
L8:
	;
	v36 = v32 + int32(1)
	if v36&int32(3) == int32(0) {
		v43 = v36
		goto L3
	} else {
		goto L10
	}
L9:
	;
	v68 = v36
	goto L2
L10:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	if v41 != 0 {
		v32 = v36
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v56 = int32(-2139062144)
	if (int32(16843008)-v53|v53)&v56 == v56 {
		v47 = v47 + int32(4)
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v62 = v47
	goto L15
L14:
	;
	goto L13
L15:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	if v66 != 0 {
		v62 = v62 + int32(1)
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v68 = v62
	goto L2
L17:
	;
	goto L16
L18:
	;
	F_FreeDir(m, v98)
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L23
	} else {
		goto L150
	}
L19:
	;
	v98 = F_AllocateDir(m, v97)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L23
	} else {
		goto L30
	}
L20:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v81 = F_pstrdup(m, v80)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	if v85 == int32(47) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	return int32(0)
L24:
	;
	v97 = v81
	goto L19
L25:
	;
	v88 = F_pstrdup(m, v77)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L23
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v90
	v94 = F_psprintf(m, int32(_a_F_get_ext_ver_list_0), v17)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L23
	} else {
		goto L29
	}
L28:
	;
	v97 = v88
	goto L19
L29:
	;
	v97 = v94
	goto L19
L30:
	;
	v100 = F_ReadDir(m, v98, v97)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L23
	} else {
		goto L31
	}
L31:
	;
	if v100 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v607 = v2
	goto L18
L33:
	;
	goto L34
L34:
	;
	v107 = v100
	v109 = v2
	goto L35
L35:
	;
	v121 = v107 + int32(19)
	v125 = F_strlen(m, v121)
	mBase = m.M
	v132 = v125 + int32(1)
	goto L40
L36:
	;
	v607 = v591
	goto L18
L37:
	;
	v602 = F_ReadDir(m, v98, v97)
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L23
	} else {
		goto L148
	}
L38:
	;
	if v144 == int32(0) {
		v591 = v109
		goto L37
	} else {
		goto L44
	}
L39:
	;
	goto L38
L40:
	;
	v134 = int32(0)
	if v132 == v134 {
		v144 = v134
		goto L39
	} else {
		goto L42
	}
L41:
	;
	v144 = v139
	goto L39
L42:
	;
	v138 = v132 - int32(1)
	v139 = v121 + v138
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139))))
	if v140 != int32(46) {
		v132 = v138
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v147 = int32(_a_F_get_ext_ver_list_1)
	v150 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_get_ext_ver_list[0])))
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144))))
	if v151 == int32(0) {
		v170 = v150
		v171 = v151
		goto L46
	} else {
		goto L47
	}
L45:
	;
	if v171-v170 != 0 {
		v591 = v109
		goto L37
	} else {
		goto L53
	}
L46:
	;
	goto L45
L47:
	;
	if v150 != v151 {
		v170 = v150
		v171 = v151
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v155 = v144
	v156 = v147
	goto L49
L49:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156)+1)))
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+1)))
	if v160 == int32(0) {
		v170 = v159
		v171 = v160
		goto L46
	} else {
		goto L51
	}
L50:
	;
	v170 = v159
	v171 = v160
	goto L46
L51:
	;
	v163 = int32(1)
	if v159 == v160 {
		v155 = v155 + v163
		v156 = v156 + v163
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v76 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	if v217 != 0 {
		v591 = v109
		goto L37
	} else {
		goto L68
	}
L55:
	;
	v217 = int32(0)
	goto L54
L56:
	;
	goto L57
L57:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
	if v179 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v180 = v121
	v181 = v173
	v182 = v76
	v183 = v179
	goto L62
L59:
	;
	v205 = v173
	v209 = int32(0)
	goto L60
L60:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
	v217 = v209 - v210
	goto L54
L61:
	;
	v205 = v200
	v209 = v202
	goto L60
L62:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181))))
	if v183 != v185 {
		v200 = v181
		v202 = v183
		goto L61
	} else {
		goto L64
	}
L63:
	;
	v200 = v194
	v202 = int32(0)
	goto L61
L64:
	;
	if v185 == int32(0) {
		v200 = v181
		v202 = v183
		goto L61
	} else {
		goto L65
	}
L65:
	;
	v190 = v182 - int32(1)
	if v190 == int32(0) {
		v200 = v181
		v202 = v183
		goto L61
	} else {
		goto L66
	}
L66:
	;
	v193 = int32(1)
	v194 = v181 + v193
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+1)))
	if v195 != 0 {
		v180 = v180 + v193
		v181 = v194
		v182 = v190
		v183 = v195
		goto L62
	} else {
		goto L67
	}
L67:
	;
	goto L63
L68:
	;
	v218 = v121 + v76
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218))))
	if v219 != int32(45) {
		v591 = v109
		goto L37
	} else {
		goto L69
	}
L69:
	;
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121+(v76+int32(1))))))
	if v223 != int32(45) {
		v591 = v109
		goto L37
	} else {
		goto L70
	}
L70:
	;
	v228 = F_pstrdup(m, v218+int32(2))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L23
	} else {
		goto L71
	}
L71:
	;
	v233 = F_strlen(m, v228)
	mBase = m.M
	v240 = v233 + int32(1)
	goto L74
L72:
	;
	v253 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v252))) = uint8(v253)
	v256 = F_strstr(m, v228, int32(_a_F_get_ext_ver_list_2))
	mBase = m.M
	if v256 == v253 {
		goto L78
	} else {
		goto L79
	}
L73:
	;
	goto L72
L74:
	;
	v242 = int32(0)
	if v240 == v242 {
		v252 = v242
		goto L73
	} else {
		goto L76
	}
L75:
	;
	v252 = v247
	goto L73
L76:
	;
	v246 = v240 - int32(1)
	v247 = v228 + v246
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v247))))
	if v248 != int32(46) {
		v240 = v246
		goto L74
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	if v109 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L79:
	;
	goto L80
L80:
	;
	v359 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v256))) = uint8(v359)
	v362 = v256 + int32(2)
	v364 = F_strstr(m, v362, int32(_a_F_get_ext_ver_list_2))
	mBase = m.M
	if v364 != 0 {
		v591 = v109
		goto L37
	} else {
		goto L100
	}
L81:
	;
	v357 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v347)+8)) = uint8(v357)
	v591 = v346
	goto L37
L82:
	;
	v330 = F_palloc(m, int32(20))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L23
	} else {
		goto L97
	}
L83:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	if v261 <= int32(0) {
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
	v267 = int32(0)
	goto L85
L85:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v264+v267<<(uint(int32(2))%32))))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v283)))
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228))))
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284))))
	if v288 == int32(0) {
		v307 = v287
		v308 = v288
		goto L88
	} else {
		goto L89
	}
L86:
	;
	goto L82
L87:
	;
	if v308-v307 == int32(0) {
		v346 = v109
		v347 = v283
		goto L81
	} else {
		goto L95
	}
L88:
	;
	goto L87
L89:
	;
	if v287 != v288 {
		v307 = v287
		v308 = v288
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v292 = v284
	v293 = v228
	goto L91
L91:
	;
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v293)+1)))
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292)+1)))
	if v297 == int32(0) {
		v307 = v296
		v308 = v297
		goto L88
	} else {
		goto L93
	}
L92:
	;
	v307 = v296
	v308 = v297
	goto L88
L93:
	;
	v300 = int32(1)
	if v296 == v297 {
		v292 = v292 + v300
		v293 = v293 + v300
		goto L91
	} else {
		goto L94
	}
L94:
	;
	goto L92
L95:
	;
	v313 = v267 + int32(1)
	if v261 != v313 {
		v267 = v313
		goto L85
	} else {
		goto L96
	}
L96:
	;
	goto L86
L97:
	;
	v332 = F_pstrdup(m, v228)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L23
	} else {
		goto L98
	}
L98:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v330)+12)) = int64(2147483647)
	v336 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v330)+8)) = uint16(v336)
	*(*int32)(unsafe.Add(mBase, uint32(v330)+4)) = v336
	*(*int32)(unsafe.Add(mBase, uint32(v330))) = v332
	v341 = F_lappend(m, v109, v330)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L23
	} else {
		goto L99
	}
L99:
	;
	v346 = v341
	v347 = v330
	goto L81
L100:
	;
	if v109 == int32(0) {
		goto L104
	} else {
		goto L105
	}
L101:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v574)+4))
	v585 = F_lappend(m, v584, v572)
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L23
	} else {
		goto L147
	}
L102:
	;
	v557 = F_palloc(m, int32(20))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L23
	} else {
		goto L144
	}
L103:
	;
	if int32(0) < v466 {
		goto L129
	} else {
		goto L130
	}
L104:
	;
	v440 = F_palloc(m, int32(20))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L23
	} else {
		goto L122
	}
L105:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	if v367 <= int32(0) {
		goto L104
	} else {
		goto L106
	}
L106:
	;
	v370 = int32(0)
	if v370 < v367 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v373 = v367
	goto L109
L108:
	;
	v373 = v370
	goto L109
L109:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
	v377 = int32(0)
	goto L110
L110:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v374+v377<<(uint(int32(2))%32))))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v393)))
	v397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228))))
	v398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v394))))
	if v398 == int32(0) {
		v417 = v397
		v418 = v398
		goto L113
	} else {
		goto L114
	}
L111:
	;
	goto L104
L112:
	;
	if v418-v417 == int32(0) {
		v464 = v109
		v465 = v393
		v466 = v367
		v472 = v373
		goto L103
	} else {
		goto L120
	}
L113:
	;
	goto L112
L114:
	;
	if v397 != v398 {
		v417 = v397
		v418 = v398
		goto L113
	} else {
		goto L115
	}
L115:
	;
	v402 = v394
	v403 = v228
	goto L116
L116:
	;
	v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v403)+1)))
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v402)+1)))
	if v407 == int32(0) {
		v417 = v406
		v418 = v407
		goto L113
	} else {
		goto L118
	}
L117:
	;
	v417 = v406
	v418 = v407
	goto L113
L118:
	;
	v410 = int32(1)
	if v406 == v407 {
		v402 = v402 + v410
		v403 = v403 + v410
		goto L116
	} else {
		goto L119
	}
L119:
	;
	goto L117
L120:
	;
	v423 = v377 + int32(1)
	if v367 != v423 {
		v377 = v423
		goto L110
	} else {
		goto L121
	}
L121:
	;
	goto L111
L122:
	;
	v442 = F_pstrdup(m, v228)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L23
	} else {
		goto L123
	}
L123:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v440)+12)) = int64(2147483647)
	v446 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v440)+8)) = uint16(v446)
	*(*int32)(unsafe.Add(mBase, uint32(v440)+4)) = v446
	*(*int32)(unsafe.Add(mBase, uint32(v440))) = v442
	v452 = F_lappend(m, v109, v440)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L23
	} else {
		goto L124
	}
L124:
	;
	if v452 == int32(0) {
		v543 = v446
		v546 = v440
		goto L102
	} else {
		goto L125
	}
L125:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v452)+4))
	v457 = int32(0)
	if v457 < v456 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v460 = v456
	goto L128
L127:
	;
	v460 = v457
	goto L128
L128:
	;
	v464 = v452
	v465 = v440
	v466 = v456
	v472 = v460
	goto L103
L129:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v464)+12))
	v480 = int32(0)
	goto L132
L130:
	;
	goto L131
L131:
	;
	v543 = v464
	v546 = v465
	goto L102
L132:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v477+v480<<(uint(int32(2))%32))))
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v496)))
	v500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v362))))
	v501 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497))))
	if v501 == int32(0) {
		v520 = v500
		v521 = v501
		goto L135
	} else {
		goto L136
	}
L133:
	;
	goto L131
L134:
	;
	if v521-v520 == int32(0) {
		v572 = v496
		v573 = v464
		v574 = v465
		goto L101
	} else {
		goto L142
	}
L135:
	;
	goto L134
L136:
	;
	if v500 != v501 {
		v520 = v500
		v521 = v501
		goto L135
	} else {
		goto L137
	}
L137:
	;
	v505 = v497
	v506 = v362
	goto L138
L138:
	;
	v509 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v506)+1)))
	v510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v505)+1)))
	if v510 == int32(0) {
		v520 = v509
		v521 = v510
		goto L135
	} else {
		goto L140
	}
L139:
	;
	v520 = v509
	v521 = v510
	goto L135
L140:
	;
	v513 = int32(1)
	if v509 == v510 {
		v505 = v505 + v513
		v506 = v506 + v513
		goto L138
	} else {
		goto L141
	}
L141:
	;
	goto L139
L142:
	;
	v526 = v480 + int32(1)
	if v526 != v472 {
		v480 = v526
		goto L132
	} else {
		goto L143
	}
L143:
	;
	goto L133
L144:
	;
	v559 = F_pstrdup(m, v362)
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L23
	} else {
		goto L145
	}
L145:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v557)+12)) = int64(2147483647)
	v563 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v557)+8)) = uint16(v563)
	*(*int32)(unsafe.Add(mBase, uint32(v557)+4)) = v563
	*(*int32)(unsafe.Add(mBase, uint32(v557))) = v559
	v568 = F_lappend(m, v543, v557)
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L23
	} else {
		goto L146
	}
L146:
	;
	v572 = v557
	v573 = v568
	v574 = v546
	goto L101
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v574)+4)) = v585
	v591 = v573
	goto L37
L148:
	;
	if v602 != 0 {
		v107 = v602
		v109 = v591
		goto L35
	} else {
		goto L149
	}
L149:
	;
	goto L36
L150:
	;
	m.G0 = v17 + int32(16)
	return v607
}
func F_get_number_of_groups(m *base.Module, l0 int32, l1 float64, l2 int32, l3 int32) float64 {
	mBase := m.M
	_ = mBase
	var v16 float64
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v42 int32
	_ = v42
	var v46 float64
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v80 float64
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 float64
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v113 float64
	_ = v113
	var v114 int32
	_ = v114
	var v116 float64
	_ = v116
	var v117 float64
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 float64
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v148 float64
	_ = v148
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v200 float64
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v208 float64
	_ = v208
	var v209 int32
	_ = v209
	var v211 float64
	_ = v211
	var v231 float64
	_ = v231
	v16 = float64(0)
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+108))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v22)+100))
	if v24 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v20 + int32(16)
	return v231
L2:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v150 == int32(0) {
		v231 = v148
		goto L1
	} else {
		goto L35
	}
L3:
	;
	if v23 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	if v23 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L6:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v25 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	v122 = F_get_sortgrouplist_exprs(m, v121, l3)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L15
	} else {
		goto L30
	}
L9:
	;
	v148 = v16
	goto L2
L10:
	;
	goto L11
L11:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v28 <= int32(0) {
		v148 = v16
		goto L2
	} else {
		goto L12
	}
L12:
	;
	v42 = int32(0)
	v46 = v16
	goto L13
L13:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v48+v42<<(uint(int32(2))%32))))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v54 = F_get_sortgrouplist_exprs(m, v53, l3)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return float64(0)
L16:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v52)+16)) = int64(0)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	v68 = int32(0)
	v80 = float64(0)
	goto L17
L17:
	;
	v81 = int32(0)
	if v61 == v81 {
		v90 = v81
		goto L19
	} else {
		goto L20
	}
L19:
	;
	if v60 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v84 <= v68 {
		v90 = v81
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v90 = v86 + v68<<(uint(int32(2))%32)
	goto L19
L22:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v107
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	v113 = F_estimate_num_groups(m, l0, v54, l1, v20+int32(12), int32(0))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L15
	} else {
		goto L29
	}
L23:
	;
	v102 = base.F64_add(v46, v80)
	v104 = v42 + int32(1)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v105 <= v104 {
		v148 = v102
		goto L2
	} else {
		goto L28
	}
L24:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	if v93 <= v68 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	if v90 == int32(0) {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
	v100 = v97 + v68<<(uint(int32(2))%32)
	if v100 != 0 {
		goto L22
	} else {
		goto L27
	}
L27:
	;
	goto L23
L28:
	;
	v42 = v104
	v46 = v102
	goto L13
L29:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v109)+8)) = v113
	v116 = *(*float64)(unsafe.Add(mBase, uint32(v52)+16))
	v117 = base.F64_add(v113, v116)
	*(*float64)(unsafe.Add(mBase, uint32(v52)+16)) = v117
	v68 = v68 + int32(1)
	v80 = v117
	goto L17
L30:
	;
	v124 = int32(0)
	v126 = F_estimate_num_groups(m, l0, v122, l1, v124, v124)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L15
	} else {
		goto L31
	}
L31:
	;
	v231 = v126
	goto L1
L32:
	;
	v231 = float64(1)
	goto L1
L33:
	;
	goto L34
L34:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v231 = base.F64_convert_i32_s(v131)
	goto L1
L35:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2)+8)) = int64(0)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v22)+100))
	v156 = F_get_sortgrouplist_exprs(m, v155, l3)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L15
	} else {
		goto L36
	}
L36:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v165 = int32(0)
	goto L37
L37:
	;
	v178 = int32(0)
	if v159 == v178 {
		v188 = v178
		goto L39
	} else {
		goto L40
	}
L39:
	;
	if v158 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v159)+4))
	if v182 <= v165 {
		v188 = int32(0)
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v159)+12))
	v188 = v184 + v165<<(uint(int32(2))%32)
	goto L39
L42:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v202
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
	v208 = F_estimate_num_groups(m, l0, v156, l1, v20+int32(8), int32(0))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L15
	} else {
		goto L48
	}
L43:
	;
	v200 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
	v231 = base.F64_add(v148, v200)
	goto L1
L44:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
	if v191 <= v165 {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	if v188 == int32(0) {
		goto L43
	} else {
		goto L46
	}
L46:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v158)+12))
	v198 = v195 + v165<<(uint(int32(2))%32)
	if v198 != 0 {
		goto L42
	} else {
		goto L47
	}
L47:
	;
	goto L43
L48:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v204)+8)) = v208
	v211 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
	*(*float64)(unsafe.Add(mBase, uint32(l2)+8)) = base.F64_add(v208, v211)
	v165 = v165 + int32(1)
	goto L37
}
func F_get_returning_clause(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_appendContextKeyword(m, l1, int32(_a_F_get_returning_clause_0), int32(-8), int32(8), int32(1))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	m.G0 = v9 + int32(32)
	return
L4:
	;
	return
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v19 == int32(0) {
		v59 = v3
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	if v60 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L7:
	;
	v22 = int32(_a_F_get_returning_clause_1)
	v25 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_get_returning_clause[0])))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v26 == int32(0) {
		v45 = v25
		v46 = v26
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if v46-v45 == int32(0) {
		v59 = v3
		goto L6
	} else {
		goto L16
	}
L9:
	;
	goto L8
L10:
	;
	if v25 != v26 {
		v45 = v25
		v46 = v26
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v30 = v19
	v31 = v22
	goto L12
L12:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+1)))
	if v35 == int32(0) {
		v45 = v34
		v46 = v35
		goto L9
	} else {
		goto L14
	}
L13:
	;
	v45 = v34
	v46 = v35
	goto L9
L14:
	;
	v38 = int32(1)
	if v34 == v35 {
		v30 = v30 + v38
		v31 = v31 + v38
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	v50 = F_quote_identifier(m, v19)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v50
	F_appendStringInfo(m, v12, int32(_a_F_get_returning_clause_2), v9+int32(16))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	v59 = int32(1)
	goto L6
L19:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	F_get_target_list(m, v104, l1)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L4
	} else {
		goto L39
	}
L20:
	;
	F_appendStringInfoChar(m, v12, int32(41))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L4
	} else {
		goto L38
	}
L21:
	;
	if v59 == int32(0) {
		goto L19
	} else {
		goto L37
	}
L22:
	;
	v63 = int32(_a_F_get_returning_clause_3)
	v66 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_get_returning_clause[1])))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
	if v67 == int32(0) {
		v86 = v66
		v87 = v67
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if v87-v86 == int32(0) {
		goto L21
	} else {
		goto L31
	}
L24:
	;
	goto L23
L25:
	;
	if v66 != v67 {
		v86 = v66
		v87 = v67
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v71 = v60
	v72 = v63
	goto L27
L27:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+1)))
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+1)))
	if v76 == int32(0) {
		v86 = v75
		v87 = v76
		goto L24
	} else {
		goto L29
	}
L28:
	;
	v86 = v75
	v87 = v76
	goto L24
L29:
	;
	v79 = int32(1)
	if v75 == v76 {
		v71 = v71 + v79
		v72 = v72 + v79
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v91 = F_quote_identifier(m, v60)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v91
	if v59 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v96 = int32(_a_F_get_returning_clause_4)
	goto L35
L34:
	;
	v96 = int32(_a_F_get_returning_clause_5)
	goto L35
L35:
	;
	F_appendStringInfo(m, v12, v96, v9)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	goto L20
L37:
	;
	goto L20
L38:
	;
	goto L19
L39:
	;
	goto L3
}
func F_get_rolespec_tuple(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	v4 = m.G0
	v6 = v4 - int32(80)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v8 {
	case 0:
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v11 = F_SearchSysCache1(m, int32(10), v10)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			if v11 != 0 {
				v111 = v11
				m.G0 = v6 + int32(80)
				return v111
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(67137668))
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v22
						F_errmsg(m, int32(_a_F_get_rolespec_tuple_0), v6+int32(16))
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_get_rolespec_tuple_1), int32(_a_F_get_rolespec_tuple_2), int32(_a_F_get_rolespec_tuple_3))
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		}
	case 1, 2:
		v36 = *(*int32)(unsafe.Add(mBase, _c_F_get_rolespec_tuple[0]))
		v37 = F_SearchSysCache1(m, int32(11), v36)
		mBase = m.M
		v38 = m.ExcPending
		if v38 != 0 {
			return int32(0)
		} else {
			if v37 != 0 {
				v111 = v37
				m.G0 = v6 + int32(80)
				return v111
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, _c_F_get_rolespec_tuple[0]))
					*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = v44
					F_errmsg_internal(m, int32(_a_F_get_rolespec_tuple_4), v6+int32(32))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_get_rolespec_tuple_1), int32(_a_F_get_rolespec_tuple_5), int32(_a_F_get_rolespec_tuple_3))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	case 3:
		v58 = *(*int32)(unsafe.Add(mBase, _c_F_get_rolespec_tuple[1]))
		v59 = F_SearchSysCache1(m, int32(11), v58)
		mBase = m.M
		v60 = m.ExcPending
		if v60 != 0 {
			return int32(0)
		} else {
			if v59 != 0 {
				v111 = v59
				m.G0 = v6 + int32(80)
				return v111
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return int32(0)
				} else {
					v66 = *(*int32)(unsafe.Add(mBase, _c_F_get_rolespec_tuple[1]))
					*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = v66
					F_errmsg_internal(m, int32(_a_F_get_rolespec_tuple_4), v6+int32(48))
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_get_rolespec_tuple_1), int32(_a_F_get_rolespec_tuple_6), int32(_a_F_get_rolespec_tuple_3))
						mBase = m.M
						v77 = m.ExcPending
						if v77 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	case 4:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v81 = m.ExcPending
		if v81 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(67137668))
			mBase = m.M
			v84 = m.ExcPending
			if v84 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6)+64)) = int32(_a_F_get_rolespec_tuple_7)
				F_errmsg(m, int32(_a_F_get_rolespec_tuple_0), v6-int32(-64))
				mBase = m.M
				v91 = m.ExcPending
				if v91 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_get_rolespec_tuple_1), int32(_a_F_get_rolespec_tuple_8), int32(_a_F_get_rolespec_tuple_3))
					mBase = m.M
					v96 = m.ExcPending
					if v96 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v100 = m.ExcPending
		if v100 != 0 {
			return int32(0)
		} else {
			v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = v101
			F_errmsg_internal(m, int32(_a_F_get_rolespec_tuple_9), v6)
			mBase = m.M
			v105 = m.ExcPending
			if v105 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_get_rolespec_tuple_1), int32(_a_F_get_rolespec_tuple_10), int32(_a_F_get_rolespec_tuple_3))
				mBase = m.M
				v110 = m.ExcPending
				if v110 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_get_sortgroupclause_tle(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	if l1 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L10
	} else {
		goto L11
	}
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v8 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v15 = int32(0)
	goto L4
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v12+v15<<(uint(int32(2))%32))))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	if v11 != v23 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	return v22
L6:
	;
	v26 = v15 + int32(1)
	if v26 != v8 {
		v15 = v26
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	goto L5
L9:
	;
	goto L1
L10:
	;
	return int32(0)
L11:
	;
	F_errmsg_internal(m, int32(_a_F_get_sortgroupclause_tle_0), int32(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	F_errfinish(m, int32(_a_F_get_sortgroupclause_tle_1), int32(366), int32(_a_F_get_sortgroupclause_tle_2))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_str_from_var(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v11 = int32(1)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v16 = v12<<(uint(int32(2))%32) + int32(4)
	if v16 <= v11 {
		v19 = v11
	} else {
		v19 = v16
	}
	v23 = F_palloc(m, v10+v19+int32(6))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		return int32(0)
	} else {
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v27 == int32(_a_F_get_str_from_var_0) {
			v30 = int32(45)
			*(*uint8)(unsafe.Add(mBase, uint32(v23))) = uint8(v30)
			v34 = v23 + int32(1)
		} else {
			v34 = v23
		}
		v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if int32(0) <= v35 {
			v40 = v34
			v42 = int32(0)
			for {
				v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				if v42 < v49 {
					v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v55 = int32(*(*int16)(unsafe.Add(mBase, uint32(v51+v42<<(uint(int32(1))%32)))))
					v56 = v55
				} else {
					v56 = int32(0)
				}
				v57 = int32(1000)
				v58 = base.I32_div_s(v56, v57)
				v63 = v58&int32(_a_F_get_str_from_var_1)*int32(_a_F_get_str_from_var_2) + v56
				v66 = int32(0)
				if base.B2i32(v56 < v57)&base.B2i32(v42 <= v66) == v66 {
					v72 = v58 + int32(48)
					*(*uint8)(unsafe.Add(mBase, uint32(v40))) = uint8(v72)
					v78 = base.I32_div_s(base.I32_extend16_s(v63), int32(100))
					v94 = v40 + int32(1)
					v95 = v78&int32(_a_F_get_str_from_var_1)*int32(_a_F_get_str_from_var_3) + v63
					v96 = v78
					v99 = v96 + int32(48)
					*(*uint8)(unsafe.Add(mBase, uint32(v94))) = uint8(v99)
					v105 = base.I32_div_s(base.I32_extend16_s(v95), int32(10))
					v121 = v94 + int32(1)
					v122 = v105&int32(_a_F_get_str_from_var_1)*int32(246) + v95
					v124 = v105
					v127 = v124 + int32(48)
					*(*uint8)(unsafe.Add(mBase, uint32(v121))) = uint8(v127)
					v131 = v121 + int32(1)
					v132 = v122
				} else {
					v84 = base.I32_extend16_s(v63)
					v85 = int32(100)
					v86 = base.I32_div_s(v84, v85)
					v91 = v86&int32(_a_F_get_str_from_var_1)*int32(_a_F_get_str_from_var_3) + v63
					if v84 < v85 {
						v111 = base.I32_extend16_s(v91)
						v112 = int32(10)
						v113 = base.I32_div_s(v111, v112)
						v118 = v113&int32(_a_F_get_str_from_var_1)*int32(246) + v91
						if v111 < v112 {
							v131 = v40
							v132 = v118
						} else {
							v121 = v40
							v122 = v118
							v124 = v113
							v127 = v124 + int32(48)
							*(*uint8)(unsafe.Add(mBase, uint32(v121))) = uint8(v127)
							v131 = v121 + int32(1)
							v132 = v122
						}
					} else {
						v94 = v40
						v95 = v91
						v96 = v86
						v99 = v96 + int32(48)
						*(*uint8)(unsafe.Add(mBase, uint32(v94))) = uint8(v99)
						v105 = base.I32_div_s(base.I32_extend16_s(v95), int32(10))
						v121 = v94 + int32(1)
						v122 = v105&int32(_a_F_get_str_from_var_1)*int32(246) + v95
						v124 = v105
						v127 = v124 + int32(48)
						*(*uint8)(unsafe.Add(mBase, uint32(v121))) = uint8(v127)
						v131 = v121 + int32(1)
						v132 = v122
					}
				}
				v137 = v132 + int32(48)
				*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v137)
				v139 = int32(1)
				v140 = v131 + v139
				v142 = v42 + v139
				v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if v142 <= v143 {
					v40 = v140
					v42 = v142
					continue
				} else {
					break
				}
				break
			}
			v152 = v140
			v154 = v142
		} else {
			v145 = int32(48)
			*(*uint8)(unsafe.Add(mBase, uint32(v34))) = uint8(v145)
			v147 = int32(1)
			v152 = v34 + v147
			v154 = v35 + v147
		}
		if int32(0) < v10 {
			v162 = int32(46)
			*(*uint8)(unsafe.Add(mBase, uint32(v152))) = uint8(v162)
			v166 = v152 + int32(1)
			v168 = v166
			v170 = v154
			v172 = int32(0)
			for {
				v176 = int32(0)
				if v170 < v176 {
					v187 = v176
				} else {
					v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					if v180 <= v170 {
						v187 = int32(0)
					} else {
						v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v186 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v182+v170<<(uint(int32(1))%32)))))
						v187 = v186
					}
				}
				v190 = base.I32_div_s(base.I32_extend16_s(v187), int32(1000))
				v191 = int32(48)
				v192 = v190 + v191
				*(*uint8)(unsafe.Add(mBase, uint32(v168))) = uint8(v192)
				v194 = int32(_a_F_get_str_from_var_1)
				v198 = v190&v194*int32(_a_F_get_str_from_var_2) + v187
				v201 = base.I32_div_s(base.I32_extend16_s(v198), int32(100))
				v203 = v201 + v191
				*(*uint8)(unsafe.Add(mBase, uint32(v168)+1)) = uint8(v203)
				v209 = v201&v194*int32(_a_F_get_str_from_var_3) + v198
				v212 = base.I32_div_s(base.I32_extend16_s(v209), int32(10))
				v214 = v212 + v191
				*(*uint8)(unsafe.Add(mBase, uint32(v168)+2)) = uint8(v214)
				v220 = v212*int32(246) + v209 + v191
				*(*uint8)(unsafe.Add(mBase, uint32(v168)+3)) = uint8(v220)
				v224 = int32(4)
				v227 = v172 + v224
				if v227 < v10 {
					v168 = v168 + v224
					v170 = v170 + int32(1)
					v172 = v227
					continue
				} else {
					break
				}
				break
			}
			v239 = v10 + v166
		} else {
			v239 = v152
		}
		v240 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v239))) = uint8(v240)
		return v23
	}
}
func F_get_user_default_acl(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	v4 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_get_user_default_acl[0]))
	if v15 == v4 {
		v205 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v12 + int32(16)
	return v205
L2:
	;
	switch l0 - int32(19) {
	case 0:
		goto L7
	default:
		v205 = v4
		goto L1
	case 3:
		goto L4
	case 17:
		goto L5
	case 18:
		goto L8
	case 22:
		v26 = int32(114)
		goto L3
	case 30:
		goto L6
	}
L3:
	;
	v29 = F_SearchSysCache3(m, int32(22), l1, int32(0), v26)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	v26 = int32(76)
	goto L3
L5:
	;
	v26 = int32(110)
	goto L3
L6:
	;
	v26 = int32(84)
	goto L3
L7:
	;
	v26 = int32(102)
	goto L3
L8:
	;
	v26 = int32(83)
	goto L3
L9:
	;
	return int32(0)
L10:
	;
	if v29 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v37 = F_SysCacheGetAttr(m, int32(22), v29, int32(5), v12+int32(14))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L9
	} else {
		goto L14
	}
L12:
	;
	v47 = v4
	goto L13
L13:
	;
	v50 = F_SearchSysCache3(m, int32(22), l1, l2, v26)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L9
	} else {
		goto L20
	}
L14:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+14)))
	if v39 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v42 = F_pg_detoast_datum_copy(m, v37)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L9
	} else {
		goto L18
	}
L16:
	;
	v44 = v4
	goto L17
L17:
	;
	F_ReleaseCatCache(m, v29)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L9
	} else {
		goto L19
	}
L18:
	;
	v44 = v42
	goto L17
L19:
	;
	v47 = v44
	goto L13
L20:
	;
	if v50 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v56 = F_SysCacheGetAttr(m, int32(22), v50, int32(5), v12+int32(15))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L9
	} else {
		goto L24
	}
L22:
	;
	v66 = v4
	goto L23
L23:
	;
	v68 = int32(0)
	if v66|v47 == v68 {
		v205 = v68
		goto L1
	} else {
		goto L30
	}
L24:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
	if v58 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v61 = F_pg_detoast_datum_copy(m, v56)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L9
	} else {
		goto L28
	}
L26:
	;
	v63 = v4
	goto L27
L27:
	;
	F_ReleaseCatCache(m, v50)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L9
	} else {
		goto L29
	}
L28:
	;
	v63 = v61
	goto L27
L29:
	;
	v66 = v63
	goto L23
L30:
	;
	v72 = F_acldefault(m, l0, l1)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L9
	} else {
		goto L32
	}
L31:
	;
	F_aclitemsort(m, v149)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L9
	} else {
		goto L63
	}
L32:
	;
	if v47 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v74 = v47
	goto L35
L34:
	;
	v74 = v72
	goto L35
L35:
	;
	v75 = int32(0)
	if v74 != 0 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	v149 = v131
	goto L31
L37:
	;
	if v66 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L38:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v74)+16))
	if v77 != 0 {
		goto L37
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	if v66 == int32(0) {
		v131 = v75
		goto L36
	} else {
		goto L42
	}
L41:
	;
	goto L40
L42:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v66)+16))
	if v80 == int32(0) {
		v131 = v75
		goto L36
	} else {
		goto L43
	}
L43:
	;
	v83 = F_aclcopy(m, v66)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L9
	} else {
		goto L44
	}
L44:
	;
	v149 = v83
	goto L31
L45:
	;
	v87 = F_aclcopy(m, v74)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L9
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v66)+16))
	if v89 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v149 = v87
	goto L31
L49:
	;
	v92 = F_aclcopy(m, v74)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L9
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v94 = F_aclcopy(m, v74)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L9
	} else {
		goto L53
	}
L52:
	;
	v149 = v92
	goto L31
L53:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
	if v96 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	v106 = (v99<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L56
L55:
	;
	v106 = v96
	goto L56
L56:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v66)+16))
	if v107 <= int32(0) {
		v149 = v94
		goto L31
	} else {
		goto L57
	}
L57:
	;
	v113 = v94
	v114 = v106 + v66
	v117 = v75
	goto L58
L58:
	;
	v122 = F_aclupdate(m, v113, v114, int32(1), l1, int32(0))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L9
	} else {
		goto L60
	}
L59:
	;
	v131 = v122
	goto L36
L60:
	;
	F_pfree(m, v113)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L9
	} else {
		goto L61
	}
L61:
	;
	v129 = v117 + int32(1)
	if v129 != v107 {
		v113 = v122
		v114 = v114 + int32(16)
		v117 = v129
		goto L58
	} else {
		goto L62
	}
L62:
	;
	goto L59
L63:
	;
	F_aclitemsort(m, v72)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L9
	} else {
		goto L64
	}
L64:
	;
	v154 = int32(0)
	if v149 != 0 {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	if v199 != 0 {
		goto L83
	} else {
		goto L84
	}
L66:
	;
	if v72 == int32(0) {
		v195 = v154
		goto L74
	} else {
		goto L75
	}
L67:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v149)+16))
	if v157 != 0 {
		goto L66
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	if v72 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	goto L69
L71:
	;
	v199 = int32(1)
	goto L65
L72:
	;
	goto L73
L73:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v72)+16))
	v199 = base.B2i32(v162 == int32(0))
	goto L65
L74:
	;
	v199 = v195
	goto L65
L75:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v72)+16))
	if v157 != v167 {
		v195 = v154
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v149)+8))
	if v169 != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v177 = v169
	goto L79
L78:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v149)+4))
	v177 = (v170<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L79
L79:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v72)+8))
	if v179 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v187 = v179
	goto L82
L81:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	v187 = (v180<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L82
L82:
	;
	v191 = F_memcmp(m, v177+v149, v187+v72, v157<<(uint(int32(4))%32))
	mBase = m.M
	v195 = base.B2i32(v191 == int32(0))
	goto L74
L83:
	;
	v200 = v154
	goto L85
L84:
	;
	v200 = v149
	goto L85
L85:
	;
	v205 = v200
	goto L1
}
func F_get_with_clause(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v118 int32
	_ = v118
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v236 int32
	_ = v236
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v255 int32
	_ = v255
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v320 int32
	_ = v320
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	v3 = int32(0)
	v10 = m.G0
	v12 = v10 + int32(-64)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v14 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v12 - int32(-64)
	return
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	if v18&int32(2) != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	if v407&int32(2) != 0 {
		goto L110
	} else {
		goto L111
	}
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v21 + int32(8)
	F_appendStringInfoChar(m, v17, int32(32))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v31 = v14
	goto L6
L6:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v32 <= int32(0) {
		goto L3
	} else {
		goto L10
	}
L7:
	;
	return
L8:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v28 == int32(0) {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v31 = v28
	goto L6
L10:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+41)))
	if v37 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v38 = int32(_a_F_get_with_clause_0)
	goto L13
L12:
	;
	v38 = int32(_a_F_get_with_clause_1)
	goto L13
L13:
	;
	v41 = v38
	v47 = v3
	goto L14
L14:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v48+v47<<(uint(int32(2))%32))))
	F_appendStringInfoString(m, v17, v41)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L7
	} else {
		goto L16
	}
L15:
	;
	goto L3
L16:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v56 = F_quote_identifier(m, v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	F_appendStringInfoString(m, v17, v56)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L7
	} else {
		goto L18
	}
L18:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	if v60 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	F_appendStringInfoChar(m, v17, int32(40))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L7
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	F_appendStringInfoString(m, v17, int32(_a_F_get_with_clause_2))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L7
	} else {
		goto L36
	}
L22:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	if v64 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	F_appendStringInfoChar(m, v17, int32(41))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L7
	} else {
		goto L35
	}
L24:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	if v67 <= int32(0) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v64)+12))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	v73 = F_quote_identifier(m, v72)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L7
	} else {
		goto L26
	}
L26:
	;
	F_appendStringInfoString(m, v17, v73)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L7
	} else {
		goto L27
	}
L27:
	;
	v77 = int32(1)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	if v78 <= v77 {
		goto L23
	} else {
		goto L28
	}
L28:
	;
	v83 = v77
	goto L29
L29:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v64)+12))
	F_appendStringInfoString(m, v17, int32(_a_F_get_with_clause_3))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L7
	} else {
		goto L31
	}
L30:
	;
	goto L23
L31:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v90+v83<<(uint(int32(2))%32))))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
	v99 = F_quote_identifier(m, v98)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L7
	} else {
		goto L32
	}
L32:
	;
	F_appendStringInfoString(m, v17, v99)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L7
	} else {
		goto L33
	}
L33:
	;
	v104 = v83 + int32(1)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	if v104 < v105 {
		v83 = v104
		goto L29
	} else {
		goto L34
	}
L34:
	;
	goto L30
L35:
	;
	goto L21
L36:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
	switch v132 - int32(1) {
	case 0:
		v136 = int32(_a_F_get_with_clause_4)
		goto L38
	case 1:
		goto L39
	default:
		goto L37
	}
L37:
	;
	F_appendStringInfoChar(m, v17, int32(40))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L7
	} else {
		goto L41
	}
L38:
	;
	F_appendStringInfoString(m, v17, v136)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L7
	} else {
		goto L40
	}
L39:
	;
	v136 = int32(_a_F_get_with_clause_5)
	goto L38
L40:
	;
	goto L37
L41:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v143&int32(2) != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v147 = int32(0)
	F_appendContextKeyword(m, l1, int32(_a_F_get_with_clause_6), v147, v147, v147)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L7
	} else {
		goto L45
	}
L43:
	;
	v153 = v143
	goto L44
L44:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	F_get_query_def(m, v154, v17, v155, int32(0), int32(1), v153, v158, v159)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L7
	} else {
		goto L46
	}
L45:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v153 = v152
	goto L44
L46:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	if v162&int32(2) != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v166 = int32(0)
	F_appendContextKeyword(m, l1, int32(_a_F_get_with_clause_6), v166, v166, v166)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L7
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	F_appendStringInfoChar(m, v17, int32(41))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L7
	} else {
		goto L51
	}
L50:
	;
	goto L49
L51:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v52)+20))
	if v174 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+8)))
	if v177 != 0 {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	goto L54
L54:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v52)+24))
	if v265 != 0 {
		goto L75
	} else {
		goto L76
	}
L55:
	;
	v178 = int32(_a_F_get_with_clause_7)
	goto L57
L56:
	;
	v178 = int32(_a_F_get_with_clause_8)
	goto L57
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v178
	F_appendStringInfo(m, v17, int32(_a_F_get_with_clause_9), v10+int32(-16))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L7
	} else {
		goto L58
	}
L58:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v52)+20))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v185)+4))
	if v186 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v186)+4))
	if v187 <= int32(0) {
		goto L62
	} else {
		goto L63
	}
L60:
	;
	v246 = v185
	goto L61
L61:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v246)+12))
	v248 = F_quote_identifier(m, v247)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L7
	} else {
		goto L73
	}
L62:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v52)+20))
	v246 = v236
	goto L61
L63:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v186)+12))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v191)+4))
	v193 = F_quote_identifier(m, v192)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L7
	} else {
		goto L64
	}
L64:
	;
	F_appendStringInfoString(m, v17, v193)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L7
	} else {
		goto L65
	}
L65:
	;
	v197 = int32(1)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v186)+4))
	if v198 <= v197 {
		goto L62
	} else {
		goto L66
	}
L66:
	;
	v203 = v197
	goto L67
L67:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v186)+12))
	F_appendStringInfoString(m, v17, int32(_a_F_get_with_clause_3))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L7
	} else {
		goto L69
	}
L68:
	;
	goto L62
L69:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v210+v203<<(uint(int32(2))%32))))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v217)+4))
	v219 = F_quote_identifier(m, v218)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L7
	} else {
		goto L70
	}
L70:
	;
	F_appendStringInfoString(m, v17, v219)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L7
	} else {
		goto L71
	}
L71:
	;
	v224 = v203 + int32(1)
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v186)+4))
	if v224 < v225 {
		v203 = v224
		goto L67
	} else {
		goto L72
	}
L72:
	;
	goto L68
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v248
	F_appendStringInfo(m, v17, int32(_a_F_get_with_clause_10), v10+int32(-32))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L7
	} else {
		goto L74
	}
L74:
	;
	goto L54
L75:
	;
	F_appendStringInfoString(m, v17, int32(_a_F_get_with_clause_11))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L7
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v395 = v47 + int32(1)
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v395 < v396 {
		v41 = int32(_a_F_get_with_clause_3)
		v47 = v395
		goto L14
	} else {
		goto L109
	}
L78:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v52)+24))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v269)+4))
	if v270 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v270)+4))
	if v271 <= int32(0) {
		goto L82
	} else {
		goto L83
	}
L80:
	;
	v330 = v269
	goto L81
L81:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v330)+8))
	v332 = F_quote_identifier(m, v331)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L7
	} else {
		goto L93
	}
L82:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v52)+24))
	v330 = v320
	goto L81
L83:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v274)))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v275)+4))
	v277 = F_quote_identifier(m, v276)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L7
	} else {
		goto L84
	}
L84:
	;
	F_appendStringInfoString(m, v17, v277)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L7
	} else {
		goto L85
	}
L85:
	;
	v281 = int32(1)
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v270)+4))
	if v282 <= v281 {
		goto L82
	} else {
		goto L86
	}
L86:
	;
	v287 = v281
	goto L87
L87:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
	F_appendStringInfoString(m, v17, int32(_a_F_get_with_clause_3))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L7
	} else {
		goto L89
	}
L88:
	;
	goto L82
L89:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v294+v287<<(uint(int32(2))%32))))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v301)+4))
	v303 = F_quote_identifier(m, v302)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L7
	} else {
		goto L90
	}
L90:
	;
	F_appendStringInfoString(m, v17, v303)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L7
	} else {
		goto L91
	}
L91:
	;
	v308 = v287 + int32(1)
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v270)+4))
	if v308 < v309 {
		v287 = v308
		goto L87
	} else {
		goto L92
	}
L92:
	;
	goto L88
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v332
	F_appendStringInfo(m, v17, int32(_a_F_get_with_clause_10), v10+int32(-48))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L7
	} else {
		goto L94
	}
L94:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v52)+24))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v340)+12))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v341)+4))
	if v342 != int32(16) {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v375)+20))
	v378 = F_quote_identifier(m, v377)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L7
	} else {
		goto L107
	}
L96:
	;
	F_appendStringInfoString(m, v17, int32(_a_F_get_with_clause_12))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L7
	} else {
		goto L103
	}
L97:
	;
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341)+24)))
	if v345 != 0 {
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v341)+20))
	if v346 == int32(0) {
		goto L96
	} else {
		goto L99
	}
L99:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v340)+16))
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v349)+4))
	if v350 != int32(16) {
		goto L96
	} else {
		goto L100
	}
L100:
	;
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349)+24)))
	if v353 != 0 {
		goto L96
	} else {
		goto L101
	}
L101:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v349)+20))
	if v354 == int32(0) {
		v375 = v340
		goto L95
	} else {
		goto L102
	}
L102:
	;
	goto L96
L103:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v52)+24))
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v361)+12))
	F_get_rule_expr(m, v362, l1, int32(0))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L7
	} else {
		goto L104
	}
L104:
	;
	F_appendStringInfoString(m, v17, int32(_a_F_get_with_clause_13))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L7
	} else {
		goto L105
	}
L105:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v52)+24))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v369)+16))
	F_get_rule_expr(m, v370, l1, int32(0))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L7
	} else {
		goto L106
	}
L106:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v52)+24))
	v375 = v374
	goto L95
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v378
	F_appendStringInfo(m, v17, int32(_a_F_get_with_clause_14), v12)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L7
	} else {
		goto L108
	}
L108:
	;
	goto L77
L109:
	;
	goto L15
L110:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v410 - int32(8)
	v415 = int32(0)
	F_appendContextKeyword(m, l1, int32(_a_F_get_with_clause_6), v415, v415, v415)
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L7
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	F_appendStringInfoChar(m, v17, int32(32))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L7
	} else {
		goto L114
	}
L113:
	;
	goto L1
L114:
	;
	goto L1
}
func F_getmissingattr(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v128 int32
	_ = v128
	v7 = m.G0
	v9 = v7 + int32(-64)
	m.G0 = v9
	v11 = int32(1)
	v12 = l1 - v11
	v15 = l0 + v12<<(uint(int32(4))%32)
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+28)))
	if v16 != v11 {
		v120 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v120)
		v128 = int32(0)
		m.G0 = v9 - int32(-64)
		return v128
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
		v23 = v20 + v12<<(uint(int32(3))%32)
		v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
		if v24 != int32(1) {
			v120 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v120)
			v128 = int32(0)
			m.G0 = v9 - int32(-64)
			return v128
		} else {
			v27 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v27)
			v30 = v15 + int32(20)
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+6)))
			if v31 != 0 {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
				v128 = v32
				m.G0 = v9 - int32(-64)
				return v128
			} else {
				v34 = *(*int32)(unsafe.Add(mBase, _c_F_getmissingattr[0]))
				if v34 == int32(0) {
					*(*int64)(unsafe.Add(mBase, uint32(v9)+32)) = int64(34359738376)
					*(*int32)(unsafe.Add(mBase, uint32(v9)+44)) = int32(23)
					*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = int32(24)
					v44 = *(*int32)(unsafe.Add(mBase, _c_F_getmissingattr[1]))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = v44
					v52 = F_hash_create(m, int32(_a_F_getmissingattr_0), int32(32), v7+int32(-48), int32(1224))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_getmissingattr[0])) = v52
						v57 = v52
						v58 = int32(*(*int16)(unsafe.Add(mBase, uint32(v30)+4)))
						if int32(0) < v58 {
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
							v89 = v58
							v90 = v61
						} else {
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
							v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
							if v63 == int32(1) {
								v66 = int32(6)
								v68 = int32(18)
								v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+1)))
								if v70 == v68 {
									v73 = v68
								} else {
									v73 = int32(2)
								}
								if v70&int32(254) == int32(2) {
									v78 = v66
								} else {
									v78 = v73
								}
								if v70 == int32(1) {
									v81 = v66
								} else {
									v81 = v78
								}
								v89 = v81
								v90 = v62
							} else {
								if v63&int32(1) != 0 {
									v89 = int32(base.Ui32(v63) >> (uint(int32(1)) % 32))
									v90 = v62
								} else {
									v86 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
									v89 = int32(base.Ui32(v86) >> (uint(int32(2)) % 32))
									v90 = v62
								}
							}
						}
						*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v90
						*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v89
						v98 = F_hash_search(m, v57, v7+int32(-48), int32(1), v7+int32(-49))
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return int32(0)
						} else {
							v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
							if v100 == int32(0) {
								v103 = int32(_a_F_getmissingattr_1)
								v104 = *(*int32)(unsafe.Add(mBase, _c_F_getmissingattr[2]))
								v107 = *(*int32)(unsafe.Add(mBase, _c_F_getmissingattr[1]))
								*(*int32)(unsafe.Add(mBase, _c_F_getmissingattr[2])) = v107
								v109 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
								v111 = int32(*(*int16)(unsafe.Add(mBase, uint32(v30)+4)))
								v112 = F_datumCopy(m, v109, int32(0), v111)
								mBase = m.M
								v113 = m.ExcPending
								if v113 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v98)+4)) = v112
									*(*int32)(unsafe.Add(mBase, _c_F_getmissingattr[2])) = v104
									v118 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
									v128 = v118
									m.G0 = v9 - int32(-64)
									return v128
								}
							} else {
								v118 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
								v128 = v118
								m.G0 = v9 - int32(-64)
								return v128
							}
						}
					}
				} else {
					v57 = v34
					v58 = int32(*(*int16)(unsafe.Add(mBase, uint32(v30)+4)))
					if int32(0) < v58 {
						v61 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
						v89 = v58
						v90 = v61
					} else {
						v62 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
						v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
						if v63 == int32(1) {
							v66 = int32(6)
							v68 = int32(18)
							v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+1)))
							if v70 == v68 {
								v73 = v68
							} else {
								v73 = int32(2)
							}
							if v70&int32(254) == int32(2) {
								v78 = v66
							} else {
								v78 = v73
							}
							if v70 == int32(1) {
								v81 = v66
							} else {
								v81 = v78
							}
							v89 = v81
							v90 = v62
						} else {
							if v63&int32(1) != 0 {
								v89 = int32(base.Ui32(v63) >> (uint(int32(1)) % 32))
								v90 = v62
							} else {
								v86 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
								v89 = int32(base.Ui32(v86) >> (uint(int32(2)) % 32))
								v90 = v62
							}
						}
					}
					*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v90
					*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v89
					v98 = F_hash_search(m, v57, v7+int32(-48), int32(1), v7+int32(-49))
					mBase = m.M
					v99 = m.ExcPending
					if v99 != 0 {
						return int32(0)
					} else {
						v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
						if v100 == int32(0) {
							v103 = int32(_a_F_getmissingattr_1)
							v104 = *(*int32)(unsafe.Add(mBase, _c_F_getmissingattr[2]))
							v107 = *(*int32)(unsafe.Add(mBase, _c_F_getmissingattr[1]))
							*(*int32)(unsafe.Add(mBase, _c_F_getmissingattr[2])) = v107
							v109 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
							v111 = int32(*(*int16)(unsafe.Add(mBase, uint32(v30)+4)))
							v112 = F_datumCopy(m, v109, int32(0), v111)
							mBase = m.M
							v113 = m.ExcPending
							if v113 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v98)+4)) = v112
								*(*int32)(unsafe.Add(mBase, _c_F_getmissingattr[2])) = v104
								v118 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
								v128 = v118
								m.G0 = v9 - int32(-64)
								return v128
							}
						} else {
							v118 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
							v128 = v118
							m.G0 = v9 - int32(-64)
							return v128
						}
					}
				}
			}
		}
	}
}
func F_getopt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v122 int32
	_ = v122
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	var v274 int32
	_ = v274
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v378 int32
	_ = v378
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v406 int32
	_ = v406
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_getopt[0]))
	if v16 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v32 = int32(-1)
	if l0 <= v31 {
		v406 = v32
		goto L6
	} else {
		goto L7
	}
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_getopt[1]))
	if v18 == int32(0) {
		v31 = v16
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v21 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_getopt[0])) = v21
	v26 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_getopt[2])) = v26
	*(*int32)(unsafe.Add(mBase, _c_F_getopt[1])) = v26
	v31 = v21
	goto L1
L5:
	;
	goto L4
L6:
	;
	m.G0 = v13 + int32(16)
	return v406
L7:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1+v31<<(uint(int32(2))%32))))
	if v37 == int32(0) {
		v406 = v32
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
	if v40 != int32(45) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v43 != int32(45) {
		v406 = v32
		goto L6
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+1)))
	if v53 == int32(0) {
		v406 = v32
		goto L6
	} else {
		goto L13
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_getopt[3])) = v37
	v48 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_getopt[0])) = v31 + v48
	v406 = v48
	goto L6
L13:
	;
	if v53 != int32(45) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v64 = v13 + int32(12)
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_getopt[2]))
	if v66 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+2)))
	if v58 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_getopt[0])) = v31 + int32(1)
	v406 = v32
	goto L6
L17:
	;
	v71 = v66
	goto L19
L18:
	;
	v68 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_getopt[2])) = v68
	v71 = v68
	goto L19
L19:
	;
	v72 = v71 + v37
	if v72 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	if v172 < int32(0) {
		goto L49
	} else {
		goto L50
	}
L21:
	;
	v172 = int32(0)
	goto L20
L22:
	;
	goto L23
L23:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	v80 = base.I32_extend8_s(v79)
	if int32(0) <= v80 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v172 = v168
	goto L20
L25:
	;
	if v64 != 0 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _c_F_getopt[4]))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	if v88 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v79
	goto L30
L29:
	;
	goto L30
L30:
	;
	v172 = base.B2i32(v80 != int32(0))
	goto L20
L31:
	;
	if v64 == int32(0) {
		v168 = int32(1)
		goto L24
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v99 = v79 - int32(194)
	if base.Ui32(int32(50)) < base.Ui32(v99) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v80 & int32(_a_F_getopt_0)
	v172 = int32(1)
	goto L20
L35:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_getopt[5])) = int32(25)
	v168 = int32(-1)
	goto L24
L36:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+1)))
	v104 = int32(base.Ui32(v102) >> (uint(int32(3)) % 32))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v99<<(uint(int32(2))%32))+uint32(_c_F_getopt[6])))
	if base.Ui32(int32(7)) < base.Ui32(v104-int32(16)|(v104+v111>>(uint(int32(26))%32))) {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v122 = v102 - int32(128) | v111<<(uint(int32(6))%32)
	if int32(0) <= v122 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	if v64 == int32(0) {
		v168 = int32(2)
		goto L24
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+2)))
	v132 = v130 - int32(128)
	if base.Ui32(int32(63)) < base.Ui32(v132) {
		goto L35
	} else {
		goto L42
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v122
	v172 = int32(2)
	goto L20
L42:
	;
	v136 = v122 << (uint(int32(6)) % 32)
	v137 = v132 | v136
	if int32(0) <= v136 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	if v64 == int32(0) {
		v168 = int32(3)
		goto L24
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+3)))
	v147 = v145 - int32(128)
	if base.Ui32(int32(63)) < base.Ui32(v147) {
		goto L35
	} else {
		goto L47
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v137
	v172 = int32(3)
	goto L20
L47:
	;
	if v64 == int32(0) {
		v168 = int32(4)
		goto L24
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v147 | v137<<(uint(int32(6))%32)
	v172 = int32(4)
	goto L20
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = int32(_a_F_getopt_1)
	v178 = int32(1)
	goto L51
L50:
	;
	v178 = v172
	goto L51
L51:
	;
	v180 = *(*int32)(unsafe.Add(mBase, _c_F_getopt[0]))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l1+v180<<(uint(int32(2))%32))))
	v185 = int32(_a_F_getopt_2)
	v187 = *(*int32)(unsafe.Add(mBase, _c_F_getopt[2]))
	v188 = v187 + v178
	*(*int32)(unsafe.Add(mBase, _c_F_getopt[2])) = v188
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184+v188))))
	if v191 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_getopt[0])) = v180 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_getopt[2])) = int32(0)
	goto L54
L53:
	;
	goto L54
L54:
	;
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	switch v201 - int32(43) {
	case 0, 2:
		goto L56
	default:
		v206 = l2
		goto L55
	}
L55:
	;
	v207 = v184 + v187
	v208 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v208
	v214 = v208
	goto L57
L56:
	;
	v206 = l2 + int32(1)
	goto L55
L57:
	;
	v223 = v13 + int32(8)
	v224 = v206 + v214
	if v224 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L58:
	;
	if base.B2i32(v330 == v329)&base.B2i32(v329 != int32(58)) == int32(0) {
		goto L94
	} else {
		goto L95
	}
L59:
	;
	goto L58
L60:
	;
	if v324 <= int32(1) {
		goto L89
	} else {
		goto L90
	}
L61:
	;
	v324 = int32(0)
	goto L60
L62:
	;
	goto L63
L63:
	;
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224))))
	v232 = base.I32_extend8_s(v231)
	if int32(0) <= v232 {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v324 = v320
	goto L60
L65:
	;
	if v223 != 0 {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	goto L67
L67:
	;
	v239 = *(*int32)(unsafe.Add(mBase, _c_F_getopt[4]))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v239)))
	if v240 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v223))) = v231
	goto L70
L69:
	;
	goto L70
L70:
	;
	v324 = base.B2i32(v232 != int32(0))
	goto L60
L71:
	;
	if v223 == int32(0) {
		v320 = int32(1)
		goto L64
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v251 = v231 - int32(194)
	if base.Ui32(int32(50)) < base.Ui32(v251) {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v223))) = v232 & int32(_a_F_getopt_0)
	v324 = int32(1)
	goto L60
L75:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_getopt[5])) = int32(25)
	v320 = int32(-1)
	goto L64
L76:
	;
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224)+1)))
	v256 = int32(base.Ui32(v254) >> (uint(int32(3)) % 32))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v251<<(uint(int32(2))%32))+uint32(_c_F_getopt[6])))
	if base.Ui32(int32(7)) < base.Ui32(v256-int32(16)|(v256+v263>>(uint(int32(26))%32))) {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v274 = v254 - int32(128) | v263<<(uint(int32(6))%32)
	if int32(0) <= v274 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	if v223 == int32(0) {
		v320 = int32(2)
		goto L64
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224)+2)))
	v284 = v282 - int32(128)
	if base.Ui32(int32(63)) < base.Ui32(v284) {
		goto L75
	} else {
		goto L82
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v223))) = v274
	v324 = int32(2)
	goto L60
L82:
	;
	v288 = v274 << (uint(int32(6)) % 32)
	v289 = v284 | v288
	if int32(0) <= v288 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	if v223 == int32(0) {
		v320 = int32(3)
		goto L64
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224)+3)))
	v299 = v297 - int32(128)
	if base.Ui32(int32(63)) < base.Ui32(v299) {
		goto L75
	} else {
		goto L87
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v223))) = v289
	v324 = int32(3)
	goto L60
L87:
	;
	if v223 == int32(0) {
		v320 = int32(4)
		goto L64
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v223))) = v299 | v289<<(uint(int32(6))%32)
	v324 = int32(4)
	goto L60
L89:
	;
	v327 = int32(1)
	goto L91
L90:
	;
	v327 = v324
	goto L91
L91:
	;
	v328 = v327 + v214
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if v324 == int32(0) {
		goto L59
	} else {
		goto L92
	}
L92:
	;
	if v330 != v329 {
		v214 = v328
		goto L57
	} else {
		goto L93
	}
L93:
	;
	goto L59
L94:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_getopt[7])) = v329
	v342 = int32(63)
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206))))
	if v343 == int32(58) {
		v406 = v342
		goto L6
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	v356 = v206 + v328
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v356))))
	if v357 != int32(58) {
		v406 = v330
		goto L6
	} else {
		goto L101
	}
L97:
	;
	v347 = *(*int32)(unsafe.Add(mBase, _c_F_getopt[8]))
	if v347 == int32(0) {
		v406 = v342
		goto L6
	} else {
		goto L98
	}
L98:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F___getopt_msg(m, v350, int32(_a_F_getopt_3), v207, v178)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	return int32(0)
L100:
	;
	v406 = v342
	goto L6
L101:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_getopt[3])) = int32(0)
	v364 = *(*int32)(unsafe.Add(mBase, _c_F_getopt[0]))
	v366 = *(*int32)(unsafe.Add(mBase, _c_F_getopt[2]))
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v356)+1)))
	if v367 != int32(58) {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	if v384 <= l0 {
		v406 = v330
		goto L6
	} else {
		goto L106
	}
L103:
	;
	v372 = v364 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_getopt[0])) = v372
	v378 = *(*int32)(unsafe.Add(mBase, uint32(l1+v364<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, _c_F_getopt[3])) = v378 + v366
	*(*int32)(unsafe.Add(mBase, _c_F_getopt[2])) = int32(0)
	v384 = v372
	goto L102
L104:
	;
	if v366 != 0 {
		goto L103
	} else {
		goto L105
	}
L105:
	;
	v384 = v364
	goto L102
L106:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_getopt[7])) = v330
	v388 = int32(58)
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206))))
	if v389 == v388 {
		v406 = v388
		goto L6
	} else {
		goto L107
	}
L107:
	;
	v392 = int32(63)
	v394 = *(*int32)(unsafe.Add(mBase, _c_F_getopt[8]))
	if v394 == int32(0) {
		v406 = v392
		goto L6
	} else {
		goto L108
	}
L108:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F___getopt_msg(m, v397, int32(_a_F_getopt_4), v207, v178)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L99
	} else {
		goto L109
	}
L109:
	;
	v406 = v392
	goto L6
}
func F_getrusage(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v30 int32
	_ = v30
	var v32 int64
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int64
	_ = v38
	var v39 int32
	_ = v39
	v6 = m.G0
	v7 = int32(16)
	v8 = v6 - v7
	m.G0 = v8
	v11 = l0 + v7
	v17 = F__emscripten_memset_bulkmem(m, l0+int32(24), base.I32_extend8_s(int32(0)), int32(144))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = int32(4)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = int64(3)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = int32(2)
	*(*int64)(unsafe.Add(mBase, uint32(v11))) = int64(1)
	v30 = F__emscripten_memcpy_bulkmem(m, v8, v11, int32(16))
	mBase = m.M
	v32 = int64(*(*int32)(unsafe.Add(mBase, uint32(v8))))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v34 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v33
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v32
	v38 = int64(*(*int32)(unsafe.Add(mBase, uint32(v8)+8)))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v39
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v38
	m.G0 = v8 + int32(16)
	return
}
func F_getsubdfa(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v5+v6<<(uint(int32(2))%32))))
	if v10 != 0 {
		v40 = v10
		return v40
	} else {
		v11 = int32(0)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v18 = F_newdfa(m, l0, l1+int32(36), v14+int32(72), v11)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			if v18 == int32(0) {
				v40 = v11
			} else {
				v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
				if v24 == int32(98) {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v27
					v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+16)))
					*(*uint16)(unsafe.Add(mBase, uint32(v18)+64)) = uint16(v29)
					v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+18)))
					*(*uint16)(unsafe.Add(mBase, uint32(v18)+66)) = uint16(v31)
				} else {
				}
				v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v33+v34<<(uint(int32(2))%32)))) = v18
				v40 = v18
			}
			return v40
		}
	}
}
func F_gettoken_query_plain(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	v7 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l4))) = uint16(v7)
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v7)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v12
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v15&int32(3) == int32(0) {
		v39 = v15
		goto L6
	} else {
		goto L7
	}
L2:
	;
	v83 = v7
	goto L3
L3:
	;
	return v83
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v72
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v72 + v74
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v77 + int32(1)
	v83 = int32(2)
	goto L3
L5:
	;
	v72 = v64 - v15
	goto L4
L6:
	;
	v43 = v39
	goto L15
L7:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v23 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v72 = int32(0)
	goto L4
L9:
	;
	goto L10
L10:
	;
	v28 = v15
	goto L11
L11:
	;
	v32 = v28 + int32(1)
	if v32&int32(3) == int32(0) {
		v39 = v32
		goto L6
	} else {
		goto L13
	}
L12:
	;
	v64 = v32
	goto L5
L13:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	if v37 != 0 {
		v28 = v32
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	v52 = int32(-2139062144)
	if (int32(16843008)-v49|v49)&v52 == v52 {
		v43 = v43 + int32(4)
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v58 = v43
	goto L18
L17:
	;
	goto L16
L18:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	if v62 != 0 {
		v58 = v58 + int32(1)
		goto L18
	} else {
		goto L20
	}
L19:
	;
	v64 = v58
	goto L5
L20:
	;
	goto L19
}
func F_ginbuildempty(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int64
	_ = v13
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int64
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v72 int64
	_ = v72
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	v4 = m.G0
	v6 = v4 + int32(-64)
	m.G0 = v6
	*(*int64)(unsafe.Add(mBase, uint32(v6)+56)) = int64(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = v10
	*(*int32)(unsafe.Add(mBase, uint32(v6)+52)) = l0
	v13 = *(*int64)(unsafe.Add(mBase, uint32(v6)+52))
	*(*int64)(unsafe.Add(mBase, uint32(v6)+24)) = v13
	v20 = F_ExtendBufferedRel(m, v4+int32(-40), int32(3), int32(0), int32(9))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v6)+44)) = int64(0)
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v6)+48))
		*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v24
		*(*int32)(unsafe.Add(mBase, uint32(v6)+40)) = l0
		v27 = *(*int64)(unsafe.Add(mBase, uint32(v6)+40))
		*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = v27
		v34 = F_ExtendBufferedRel(m, v4+int32(-56), int32(3), int32(0), int32(9))
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return
		} else {
			v36 = int32(_a_F_ginbuildempty_0)
			v38 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuildempty[0]))
			*(*int32)(unsafe.Add(mBase, _c_F_ginbuildempty[0])) = v38 + int32(1)
			if v20 < int32(0) {
				v46 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuildempty[1]))
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v46+(v20^int32(-1))<<(uint(int32(2))%32))))
				v60 = v52
			} else {
				v54 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuildempty[2]))
				v60 = v54 + v20<<(uint(int32(13))%32) + int32(-8192)
			}
			v62 = int32(8)
			F_PageInit(m, v60, int32(_a_F_ginbuildempty_1), v62)
			mBase = m.M
			v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60)+16)))
			v65 = v60 + v64
			*(*int32)(unsafe.Add(mBase, uint32(v65))) = int32(-1)
			*(*uint16)(unsafe.Add(mBase, uint32(v65)+6)) = uint16(v62)
			v72 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v60-int32(-64)))) = v72
			*(*int64)(unsafe.Add(mBase, uint32(v60)+24)) = int64(-1)
			*(*int64)(unsafe.Add(mBase, uint32(v60)+32)) = v72
			*(*int64)(unsafe.Add(mBase, uint32(v60)+40)) = v72
			*(*int64)(unsafe.Add(mBase, uint32(v60)+48)) = v72
			*(*int32)(unsafe.Add(mBase, uint32(v60)+56)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v60)+72)) = int32(2)
			v86 = int32(80)
			*(*uint16)(unsafe.Add(mBase, uint32(v60)+12)) = uint16(v86)
			F_MarkBufferDirty(m, v20)
			mBase = m.M
			v89 = m.ExcPending
			if v89 != 0 {
				return
			} else {
				F_log_newpage_buffer(m, v20, int32(1))
				mBase = m.M
				v92 = m.ExcPending
				if v92 != 0 {
					return
				} else {
					v93 = int32(2)
					if v34 < int32(0) {
						v97 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuildempty[1]))
						v103 = *(*int32)(unsafe.Add(mBase, uint32(v97+(v34^int32(-1))<<(uint(int32(2))%32))))
						v111 = v103
					} else {
						v105 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuildempty[2]))
						v111 = v105 + v34<<(uint(int32(13))%32) + int32(-8192)
					}
					F_PageInit(m, v111, int32(_a_F_ginbuildempty_1), int32(8))
					mBase = m.M
					v115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v111)+16)))
					v116 = v111 + v115
					*(*int32)(unsafe.Add(mBase, uint32(v116))) = int32(-1)
					*(*uint16)(unsafe.Add(mBase, uint32(v116)+6)) = uint16(v93)
					F_MarkBufferDirty(m, v34)
					mBase = m.M
					v121 = m.ExcPending
					if v121 != 0 {
						return
					} else {
						F_log_newpage_buffer(m, v34, int32(0))
						mBase = m.M
						v124 = m.ExcPending
						if v124 != 0 {
							return
						} else {
							v125 = int32(_a_F_ginbuildempty_0)
							v127 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuildempty[0]))
							*(*int32)(unsafe.Add(mBase, _c_F_ginbuildempty[0])) = v127 - int32(1)
							F_UnlockReleaseBuffer(m, v20)
							mBase = m.M
							v132 = m.ExcPending
							if v132 != 0 {
								return
							} else {
								F_UnlockReleaseBuffer(m, v34)
								mBase = m.M
								v134 = m.ExcPending
								if v134 != 0 {
									return
								} else {
									m.G0 = v6 - int32(-64)
									return
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_gincostestimate(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v177 int32
	_ = v177
	var v192 int32
	_ = v192
	var v214 float64
	_ = v214
	var v215 int32
	_ = v215
	var v216 float64
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int64
	_ = v233
	var v243 float64
	_ = v243
	var v246 float64
	_ = v246
	var v249 int32
	_ = v249
	var v250 float64
	_ = v250
	var v259 int32
	_ = v259
	var v262 int64
	_ = v262
	var v265 float64
	_ = v265
	var v266 int32
	_ = v266
	var v269 float64
	_ = v269
	var v270 float64
	_ = v270
	var v273 float64
	_ = v273
	var v275 float64
	_ = v275
	var v276 float64
	_ = v276
	var v278 float64
	_ = v278
	var v285 float64
	_ = v285
	var v288 float64
	_ = v288
	var v289 float64
	_ = v289
	var v292 float64
	_ = v292
	var v302 float64
	_ = v302
	var v303 float64
	_ = v303
	var v305 float64
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v386 int32
	_ = v386
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v420 int32
	_ = v420
	var v447 float64
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v452 float64
	_ = v452
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v460 int32
	_ = v460
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v490 int32
	_ = v490
	var v496 int32
	_ = v496
	var v515 int32
	_ = v515
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v542 int32
	_ = v542
	var v562 int32
	_ = v562
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v587 float64
	_ = v587
	var v588 float64
	_ = v588
	var v591 float64
	_ = v591
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v618 float64
	_ = v618
	var v619 float64
	_ = v619
	var v622 float64
	_ = v622
	var v626 float64
	_ = v626
	var v627 int32
	_ = v627
	var v628 float64
	_ = v628
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v663 float64
	_ = v663
	var v664 int32
	_ = v664
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v693 float64
	_ = v693
	var v694 float64
	_ = v694
	var v696 float64
	_ = v696
	var v703 int32
	_ = v703
	var v705 int32
	_ = v705
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v716 int32
	_ = v716
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v723 int32
	_ = v723
	var v726 int32
	_ = v726
	var v733 float64
	_ = v733
	var v735 float64
	_ = v735
	var v737 float64
	_ = v737
	var v739 int32
	_ = v739
	var v740 float64
	_ = v740
	var v741 float64
	_ = v741
	var v742 float64
	_ = v742
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v749 float64
	_ = v749
	var v751 float64
	_ = v751
	var v755 float64
	_ = v755
	var v759 float64
	_ = v759
	var v762 float64
	_ = v762
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v809 int32
	_ = v809
	var v814 int32
	_ = v814
	var v850 int32
	_ = v850
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v865 int32
	_ = v865
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v934 int32
	_ = v934
	var v949 int32
	_ = v949
	var v978 int32
	_ = v978
	var v982 int32
	_ = v982
	var v986 int32
	_ = v986
	var v1026 float64
	_ = v1026
	var v1027 float64
	_ = v1027
	var v1028 int64
	_ = v1028
	var v1046 int32
	_ = v1046
	var v1086 int32
	_ = v1086
	var v1099 float64
	_ = v1099
	var v1109 float64
	_ = v1109
	var v1110 int64
	_ = v1110
	var v1115 float64
	_ = v1115
	var v1120 float64
	_ = v1120
	var v1121 float64
	_ = v1121
	var v1124 float64
	_ = v1124
	var v1127 float64
	_ = v1127
	var v1129 float64
	_ = v1129
	var v1130 float64
	_ = v1130
	var v1133 float64
	_ = v1133
	var v1138 float64
	_ = v1138
	var v1139 float64
	_ = v1139
	var v1141 float64
	_ = v1141
	var v1146 float64
	_ = v1146
	var v1150 float64
	_ = v1150
	var v1152 int32
	_ = v1152
	var v1153 float64
	_ = v1153
	var v1155 float64
	_ = v1155
	var v1162 float64
	_ = v1162
	var v1164 float64
	_ = v1164
	var v1168 float64
	_ = v1168
	var v1172 float64
	_ = v1172
	var v1181 float64
	_ = v1181
	var v1183 float64
	_ = v1183
	var v1186 float64
	_ = v1186
	var v1190 int32
	_ = v1190
	var v1191 float64
	_ = v1191
	var v1192 float64
	_ = v1192
	var v1198 int32
	_ = v1198
	var v1200 int32
	_ = v1200
	var v1204 int32
	_ = v1204
	var v1207 int32
	_ = v1207
	var v1208 float64
	_ = v1208
	var v1209 float64
	_ = v1209
	var v1210 float64
	_ = v1210
	var v1212 int32
	_ = v1212
	var v1215 float64
	_ = v1215
	var v1216 float64
	_ = v1216
	var v1220 float64
	_ = v1220
	var v1221 float64
	_ = v1221
	var v1225 float64
	_ = v1225
	var v1229 float64
	_ = v1229
	var v1234 float64
	_ = v1234
	var v1244 float64
	_ = v1244
	var v1246 float64
	_ = v1246
	var v1252 float64
	_ = v1252
	var v1254 float64
	_ = v1254
	var v1260 int32
	_ = v1260
	var v1262 int32
	_ = v1262
	var v1266 int32
	_ = v1266
	var v1269 int32
	_ = v1269
	var v1270 float64
	_ = v1270
	var v1271 float64
	_ = v1271
	var v1272 float64
	_ = v1272
	var v1274 int32
	_ = v1274
	var v1277 float64
	_ = v1277
	var v1278 float64
	_ = v1278
	var v1282 float64
	_ = v1282
	var v1283 float64
	_ = v1283
	var v1287 float64
	_ = v1287
	var v1291 float64
	_ = v1291
	var v1296 float64
	_ = v1296
	var v1306 float64
	_ = v1306
	var v1308 float64
	_ = v1308
	var v1314 float64
	_ = v1314
	var v1316 float64
	_ = v1316
	var v1318 float64
	_ = v1318
	var v1320 float64
	_ = v1320
	var v1322 float64
	_ = v1322
	var v1323 float64
	_ = v1323
	var v1325 float64
	_ = v1325
	var v1326 float64
	_ = v1326
	var v1328 float64
	_ = v1328
	var v1337 float64
	_ = v1337
	var v1340 float64
	_ = v1340
	var v1342 float64
	_ = v1342
	var v1347 float64
	_ = v1347
	var v1349 float64
	_ = v1349
	var v1350 float64
	_ = v1350
	var v1353 float64
	_ = v1353
	var v1359 int32
	_ = v1359
	var v1361 int32
	_ = v1361
	var v1365 int32
	_ = v1365
	var v1368 int32
	_ = v1368
	var v1369 float64
	_ = v1369
	var v1370 float64
	_ = v1370
	var v1371 float64
	_ = v1371
	var v1373 int32
	_ = v1373
	var v1376 float64
	_ = v1376
	var v1377 float64
	_ = v1377
	var v1381 float64
	_ = v1381
	var v1382 float64
	_ = v1382
	var v1386 float64
	_ = v1386
	var v1390 float64
	_ = v1390
	var v1395 float64
	_ = v1395
	var v1405 float64
	_ = v1405
	var v1407 float64
	_ = v1407
	var v1413 float64
	_ = v1413
	var v1415 float64
	_ = v1415
	var v1416 float64
	_ = v1416
	var v1417 float64
	_ = v1417
	var v1418 float64
	_ = v1418
	var v1419 float64
	_ = v1419
	var v1421 float64
	_ = v1421
	var v1425 float64
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1428 float64
	_ = v1428
	var v1432 int32
	_ = v1432
	var v1434 float64
	_ = v1434
	var v1435 float64
	_ = v1435
	var v1441 float64
	_ = v1441
	var v1443 float64
	_ = v1443
	var v1445 float64
	_ = v1445
	var v1448 float64
	_ = v1448
	v9 = int32(0)
	v36 = m.G0
	v38 = v36 - int32(256)
	m.G0 = v38
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v41 == v9 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v214 = *(*float64)(unsafe.Add(mBase, uint32(v40)+24))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
	v216 = base.F64_convert_i32_u(v215)
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+104)))
	if v217 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L2:
	;
	v192 = v9
	goto L1
L3:
	;
	goto L4
L4:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	if v44 <= int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v192 = v9
	goto L1
L6:
	;
	goto L7
L7:
	;
	v58 = v44
	v59 = v9
	v60 = v9
	goto L8
L8:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v82+v59<<(uint(int32(2))%32))))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+8))
	if v87 == int32(0) {
		v152 = v58
		v154 = v60
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v192 = v154
	goto L1
L10:
	;
	v177 = v59 + int32(1)
	if v177 < v152 {
		v58 = v152
		v59 = v177
		v60 = v154
		goto L8
	} else {
		goto L18
	}
L11:
	;
	v90 = int32(0)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	if v91 <= v90 {
		v152 = v58
		v154 = v60
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v103 = v90
	v107 = v60
	goto L13
L13:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v129+v103<<(uint(int32(2))%32))))
	v134 = F_lappend(m, v107, v133)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	v152 = v140
	v154 = v134
	goto L10
L15:
	;
	return
L16:
	;
	v137 = v103 + int32(1)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	if v137 < v138 {
		v103 = v137
		v107 = v134
		goto L13
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	goto L9
L19:
	;
	if base.F64_gt(v216, v243) != 0 {
		goto L26
	} else {
		goto L27
	}
L20:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v222 = F_index_open(m, v220, int32(0))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L15
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v233 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v38)+32)) = v233
	*(*int64)(unsafe.Add(mBase, uint32(v38)+24)) = v233
	*(*int64)(unsafe.Add(mBase, uint32(v38)+16)) = v233
	*(*int64)(unsafe.Add(mBase, uint32(v38)+8)) = v233
	v243 = float64(0)
	goto L19
L23:
	;
	F_ginGetStats(m, v222, v38+int32(8))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L15
	} else {
		goto L24
	}
L24:
	;
	F_relation_close(m, v222, int32(0))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L15
	} else {
		goto L25
	}
L25:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	v243 = base.F64_convert_i32_u(v231)
	goto L19
L26:
	;
	v246 = v243
	goto L28
L27:
	;
	v246 = float64(0)
	goto L28
L28:
	;
	if v215 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v40)+88))
	if v308 != 0 {
		goto L45
	} else {
		goto L46
	}
L30:
	;
	v285 = float64(10)
	if base.F64_gt(v216, v285) != 0 {
		goto L42
	} else {
		goto L43
	}
L31:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	v250 = base.F64_convert_i32_u(v249)
	if base.F64_le(v250, v216) == int32(0) {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	if base.F64_lt(base.F64_mul(v216, float64(0.25)), v250) == int32(0) {
		goto L30
	} else {
		goto L33
	}
L33:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
	if v259 == int32(0) {
		goto L30
	} else {
		goto L34
	}
L34:
	;
	v262 = *(*int64)(unsafe.Add(mBase, uint32(v38)+24))
	if v262 <= int64(0) {
		goto L30
	} else {
		goto L35
	}
L35:
	;
	v265 = base.F64_div(v216, v250)
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
	v269 = base.F64_ceil(base.F64_mul(v265, base.F64_convert_i32_u(v266)))
	v270 = base.F64_sub(v216, v246)
	v273 = base.F64_ceil(base.F64_mul(v265, base.F64_convert_i32_u(v259)))
	if base.F64_gt(v270, v273) != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v275 = v273
	goto L38
L37:
	;
	v275 = v270
	goto L38
L38:
	;
	v276 = base.F64_sub(v270, v275)
	if base.F64_gt(v276, v269) != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v278 = v269
	goto L41
L40:
	;
	v278 = v276
	goto L41
L41:
	;
	v302 = v278
	v303 = v275
	v305 = base.F64_ceil(base.F64_mul(v265, base.F64_convert_i64_u(v262)))
	goto L29
L42:
	;
	v288 = v216
	goto L44
L43:
	;
	v288 = v285
	goto L44
L44:
	;
	v289 = base.F64_sub(v288, v246)
	v292 = base.F64_floor(base.F64_mul(v289, float64(0.9)))
	v302 = base.F64_sub(v289, v292)
	v303 = v292
	v305 = base.F64_floor(base.F64_mul(v292, float64(100)))
	goto L29
L45:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v308)+4))
	if v309 <= int32(0) {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	v420 = v192
	goto L47
L47:
	;
	if base.F64_lt(v305, float64(1)) != 0 {
		goto L62
	} else {
		goto L63
	}
L48:
	;
	v409 = F_list_concat(m, v386, v192)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L15
	} else {
		goto L61
	}
L49:
	;
	v386 = int32(0)
	goto L48
L50:
	;
	goto L51
L51:
	;
	v313 = int32(0)
	v324 = v313
	v327 = v313
	goto L52
L52:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v308)+12))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v350+v324<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+4)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v38)+144)) = v354
	v360 = F_list_make1_impl(m, int32(1), v38+int32(4))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L15
	} else {
		goto L54
	}
L53:
	;
	v386 = v369
	goto L48
L54:
	;
	v363 = F_predicate_implied_by(m, v360, v192, int32(0))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L15
	} else {
		goto L55
	}
L55:
	;
	if v363 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v367 = F_list_concat(m, v327, v360)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L15
	} else {
		goto L59
	}
L57:
	;
	v369 = v327
	goto L58
L58:
	;
	v371 = v324 + int32(1)
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v308)+4))
	if v371 < v372 {
		v324 = v371
		v327 = v369
		goto L52
	} else {
		goto L60
	}
L59:
	;
	v369 = v367
	goto L58
L60:
	;
	goto L53
L61:
	;
	v420 = v409
	goto L47
L62:
	;
	v447 = float64(1)
	goto L64
L63:
	;
	v447 = v305
	goto L64
L64:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v448)+68))
	v450 = int32(0)
	v452 = F_clauselist_selectivity(m, l0, v420, v449, v450, v450)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L15
	} else {
		goto L65
	}
L65:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l5))) = v452
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	F_get_tablespace_page_costs(m, v455, v38+int32(40), int32(0))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L15
	} else {
		goto L66
	}
L66:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l6))) = int64(0)
	v468 = F__emscripten_memset_bulkmem(m, v38+int32(48), base.I32_extend8_s(int32(0)), int32(88))
	mBase = m.M
	goto L67
L67:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v38)+136)) = int64(4607182418800017408)
	v471 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v471 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L68:
	;
	m.G0 = v38 + int32(256)
	return
L69:
	;
	v1110 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l3))) = v1110
	*(*int64)(unsafe.Add(mBase, uint32(l4))) = v1110
	v1115 = F_pow(m, v303, float64(0.15))
	mBase = m.M
	v1120 = base.F64_div(v1099, v447)
	v1121 = float64(1)
	if base.F64_lt(v1120, v1121) != 0 {
		goto L138
	} else {
		goto L139
	}
L70:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v38)+128)) = v447
	*(*float64)(unsafe.Add(mBase, uint32(v38)+120)) = v447
	*(*int64)(unsafe.Add(mBase, uint32(v38)+112)) = int64(0)
	v1086 = v1046
	v1099 = float64(0)
	v1109 = v447
	goto L69
L71:
	;
	v1028 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l3))) = v1028
	*(*int64)(unsafe.Add(mBase, uint32(l4))) = v1028
	*(*int64)(unsafe.Add(mBase, uint32(l5))) = v1028
	goto L68
L72:
	;
	v934 = *(*int32)(unsafe.Add(mBase, uint32(v40)+40))
	if int32(0) < v934 {
		goto L128
	} else {
		goto L129
	}
L73:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v471)+4))
	if v474 <= int32(0) {
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v490 = int32(1)
	v496 = v9
	goto L75
L75:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v471)+12))
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v515+v496<<(uint(int32(2))%32))))
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v519)+8))
	if v520 == int32(0) {
		v865 = v490
		goto L77
	} else {
		goto L78
	}
L76:
	;
	if v865&int32(1) == int32(0) {
		goto L71
	} else {
		goto L127
	}
L77:
	;
	v891 = v496 + int32(1)
	v892 = *(*int32)(unsafe.Add(mBase, uint32(v471)+4))
	if v891 < v892 {
		v490 = v865
		v496 = v891
		goto L75
	} else {
		goto L126
	}
L78:
	;
	v523 = int32(0)
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v520)+4))
	if v524 <= v523 {
		v865 = v490
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v542 = v523
	goto L80
L80:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v520)+12))
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v562+v542<<(uint(int32(2))%32))))
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v566)+4))
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v567)))
	switch v568 - int32(17) {
	case 0:
		goto L85
	default:
		goto L83
	case 3:
		goto L84
	}
L81:
	;
	v865 = v850
	goto L77
L82:
	;
	v850 = int32(1)
	v852 = v542 + v850
	v853 = *(*int32)(unsafe.Add(mBase, uint32(v520)+4))
	if v852 < v853 {
		v542 = v852
		goto L80
	} else {
		goto L125
	}
L83:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		goto L15
	} else {
		goto L122
	}
L84:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v567)+4))
	v603 = int32(*(*int16)(unsafe.Add(mBase, uint32(v519)+14)))
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v567)+28))
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v604)+12))
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v605)+4))
	v607 = F_estimate_expression_value(m, l0, v606)
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L15
	} else {
		goto L96
	}
L85:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v567)+4))
	v572 = int32(*(*int16)(unsafe.Add(mBase, uint32(v519)+14)))
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v567)+28))
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v573)+12))
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v574)+4))
	v576 = F_estimate_expression_value(m, l0, v575)
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L15
	} else {
		goto L86
	}
L86:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v576)))
	if v578 == int32(27) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v576)+4))
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v581)))
	v583 = v581
	v584 = v582
	goto L89
L88:
	;
	v583 = v576
	v584 = v578
	goto L89
L89:
	;
	if v584 != int32(7) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v587 = *(*float64)(unsafe.Add(mBase, uint32(v38)+120))
	v588 = float64(1)
	*(*float64)(unsafe.Add(mBase, uint32(v38)+120)) = base.F64_add(v587, v588)
	v591 = *(*float64)(unsafe.Add(mBase, uint32(v38)+128))
	*(*float64)(unsafe.Add(mBase, uint32(v38)+128)) = base.F64_add(v591, v588)
	goto L82
L91:
	;
	goto L92
L92:
	;
	v595 = int32(0)
	v596 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v583)+24)))
	if v596 != 0 {
		v865 = v595
		goto L77
	} else {
		goto L93
	}
L93:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v583)+20))
	v600 = F_gincost_pattern(m, v40, v572, v571, v597, v38+int32(48))
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L15
	} else {
		goto L94
	}
L94:
	;
	if v600 != 0 {
		goto L82
	} else {
		goto L95
	}
L95:
	;
	v865 = v595
	goto L77
L96:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v607)))
	if v609 == int32(27) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v607)+4))
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v612)))
	v614 = v612
	v615 = v613
	goto L99
L98:
	;
	v614 = v607
	v615 = v609
	goto L99
L99:
	;
	if v615 != int32(7) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v618 = *(*float64)(unsafe.Add(mBase, uint32(v38)+120))
	v619 = float64(1)
	*(*float64)(unsafe.Add(mBase, uint32(v38)+120)) = base.F64_add(v618, v619)
	v622 = *(*float64)(unsafe.Add(mBase, uint32(v38)+128))
	*(*float64)(unsafe.Add(mBase, uint32(v38)+128)) = base.F64_add(v622, v619)
	v626 = F_estimate_array_length(m, l0, v614)
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L15
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614)+24)))
	if v631 != 0 {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	v628 = *(*float64)(unsafe.Add(mBase, uint32(v38)+136))
	*(*float64)(unsafe.Add(mBase, uint32(v38)+136)) = base.F64_mul(v626, v628)
	goto L82
L104:
	;
	v865 = int32(0)
	goto L77
L105:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v614)+20))
	v633 = F_pg_detoast_datum(m, v632)
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L15
	} else {
		goto L106
	}
L106:
	;
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v633)+12))
	F_get_typlenbyvalalign(m, v635, v38+int32(254), v38+int32(253), v38+int32(252))
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L15
	} else {
		goto L107
	}
L107:
	;
	v645 = int32(*(*int16)(unsafe.Add(mBase, uint32(v38)+254)))
	v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+253)))
	v647 = int32(*(*int8)(unsafe.Add(mBase, uint32(v38)+252)))
	F_deconstruct_array(m, v633, v645, v646, v647, v38+int32(244), v38+int32(240), v38+int32(248))
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L15
	} else {
		goto L108
	}
L108:
	;
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v38)+248))
	if v656 <= int32(0) {
		goto L104
	} else {
		goto L109
	}
L109:
	;
	v663 = float64(0)
	v664 = int32(0)
	v677 = v664
	v679 = v664
	v693 = v663
	v694 = v663
	v696 = v663
	goto L110
L110:
	;
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v38)+240))
	v705 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v703+v677))))
	if v705 != 0 {
		v739 = v679
		v740 = v693
		v741 = v694
		v742 = v696
		goto L112
	} else {
		goto L113
	}
L111:
	;
	if v739 == int32(0) {
		goto L104
	} else {
		goto L121
	}
L112:
	;
	v744 = v677 + int32(1)
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v38)+248))
	if v744 < v745 {
		v677 = v744
		v679 = v739
		v693 = v740
		v694 = v741
		v696 = v742
		goto L110
	} else {
		goto L120
	}
L113:
	;
	v711 = F__emscripten_memset_bulkmem(m, v38+int32(144), base.I32_extend8_s(int32(0)), int32(96))
	mBase = m.M
	goto L114
L114:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v38)+244))
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v712+v677<<(uint(int32(2))%32))))
	v719 = F_gincost_pattern(m, v40, v603, v602, v716, v38+int32(144))
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L15
	} else {
		goto L115
	}
L115:
	;
	if v719 == int32(0) {
		v739 = v679
		v740 = v693
		v741 = v694
		v742 = v696
		goto L112
	} else {
		goto L116
	}
L116:
	;
	v723 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38+int32(144)+v603))))
	if v723 != int32(1) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v733 = *(*float64)(unsafe.Add(mBase, uint32(v38)+224))
	v735 = *(*float64)(unsafe.Add(mBase, uint32(v38)+216))
	v737 = *(*float64)(unsafe.Add(mBase, uint32(v38)+208))
	v739 = v679 + int32(1)
	v740 = base.F64_add(v693, v733)
	v741 = base.F64_add(v694, v737)
	v742 = base.F64_add(v696, v735)
	goto L112
L118:
	;
	v726 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v603+(v38+int32(176))))))
	if v726 != 0 {
		goto L117
	} else {
		goto L119
	}
L119:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v38)+224)) = v447
	*(*float64)(unsafe.Add(mBase, uint32(v38)+216)) = v447
	*(*int64)(unsafe.Add(mBase, uint32(v38)+208)) = int64(0)
	goto L117
L120:
	;
	goto L111
L121:
	;
	v749 = base.F64_convert_i32_s(v739)
	v751 = *(*float64)(unsafe.Add(mBase, uint32(v38)+112))
	*(*float64)(unsafe.Add(mBase, uint32(v38)+112)) = base.F64_add(base.F64_div(v741, v749), v751)
	v755 = *(*float64)(unsafe.Add(mBase, uint32(v38)+120))
	*(*float64)(unsafe.Add(mBase, uint32(v38)+120)) = base.F64_add(base.F64_div(v742, v749), v755)
	v759 = *(*float64)(unsafe.Add(mBase, uint32(v38)+128))
	*(*float64)(unsafe.Add(mBase, uint32(v38)+128)) = base.F64_add(base.F64_div(v740, v749), v759)
	v762 = *(*float64)(unsafe.Add(mBase, uint32(v38)+136))
	*(*float64)(unsafe.Add(mBase, uint32(v38)+136)) = base.F64_mul(v762, v749)
	goto L82
L122:
	;
	v805 = *(*int32)(unsafe.Add(mBase, uint32(v567)))
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = v805
	F_errmsg_internal(m, int32(_a_F_gincostestimate_0), v38)
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L15
	} else {
		goto L123
	}
L123:
	;
	F_errfinish(m, int32(_a_F_gincostestimate_1), int32(_a_F_gincostestimate_2), int32(_a_F_gincostestimate_3))
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L15
	} else {
		goto L124
	}
L124:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L125:
	;
	goto L81
L126:
	;
	goto L76
L127:
	;
	goto L72
L128:
	;
	v949 = int32(0)
	goto L131
L129:
	;
	goto L130
L130:
	;
	if v192 == int32(0) {
		v1046 = int32(1)
		goto L70
	} else {
		goto L137
	}
L131:
	;
	v978 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38+int32(48)+v949))))
	if v978 != int32(1) {
		goto L133
	} else {
		goto L134
	}
L132:
	;
	goto L130
L133:
	;
	v986 = v949 + int32(1)
	if v986 != v934 {
		v949 = v986
		goto L131
	} else {
		goto L136
	}
L134:
	;
	v982 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v949+(v38+int32(80))))))
	if v982 != 0 {
		goto L133
	} else {
		goto L135
	}
L135:
	;
	v1046 = base.B2i32(v192 == int32(0))
	goto L70
L136:
	;
	goto L132
L137:
	;
	v1026 = *(*float64)(unsafe.Add(mBase, uint32(v38)+112))
	v1027 = *(*float64)(unsafe.Add(mBase, uint32(v38)+128))
	v1086 = int32(0)
	v1099 = v1026
	v1109 = v1027
	goto L69
L138:
	;
	v1124 = v1120
	goto L140
L139:
	;
	v1124 = v1121
	goto L140
L140:
	;
	v1127 = base.F64_add(base.F64_add(v246, base.F64_ceil(base.F64_mul(base.F64_nearest(v1115), v1109))), base.F64_ceil(base.F64_mul(v303, v1124)))
	v1129 = base.F64_ceil(base.F64_mul(v302, v1124))
	v1130 = *(*float64)(unsafe.Add(mBase, uint32(v38)+136))
	if base.F64_gt(v447, float64(1)) != 0 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v1133 = F_log(m, v447)
	mBase = m.M
	v1138 = *(*float64)(unsafe.Add(mBase, _c_F_gincostestimate[0]))
	v1139 = base.F64_mul(base.F64_ceil(base.F64_div(v1133, float64(0.6931471805599453))), v1138)
	v1141 = *(*float64)(unsafe.Add(mBase, uint32(l3)))
	*(*float64)(unsafe.Add(mBase, uint32(l3))) = base.F64_add(base.F64_mul(v1139, v1109), v1141)
	v1146 = *(*float64)(unsafe.Add(mBase, uint32(l4)))
	*(*float64)(unsafe.Add(mBase, uint32(l4))) = base.F64_add(base.F64_mul(base.F64_mul(v1139, v1130), v1109), v1146)
	goto L143
L142:
	;
	goto L143
L143:
	;
	v1150 = float64(50)
	v1152 = int32(_a_F_gincostestimate_4)
	v1153 = *(*float64)(unsafe.Add(mBase, _c_F_gincostestimate[0]))
	v1155 = *(*float64)(unsafe.Add(mBase, uint32(l3)))
	*(*float64)(unsafe.Add(mBase, uint32(l3))) = base.F64_add(base.F64_mul(base.F64_mul(v1127, v1150), v1153), v1155)
	v1162 = *(*float64)(unsafe.Add(mBase, _c_F_gincostestimate[0]))
	v1164 = *(*float64)(unsafe.Add(mBase, uint32(l4)))
	*(*float64)(unsafe.Add(mBase, uint32(l4))) = base.F64_add(base.F64_mul(base.F64_mul(base.F64_mul(v1127, v1130), v1150), v1162), v1164)
	v1168 = *(*float64)(unsafe.Add(mBase, _c_F_gincostestimate[0]))
	v1172 = *(*float64)(unsafe.Add(mBase, uint32(l3)))
	*(*float64)(unsafe.Add(mBase, uint32(l3))) = base.F64_add(base.F64_mul(base.F64_mul(v1168, v1150), v1129), v1172)
	v1181 = *(*float64)(unsafe.Add(mBase, _c_F_gincostestimate[0]))
	v1183 = *(*float64)(unsafe.Add(mBase, uint32(l4)))
	*(*float64)(unsafe.Add(mBase, uint32(l4))) = base.F64_add(base.F64_mul(base.F64_mul(base.F64_mul(v1129, base.F64_add(v1130, float64(-1))), v1150), v1181), v1183)
	v1186 = float64(1)
	v1190 = base.F64_gt(l2, v1186) | base.F64_gt(v1130, v1186)
	if v1190 != 0 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v1191 = base.F64_mul(l2, v1130)
	v1192 = base.F64_mul(v1127, v1191)
	if base.F64_lt(v303, float64(4.294967296e+09))&base.F64_ge(v303, float64(0)) != 0 {
		goto L148
	} else {
		goto L149
	}
L145:
	;
	v1316 = v1127
	v1318 = v1129
	goto L146
L146:
	;
	v1320 = *(*float64)(unsafe.Add(mBase, uint32(v38)+40))
	v1322 = *(*float64)(unsafe.Add(mBase, uint32(l3)))
	v1323 = base.F64_add(base.F64_mul(base.F64_add(v1318, v1316), v1320), v1322)
	*(*float64)(unsafe.Add(mBase, uint32(l3))) = v1323
	v1325 = *(*float64)(unsafe.Add(mBase, uint32(l5)))
	v1326 = *(*float64)(unsafe.Add(mBase, uint32(v38)+120))
	v1328 = *(*float64)(unsafe.Add(mBase, _c_F_gincostestimate[0]))
	*(*float64)(unsafe.Add(mBase, uint32(l3))) = base.F64_add(base.F64_mul(base.F64_mul(v1328, float64(50)), v1109), v1323)
	v1337 = base.F64_ceil(base.F64_mul(v1325, base.F64_div(v214, float64(2730))))
	v1340 = base.F64_ceil(base.F64_div(base.F64_mul(v302, v1326), v447))
	if base.F64_gt(v1337, v1340) != 0 {
		goto L191
	} else {
		goto L192
	}
L147:
	;
	v1204 = int32(1)
	if base.Ui32(v1200) <= base.Ui32(v1204) {
		goto L152
	} else {
		goto L153
	}
L148:
	;
	v1198 = base.I32_trunc_f64_u(v303)
	v1200 = v1198
	goto L147
L149:
	;
	goto L150
L150:
	;
	v1200 = int32(0)
	goto L147
L151:
	;
	v1254 = base.F64_mul(v1129, v1191)
	if base.F64_lt(v302, float64(4.294967296e+09))&base.F64_ge(v302, float64(0)) != 0 {
		goto L170
	} else {
		goto L171
	}
L152:
	;
	v1207 = v1204
	goto L154
L153:
	;
	v1207 = v1200
	goto L154
L154:
	;
	v1208 = base.F64_convert_i32_u(v1207)
	v1209 = base.F64_add(v1208, v1208)
	v1210 = float64(1)
	v1212 = *(*int32)(unsafe.Add(mBase, _c_F_gincostestimate[1]))
	v1215 = *(*float64)(unsafe.Add(mBase, uint32(l0)+288))
	v1216 = base.F64_add(v303, v1215)
	if base.F64_gt(v1216, v1210) != 0 {
		goto L156
	} else {
		goto L157
	}
L155:
	;
	v1252 = v1246
	goto L151
L156:
	;
	v1220 = v1216
	goto L158
L157:
	;
	v1220 = v1210
	goto L158
L158:
	;
	v1221 = base.F64_div(base.F64_mul(v1208, base.F64_convert_i32_s(v1212)), v1220)
	if base.F64_le(v1221, float64(1)) != 0 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v1225 = v1210
	goto L161
L160:
	;
	v1225 = base.F64_ceil(v1221)
	goto L161
L161:
	;
	if base.F64_le(v1208, v1225) != 0 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v1229 = base.F64_div(base.F64_mul(v1192, v1209), base.F64_add(v1209, v1192))
	if base.F64_ge(v1229, v1208) != 0 {
		v1246 = v1208
		goto L155
	} else {
		goto L165
	}
L163:
	;
	goto L164
L164:
	;
	v1234 = base.F64_div(base.F64_mul(v1209, v1225), base.F64_sub(v1209, v1225))
	if base.F64_ge(v1234, v1192) != 0 {
		goto L166
	} else {
		goto L167
	}
L165:
	;
	v1252 = base.F64_ceil(v1229)
	goto L151
L166:
	;
	v1244 = base.F64_div(base.F64_mul(v1192, v1209), base.F64_add(v1209, v1192))
	goto L168
L167:
	;
	v1244 = base.F64_add(v1225, base.F64_div(base.F64_mul(base.F64_sub(v1208, v1225), base.F64_sub(v1192, v1234)), v1208))
	goto L168
L168:
	;
	v1246 = base.F64_ceil(v1244)
	goto L155
L169:
	;
	v1266 = int32(1)
	if base.Ui32(v1262) <= base.Ui32(v1266) {
		goto L174
	} else {
		goto L175
	}
L170:
	;
	v1260 = base.I32_trunc_f64_u(v302)
	v1262 = v1260
	goto L169
L171:
	;
	goto L172
L172:
	;
	v1262 = int32(0)
	goto L169
L173:
	;
	v1316 = base.F64_div(v1252, l2)
	v1318 = base.F64_div(v1314, l2)
	goto L146
L174:
	;
	v1269 = v1266
	goto L176
L175:
	;
	v1269 = v1262
	goto L176
L176:
	;
	v1270 = base.F64_convert_i32_u(v1269)
	v1271 = base.F64_add(v1270, v1270)
	v1272 = float64(1)
	v1274 = *(*int32)(unsafe.Add(mBase, _c_F_gincostestimate[1]))
	v1277 = *(*float64)(unsafe.Add(mBase, uint32(l0)+288))
	v1278 = base.F64_add(v302, v1277)
	if base.F64_gt(v1278, v1272) != 0 {
		goto L178
	} else {
		goto L179
	}
L177:
	;
	v1314 = v1308
	goto L173
L178:
	;
	v1282 = v1278
	goto L180
L179:
	;
	v1282 = v1272
	goto L180
L180:
	;
	v1283 = base.F64_div(base.F64_mul(v1270, base.F64_convert_i32_s(v1274)), v1282)
	if base.F64_le(v1283, float64(1)) != 0 {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v1287 = v1272
	goto L183
L182:
	;
	v1287 = base.F64_ceil(v1283)
	goto L183
L183:
	;
	if base.F64_le(v1270, v1287) != 0 {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v1291 = base.F64_div(base.F64_mul(v1254, v1271), base.F64_add(v1271, v1254))
	if base.F64_ge(v1291, v1270) != 0 {
		v1308 = v1270
		goto L177
	} else {
		goto L187
	}
L185:
	;
	goto L186
L186:
	;
	v1296 = base.F64_div(base.F64_mul(v1271, v1287), base.F64_sub(v1271, v1287))
	if base.F64_ge(v1296, v1254) != 0 {
		goto L188
	} else {
		goto L189
	}
L187:
	;
	v1314 = base.F64_ceil(v1291)
	goto L173
L188:
	;
	v1306 = base.F64_div(base.F64_mul(v1254, v1271), base.F64_add(v1271, v1254))
	goto L190
L189:
	;
	v1306 = base.F64_add(v1287, base.F64_div(base.F64_mul(base.F64_sub(v1270, v1287), base.F64_sub(v1254, v1296)), v1270))
	goto L190
L190:
	;
	v1308 = base.F64_ceil(v1306)
	goto L177
L191:
	;
	v1342 = v1337
	goto L193
L192:
	;
	v1342 = v1340
	goto L193
L193:
	;
	v1347 = *(*float64)(unsafe.Add(mBase, _c_F_gincostestimate[0]))
	v1349 = *(*float64)(unsafe.Add(mBase, uint32(l4)))
	v1350 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_mul(v1130, v1342), float64(50)), v1347), v1349)
	*(*float64)(unsafe.Add(mBase, uint32(l4))) = v1350
	if v1190 != 0 {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	v1353 = base.F64_mul(base.F64_mul(l2, v1130), v1342)
	if base.F64_lt(v302, float64(4.294967296e+09))&base.F64_ge(v302, float64(0)) != 0 {
		goto L198
	} else {
		goto L199
	}
L195:
	;
	v1417 = v1342
	v1418 = v1320
	v1419 = v1350
	goto L196
L196:
	;
	v1421 = *(*float64)(unsafe.Add(mBase, uint32(l3)))
	*(*float64)(unsafe.Add(mBase, uint32(l4))) = base.F64_add(v1419, base.F64_add(base.F64_mul(v1417, v1418), v1421))
	v1425 = F_index_other_operands_eval_cost(m, l0, v192)
	mBase = m.M
	v1426 = m.ExcPending
	if v1426 != 0 {
		goto L15
	} else {
		goto L219
	}
L197:
	;
	v1365 = int32(1)
	if base.Ui32(v1361) <= base.Ui32(v1365) {
		goto L202
	} else {
		goto L203
	}
L198:
	;
	v1359 = base.I32_trunc_f64_u(v302)
	v1361 = v1359
	goto L197
L199:
	;
	goto L200
L200:
	;
	v1361 = int32(0)
	goto L197
L201:
	;
	v1415 = *(*float64)(unsafe.Add(mBase, uint32(v38)+40))
	v1416 = *(*float64)(unsafe.Add(mBase, uint32(l4)))
	v1417 = base.F64_div(v1413, l2)
	v1418 = v1415
	v1419 = v1416
	goto L196
L202:
	;
	v1368 = v1365
	goto L204
L203:
	;
	v1368 = v1361
	goto L204
L204:
	;
	v1369 = base.F64_convert_i32_u(v1368)
	v1370 = base.F64_add(v1369, v1369)
	v1371 = float64(1)
	v1373 = *(*int32)(unsafe.Add(mBase, _c_F_gincostestimate[1]))
	v1376 = *(*float64)(unsafe.Add(mBase, uint32(l0)+288))
	v1377 = base.F64_add(v302, v1376)
	if base.F64_gt(v1377, v1371) != 0 {
		goto L206
	} else {
		goto L207
	}
L205:
	;
	v1413 = v1407
	goto L201
L206:
	;
	v1381 = v1377
	goto L208
L207:
	;
	v1381 = v1371
	goto L208
L208:
	;
	v1382 = base.F64_div(base.F64_mul(v1369, base.F64_convert_i32_s(v1373)), v1381)
	if base.F64_le(v1382, float64(1)) != 0 {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	v1386 = v1371
	goto L211
L210:
	;
	v1386 = base.F64_ceil(v1382)
	goto L211
L211:
	;
	if base.F64_le(v1369, v1386) != 0 {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	v1390 = base.F64_div(base.F64_mul(v1353, v1370), base.F64_add(v1370, v1353))
	if base.F64_ge(v1390, v1369) != 0 {
		v1407 = v1369
		goto L205
	} else {
		goto L215
	}
L213:
	;
	goto L214
L214:
	;
	v1395 = base.F64_div(base.F64_mul(v1370, v1386), base.F64_sub(v1370, v1386))
	if base.F64_ge(v1395, v1353) != 0 {
		goto L216
	} else {
		goto L217
	}
L215:
	;
	v1413 = base.F64_ceil(v1390)
	goto L201
L216:
	;
	v1405 = base.F64_div(base.F64_mul(v1353, v1370), base.F64_add(v1370, v1353))
	goto L218
L217:
	;
	v1405 = base.F64_add(v1386, base.F64_div(base.F64_mul(base.F64_sub(v1369, v1386), base.F64_sub(v1353, v1395)), v1369))
	goto L218
L218:
	;
	v1407 = base.F64_ceil(v1405)
	goto L205
L219:
	;
	v1428 = *(*float64)(unsafe.Add(mBase, _c_F_gincostestimate[0]))
	if v1086 == int32(0) {
		goto L220
	} else {
		goto L221
	}
L220:
	;
	v1432 = *(*int32)(unsafe.Add(mBase, uint32(v192)+4))
	v1434 = base.F64_convert_i32_s(v1432)
	goto L222
L221:
	;
	v1434 = float64(0)
	goto L222
L222:
	;
	v1435 = *(*float64)(unsafe.Add(mBase, uint32(l3)))
	*(*float64)(unsafe.Add(mBase, uint32(l3))) = base.F64_add(v1425, v1435)
	v1441 = *(*float64)(unsafe.Add(mBase, uint32(l4)))
	v1443 = base.F64_add(base.F64_mul(base.F64_mul(v1109, v1130), base.F64_mul(v1428, v1434)), base.F64_add(v1425, v1441))
	*(*float64)(unsafe.Add(mBase, uint32(l4))) = v1443
	v1445 = *(*float64)(unsafe.Add(mBase, uint32(l5)))
	v1448 = *(*float64)(unsafe.Add(mBase, _c_F_gincostestimate[2]))
	*(*float64)(unsafe.Add(mBase, uint32(l4))) = base.F64_add(base.F64_mul(base.F64_mul(v214, v1445), v1448), v1443)
	*(*float64)(unsafe.Add(mBase, uint32(l7))) = v1417
	goto L68
}
func F_gininsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v70 int32
	_ = v70
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v110 int32
	_ = v110
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int64
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v177 int32
	_ = v177
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v272 int32
	_ = v272
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int64
	_ = v413
	var v415 int32
	_ = v415
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v506 int32
	_ = v506
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v586 int32
	_ = v586
	var v602 int32
	_ = v602
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v666 int32
	_ = v666
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v687 int32
	_ = v687
	var v691 int32
	_ = v691
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v700 int32
	_ = v700
	var v701 int64
	_ = v701
	var v708 int32
	_ = v708
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v726 int32
	_ = v726
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	var v763 int64
	_ = v763
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v781 int32
	_ = v781
	var v786 int32
	_ = v786
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v805 int32
	_ = v805
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v815 int32
	_ = v815
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v846 int32
	_ = v846
	var v852 int32
	_ = v852
	var v854 int32
	_ = v854
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v863 int32
	_ = v863
	var v869 int32
	_ = v869
	var v875 int32
	_ = v875
	var v880 int32
	_ = v880
	var v885 int32
	_ = v885
	var v888 int32
	_ = v888
	var v889 int64
	_ = v889
	var v893 int64
	_ = v893
	var v895 int64
	_ = v895
	var v897 int64
	_ = v897
	var v899 int64
	_ = v899
	var v901 int64
	_ = v901
	var v903 int64
	_ = v903
	var v908 int32
	_ = v908
	var v913 int32
	_ = v913
	var v916 int64
	_ = v916
	var v917 int32
	_ = v917
	var v925 int64
	_ = v925
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v933 int32
	_ = v933
	var v939 int32
	_ = v939
	var v955 int32
	_ = v955
	var v958 int32
	_ = v958
	var v965 int32
	_ = v965
	var v982 int32
	_ = v982
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1016 int32
	_ = v1016
	var v1025 int32
	_ = v1025
	var v1030 int32
	_ = v1030
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1064 int32
	_ = v1064
	var v1069 int32
	_ = v1069
	var v1093 int32
	_ = v1093
	v21 = m.G0
	v23 = v21 - int32(32)
	m.G0 = v23
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_gininsert[0]))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l7)+136))
	if v27 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l7)+140))
	*(*int32)(unsafe.Add(mBase, _c_F_gininsert[0])) = v31
	v34 = F_palloc(m, int32(_a_F_gininsert_0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v43 = v27
	goto L3
L3:
	;
	v48 = F_AllocSetContextCreateInternal(m, v26, int32(_a_F_gininsert_1), int32(0), int32(_a_F_gininsert_2), int32(_a_F_gininsert_3))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L4
	} else {
		goto L7
	}
L4:
	;
	return int32(0)
L5:
	;
	F_initGinState(m, v34, l0)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7)+136)) = v34
	*(*int32)(unsafe.Add(mBase, _c_F_gininsert[0])) = v26
	v43 = v34
	goto L3
L7:
	;
	v50 = int32(_a_F_gininsert_4)
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_gininsert[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_gininsert[0])) = v48
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v54 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_gininsert[0])) = v51
	F_MemoryContextDelete(m, v48)
	mBase = m.M
	v1093 = m.ExcPending
	if v1093 != 0 {
		goto L4
	} else {
		goto L199
	}
L9:
	;
	v161 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+16)) = v161
	*(*int64)(unsafe.Add(mBase, uint32(v23)+8)) = v161
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	if int32(0) < v166 {
		goto L24
	} else {
		goto L25
	}
L10:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+4)))
	if v57 != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v58 = int32(0)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	if v60 <= v58 {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v70 = v58
	goto L13
L13:
	;
	v84 = v70 + int32(1)
	v86 = v84 & int32(_a_F_gininsert_5)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l1+v70<<(uint(int32(2))%32))))
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v70))))
	v97 = F_ginExtractEntries(m, v43, v86, v90, v92, v23+int32(28), v23+int32(8))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L4
	} else {
		goto L15
	}
L14:
	;
	goto L8
L15:
	;
	v99 = int32(0)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
	if v99 < v100 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v110 = v99
	goto L19
L17:
	;
	goto L18
L18:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	if v84 < v159 {
		v70 = v84
		goto L13
	} else {
		goto L23
	}
L19:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v97+v110<<(uint(int32(2))%32))))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v129 = int32(*(*int8)(unsafe.Add(mBase, uint32(v127+v110))))
	F_ginEntryInsert(m, v43, v86, v126, v129, l3, int32(1), int32(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L4
	} else {
		goto L21
	}
L20:
	;
	goto L18
L21:
	;
	v135 = v110 + int32(1)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
	if v135 < v136 {
		v110 = v135
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	goto L14
L24:
	;
	v177 = int32(0)
	goto L27
L25:
	;
	goto L26
L26:
	;
	v380 = int32(0)
	v385 = m.G0
	v387 = v385 - int32(96)
	m.G0 = v387
	v390 = v23 + int32(8)
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v390)+4))
	if v391 == v380 {
		goto L64
	} else {
		goto L65
	}
L27:
	;
	v190 = int32(8)
	v191 = v23 + v190
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l1+v177<<(uint(int32(2))%32))))
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v177))))
	v198 = m.G0
	v200 = v198 - int32(16)
	m.G0 = v200
	v203 = v177 + int32(1)
	v205 = v203 & int32(_a_F_gininsert_5)
	v210 = F_ginExtractEntries(m, v43, v205, v195, v197, v200+v190, v200+int32(12))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L4
	} else {
		goto L29
	}
L28:
	;
	goto L26
L29:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v200)+8))
	if v212 < int32(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v357)))
	if v203 < v358 {
		v177 = v203
		goto L27
	} else {
		goto L61
	}
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L4
	} else {
		goto L58
	}
L32:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v191)+4))
	v216 = v215 + v212
	if base.Ui32(int32(268435456)) <= base.Ui32(v216) {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v191)))
	if v219 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v200)+8))
	if int32(0) < v260 {
		goto L51
	} else {
		goto L52
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v191))) = v257
	goto L34
L36:
	;
	v224 = int32(16)
	if base.Ui32(v212) <= base.Ui32(v224) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v191)+8))
	if base.Ui32(v216) <= base.Ui32(v240) {
		goto L34
	} else {
		goto L46
	}
L39:
	;
	v227 = v224
	goto L41
L40:
	;
	v227 = v212
	goto L41
L41:
	;
	if v227&(v227-int32(1)) != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v234 = int32(1) << (uint(int32(32)-base.I32_clz(v227)) % 32)
	goto L44
L43:
	;
	v234 = v227
	goto L44
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v191)+8)) = v234
	v238 = F_palloc(m, v234<<(uint(int32(2))%32))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	v257 = v238
	goto L35
L46:
	;
	v242 = int32(1)
	if v216&(v216-v242) != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v250 = v242 << (uint(int32(32)-base.I32_clz(v216)) % 32)
	goto L49
L48:
	;
	v250 = v216
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v191)+8)) = v250
	v254 = F_repalloc(m, v219, v250<<(uint(int32(2))%32))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	v257 = v254
	goto L35
L51:
	;
	v272 = int32(0)
	goto L54
L52:
	;
	goto L53
L53:
	;
	m.G0 = v200 + int32(16)
	goto L30
L54:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v210+v272<<(uint(int32(2))%32))))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v200)+12))
	v290 = int32(*(*int8)(unsafe.Add(mBase, uint32(v288+v272))))
	v291 = int32(0)
	v295 = F_GinFormTuple(m, v43, v205, v287, v290, v291, v291, v291, int32(1))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L4
	} else {
		goto L56
	}
L55:
	;
	goto L53
L56:
	;
	v297 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v295)+4)) = uint16(v297)
	v299 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v295))) = v299
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v191)+4))
	v302 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v191)+4)) = v301 + v302
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v191)))
	*(*int32)(unsafe.Add(mBase, uint32(v305+v301<<(uint(int32(2))%32)))) = v295
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v191)+12))
	v311 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v295)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v191)+12)) = v310 + v311&int32(_a_F_gininsert_6)
	v317 = v272 + v302
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v200)+8))
	if v317 < v318 {
		v272 = v317
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	F_errmsg_internal(m, int32(_a_F_gininsert_7), int32(0))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F_gininsert_8), int32(504), int32(_a_F_gininsert_9))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L4
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L61:
	;
	goto L28
L62:
	;
	goto L8
L63:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1057 = m.ExcPending
	if v1057 != 0 {
		goto L4
	} else {
		goto L196
	}
L64:
	;
	m.G0 = v387 + int32(96)
	goto L62
L65:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v395)+48))
	v397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v396)+118)))
	if v397 != int32(112) {
		v410 = int32(0)
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v395)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v387)+16)) = v411
	v413 = *(*int64)(unsafe.Add(mBase, uint32(v395)))
	*(*int64)(unsafe.Add(mBase, uint32(v387)+8)) = v413
	v415 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v387)+88)) = v415
	*(*int64)(unsafe.Add(mBase, uint32(v387)+80)) = int64(-1)
	v420 = F_ReadBuffer(m, v395, v415)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L4
	} else {
		goto L71
	}
L67:
	;
	v402 = *(*int32)(unsafe.Add(mBase, _c_F_gininsert[1]))
	if int32(0) < v402 {
		v410 = int32(1)
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v395)+32))
	if v406 != 0 {
		v410 = int32(0)
		goto L66
	} else {
		goto L69
	}
L69:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v395)+40))
	v410 = base.B2i32(v407 == int32(0))
	goto L66
L70:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v390)+12))
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v390)+4))
	if base.Ui32(v440+v441<<(uint(int32(2))%32)) <= base.Ui32(int32(_a_F_gininsert_10)) {
		goto L81
	} else {
		goto L82
	}
L71:
	;
	if v420 < int32(0) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v425 = *(*int32)(unsafe.Add(mBase, _c_F_gininsert[2]))
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v425+(v420^int32(-1))<<(uint(int32(2))%32))))
	v439 = v431
	goto L70
L73:
	;
	goto L74
L74:
	;
	v433 = *(*int32)(unsafe.Add(mBase, _c_F_gininsert[3]))
	v439 = v433 + v420<<(uint(int32(13))%32) + int32(-8192)
	goto L70
L75:
	;
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(v395)+180))
	if v1003 != 0 {
		goto L189
	} else {
		goto L190
	}
L76:
	;
	F_UnlockReleaseBuffer(m, v965)
	mBase = m.M
	v982 = m.ExcPending
	if v982 != 0 {
		goto L4
	} else {
		goto L187
	}
L77:
	;
	v955 = int32(80)
	*(*uint16)(unsafe.Add(mBase, uint32(v439)+12)) = uint16(v955)
	F_MarkBufferDirty(m, v420)
	mBase = m.M
	v958 = m.ExcPending
	if v958 != 0 {
		goto L4
	} else {
		goto L185
	}
L78:
	;
	v927 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v740)+14)))
	v928 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v740)+12)))
	v929 = v927 - v928
	v930 = int32(0)
	if v930 < v929 {
		goto L182
	} else {
		goto L183
	}
L79:
	;
	v885 = int32(80)
	*(*uint16)(unsafe.Add(mBase, uint32(v439)+12)) = uint16(v885)
	F_MarkBufferDirty(m, v420)
	mBase = m.M
	v888 = m.ExcPending
	if v888 != 0 {
		goto L4
	} else {
		goto L176
	}
L80:
	;
	v713 = int32(0)
	F_CheckForSerializableConflictIn(m, v395, v713, v713)
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L4
	} else {
		goto L141
	}
L81:
	;
	F_LockBuffer(m, v420, int32(2))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L4
	} else {
		goto L84
	}
L82:
	;
	v467 = v441
	goto L83
L83:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v390)))
	v470 = int32(0)
	if v470 < v467 {
		goto L90
	} else {
		goto L91
	}
L84:
	;
	v451 = v439 + int32(24)
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v451)))
	if v452 != int32(-1) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v439)+32))
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v390)+12))
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v390)+4))
	if base.Ui32(v456+v457<<(uint(int32(2))%32)) <= base.Ui32(v455) {
		goto L80
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	F_LockBuffer(m, v420, int32(0))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L4
	} else {
		goto L89
	}
L88:
	;
	goto L87
L89:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v390)+4))
	v467 = v465
	goto L83
L90:
	;
	v473 = int32(0)
	v476 = v380
	v477 = v380
	v478 = v473
	v479 = v380
	v482 = v470
	v483 = v380
	v485 = v473
	goto L93
L91:
	;
	v580 = v380
	v581 = v380
	v582 = int32(1)
	v586 = v470
	goto L92
L92:
	;
	if v586 < int32(0) {
		goto L120
	} else {
		goto L121
	}
L93:
	;
	if v482 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v580 = v550
	v581 = v551
	v582 = v555 + int32(1)
	v586 = v574
	goto L92
L95:
	;
	v497 = F_GinNewBuffer(m, v395)
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L4
	} else {
		goto L98
	}
L96:
	;
	v550 = v476
	v551 = v477
	v552 = v479
	v553 = v482
	v554 = v483
	v555 = v485
	goto L97
L97:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v469+v478<<(uint(int32(2))%32))))
	v560 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v559)+6)))
	v569 = v554 + (v560&int32(_a_F_gininsert_6)+int32(7))&int32(_a_F_gininsert_11) + int32(4)
	v571 = base.B2i32(base.Ui32(v569) < base.Ui32(int32(_a_F_gininsert_12)))
	if base.Ui32(v569) < base.Ui32(int32(_a_F_gininsert_12)) {
		goto L112
	} else {
		goto L113
	}
L98:
	;
	if v479 != 0 {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	v550 = v478
	v551 = v547
	v552 = v497
	v553 = v497
	v554 = v548
	v555 = v549
	goto L97
L100:
	;
	if v497 < int32(0) {
		goto L104
	} else {
		goto L105
	}
L101:
	;
	goto L102
L102:
	;
	v527 = int32(0)
	if v497 < v527 {
		goto L109
	} else {
		goto L110
	}
L103:
	;
	v522 = F_writeListPage(m, v395, v479, v469+v476<<(uint(int32(2))%32), v478-v476, v521)
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L4
	} else {
		goto L107
	}
L104:
	;
	v506 = *(*int32)(unsafe.Add(mBase, _c_F_gininsert[4]))
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v506+(v497^int32(-1))<<(uint(int32(6))%32))+16))
	v521 = v512
	goto L103
L105:
	;
	goto L106
L106:
	;
	v514 = *(*int32)(unsafe.Add(mBase, _c_F_gininsert[5]))
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v514+v497<<(uint(int32(6))%32)+int32(-64))+16))
	v521 = v520
	goto L103
L107:
	;
	v547 = v477
	v548 = int32(0)
	v549 = v485 + int32(1)
	goto L99
L108:
	;
	v547 = v546
	v548 = v527
	v549 = v485
	goto L99
L109:
	;
	v531 = *(*int32)(unsafe.Add(mBase, _c_F_gininsert[4]))
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v531+(v497^int32(-1))<<(uint(int32(6))%32))+16))
	v546 = v537
	goto L108
L110:
	;
	goto L111
L111:
	;
	v539 = *(*int32)(unsafe.Add(mBase, _c_F_gininsert[5]))
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v539+v497<<(uint(int32(6))%32)+int32(-64))+16))
	v546 = v545
	goto L108
L112:
	;
	v572 = v569
	goto L114
L113:
	;
	v572 = v554
	goto L114
L114:
	;
	if base.Ui32(v569) < base.Ui32(int32(_a_F_gininsert_12)) {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v574 = v553
	goto L117
L116:
	;
	v574 = int32(0)
	goto L117
L117:
	;
	v575 = v478 + v571
	if v575 < v467 {
		v476 = v550
		v477 = v551
		v478 = v575
		v479 = v552
		v482 = v574
		v483 = v572
		v485 = v555
		goto L93
	} else {
		goto L118
	}
L118:
	;
	goto L94
L119:
	;
	v623 = F_writeListPage(m, v395, v586, v469+v580<<(uint(int32(2))%32), v467-v580, int32(-1))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L4
	} else {
		goto L123
	}
L120:
	;
	v602 = *(*int32)(unsafe.Add(mBase, _c_F_gininsert[4]))
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v602+(v586^int32(-1))<<(uint(int32(6))%32))+16))
	v617 = v608
	goto L119
L121:
	;
	goto L122
L122:
	;
	v610 = *(*int32)(unsafe.Add(mBase, _c_F_gininsert[5]))
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v610+v586<<(uint(int32(6))%32)+int32(-64))+16))
	v617 = v616
	goto L119
L123:
	;
	F_LockBuffer(m, v420, int32(2))
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L4
	} else {
		goto L124
	}
L124:
	;
	v628 = int32(0)
	F_CheckForSerializableConflictIn(m, v395, v628, v628)
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L4
	} else {
		goto L125
	}
L125:
	;
	v633 = v439 + int32(24)
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v633)))
	if v634 == int32(-1) {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v637 = int32(0)
	v638 = int32(_a_F_gininsert_13)
	v640 = *(*int32)(unsafe.Add(mBase, _c_F_gininsert[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_gininsert[6])) = v640 + int32(1)
	*(*int64)(unsafe.Add(mBase, uint32(v439)+40)) = int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v439)+36)) = v582
	*(*int32)(unsafe.Add(mBase, uint32(v439)+32)) = v623
	*(*int32)(unsafe.Add(mBase, uint32(v439)+28)) = v617
	*(*int32)(unsafe.Add(mBase, uint32(v439)+24)) = v581
	if v410 == v637 {
		v939 = v637
		goto L77
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v439)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v387)+84)) = v581
	*(*int32)(unsafe.Add(mBase, uint32(v387)+80)) = v655
	v658 = F_ReadBuffer(m, v395, v655)
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L4
	} else {
		goto L131
	}
L129:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L4
	} else {
		goto L130
	}
L130:
	;
	v869 = v637
	v875 = v633
	v880 = int32(0)
	goto L79
L131:
	;
	F_LockBuffer(m, v658, int32(2))
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L4
	} else {
		goto L132
	}
L132:
	;
	if v658 < int32(0) {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	v681 = int32(_a_F_gininsert_13)
	v683 = *(*int32)(unsafe.Add(mBase, _c_F_gininsert[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_gininsert[6])) = v683 + int32(1)
	v687 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v680)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v680+v687))) = v581
	F_MarkBufferDirty(m, v658)
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L4
	} else {
		goto L137
	}
L134:
	;
	v666 = *(*int32)(unsafe.Add(mBase, _c_F_gininsert[2]))
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v666+(v658^int32(-1))<<(uint(int32(2))%32))))
	v680 = v672
	goto L133
L135:
	;
	goto L136
L136:
	;
	v674 = *(*int32)(unsafe.Add(mBase, _c_F_gininsert[3]))
	v680 = v674 + v658<<(uint(int32(13))%32) + int32(-8192)
	goto L133
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v439)+32)) = v623
	*(*int32)(unsafe.Add(mBase, uint32(v439)+28)) = v617
	v695 = v439 + int32(36)
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v695)))
	*(*int32)(unsafe.Add(mBase, uint32(v695))) = v696 + v582
	v700 = v439 + int32(40)
	v701 = *(*int64)(unsafe.Add(mBase, uint32(v700)))
	*(*int64)(unsafe.Add(mBase, uint32(v700))) = v701 + int64(1)
	if v410 == int32(0) {
		v939 = v658
		goto L77
	} else {
		goto L138
	}
L138:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L4
	} else {
		goto L139
	}
L139:
	;
	F_XLogRegisterBuffer(m, int32(1), v658, int32(8))
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L4
	} else {
		goto L140
	}
L140:
	;
	v869 = v658
	v875 = v633
	v880 = v680
	goto L79
L141:
	;
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v439)+28))
	v718 = F_ReadBuffer(m, v395, v717)
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L4
	} else {
		goto L142
	}
L142:
	;
	F_LockBuffer(m, v718, int32(2))
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L4
	} else {
		goto L143
	}
L143:
	;
	if v718 < int32(0) {
		goto L145
	} else {
		goto L146
	}
L144:
	;
	v741 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v740)+12)))
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v390)+12))
	v743 = F_palloc(m, v742)
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L4
	} else {
		goto L148
	}
L145:
	;
	v726 = *(*int32)(unsafe.Add(mBase, _c_F_gininsert[2]))
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v726+(v718^int32(-1))<<(uint(int32(2))%32))))
	v740 = v732
	goto L144
L146:
	;
	goto L147
L147:
	;
	v734 = *(*int32)(unsafe.Add(mBase, _c_F_gininsert[3]))
	v740 = v734 + v718<<(uint(int32(13))%32) + int32(-8192)
	goto L144
L148:
	;
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v390)+4))
	v746 = int32(_a_F_gininsert_13)
	v748 = *(*int32)(unsafe.Add(mBase, _c_F_gininsert[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_gininsert[6])) = v748 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v387)+88)) = v745
	if v410 != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L4
	} else {
		goto L152
	}
L150:
	;
	goto L151
L151:
	;
	v755 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v740)+16)))
	v756 = v740 + v755
	v757 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v756)+4)))
	v759 = v757 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v756)+4)) = uint16(v759)
	v762 = v439 + int32(40)
	v763 = *(*int64)(unsafe.Add(mBase, uint32(v762)))
	*(*int64)(unsafe.Add(mBase, uint32(v762))) = v763 + int64(1)
	v767 = *(*int32)(unsafe.Add(mBase, uint32(v390)+4))
	if v767 != 0 {
		goto L153
	} else {
		goto L154
	}
L152:
	;
	goto L151
L153:
	;
	v768 = int32(1)
	if base.Ui32(v741) < base.Ui32(int32(25)) {
		goto L156
	} else {
		goto L157
	}
L154:
	;
	goto L155
L155:
	;
	F_MarkBufferDirty(m, v718)
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L4
	} else {
		goto L168
	}
L156:
	;
	v777 = v768
	goto L158
L157:
	;
	v777 = int32(base.Ui32(v741+int32(_a_F_gininsert_14))>>(uint(int32(2))%32)) + v768
	goto L158
L158:
	;
	v778 = v743
	v781 = v380
	v786 = v777
	goto L159
L159:
	;
	v799 = v781 << (uint(int32(2)) % 32)
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v390)))
	v802 = *(*int32)(unsafe.Add(mBase, uint32(v799+v800)))
	v803 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v802)+6)))
	v805 = v803 & int32(_a_F_gininsert_6)
	v809 = F_PageAddItemExtended(m, v740, v802, v805, v786&int32(_a_F_gininsert_5), int32(0))
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L4
	} else {
		goto L161
	}
L160:
	;
	goto L155
L161:
	;
	if v809 == int32(0) {
		goto L63
	} else {
		goto L162
	}
L162:
	;
	v815 = *(*int32)(unsafe.Add(mBase, uint32(v390)))
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v815+v799)))
	if v805 != 0 {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	v822 = v781 + int32(1)
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v390)+4))
	if base.Ui32(v822) < base.Ui32(v823) {
		v778 = v819 + v805
		v781 = v822
		v786 = v786 + int32(1)
		goto L159
	} else {
		goto L167
	}
L164:
	;
	v818 = F__emscripten_memcpy_bulkmem(m, v778, v817, v805)
	mBase = m.M
	v819 = v818
	goto L166
L165:
	;
	v819 = v778
	goto L166
L166:
	;
	goto L163
L167:
	;
	goto L160
L168:
	;
	if v410 == int32(0) {
		goto L78
	} else {
		goto L169
	}
L169:
	;
	F_XLogRegisterBuffer(m, int32(1), v718, int32(8))
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
		goto L4
	} else {
		goto L170
	}
L170:
	;
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v390)+12))
	F_XLogRegisterBufData(m, int32(1), v743, v854)
	mBase = m.M
	v856 = m.ExcPending
	if v856 != 0 {
		goto L4
	} else {
		goto L171
	}
L171:
	;
	v857 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v740)+14)))
	v858 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v740)+12)))
	v859 = v857 - v858
	v860 = int32(0)
	if v860 < v859 {
		goto L173
	} else {
		goto L174
	}
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v439)+32)) = v863
	v869 = v718
	v875 = v451
	v880 = v740
	goto L79
L173:
	;
	v863 = v859
	goto L175
L174:
	;
	v863 = v860
	goto L175
L175:
	;
	goto L172
L176:
	;
	v889 = *(*int64)(unsafe.Add(mBase, uint32(v875)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v387)+72)) = v889
	v893 = *(*int64)(unsafe.Add(mBase, uint32(v875)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v387-int32(-64)))) = v893
	v895 = *(*int64)(unsafe.Add(mBase, uint32(v875)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v387)+56)) = v895
	v897 = *(*int64)(unsafe.Add(mBase, uint32(v875)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v387)+48)) = v897
	v899 = *(*int64)(unsafe.Add(mBase, uint32(v875)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v387)+40)) = v899
	v901 = *(*int64)(unsafe.Add(mBase, uint32(v875)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v387)+32)) = v901
	v903 = *(*int64)(unsafe.Add(mBase, uint32(v875)))
	*(*int64)(unsafe.Add(mBase, uint32(v387)+24)) = v903
	F_XLogRegisterBuffer(m, int32(0), v420, int32(14))
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
		goto L4
	} else {
		goto L177
	}
L177:
	;
	F_XLogRegisterData(m, v387+int32(8), int32(88))
	mBase = m.M
	v913 = m.ExcPending
	if v913 != 0 {
		goto L4
	} else {
		goto L178
	}
L178:
	;
	v916 = F_XLogInsert(m, int32(13), int32(96))
	mBase = m.M
	v917 = m.ExcPending
	if v917 != 0 {
		goto L4
	} else {
		goto L179
	}
L179:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v439))) = base.I64_rotr(v916, int64(32))
	if v869 == int32(0) {
		goto L75
	} else {
		goto L180
	}
L180:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v880)+4)) = uint32(v916)
	v925 = int64(base.Ui64(v916) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v880))) = uint32(v925)
	v965 = v869
	goto L76
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v439)+32)) = v933
	v939 = v718
	goto L77
L182:
	;
	v933 = v929
	goto L184
L183:
	;
	v933 = v930
	goto L184
L184:
	;
	goto L181
L185:
	;
	if v939 == int32(0) {
		goto L75
	} else {
		goto L186
	}
L186:
	;
	v965 = v939
	goto L76
L187:
	;
	goto L75
L188:
	;
	v1011 = *(*int32)(unsafe.Add(mBase, uint32(v439)+36))
	F_UnlockReleaseBuffer(m, v420)
	mBase = m.M
	v1013 = m.ExcPending
	if v1013 != 0 {
		goto L4
	} else {
		goto L193
	}
L189:
	;
	v1004 = *(*int32)(unsafe.Add(mBase, uint32(v1003)+8))
	if v1004 != int32(-1) {
		v1010 = v1004
		goto L188
	} else {
		goto L192
	}
L190:
	;
	goto L191
L191:
	;
	v1009 = *(*int32)(unsafe.Add(mBase, _c_F_gininsert[7]))
	v1010 = v1009
	goto L188
L192:
	;
	goto L191
L193:
	;
	v1014 = int32(_a_F_gininsert_13)
	v1016 = *(*int32)(unsafe.Add(mBase, _c_F_gininsert[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_gininsert[6])) = v1016 - int32(1)
	if base.Ui32(v1011*int32(_a_F_gininsert_10)) <= base.Ui32(v1010<<(uint(int32(10))%32)) {
		goto L64
	} else {
		goto L194
	}
L194:
	;
	v1025 = int32(0)
	F_ginInsertCleanup(m, v43, v1025, int32(1), v1025, v1025)
	mBase = m.M
	v1030 = m.ExcPending
	if v1030 != 0 {
		goto L4
	} else {
		goto L195
	}
L195:
	;
	goto L64
L196:
	;
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(v395)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v387))) = v1058 + int32(4)
	F_errmsg_internal(m, int32(_a_F_gininsert_15), v387)
	mBase = m.M
	v1064 = m.ExcPending
	if v1064 != 0 {
		goto L4
	} else {
		goto L197
	}
L197:
	;
	F_errfinish(m, int32(_a_F_gininsert_8), int32(391), int32(_a_F_gininsert_16))
	mBase = m.M
	v1069 = m.ExcPending
	if v1069 != 0 {
		goto L4
	} else {
		goto L198
	}
L198:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L199:
	;
	m.G0 = v23 + int32(32)
	return int32(0)
}
func F_gintuple_get_key(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v11 == int32(1) {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v18 = F_index_getattr_1(m, l1, int32(1), v15, v9+int32(14))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v39 = v18
			v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+14)))
			if v40 == int32(1) {
				v45 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
				if int32(0) <= v45 {
					v48 = int32(8)
				} else {
					v48 = int32(16)
				}
				v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
				if v52 != 0 {
					v53 = int32(0)
				} else {
					v53 = int32(2)
				}
				v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v48+v53))))
				v56 = v55
			} else {
				v56 = int32(0)
			}
			*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v56)
			m.G0 = v9 + int32(16)
			return v39
		}
	} else {
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v27 = F_index_getattr_1(m, l1, int32(1), v24, v9+int32(15))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v27&int32(_a_F_gintuple_get_key_0)<<(uint(int32(2))%32)+l0)+8))
			v37 = F_index_getattr_1(m, l1, int32(2), v34, v9+int32(14))
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int32(0)
			} else {
				v39 = v37
				v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+14)))
				if v40 == int32(1) {
					v45 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
					if int32(0) <= v45 {
						v48 = int32(8)
					} else {
						v48 = int32(16)
					}
					v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
					if v52 != 0 {
						v53 = int32(0)
					} else {
						v53 = int32(2)
					}
					v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v48+v53))))
					v56 = v55
				} else {
					v56 = int32(0)
				}
				*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v56)
				m.G0 = v9 + int32(16)
				return v39
			}
		}
	}
}
func F_gistbufferinginserttuples(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v129 int32
	_ = v129
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v225 int32
	_ = v225
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v405 int32
	_ = v405
	var v411 int32
	_ = v411
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v444 int32
	_ = v444
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v532 int32
	_ = v532
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v547 int32
	_ = v547
	var v552 int32
	_ = v552
	var v560 int32
	_ = v560
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v593 int32
	_ = v593
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v621 int64
	_ = v621
	var v623 int64
	_ = v623
	var v625 int64
	_ = v625
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v651 int32
	_ = v651
	var v670 int32
	_ = v670
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v678 int32
	_ = v678
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v687 int32
	_ = v687
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v724 int32
	_ = v724
	var v729 int32
	_ = v729
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v765 int32
	_ = v765
	var v771 int32
	_ = v771
	var v774 int32
	_ = v774
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v805 int32
	_ = v805
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v834 int32
	_ = v834
	var v837 int32
	_ = v837
	var v844 int32
	_ = v844
	var v845 float32
	_ = v845
	var v846 int32
	_ = v846
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v850 float32
	_ = v850
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v873 int32
	_ = v873
	var v875 int32
	_ = v875
	var v877 int32
	_ = v877
	var v882 int32
	_ = v882
	var v887 int32
	_ = v887
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v953 int32
	_ = v953
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v983 int32
	_ = v983
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v998 int32
	_ = v998
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1029 int32
	_ = v1029
	var v1035 int32
	_ = v1035
	var v1037 int32
	_ = v1037
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1048 int32
	_ = v1048
	var v1054 int32
	_ = v1054
	var v1056 int32
	_ = v1056
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1065 int32
	_ = v1065
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1076 int32
	_ = v1076
	var v1080 int32
	_ = v1080
	var v1086 int32
	_ = v1086
	var v1088 int32
	_ = v1088
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1099 int32
	_ = v1099
	var v1105 int32
	_ = v1105
	var v1107 int32
	_ = v1107
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1118 int32
	_ = v1118
	var v1135 int32
	_ = v1135
	var v1158 int32
	_ = v1158
	var v1161 int32
	_ = v1161
	var v1162 int32
	_ = v1162
	var v1165 int32
	_ = v1165
	var v1168 int32
	_ = v1168
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1203 int32
	_ = v1203
	var v1205 int32
	_ = v1205
	var v1207 int32
	_ = v1207
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1242 int32
	_ = v1242
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1246 int32
	_ = v1246
	var v1270 int32
	_ = v1270
	v9 = int32(0)
	v24 = m.G0
	v26 = v24 - int32(80)
	m.G0 = v26
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+64)) = int32(-1)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v42 = F_gistplacetopage(m, v31, v32, v33, l1, l3, l4, l5, v26-int32(-64), v9, v26+int32(68), v9, v40, int32(1))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v26)+68))
	if v332 != 0 {
		goto L49
	} else {
		goto L50
	}
L2:
	;
	return int32(0)
L3:
	;
	if v42 == int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if l1 < int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v66 != 0 {
		goto L1
	} else {
		goto L9
	}
L6:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_gistbufferinginserttuples[0]))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v51+(l1^int32(-1))<<(uint(int32(6))%32))+16))
	v66 = v57
	goto L5
L7:
	;
	goto L8
L8:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_gistbufferinginserttuples[1]))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v59+l1<<(uint(int32(6))%32)+int32(-64))+16))
	v66 = v65
	goto L5
L9:
	;
	if l1 < int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v28)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+60)) = v85 + int32(1)
	v91 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L2
	} else {
		goto L14
	}
L11:
	;
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_gistbufferinginserttuples[2]))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v70+(l1^int32(-1))<<(uint(int32(2))%32))))
	v84 = v76
	goto L10
L12:
	;
	goto L13
L13:
	;
	v78 = *(*int32)(unsafe.Add(mBase, _c_F_gistbufferinginserttuples[3]))
	v84 = v78 + l1<<(uint(int32(13))%32) + int32(-8192)
	goto L10
L14:
	;
	if v91 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v28)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+48)) = v93
	F_errmsg_internal(m, int32(_a_F_gistbufferinginserttuples_0), v26+int32(48))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L2
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v28)+60))
	if v105 < int32(2) {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	F_errfinish(m, int32(_a_F_gistbufferinginserttuples_1), int32(1090), int32(_a_F_gistbufferinginserttuples_2))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v84)+12)))
	if base.Ui32(v108) < base.Ui32(int32(25)) {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v112 = v108 + int32(_a_F_gistbufferinginserttuples_3)
	if v112&int32(_a_F_gistbufferinginserttuples_4) == int32(0) {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v129 = int32(1)
	goto L23
L23:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v129<<(uint(int32(2))%32)+(v84+int32(24))-int32(4))))
	v156 = v84 + v153&int32(_a_F_gistbufferinginserttuples_5)
	v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v156))))
	v160 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v156)+2)))
	v161 = v157<<(uint(int32(16))%32) | v160
	v162 = F_ReadBuffer(m, v147, v161)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L2
	} else {
		goto L25
	}
L24:
	;
	goto L1
L25:
	;
	F_LockBuffer(m, v162, int32(1))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L2
	} else {
		goto L26
	}
L26:
	;
	if v162 < int32(0) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	if v162 < int32(0) {
		goto L33
	} else {
		goto L34
	}
L28:
	;
	v170 = *(*int32)(unsafe.Add(mBase, _c_F_gistbufferinginserttuples[0]))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v170+(v162^int32(-1))<<(uint(int32(6))%32))+16))
	v185 = v176
	goto L27
L29:
	;
	goto L30
L30:
	;
	v178 = *(*int32)(unsafe.Add(mBase, _c_F_gistbufferinginserttuples[1]))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v178+v162<<(uint(int32(6))%32)+int32(-64))+16))
	v185 = v184
	goto L27
L31:
	;
	F_UnlockReleaseBuffer(m, v162)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L2
	} else {
		goto L42
	}
L32:
	;
	v204 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v203)+12)))
	if base.Ui32(v204) < base.Ui32(int32(25)) {
		goto L31
	} else {
		goto L36
	}
L33:
	;
	v189 = *(*int32)(unsafe.Add(mBase, _c_F_gistbufferinginserttuples[2]))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v189+(v162^int32(-1))<<(uint(int32(2))%32))))
	v203 = v195
	goto L32
L34:
	;
	goto L35
L35:
	;
	v197 = *(*int32)(unsafe.Add(mBase, _c_F_gistbufferinginserttuples[3]))
	v203 = v197 + v162<<(uint(int32(13))%32) + int32(-8192)
	goto L32
L36:
	;
	v208 = v204 + int32(_a_F_gistbufferinginserttuples_3)
	if v208&int32(_a_F_gistbufferinginserttuples_4) == int32(0) {
		goto L31
	} else {
		goto L37
	}
L37:
	;
	v225 = int32(1)
	goto L38
L38:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v225<<(uint(int32(2))%32)+(v203+int32(24))-int32(4))))
	v251 = v203 + v248&int32(_a_F_gistbufferinginserttuples_5)
	v252 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v251))))
	v255 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v251)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+76)) = v252<<(uint(int32(16))%32) | v255
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v264 = F_hash_search(m, v258, v26+int32(76), int32(1), v26+int32(75))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L2
	} else {
		goto L40
	}
L39:
	;
	goto L31
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v264)+4)) = v185
	if v225 != int32(base.Ui32(v208)>>(uint(int32(2))%32))&int32(_a_F_gistbufferinginserttuples_6) {
		v225 = v225 + int32(1)
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+76)) = v161
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v302 = F_hash_search(m, v296, v26+int32(76), int32(1), v26+int32(75))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L2
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v302)+4)) = int32(0)
	if v129 != int32(base.Ui32(v112)>>(uint(int32(2))%32))&int32(_a_F_gistbufferinginserttuples_6) {
		v129 = v129 + int32(1)
		goto L23
	} else {
		goto L44
	}
L44:
	;
	goto L24
L45:
	;
	v1270 = *(*int32)(unsafe.Add(mBase, uint32(v26)+64))
	m.G0 = v26 + int32(80)
	return v1270
L46:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v577 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v26)+68))
	v579 = int32(0)
	v580 = m.G0
	v582 = v580 - int32(720)
	m.G0 = v582
	if l2 == v579 {
		goto L95
	} else {
		goto L96
	}
L47:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L2
	} else {
		goto L92
	}
L48:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L2
	} else {
		goto L89
	}
L49:
	;
	if l1 < int32(0) {
		goto L53
	} else {
		goto L54
	}
L50:
	;
	goto L51
L51:
	;
	F_UnlockReleaseBuffer(m, l1)
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L2
	} else {
		goto L88
	}
L52:
	;
	if int32(0) < l2 {
		goto L57
	} else {
		goto L58
	}
L53:
	;
	v336 = *(*int32)(unsafe.Add(mBase, _c_F_gistbufferinginserttuples[0]))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v336+(l1^int32(-1))<<(uint(int32(6))%32))+16))
	v351 = v342
	goto L52
L54:
	;
	goto L55
L55:
	;
	v344 = *(*int32)(unsafe.Add(mBase, _c_F_gistbufferinginserttuples[1]))
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v344+l1<<(uint(int32(6))%32)+int32(-64))+16))
	v351 = v350
	goto L52
L56:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v371 = F_ReadBuffer(m, v370, v369)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L2
	} else {
		goto L64
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+76)) = v351
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v361 = F_hash_search(m, v355, v26+int32(76), int32(0), v26+int32(75))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L2
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	if l6 == int32(-1) {
		goto L47
	} else {
		goto L62
	}
L60:
	;
	v363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+75)))
	if v363 == int32(0) {
		goto L48
	} else {
		goto L61
	}
L61:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v361)+4))
	v369 = v366
	goto L56
L62:
	;
	v369 = l6
	goto L56
L63:
	;
	F_LockBuffer(m, v371, int32(2))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L2
	} else {
		goto L68
	}
L64:
	;
	if v371 < int32(0) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v376 = *(*int32)(unsafe.Add(mBase, _c_F_gistbufferinginserttuples[2]))
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v376+(v371^int32(-1))<<(uint(int32(2))%32))))
	v390 = v382
	goto L63
L66:
	;
	goto L67
L67:
	;
	v384 = *(*int32)(unsafe.Add(mBase, _c_F_gistbufferinginserttuples[3]))
	v390 = v384 + v371<<(uint(int32(13))%32) + int32(-8192)
	goto L63
L68:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_gistcheckpage(m, v394, v371)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L2
	} else {
		goto L69
	}
L69:
	;
	v397 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v390)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v397) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v405 = int32(base.Ui32(v397+int32(_a_F_gistbufferinginserttuples_3)) >> (uint(int32(2)) % 32))
	goto L72
L71:
	;
	v405 = int32(0)
	goto L72
L72:
	;
	if l6 == int32(-1) {
		goto L75
	} else {
		goto L76
	}
L73:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L2
	} else {
		goto L85
	}
L74:
	;
	v444 = int32(1)
	goto L81
L75:
	;
	if v405&int32(_a_F_gistbufferinginserttuples_6) == int32(0) {
		goto L73
	} else {
		goto L80
	}
L76:
	;
	if v369 != l6 {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v411 = int32(_a_F_gistbufferinginserttuples_6)
	if base.Ui32(v405&v411) <= base.Ui32((l7-int32(1))&v411) {
		goto L75
	} else {
		goto L78
	}
L78:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(l7<<(uint(int32(2))%32)+v390)+20))
	v422 = v390 + v419&int32(_a_F_gistbufferinginserttuples_5)
	v423 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v422))))
	v426 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v422)+2)))
	if v423<<(uint(int32(16))%32)|v426 != v351 {
		goto L74
	} else {
		goto L79
	}
L79:
	;
	v560 = l7
	goto L46
L80:
	;
	goto L74
L81:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v444&int32(_a_F_gistbufferinginserttuples_6)<<(uint(int32(2))%32)+(v390+int32(24))-int32(4))))
	v470 = v390 + v467&int32(_a_F_gistbufferinginserttuples_5)
	v471 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v470))))
	v474 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v470)+2)))
	if v471<<(uint(int32(16))%32)|v474 == v351 {
		v560 = v444
		goto L46
	} else {
		goto L83
	}
L82:
	;
	goto L73
L83:
	;
	v478 = v444 + int32(1)
	v479 = int32(_a_F_gistbufferinginserttuples_6)
	if base.Ui32(v478&v479) <= base.Ui32(v405&v479) {
		v444 = v478
		goto L81
	} else {
		goto L84
	}
L84:
	;
	goto L82
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v351
	F_errmsg_internal(m, int32(_a_F_gistbufferinginserttuples_7), v26)
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L2
	} else {
		goto L86
	}
L86:
	;
	F_errfinish(m, int32(_a_F_gistbufferinginserttuples_1), int32(1289), int32(_a_F_gistbufferinginserttuples_8))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L2
	} else {
		goto L87
	}
L87:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L88:
	;
	goto L45
L89:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v26)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v526
	F_errmsg_internal(m, int32(_a_F_gistbufferinginserttuples_9), v26+int32(16))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L2
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(_a_F_gistbufferinginserttuples_1), int32(1578), int32(_a_F_gistbufferinginserttuples_10))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L2
	} else {
		goto L91
	}
L91:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v351
	F_errmsg_internal(m, int32(_a_F_gistbufferinginserttuples_11), v26+int32(32))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L2
	} else {
		goto L93
	}
L93:
	;
	F_errfinish(m, int32(_a_F_gistbufferinginserttuples_1), int32(1245), int32(_a_F_gistbufferinginserttuples_8))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L2
	} else {
		goto L94
	}
L94:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L95:
	;
	m.G0 = v582 + int32(720)
	v980 = *(*int32)(unsafe.Add(mBase, uint32(v26)+68))
	if v980 != 0 {
		goto L157
	} else {
		goto L158
	}
L96:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v28)+32))
	v587 = base.I32_rem_s(l2, v586)
	if v587 != 0 {
		goto L95
	} else {
		goto L97
	}
L97:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v28)+60))
	if l2 == v588 {
		goto L95
	} else {
		goto L98
	}
L98:
	;
	if l1 < int32(0) {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v582)+712)) = v608
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v28)+24))
	v616 = F_hash_search(m, v610, v582+int32(712), int32(0), v582+int32(719))
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L2
	} else {
		goto L103
	}
L100:
	;
	v593 = *(*int32)(unsafe.Add(mBase, _c_F_gistbufferinginserttuples[0]))
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v593+(l1^int32(-1))<<(uint(int32(6))%32))+16))
	v608 = v599
	goto L99
L101:
	;
	goto L102
L102:
	;
	v601 = *(*int32)(unsafe.Add(mBase, _c_F_gistbufferinginserttuples[1]))
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v601+l1<<(uint(int32(6))%32)+int32(-64))+16))
	v608 = v607
	goto L99
L103:
	;
	v618 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v582)+719)))
	if v618 != int32(1) {
		goto L95
	} else {
		goto L104
	}
L104:
	;
	v621 = *(*int64)(unsafe.Add(mBase, uint32(v616)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v582)+152)) = v621
	v623 = *(*int64)(unsafe.Add(mBase, uint32(v616)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v582)+144)) = v623
	v625 = *(*int64)(unsafe.Add(mBase, uint32(v616)))
	*(*int64)(unsafe.Add(mBase, uint32(v582)+136)) = v625
	v627 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v582)+153)) = uint8(v627)
	v629 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v616)+12)) = v629
	*(*int64)(unsafe.Add(mBase, uint32(v616)+4)) = int64(-4294967296)
	if v578 == v629 {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	v738 = F_gistPopItupFromNodeBuffer(m, v28, v582+int32(136), v582+int32(708))
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L2
	} else {
		goto L121
	}
L106:
	;
	v636 = F_palloc(m, int32(0))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L2
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v578)+4))
	v641 = F_palloc(m, v638*int32(552))
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L2
	} else {
		goto L110
	}
L109:
	;
	v724 = v636
	v729 = v579
	goto L105
L110:
	;
	v643 = int32(0)
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v578)+4))
	if v644 <= v643 {
		v724 = v641
		v729 = v638
		goto L105
	} else {
		goto L111
	}
L111:
	;
	v651 = v643
	goto L112
L112:
	;
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v578)+12))
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v670+v651<<(uint(int32(2))%32))))
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v674)+4))
	v678 = v641 + v651*int32(552)
	F_gistDeCompressAtt(m, v576, v577, v675, v678, v678+int32(512))
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L2
	} else {
		goto L114
	}
L113:
	;
	v724 = v641
	v729 = v638
	goto L105
L114:
	;
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v674)))
	if v683 < int32(0) {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	v703 = F_gistGetNodeBuffer(m, v28, v702, l2)
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L2
	} else {
		goto L119
	}
L116:
	;
	v687 = *(*int32)(unsafe.Add(mBase, _c_F_gistbufferinginserttuples[0]))
	v693 = *(*int32)(unsafe.Add(mBase, uint32(v687+(v683^int32(-1))<<(uint(int32(6))%32))+16))
	v702 = v693
	goto L115
L117:
	;
	goto L118
L118:
	;
	v695 = *(*int32)(unsafe.Add(mBase, _c_F_gistbufferinginserttuples[1]))
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v695+v683<<(uint(int32(6))%32)+int32(-64))+16))
	v702 = v701
	goto L115
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v678)+544)) = v674
	*(*int32)(unsafe.Add(mBase, uint32(v678)+548)) = v703
	v708 = v651 + int32(1)
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v578)+4))
	if v708 < v709 {
		v651 = v708
		goto L112
	} else {
		goto L120
	}
L120:
	;
	goto L113
L121:
	;
	if v738 != 0 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	goto L125
L123:
	;
	goto L124
L124:
	;
	F_pfree(m, v724)
	mBase = m.M
	v953 = m.ExcPending
	if v953 != 0 {
		goto L2
	} else {
		goto L156
	}
L125:
	;
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v582)+708))
	F_gistDeCompressAtt(m, v576, v577, v765, v582+int32(192), v582+int32(160))
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L2
	} else {
		goto L127
	}
L126:
	;
	goto L124
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v582))) = int32(-1082130432)
	v774 = int32(0)
	if v729 <= int32(0) {
		v887 = v774
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v909 = v724 + v887*int32(552)
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v909)+548))
	F_gistPushItupToNodeBuffer(m, v28, v910, v765)
	mBase = m.M
	v912 = m.ExcPending
	if v912 != 0 {
		goto L2
	} else {
		goto L148
	}
L129:
	;
	v777 = v774
	v779 = v774
	goto L130
L130:
	;
	v799 = *(*int32)(unsafe.Add(mBase, uint32(v577)+192))
	v800 = int32(*(*int16)(unsafe.Add(mBase, uint32(v799)+10)))
	if v800 <= int32(0) {
		v887 = v779
		goto L128
	} else {
		goto L132
	}
L131:
	;
	v887 = v877
	goto L128
L132:
	;
	v805 = v724 + v777*int32(552)
	v813 = v779
	v814 = int32(1)
	v815 = int32(0)
	goto L134
L133:
	;
	v882 = v777 + int32(1)
	if v882 != v729 {
		v777 = v882
		v779 = v877
		goto L130
	} else {
		goto L147
	}
L134:
	;
	v834 = v815 << (uint(int32(4)) % 32)
	v837 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v815+(v805+int32(512))))))
	v844 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v582+int32(160)+v815))))
	v845 = F_gistpenalty(m, v576, v815, v805+v834, v837, v582+int32(192)+v834, v844)
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L2
	} else {
		goto L136
	}
L135:
	;
	if v873 != 0 {
		v887 = v867
		goto L128
	} else {
		goto L146
	}
L136:
	;
	v848 = v815 << (uint(int32(2)) % 32)
	v849 = v582 + v848
	v850 = *(*float32)(unsafe.Add(mBase, uint32(v849)))
	if base.F32_lt(v850, float32(0))|base.F32_lt(v845, v850) != 0 {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	v873 = base.B2i32(base.F32_gt(v845, float32(0)) == int32(0)) & v814
	v875 = v815 + int32(1)
	if v875 < v868 {
		v813 = v867
		v814 = v873
		v815 = v875
		goto L134
	} else {
		goto L145
	}
L138:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v849))) = v845
	v856 = *(*int32)(unsafe.Add(mBase, uint32(v577)+192))
	v857 = int32(*(*int16)(unsafe.Add(mBase, uint32(v856)+10)))
	if v857-int32(1) <= v815 {
		goto L141
	} else {
		goto L142
	}
L139:
	;
	goto L140
L140:
	;
	if base.F32_ne(v845, v850) != 0 {
		v877 = v813
		goto L133
	} else {
		goto L144
	}
L141:
	;
	v867 = v777
	v868 = v857
	goto L137
L142:
	;
	goto L143
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v582+v848)+4)) = int32(-1082130432)
	v867 = v777
	v868 = v857
	goto L137
L144:
	;
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v577)+192))
	v866 = int32(*(*int16)(unsafe.Add(mBase, uint32(v865)+10)))
	v867 = v813
	v868 = v866
	goto L137
L145:
	;
	goto L135
L146:
	;
	v877 = v867
	goto L133
L147:
	;
	goto L131
L148:
	;
	v913 = *(*int32)(unsafe.Add(mBase, uint32(v909)+544))
	v914 = *(*int32)(unsafe.Add(mBase, uint32(v913)+4))
	v915 = F_gistgetadjusted(m, v577, v914, v765, v576)
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L2
	} else {
		goto L149
	}
L149:
	;
	if v915 != 0 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	F_gistDeCompressAtt(m, v576, v577, v915, v909, v909+int32(512))
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L2
	} else {
		goto L153
	}
L151:
	;
	goto L152
L152:
	;
	v927 = F_gistPopItupFromNodeBuffer(m, v28, v582+int32(136), v582+int32(708))
	mBase = m.M
	v928 = m.ExcPending
	if v928 != 0 {
		goto L2
	} else {
		goto L154
	}
L153:
	;
	v921 = *(*int32)(unsafe.Add(mBase, uint32(v909)+544))
	*(*int32)(unsafe.Add(mBase, uint32(v921)+4)) = v915
	goto L152
L154:
	;
	if v927 != 0 {
		goto L125
	} else {
		goto L155
	}
L155:
	;
	goto L126
L156:
	;
	goto L95
L157:
	;
	v981 = *(*int32)(unsafe.Add(mBase, uint32(v980)+4))
	v983 = v981
	goto L159
L158:
	;
	v983 = int32(0)
	goto L159
L159:
	;
	v986 = F_palloc(m, v983<<(uint(int32(2))%32))
	mBase = m.M
	v987 = m.ExcPending
	if v987 != 0 {
		goto L2
	} else {
		goto L160
	}
L160:
	;
	v988 = *(*int32)(unsafe.Add(mBase, uint32(v26)+68))
	if v988 == int32(0) {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v1242 = F_gistbufferinginserttuples(m, l0, v371, l2+int32(1), v986, v983, v560&int32(_a_F_gistbufferinginserttuples_6), int32(-1), int32(0))
	mBase = m.M
	v1243 = m.ExcPending
	if v1243 != 0 {
		goto L2
	} else {
		goto L194
	}
L162:
	;
	v991 = int32(0)
	v992 = *(*int32)(unsafe.Add(mBase, uint32(v988)+4))
	if v992 <= v991 {
		goto L161
	} else {
		goto L163
	}
L163:
	;
	v998 = v991
	goto L164
L164:
	;
	v1021 = v998 << (uint(int32(2)) % 32)
	v1022 = *(*int32)(unsafe.Add(mBase, uint32(v988)+12))
	v1024 = *(*int32)(unsafe.Add(mBase, uint32(v1021+v1022)))
	if l2 <= int32(0) {
		goto L166
	} else {
		goto L167
	}
L165:
	;
	goto L161
L166:
	;
	v1203 = *(*int32)(unsafe.Add(mBase, uint32(v1024)))
	F_UnlockReleaseBuffer(m, v1203)
	mBase = m.M
	v1205 = m.ExcPending
	if v1205 != 0 {
		goto L2
	} else {
		goto L192
	}
L167:
	;
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(v1024)))
	if v1025 < int32(0) {
		goto L169
	} else {
		goto L170
	}
L168:
	;
	if v371 < int32(0) {
		goto L173
	} else {
		goto L174
	}
L169:
	;
	v1029 = *(*int32)(unsafe.Add(mBase, _c_F_gistbufferinginserttuples[0]))
	v1035 = *(*int32)(unsafe.Add(mBase, uint32(v1029+(v1025^int32(-1))<<(uint(int32(6))%32))+16))
	v1044 = v1035
	goto L168
L170:
	;
	goto L171
L171:
	;
	v1037 = *(*int32)(unsafe.Add(mBase, _c_F_gistbufferinginserttuples[1]))
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(v1037+v1025<<(uint(int32(6))%32)+int32(-64))+16))
	v1044 = v1043
	goto L168
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+76)) = v1044
	v1065 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v1071 = F_hash_search(m, v1065, v26+int32(76), int32(1), v26+int32(75))
	mBase = m.M
	v1072 = m.ExcPending
	if v1072 != 0 {
		goto L2
	} else {
		goto L176
	}
L173:
	;
	v1048 = *(*int32)(unsafe.Add(mBase, _c_F_gistbufferinginserttuples[0]))
	v1054 = *(*int32)(unsafe.Add(mBase, uint32(v1048+(v371^int32(-1))<<(uint(int32(6))%32))+16))
	v1063 = v1054
	goto L172
L174:
	;
	goto L175
L175:
	;
	v1056 = *(*int32)(unsafe.Add(mBase, _c_F_gistbufferinginserttuples[1]))
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(v1056+v371<<(uint(int32(6))%32)+int32(-64))+16))
	v1063 = v1062
	goto L172
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1071)+4)) = v1063
	if l2 == int32(1) {
		goto L166
	} else {
		goto L177
	}
L177:
	;
	v1076 = *(*int32)(unsafe.Add(mBase, uint32(v1024)))
	if v1076 < int32(0) {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	if v1076 < int32(0) {
		goto L183
	} else {
		goto L184
	}
L179:
	;
	v1080 = *(*int32)(unsafe.Add(mBase, _c_F_gistbufferinginserttuples[0]))
	v1086 = *(*int32)(unsafe.Add(mBase, uint32(v1080+(v1076^int32(-1))<<(uint(int32(6))%32))+16))
	v1095 = v1086
	goto L178
L180:
	;
	goto L181
L181:
	;
	v1088 = *(*int32)(unsafe.Add(mBase, _c_F_gistbufferinginserttuples[1]))
	v1094 = *(*int32)(unsafe.Add(mBase, uint32(v1088+v1076<<(uint(int32(6))%32)+int32(-64))+16))
	v1095 = v1094
	goto L178
L182:
	;
	v1114 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1113)+12)))
	if base.Ui32(v1114) < base.Ui32(int32(25)) {
		goto L166
	} else {
		goto L186
	}
L183:
	;
	v1099 = *(*int32)(unsafe.Add(mBase, _c_F_gistbufferinginserttuples[2]))
	v1105 = *(*int32)(unsafe.Add(mBase, uint32(v1099+(v1076^int32(-1))<<(uint(int32(2))%32))))
	v1113 = v1105
	goto L182
L184:
	;
	goto L185
L185:
	;
	v1107 = *(*int32)(unsafe.Add(mBase, _c_F_gistbufferinginserttuples[3]))
	v1113 = v1107 + v1076<<(uint(int32(13))%32) + int32(-8192)
	goto L182
L186:
	;
	v1118 = v1114 + int32(_a_F_gistbufferinginserttuples_3)
	if v1118&int32(_a_F_gistbufferinginserttuples_4) == int32(0) {
		goto L166
	} else {
		goto L187
	}
L187:
	;
	v1135 = int32(1)
	goto L188
L188:
	;
	v1158 = *(*int32)(unsafe.Add(mBase, uint32(v1135<<(uint(int32(2))%32)+(v1113+int32(24))-int32(4))))
	v1161 = v1113 + v1158&int32(_a_F_gistbufferinginserttuples_5)
	v1162 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1161))))
	v1165 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1161)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+76)) = v1162<<(uint(int32(16))%32) | v1165
	v1168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v1174 = F_hash_search(m, v1168, v26+int32(76), int32(1), v26+int32(75))
	mBase = m.M
	v1175 = m.ExcPending
	if v1175 != 0 {
		goto L2
	} else {
		goto L190
	}
L189:
	;
	goto L166
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1174)+4)) = v1095
	if v1135 != int32(base.Ui32(v1118)>>(uint(int32(2))%32))&int32(_a_F_gistbufferinginserttuples_6) {
		v1135 = v1135 + int32(1)
		goto L188
	} else {
		goto L191
	}
L191:
	;
	goto L189
L192:
	;
	v1207 = *(*int32)(unsafe.Add(mBase, uint32(v1024)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v986+v1021))) = v1207
	v1210 = v998 + int32(1)
	v1211 = *(*int32)(unsafe.Add(mBase, uint32(v988)+4))
	if v1210 < v1211 {
		v998 = v1210
		goto L164
	} else {
		goto L193
	}
L193:
	;
	goto L165
L194:
	;
	v1244 = *(*int32)(unsafe.Add(mBase, uint32(v26)+68))
	F_list_free_deep(m, v1244)
	mBase = m.M
	v1246 = m.ExcPending
	if v1246 != 0 {
		goto L2
	} else {
		goto L195
	}
L195:
	;
	goto L45
}
func F_gistcanreturn(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	v4 = int32(1)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v6 = int32(*(*int16)(unsafe.Add(mBase, uint32(v5)+10)))
	if v6 < l1 {
		v44 = v4
	} else {
		v8 = base.I32_extend16_s(l1)
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
		v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+6)))
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v10+v12*(v8-int32(1))<<(uint(int32(2))%32)+int32(36)-int32(4))))
		if v24 != 0 {
			v44 = v4
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
			v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27)+6)))
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v26+v28*(v8-int32(1))<<(uint(int32(2))%32)+int32(12)-int32(4))))
			v44 = base.B2i32(v40 == int32(0))
		}
	}
	return v44
}
func F_gisthandler(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v25 int32
	_ = v25
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v61 int64
	_ = v61
	v3 = F_palloc0(m, int32(140))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v3)+8)) = int64(16777226)
		*(*int64)(unsafe.Add(mBase, uint32(v3))) = int64(3377699720528310)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+108)) = int32(76)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+104)) = int32(77)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+100)) = int32(78)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+96)) = int32(79)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+92)) = int32(80)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+88)) = int32(81)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+84)) = int32(82)
		v25 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+80)) = v25
		*(*int32)(unsafe.Add(mBase, uint32(v3)+76)) = int32(83)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+72)) = int32(84)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+68)) = v25
		*(*int32)(unsafe.Add(mBase, uint32(v3)+64)) = int32(85)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+60)) = int32(86)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+56)) = int32(87)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+52)) = int32(88)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+48)) = v25
		*(*int32)(unsafe.Add(mBase, uint32(v3)+44)) = int32(89)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+40)) = int32(90)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+36)) = int32(91)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+32)) = v25
		v51 = int32(768)
		*(*uint16)(unsafe.Add(mBase, uint32(v3)+28)) = uint16(v51)
		*(*int64)(unsafe.Add(mBase, uint32(v3)+20)) = int64(281474993553665)
		*(*uint8)(unsafe.Add(mBase, uint32(v3)+19)) = uint8(v25)
		v57 = int32(257)
		*(*uint16)(unsafe.Add(mBase, uint32(v3)+17)) = uint16(v57)
		*(*uint8)(unsafe.Add(mBase, uint32(v3)+16)) = uint8(v25)
		v61 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v3)+128)) = v61
		*(*int64)(unsafe.Add(mBase, uint32(v3)+120)) = v61
		*(*int64)(unsafe.Add(mBase, uint32(v3)+112)) = v61
		*(*int32)(unsafe.Add(mBase, uint32(v3)+136)) = int32(92)
		return v3
	}
}
func F_gistrescan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v90 int32
	_ = v90
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v199 int32
	_ = v199
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int64
	_ = v268
	var v270 int64
	_ = v270
	var v272 int32
	_ = v272
	var v274 int64
	_ = v274
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v309 int32
	_ = v309
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v342 int32
	_ = v342
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v429 int64
	_ = v429
	var v431 int64
	_ = v431
	var v433 int32
	_ = v433
	var v435 int64
	_ = v435
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v461 int32
	_ = v461
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v493 int32
	_ = v493
	var v498 int32
	_ = v498
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	if v17 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v35 != int32(1) {
		goto L9
	} else {
		goto L10
	}
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if v20 == v22 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v28 = F_AllocSetContextCreateInternal(m, v22, int32(_a_F_gistrescan_0), int32(0), int32(_a_F_gistrescan_1), int32(_a_F_gistrescan_2))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	F_MemoryContextReset(m, v20)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L6
	} else {
		goto L8
	}
L6:
	;
	return
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v28
	goto L1
L8:
	;
	goto L1
L9:
	;
	v165 = int32(_a_F_gistrescan_3)
	v166 = *(*int32)(unsafe.Add(mBase, _c_F_gistrescan[0]))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	*(*int32)(unsafe.Add(mBase, _c_F_gistrescan[0])) = v168
	v171 = F_pairingheap_allocate(m, int32(119), l0)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L6
	} else {
		goto L28
	}
L10:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v38 != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+192))
	v41 = int32(*(*int16)(unsafe.Add(mBase, uint32(v40)+10)))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v43 = int32(*(*int16)(unsafe.Add(mBase, uint32(v42)+120)))
	v44 = F_CreateTemplateTupleDesc(m, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+16)) = v44
	if int32(0) < v41 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v51 = int32(1)
	v56 = v51
	goto L16
L14:
	;
	v90 = int32(1)
	goto L15
L15:
	;
	if v90 <= v43 {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+16))
	v68 = int32(0)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+212))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v70+v56<<(uint(int32(2))%32)-int32(4))))
	F_TupleDescInitEntry(m, v66, base.I32_extend16_s(v56), v68, v76, int32(-1), v68)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L6
	} else {
		goto L18
	}
L17:
	;
	v90 = v41 + v51
	goto L15
L18:
	;
	if base.B2i32(v56 == v41) == int32(0) {
		v56 = v56 + int32(1)
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v102 = v90
	goto L23
L21:
	;
	goto L22
L22:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v143
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
	v151 = F_AllocSetContextCreateInternal(m, v146, int32(_a_F_gistrescan_4), int32(0), int32(_a_F_gistrescan_1), int32(_a_F_gistrescan_2))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L6
	} else {
		goto L27
	}
L23:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+16))
	v112 = int32(0)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v113+v114<<(uint(int32(4))%32)+v102*int32(100)-int32(12))))
	F_TupleDescInitEntry(m, v110, base.I32_extend16_s(v102), v112, v123, int32(-1), v112)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L6
	} else {
		goto L25
	}
L24:
	;
	goto L22
L25:
	;
	if v102 != v43 {
		v102 = v102 + int32(1)
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+uint32(_c_F_gistrescan[1]))) = v151
	goto L9
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v171
	*(*int32)(unsafe.Add(mBase, _c_F_gistrescan[0])) = v166
	v176 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+17)) = uint8(v176)
	if l1 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	if l3 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L30:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v180 <= int32(0) {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v183 = int32(0)
	if v17 == v183 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v232 = v225 * int32(48)
	if v232 != 0 {
		goto L42
	} else {
		goto L43
	}
L33:
	;
	v225 = v180
	v226 = int32(0)
	goto L32
L34:
	;
	goto L35
L35:
	;
	v189 = F_palloc(m, v180<<(uint(int32(2))%32))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L6
	} else {
		goto L36
	}
L36:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v191 <= int32(0) {
		v225 = v191
		v226 = v189
		goto L32
	} else {
		goto L37
	}
L37:
	;
	v199 = int32(0)
	goto L38
L38:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v209+v199*int32(48))+32))
	*(*int32)(unsafe.Add(mBase, uint32(v189+v199<<(uint(int32(2))%32)))) = v213
	v216 = v199 + int32(1)
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v216 < v217 {
		v199 = v216
		goto L38
	} else {
		goto L40
	}
L39:
	;
	v225 = v217
	v226 = v189
	goto L32
L40:
	;
	goto L39
L41:
	;
	v235 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+16)) = uint8(v235)
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if int32(0) < v237 {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	v233 = F__emscripten_memcpy_bulkmem(m, v230, l1, v232)
	mBase = m.M
	goto L44
L43:
	;
	goto L44
L44:
	;
	goto L41
L45:
	;
	v242 = v183
	goto L48
L46:
	;
	goto L47
L47:
	;
	if v17 == int32(0) {
		goto L29
	} else {
		goto L58
	}
L48:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v254 = v251 + v242*int32(48)
	v256 = v254 + int32(16)
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v258 = int32(*(*int16)(unsafe.Add(mBase, uint32(v254)+4)))
	v263 = v257 + v258*int32(28) - int32(8)
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v257)))
	v267 = v254 + int32(32)
	v268 = *(*int64)(unsafe.Add(mBase, uint32(v263)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v267))) = v268
	v270 = *(*int64)(unsafe.Add(mBase, uint32(v263)))
	*(*int64)(unsafe.Add(mBase, uint32(v256))) = v270
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v263)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v256)+24)) = v272
	v274 = *(*int64)(unsafe.Add(mBase, uint32(v263)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v256)+8)) = v274
	*(*int32)(unsafe.Add(mBase, uint32(v256)+20)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v267))) = int32(0)
	goto L50
L49:
	;
	goto L47
L50:
	;
	if v17 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v226+v242<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v254)+32)) = v282
	goto L53
L52:
	;
	goto L53
L53:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v254)))
	if v284&int32(193) == int32(1) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v289 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+16)) = uint8(v289)
	goto L56
L55:
	;
	goto L56
L56:
	;
	v292 = v242 + int32(1)
	v293 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v292 < v293 {
		v242 = v292
		goto L48
	} else {
		goto L57
	}
L57:
	;
	goto L49
L58:
	;
	F_pfree(m, v226)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L6
	} else {
		goto L59
	}
L59:
	;
	goto L29
L60:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L6
	} else {
		goto L92
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = int32(0)
	m.G0 = v14 + int32(16)
	return
L62:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v323 <= int32(0) {
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v326 = int32(0)
	if v17 == v326 {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v375 = v368 * int32(48)
	if v375 != 0 {
		goto L74
	} else {
		goto L75
	}
L65:
	;
	v363 = int32(0)
	v368 = v323
	goto L64
L66:
	;
	goto L67
L67:
	;
	v332 = F_palloc(m, v323<<(uint(int32(2))%32))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L6
	} else {
		goto L68
	}
L68:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v334 <= int32(0) {
		v363 = v332
		v368 = v334
		goto L64
	} else {
		goto L69
	}
L69:
	;
	v342 = int32(0)
	goto L70
L70:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v352+v342*int32(48))+32))
	*(*int32)(unsafe.Add(mBase, uint32(v332+v342<<(uint(int32(2))%32)))) = v356
	v359 = v342 + int32(1)
	v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v359 < v360 {
		v342 = v359
		goto L70
	} else {
		goto L72
	}
L71:
	;
	v363 = v332
	v368 = v360
	goto L64
L72:
	;
	goto L71
L73:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v381 = F_palloc(m, v378<<(uint(int32(2))%32))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L6
	} else {
		goto L77
	}
L74:
	;
	v376 = F__emscripten_memcpy_bulkmem(m, v373, l3, v375)
	mBase = m.M
	goto L76
L75:
	;
	goto L76
L76:
	;
	goto L73
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v381
	v384 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if int32(0) < v384 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v389 = v326
	goto L81
L79:
	;
	goto L80
L80:
	;
	if v17 == int32(0) {
		goto L61
	} else {
		goto L90
	}
L81:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v399 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v402 = v399 + v389*int32(48)
	v403 = int32(*(*int16)(unsafe.Add(mBase, uint32(v402)+4)))
	v406 = v398 + v403*int32(28)
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v406)+uint32(_c_F_gistrescan[2])))
	if v409 == int32(0) {
		goto L60
	} else {
		goto L83
	}
L82:
	;
	goto L80
L83:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v402)+20))
	v413 = F_get_func_rettype(m, v412)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L6
	} else {
		goto L84
	}
L84:
	;
	v416 = v389 << (uint(int32(2)) % 32)
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v416+v417))) = v413
	v421 = v402 + int32(16)
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v424)))
	v428 = v402 + int32(32)
	v429 = *(*int64)(unsafe.Add(mBase, uint32(v406)+uint32(_c_F_gistrescan[3])))
	*(*int64)(unsafe.Add(mBase, uint32(v428))) = v429
	v431 = *(*int64)(unsafe.Add(mBase, uint32(v406)+uint32(_c_F_gistrescan[4])))
	*(*int64)(unsafe.Add(mBase, uint32(v421))) = v431
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v406)+uint32(_c_F_gistrescan[5])))
	*(*int32)(unsafe.Add(mBase, uint32(v421)+24)) = v433
	v435 = *(*int64)(unsafe.Add(mBase, uint32(v406)+uint32(_c_F_gistrescan[6])))
	*(*int64)(unsafe.Add(mBase, uint32(v421)+8)) = v435
	*(*int32)(unsafe.Add(mBase, uint32(v421)+20)) = v425
	*(*int32)(unsafe.Add(mBase, uint32(v428))) = int32(0)
	goto L85
L85:
	;
	if v17 != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v363+v416)))
	*(*int32)(unsafe.Add(mBase, uint32(v402)+32)) = v441
	goto L88
L87:
	;
	goto L88
L88:
	;
	v444 = v389 + int32(1)
	v445 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v444 < v445 {
		v389 = v444
		goto L81
	} else {
		goto L89
	}
L89:
	;
	goto L82
L90:
	;
	F_pfree(m, v363)
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L6
	} else {
		goto L91
	}
L91:
	;
	goto L61
L92:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v482)+48))
	v484 = int32(*(*int16)(unsafe.Add(mBase, uint32(v402)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v484
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v483 + int32(4)
	F_errmsg_internal(m, int32(_a_F_gistrescan_5), v14)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L6
	} else {
		goto L93
	}
L93:
	;
	F_errfinish(m, int32(_a_F_gistrescan_6), int32(311), int32(_a_F_gistrescan_7))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L6
	} else {
		goto L94
	}
L94:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_gistvacuumscan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v26 int64
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int64
	_ = v72
	var v73 int32
	_ = v73
	var v74 int64
	_ = v74
	var v75 int32
	_ = v75
	var v76 int64
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int64
	_ = v200
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v265 int32
	_ = v265
	var v268 int64
	_ = v268
	var v270 int64
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v311 int64
	_ = v311
	var v312 int64
	_ = v312
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v431 int64
	_ = v431
	var v432 int32
	_ = v432
	var v433 int64
	_ = v433
	var v434 int32
	_ = v434
	var v435 int64
	_ = v435
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v445 float64
	_ = v445
	var v449 int32
	_ = v449
	var v475 int32
	_ = v475
	var v487 int32
	_ = v487
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v547 int32
	_ = v547
	var v574 int32
	_ = v574
	var v578 int32
	_ = v578
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v593 int32
	_ = v593
	var v597 int32
	_ = v597
	var v601 int32
	_ = v601
	var v606 int32
	_ = v606
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v645 float64
	_ = v645
	var v680 int32
	_ = v680
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v700 int64
	_ = v700
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
	var v707 int32
	_ = v707
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v724 int32
	_ = v724
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v753 int32
	_ = v753
	var v756 int32
	_ = v756
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v762 int32
	_ = v762
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v790 int32
	_ = v790
	var v796 int32
	_ = v796
	var v797 int64
	_ = v797
	var v799 int64
	_ = v799
	var v806 int32
	_ = v806
	var v809 int32
	_ = v809
	var v812 int64
	_ = v812
	var v815 int64
	_ = v815
	var v818 int64
	_ = v818
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v831 int32
	_ = v831
	var v835 int32
	_ = v835
	var v837 int32
	_ = v837
	var v858 int64
	_ = v858
	var v859 int64
	_ = v859
	var v865 int32
	_ = v865
	var v868 int64
	_ = v868
	var v869 int64
	_ = v869
	var v871 int64
	_ = v871
	var v875 int64
	_ = v875
	var v877 int64
	_ = v877
	var v881 int64
	_ = v881
	var v883 int64
	_ = v883
	var v887 int64
	_ = v887
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v891 int64
	_ = v891
	var v893 int32
	_ = v893
	var v897 int32
	_ = v897
	var v920 int64
	_ = v920
	var v921 int64
	_ = v921
	var v927 int32
	_ = v927
	var v930 int32
	_ = v930
	var v950 int64
	_ = v950
	var v951 int64
	_ = v951
	var v961 int64
	_ = v961
	var v963 int32
	_ = v963
	var v967 int32
	_ = v967
	var v1004 int32
	_ = v1004
	var v1033 int32
	_ = v1033
	var v1037 int32
	_ = v1037
	var v1040 int32
	_ = v1040
	var v1043 int32
	_ = v1043
	var v1045 int32
	_ = v1045
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1053 int32
	_ = v1053
	var v1058 int32
	_ = v1058
	var v1085 int32
	_ = v1085
	var v1088 int32
	_ = v1088
	var v1092 int64
	_ = v1092
	var v1125 int32
	_ = v1125
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1133 int32
	_ = v1133
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1142 int32
	_ = v1142
	var v1148 int32
	_ = v1148
	var v1150 int32
	_ = v1150
	var v1156 int32
	_ = v1156
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1160 int32
	_ = v1160
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1175 int32
	_ = v1175
	var v1177 int32
	_ = v1177
	var v1180 int32
	_ = v1180
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1194 int32
	_ = v1194
	var v1198 int32
	_ = v1198
	var v1227 int32
	_ = v1227
	var v1230 int32
	_ = v1230
	var v1231 int32
	_ = v1231
	var v1234 int32
	_ = v1234
	var v1235 int32
	_ = v1235
	var v1236 int64
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1238 int32
	_ = v1238
	var v1241 int64
	_ = v1241
	var v1244 int32
	_ = v1244
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1278 int32
	_ = v1278
	var v1279 int32
	_ = v1279
	var v1283 int64
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1288 int32
	_ = v1288
	var v1295 int64
	_ = v1295
	var v1297 int32
	_ = v1297
	var v1300 int32
	_ = v1300
	var v1302 int32
	_ = v1302
	var v1312 int32
	_ = v1312
	var v1327 int32
	_ = v1327
	var v1335 int32
	_ = v1335
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1378 int64
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1380 int32
	_ = v1380
	var v1383 int32
	_ = v1383
	var v1385 int32
	_ = v1385
	var v1388 int32
	_ = v1388
	var v1393 int32
	_ = v1393
	var v1403 int32
	_ = v1403
	var v1426 int32
	_ = v1426
	var v1430 int32
	_ = v1430
	var v1437 int32
	_ = v1437
	var v1438 int32
	_ = v1438
	var v1464 int32
	_ = v1464
	var v1465 int32
	_ = v1465
	var v1469 int64
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1471 int32
	_ = v1471
	var v1474 int32
	_ = v1474
	var v1483 int32
	_ = v1483
	var v1484 int64
	_ = v1484
	var v1486 int64
	_ = v1486
	var v1493 int32
	_ = v1493
	var v1496 int32
	_ = v1496
	var v1502 int32
	_ = v1502
	var v1505 int32
	_ = v1505
	var v1508 int64
	_ = v1508
	var v1511 int64
	_ = v1511
	var v1521 int32
	_ = v1521
	var v1541 int64
	_ = v1541
	var v1542 int64
	_ = v1542
	var v1548 int64
	_ = v1548
	var v1552 int32
	_ = v1552
	var v1562 int32
	_ = v1562
	var v1587 int32
	_ = v1587
	var v1613 int32
	_ = v1613
	var v1649 int32
	_ = v1649
	var v1652 int32
	_ = v1652
	var v1664 int32
	_ = v1664
	var v1667 int32
	_ = v1667
	var v1674 int32
	_ = v1674
	var v1679 int32
	_ = v1679
	var v1687 int32
	_ = v1687
	var v1692 int32
	_ = v1692
	var v1715 int32
	_ = v1715
	var v1724 int32
	_ = v1724
	var v1730 int32
	_ = v1730
	var v1732 int32
	_ = v1732
	var v1733 int32
	_ = v1733
	var v1734 int32
	_ = v1734
	var v1737 int32
	_ = v1737
	var v1739 int32
	_ = v1739
	var v1742 int32
	_ = v1742
	var v1751 int32
	_ = v1751
	var v1753 int32
	_ = v1753
	var v1755 int32
	_ = v1755
	var v1759 int32
	_ = v1759
	var v1760 int32
	_ = v1760
	var v1764 int32
	_ = v1764
	var v1770 int32
	_ = v1770
	var v1772 int32
	_ = v1772
	var v1778 int32
	_ = v1778
	var v1779 int32
	_ = v1779
	var v1781 int32
	_ = v1781
	var v1786 int32
	_ = v1786
	var v1794 int32
	_ = v1794
	var v1795 int32
	_ = v1795
	var v1798 int32
	_ = v1798
	var v1800 int32
	_ = v1800
	var v1803 int32
	_ = v1803
	var v1811 int32
	_ = v1811
	var v1812 int32
	_ = v1812
	var v1813 int32
	_ = v1813
	var v1816 int32
	_ = v1816
	var v1823 int32
	_ = v1823
	var v1827 int32
	_ = v1827
	var v1833 int32
	_ = v1833
	var v1835 int32
	_ = v1835
	var v1841 int32
	_ = v1841
	var v1842 int32
	_ = v1842
	var v1845 int32
	_ = v1845
	var v1846 int32
	_ = v1846
	var v1849 int32
	_ = v1849
	var v1852 int64
	_ = v1852
	var v1853 int32
	_ = v1853
	var v1854 int32
	_ = v1854
	var v1856 int32
	_ = v1856
	var v1861 int32
	_ = v1861
	var v1862 int32
	_ = v1862
	var v1863 int32
	_ = v1863
	var v1864 int32
	_ = v1864
	var v1866 int32
	_ = v1866
	var v1869 int32
	_ = v1869
	var v1871 int32
	_ = v1871
	var v1872 int32
	_ = v1872
	var v1875 int32
	_ = v1875
	var v1880 int32
	_ = v1880
	var v1882 int32
	_ = v1882
	var v1883 int32
	_ = v1883
	var v1884 int32
	_ = v1884
	var v1885 int32
	_ = v1885
	var v1889 int32
	_ = v1889
	var v1892 int32
	_ = v1892
	var v1893 int32
	_ = v1893
	var v1894 int32
	_ = v1894
	var v1896 int32
	_ = v1896
	var v1901 int32
	_ = v1901
	var v1904 int32
	_ = v1904
	var v1908 int32
	_ = v1908
	var v1912 int32
	_ = v1912
	var v1915 int64
	_ = v1915
	var v1916 int32
	_ = v1916
	var v1920 int64
	_ = v1920
	var v1921 int32
	_ = v1921
	var v1924 int64
	_ = v1924
	var v1925 int64
	_ = v1925
	var v1930 int64
	_ = v1930
	var v1932 int32
	_ = v1932
	var v1934 int32
	_ = v1934
	var v1935 int32
	_ = v1935
	var v1942 int32
	_ = v1942
	var v1947 int32
	_ = v1947
	var v1949 int32
	_ = v1949
	var v1951 int32
	_ = v1951
	var v1961 int32
	_ = v1961
	var v1984 int32
	_ = v1984
	var v2004 int32
	_ = v2004
	var v2030 int32
	_ = v2030
	var v2039 int32
	_ = v2039
	var v2047 int32
	_ = v2047
	v5 = int32(0)
	v26 = int64(0)
	v31 = m.G0
	v33 = v31 - int32(_a_F_gistvacuumscan_0)
	m.G0 = v33
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+28)) = v26
	*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = v26
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)) = uint8(v5)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v5
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_gistvacuumscan[0]))
	v47 = int32(_a_F_gistvacuumscan_1)
	v50 = F_GenerationContextCreate(m, v45, int32(_a_F_gistvacuumscan_2), v47, v47, v47)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v52 = int32(_a_F_gistvacuumscan_3)
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_gistvacuumscan[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_gistvacuumscan[0])) = v50
	v56 = F_intset_create(m)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v58 = F_intset_create(m)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_gistvacuumscan[0])) = v53
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v35)+48))
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+118)))
	if v63 != int32(112) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+24)))
	if v77 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L6:
	;
	v74 = F_gistGetFakeLSN(m, v35)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L14
	}
L7:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_gistvacuumscan[1]))
	if v67 <= int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v35)+32))
	if v70 != 0 {
		goto L6
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v72 = F_GetInsertRecPtr(m)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v35)+40))
	if v71 != 0 {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v76 = v72
	goto L5
L14:
	;
	v76 = v74
	goto L5
L15:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v35)+32))
	v83 = base.B2i32(v80 == int32(0))
	goto L17
L16:
	;
	v83 = v5
	goto L17
L17:
	;
	v84 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+8)) = v84
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v92 = F_read_stream_begin_relation(m, int32(13), v87, v35, int32(120), v33+int32(8), v84)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	goto L19
L19:
	;
	if v83 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	F_read_stream_end(m, v92)
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L1
	} else {
		goto L136
	}
L21:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	if base.Ui32(v139) < base.Ui32(v138) {
		goto L29
	} else {
		goto L30
	}
L22:
	;
	v127 = F_RelationGetNumberOfBlocksInFork(m, v35, int32(0))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	F_LockRelationForExtension(m, v35, int32(7))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L26
	}
L25:
	;
	v138 = v127
	goto L21
L26:
	;
	v133 = F_RelationGetNumberOfBlocksInFork(m, v35, int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_UnlockRelationForExtension(m, v35, int32(7))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v138 = v133
	goto L21
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+12)) = v138
	goto L32
L30:
	;
	goto L31
L31:
	;
	goto L20
L32:
	;
	F_vacuum_delay_point(m, int32(0))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L35
	}
L33:
	;
	F_read_stream_reset(m, v92)
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L1
	} else {
		goto L135
	}
L34:
	;
	goto L33
L35:
	;
	v176 = F_read_stream_next_buffer(m, v92, int32(0))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	if v176 == int32(0) {
		goto L34
	} else {
		goto L37
	}
L37:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v176 < int32(0) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v200 = base.I64_extend_i32_u(v199)
	v213 = v176
	v219 = v199
	goto L42
L39:
	;
	v184 = *(*int32)(unsafe.Add(mBase, _c_F_gistvacuumscan[2]))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v184+(v176^int32(-1))<<(uint(int32(6))%32))+16))
	v199 = v190
	goto L38
L40:
	;
	goto L41
L41:
	;
	v192 = *(*int32)(unsafe.Add(mBase, _c_F_gistvacuumscan[3]))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v192+v176<<(uint(int32(6))%32)+int32(-64))+16))
	v199 = v198
	goto L38
L42:
	;
	F_LockBuffer(m, v213, int32(2))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	if v213 < int32(0) {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	if v273 != 0 {
		goto L56
	} else {
		goto L57
	}
L46:
	;
	v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v252)+14)))
	if v253 == int32(0) {
		v273 = int32(1)
		goto L45
	} else {
		goto L50
	}
L47:
	;
	v238 = *(*int32)(unsafe.Add(mBase, _c_F_gistvacuumscan[4]))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v238+(v213^int32(-1))<<(uint(int32(2))%32))))
	v252 = v244
	goto L46
L48:
	;
	goto L49
L49:
	;
	v246 = *(*int32)(unsafe.Add(mBase, _c_F_gistvacuumscan[5]))
	v252 = v246 + v213<<(uint(int32(13))%32) + int32(-8192)
	goto L46
L50:
	;
	v256 = int32(0)
	v257 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v252)+16)))
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v252+v257)+12)))
	if v259&int32(2) == v256 {
		v273 = v256
		goto L45
	} else {
		goto L51
	}
L51:
	;
	v265 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v252)+12)))
	if base.Ui32(int32(32)) <= base.Ui32(v265) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v268 = *(*int64)(unsafe.Add(mBase, uint32(v252)+24))
	v270 = v268
	goto L54
L53:
	;
	v270 = int64(3)
	goto L54
L54:
	;
	v271 = F_GlobalVisCheckRemovableFullXid(m, int32(0), v270)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v273 = v271
	goto L45
L56:
	;
	F_RecordFreeIndexPage(m, v180, v219)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v286 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v252)+16)))
	v287 = v252 + v286
	v288 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v287)+12)))
	if v288&int32(2) != 0 {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v277 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v276 + v277
	v280 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v280 + v277
	F_UnlockReleaseBuffer(m, v213)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	goto L32
L61:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v291 + int32(1)
	F_UnlockReleaseBuffer(m, v213)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	if v288&int32(1) != 0 {
		goto L67
	} else {
		goto L68
	}
L64:
	;
	goto L32
L65:
	;
	F_UnlockReleaseBuffer(m, v213)
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L1
	} else {
		goto L131
	}
L66:
	;
	v645 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = base.F64_add(v645, base.F64_convert_i32_u(v487))
	goto L65
L67:
	;
	v299 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v252)+12)))
	if v288&int32(8) == int32(0) {
		goto L71
	} else {
		goto L72
	}
L68:
	;
	goto L69
L69:
	;
	v521 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v252)+12)))
	if base.Ui32(v521) < base.Ui32(int32(25)) {
		goto L112
	} else {
		goto L113
	}
L70:
	;
	if base.Ui32(v299) < base.Ui32(int32(25)) {
		goto L81
	} else {
		goto L82
	}
L71:
	;
	v311 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v287)+4)))
	v312 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v287))))
	if base.Ui64(v311|v312<<(uint(int64(32))%64)) <= base.Ui64(v76) {
		v326 = int32(-1)
		goto L70
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v317 = int32(-1)
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v287)+8))
	if base.Ui32(v199) <= base.Ui32(v319) {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	goto L73
L75:
	;
	v321 = v317
	goto L77
L76:
	;
	v321 = v319
	goto L77
L77:
	;
	if v319 == int32(-1) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v324 = v317
	goto L80
L79:
	;
	v324 = v321
	goto L80
L80:
	;
	v326 = v324
	goto L70
L81:
	;
	v328 = int32(0)
	goto L83
L82:
	;
	v328 = int32(base.Ui32(v299+int32(_a_F_gistvacuumscan_4)) >> (uint(int32(2)) % 32))
	goto L83
L83:
	;
	if l2 == int32(0) {
		v475 = v328
		goto L85
	} else {
		goto L86
	}
L84:
	;
	if v199 != v219 {
		goto L65
	} else {
		goto L110
	}
L85:
	;
	v487 = v475 & int32(_a_F_gistvacuumscan_5)
	if v487 != 0 {
		goto L66
	} else {
		goto L109
	}
L86:
	;
	v332 = v328 & int32(_a_F_gistvacuumscan_5)
	if v332 == int32(0) {
		goto L84
	} else {
		goto L87
	}
L87:
	;
	v343 = int32(1)
	v347 = int32(0)
	goto L88
L88:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v343&int32(_a_F_gistvacuumscan_5)<<(uint(int32(2))%32)+(v252+int32(24))-int32(4))))
	v380 = m.T0[l2].(func(*base.Module, int32, int32) int32)(m, v252+v376&int32(_a_F_gistvacuumscan_6), l3)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L1
	} else {
		goto L90
	}
L89:
	;
	if v390 <= int32(0) {
		v475 = v328
		goto L85
	} else {
		goto L95
	}
L90:
	;
	if v380 != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v384 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v33+int32(16)+v347<<(uint(v384)%32)))) = uint16(v343)
	v390 = v347 + v384
	goto L93
L92:
	;
	v390 = v347
	goto L93
L93:
	;
	v392 = v343 + int32(1)
	if base.Ui32(v392&int32(_a_F_gistvacuumscan_5)) <= base.Ui32(v332) {
		v343 = v392
		v347 = v390
		goto L88
	} else {
		goto L94
	}
L94:
	;
	goto L89
L95:
	;
	v398 = int32(_a_F_gistvacuumscan_7)
	v400 = *(*int32)(unsafe.Add(mBase, _c_F_gistvacuumscan[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_gistvacuumscan[6])) = v400 + int32(1)
	F_MarkBufferDirty(m, v213)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	F_PageIndexMultiDelete(m, v252, v33+int32(16), v390)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	v410 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v252)+16)))
	v411 = v252 + v410
	v412 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v411)+12)))
	v414 = v412 | int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v411)+12)) = uint16(v414)
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v180)+48))
	v417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v416)+118)))
	if v417 != int32(112) {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v252))) = base.I64_rotr(v435, int64(32))
	v439 = int32(_a_F_gistvacuumscan_7)
	v441 = *(*int32)(unsafe.Add(mBase, _c_F_gistvacuumscan[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_gistvacuumscan[6])) = v441 - int32(1)
	v445 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*float64)(unsafe.Add(mBase, uint32(l1)+16)) = base.F64_add(v445, base.F64_convert_i32_u(v390))
	v449 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v252)+12)))
	if base.Ui32(v449) < base.Ui32(int32(25)) {
		goto L84
	} else {
		goto L108
	}
L99:
	;
	v433 = F_gistGetFakeLSN(m, v180)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L1
	} else {
		goto L107
	}
L100:
	;
	v421 = *(*int32)(unsafe.Add(mBase, _c_F_gistvacuumscan[1]))
	if v421 <= int32(0) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v180)+32))
	if v424 != 0 {
		goto L99
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	v428 = int32(0)
	v431 = F_gistXLogUpdate(m, v213, v33+int32(16), v390, v428, v428, v428)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L1
	} else {
		goto L106
	}
L104:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v180)+40))
	if v425 != 0 {
		goto L99
	} else {
		goto L105
	}
L105:
	;
	goto L103
L106:
	;
	v435 = v431
	goto L98
L107:
	;
	v435 = v433
	goto L98
L108:
	;
	v475 = int32(base.Ui32(v449+int32(_a_F_gistvacuumscan_4)) >> (uint(int32(2)) % 32))
	goto L85
L109:
	;
	goto L84
L110:
	;
	F_intset_add_member(m, v58, v200)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	goto L65
L112:
	;
	if v199 == v219 {
		goto L126
	} else {
		goto L127
	}
L113:
	;
	v525 = v521 + int32(_a_F_gistvacuumscan_4)
	if v525&int32(_a_F_gistvacuumscan_8) == int32(0) {
		goto L112
	} else {
		goto L114
	}
L114:
	;
	v547 = int32(1)
	goto L115
L115:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v547<<(uint(int32(2))%32)+(v252+int32(24))-int32(4))))
	v578 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v252+int32(4)+v574&int32(_a_F_gistvacuumscan_6)))))
	if v578 != int32(_a_F_gistvacuumscan_9) {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	goto L112
L117:
	;
	if v547 != int32(base.Ui32(v525)>>(uint(int32(2))%32))&int32(_a_F_gistvacuumscan_5) {
		v547 = v547 + int32(1)
		goto L115
	} else {
		goto L125
	}
L118:
	;
	v583 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	if v583 == int32(0) {
		goto L117
	} else {
		goto L120
	}
L120:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v180)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v587 + int32(4)
	F_errmsg(m, int32(_a_F_gistvacuumscan_10), v33)
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	F_errdetail(m, int32(_a_F_gistvacuumscan_11), int32(0))
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	F_errhint(m, int32(_a_F_gistvacuumscan_12), int32(0))
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	F_errfinish(m, int32(_a_F_gistvacuumscan_13), int32(466), int32(_a_F_gistvacuumscan_14))
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	goto L117
L125:
	;
	goto L116
L126:
	;
	F_intset_add_member(m, v56, v200)
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L1
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	F_UnlockReleaseBuffer(m, v213)
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L1
	} else {
		goto L130
	}
L129:
	;
	goto L128
L130:
	;
	goto L32
L131:
	;
	if v326 == int32(-1) {
		goto L32
	} else {
		goto L132
	}
L132:
	;
	F_vacuum_delay_point(m, int32(0))
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	v686 = int32(0)
	v688 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v689 = F_ReadBufferExtended(m, v180, v686, v326, v686, v688)
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	v213 = v689
	v219 = v326
	goto L42
L135:
	;
	goto L19
L136:
	;
	v695 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v695 != 0 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	F_FreeSpaceMapVacuum(m, v35)
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L1
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v138
	v699 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v700 = *(*int64)(unsafe.Add(mBase, uint32(v58)+16))
	v701 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+3948)) = uint8(v701)
	v703 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v56)+3968)) = v703
	*(*int64)(unsafe.Add(mBase, uint32(v56)+3956)) = int64(0)
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v56)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+3964)) = v707
	*(*int32)(unsafe.Add(mBase, uint32(v56)+3952)) = v56 + int32(3976)
	v712 = base.I32_wrap_i64(v700)
	if v712 == v703 {
		v2030 = v33
		v2039 = v50
		goto L141
	} else {
		goto L142
	}
L140:
	;
	goto L139
L141:
	;
	F_MemoryContextDelete(m, v2039)
	mBase = m.M
	v2047 = m.ExcPending
	if v2047 != 0 {
		goto L1
	} else {
		goto L320
	}
L142:
	;
	v715 = l0
	v716 = l1
	v724 = v56
	v729 = v33
	v731 = v699
	v733 = v712
	v736 = v58
	v738 = v50
	goto L143
L143:
	;
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v724)+3960))
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v724)+3956))
	if v747 < v748 {
		v1058 = v747
		goto L146
	} else {
		goto L147
	}
L144:
	;
	v2030 = v729
	v2039 = v738
	goto L141
L145:
	;
	if v1125 == int32(0) {
		v2030 = v729
		v2039 = v738
		goto L141
	} else {
		goto L177
	}
L146:
	;
	v1085 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v724)+3960)) = v1058 + v1085
	v1088 = *(*int32)(unsafe.Add(mBase, uint32(v724)+3952))
	v1092 = *(*int64)(unsafe.Add(mBase, uint32(v1088+v1058<<(uint(int32(3))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v729)+uint32(_c_F_gistvacuumscan[7]))) = v1092
	v1125 = v1085
	goto L145
L147:
	;
	v753 = v724 + int32(3984)
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v724)+3964))
	v759 = v756
	v760 = v747
	v762 = v748
	goto L148
L148:
	;
	if v759 != 0 {
		goto L151
	} else {
		goto L152
	}
L149:
	;
	v1058 = v1051
	goto L146
L150:
	;
	if v1053 <= v1051 {
		v759 = v1050
		v760 = v1051
		v762 = v1053
		goto L148
	} else {
		goto L176
	}
L151:
	;
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v724)+3968))
	v788 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v759)+2)))
	if v787 < v788 {
		goto L154
	} else {
		goto L155
	}
L152:
	;
	goto L153
L153:
	;
	v1037 = *(*int32)(unsafe.Add(mBase, uint32(v724)+3952))
	if v724+int32(3976) == v1037 {
		goto L173
	} else {
		goto L174
	}
L154:
	;
	v790 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v724)+3968)) = v787 + v790
	v796 = v759 + v787<<(uint(int32(4))%32)
	v797 = *(*int64)(unsafe.Add(mBase, uint32(v796)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v724)+3976)) = v797
	v799 = *(*int64)(unsafe.Add(mBase, uint32(v796)+16))
	if v799 == int64(1152921504606846975) {
		v1004 = v790
		goto L157
	} else {
		goto L158
	}
L155:
	;
	goto L156
L156:
	;
	v1033 = *(*int32)(unsafe.Add(mBase, uint32(v759)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v724)+3968)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v724)+3964)) = v1033
	v1050 = v1033
	v1051 = v760
	v1053 = v762
	goto L150
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v724)+3956)) = v1004
	v1058 = int32(0)
	goto L146
L158:
	;
	v806 = base.I32_wrap_i64(int64(base.Ui64(v799)>>(uint(int64(60))%64))) << (uint(int32(1)) % 32)
	v809 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v806)+uint32(_c_F_gistvacuumscan[8]))))
	if v809 == int32(0) {
		v1004 = v790
		goto L157
	} else {
		goto L159
	}
L159:
	;
	v812 = int64(-1)
	v815 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v806)+uint32(_c_F_gistvacuumscan[9]))))
	v818 = v812<<(uint(v815)%64) ^ v812
	v819 = int32(3)
	v820 = v809 & v819
	if base.Ui32((v809-int32(1))&int32(255)) < base.Ui32(v819) {
		goto L161
	} else {
		goto L162
	}
L160:
	;
	if v820 != 0 {
		goto L167
	} else {
		goto L168
	}
L161:
	;
	v897 = int32(0)
	v920 = v799
	v921 = v797
	goto L160
L162:
	;
	goto L163
L163:
	;
	v831 = int32(0)
	v835 = v831
	v837 = v831
	v858 = v799
	v859 = v797
	goto L164
L164:
	;
	v865 = v753 + v835<<(uint(int32(3))%32)
	v868 = int64(1)
	v869 = v859 + v858&v818 + v868
	*(*int64)(unsafe.Add(mBase, uint32(v865))) = v869
	v871 = int64(base.Ui64(v858) >> (uint(v815) % 64))
	v875 = v869 + v871&v818 + v868
	*(*int64)(unsafe.Add(mBase, uint32(v865)+8)) = v875
	v877 = int64(base.Ui64(v871) >> (uint(v815) % 64))
	v881 = v875 + v877&v818 + v868
	*(*int64)(unsafe.Add(mBase, uint32(v865)+16)) = v881
	v883 = int64(base.Ui64(v877) >> (uint(v815) % 64))
	v887 = v881 + v883&v818 + v868
	*(*int64)(unsafe.Add(mBase, uint32(v865)+24)) = v887
	v889 = int32(4)
	v890 = v835 + v889
	v891 = int64(base.Ui64(v883) >> (uint(v815) % 64))
	v893 = v837 + v889
	if v893 != v809&int32(252) {
		v835 = v890
		v837 = v893
		v858 = v891
		v859 = v887
		goto L164
	} else {
		goto L166
	}
L165:
	;
	v897 = v890
	v920 = v891
	v921 = v887
	goto L160
L166:
	;
	goto L165
L167:
	;
	v927 = v897
	v930 = int32(0)
	v950 = v920
	v951 = v921
	goto L170
L168:
	;
	goto L169
L169:
	;
	v1004 = v809 + int32(1)
	goto L157
L170:
	;
	v961 = v951 + v950&v818 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v753+v927<<(uint(int32(3))%32)))) = v961
	v963 = int32(1)
	v967 = v930 + v963
	if v967 != v820 {
		v927 = v927 + v963
		v930 = v967
		v950 = int64(base.Ui64(v950) >> (uint(v815) % 64))
		v951 = v961
		goto L170
	} else {
		goto L172
	}
L171:
	;
	goto L169
L172:
	;
	goto L171
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v724)+3952)) = v724 + int32(88)
	v1040 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v724)+3960)) = v1040
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(v724)+3944))
	*(*int32)(unsafe.Add(mBase, uint32(v724)+3956)) = v1043
	v1050 = v759
	v1051 = v1040
	v1053 = v1043
	goto L150
L174:
	;
	goto L175
L175:
	;
	v1045 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v724)+3948)) = uint8(v1045)
	*(*int64)(unsafe.Add(mBase, uint32(v729)+uint32(_c_F_gistvacuumscan[7]))) = int64(0)
	v1125 = v1045
	goto L145
L176:
	;
	goto L149
L177:
	;
	v1128 = int32(0)
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(v729)+uint32(_c_F_gistvacuumscan[7])))
	v1131 = *(*int32)(unsafe.Add(mBase, uint32(v715)+24))
	v1132 = F_ReadBufferExtended(m, v731, v1128, v1129, v1128, v1131)
	mBase = m.M
	v1133 = m.ExcPending
	if v1133 != 0 {
		goto L1
	} else {
		goto L178
	}
L178:
	;
	F_LockBuffer(m, v1132, int32(1))
	mBase = m.M
	v1136 = m.ExcPending
	if v1136 != 0 {
		goto L1
	} else {
		goto L179
	}
L179:
	;
	v1137 = int32(0)
	v1138 = base.B2i32(v1137 <= v1132)
	if v1138 == v1137 {
		goto L183
	} else {
		goto L184
	}
L180:
	;
	if v2004 != 0 {
		v733 = v2004
		goto L143
	} else {
		goto L319
	}
L181:
	;
	v1167 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1156)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v1167) {
		goto L192
	} else {
		goto L193
	}
L182:
	;
	v1157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1156)+14)))
	if v1157 != 0 {
		goto L186
	} else {
		goto L187
	}
L183:
	;
	v1142 = *(*int32)(unsafe.Add(mBase, _c_F_gistvacuumscan[4]))
	v1148 = *(*int32)(unsafe.Add(mBase, uint32(v1142+(v1132^int32(-1))<<(uint(int32(2))%32))))
	v1156 = v1148
	goto L182
L184:
	;
	goto L185
L185:
	;
	v1150 = *(*int32)(unsafe.Add(mBase, _c_F_gistvacuumscan[5]))
	v1156 = v1150 + v1132<<(uint(int32(13))%32) + int32(-8192)
	goto L182
L186:
	;
	v1158 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1156)+16)))
	v1160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1156+v1158)+12)))
	if v1160&int32(3) == int32(0) {
		goto L181
	} else {
		goto L189
	}
L187:
	;
	goto L188
L188:
	;
	F_UnlockReleaseBuffer(m, v1132)
	mBase = m.M
	v1166 = m.ExcPending
	if v1166 != 0 {
		goto L1
	} else {
		goto L190
	}
L189:
	;
	goto L188
L190:
	;
	v2004 = v733
	goto L180
L191:
	;
	F_ReleaseBuffer(m, v1132)
	mBase = m.M
	v1984 = m.ExcPending
	if v1984 != 0 {
		goto L1
	} else {
		goto L318
	}
L192:
	;
	v1175 = int32(base.Ui32(v1167+int32(_a_F_gistvacuumscan_4)) >> (uint(int32(2)) % 32))
	goto L194
L193:
	;
	v1175 = int32(0)
	goto L194
L194:
	;
	v1177 = v1175 & int32(_a_F_gistvacuumscan_5)
	if base.Ui32(v1177) <= base.Ui32(int32(1)) {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v1180 = int32(0)
	F_LockBuffer(m, v1132, v1180)
	mBase = m.M
	v1183 = m.ExcPending
	if v1183 != 0 {
		goto L1
	} else {
		goto L198
	}
L196:
	;
	goto L197
L197:
	;
	v1184 = int32(1)
	v1194 = v1184
	v1198 = int32(0)
	goto L199
L198:
	;
	v1961 = v1180
	goto L191
L199:
	;
	v1227 = *(*int32)(unsafe.Add(mBase, uint32(v1194&int32(_a_F_gistvacuumscan_5)<<(uint(int32(2))%32)+(v1156+int32(24))-int32(4))))
	v1230 = v1156 + v1227&int32(_a_F_gistvacuumscan_6)
	v1231 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1230))))
	v1234 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1230)+2)))
	v1235 = v1231<<(uint(int32(16))%32) | v1234
	v1236 = base.I64_extend_i32_u(v1235)
	v1237 = int32(0)
	v1238 = *(*int32)(unsafe.Add(mBase, uint32(v736)+3944))
	if v1238 <= v1237 {
		goto L202
	} else {
		goto L203
	}
L200:
	;
	F_LockBuffer(m, v1132, int32(0))
	mBase = m.M
	v1674 = m.ExcPending
	if v1674 != 0 {
		goto L1
	} else {
		goto L259
	}
L201:
	;
	if v1649 != 0 {
		goto L255
	} else {
		goto L256
	}
L202:
	;
	v1297 = *(*int32)(unsafe.Add(mBase, uint32(v736)+36))
	if v1297 == int32(0) {
		v1613 = v1237
		goto L215
	} else {
		goto L216
	}
L203:
	;
	v1241 = *(*int64)(unsafe.Add(mBase, uint32(v736)+88))
	if base.Ui64(v1236) < base.Ui64(v1241) {
		goto L202
	} else {
		goto L204
	}
L204:
	;
	v1244 = v736 + int32(88)
	v1251 = v1238
	v1252 = int32(0)
	goto L205
L205:
	;
	v1278 = base.I32_div_s(v1251-v1252, int32(2))
	v1279 = v1278 + v1252
	v1283 = *(*int64)(unsafe.Add(mBase, uint32(v1244+v1279<<(uint(int32(3))%32))))
	v1284 = base.B2i32(base.Ui64(v1283) < base.Ui64(v1236))
	if base.Ui64(v1283) < base.Ui64(v1236) {
		goto L207
	} else {
		goto L208
	}
L206:
	;
	if v1238 <= v1288 {
		v1649 = int32(0)
		goto L201
	} else {
		goto L214
	}
L207:
	;
	v1285 = v1251
	goto L209
L208:
	;
	v1285 = v1279
	goto L209
L209:
	;
	if base.Ui64(v1283) < base.Ui64(v1236) {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	v1288 = v1279 + int32(1)
	goto L212
L211:
	;
	v1288 = v1252
	goto L212
L212:
	;
	if v1288 < v1285 {
		v1251 = v1285
		v1252 = v1288
		goto L205
	} else {
		goto L213
	}
L213:
	;
	goto L206
L214:
	;
	v1295 = *(*int64)(unsafe.Add(mBase, uint32(v1244+v1288<<(uint(int32(3))%32))))
	v1649 = base.B2i32(v1295 == v1236)
	goto L201
L215:
	;
	v1649 = v1613
	goto L201
L216:
	;
	v1300 = *(*int32)(unsafe.Add(mBase, uint32(v736)+32))
	v1302 = v1300 - int32(1)
	if int32(0) < v1302 {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	v1312 = v1297
	v1327 = v1302
	goto L220
L218:
	;
	v1403 = v1297
	goto L219
L219:
	;
	v1426 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1403)+2)))
	if v1426 == int32(0) {
		v1613 = v1237
		goto L215
	} else {
		goto L234
	}
L220:
	;
	v1335 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1312)+2)))
	if v1335 == int32(0) {
		v1613 = v1237
		goto L215
	} else {
		goto L222
	}
L221:
	;
	v1403 = v1393
	goto L219
L222:
	;
	v1346 = v1335
	v1347 = int32(0)
	goto L223
L223:
	;
	v1373 = base.I32_div_s(v1346-v1347, int32(2))
	v1374 = v1373 + v1347
	v1378 = *(*int64)(unsafe.Add(mBase, uint32(v1312+int32(8)+v1374<<(uint(int32(3))%32))))
	v1379 = base.B2i32(base.Ui64(v1236) < base.Ui64(v1378))
	if base.Ui64(v1236) < base.Ui64(v1378) {
		goto L225
	} else {
		goto L226
	}
L224:
	;
	v1385 = int32(0)
	if v1383 == v1385 {
		v1649 = v1385
		goto L201
	} else {
		goto L232
	}
L225:
	;
	v1380 = v1374
	goto L227
L226:
	;
	v1380 = v1346
	goto L227
L227:
	;
	if base.Ui64(v1236) < base.Ui64(v1378) {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	v1383 = v1347
	goto L230
L229:
	;
	v1383 = v1374 + int32(1)
	goto L230
L230:
	;
	if v1383 < v1380 {
		v1346 = v1380
		v1347 = v1383
		goto L223
	} else {
		goto L231
	}
L231:
	;
	goto L224
L232:
	;
	v1388 = int32(1)
	v1393 = *(*int32)(unsafe.Add(mBase, uint32(v1383<<(uint(int32(2))%32)+v1312)+516))
	if v1388 < v1327 {
		v1312 = v1393
		v1327 = v1327 - v1388
		goto L220
	} else {
		goto L233
	}
L233:
	;
	goto L221
L234:
	;
	v1430 = v1403 + int32(8)
	v1437 = v1426
	v1438 = int32(0)
	goto L235
L235:
	;
	v1464 = base.I32_div_s(v1437-v1438, int32(2))
	v1465 = v1464 + v1438
	v1469 = *(*int64)(unsafe.Add(mBase, uint32(v1430+v1465<<(uint(int32(4))%32))))
	v1470 = base.B2i32(base.Ui64(v1236) < base.Ui64(v1469))
	if base.Ui64(v1236) < base.Ui64(v1469) {
		goto L237
	} else {
		goto L238
	}
L236:
	;
	if v1474 == int32(0) {
		v1613 = v1237
		goto L215
	} else {
		goto L244
	}
L237:
	;
	v1471 = v1465
	goto L239
L238:
	;
	v1471 = v1437
	goto L239
L239:
	;
	if base.Ui64(v1236) < base.Ui64(v1469) {
		goto L240
	} else {
		goto L241
	}
L240:
	;
	v1474 = v1438
	goto L242
L241:
	;
	v1474 = v1465 + int32(1)
	goto L242
L242:
	;
	if v1474 < v1471 {
		v1437 = v1471
		v1438 = v1474
		goto L235
	} else {
		goto L243
	}
L243:
	;
	goto L236
L244:
	;
	v1483 = v1474<<(uint(int32(4))%32) + v1430 - int32(16)
	v1484 = *(*int64)(unsafe.Add(mBase, uint32(v1483)))
	if v1484 == v1236 {
		v1649 = int32(1)
		goto L201
	} else {
		goto L245
	}
L245:
	;
	v1486 = *(*int64)(unsafe.Add(mBase, uint32(v1483)+8))
	if v1486 == int64(1152921504606846975) {
		v1613 = v1237
		goto L215
	} else {
		goto L246
	}
L246:
	;
	v1493 = base.I32_wrap_i64(int64(base.Ui64(v1486)>>(uint(int64(60))%64))) << (uint(int32(1)) % 32)
	v1496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1493)+uint32(_c_F_gistvacuumscan[8]))))
	v1502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1493)+uint32(_c_F_gistvacuumscan[9]))))
	if v1502 == int32(0) {
		v1649 = base.B2i32(base.Ui64(v1236-v1484) <= base.Ui64(base.I64_extend_i32_u(v1496)))
		goto L201
	} else {
		goto L247
	}
L247:
	;
	v1505 = int32(0)
	if v1496 == v1505 {
		v1562 = v1496
		v1587 = v1505
		goto L248
	} else {
		goto L249
	}
L248:
	;
	v1613 = v1587 & v1562
	goto L215
L249:
	;
	v1508 = int64(-1)
	v1511 = base.I64_extend_i32_u(v1502) & int64(255)
	v1521 = v1505
	v1541 = v1484
	v1542 = v1486
	goto L251
L250:
	;
	v1562 = base.B2i32(v1236 == v1548)
	v1587 = int32(1)
	goto L248
L251:
	;
	v1548 = v1541 + v1542&(v1508<<(uint(v1511)%64)^v1508) + int64(1)
	if base.Ui64(v1236) <= base.Ui64(v1548) {
		goto L250
	} else {
		goto L253
	}
L252:
	;
	v1562 = v1496
	v1587 = int32(0)
	goto L248
L253:
	;
	v1552 = v1521 + int32(1)
	if v1552 != v1496 {
		v1521 = v1552
		v1541 = v1548
		v1542 = int64(base.Ui64(v1542) >> (uint(v1511) % 64))
		goto L251
	} else {
		goto L254
	}
L254:
	;
	goto L252
L255:
	;
	v1652 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v729+int32(_a_F_gistvacuumscan_15)+v1198<<(uint(v1652)%32)))) = uint16(v1194)
	*(*int32)(unsafe.Add(mBase, uint32(v729+int32(16)+v1198<<(uint(int32(2))%32)))) = v1235
	v1664 = v1198 + v1652
	goto L257
L256:
	;
	v1664 = v1198
	goto L257
L257:
	;
	v1667 = v1194 + int32(1)
	if base.B2i32(v1664 < v1177-v1184)&base.B2i32(base.Ui32(v1667&int32(_a_F_gistvacuumscan_5)) <= base.Ui32(v1177)) != 0 {
		v1194 = v1667
		v1198 = v1664
		goto L199
	} else {
		goto L258
	}
L258:
	;
	goto L200
L259:
	;
	if v1664 <= int32(0) {
		v1961 = v1664
		goto L191
	} else {
		goto L260
	}
L260:
	;
	v1679 = int32(0)
	v1687 = v1679
	v1692 = v1679
	goto L261
L261:
	;
	v1715 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1156)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v1715) {
		goto L263
	} else {
		goto L264
	}
L262:
	;
	v1961 = v1664
	goto L191
L263:
	;
	if (v1715+int32(_a_F_gistvacuumscan_4))&int32(_a_F_gistvacuumscan_8) == int32(4) {
		v1961 = v1664
		goto L191
	} else {
		goto L266
	}
L264:
	;
	goto L265
L265:
	;
	v1724 = int32(0)
	v1730 = *(*int32)(unsafe.Add(mBase, uint32(v729+int32(16)+v1687<<(uint(int32(2))%32))))
	v1732 = *(*int32)(unsafe.Add(mBase, uint32(v715)+24))
	v1733 = F_ReadBufferExtended(m, v731, v1724, v1730, v1724, v1732)
	mBase = m.M
	v1734 = m.ExcPending
	if v1734 != 0 {
		goto L1
	} else {
		goto L267
	}
L266:
	;
	goto L265
L267:
	;
	F_LockBuffer(m, v1733, int32(2))
	mBase = m.M
	v1737 = m.ExcPending
	if v1737 != 0 {
		goto L1
	} else {
		goto L268
	}
L268:
	;
	F_gistcheckpage(m, v731, v1733)
	mBase = m.M
	v1739 = m.ExcPending
	if v1739 != 0 {
		goto L1
	} else {
		goto L269
	}
L269:
	;
	F_LockBuffer(m, v1132, int32(2))
	mBase = m.M
	v1742 = m.ExcPending
	if v1742 != 0 {
		goto L1
	} else {
		goto L270
	}
L270:
	;
	if v1138 == int32(0) {
		goto L272
	} else {
		goto L273
	}
L271:
	;
	v1760 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v729+int32(_a_F_gistvacuumscan_15)+v1687<<(uint(int32(1))%32)))))
	if v1733 < int32(0) {
		goto L277
	} else {
		goto L278
	}
L272:
	;
	v1751 = *(*int32)(unsafe.Add(mBase, _c_F_gistvacuumscan[4]))
	v1753 = *(*int32)(unsafe.Add(mBase, uint32(v1751+(v1132^int32(-1))<<(uint(int32(2))%32))))
	v1759 = v1753
	goto L271
L273:
	;
	goto L274
L274:
	;
	v1755 = *(*int32)(unsafe.Add(mBase, _c_F_gistvacuumscan[5]))
	v1759 = v1755 + v1132<<(uint(int32(13))%32) + int32(-8192)
	goto L271
L275:
	;
	F_LockBuffer(m, v1132, int32(0))
	mBase = m.M
	v1947 = m.ExcPending
	if v1947 != 0 {
		goto L1
	} else {
		goto L315
	}
L276:
	;
	v1779 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1778)+16)))
	v1781 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1779+v1778)+12)))
	if v1781&int32(9) != int32(1) {
		v1942 = v1692
		goto L275
	} else {
		goto L280
	}
L277:
	;
	v1764 = *(*int32)(unsafe.Add(mBase, _c_F_gistvacuumscan[4]))
	v1770 = *(*int32)(unsafe.Add(mBase, uint32(v1764+(v1733^int32(-1))<<(uint(int32(2))%32))))
	v1778 = v1770
	goto L276
L278:
	;
	goto L279
L279:
	;
	v1772 = *(*int32)(unsafe.Add(mBase, _c_F_gistvacuumscan[5]))
	v1778 = v1772 + v1733<<(uint(int32(13))%32) + int32(-8192)
	goto L276
L280:
	;
	v1786 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1778)+12)))
	if (v1786+int32(_a_F_gistvacuumscan_4))&int32(_a_F_gistvacuumscan_8) != 0 {
		goto L281
	} else {
		goto L282
	}
L281:
	;
	v1794 = base.B2i32(base.Ui32(int32(25)) <= base.Ui32(v1786))
	goto L283
L282:
	;
	v1794 = int32(0)
	goto L283
L283:
	;
	if v1794 != 0 {
		v1942 = v1692
		goto L275
	} else {
		goto L284
	}
L284:
	;
	v1795 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1759)+14)))
	if v1795 == int32(0) {
		v1942 = v1692
		goto L275
	} else {
		goto L285
	}
L285:
	;
	v1798 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1759)+16)))
	v1800 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1759+v1798)+12)))
	if v1800&int32(3) != 0 {
		v1942 = v1692
		goto L275
	} else {
		goto L286
	}
L286:
	;
	v1803 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1759)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v1803) {
		goto L287
	} else {
		goto L288
	}
L287:
	;
	v1811 = int32(base.Ui32(v1803+int32(_a_F_gistvacuumscan_4)) >> (uint(int32(2)) % 32))
	goto L289
L288:
	;
	v1811 = int32(0)
	goto L289
L289:
	;
	v1812 = int32(_a_F_gistvacuumscan_5)
	v1813 = v1811 & v1812
	v1816 = (v1760 - v1692) & v1812
	if base.Ui32(v1813) < base.Ui32(v1816) {
		v1942 = v1692
		goto L275
	} else {
		goto L290
	}
L290:
	;
	if base.Ui32(v1813) < base.Ui32(int32(2)) {
		v1942 = v1692
		goto L275
	} else {
		goto L291
	}
L291:
	;
	v1823 = *(*int32)(unsafe.Add(mBase, uint32(v1816<<(uint(int32(2))%32)+v1759)+20))
	if v1733 < int32(0) {
		goto L293
	} else {
		goto L294
	}
L292:
	;
	v1845 = v1759 + v1823&int32(_a_F_gistvacuumscan_6)
	v1846 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1845))))
	v1849 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1845)+2)))
	if v1842 != v1846<<(uint(int32(16))%32)|v1849 {
		v1942 = v1692
		goto L275
	} else {
		goto L296
	}
L293:
	;
	v1827 = *(*int32)(unsafe.Add(mBase, _c_F_gistvacuumscan[2]))
	v1833 = *(*int32)(unsafe.Add(mBase, uint32(v1827+(v1733^int32(-1))<<(uint(int32(6))%32))+16))
	v1842 = v1833
	goto L292
L294:
	;
	goto L295
L295:
	;
	v1835 = *(*int32)(unsafe.Add(mBase, _c_F_gistvacuumscan[3]))
	v1841 = *(*int32)(unsafe.Add(mBase, uint32(v1835+v1733<<(uint(int32(6))%32)+int32(-64))+16))
	v1842 = v1841
	goto L292
L296:
	;
	v1852 = F_ReadNextFullTransactionId(m)
	mBase = m.M
	v1853 = m.ExcPending
	if v1853 != 0 {
		goto L1
	} else {
		goto L297
	}
L297:
	;
	v1854 = int32(_a_F_gistvacuumscan_7)
	v1856 = *(*int32)(unsafe.Add(mBase, _c_F_gistvacuumscan[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_gistvacuumscan[6])) = v1856 + int32(1)
	F_MarkBufferDirty(m, v1733)
	mBase = m.M
	v1861 = m.ExcPending
	if v1861 != 0 {
		goto L1
	} else {
		goto L298
	}
L298:
	;
	v1862 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1778)+16)))
	v1863 = v1778 + v1862
	v1864 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1863)+12)))
	v1866 = v1864 | int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v1863)+12)) = uint16(v1866)
	*(*int64)(unsafe.Add(mBase, uint32(v1778)+24)) = v1852
	v1869 = int32(32)
	*(*uint16)(unsafe.Add(mBase, uint32(v1778)+12)) = uint16(v1869)
	v1871 = *(*int32)(unsafe.Add(mBase, uint32(v716)+24))
	v1872 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v716)+24)) = v1871 + v1872
	v1875 = *(*int32)(unsafe.Add(mBase, uint32(v716)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v716)+28)) = v1875 + v1872
	F_MarkBufferDirty(m, v1132)
	mBase = m.M
	v1880 = m.ExcPending
	if v1880 != 0 {
		goto L1
	} else {
		goto L299
	}
L299:
	;
	F_PageIndexTupleDelete(m, v1759, v1816)
	mBase = m.M
	v1882 = m.ExcPending
	if v1882 != 0 {
		goto L1
	} else {
		goto L300
	}
L300:
	;
	v1883 = *(*int32)(unsafe.Add(mBase, uint32(v715)))
	v1884 = *(*int32)(unsafe.Add(mBase, uint32(v1883)+48))
	v1885 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1884)+118)))
	if v1885 != int32(112) {
		goto L302
	} else {
		goto L303
	}
L301:
	;
	v1925 = int64(32)
	*(*int64)(unsafe.Add(mBase, uint32(v1759))) = base.I64_rotr(v1924, v1925)
	*(*uint32)(unsafe.Add(mBase, uint32(v1778)+4)) = uint32(v1924)
	v1930 = int64(base.Ui64(v1924) >> (uint(v1925) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v1778))) = uint32(v1930)
	v1932 = int32(_a_F_gistvacuumscan_7)
	v1934 = *(*int32)(unsafe.Add(mBase, _c_F_gistvacuumscan[6]))
	v1935 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_gistvacuumscan[6])) = v1934 - v1935
	v1942 = v1692 + v1935
	goto L275
L302:
	;
	v1920 = F_gistGetFakeLSN(m, v1883)
	mBase = m.M
	v1921 = m.ExcPending
	if v1921 != 0 {
		goto L1
	} else {
		goto L314
	}
L303:
	;
	v1889 = *(*int32)(unsafe.Add(mBase, _c_F_gistvacuumscan[1]))
	if v1889 <= int32(0) {
		goto L304
	} else {
		goto L305
	}
L304:
	;
	v1892 = *(*int32)(unsafe.Add(mBase, uint32(v1883)+32))
	if v1892 != 0 {
		goto L302
	} else {
		goto L307
	}
L305:
	;
	goto L306
L306:
	;
	v1894 = m.G0
	v1896 = v1894 - int32(16)
	m.G0 = v1896
	*(*uint16)(unsafe.Add(mBase, uint32(v1896)+8)) = uint16(v1816)
	*(*int64)(unsafe.Add(mBase, uint32(v1896))) = v1852
	F_XLogBeginInsert(m)
	mBase = m.M
	v1901 = m.ExcPending
	if v1901 != 0 {
		goto L1
	} else {
		goto L309
	}
L307:
	;
	v1893 = *(*int32)(unsafe.Add(mBase, uint32(v1883)+40))
	if v1893 != 0 {
		goto L302
	} else {
		goto L308
	}
L308:
	;
	goto L306
L309:
	;
	F_XLogRegisterData(m, v1896, int32(10))
	mBase = m.M
	v1904 = m.ExcPending
	if v1904 != 0 {
		goto L1
	} else {
		goto L310
	}
L310:
	;
	F_XLogRegisterBuffer(m, int32(0), v1733, int32(8))
	mBase = m.M
	v1908 = m.ExcPending
	if v1908 != 0 {
		goto L1
	} else {
		goto L311
	}
L311:
	;
	F_XLogRegisterBuffer(m, int32(1), v1132, int32(8))
	mBase = m.M
	v1912 = m.ExcPending
	if v1912 != 0 {
		goto L1
	} else {
		goto L312
	}
L312:
	;
	v1915 = F_XLogInsert(m, int32(14), int32(96))
	mBase = m.M
	v1916 = m.ExcPending
	if v1916 != 0 {
		goto L1
	} else {
		goto L313
	}
L313:
	;
	m.G0 = v1896 + int32(16)
	v1924 = v1915
	goto L301
L314:
	;
	v1924 = v1920
	goto L301
L315:
	;
	F_UnlockReleaseBuffer(m, v1733)
	mBase = m.M
	v1949 = m.ExcPending
	if v1949 != 0 {
		goto L1
	} else {
		goto L316
	}
L316:
	;
	v1951 = v1687 + int32(1)
	if v1951 != v1664 {
		v1687 = v1951
		v1692 = v1942
		goto L261
	} else {
		goto L317
	}
L317:
	;
	goto L262
L318:
	;
	v2004 = v733 - v1961
	goto L180
L319:
	;
	goto L144
L320:
	;
	m.G0 = v2030 + int32(_a_F_gistvacuumscan_0)
	return
}
func F_gtsquery_compress(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int64
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int64
	_ = v40
	var v43 int64
	_ = v43
	var v44 int32
	_ = v44
	var v48 int64
	_ = v48
	var v51 int64
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int64
	_ = v58
	var v64 int32
	_ = v64
	var v68 int64
	_ = v68
	var v71 int64
	_ = v71
	var v77 int64
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+14)))
	if v4 != int32(1) {
		return v3
	} else {
		v9 = F_palloc(m, int32(16))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
			v14 = int64(0)
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
			if v18 <= int32(0) {
				v77 = int64(0)
			} else {
				v22 = int32(1)
				v25 = v13 + int32(8)
				if v18 == v22 {
					v57 = v25
					v58 = v14
				} else {
					v31 = v25
					v32 = v14
					v33 = int32(0)
					for {
						v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
						if v36 == int32(1) {
							v40 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v31)+4)))
							v43 = int64(1)<<(uint(v40)%64) | v32
						} else {
							v43 = v32
						}
						v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+12)))
						if v44 == int32(1) {
							v48 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v31)+16)))
							v51 = int64(1)<<(uint(v48)%64) | v43
						} else {
							v51 = v43
						}
						v53 = v31 + int32(24)
						v55 = v33 + int32(2)
						if v55 != v18&int32(2147483646) {
							v31 = v53
							v32 = v51
							v33 = v55
							continue
						} else {
							break
						}
						break
					}
					v57 = v53
					v58 = v51
				}
				if v18&v22 == int32(0) {
					v71 = v58
				} else {
					v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
					if v64 != int32(1) {
						v71 = v58
					} else {
						v68 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v57)+4)))
						v71 = int64(1)<<(uint(v68)%64) | v58
					}
				}
				v77 = v71
			}
			v78 = F_Int64GetDatum(m, v77)
			mBase = m.M
			v79 = m.ExcPending
			if v79 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = v78
				v81 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v81
				v83 = *(*int32)(unsafe.Add(mBase, uint32(v3)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v83
				v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3)+12)))
				v86 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v9)+14)) = uint8(v86)
				*(*uint16)(unsafe.Add(mBase, uint32(v9)+12)) = uint16(v85)
				return v9
			}
		}
	}
}
func F_gtsquery_consistent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int64
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int64
	_ = v39
	var v42 int64
	_ = v42
	var v43 int32
	_ = v43
	var v47 int64
	_ = v47
	var v50 int64
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int64
	_ = v57
	var v63 int32
	_ = v63
	var v67 int64
	_ = v67
	var v70 int64
	_ = v70
	var v76 int64
	_ = v76
	var v77 int32
	_ = v77
	var v82 int64
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v94 int64
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	v2 = int64(0)
	v7 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+36)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v11 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v17 <= int32(0) {
		v76 = int64(0)
	} else {
		v21 = int32(1)
		v24 = v12 + int32(8)
		if v17 == v21 {
			v56 = v24
			v57 = v2
		} else {
			v30 = v24
			v31 = v2
			v32 = int32(0)
			for {
				v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
				if v35 == int32(1) {
					v39 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v30)+4)))
					v42 = int64(1)<<(uint(v39)%64) | v31
				} else {
					v42 = v31
				}
				v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+12)))
				if v43 == int32(1) {
					v47 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v30)+16)))
					v50 = int64(1)<<(uint(v47)%64) | v42
				} else {
					v50 = v42
				}
				v52 = v30 + int32(24)
				v54 = v32 + int32(2)
				if v54 != v17&int32(2147483646) {
					v30 = v52
					v31 = v50
					v32 = v54
					continue
				} else {
					break
				}
				break
			}
			v56 = v52
			v57 = v50
		}
		if v17&v21 == int32(0) {
			v70 = v57
		} else {
			v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
			if v63 != int32(1) {
				v70 = v57
			} else {
				v67 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v56)+4)))
				v70 = int64(1)<<(uint(v67)%64) | v57
			}
		}
		v76 = v70
	}
	v77 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v77)
	switch v7 - int32(7) {
	case 0:
		v82 = v11 & v76
		v83 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
		v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v83)+16)))
		v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83+v84)+12)))
		if v86&int32(1) != 0 {
			return base.B2i32(v82 == v76)
		} else {
			return base.B2i32(v82 != int64(0))
		}
	case 1:
		v94 = v11 & v76
		v95 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
		v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v95)+16)))
		v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95+v96)+12)))
		if v98&int32(1) != 0 {
			return base.B2i32(v11 == v94)
		} else {
			v105 = base.B2i32(v94 != int64(0))
			return v105
		}
	default:
		v105 = int32(0)
		return v105
	}
}
