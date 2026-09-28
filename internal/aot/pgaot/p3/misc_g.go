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
func F_GetComment(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
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
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v80 int64
	_ = v80
	var v81 int64
	_ = v81
	var v82 int64
	_ = v82
	var v83 int64
	_ = v83
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v99 int64
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v107 int64
	_ = v107
	var v108 int32
	_ = v108
	var v112 int64
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int64
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	v9 = m.G0
	v11 = v9 - int32(192)
	m.G0 = v11
	v14 = v11 + int32(16)
	F_ScanKeyInit(m, v14, int32(1), int32(3), int32(184), base.I64_extend_i32_u(l0))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		F_ScanKeyInit(m, v11+int32(72), int32(2), int32(3), int32(184), base.I64_extend_i32_u(l1))
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return int32(0)
		} else {
			v33 = int32(3)
			F_ScanKeyInit(m, v11+int32(128), v33, v33, int32(65), base.I64_extend_i32_s(l2))
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int32(0)
			} else {
				v41 = F_table_open(m, int32(2609), int32(1))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v41)+52))
					v44 = int32(0)
					v49 = F_systable_beginscan(m, v41, int32(2675), int32(1), v44, int32(3), v14)
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						v51 = F_systable_getnext(m, v49)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							if v51 == int32(0) {
								v121 = v44
								F_systable_endscan(m, v49)
								mBase = m.M
								v126 = m.ExcPending
								if v126 != 0 {
									return int32(0)
								} else {
									F_relation_close(m, v41, int32(1))
									mBase = m.M
									v129 = m.ExcPending
									if v129 != 0 {
										return int32(0)
									} else {
										m.G0 = v11 + int32(192)
										return v121
									}
								}
							} else {
								v55 = *(*int32)(unsafe.Add(mBase, uint32(v51)+16))
								v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v55)+18)))
								if v56&int32(2044) != 0 {
									v59 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)) = uint8(v59)
									v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+20)))
									if v61&int32(1) == v59 {
										v66 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+52)))
										if int32(0) <= v66 {
											v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+22)))
											v71 = v55 + v69 + v66
											v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+56)))
											if v72 == int32(1) {
												v75 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+54)))
												if base.I32_popcnt(v75) != int32(1) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v87 = m.ExcPending
													if v87 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v11))) = v75
														F_errmsg_internal(m, int32(_a_F_GetComment_0), v11)
														mBase = m.M
														v91 = m.ExcPending
														if v91 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_GetComment_1), int32(123), int32(_a_F_GetComment_2))
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
												} else {
													switch base.I32_ctz(v75) {
													case 0:
														v80 = int64(*(*int8)(unsafe.Add(mBase, uint32(v71))))
														v117 = v80
														v119 = F_text_to_cstring(m, base.I32_wrap_i64(v117))
														mBase = m.M
														v120 = m.ExcPending
														if v120 != 0 {
															return int32(0)
														} else {
															v121 = v119
															F_systable_endscan(m, v49)
															mBase = m.M
															v126 = m.ExcPending
															if v126 != 0 {
																return int32(0)
															} else {
																F_relation_close(m, v41, int32(1))
																mBase = m.M
																v129 = m.ExcPending
																if v129 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v11 + int32(192)
																	return v121
																}
															}
														}
													case 1:
														v81 = int64(*(*int16)(unsafe.Add(mBase, uint32(v71))))
														v117 = v81
														v119 = F_text_to_cstring(m, base.I32_wrap_i64(v117))
														mBase = m.M
														v120 = m.ExcPending
														if v120 != 0 {
															return int32(0)
														} else {
															v121 = v119
															F_systable_endscan(m, v49)
															mBase = m.M
															v126 = m.ExcPending
															if v126 != 0 {
																return int32(0)
															} else {
																F_relation_close(m, v41, int32(1))
																mBase = m.M
																v129 = m.ExcPending
																if v129 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v11 + int32(192)
																	return v121
																}
															}
														}
													case 2:
														v82 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v71))))
														v117 = v82
														v119 = F_text_to_cstring(m, base.I32_wrap_i64(v117))
														mBase = m.M
														v120 = m.ExcPending
														if v120 != 0 {
															return int32(0)
														} else {
															v121 = v119
															F_systable_endscan(m, v49)
															mBase = m.M
															v126 = m.ExcPending
															if v126 != 0 {
																return int32(0)
															} else {
																F_relation_close(m, v41, int32(1))
																mBase = m.M
																v129 = m.ExcPending
																if v129 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v11 + int32(192)
																	return v121
																}
															}
														}
													case 3:
														v83 = *(*int64)(unsafe.Add(mBase, uint32(v71)))
														v117 = v83
														v119 = F_text_to_cstring(m, base.I32_wrap_i64(v117))
														mBase = m.M
														v120 = m.ExcPending
														if v120 != 0 {
															return int32(0)
														} else {
															v121 = v119
															F_systable_endscan(m, v49)
															mBase = m.M
															v126 = m.ExcPending
															if v126 != 0 {
																return int32(0)
															} else {
																F_relation_close(m, v41, int32(1))
																mBase = m.M
																v129 = m.ExcPending
																if v129 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v11 + int32(192)
																	return v121
																}
															}
														}
													default:
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v87 = m.ExcPending
														if v87 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v11))) = v75
															F_errmsg_internal(m, int32(_a_F_GetComment_0), v11)
															mBase = m.M
															v91 = m.ExcPending
															if v91 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_GetComment_1), int32(123), int32(_a_F_GetComment_2))
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
												}
											} else {
												v117 = base.I64_extend_i32_u(v71)
												v119 = F_text_to_cstring(m, base.I32_wrap_i64(v117))
												mBase = m.M
												v120 = m.ExcPending
												if v120 != 0 {
													return int32(0)
												} else {
													v121 = v119
													F_systable_endscan(m, v49)
													mBase = m.M
													v126 = m.ExcPending
													if v126 != 0 {
														return int32(0)
													} else {
														F_relation_close(m, v41, int32(1))
														mBase = m.M
														v129 = m.ExcPending
														if v129 != 0 {
															return int32(0)
														} else {
															m.G0 = v11 + int32(192)
															return v121
														}
													}
												}
											}
										} else {
											v99 = F_nocachegetattr(m, v51, int32(4), v43)
											mBase = m.M
											v100 = m.ExcPending
											if v100 != 0 {
												return int32(0)
											} else {
												v117 = v99
												v119 = F_text_to_cstring(m, base.I32_wrap_i64(v117))
												mBase = m.M
												v120 = m.ExcPending
												if v120 != 0 {
													return int32(0)
												} else {
													v121 = v119
													F_systable_endscan(m, v49)
													mBase = m.M
													v126 = m.ExcPending
													if v126 != 0 {
														return int32(0)
													} else {
														F_relation_close(m, v41, int32(1))
														mBase = m.M
														v129 = m.ExcPending
														if v129 != 0 {
															return int32(0)
														} else {
															m.G0 = v11 + int32(192)
															return v121
														}
													}
												}
											}
										}
									} else {
										v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+23)))
										if v101&int32(8) == int32(0) {
											v121 = v44
											F_systable_endscan(m, v49)
											mBase = m.M
											v126 = m.ExcPending
											if v126 != 0 {
												return int32(0)
											} else {
												F_relation_close(m, v41, int32(1))
												mBase = m.M
												v129 = m.ExcPending
												if v129 != 0 {
													return int32(0)
												} else {
													m.G0 = v11 + int32(192)
													return v121
												}
											}
										} else {
											v107 = F_nocachegetattr(m, v51, int32(4), v43)
											mBase = m.M
											v108 = m.ExcPending
											if v108 != 0 {
												return int32(0)
											} else {
												v117 = v107
												v119 = F_text_to_cstring(m, base.I32_wrap_i64(v117))
												mBase = m.M
												v120 = m.ExcPending
												if v120 != 0 {
													return int32(0)
												} else {
													v121 = v119
													F_systable_endscan(m, v49)
													mBase = m.M
													v126 = m.ExcPending
													if v126 != 0 {
														return int32(0)
													} else {
														F_relation_close(m, v41, int32(1))
														mBase = m.M
														v129 = m.ExcPending
														if v129 != 0 {
															return int32(0)
														} else {
															m.G0 = v11 + int32(192)
															return v121
														}
													}
												}
											}
										}
									}
								} else {
									v112 = F_getmissingattr(m, v43, int32(4), v11+int32(15))
									mBase = m.M
									v113 = m.ExcPending
									if v113 != 0 {
										return int32(0)
									} else {
										v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
										if v114 != 0 {
											v121 = v44
											F_systable_endscan(m, v49)
											mBase = m.M
											v126 = m.ExcPending
											if v126 != 0 {
												return int32(0)
											} else {
												F_relation_close(m, v41, int32(1))
												mBase = m.M
												v129 = m.ExcPending
												if v129 != 0 {
													return int32(0)
												} else {
													m.G0 = v11 + int32(192)
													return v121
												}
											}
										} else {
											v117 = v112
											v119 = F_text_to_cstring(m, base.I32_wrap_i64(v117))
											mBase = m.M
											v120 = m.ExcPending
											if v120 != 0 {
												return int32(0)
											} else {
												v121 = v119
												F_systable_endscan(m, v49)
												mBase = m.M
												v126 = m.ExcPending
												if v126 != 0 {
													return int32(0)
												} else {
													F_relation_close(m, v41, int32(1))
													mBase = m.M
													v129 = m.ExcPending
													if v129 != 0 {
														return int32(0)
													} else {
														m.G0 = v11 + int32(192)
														return v121
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
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v122 int32
	_ = v122
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	v2 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_GetNonHistoricCatalogSnapshot[0]))
	if v8 == v2 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v141
L2:
	;
	v131 = F_GetSnapshotData(m, int32(_a_F_GetNonHistoricCatalogSnapshot_0))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L30
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
	v122 = *(*int32)(unsafe.Add(mBase, _c_F_GetNonHistoricCatalogSnapshot[0]))
	if v122 != 0 {
		v141 = v122
		goto L1
	} else {
		goto L42
	}
L5:
	;
	if v29 != 0 {
		goto L4
	} else {
		goto L15
	}
L6:
	;
	goto L5
L7:
	;
	v29 = int32(0)
	goto L6
L8:
	;
	if base.B2i32(l0 == int32(1214))|base.B2i32(base.Ui32(l0-int32(2608)) < base.Ui32(int32(2))) != 0 {
		v29 = v12
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
		v29 = v12
		goto L6
	case 1, 2, 3:
		goto L7
	default:
		goto L13
	}
L11:
	;
	if l0 != int32(2396) {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v29 = v12
	goto L6
L13:
	;
	if l0 == int32(2964) {
		v29 = v12
		goto L6
	} else {
		goto L14
	}
L14:
	;
	goto L7
L15:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_GetNonHistoricCatalogSnapshot[1]))
	v33 = v31 - int32(1)
	if v33 < int32(0) {
		v65 = v2
		goto L16
	} else {
		goto L17
	}
L16:
	;
	if v65 != 0 {
		goto L4
	} else {
		goto L28
	}
L17:
	;
	v37 = v33
	v38 = v2
	goto L18
L18:
	;
	v43 = int32(2)
	v44 = base.I32_div_s(v37-v38, v43)
	v45 = v44 + v38
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v45<<(uint(v43)%32))+uint32(_c_F_GetNonHistoricCatalogSnapshot[2])))
	v51 = base.B2i32(v50 == l0)
	if v50 == l0 {
		v65 = v51
		goto L16
	} else {
		goto L20
	}
L19:
	;
	v65 = v51
	goto L16
L20:
	;
	v54 = base.B2i32(base.Ui32(v50) < base.Ui32(l0))
	if base.Ui32(v50) < base.Ui32(l0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v55 = v45 + int32(1)
	goto L23
L22:
	;
	v55 = v38
	goto L23
L23:
	;
	if base.Ui32(v50) < base.Ui32(l0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v58 = v37
	goto L26
L25:
	;
	v58 = v45 - int32(1)
	goto L26
L26:
	;
	if v55 <= v58 {
		v37 = v58
		v38 = v55
		goto L18
	} else {
		goto L27
	}
L27:
	;
	goto L19
L28:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_GetNonHistoricCatalogSnapshot[0]))
	if v67 == int32(0) {
		goto L2
	} else {
		goto L29
	}
L29:
	;
	F_pairingheap_remove(m, int32(_a_F_GetNonHistoricCatalogSnapshot_1), v67+int32(52))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	return int32(0)
L31:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_GetNonHistoricCatalogSnapshot[0])) = int32(0)
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_GetNonHistoricCatalogSnapshot[3]))
	if v81 != 0 {
		goto L2
	} else {
		goto L32
	}
L32:
	;
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_GetNonHistoricCatalogSnapshot[4]))
	if v83 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v87 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_GetNonHistoricCatalogSnapshot[5])) = v87
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_GetNonHistoricCatalogSnapshot[6]))
	*(*int32)(unsafe.Add(mBase, uint32(v90)+52)) = v87
	goto L4
L34:
	;
	goto L35
L35:
	;
	v94 = *(*int32)(unsafe.Add(mBase, _c_F_GetNonHistoricCatalogSnapshot[6]))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+52))
	v96 = int32(3)
	v99 = *(*int32)(unsafe.Add(mBase, _c_F_GetNonHistoricCatalogSnapshot[4]))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v99-int32(48))))
	if base.B2i32(base.Ui32(v95) < base.Ui32(v96))|base.B2i32(base.Ui32(v102) < base.Ui32(v96)) == int32(0) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_GetNonHistoricCatalogSnapshot[5])) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v94)+52)) = v102
	goto L4
L37:
	;
	if v95-v102 < int32(0) {
		goto L36
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	if base.Ui32(v102) <= base.Ui32(v95) {
		goto L4
	} else {
		goto L41
	}
L40:
	;
	goto L4
L41:
	;
	goto L36
L42:
	;
	goto L2
L43:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_GetNonHistoricCatalogSnapshot[0])) = v131
	F_pairingheap_add(m, int32(_a_F_GetNonHistoricCatalogSnapshot_1), v131+int32(52))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L30
	} else {
		goto L44
	}
L44:
	;
	v140 = *(*int32)(unsafe.Add(mBase, _c_F_GetNonHistoricCatalogSnapshot[0]))
	v141 = v140
	goto L1
}
func F_GetRunningTransactionData(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
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
	var v57 int64
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int64
	_ = v230
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	v1 = int32(0)
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_GetRunningTransactionData[0]))
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_GetRunningTransactionData[1]))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_GetRunningTransactionData[2]))
	if v22 == v1 {
		v27 = *(*int32)(unsafe.Add(mBase, _c_F_GetRunningTransactionData[3]))
		v29 = *(*int32)(unsafe.Add(mBase, _c_F_GetRunningTransactionData[4]))
		v33 = F_emscripten_builtin_malloc(m, (v27+v29)*int32(260))
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, _c_F_GetRunningTransactionData[2])) = v33
		if v33 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v244 = m.ExcPending
			if v244 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(_a_F_GetRunningTransactionData_0))
				mBase = m.M
				v247 = m.ExcPending
				if v247 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_GetRunningTransactionData_1), int32(0))
					mBase = m.M
					v251 = m.ExcPending
					if v251 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_GetRunningTransactionData_2), int32(2679), int32(_a_F_GetRunningTransactionData_3))
						mBase = m.M
						v256 = m.ExcPending
						if v256 != 0 {
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
			v37 = v33
			v39 = *(*int32)(unsafe.Add(mBase, _c_F_GetRunningTransactionData[5]))
			v43 = F_LWLockAcquire(m, v39+int32(512), int32(1))
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return int32(0)
			} else {
				v48 = *(*int32)(unsafe.Add(mBase, _c_F_GetRunningTransactionData[5]))
				v52 = F_LWLockAcquire(m, v48+int32(384), int32(1))
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					v55 = *(*int32)(unsafe.Add(mBase, _c_F_GetRunningTransactionData[6]))
					v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
					v57 = *(*int64)(unsafe.Add(mBase, uint32(v55)+48))
					v58 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
					if v58 <= int32(0) {
						v206 = v1
						v208 = v1
						v209 = v56
						v210 = v56
						v212 = v1
					} else {
						v64 = *(*int32)(unsafe.Add(mBase, _c_F_GetRunningTransactionData[7]))
						v66 = v58
						v67 = v1
						v68 = v56
						v69 = v56
						v70 = v1
						v71 = v1
						for {
							v81 = v70 << (uint(int32(2)) % 32)
							v83 = *(*int32)(unsafe.Add(mBase, uint32(v20+v81)))
							if v83 != 0 {
								if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v83))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v69)) != 0 {
									v93 = int32(base.Ui32(v83-v69) >> (uint(int32(31)) % 32))
								} else {
									v93 = base.B2i32(base.Ui32(v83) < base.Ui32(v69))
								}
								v94 = int32(3)
								if base.B2i32(base.Ui32(v68) < base.Ui32(v94))|base.B2i32(base.Ui32(v83) < base.Ui32(v94)) == int32(0) {
									if v83-v68 < int32(0) {
										v106 = *(*int32)(unsafe.Add(mBase, uint32(v17+int32(36)+v81)))
										v110 = *(*int32)(unsafe.Add(mBase, uint32(v64+v106*int32(768))+20))
										v112 = *(*int32)(unsafe.Add(mBase, _c_F_GetRunningTransactionData[8]))
										if v110 == v112 {
											v114 = v83
										} else {
											v114 = v68
										}
										v115 = v114
									} else {
										v115 = v68
									}
								} else {
									if base.Ui32(v68) <= base.Ui32(v83) {
										v115 = v68
									} else {
										v106 = *(*int32)(unsafe.Add(mBase, uint32(v17+int32(36)+v81)))
										v110 = *(*int32)(unsafe.Add(mBase, uint32(v64+v106*int32(768))+20))
										v112 = *(*int32)(unsafe.Add(mBase, _c_F_GetRunningTransactionData[8]))
										if v110 == v112 {
											v114 = v83
										} else {
											v114 = v68
										}
										v115 = v114
									}
								}
								if v93 != 0 {
									v116 = v83
								} else {
									v116 = v69
								}
								v118 = *(*int32)(unsafe.Add(mBase, _c_F_GetRunningTransactionData[1]))
								v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)+8))
								v120 = int32(1)
								v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119+v70<<(uint(v120)%32))+1)))
								*(*int32)(unsafe.Add(mBase, uint32(v37+v71<<(uint(int32(2))%32)))) = v83
								v129 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
								v132 = v129
								v133 = v123 | v67
								v134 = v115
								v135 = v116
								v136 = v71 + v120
							} else {
								v132 = v66
								v133 = v67
								v134 = v68
								v135 = v69
								v136 = v71
							}
							v138 = v70 + int32(1)
							if v138 < v132 {
								v66 = v132
								v67 = v133
								v68 = v134
								v69 = v135
								v70 = v138
								v71 = v136
								continue
							} else {
								break
							}
							break
						}
						if v133&int32(1) != 0 {
							v206 = int32(2)
							v208 = int32(0)
							v209 = v134
							v210 = v135
							v212 = v136
						} else {
							v144 = int32(0)
							if v132 <= v144 {
								v206 = v144
								v208 = int32(0)
								v209 = v134
								v210 = v135
								v212 = v136
							} else {
								v149 = *(*int32)(unsafe.Add(mBase, _c_F_GetRunningTransactionData[1]))
								v150 = *(*int32)(unsafe.Add(mBase, uint32(v149)+8))
								v154 = v144
								v155 = v132
								v156 = int32(0)
								v160 = v136
								for {
									v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150+v154<<(uint(int32(1))%32)))))
									if v172 != 0 {
										v173 = int32(2)
										v176 = *(*int32)(unsafe.Add(mBase, uint32(v17+int32(36)+v154<<(uint(v173)%32))))
										v178 = *(*int32)(unsafe.Add(mBase, _c_F_GetRunningTransactionData[7]))
										v179 = int32(0)
										v182 = base.AtomicRmwOr32(m, v179, int32(_a_F_GetRunningTransactionData_4), v179)
										v184 = v172 << (uint(v173) % 32)
										if v184 != 0 {
											base.MemoryCopy(m, v37+v160<<(uint(int32(2))%32), v178+v176*int32(768)+int32(60), v184)
										} else {
										}
										v195 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
										v197 = v195
										v198 = v156 + v172
										v199 = v172 + v160
									} else {
										v197 = v155
										v198 = v156
										v199 = v160
									}
									v203 = v154 + int32(1)
									if v203 < v197 {
										v154 = v203
										v155 = v197
										v156 = v198
										v160 = v199
										continue
									} else {
										break
									}
									break
								}
								v206 = int32(0)
								v208 = v198
								v209 = v134
								v210 = v135
								v212 = v199
							}
						}
					}
					*(*int32)(unsafe.Add(mBase, _c_F_GetRunningTransactionData[9])) = v206
					*(*int32)(unsafe.Add(mBase, _c_F_GetRunningTransactionData[10])) = v208
					v225 = int32(_a_F_GetRunningTransactionData_5)
					*(*int32)(unsafe.Add(mBase, _c_F_GetRunningTransactionData[11])) = v212 - v208
					v229 = *(*int32)(unsafe.Add(mBase, _c_F_GetRunningTransactionData[6]))
					v230 = *(*int64)(unsafe.Add(mBase, uint32(v229)+8))
					*(*uint32)(unsafe.Add(mBase, _c_F_GetRunningTransactionData[12])) = uint32(v57)
					*(*int32)(unsafe.Add(mBase, _c_F_GetRunningTransactionData[13])) = v209
					*(*int32)(unsafe.Add(mBase, _c_F_GetRunningTransactionData[14])) = v210
					*(*uint32)(unsafe.Add(mBase, _c_F_GetRunningTransactionData[15])) = uint32(v230)
					return v225
				}
			}
		}
	} else {
		v37 = v22
		v39 = *(*int32)(unsafe.Add(mBase, _c_F_GetRunningTransactionData[5]))
		v43 = F_LWLockAcquire(m, v39+int32(512), int32(1))
		mBase = m.M
		v46 = m.ExcPending
		if v46 != 0 {
			return int32(0)
		} else {
			v48 = *(*int32)(unsafe.Add(mBase, _c_F_GetRunningTransactionData[5]))
			v52 = F_LWLockAcquire(m, v48+int32(384), int32(1))
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return int32(0)
			} else {
				v55 = *(*int32)(unsafe.Add(mBase, _c_F_GetRunningTransactionData[6]))
				v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
				v57 = *(*int64)(unsafe.Add(mBase, uint32(v55)+48))
				v58 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
				if v58 <= int32(0) {
					v206 = v1
					v208 = v1
					v209 = v56
					v210 = v56
					v212 = v1
				} else {
					v64 = *(*int32)(unsafe.Add(mBase, _c_F_GetRunningTransactionData[7]))
					v66 = v58
					v67 = v1
					v68 = v56
					v69 = v56
					v70 = v1
					v71 = v1
					for {
						v81 = v70 << (uint(int32(2)) % 32)
						v83 = *(*int32)(unsafe.Add(mBase, uint32(v20+v81)))
						if v83 != 0 {
							if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v83))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v69)) != 0 {
								v93 = int32(base.Ui32(v83-v69) >> (uint(int32(31)) % 32))
							} else {
								v93 = base.B2i32(base.Ui32(v83) < base.Ui32(v69))
							}
							v94 = int32(3)
							if base.B2i32(base.Ui32(v68) < base.Ui32(v94))|base.B2i32(base.Ui32(v83) < base.Ui32(v94)) == int32(0) {
								if v83-v68 < int32(0) {
									v106 = *(*int32)(unsafe.Add(mBase, uint32(v17+int32(36)+v81)))
									v110 = *(*int32)(unsafe.Add(mBase, uint32(v64+v106*int32(768))+20))
									v112 = *(*int32)(unsafe.Add(mBase, _c_F_GetRunningTransactionData[8]))
									if v110 == v112 {
										v114 = v83
									} else {
										v114 = v68
									}
									v115 = v114
								} else {
									v115 = v68
								}
							} else {
								if base.Ui32(v68) <= base.Ui32(v83) {
									v115 = v68
								} else {
									v106 = *(*int32)(unsafe.Add(mBase, uint32(v17+int32(36)+v81)))
									v110 = *(*int32)(unsafe.Add(mBase, uint32(v64+v106*int32(768))+20))
									v112 = *(*int32)(unsafe.Add(mBase, _c_F_GetRunningTransactionData[8]))
									if v110 == v112 {
										v114 = v83
									} else {
										v114 = v68
									}
									v115 = v114
								}
							}
							if v93 != 0 {
								v116 = v83
							} else {
								v116 = v69
							}
							v118 = *(*int32)(unsafe.Add(mBase, _c_F_GetRunningTransactionData[1]))
							v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)+8))
							v120 = int32(1)
							v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119+v70<<(uint(v120)%32))+1)))
							*(*int32)(unsafe.Add(mBase, uint32(v37+v71<<(uint(int32(2))%32)))) = v83
							v129 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
							v132 = v129
							v133 = v123 | v67
							v134 = v115
							v135 = v116
							v136 = v71 + v120
						} else {
							v132 = v66
							v133 = v67
							v134 = v68
							v135 = v69
							v136 = v71
						}
						v138 = v70 + int32(1)
						if v138 < v132 {
							v66 = v132
							v67 = v133
							v68 = v134
							v69 = v135
							v70 = v138
							v71 = v136
							continue
						} else {
							break
						}
						break
					}
					if v133&int32(1) != 0 {
						v206 = int32(2)
						v208 = int32(0)
						v209 = v134
						v210 = v135
						v212 = v136
					} else {
						v144 = int32(0)
						if v132 <= v144 {
							v206 = v144
							v208 = int32(0)
							v209 = v134
							v210 = v135
							v212 = v136
						} else {
							v149 = *(*int32)(unsafe.Add(mBase, _c_F_GetRunningTransactionData[1]))
							v150 = *(*int32)(unsafe.Add(mBase, uint32(v149)+8))
							v154 = v144
							v155 = v132
							v156 = int32(0)
							v160 = v136
							for {
								v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150+v154<<(uint(int32(1))%32)))))
								if v172 != 0 {
									v173 = int32(2)
									v176 = *(*int32)(unsafe.Add(mBase, uint32(v17+int32(36)+v154<<(uint(v173)%32))))
									v178 = *(*int32)(unsafe.Add(mBase, _c_F_GetRunningTransactionData[7]))
									v179 = int32(0)
									v182 = base.AtomicRmwOr32(m, v179, int32(_a_F_GetRunningTransactionData_4), v179)
									v184 = v172 << (uint(v173) % 32)
									if v184 != 0 {
										base.MemoryCopy(m, v37+v160<<(uint(int32(2))%32), v178+v176*int32(768)+int32(60), v184)
									} else {
									}
									v195 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
									v197 = v195
									v198 = v156 + v172
									v199 = v172 + v160
								} else {
									v197 = v155
									v198 = v156
									v199 = v160
								}
								v203 = v154 + int32(1)
								if v203 < v197 {
									v154 = v203
									v155 = v197
									v156 = v198
									v160 = v199
									continue
								} else {
									break
								}
								break
							}
							v206 = int32(0)
							v208 = v198
							v209 = v134
							v210 = v135
							v212 = v199
						}
					}
				}
				*(*int32)(unsafe.Add(mBase, _c_F_GetRunningTransactionData[9])) = v206
				*(*int32)(unsafe.Add(mBase, _c_F_GetRunningTransactionData[10])) = v208
				v225 = int32(_a_F_GetRunningTransactionData_5)
				*(*int32)(unsafe.Add(mBase, _c_F_GetRunningTransactionData[11])) = v212 - v208
				v229 = *(*int32)(unsafe.Add(mBase, _c_F_GetRunningTransactionData[6]))
				v230 = *(*int64)(unsafe.Add(mBase, uint32(v229)+8))
				*(*uint32)(unsafe.Add(mBase, _c_F_GetRunningTransactionData[12])) = uint32(v57)
				*(*int32)(unsafe.Add(mBase, _c_F_GetRunningTransactionData[13])) = v209
				*(*int32)(unsafe.Add(mBase, _c_F_GetRunningTransactionData[14])) = v210
				*(*uint32)(unsafe.Add(mBase, _c_F_GetRunningTransactionData[15])) = uint32(v230)
				return v225
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
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = F_SearchSysCache1(m, int32(11), base.I64_extend_i32_u(l0))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if v12 == int32(0) {
			if l1 != 0 {
				v44 = int32(0)
				m.G0 = v8 + int32(16)
				return v44
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(67137668))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
						F_errmsg(m, int32(_a_F_GetUserNameFromId_0), v8)
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_GetUserNameFromId_1), int32(1000), int32(_a_F_GetUserNameFromId_2))
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
		} else {
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
			v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+22)))
			v39 = F_pstrdup(m, v34+v35+int32(4))
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int32(0)
			} else {
				F_ReleaseCatCache(m, v12)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					v44 = v39
					m.G0 = v8 + int32(16)
					return v44
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
	var v62 int32
	_ = v62
	var v63 int64
	_ = v63
	var v67 int64
	_ = v67
	var v68 int64
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v84 int64
	_ = v84
	var v86 int32
	_ = v86
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_GlobalVisCheckRemovableXid[0]))
	v63 = *(*int64)(unsafe.Add(mBase, uint32(v62)))
	v67 = v63 + base.I64_extend_i32_s(l0-base.I32_wrap_i64(v63))
	v68 = *(*int64)(unsafe.Add(mBase, uint32(v62)+8))
	if base.Ui64(v67) < base.Ui64(v68) {
		v86 = int32(1)
		m.G0 = v9 + int32(48)
		return v86
	} else {
		v71 = int32(0)
		if base.Ui64(v63) <= base.Ui64(v67) {
			v86 = v71
			m.G0 = v9 + int32(48)
			return v86
		} else {
			v74 = *(*int32)(unsafe.Add(mBase, _c_F_GlobalVisCheckRemovableXid[1]))
			if v74 != 0 {
				v76 = *(*int32)(unsafe.Add(mBase, _c_F_GlobalVisCheckRemovableXid[2]))
				if v76 == v74 {
					v86 = v71
					m.G0 = v9 + int32(48)
					return v86
				} else {
					F_ComputeXidHorizons(m, v9+int32(8))
					mBase = m.M
					v83 = m.ExcPending
					if v83 != 0 {
						return int32(0)
					} else {
						v84 = *(*int64)(unsafe.Add(mBase, uint32(v62)+8))
						v86 = base.B2i32(base.Ui64(v67) < base.Ui64(v84))
						m.G0 = v9 + int32(48)
						return v86
					}
				}
			} else {
				F_ComputeXidHorizons(m, v9+int32(8))
				mBase = m.M
				v83 = m.ExcPending
				if v83 != 0 {
					return int32(0)
				} else {
					v84 = *(*int64)(unsafe.Add(mBase, uint32(v62)+8))
					v86 = base.B2i32(base.Ui64(v67) < base.Ui64(v84))
					m.G0 = v9 + int32(48)
					return v86
				}
			}
		}
	}
}
func F_g_intbig_penalty(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	v3 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+56)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v9 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = int32(0)
	if v11 == v12 {
		v28 = v12
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
		if v15 == int32(0) {
			v28 = v12
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
			if v18 != int32(7) {
				v28 = v12
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
				if v21 != int32(17) {
					v28 = v12
				} else {
					v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+32)))
					v28 = v24 ^ int32(1)
				}
			}
		}
	}
	if v28&int32(1) != 0 {
		v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v32 = F_get_fn_opclass_options(m, v31)
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return int64(0)
		} else {
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
			v38 = v36
			v40 = Fn14312(m, base.I32_wrap_i64(v6), base.I32_wrap_i64(v9), v38, int32(4))
			mBase = m.M
			*(*float32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v3)))) = base.F32_convert_i32_s(v40)
			return v3
		}
	} else {
		v38 = int32(252)
		v40 = Fn14312(m, base.I32_wrap_i64(v6), base.I32_wrap_i64(v9), v38, int32(4))
		mBase = m.M
		*(*float32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v3)))) = base.F32_convert_i32_s(v40)
		return v3
	}
}
func F_generate_useful_gather_paths(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 float64
	_ = v34
	var v35 int32
	_ = v35
	var v36 float64
	_ = v36
	var v38 int32
	_ = v38
	var v44 float64
	_ = v44
	var v48 float64
	_ = v48
	var v50 float64
	_ = v50
	var v52 float64
	_ = v52
	var v53 float64
	_ = v53
	var v62 float64
	_ = v62
	var v66 float64
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 float64
	_ = v107
	var v108 int32
	_ = v108
	var v109 float64
	_ = v109
	var v111 int32
	_ = v111
	var v117 float64
	_ = v117
	var v121 float64
	_ = v121
	var v123 float64
	_ = v123
	var v125 float64
	_ = v125
	var v126 float64
	_ = v126
	var v135 float64
	_ = v135
	var v139 float64
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v179 int32
	_ = v179
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v286 int32
	_ = v286
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v313 int32
	_ = v313
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v367 int32
	_ = v367
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v450 int32
	_ = v450
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v491 float64
	_ = v491
	var v492 int32
	_ = v492
	var v493 float64
	_ = v493
	var v495 int32
	_ = v495
	var v501 float64
	_ = v501
	var v505 float64
	_ = v505
	var v507 float64
	_ = v507
	var v509 float64
	_ = v509
	var v510 float64
	_ = v510
	var v519 float64
	_ = v519
	var v523 float64
	_ = v523
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	v4 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v19 == v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v17 + int32(16)
	return
L2:
	;
	v22 = m.G0
	v24 = v22 - int32(16)
	m.G0 = v24
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v26 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v24 + int32(16)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v168 == int32(0) {
		goto L1
	} else {
		goto L41
	}
L4:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v34 = *(*float64)(unsafe.Add(mBase, uint32(v30)+32))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v30)+24))
	v36 = base.F64_convert_i32_s(v35)
	v38 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_generate_useful_gather_paths[0])))
	if v38 == int32(1) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v24)+8)) = v66
	v68 = int32(0)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if l2 != 0 {
		goto L15
	} else {
		goto L16
	}
L6:
	;
	v44 = base.F64_add(base.F64_mul(v36, float64(-0.3)), float64(1))
	if base.F64_gt(v44, float64(0)) != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v50 = v36
	goto L8
L8:
	;
	v52 = float64(1e+100)
	v53 = base.F64_mul(v34, v50)
	if base.F64_gt(v53, v52)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v53)&int64(9223372036854775807))) != 0 {
		v66 = v52
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v48 = v44
	goto L11
L10:
	;
	v48 = math.Float64frombits(uint64(0x8000000000000000))
	goto L11
L11:
	;
	v50 = base.F64_add(v48, v36)
	goto L8
L12:
	;
	goto L5
L13:
	;
	v62 = float64(1)
	if base.F64_le(v53, v62) != 0 {
		v66 = v62
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v66 = base.F64_nearest(v53)
	goto L12
L15:
	;
	v73 = v24 + int32(8)
	goto L17
L16:
	;
	v73 = v68
	goto L17
L17:
	;
	v74 = F_create_gather_path(m, l0, l1, v30, v69, v73)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return
L19:
	;
	F_add_path(m, l1, v74)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v78 == int32(0) {
		goto L3
	} else {
		goto L21
	}
L21:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	if v81 <= int32(0) {
		goto L3
	} else {
		goto L22
	}
L22:
	;
	v89 = v68
	goto L23
L23:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v98+v89<<(uint(int32(2))%32))))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+64))
	if v103 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	goto L3
L25:
	;
	v107 = *(*float64)(unsafe.Add(mBase, uint32(v102)+32))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v102)+24))
	v109 = base.F64_convert_i32_s(v108)
	v111 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_generate_useful_gather_paths[0])))
	if v111 == int32(1) {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	goto L27
L27:
	;
	v148 = v89 + int32(1)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	if v148 < v149 {
		v89 = v148
		goto L23
	} else {
		goto L40
	}
L28:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v24)+8)) = v139
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v102)+64))
	v143 = F_create_gather_merge_path(m, l0, l1, v102, v141, v142, v73)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L18
	} else {
		goto L38
	}
L29:
	;
	v117 = base.F64_add(base.F64_mul(v109, float64(-0.3)), float64(1))
	if base.F64_gt(v117, float64(0)) != 0 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v123 = v109
	goto L31
L31:
	;
	v125 = float64(1e+100)
	v126 = base.F64_mul(v107, v123)
	if base.F64_gt(v126, v125)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v126)&int64(9223372036854775807))) != 0 {
		v139 = v125
		goto L35
	} else {
		goto L36
	}
L32:
	;
	v121 = v117
	goto L34
L33:
	;
	v121 = math.Float64frombits(uint64(0x8000000000000000))
	goto L34
L34:
	;
	v123 = base.F64_add(v121, v109)
	goto L31
L35:
	;
	goto L28
L36:
	;
	v135 = float64(1)
	if base.F64_le(v126, v135) != 0 {
		v139 = v135
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v139 = base.F64_nearest(v126)
	goto L35
L38:
	;
	F_add_path(m, l1, v143)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L18
	} else {
		goto L39
	}
L39:
	;
	goto L27
L40:
	;
	goto L24
L41:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v168)+4))
	if int32(0) < v171 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v323 = F_lappend(m, int32(0), v313)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L18
	} else {
		goto L80
	}
L43:
	;
	v179 = int32(0)
	goto L46
L44:
	;
	goto L45
L45:
	;
	if v171 != 0 {
		goto L1
	} else {
		goto L79
	}
L46:
	;
	v189 = int32(0)
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v168)+12))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v191+v179<<(uint(int32(2))%32))))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v195)+4))
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196)+41)))
	if v197 != 0 {
		v286 = v189
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v299 != 0 {
		goto L73
	} else {
		goto L74
	}
L48:
	;
	if v286 != 0 {
		goto L69
	} else {
		goto L70
	}
L49:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v198)+4))
	if v199 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v199)+4))
	if int32(0) < v200 {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	v257 = v189
	goto L52
L52:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v271 = F_find_computable_ec_member(m, l0, v196, v257, v269, int32(1))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L18
	} else {
		goto L66
	}
L53:
	;
	v206 = v189
	goto L56
L54:
	;
	goto L55
L55:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v198)+4))
	v257 = v253
	goto L52
L56:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v199)+12))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v217+v206<<(uint(int32(2))%32))))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v223 = F_find_ec_member_matching_expr(m, v196, v221, v222)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L18
	} else {
		goto L59
	}
L57:
	;
	goto L55
L58:
	;
	v236 = v206 + int32(1)
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v199)+4))
	if v236 < v237 {
		v206 = v236
		goto L56
	} else {
		goto L65
	}
L59:
	;
	if v223 == int32(0) {
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v223)+4))
	v228 = F_expression_returns_set(m, v227)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L18
	} else {
		goto L61
	}
L61:
	;
	if v228 != 0 {
		goto L58
	} else {
		goto L62
	}
L62:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v223)+4))
	v232 = F_is_parallel_safe(m, l0, v231)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L18
	} else {
		goto L63
	}
L63:
	;
	if v232 != 0 {
		v286 = int32(1)
		goto L48
	} else {
		goto L64
	}
L64:
	;
	goto L58
L65:
	;
	goto L57
L66:
	;
	if v271 == int32(0) {
		v286 = int32(0)
		goto L48
	} else {
		goto L67
	}
L67:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v271)+4))
	v276 = F_expression_returns_set(m, v275)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L18
	} else {
		goto L68
	}
L68:
	;
	v286 = v276 ^ int32(1)
	goto L48
L69:
	;
	v295 = v179 + int32(1)
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v168)+4))
	if v295 < v296 {
		v179 = v295
		goto L46
	} else {
		goto L72
	}
L70:
	;
	v298 = v179
	goto L71
L71:
	;
	goto L47
L72:
	;
	v298 = v295
	goto L71
L73:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v299)+4))
	v302 = v300
	goto L75
L74:
	;
	v302 = int32(0)
	goto L75
L75:
	;
	if v302 == v298 {
		v313 = v299
		goto L42
	} else {
		goto L76
	}
L76:
	;
	if v298 <= int32(0) {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v306 = F_list_copy_head(m, v299, v298)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L18
	} else {
		goto L78
	}
L78:
	;
	v313 = v306
	goto L42
L79:
	;
	v313 = v168
	goto L42
L80:
	;
	if v323 == int32(0) {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v323)+4))
	if v327 <= int32(0) {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	if l2 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v333 = v17 + int32(8)
	goto L85
L84:
	;
	v333 = int32(0)
	goto L85
L85:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v334)+12))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v335)))
	v349 = v4
	goto L86
L86:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v351 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	goto L1
L88:
	;
	v553 = v349 + int32(1)
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v323)+4))
	if v553 < v554 {
		v349 = v553
		goto L86
	} else {
		goto L155
	}
L89:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v351)+4))
	if v354 <= int32(0) {
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v323)+12))
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v357+v349<<(uint(int32(2))%32))))
	v367 = int32(0)
	goto L91
L91:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v351)+12))
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v377+v367<<(uint(int32(2))%32))))
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v381)+64))
	v384 = v17 + int32(4)
	if v361 == v382 {
		goto L96
	} else {
		goto L97
	}
L92:
	;
	goto L88
L93:
	;
	v535 = v367 + int32(1)
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v351)+4))
	if v535 < v536 {
		v367 = v535
		goto L91
	} else {
		goto L154
	}
L94:
	;
	if v462 != 0 {
		goto L93
	} else {
		goto L126
	}
L95:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v361)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v384))) = v450
	v462 = int32(1)
	goto L94
L96:
	;
	if v361 != 0 {
		goto L95
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	if v361 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v384))) = int32(0)
	v462 = int32(1)
	goto L94
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v384))) = int32(0)
	v462 = int32(1)
	goto L94
L101:
	;
	goto L102
L102:
	;
	if v382 == int32(0) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v402 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v384))) = v402
	v462 = v402
	goto L94
L104:
	;
	goto L105
L105:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v382)+4))
	v406 = int32(0)
	if v406 < v405 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v409 = v405
	goto L108
L107:
	;
	v409 = v406
	goto L108
L108:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v361)+4))
	v415 = int32(0)
	goto L109
L109:
	;
	if v415 < v410 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v361)+12))
	v426 = v422 + v415<<(uint(int32(2))%32)
	goto L113
L112:
	;
	v426 = int32(0)
	goto L113
L113:
	;
	if v415 == v409 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v384))) = v409
	v462 = base.B2i32(v426 == int32(0))
	goto L94
L115:
	;
	goto L116
L116:
	;
	v432 = base.B2i32(v426 == int32(0))
	if v426 == int32(0) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v384))) = v415
	v462 = v432
	goto L94
L118:
	;
	goto L119
L119:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v382)+12))
	if v436 == int32(0) {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v384))) = v415
	v462 = v432
	goto L94
L121:
	;
	goto L122
L122:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v426)))
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v436+v415<<(uint(int32(2))%32))))
	if v440 != v444 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v384))) = v415
	v462 = int32(0)
	goto L94
L124:
	;
	v415 = v415 + int32(1)
	goto L109
L126:
	;
	v464 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_generate_useful_gather_paths[1])))
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v381 == v336 {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	if v474&int32(1) != 0 {
		goto L134
	} else {
		goto L135
	}
L128:
	;
	v474 = v464
	goto L127
L129:
	;
	goto L130
L130:
	;
	if v465 == int32(0) {
		goto L93
	} else {
		goto L131
	}
L131:
	;
	v469 = int32(1)
	if v464&v469 == int32(0) {
		goto L93
	} else {
		goto L132
	}
L132:
	;
	v474 = v469
	goto L127
L133:
	;
	v491 = *(*float64)(unsafe.Add(mBase, uint32(v487)+32))
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v487)+24))
	v493 = base.F64_convert_i32_s(v492)
	v495 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_generate_useful_gather_paths[0])))
	if v495 == int32(1) {
		goto L143
	} else {
		goto L144
	}
L134:
	;
	v478 = v465
	goto L136
L135:
	;
	v478 = int32(0)
	goto L136
L136:
	;
	if v478 == int32(0) {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v482 = F_create_sort_path(m, l1, v381, v361, float64(-1))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L18
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	v485 = F_create_incremental_sort_path(m, l0, l1, v381, v361, v465, float64(-1))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L18
	} else {
		goto L141
	}
L140:
	;
	v487 = v482
	goto L133
L141:
	;
	v487 = v485
	goto L133
L142:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v17)+8)) = v523
	v525 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v487)+64))
	v527 = F_create_gather_merge_path(m, l0, l1, v487, v525, v526, v333)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L18
	} else {
		goto L152
	}
L143:
	;
	v501 = base.F64_add(base.F64_mul(v493, float64(-0.3)), float64(1))
	if base.F64_gt(v501, float64(0)) != 0 {
		goto L146
	} else {
		goto L147
	}
L144:
	;
	v507 = v493
	goto L145
L145:
	;
	v509 = float64(1e+100)
	v510 = base.F64_mul(v491, v507)
	if base.F64_gt(v510, v509)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v510)&int64(9223372036854775807))) != 0 {
		v523 = v509
		goto L149
	} else {
		goto L150
	}
L146:
	;
	v505 = v501
	goto L148
L147:
	;
	v505 = math.Float64frombits(uint64(0x8000000000000000))
	goto L148
L148:
	;
	v507 = base.F64_add(v505, v493)
	goto L145
L149:
	;
	goto L142
L150:
	;
	v519 = float64(1)
	if base.F64_le(v510, v519) != 0 {
		v523 = v519
		goto L149
	} else {
		goto L151
	}
L151:
	;
	v523 = base.F64_nearest(v510)
	goto L149
L152:
	;
	F_add_path(m, l1, v527)
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L18
	} else {
		goto L153
	}
L153:
	;
	goto L93
L154:
	;
	goto L92
L155:
	;
	goto L87
}
func F_geqo_randint(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v14 int64
	_ = v14
	var v21 int64
	_ = v21
	var v23 int64
	_ = v23
	var v24 int64
	_ = v24
	var v27 int64
	_ = v27
	var v29 int64
	_ = v29
	var v33 int64
	_ = v33
	var v35 int64
	_ = v35
	var v43 int64
	_ = v43
	var v48 int64
	_ = v48
	var v61 int64
	_ = v61
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+388))
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_geqo_randint[0]))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v4+v6<<(uint(int32(2))%32))))
	v12 = v10 + int32(8)
	v13 = base.I64_extend_i32_s(l2)
	v14 = base.I64_extend_i32_s(l1)
	if base.Ui64(v14) <= base.Ui64(v13) {
		v61 = v13
	} else {
		v21 = v14 - v13
		v23 = *(*int64)(unsafe.Add(mBase, uint32(v12)+8))
		v24 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
		v27 = v24
		v29 = v23
		for {
			v33 = v27 ^ v29
			v35 = base.I64_rotl(v33, int64(37))
			v43 = v33 ^ (v33<<(uint(int64(16))%64) ^ base.I64_rotl(v27, int64(24)))
			v48 = int64(base.Ui64(base.I64_rotl(v27*int64(5), int64(7))*int64(9)) >> (uint(base.I64_clz(v21)) % 64))
			if base.Ui64(v21) < base.Ui64(v48) {
				v27 = v43
				v29 = v35
				continue
			} else {
				break
			}
			break
		}
		*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v35
		*(*int64)(unsafe.Add(mBase, uint32(v12))) = v43
		v61 = v13 + v48
	}
	return base.I32_wrap_i64(v61)
}
func F_german_UTF_8_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
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
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v557 int32
	_ = v557
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v597 int32
	_ = v597
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v619 int32
	_ = v619
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v644 int32
	_ = v644
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v657 int32
	_ = v657
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v666 int32
	_ = v666
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v682 int32
	_ = v682
	var v695 int32
	_ = v695
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v715 int32
	_ = v715
	var v721 int32
	_ = v721
	var v728 int32
	_ = v728
	var v739 int32
	_ = v739
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v766 int32
	_ = v766
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v779 int32
	_ = v779
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v788 int32
	_ = v788
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v804 int32
	_ = v804
	var v817 int32
	_ = v817
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v837 int32
	_ = v837
	var v843 int32
	_ = v843
	var v851 int32
	_ = v851
	var v862 int32
	_ = v862
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v869 int32
	_ = v869
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v892 int32
	_ = v892
	var v899 int32
	_ = v899
	var v901 int32
	_ = v901
	var v905 int32
	_ = v905
	var v908 int32
	_ = v908
	var v910 int32
	_ = v910
	var v914 int32
	_ = v914
	var v924 int32
	_ = v924
	var v926 int32
	_ = v926
	var v930 int32
	_ = v930
	var v943 int32
	_ = v943
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v963 int32
	_ = v963
	var v969 int32
	_ = v969
	var v976 int32
	_ = v976
	var v987 int32
	_ = v987
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1014 int32
	_ = v1014
	var v1021 int32
	_ = v1021
	var v1023 int32
	_ = v1023
	var v1027 int32
	_ = v1027
	var v1030 int32
	_ = v1030
	var v1032 int32
	_ = v1032
	var v1036 int32
	_ = v1036
	var v1046 int32
	_ = v1046
	var v1048 int32
	_ = v1048
	var v1052 int32
	_ = v1052
	var v1065 int32
	_ = v1065
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1085 int32
	_ = v1085
	var v1091 int32
	_ = v1091
	var v1099 int32
	_ = v1099
	var v1110 int32
	_ = v1110
	var v1113 int32
	_ = v1113
	var v1119 int32
	_ = v1119
	var v1123 int32
	_ = v1123
	var v1125 int32
	_ = v1125
	var v1127 int32
	_ = v1127
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1146 int32
	_ = v1146
	var v1148 int32
	_ = v1148
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1155 int32
	_ = v1155
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1161 int32
	_ = v1161
	var v1164 int32
	_ = v1164
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1173 int32
	_ = v1173
	var v1176 int32
	_ = v1176
	var v1179 int32
	_ = v1179
	var v1182 int32
	_ = v1182
	var v1184 int32
	_ = v1184
	var v1186 int32
	_ = v1186
	var v1190 int32
	_ = v1190
	var v1194 int32
	_ = v1194
	var v1197 int32
	_ = v1197
	var v1199 int32
	_ = v1199
	var v1202 int32
	_ = v1202
	var v1205 int32
	_ = v1205
	var v1208 int32
	_ = v1208
	var v1212 int32
	_ = v1212
	var v1215 int32
	_ = v1215
	var v1230 int32
	_ = v1230
	var v1231 int32
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1251 int32
	_ = v1251
	var v1253 int32
	_ = v1253
	var v1260 int32
	_ = v1260
	var v1262 int32
	_ = v1262
	var v1264 int32
	_ = v1264
	var v1266 int32
	_ = v1266
	var v1279 int32
	_ = v1279
	var v1281 int32
	_ = v1281
	var v1283 int32
	_ = v1283
	var v1301 int32
	_ = v1301
	var v1303 int32
	_ = v1303
	var v1311 int32
	_ = v1311
	var v1315 int32
	_ = v1315
	var v1317 int32
	_ = v1317
	var v1323 int32
	_ = v1323
	var v1339 int32
	_ = v1339
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1352 int32
	_ = v1352
	var v1353 int32
	_ = v1353
	var v1358 int32
	_ = v1358
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1365 int32
	_ = v1365
	var v1367 int32
	_ = v1367
	var v1382 int32
	_ = v1382
	var v1383 int32
	_ = v1383
	var v1386 int32
	_ = v1386
	var v1388 int32
	_ = v1388
	var v1392 int32
	_ = v1392
	var v1407 int32
	_ = v1407
	var v1408 int32
	_ = v1408
	var v1409 int32
	_ = v1409
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1428 int32
	_ = v1428
	var v1430 int32
	_ = v1430
	var v1437 int32
	_ = v1437
	var v1439 int32
	_ = v1439
	var v1441 int32
	_ = v1441
	var v1443 int32
	_ = v1443
	var v1456 int32
	_ = v1456
	var v1458 int32
	_ = v1458
	var v1460 int32
	_ = v1460
	var v1478 int32
	_ = v1478
	var v1480 int32
	_ = v1480
	var v1488 int32
	_ = v1488
	var v1492 int32
	_ = v1492
	var v1494 int32
	_ = v1494
	var v1500 int32
	_ = v1500
	var v1516 int32
	_ = v1516
	var v1523 int32
	_ = v1523
	var v1524 int32
	_ = v1524
	var v1525 int32
	_ = v1525
	var v1526 int32
	_ = v1526
	var v1533 int32
	_ = v1533
	var v1535 int32
	_ = v1535
	var v1540 int32
	_ = v1540
	var v1542 int32
	_ = v1542
	var v1548 int32
	_ = v1548
	var v1553 int32
	_ = v1553
	var v1557 int32
	_ = v1557
	var v1560 int32
	_ = v1560
	var v1564 int32
	_ = v1564
	var v1578 int32
	_ = v1578
	var v1582 int32
	_ = v1582
	var v1585 int32
	_ = v1585
	var v1598 int32
	_ = v1598
	var v1599 int32
	_ = v1599
	var v1600 int32
	_ = v1600
	var v1616 int32
	_ = v1616
	var v1617 int32
	_ = v1617
	var v1619 int32
	_ = v1619
	var v1621 int32
	_ = v1621
	var v1628 int32
	_ = v1628
	var v1630 int32
	_ = v1630
	var v1632 int32
	_ = v1632
	var v1634 int32
	_ = v1634
	var v1647 int32
	_ = v1647
	var v1649 int32
	_ = v1649
	var v1651 int32
	_ = v1651
	var v1669 int32
	_ = v1669
	var v1671 int32
	_ = v1671
	var v1679 int32
	_ = v1679
	var v1683 int32
	_ = v1683
	var v1685 int32
	_ = v1685
	var v1691 int32
	_ = v1691
	var v1707 int32
	_ = v1707
	var v1714 int32
	_ = v1714
	var v1715 int32
	_ = v1715
	var v1716 int32
	_ = v1716
	var v1717 int32
	_ = v1717
	var v1720 int32
	_ = v1720
	var v1721 int32
	_ = v1721
	var v1723 int32
	_ = v1723
	var v1725 int32
	_ = v1725
	var v1740 int32
	_ = v1740
	var v1741 int32
	_ = v1741
	var v1742 int32
	_ = v1742
	var v1744 int32
	_ = v1744
	var v1747 int32
	_ = v1747
	var v1753 int32
	_ = v1753
	var v1757 int32
	_ = v1757
	var v1758 int32
	_ = v1758
	var v1760 int32
	_ = v1760
	var v1762 int32
	_ = v1762
	var v1777 int32
	_ = v1777
	var v1778 int32
	_ = v1778
	var v1781 int32
	_ = v1781
	var v1783 int32
	_ = v1783
	var v1787 int32
	_ = v1787
	var v1790 int32
	_ = v1790
	var v1792 int32
	_ = v1792
	var v1794 int32
	_ = v1794
	var v1797 int32
	_ = v1797
	var v1800 int32
	_ = v1800
	var v1803 int32
	_ = v1803
	var v1807 int32
	_ = v1807
	var v1810 int32
	_ = v1810
	var v1812 int32
	_ = v1812
	var v1814 int32
	_ = v1814
	var v1818 int32
	_ = v1818
	var v1821 int32
	_ = v1821
	var v1823 int32
	_ = v1823
	var v1826 int32
	_ = v1826
	var v1828 int32
	_ = v1828
	var v1832 int32
	_ = v1832
	var v1835 int32
	_ = v1835
	var v1838 int32
	_ = v1838
	var v1841 int32
	_ = v1841
	var v1843 int32
	_ = v1843
	var v1844 int32
	_ = v1844
	var v1846 int32
	_ = v1846
	var v1849 int32
	_ = v1849
	var v1852 int32
	_ = v1852
	var v1855 int32
	_ = v1855
	var v1859 int32
	_ = v1859
	var v1862 int32
	_ = v1862
	var v1864 int32
	_ = v1864
	var v1866 int32
	_ = v1866
	var v1868 int32
	_ = v1868
	var v1871 int32
	_ = v1871
	var v1874 int32
	_ = v1874
	var v1877 int32
	_ = v1877
	var v1881 int32
	_ = v1881
	var v1884 int32
	_ = v1884
	var v1886 int32
	_ = v1886
	var v1888 int32
	_ = v1888
	var v1891 int32
	_ = v1891
	var v1894 int32
	_ = v1894
	var v1897 int32
	_ = v1897
	var v1898 int32
	_ = v1898
	var v1900 int32
	_ = v1900
	var v1902 int32
	_ = v1902
	var v1909 int32
	_ = v1909
	var v1913 int32
	_ = v1913
	var v1914 int32
	_ = v1914
	var v1917 int32
	_ = v1917
	var v1919 int32
	_ = v1919
	var v1921 int32
	_ = v1921
	var v1925 int32
	_ = v1925
	var v1931 int32
	_ = v1931
	var v1934 int32
	_ = v1934
	var v1942 int32
	_ = v1942
	var v1943 int32
	_ = v1943
	var v1944 int32
	_ = v1944
	var v1950 int32
	_ = v1950
	var v1951 int32
	_ = v1951
	var v1956 int32
	_ = v1956
	var v1957 int32
	_ = v1957
	var v1962 int32
	_ = v1962
	var v1963 int32
	_ = v1963
	var v1968 int32
	_ = v1968
	var v1969 int32
	_ = v1969
	var v1972 int32
	_ = v1972
	var v1973 int32
	_ = v1973
	var v1980 int32
	_ = v1980
	var v1982 int32
	_ = v1982
	var v1987 int32
	_ = v1987
	var v1989 int32
	_ = v1989
	var v1996 int32
	_ = v1996
	var v1999 int32
	_ = v1999
	var v2003 int32
	_ = v2003
	var v2010 int32
	_ = v2010
	var v2011 int32
	_ = v2011
	var v2025 int32
	_ = v2025
	var v2030 int32
	_ = v2030
	var v2035 int32
	_ = v2035
	var v2036 int32
	_ = v2036
	var v2040 int32
	_ = v2040
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v11 = v6
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v11
	goto L1
L28:
	;
	return v2040
L29:
	;
	v2035 = F_slice_from_s(m, l0, int32(1), int32(_a_F_german_UTF_8_stem_2))
	mBase = m.M
	v2036 = m.ExcPending
	if v2036 != 0 {
		goto L112
	} else {
		goto L510
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v11
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
	v407 = v11
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
	v11 = v452
	goto L1
L112:
	;
	return int32(0)
L113:
	;
	if v458 < int32(0) {
		v2040 = v458
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
	v474 = F_find_among(m, l0, int32(_a_F_german_UTF_8_stem_4), int32(6), int32(0))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L112
	} else {
		goto L118
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6
	v564 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v564
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v564
	v567 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L157
L117:
	;
	goto L116
L118:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v476
	switch v474 - int32(1) {
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
	v562 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v466 = v562
	goto L115
L120:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v505 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L135
L121:
	;
	v500 = F_slice_from_s(m, l0, int32(2), int32(_a_F_german_UTF_8_stem_5))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L112
	} else {
		goto L131
	}
L122:
	;
	v494 = F_slice_from_s(m, l0, int32(2), int32(_a_F_german_UTF_8_stem_6))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L112
	} else {
		goto L129
	}
L123:
	;
	v488 = F_slice_from_s(m, l0, int32(2), int32(_a_F_german_UTF_8_stem_7))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L112
	} else {
		goto L127
	}
L124:
	;
	v482 = F_slice_from_s(m, l0, int32(2), int32(_a_F_german_UTF_8_stem_8))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L112
	} else {
		goto L125
	}
L125:
	;
	if int32(0) <= v482 {
		goto L119
	} else {
		goto L126
	}
L126:
	;
	v2040 = v482
	goto L28
L127:
	;
	if int32(0) <= v488 {
		goto L119
	} else {
		goto L128
	}
L128:
	;
	v2040 = v488
	goto L28
L129:
	;
	if int32(0) <= v494 {
		goto L119
	} else {
		goto L130
	}
L130:
	;
	v2040 = v494
	goto L28
L131:
	;
	if int32(0) <= v500 {
		goto L119
	} else {
		goto L132
	}
L132:
	;
	v2040 = v500
	goto L28
L133:
	;
	if v557 < int32(0) {
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
	v512 = v476
	v514 = int32(1)
	goto L140
L139:
	;
	v557 = v542
	goto L133
L140:
	;
	if v505 <= v512 {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	goto L139
L142:
	;
	v557 = int32(-1)
	goto L133
L143:
	;
	goto L144
L144:
	;
	v519 = v512 + int32(1)
	v521 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v504+v512))))
	if base.Ui32(v521) < base.Ui32(int32(192)) {
		v542 = v519
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v543 = int32(1)
	if v543 < v514 {
		v512 = v542
		v514 = v514 - v543
		goto L140
	} else {
		goto L152
	}
L146:
	;
	if v505 <= v519 {
		v542 = v519
		goto L145
	} else {
		goto L147
	}
L147:
	;
	v528 = v519
	goto L148
L148:
	;
	v531 = int32(*(*int8)(unsafe.Add(mBase, uint32(v504+v528))))
	if int32(-65) < v531 {
		v542 = v528
		goto L145
	} else {
		goto L150
	}
L149:
	;
	v542 = v505
	goto L145
L150:
	;
	v535 = v528 + int32(1)
	if v535 != v505 {
		v528 = v535
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v557
	goto L119
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v6
	v1119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1119
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1119
	if v1119 <= v6 {
		goto L281
	} else {
		goto L282
	}
L155:
	;
	if v619 < int32(0) {
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
	v574 = v6
	v576 = int32(3)
	goto L162
L161:
	;
	v619 = v604
	goto L155
L162:
	;
	if v564 <= v574 {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	goto L161
L164:
	;
	v619 = int32(-1)
	goto L155
L165:
	;
	goto L166
L166:
	;
	v581 = v574 + int32(1)
	v583 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v567+v574))))
	if base.Ui32(v583) < base.Ui32(int32(192)) {
		v604 = v581
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v605 = int32(1)
	if v605 < v576 {
		v574 = v604
		v576 = v576 - v605
		goto L162
	} else {
		goto L174
	}
L168:
	;
	if v564 <= v581 {
		v604 = v581
		goto L167
	} else {
		goto L169
	}
L169:
	;
	v590 = v581
	goto L170
L170:
	;
	v593 = int32(*(*int8)(unsafe.Add(mBase, uint32(v567+v590))))
	if int32(-65) < v593 {
		v604 = v590
		goto L167
	} else {
		goto L172
	}
L171:
	;
	v604 = v564
	goto L167
L172:
	;
	v597 = v590 + int32(1)
	if v597 != v564 {
		v590 = v597
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6
	v635 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v636 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v644 = v6
	goto L178
L176:
	;
	if v739 < int32(0) {
		goto L154
	} else {
		goto L201
	}
L177:
	;
	v739 = v711
	goto L176
L178:
	;
	if v635 <= v644 {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v739 = int32(-1)
	goto L176
L181:
	;
	goto L182
L182:
	;
	v651 = int32(1)
	v653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v644+v636))))
	if base.Ui32(v653) < base.Ui32(int32(192)) {
		v710 = v653
		v711 = v651
		goto L183
	} else {
		goto L184
	}
L183:
	;
	if int32(252) < v710 {
		goto L196
	} else {
		goto L197
	}
L184:
	;
	v657 = v644 + int32(1)
	if v657 == v635 {
		v710 = v653
		v711 = v651
		goto L183
	} else {
		goto L185
	}
L185:
	;
	v660 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v657+v636))))
	v662 = v660 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v653) {
		goto L187
	} else {
		goto L188
	}
L186:
	;
	v676 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v666+v636))))
	v678 = v676 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v653) {
		goto L192
	} else {
		goto L193
	}
L187:
	;
	v666 = v644 + int32(2)
	if v666 != v635 {
		goto L186
	} else {
		goto L190
	}
L188:
	;
	goto L189
L189:
	;
	v710 = v653<<(uint(int32(6))%32)&int32(1984) | v662
	v711 = int32(2)
	goto L183
L190:
	;
	goto L189
L191:
	;
	v695 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v636+v682))))
	v710 = v695&int32(63) | (v653<<(uint(int32(18))%32)&int32(_a_F_german_UTF_8_stem_0) | v662<<(uint(int32(12))%32) | v678<<(uint(int32(6))%32))
	v711 = int32(4)
	goto L183
L192:
	;
	v682 = v644 + int32(3)
	if v682 != v635 {
		goto L191
	} else {
		goto L195
	}
L193:
	;
	goto L194
L194:
	;
	v710 = v653<<(uint(int32(12))%32)&int32(_a_F_german_UTF_8_stem_1) | v662<<(uint(int32(6))%32) | v678
	v711 = int32(3)
	goto L183
L195:
	;
	goto L194
L196:
	;
	v728 = v711 + v644
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v728
	v644 = v728
	goto L178
L197:
	;
	v715 = v710 - int32(97)
	if v715 < int32(0) {
		goto L196
	} else {
		goto L198
	}
L198:
	;
	v721 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v715)>>(uint(int32(3))%32)))+uint32(_c_F_german_UTF_8_stem[0]))))
	if int32(base.Ui32(v721)>>(uint(v715&int32(7))%32))&int32(1) != 0 {
		goto L177
	} else {
		goto L199
	}
L199:
	;
	goto L196
L201:
	;
	v742 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v743 = v742 + v739
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v743
	v757 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v758 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v766 = v743
	goto L204
L202:
	;
	if v862 < int32(0) {
		goto L154
	} else {
		goto L226
	}
L203:
	;
	v862 = v833
	goto L202
L204:
	;
	if v757 <= v766 {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	v862 = int32(-1)
	goto L202
L207:
	;
	goto L208
L208:
	;
	v773 = int32(1)
	v775 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v766+v758))))
	if base.Ui32(v775) < base.Ui32(int32(192)) {
		v832 = v775
		v833 = v773
		goto L209
	} else {
		goto L210
	}
L209:
	;
	if int32(252) < v832 {
		goto L203
	} else {
		goto L222
	}
L210:
	;
	v779 = v766 + int32(1)
	if v779 == v757 {
		v832 = v775
		v833 = v773
		goto L209
	} else {
		goto L211
	}
L211:
	;
	v782 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v779+v758))))
	v784 = v782 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v775) {
		goto L213
	} else {
		goto L214
	}
L212:
	;
	v798 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v788+v758))))
	v800 = v798 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v775) {
		goto L218
	} else {
		goto L219
	}
L213:
	;
	v788 = v766 + int32(2)
	if v788 != v757 {
		goto L212
	} else {
		goto L216
	}
L214:
	;
	goto L215
L215:
	;
	v832 = v775<<(uint(int32(6))%32)&int32(1984) | v784
	v833 = int32(2)
	goto L209
L216:
	;
	goto L215
L217:
	;
	v817 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v758+v804))))
	v832 = v817&int32(63) | (v775<<(uint(int32(18))%32)&int32(_a_F_german_UTF_8_stem_0) | v784<<(uint(int32(12))%32) | v800<<(uint(int32(6))%32))
	v833 = int32(4)
	goto L209
L218:
	;
	v804 = v766 + int32(3)
	if v804 != v757 {
		goto L217
	} else {
		goto L221
	}
L219:
	;
	goto L220
L220:
	;
	v832 = v775<<(uint(int32(12))%32)&int32(_a_F_german_UTF_8_stem_1) | v784<<(uint(int32(6))%32) | v800
	v833 = int32(3)
	goto L209
L221:
	;
	goto L220
L222:
	;
	v837 = v832 - int32(97)
	if v837 < int32(0) {
		goto L203
	} else {
		goto L223
	}
L223:
	;
	v843 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v837)>>(uint(int32(3))%32)))+uint32(_c_F_german_UTF_8_stem[0]))))
	if int32(base.Ui32(v843)>>(uint(v837&int32(7))%32))&int32(1) == int32(0) {
		goto L203
	} else {
		goto L224
	}
L224:
	;
	v851 = v833 + v766
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v851
	v766 = v851
	goto L204
L226:
	;
	v865 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v866 = v865 + v862
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v866
	if v619 < v866 {
		goto L227
	} else {
		goto L228
	}
L227:
	;
	v869 = v866
	goto L229
L228:
	;
	v869 = v619
	goto L229
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v869
	v882 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v883 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v884 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v892 = v882
	goto L232
L230:
	;
	if v987 < int32(0) {
		goto L154
	} else {
		goto L255
	}
L231:
	;
	v987 = v959
	goto L230
L232:
	;
	if v883 <= v892 {
		goto L234
	} else {
		goto L235
	}
L234:
	;
	v987 = int32(-1)
	goto L230
L235:
	;
	goto L236
L236:
	;
	v899 = int32(1)
	v901 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v892+v884))))
	if base.Ui32(v901) < base.Ui32(int32(192)) {
		v958 = v901
		v959 = v899
		goto L237
	} else {
		goto L238
	}
L237:
	;
	if int32(252) < v958 {
		goto L250
	} else {
		goto L251
	}
L238:
	;
	v905 = v892 + int32(1)
	if v905 == v883 {
		v958 = v901
		v959 = v899
		goto L237
	} else {
		goto L239
	}
L239:
	;
	v908 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v905+v884))))
	v910 = v908 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v901) {
		goto L241
	} else {
		goto L242
	}
L240:
	;
	v924 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v914+v884))))
	v926 = v924 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v901) {
		goto L246
	} else {
		goto L247
	}
L241:
	;
	v914 = v892 + int32(2)
	if v914 != v883 {
		goto L240
	} else {
		goto L244
	}
L242:
	;
	goto L243
L243:
	;
	v958 = v901<<(uint(int32(6))%32)&int32(1984) | v910
	v959 = int32(2)
	goto L237
L244:
	;
	goto L243
L245:
	;
	v943 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v884+v930))))
	v958 = v943&int32(63) | (v901<<(uint(int32(18))%32)&int32(_a_F_german_UTF_8_stem_0) | v910<<(uint(int32(12))%32) | v926<<(uint(int32(6))%32))
	v959 = int32(4)
	goto L237
L246:
	;
	v930 = v892 + int32(3)
	if v930 != v883 {
		goto L245
	} else {
		goto L249
	}
L247:
	;
	goto L248
L248:
	;
	v958 = v901<<(uint(int32(12))%32)&int32(_a_F_german_UTF_8_stem_1) | v910<<(uint(int32(6))%32) | v926
	v959 = int32(3)
	goto L237
L249:
	;
	goto L248
L250:
	;
	v976 = v959 + v892
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v976
	v892 = v976
	goto L232
L251:
	;
	v963 = v958 - int32(97)
	if v963 < int32(0) {
		goto L250
	} else {
		goto L252
	}
L252:
	;
	v969 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v963)>>(uint(int32(3))%32)))+uint32(_c_F_german_UTF_8_stem[0]))))
	if int32(base.Ui32(v969)>>(uint(v963&int32(7))%32))&int32(1) != 0 {
		goto L231
	} else {
		goto L253
	}
L253:
	;
	goto L250
L255:
	;
	v990 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v991 = v990 + v987
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v991
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1006 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1014 = v991
	goto L258
L256:
	;
	if v1110 < int32(0) {
		goto L154
	} else {
		goto L280
	}
L257:
	;
	v1110 = v1081
	goto L256
L258:
	;
	if v1005 <= v1014 {
		goto L260
	} else {
		goto L261
	}
L260:
	;
	v1110 = int32(-1)
	goto L256
L261:
	;
	goto L262
L262:
	;
	v1021 = int32(1)
	v1023 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1014+v1006))))
	if base.Ui32(v1023) < base.Ui32(int32(192)) {
		v1080 = v1023
		v1081 = v1021
		goto L263
	} else {
		goto L264
	}
L263:
	;
	if int32(252) < v1080 {
		goto L257
	} else {
		goto L276
	}
L264:
	;
	v1027 = v1014 + int32(1)
	if v1027 == v1005 {
		v1080 = v1023
		v1081 = v1021
		goto L263
	} else {
		goto L265
	}
L265:
	;
	v1030 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1027+v1006))))
	v1032 = v1030 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1023) {
		goto L267
	} else {
		goto L268
	}
L266:
	;
	v1046 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1036+v1006))))
	v1048 = v1046 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1023) {
		goto L272
	} else {
		goto L273
	}
L267:
	;
	v1036 = v1014 + int32(2)
	if v1036 != v1005 {
		goto L266
	} else {
		goto L270
	}
L268:
	;
	goto L269
L269:
	;
	v1080 = v1023<<(uint(int32(6))%32)&int32(1984) | v1032
	v1081 = int32(2)
	goto L263
L270:
	;
	goto L269
L271:
	;
	v1065 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1006+v1052))))
	v1080 = v1065&int32(63) | (v1023<<(uint(int32(18))%32)&int32(_a_F_german_UTF_8_stem_0) | v1032<<(uint(int32(12))%32) | v1048<<(uint(int32(6))%32))
	v1081 = int32(4)
	goto L263
L272:
	;
	v1052 = v1014 + int32(3)
	if v1052 != v1005 {
		goto L271
	} else {
		goto L275
	}
L273:
	;
	goto L274
L274:
	;
	v1080 = v1023<<(uint(int32(12))%32)&int32(_a_F_german_UTF_8_stem_1) | v1032<<(uint(int32(6))%32) | v1048
	v1081 = int32(3)
	goto L263
L275:
	;
	goto L274
L276:
	;
	v1085 = v1080 - int32(97)
	if v1085 < int32(0) {
		goto L257
	} else {
		goto L277
	}
L277:
	;
	v1091 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1085)>>(uint(int32(3))%32)))+uint32(_c_F_german_UTF_8_stem[0]))))
	if int32(base.Ui32(v1091)>>(uint(v1085&int32(7))%32))&int32(1) == int32(0) {
		goto L257
	} else {
		goto L278
	}
L278:
	;
	v1099 = v1081 + v1014
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1099
	v1014 = v1099
	goto L258
L280:
	;
	v1113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v1113 + v1110
	goto L154
L281:
	;
	v1358 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1358
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1358
	v1362 = v1358 - int32(1)
	v1363 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1362 <= v1363 {
		goto L335
	} else {
		goto L336
	}
L282:
	;
	v1123 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1125 = int32(1)
	v1127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1123+v1119-v1125))))
	if base.B2i32(v1127&int32(224) != int32(96))|base.B2i32(v1125<<(uint(v1127)%32)&int32(_a_F_german_UTF_8_stem_9) == int32(0)) != 0 {
		goto L281
	} else {
		goto L283
	}
L283:
	;
	v1142 = F_find_among_b(m, l0, int32(_a_F_german_UTF_8_stem_10), int32(11), int32(0))
	mBase = m.M
	v1143 = m.ExcPending
	if v1143 != 0 {
		goto L112
	} else {
		goto L284
	}
L284:
	;
	if v1142 == int32(0) {
		goto L281
	} else {
		goto L285
	}
L285:
	;
	v1146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1146
	v1148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1146 < v1148 {
		goto L281
	} else {
		goto L286
	}
L286:
	;
	switch v1142 - int32(1) {
	case 0:
		goto L291
	case 1:
		goto L290
	case 2:
		goto L289
	case 3:
		goto L288
	case 4:
		goto L287
	default:
		goto L281
	}
L287:
	;
	v1352 = F_slice_from_s(m, l0, int32(1), int32(_a_F_german_UTF_8_stem_11))
	mBase = m.M
	v1353 = m.ExcPending
	if v1353 != 0 {
		goto L112
	} else {
		goto L333
	}
L288:
	;
	v1230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1232 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L310
L289:
	;
	v1179 = F_slice_del(m, l0)
	mBase = m.M
	if v1179 < int32(0) {
		v2040 = v1179
		goto L28
	} else {
		goto L299
	}
L290:
	;
	v1176 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1176 {
		goto L281
	} else {
		goto L298
	}
L291:
	;
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1153 = int32(4)
	v1155 = int32(0)
	v1157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1157-v1158 < v1153 {
		v1168 = v1155
		goto L293
	} else {
		goto L294
	}
L292:
	;
	if v1168 != 0 {
		goto L281
	} else {
		goto L296
	}
L293:
	;
	goto L292
L294:
	;
	v1161 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1164 = F_memcmp(m, v1161+v1157-v1153, int32(_a_F_german_UTF_8_stem_12), v1153)
	mBase = m.M
	if v1164 != 0 {
		v1168 = v1155
		goto L293
	} else {
		goto L295
	}
L295:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1157 - v1153
	v1168 = int32(1)
	goto L293
L296:
	;
	v1169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1169 + (v1146 - v1152)
	v1173 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1173 {
		goto L281
	} else {
		goto L297
	}
L297:
	;
	v2040 = v1173
	goto L28
L298:
	;
	v2040 = v1176
	goto L28
L299:
	;
	v1182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1182
	v1184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1182 <= v1184 {
		goto L281
	} else {
		goto L300
	}
L300:
	;
	v1186 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1186+v1182-int32(1)))))
	if v1190 != int32(115) {
		goto L281
	} else {
		goto L301
	}
L301:
	;
	v1194 = v1182 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1194
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1194
	v1197 = int32(3)
	v1199 = int32(0)
	v1202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1194-v1202 < v1197 {
		v1212 = v1199
		goto L303
	} else {
		goto L304
	}
L302:
	;
	if v1212 == int32(0) {
		goto L281
	} else {
		goto L306
	}
L303:
	;
	goto L302
L304:
	;
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1208 = F_memcmp(m, v1205+v1194-v1197, int32(_a_F_german_UTF_8_stem_13), v1197)
	mBase = m.M
	if v1208 != 0 {
		v1212 = v1199
		goto L303
	} else {
		goto L305
	}
L305:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1194 - v1197
	v1212 = int32(1)
	goto L303
L306:
	;
	v1215 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1215 {
		goto L281
	} else {
		goto L307
	}
L307:
	;
	v2040 = v1215
	goto L28
L308:
	;
	if v1346 != 0 {
		goto L281
	} else {
		goto L331
	}
L309:
	;
	v1346 = v1339
	goto L308
L310:
	;
	if v1230 <= v1231 {
		v1339 = int32(-1)
		goto L309
	} else {
		goto L312
	}
L311:
	;
	v1339 = int32(0)
	goto L309
L312:
	;
	v1248 = int32(1)
	v1249 = v1230 - v1248
	v1251 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1232+v1249))))
	v1253 = v1251 & int32(255)
	if base.B2i32(v1249 == v1231)|base.B2i32(int32(0) <= v1251) != 0 {
		v1311 = v1253
		v1315 = v1248
		goto L313
	} else {
		goto L314
	}
L313:
	;
	if int32(116) < v1311 {
		goto L321
	} else {
		goto L322
	}
L314:
	;
	v1260 = v1253 & int32(63)
	v1262 = v1230 - int32(2)
	v1264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1232+v1262))))
	v1266 = v1264 << (uint(int32(6)) % 32)
	if base.B2i32(v1262 != v1231)&base.B2i32(base.Ui32(v1264) < base.Ui32(int32(192))) == int32(0) {
		goto L315
	} else {
		goto L316
	}
L315:
	;
	v1311 = v1266&int32(1984) | v1260
	v1315 = int32(2)
	goto L313
L316:
	;
	goto L317
L317:
	;
	v1279 = v1266&int32(4032) | v1260
	v1281 = v1230 - int32(3)
	v1283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1232+v1281))))
	if base.B2i32(v1281 != v1231)&base.B2i32(base.Ui32(v1283) < base.Ui32(int32(224))) == int32(0) {
		goto L318
	} else {
		goto L319
	}
L318:
	;
	v1311 = v1283<<(uint(int32(12))%32)&int32(_a_F_german_UTF_8_stem_1) | v1279
	v1315 = int32(3)
	goto L313
L319:
	;
	goto L320
L320:
	;
	v1301 = int32(4)
	v1303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1230+v1232-v1301))))
	v1311 = v1283<<(uint(int32(12))%32)&int32(_a_F_german_UTF_8_stem_14) | v1303&int32(7)<<(uint(int32(18))%32) | v1279
	v1315 = v1301
	goto L313
L321:
	;
	v1346 = v1315
	goto L308
L322:
	;
	goto L323
L323:
	;
	v1317 = v1311 - int32(98)
	if v1317 < int32(0) {
		goto L324
	} else {
		goto L325
	}
L324:
	;
	v1346 = v1315
	goto L308
L325:
	;
	goto L326
L326:
	;
	v1323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1317)>>(uint(int32(3))%32)))+uint32(_c_F_german_UTF_8_stem[1]))))
	if int32(base.Ui32(v1323)>>(uint(v1317&int32(7))%32))&int32(1) == int32(0) {
		goto L327
	} else {
		goto L328
	}
L327:
	;
	v1346 = v1315
	goto L308
L328:
	;
	goto L329
L329:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1230 - v1315
	goto L330
L330:
	;
	goto L311
L331:
	;
	v1347 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1347 {
		goto L281
	} else {
		goto L332
	}
L332:
	;
	v2040 = v1347
	goto L28
L333:
	;
	if v1352 < int32(0) {
		v2040 = v1352
		goto L28
	} else {
		goto L334
	}
L334:
	;
	goto L281
L335:
	;
	v1753 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1753
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1753
	v1757 = v1753 - int32(1)
	v1758 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1757 <= v1758 {
		goto L420
	} else {
		goto L421
	}
L336:
	;
	v1365 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1365+v1362))))
	if base.B2i32(v1367&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v1367)%32)&int32(_a_F_german_UTF_8_stem_15) == int32(0)) != 0 {
		goto L335
	} else {
		goto L337
	}
L337:
	;
	v1382 = F_find_among_b(m, l0, int32(_a_F_german_UTF_8_stem_16), int32(5), int32(0))
	mBase = m.M
	v1383 = m.ExcPending
	if v1383 != 0 {
		goto L112
	} else {
		goto L338
	}
L338:
	;
	if v1382 == int32(0) {
		goto L335
	} else {
		goto L339
	}
L339:
	;
	v1386 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1386
	v1388 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1386 < v1388 {
		goto L335
	} else {
		goto L340
	}
L340:
	;
	switch v1382 - int32(1) {
	case 0:
		goto L343
	case 1:
		goto L342
	case 2:
		goto L341
	default:
		goto L335
	}
L341:
	;
	v1585 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1598 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1599 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1600 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L392
L342:
	;
	v1407 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1408 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1409 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L347
L343:
	;
	v1392 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1392 {
		goto L335
	} else {
		goto L344
	}
L344:
	;
	v2040 = v1392
	goto L28
L345:
	;
	if v1523 != 0 {
		goto L335
	} else {
		goto L368
	}
L346:
	;
	v1523 = v1516
	goto L345
L347:
	;
	if v1407 <= v1408 {
		v1516 = int32(-1)
		goto L346
	} else {
		goto L349
	}
L348:
	;
	v1516 = int32(0)
	goto L346
L349:
	;
	v1425 = int32(1)
	v1426 = v1407 - v1425
	v1428 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1409+v1426))))
	v1430 = v1428 & int32(255)
	if base.B2i32(v1426 == v1408)|base.B2i32(int32(0) <= v1428) != 0 {
		v1488 = v1430
		v1492 = v1425
		goto L350
	} else {
		goto L351
	}
L350:
	;
	if int32(116) < v1488 {
		goto L358
	} else {
		goto L359
	}
L351:
	;
	v1437 = v1430 & int32(63)
	v1439 = v1407 - int32(2)
	v1441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1409+v1439))))
	v1443 = v1441 << (uint(int32(6)) % 32)
	if base.B2i32(v1439 != v1408)&base.B2i32(base.Ui32(v1441) < base.Ui32(int32(192))) == int32(0) {
		goto L352
	} else {
		goto L353
	}
L352:
	;
	v1488 = v1443&int32(1984) | v1437
	v1492 = int32(2)
	goto L350
L353:
	;
	goto L354
L354:
	;
	v1456 = v1443&int32(4032) | v1437
	v1458 = v1407 - int32(3)
	v1460 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1409+v1458))))
	if base.B2i32(v1458 != v1408)&base.B2i32(base.Ui32(v1460) < base.Ui32(int32(224))) == int32(0) {
		goto L355
	} else {
		goto L356
	}
L355:
	;
	v1488 = v1460<<(uint(int32(12))%32)&int32(_a_F_german_UTF_8_stem_1) | v1456
	v1492 = int32(3)
	goto L350
L356:
	;
	goto L357
L357:
	;
	v1478 = int32(4)
	v1480 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1407+v1409-v1478))))
	v1488 = v1460<<(uint(int32(12))%32)&int32(_a_F_german_UTF_8_stem_14) | v1480&int32(7)<<(uint(int32(18))%32) | v1456
	v1492 = v1478
	goto L350
L358:
	;
	v1523 = v1492
	goto L345
L359:
	;
	goto L360
L360:
	;
	v1494 = v1488 - int32(98)
	if v1494 < int32(0) {
		goto L361
	} else {
		goto L362
	}
L361:
	;
	v1523 = v1492
	goto L345
L362:
	;
	goto L363
L363:
	;
	v1500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1494)>>(uint(int32(3))%32)))+uint32(_c_F_german_UTF_8_stem[2]))))
	if int32(base.Ui32(v1500)>>(uint(v1494&int32(7))%32))&int32(1) == int32(0) {
		goto L364
	} else {
		goto L365
	}
L364:
	;
	v1523 = v1492
	goto L345
L365:
	;
	goto L366
L366:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1407 - v1492
	goto L367
L367:
	;
	goto L348
L368:
	;
	v1524 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1525 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1526 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L371
L369:
	;
	if v1578 < int32(0) {
		goto L335
	} else {
		goto L388
	}
L371:
	;
	goto L372
L372:
	;
	goto L373
L373:
	;
	v1533 = v1525
	v1535 = int32(3)
	goto L376
L375:
	;
	v1578 = v1560
	goto L369
L376:
	;
	if v1533 <= v1526 {
		goto L378
	} else {
		goto L379
	}
L377:
	;
	goto L375
L378:
	;
	v1578 = int32(-1)
	goto L369
L379:
	;
	goto L380
L380:
	;
	v1540 = v1533 - int32(1)
	v1542 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1524+v1540))))
	if base.B2i32(int32(0) <= v1542)|base.B2i32(v1540 <= v1526) != 0 {
		v1560 = v1540
		goto L381
	} else {
		goto L382
	}
L381:
	;
	v1564 = int32(1)
	if v1564 < v1535 {
		v1533 = v1560
		v1535 = v1535 - v1564
		goto L376
	} else {
		goto L387
	}
L382:
	;
	v1548 = v1540
	goto L383
L383:
	;
	v1553 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1524+v1548))))
	if base.Ui32(int32(191)) < base.Ui32(v1553) {
		v1560 = v1548
		goto L381
	} else {
		goto L385
	}
L384:
	;
	v1560 = v1526
	goto L381
L385:
	;
	v1557 = v1548 - int32(1)
	if v1526 < v1557 {
		v1548 = v1557
		goto L383
	} else {
		goto L386
	}
L386:
	;
	goto L384
L387:
	;
	goto L377
L388:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1578
	v1582 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1582 {
		goto L335
	} else {
		goto L389
	}
L389:
	;
	v2040 = v1582
	goto L28
L390:
	;
	if v1714 != 0 {
		goto L335
	} else {
		goto L413
	}
L391:
	;
	v1714 = v1707
	goto L390
L392:
	;
	if v1598 <= v1599 {
		v1707 = int32(-1)
		goto L391
	} else {
		goto L394
	}
L393:
	;
	v1707 = int32(0)
	goto L391
L394:
	;
	v1616 = int32(1)
	v1617 = v1598 - v1616
	v1619 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1600+v1617))))
	v1621 = v1619 & int32(255)
	if base.B2i32(v1617 == v1599)|base.B2i32(int32(0) <= v1619) != 0 {
		v1679 = v1621
		v1683 = v1616
		goto L395
	} else {
		goto L396
	}
L395:
	;
	if int32(228) < v1679 {
		goto L403
	} else {
		goto L404
	}
L396:
	;
	v1628 = v1621 & int32(63)
	v1630 = v1598 - int32(2)
	v1632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1600+v1630))))
	v1634 = v1632 << (uint(int32(6)) % 32)
	if base.B2i32(v1630 != v1599)&base.B2i32(base.Ui32(v1632) < base.Ui32(int32(192))) == int32(0) {
		goto L397
	} else {
		goto L398
	}
L397:
	;
	v1679 = v1634&int32(1984) | v1628
	v1683 = int32(2)
	goto L395
L398:
	;
	goto L399
L399:
	;
	v1647 = v1634&int32(4032) | v1628
	v1649 = v1598 - int32(3)
	v1651 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1600+v1649))))
	if base.B2i32(v1649 != v1599)&base.B2i32(base.Ui32(v1651) < base.Ui32(int32(224))) == int32(0) {
		goto L400
	} else {
		goto L401
	}
L400:
	;
	v1679 = v1651<<(uint(int32(12))%32)&int32(_a_F_german_UTF_8_stem_1) | v1647
	v1683 = int32(3)
	goto L395
L401:
	;
	goto L402
L402:
	;
	v1669 = int32(4)
	v1671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1598+v1600-v1669))))
	v1679 = v1651<<(uint(int32(12))%32)&int32(_a_F_german_UTF_8_stem_14) | v1671&int32(7)<<(uint(int32(18))%32) | v1647
	v1683 = v1669
	goto L395
L403:
	;
	v1714 = v1683
	goto L390
L404:
	;
	goto L405
L405:
	;
	v1685 = v1679 - int32(85)
	if v1685 < int32(0) {
		goto L406
	} else {
		goto L407
	}
L406:
	;
	v1714 = v1683
	goto L390
L407:
	;
	goto L408
L408:
	;
	v1691 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1685)>>(uint(int32(3))%32)))+uint32(_c_F_german_UTF_8_stem[3]))))
	if int32(base.Ui32(v1691)>>(uint(v1685&int32(7))%32))&int32(1) == int32(0) {
		goto L409
	} else {
		goto L410
	}
L409:
	;
	v1714 = v1683
	goto L390
L410:
	;
	goto L411
L411:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1598 - v1683
	goto L412
L412:
	;
	goto L393
L413:
	;
	v1715 = v1386 - v1585
	v1716 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1717 = v1715 + v1716
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1717
	v1720 = v1717 - int32(1)
	v1721 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1720 <= v1721 {
		v1744 = v1717
		goto L414
	} else {
		goto L415
	}
L414:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1744
	v1747 = F_slice_del(m, l0)
	mBase = m.M
	if v1747 < int32(0) {
		v2040 = v1747
		goto L28
	} else {
		goto L419
	}
L415:
	;
	v1723 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1725 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1723+v1720))))
	if base.B2i32(v1725&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v1725)%32)&int32(_a_F_german_UTF_8_stem_17) == int32(0)) != 0 {
		v1744 = v1717
		goto L414
	} else {
		goto L416
	}
L416:
	;
	v1740 = F_find_among_b(m, l0, int32(_a_F_german_UTF_8_stem_18), int32(5), int32(0))
	mBase = m.M
	v1741 = m.ExcPending
	if v1741 != 0 {
		goto L112
	} else {
		goto L417
	}
L417:
	;
	if v1740 != 0 {
		goto L335
	} else {
		goto L418
	}
L418:
	;
	v1742 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1744 = v1742 + v1715
	goto L414
L419:
	;
	goto L335
L420:
	;
	v1931 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1931
	v1934 = v1931
	goto L471
L421:
	;
	v1760 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1762 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1760+v1757))))
	if base.B2i32(v1762&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v1762)%32)&int32(_a_F_german_UTF_8_stem_19) == int32(0)) != 0 {
		goto L420
	} else {
		goto L422
	}
L422:
	;
	v1777 = F_find_among_b(m, l0, int32(_a_F_german_UTF_8_stem_20), int32(8), int32(0))
	mBase = m.M
	v1778 = m.ExcPending
	if v1778 != 0 {
		goto L112
	} else {
		goto L423
	}
L423:
	;
	if v1777 == int32(0) {
		goto L420
	} else {
		goto L424
	}
L424:
	;
	v1781 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1781
	v1783 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1781 < v1783 {
		goto L420
	} else {
		goto L425
	}
L425:
	;
	switch v1777 - int32(1) {
	case 0:
		goto L429
	case 1:
		goto L428
	case 2:
		goto L427
	case 3:
		goto L426
	default:
		goto L420
	}
L426:
	;
	v1891 = F_slice_del(m, l0)
	mBase = m.M
	if v1891 < int32(0) {
		v2040 = v1891
		goto L28
	} else {
		goto L462
	}
L427:
	;
	v1838 = F_slice_del(m, l0)
	mBase = m.M
	if v1838 < int32(0) {
		v2040 = v1838
		goto L28
	} else {
		goto L447
	}
L428:
	;
	v1826 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1826 < v1781 {
		goto L442
	} else {
		goto L443
	}
L429:
	;
	v1787 = F_slice_del(m, l0)
	mBase = m.M
	if v1787 < int32(0) {
		v2040 = v1787
		goto L28
	} else {
		goto L430
	}
L430:
	;
	v1790 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1790
	v1792 = int32(2)
	v1794 = int32(0)
	v1797 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1790-v1797 < v1792 {
		v1807 = v1794
		goto L432
	} else {
		goto L433
	}
L431:
	;
	if v1807 == int32(0) {
		goto L420
	} else {
		goto L435
	}
L432:
	;
	goto L431
L433:
	;
	v1800 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1803 = F_memcmp(m, v1800+v1790-v1792, int32(_a_F_german_UTF_8_stem_21), v1792)
	mBase = m.M
	if v1803 != 0 {
		v1807 = v1794
		goto L432
	} else {
		goto L434
	}
L434:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1790 - v1792
	v1807 = int32(1)
	goto L432
L435:
	;
	v1810 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1810
	v1812 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1812 < v1810 {
		goto L436
	} else {
		goto L437
	}
L436:
	;
	v1814 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1818 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1814+v1810-int32(1)))))
	if v1818 == int32(101) {
		goto L420
	} else {
		goto L439
	}
L437:
	;
	goto L438
L438:
	;
	v1821 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1810 < v1821 {
		goto L420
	} else {
		goto L440
	}
L439:
	;
	goto L438
L440:
	;
	v1823 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1823 {
		goto L420
	} else {
		goto L441
	}
L441:
	;
	v2040 = v1823
	goto L28
L442:
	;
	v1828 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1832 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1828+v1781-int32(1)))))
	if v1832 == int32(101) {
		goto L420
	} else {
		goto L445
	}
L443:
	;
	goto L444
L444:
	;
	v1835 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1835 {
		goto L420
	} else {
		goto L446
	}
L445:
	;
	goto L444
L446:
	;
	v2040 = v1835
	goto L28
L447:
	;
	v1841 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1841
	v1843 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1844 = int32(2)
	v1846 = int32(0)
	v1849 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1841-v1849 < v1844 {
		v1859 = v1846
		goto L449
	} else {
		goto L450
	}
L448:
	;
	if v1859 == int32(0) {
		goto L452
	} else {
		goto L453
	}
L449:
	;
	goto L448
L450:
	;
	v1852 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1855 = F_memcmp(m, v1852+v1841-v1844, int32(_a_F_german_UTF_8_stem_22), v1844)
	mBase = m.M
	if v1855 != 0 {
		v1859 = v1846
		goto L449
	} else {
		goto L451
	}
L451:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1841 - v1844
	v1859 = int32(1)
	goto L449
L452:
	;
	v1862 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1864 = v1862 + (v1841 - v1843)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1864
	v1866 = int32(2)
	v1868 = int32(0)
	v1871 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1864-v1871 < v1866 {
		v1881 = v1868
		goto L456
	} else {
		goto L457
	}
L453:
	;
	goto L454
L454:
	;
	v1884 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1884
	v1886 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1884 < v1886 {
		goto L420
	} else {
		goto L460
	}
L455:
	;
	if v1881 == int32(0) {
		goto L420
	} else {
		goto L459
	}
L456:
	;
	goto L455
L457:
	;
	v1874 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1877 = F_memcmp(m, v1874+v1864-v1866, int32(_a_F_german_UTF_8_stem_23), v1866)
	mBase = m.M
	if v1877 != 0 {
		v1881 = v1868
		goto L456
	} else {
		goto L458
	}
L458:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1864 - v1866
	v1881 = int32(1)
	goto L456
L459:
	;
	goto L454
L460:
	;
	v1888 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1888 {
		goto L420
	} else {
		goto L461
	}
L461:
	;
	v2040 = v1888
	goto L28
L462:
	;
	v1894 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1894
	v1897 = v1894 - int32(1)
	v1898 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1897 <= v1898 {
		goto L420
	} else {
		goto L463
	}
L463:
	;
	v1900 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1902 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1900+v1897))))
	if base.Ui32(int32(1)) < base.Ui32((v1902-int32(103))&int32(255)) {
		goto L420
	} else {
		goto L464
	}
L464:
	;
	v1909 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1913 = F_find_among_b(m, l0, int32(_a_F_german_UTF_8_stem_24), int32(2), int32(0))
	mBase = m.M
	v1914 = m.ExcPending
	if v1914 != 0 {
		goto L112
	} else {
		goto L465
	}
L465:
	;
	if v1913 == int32(0) {
		goto L420
	} else {
		goto L466
	}
L466:
	;
	v1917 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1917
	v1919 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1917 < v1919 {
		goto L467
	} else {
		goto L468
	}
L467:
	;
	v1921 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1921 + (v1894 - v1909)
	goto L420
L468:
	;
	goto L469
L469:
	;
	v1925 = F_slice_del(m, l0)
	mBase = m.M
	if v1925 < int32(0) {
		v2040 = v1925
		goto L28
	} else {
		goto L470
	}
L470:
	;
	goto L420
L471:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1934
	v1942 = F_find_among(m, l0, int32(_a_F_german_UTF_8_stem_25), int32(6), int32(0))
	mBase = m.M
	v1943 = m.ExcPending
	if v1943 != 0 {
		goto L112
	} else {
		goto L474
	}
L472:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1931
	v2040 = int32(1)
	goto L28
L473:
	;
	goto L472
L474:
	;
	v1944 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1944
	switch v1942 - int32(1) {
	case 0:
		goto L480
	case 1:
		goto L479
	case 2:
		goto L478
	case 3:
		goto L477
	case 4:
		goto L476
	default:
		goto L475
	}
L475:
	;
	v2030 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1934 = v2030
	goto L471
L476:
	;
	v1972 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1973 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L491
L477:
	;
	v1968 = F_slice_from_s(m, l0, int32(1), int32(_a_F_german_UTF_8_stem_26))
	mBase = m.M
	v1969 = m.ExcPending
	if v1969 != 0 {
		goto L112
	} else {
		goto L487
	}
L478:
	;
	v1962 = F_slice_from_s(m, l0, int32(1), int32(_a_F_german_UTF_8_stem_27))
	mBase = m.M
	v1963 = m.ExcPending
	if v1963 != 0 {
		goto L112
	} else {
		goto L485
	}
L479:
	;
	v1956 = F_slice_from_s(m, l0, int32(1), int32(_a_F_german_UTF_8_stem_28))
	mBase = m.M
	v1957 = m.ExcPending
	if v1957 != 0 {
		goto L112
	} else {
		goto L483
	}
L480:
	;
	v1950 = F_slice_from_s(m, l0, int32(1), int32(_a_F_german_UTF_8_stem_29))
	mBase = m.M
	v1951 = m.ExcPending
	if v1951 != 0 {
		goto L112
	} else {
		goto L481
	}
L481:
	;
	if int32(0) <= v1950 {
		goto L475
	} else {
		goto L482
	}
L482:
	;
	v2040 = v1950
	goto L28
L483:
	;
	if int32(0) <= v1956 {
		goto L475
	} else {
		goto L484
	}
L484:
	;
	v2040 = v1956
	goto L28
L485:
	;
	if int32(0) <= v1962 {
		goto L475
	} else {
		goto L486
	}
L486:
	;
	v2040 = v1962
	goto L28
L487:
	;
	if int32(0) <= v1968 {
		goto L475
	} else {
		goto L488
	}
L488:
	;
	v2040 = v1968
	goto L28
L489:
	;
	if v2025 < int32(0) {
		goto L473
	} else {
		goto L509
	}
L491:
	;
	goto L492
L492:
	;
	goto L493
L493:
	;
	v1980 = v1944
	v1982 = int32(1)
	goto L496
L495:
	;
	v2025 = v2010
	goto L489
L496:
	;
	if v1973 <= v1980 {
		goto L498
	} else {
		goto L499
	}
L497:
	;
	goto L495
L498:
	;
	v2025 = int32(-1)
	goto L489
L499:
	;
	goto L500
L500:
	;
	v1987 = v1980 + int32(1)
	v1989 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1972+v1980))))
	if base.Ui32(v1989) < base.Ui32(int32(192)) {
		v2010 = v1987
		goto L501
	} else {
		goto L502
	}
L501:
	;
	v2011 = int32(1)
	if v2011 < v1982 {
		v1980 = v2010
		v1982 = v1982 - v2011
		goto L496
	} else {
		goto L508
	}
L502:
	;
	if v1973 <= v1987 {
		v2010 = v1987
		goto L501
	} else {
		goto L503
	}
L503:
	;
	v1996 = v1987
	goto L504
L504:
	;
	v1999 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1972+v1996))))
	if int32(-65) < v1999 {
		v2010 = v1996
		goto L501
	} else {
		goto L506
	}
L505:
	;
	v2010 = v1973
	goto L501
L506:
	;
	v2003 = v1996 + int32(1)
	if v2003 != v1973 {
		v1996 = v2003
		goto L504
	} else {
		goto L507
	}
L507:
	;
	goto L505
L508:
	;
	goto L497
L509:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2025
	goto L475
L510:
	;
	if int32(0) <= v2035 {
		goto L27
	} else {
		goto L511
	}
L511:
	;
	v2040 = v2035
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
								F_errfinish(m, int32(_a_F_getIdentitySequence_1), int32(1167), int32(_a_F_getIdentitySequence_2))
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
												F_errfinish(m, int32(_a_F_getIdentitySequence_1), int32(1173), int32(_a_F_getIdentitySequence_2))
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
												F_errfinish(m, int32(_a_F_getIdentitySequence_1), int32(1179), int32(_a_F_getIdentitySequence_2))
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
							F_errfinish(m, int32(_a_F_getIdentitySequence_1), int32(1173), int32(_a_F_getIdentitySequence_2))
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
							F_errfinish(m, int32(_a_F_getIdentitySequence_1), int32(1179), int32(_a_F_getIdentitySequence_2))
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
	var v32 int32
	_ = v32
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
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
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	v8 = m.G0
	v10 = v8 - int32(176)
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
	F_ScanKeyInit(m, v10, int32(4), int32(3), int32(184), int64(1259))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_ScanKeyInit(m, v10+int32(56), int32(5), int32(3), int32(184), base.I64_extend_i32_u(l0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
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
	v45 = int32(0)
	v49 = F_systable_beginscan(m, v15, int32(2674), int32(1), v45, v44, v10)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L10
	}
L6:
	;
	v44 = int32(2)
	goto L5
L7:
	;
	goto L8
L8:
	;
	F_ScanKeyInit(m, v10+int32(112), int32(6), int32(3), int32(65), base.I64_extend_i32_s(l1))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v44 = int32(3)
	goto L5
L10:
	;
	v51 = F_systable_getnext(m, v49)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	if v51 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v55 = v45
	v56 = v51
	goto L15
L13:
	;
	v88 = v45
	goto L14
L14:
	;
	F_systable_endscan(m, v49)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L31
	}
L15:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v56)+16))
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+22)))
	v64 = v62 + v63
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	if v65 != int32(1259) {
		v85 = v55
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v88 = v85
	goto L14
L17:
	;
	v86 = F_systable_getnext(m, v49)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L29
	}
L18:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	if v68 != 0 {
		v85 = v55
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v64)+20))
	if v69 == int32(0) {
		v85 = v55
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+24)))
	switch v72 - int32(97) {
	case 0, 8:
		goto L21
	default:
		v85 = v55
		goto L17
	}
L21:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	v76 = F_get_rel_relkind(m, v75)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	if v76 != int32(83) {
		v85 = v55
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
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+24)))
	if v80 != l2&int32(255) {
		v85 = v55
		goto L17
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	v83 = F_lappend_oid(m, v55, v82)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L28
	}
L27:
	;
	goto L26
L28:
	;
	v85 = v83
	goto L17
L29:
	;
	if v86 != 0 {
		v55 = v85
		v56 = v86
		goto L15
	} else {
		goto L30
	}
L30:
	;
	goto L16
L31:
	;
	F_relation_close(m, v15, int32(1))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	m.G0 = v10 + int32(176)
	return v88
}
func F_get_am_name(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14286(m, l0, int32(2))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_get_am_type_oid(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v13 = F_SearchSysCache1(m, int32(1), base.I64_extend_i32_u(l0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		if v13 == int32(0) {
			v30 = int32(0)
			if v30|l2 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v88 = m.ExcPending
				if v88 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(67137668))
					mBase = m.M
					v91 = m.ExcPending
					if v91 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
						F_errmsg(m, int32(_a_F_get_am_type_oid_0), v9)
						mBase = m.M
						v95 = m.ExcPending
						if v95 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_get_am_type_oid_1), int32(154), int32(_a_F_get_am_type_oid_2))
							mBase = m.M
							v100 = m.ExcPending
							if v100 != 0 {
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
				return v30
			}
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+22)))
			v22 = v20 + v21
			if l1 != 0 {
				v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+72)))
				if v23 != l1&int32(255) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(325))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							v46 = m.G0
							v48 = v46 - int32(16)
							m.G0 = v48
							switch l1 - int32(105) {
							case 0:
								v67 = int32(_a_F_get_am_type_oid_3)
								v68 = int32(16)
								m.G0 = v48 + v68
								*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v67
								*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v22 + int32(4)
								F_errmsg(m, int32(_a_F_get_am_type_oid_4), v9+v68)
								mBase = m.M
								v79 = m.ExcPending
								if v79 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_get_am_type_oid_1), int32(145), int32(_a_F_get_am_type_oid_2))
									mBase = m.M
									v84 = m.ExcPending
									if v84 != 0 {
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
								v56 = m.ExcPending
								if v56 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v48))) = l1
									F_errmsg_internal(m, int32(_a_F_get_am_type_oid_5), v48)
									mBase = m.M
									v60 = m.ExcPending
									if v60 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_get_am_type_oid_1), int32(222), int32(_a_F_get_am_type_oid_6))
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							case 11:
								v67 = int32(_a_F_get_am_type_oid_7)
								v68 = int32(16)
								m.G0 = v48 + v68
								*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v67
								*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v22 + int32(4)
								F_errmsg(m, int32(_a_F_get_am_type_oid_4), v9+v68)
								mBase = m.M
								v79 = m.ExcPending
								if v79 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_get_am_type_oid_1), int32(145), int32(_a_F_get_am_type_oid_2))
									mBase = m.M
									v84 = m.ExcPending
									if v84 != 0 {
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
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
					F_ReleaseCatCache(m, v13)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						v30 = v27
						if v30|l2 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v88 = m.ExcPending
							if v88 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(67137668))
								mBase = m.M
								v91 = m.ExcPending
								if v91 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
									F_errmsg(m, int32(_a_F_get_am_type_oid_0), v9)
									mBase = m.M
									v95 = m.ExcPending
									if v95 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_get_am_type_oid_1), int32(154), int32(_a_F_get_am_type_oid_2))
										mBase = m.M
										v100 = m.ExcPending
										if v100 != 0 {
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
							return v30
						}
					}
				}
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
				F_ReleaseCatCache(m, v13)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					v30 = v27
					if v30|l2 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v88 = m.ExcPending
						if v88 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(67137668))
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
								F_errmsg(m, int32(_a_F_get_am_type_oid_0), v9)
								mBase = m.M
								v95 = m.ExcPending
								if v95 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_get_am_type_oid_1), int32(154), int32(_a_F_get_am_type_oid_2))
									mBase = m.M
									v100 = m.ExcPending
									if v100 != 0 {
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
						return v30
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
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v12 = F_SearchSysCache2(m, int32(7), base.I64_extend_i32_u(l0), base.I64_extend_i32_s(l1))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if v12 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l0
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
				F_errmsg_internal(m, int32(_a_F_get_attgenerated_0), v7)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_get_attgenerated_1), int32(1123), int32(_a_F_get_attgenerated_2))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
			v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+22)))
			v35 = int32(*(*int8)(unsafe.Add(mBase, uint32(v32+v33)+90)))
			F_ReleaseCatCache(m, v12)
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return int32(0)
			} else {
				m.G0 = v7 + int32(16)
				return v35
			}
		}
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
	var v47 int32
	_ = v47
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
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v105 int32
	_ = v105
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
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int64
	_ = v337
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
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
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v477 int32
	_ = v477
	var v489 int32
	_ = v489
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v568 int32
	_ = v568
	var v592 int32
	_ = v592
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v622 int32
	_ = v622
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
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v26 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L1
	} else {
		goto L134
	}
L4:
	;
	return v592
L5:
	;
	v322 = int32(0)
	if l7 == v322 {
		v592 = v322
		goto L4
	} else {
		goto L72
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
	v47 = int32(0)
	goto L8
L8:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v50+v47<<(uint(int32(2))%32))))
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
	v301 = v47 + int32(1)
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v301 < v302 {
		v47 = v301
		goto L8
	} else {
		goto L71
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
	v85 = v69
	v86 = int32(-1)
	v88 = v75
	goto L23
L23:
	;
	if v88 != 0 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v592 = v54
	goto L4
L25:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v193)))
	if v201 == int32(0) {
		goto L10
	} else {
		goto L45
	}
L26:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
	v190 = v85
	v191 = v86
	v193 = v88
	v200 = v95
	goto L25
L27:
	;
	goto L28
L28:
	;
	v105 = v86
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
	v190 = v178
	v191 = v169
	v193 = v181
	v200 = v181
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
	v120 = v105 + int32(1)
	v122 = int32(base.Ui32(v120) >> (uint(int32(5)) % 32))
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
		v105 = v169
		goto L29
	} else {
		goto L44
	}
L44:
	;
	goto L30
L45:
	;
	v205 = v193 + int32(4)
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v190)+4))
	if base.Ui32(v205) < base.Ui32(v200+v207<<(uint(int32(2))%32)) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v212 = v205
	goto L48
L47:
	;
	v212 = int32(0)
	goto L48
L48:
	;
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201)+13)))
	if v213 == int32(1) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v201)+8))
	v217 = int32(0)
	if base.B2i32(v216 == v217)|base.B2i32(l6 == v217) != 0 {
		v263 = base.B2i32(v216|l6 == v217)
		goto L53
	} else {
		goto L54
	}
L50:
	;
	goto L51
L51:
	;
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201)+12)))
	if v270 == int32(1) {
		goto L64
	} else {
		goto L65
	}
L52:
	;
	if v263 == int32(0) {
		v85 = v190
		v86 = v191
		v88 = v212
		goto L23
	} else {
		goto L63
	}
L53:
	;
	goto L52
L54:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v216)+4))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v231 != v232 {
		v263 = int32(0)
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v234 = int32(1)
	if v231 <= v234 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v237 = v234
	goto L58
L57:
	;
	v237 = v231
	goto L58
L58:
	;
	v238 = int32(8)
	v243 = int32(0)
	goto L59
L59:
	;
	v251 = v243 << (uint(int32(2)) % 32)
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v216+v238+v251)))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l6+v238+v251)))
	v256 = base.B2i32(v253 == v255)
	if v253 != v255 {
		v263 = v256
		goto L53
	} else {
		goto L61
	}
L60:
	;
	v263 = v256
	goto L53
L61:
	;
	v259 = v243 + int32(1)
	if v259 != v237 {
		v243 = v259
		goto L59
	} else {
		goto L62
	}
L62:
	;
	goto L60
L63:
	;
	goto L51
L64:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v201)+20))
	if v273 != v25 {
		v85 = v190
		v86 = v191
		v88 = v212
		goto L23
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v201)+16))
	if l3 != v275 {
		v85 = v190
		v86 = v191
		v88 = v212
		goto L23
	} else {
		goto L68
	}
L67:
	;
	goto L66
L68:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v201)+4))
	v278 = F_equal(m, v19, v277)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	if v278 == int32(0) {
		v85 = v190
		v86 = v191
		v88 = v212
		goto L23
	} else {
		goto L70
	}
L70:
	;
	goto L24
L71:
	;
	goto L9
L72:
	;
	v325 = int32(_a_F_get_eclass_for_sort_expr_0)
	v326 = *(*int32)(unsafe.Add(mBase, _c_F_get_eclass_for_sort_expr[0]))
	v328 = *(*int32)(unsafe.Add(mBase, uint32(l0)+300))
	*(*int32)(unsafe.Add(mBase, _c_F_get_eclass_for_sort_expr[0])) = v328
	v331 = F_palloc0(m, int32(60))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v331))) = int32(276)
	v335 = F_list_copy(m, l2)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	v337 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v331)+12)) = v337
	*(*int32)(unsafe.Add(mBase, uint32(v331)+8)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v331)+4)) = v335
	*(*int64)(unsafe.Add(mBase, uint32(v331)+20)) = v337
	*(*int64)(unsafe.Add(mBase, uint32(v331)+28)) = v337
	*(*int64)(unsafe.Add(mBase, uint32(v331)+33)) = v337
	v347 = F_contain_volatile_functions(m, v19)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	v349 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v331)+56)) = v349
	*(*int64)(unsafe.Add(mBase, uint32(v331)+48)) = int64(4294967295)
	*(*int32)(unsafe.Add(mBase, uint32(v331)+44)) = l5
	*(*uint8)(unsafe.Add(mBase, uint32(v331)+42)) = uint8(v349)
	*(*uint8)(unsafe.Add(mBase, uint32(v331)+41)) = uint8(v347)
	if v347 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v358 = l5
	goto L78
L77:
	;
	v358 = int32(1)
	goto L78
L78:
	;
	if v358 == int32(0) {
		goto L3
	} else {
		goto L79
	}
L79:
	;
	v361 = F_pull_varnos(m, l0, v19)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	v363 = F_copyObjectImpl(m, v19)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v366 = F_palloc0(m, int32(28))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	v368 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v366)+24)) = v368
	*(*int32)(unsafe.Add(mBase, uint32(v366)+20)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v366)+16)) = l3
	*(*uint16)(unsafe.Add(mBase, uint32(v366)+12)) = uint16(v368)
	*(*int32)(unsafe.Add(mBase, uint32(v366)+8)) = v361
	*(*int32)(unsafe.Add(mBase, uint32(v366)+4)) = v363
	*(*int32)(unsafe.Add(mBase, uint32(v366))) = int32(277)
	if v361 == v368 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v380 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v366)+12)) = uint8(v380)
	*(*uint8)(unsafe.Add(mBase, uint32(v331)+40)) = uint8(v380)
	goto L85
L84:
	;
	goto L85
L85:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v331)+16))
	v385 = F_lappend(m, v384, v366)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v331)+16)) = v385
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v331)+36))
	v389 = F_bms_add_members(m, v388, v361)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v331)+36)) = v389
	v392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331)+40)))
	if v392 != int32(1) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v409 = F_lappend(m, v408, v331)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L1
	} else {
		goto L98
	}
L89:
	;
	v395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331)+41)))
	if v395 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v404 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v331)+40)) = uint8(v404)
	*(*uint8)(unsafe.Add(mBase, uint32(v366)+12)) = uint8(v404)
	goto L88
L91:
	;
	v396 = F_expression_returns_set(m, v19)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	if v396 != 0 {
		goto L90
	} else {
		goto L93
	}
L93:
	;
	v398 = F_contain_agg_clause(m, v19)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	if v398 != 0 {
		goto L90
	} else {
		goto L95
	}
L95:
	;
	v400 = F_contain_windowfuncs(m, v19)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	if v400 == int32(0) {
		goto L88
	} else {
		goto L97
	}
L97:
	;
	goto L90
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v409
	v412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+100)))
	if v412 != int32(1) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_get_eclass_for_sort_expr[0])) = v326
	v592 = v331
	goto L4
L100:
	;
	if v409 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v409)+4))
	v419 = v416 - int32(1)
	goto L103
L102:
	;
	v419 = int32(-1)
	goto L103
L103:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v331)+36))
	if v420 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L104:
	;
	if v477 <= int32(0) {
		goto L99
	} else {
		goto L115
	}
L105:
	;
	v477 = base.I32_ctz(v463) | v464<<(uint(int32(5))%32)
	goto L104
L106:
	;
	v477 = int32(-2)
	goto L104
L107:
	;
	v428 = int32(0)
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v420)+4))
	if v431 <= v428 {
		goto L106
	} else {
		goto L108
	}
L108:
	;
	v434 = v420 + int32(8)
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v434)))
	v441 = v438 & int32(-1)
	if v441 != 0 {
		v463 = v441
		v464 = v428
		goto L105
	} else {
		goto L109
	}
L109:
	;
	v442 = int32(1)
	if v442 == v431 {
		goto L106
	} else {
		goto L110
	}
L110:
	;
	v446 = v442
	goto L111
L111:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v434+v446<<(uint(int32(2))%32))))
	if v453 != 0 {
		v463 = v453
		v464 = v446
		goto L105
	} else {
		goto L113
	}
L112:
	;
	goto L106
L113:
	;
	v455 = v446 + int32(1)
	if v455 != v431 {
		v446 = v455
		goto L111
	} else {
		goto L114
	}
L114:
	;
	goto L112
L115:
	;
	v489 = v477
	goto L116
L116:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(l0)+340))
	if v489 == v498 {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	goto L99
L118:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v331)+36))
	if v512 == int32(0) {
		goto L124
	} else {
		goto L125
	}
L119:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v500+v489<<(uint(int32(2))%32))))
	if v504 == int32(0) {
		goto L118
	} else {
		goto L120
	}
L120:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v504)+144))
	v508 = F_bms_add_member(m, v507, v419)
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v504)+144)) = v508
	goto L118
L122:
	;
	if int32(0) < v568 {
		v489 = v568
		goto L116
	} else {
		goto L133
	}
L123:
	;
	v568 = base.I32_ctz(v554) | v555<<(uint(int32(5))%32)
	goto L122
L124:
	;
	v568 = int32(-2)
	goto L122
L125:
	;
	v519 = v489 + int32(1)
	v521 = int32(base.Ui32(v519) >> (uint(int32(5)) % 32))
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v512)+4))
	if v522 <= v521 {
		goto L124
	} else {
		goto L126
	}
L126:
	;
	v525 = v512 + int32(8)
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v525+v521<<(uint(int32(2))%32))))
	v532 = v529 & (int32(-1) << (uint(v519) % 32))
	if v532 != 0 {
		v554 = v532
		v555 = v521
		goto L123
	} else {
		goto L127
	}
L127:
	;
	v534 = v521 + int32(1)
	if v534 == v522 {
		goto L124
	} else {
		goto L128
	}
L128:
	;
	v537 = v534
	goto L129
L129:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v525+v537<<(uint(int32(2))%32))))
	if v544 != 0 {
		v554 = v544
		v555 = v537
		goto L123
	} else {
		goto L131
	}
L130:
	;
	goto L124
L131:
	;
	v546 = v537 + int32(1)
	if v546 != v522 {
		v537 = v546
		goto L129
	} else {
		goto L132
	}
L132:
	;
	goto L130
L133:
	;
	goto L117
L134:
	;
	F_errmsg_internal(m, int32(_a_F_get_eclass_for_sort_expr_1), int32(0))
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	F_errfinish(m, int32(_a_F_get_eclass_for_sort_expr_2), int32(839), int32(_a_F_get_eclass_for_sort_expr_3))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
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
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
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
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
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
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
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
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v461 int32
	_ = v461
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v549 int32
	_ = v549
	v2 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v19 = F_strlen(m, v18)
	mBase = m.M
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v20 == v2 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v41 = F_AllocateDir(m, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L5
	} else {
		goto L12
	}
L2:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v24 = F_pstrdup(m, v23)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v28 == int32(47) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	return int32(0)
L6:
	;
	v40 = v24
	goto L1
L7:
	;
	v31 = F_pstrdup(m, v20)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L5
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v33
	v37 = F_psprintf(m, int32(_a_F_get_ext_ver_list_0), v16)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L5
	} else {
		goto L11
	}
L10:
	;
	v40 = v31
	goto L1
L11:
	;
	v40 = v37
	goto L1
L12:
	;
	v43 = F_ReadDir(m, v41, v40)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	if v43 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v46 = v43
	v47 = v2
	goto L17
L15:
	;
	v537 = v2
	goto L16
L16:
	;
	F_FreeDir(m, v41)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L5
	} else {
		goto L127
	}
L17:
	;
	v59 = v46 + int32(19)
	v63 = F_strlen(m, v59)
	mBase = m.M
	v70 = v63 + int32(1)
	goto L22
L18:
	;
	v537 = v522
	goto L16
L19:
	;
	v533 = F_ReadDir(m, v41, v40)
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L5
	} else {
		goto L125
	}
L20:
	;
	if v82 == int32(0) {
		v522 = v47
		goto L19
	} else {
		goto L26
	}
L21:
	;
	goto L20
L22:
	;
	v72 = int32(0)
	if v70 == v72 {
		v82 = v72
		goto L21
	} else {
		goto L24
	}
L23:
	;
	v82 = v77
	goto L21
L24:
	;
	v76 = v70 - int32(1)
	v77 = v59 + v76
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	if v78 != int32(46) {
		v70 = v76
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v85 = int32(_a_F_get_ext_ver_list_1)
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	v91 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_get_ext_ver_list[0])))
	if base.B2i32(v88 == int32(0))|base.B2i32(v88 != v91) != 0 {
		v109 = v88
		v110 = v91
		goto L28
	} else {
		goto L29
	}
L27:
	;
	if v109-v110 != 0 {
		v522 = v47
		goto L19
	} else {
		goto L34
	}
L28:
	;
	goto L27
L29:
	;
	v94 = v82
	v95 = v85
	goto L30
L30:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+1)))
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+1)))
	if v99 == int32(0) {
		v109 = v99
		v110 = v98
		goto L28
	} else {
		goto L32
	}
L31:
	;
	v109 = v99
	v110 = v98
	goto L28
L32:
	;
	v102 = int32(1)
	if v99 == v98 {
		v94 = v94 + v102
		v95 = v95 + v102
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v19 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	if v157 != 0 {
		v522 = v47
		goto L19
	} else {
		goto L48
	}
L36:
	;
	v157 = int32(0)
	goto L35
L37:
	;
	goto L38
L38:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	if v118 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v119 = v59
	v120 = v112
	v121 = v19
	v122 = v118
	goto L43
L40:
	;
	v145 = v112
	v149 = int32(0)
	goto L41
L41:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145))))
	v157 = v149 - v150
	goto L35
L42:
	;
	v145 = v140
	v149 = v142
	goto L41
L43:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120))))
	if base.B2i32(v122 != v124)|base.B2i32(v124 == int32(0)) != 0 {
		v140 = v120
		v142 = v122
		goto L42
	} else {
		goto L45
	}
L44:
	;
	v140 = v134
	v142 = int32(0)
	goto L42
L45:
	;
	v130 = v121 - int32(1)
	if v130 == int32(0) {
		v140 = v120
		v142 = v122
		goto L42
	} else {
		goto L46
	}
L46:
	;
	v133 = int32(1)
	v134 = v120 + v133
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+1)))
	if v135 != 0 {
		v119 = v119 + v133
		v120 = v134
		v121 = v130
		v122 = v135
		goto L43
	} else {
		goto L47
	}
L47:
	;
	goto L44
L48:
	;
	v158 = v59 + v19
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
	if v159 != int32(45) {
		v522 = v47
		goto L19
	} else {
		goto L49
	}
L49:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+1)))
	if v162 != int32(45) {
		v522 = v47
		goto L19
	} else {
		goto L50
	}
L50:
	;
	v167 = F_pstrdup(m, v158+int32(2))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L5
	} else {
		goto L51
	}
L51:
	;
	v172 = F_strlen(m, v167)
	mBase = m.M
	v179 = v172 + int32(1)
	goto L54
L52:
	;
	v192 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v191))) = uint8(v192)
	v195 = F_strstr(m, v167, int32(_a_F_get_ext_ver_list_2))
	mBase = m.M
	if v195 == v192 {
		goto L58
	} else {
		goto L59
	}
L53:
	;
	goto L52
L54:
	;
	v181 = int32(0)
	if v179 == v181 {
		v191 = v181
		goto L53
	} else {
		goto L56
	}
L55:
	;
	v191 = v186
	goto L53
L56:
	;
	v185 = v179 - int32(1)
	v186 = v167 + v185
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186))))
	if v187 != int32(46) {
		v179 = v185
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	if v47 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	goto L60
L60:
	;
	v296 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v195))) = uint8(v296)
	v299 = v195 + int32(2)
	v301 = F_strstr(m, v299, int32(_a_F_get_ext_ver_list_2))
	mBase = m.M
	if v301 != 0 {
		v522 = v47
		goto L19
	} else {
		goto L79
	}
L61:
	;
	v294 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v284)+8)) = uint8(v294)
	v522 = v283
	goto L19
L62:
	;
	v268 = F_palloc(m, int32(20))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L5
	} else {
		goto L76
	}
L63:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	if v200 <= int32(0) {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
	v206 = int32(0)
	goto L65
L65:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v203+v206<<(uint(int32(2))%32))))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)))
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222))))
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
	if base.B2i32(v225 == int32(0))|base.B2i32(v225 != v228) != 0 {
		v246 = v225
		v247 = v228
		goto L68
	} else {
		goto L69
	}
L66:
	;
	goto L62
L67:
	;
	if v246-v247 == int32(0) {
		v283 = v47
		v284 = v221
		goto L61
	} else {
		goto L74
	}
L68:
	;
	goto L67
L69:
	;
	v231 = v222
	v232 = v167
	goto L70
L70:
	;
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v232)+1)))
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+1)))
	if v236 == int32(0) {
		v246 = v236
		v247 = v235
		goto L68
	} else {
		goto L72
	}
L71:
	;
	v246 = v236
	v247 = v235
	goto L68
L72:
	;
	v239 = int32(1)
	if v236 == v235 {
		v231 = v231 + v239
		v232 = v232 + v239
		goto L70
	} else {
		goto L73
	}
L73:
	;
	goto L71
L74:
	;
	v252 = v206 + int32(1)
	if v200 != v252 {
		v206 = v252
		goto L65
	} else {
		goto L75
	}
L75:
	;
	goto L66
L76:
	;
	v270 = F_pstrdup(m, v167)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L5
	} else {
		goto L77
	}
L77:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v268)+12)) = int64(2147483647)
	v274 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v268)+8)) = uint16(v274)
	*(*int32)(unsafe.Add(mBase, uint32(v268)+4)) = v274
	*(*int32)(unsafe.Add(mBase, uint32(v268))) = v270
	v279 = F_lappend(m, v47, v268)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L5
	} else {
		goto L78
	}
L78:
	;
	v283 = v279
	v284 = v268
	goto L61
L79:
	;
	if v47 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L80:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v506)+4))
	v517 = F_lappend(m, v516, v508)
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L5
	} else {
		goto L124
	}
L81:
	;
	v490 = F_palloc(m, int32(20))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L5
	} else {
		goto L121
	}
L82:
	;
	if int32(0) < v403 {
		goto L107
	} else {
		goto L108
	}
L83:
	;
	v376 = F_palloc(m, int32(20))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L5
	} else {
		goto L100
	}
L84:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	if v304 <= int32(0) {
		goto L83
	} else {
		goto L85
	}
L85:
	;
	v307 = int32(0)
	if v307 < v304 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v310 = v304
	goto L88
L87:
	;
	v310 = v307
	goto L88
L88:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
	v314 = int32(0)
	goto L89
L89:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v311+v314<<(uint(int32(2))%32))))
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v329)))
	v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330))))
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
	if base.B2i32(v333 == int32(0))|base.B2i32(v333 != v336) != 0 {
		v354 = v333
		v355 = v336
		goto L92
	} else {
		goto L93
	}
L90:
	;
	goto L83
L91:
	;
	if v354-v355 == int32(0) {
		v399 = v47
		v400 = v329
		v403 = v304
		v405 = v310
		goto L82
	} else {
		goto L98
	}
L92:
	;
	goto L91
L93:
	;
	v339 = v330
	v340 = v167
	goto L94
L94:
	;
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v340)+1)))
	v344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v339)+1)))
	if v344 == int32(0) {
		v354 = v344
		v355 = v343
		goto L92
	} else {
		goto L96
	}
L95:
	;
	v354 = v344
	v355 = v343
	goto L92
L96:
	;
	v347 = int32(1)
	if v344 == v343 {
		v339 = v339 + v347
		v340 = v340 + v347
		goto L94
	} else {
		goto L97
	}
L97:
	;
	goto L95
L98:
	;
	v360 = v314 + int32(1)
	if v304 != v360 {
		v314 = v360
		goto L89
	} else {
		goto L99
	}
L99:
	;
	goto L90
L100:
	;
	v378 = F_pstrdup(m, v167)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L5
	} else {
		goto L101
	}
L101:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v376)+12)) = int64(2147483647)
	v382 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v376)+8)) = uint16(v382)
	*(*int32)(unsafe.Add(mBase, uint32(v376)+4)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v376))) = v378
	v388 = F_lappend(m, v47, v376)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L5
	} else {
		goto L102
	}
L102:
	;
	if v388 == int32(0) {
		v477 = v382
		v479 = v376
		goto L81
	} else {
		goto L103
	}
L103:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v388)+4))
	v393 = int32(0)
	if v393 < v392 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v396 = v392
	goto L106
L105:
	;
	v396 = v393
	goto L106
L106:
	;
	v399 = v388
	v400 = v376
	v403 = v392
	v405 = v396
	goto L82
L107:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v399)+12))
	v415 = int32(0)
	goto L110
L108:
	;
	goto L109
L109:
	;
	v477 = v399
	v479 = v400
	goto L81
L110:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v412+v415<<(uint(int32(2))%32))))
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v430)))
	v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v431))))
	v437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v299))))
	if base.B2i32(v434 == int32(0))|base.B2i32(v434 != v437) != 0 {
		v455 = v434
		v456 = v437
		goto L113
	} else {
		goto L114
	}
L111:
	;
	goto L109
L112:
	;
	if v455-v456 == int32(0) {
		v505 = v399
		v506 = v400
		v508 = v430
		goto L80
	} else {
		goto L119
	}
L113:
	;
	goto L112
L114:
	;
	v440 = v431
	v441 = v299
	goto L115
L115:
	;
	v444 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v441)+1)))
	v445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v440)+1)))
	if v445 == int32(0) {
		v455 = v445
		v456 = v444
		goto L113
	} else {
		goto L117
	}
L116:
	;
	v455 = v445
	v456 = v444
	goto L113
L117:
	;
	v448 = int32(1)
	if v445 == v444 {
		v440 = v440 + v448
		v441 = v441 + v448
		goto L115
	} else {
		goto L118
	}
L118:
	;
	goto L116
L119:
	;
	v461 = v415 + int32(1)
	if v461 != v405 {
		v415 = v461
		goto L110
	} else {
		goto L120
	}
L120:
	;
	goto L111
L121:
	;
	v492 = F_pstrdup(m, v299)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L5
	} else {
		goto L122
	}
L122:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v490)+12)) = int64(2147483647)
	v496 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v490)+8)) = uint16(v496)
	*(*int32)(unsafe.Add(mBase, uint32(v490)+4)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v490))) = v492
	v501 = F_lappend(m, v477, v490)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L5
	} else {
		goto L123
	}
L123:
	;
	v505 = v501
	v506 = v479
	v508 = v490
	goto L80
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v506)+4)) = v517
	v522 = v505
	goto L19
L125:
	;
	if v533 != 0 {
		v46 = v533
		v47 = v522
		goto L17
	} else {
		goto L126
	}
L126:
	;
	goto L18
L127:
	;
	m.G0 = v16 + int32(16)
	return v537
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
	var v43 int32
	_ = v43
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
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v100 float64
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v114 float64
	_ = v114
	var v115 int32
	_ = v115
	var v117 float64
	_ = v117
	var v118 float64
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 float64
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v149 float64
	_ = v149
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v199 float64
	_ = v199
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v210 float64
	_ = v210
	var v211 int32
	_ = v211
	var v213 float64
	_ = v213
	var v233 float64
	_ = v233
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
	return v233
L2:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v151 == int32(0) {
		v233 = v149
		goto L1
	} else {
		goto L32
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
		goto L29
	} else {
		goto L30
	}
L6:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v25 == int32(0) {
		v149 = v16
		goto L2
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	v123 = F_get_sortgrouplist_exprs(m, v122, l3)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L13
	} else {
		goto L27
	}
L9:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v28 <= int32(0) {
		v149 = v16
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v43 = int32(0)
	v46 = v16
	goto L11
L11:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v48+v43<<(uint(int32(2))%32))))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v54 = F_get_sortgrouplist_exprs(m, v53, l3)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return float64(0)
L14:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v52)+16)) = int64(0)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	v68 = int32(0)
	v80 = float64(0)
	goto L15
L15:
	;
	v81 = int32(0)
	if v61 == v81 {
		v90 = v81
		goto L17
	} else {
		goto L18
	}
L17:
	;
	if v60 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v84 <= v68 {
		v90 = v81
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v90 = v86 + v68<<(uint(int32(2))%32)
	goto L17
L20:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v105
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v98+v68<<(uint(int32(2))%32))))
	v114 = F_estimate_num_groups(m, l0, v54, l1, v20+int32(12), int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L13
	} else {
		goto L26
	}
L21:
	;
	v100 = base.F64_add(v46, v80)
	v102 = v43 + int32(1)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v103 <= v102 {
		v149 = v100
		goto L2
	} else {
		goto L25
	}
L22:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	if base.B2i32(v90 == int32(0))|base.B2i32(v95 <= v68) != 0 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
	if v98 != 0 {
		goto L20
	} else {
		goto L24
	}
L24:
	;
	goto L21
L25:
	;
	v43 = v102
	v46 = v100
	goto L11
L26:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v110)+8)) = v114
	v117 = *(*float64)(unsafe.Add(mBase, uint32(v52)+16))
	v118 = base.F64_add(v114, v117)
	*(*float64)(unsafe.Add(mBase, uint32(v52)+16)) = v118
	v68 = v68 + int32(1)
	v80 = v118
	goto L15
L27:
	;
	v125 = int32(0)
	v127 = F_estimate_num_groups(m, l0, v123, l1, v125, v125)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L13
	} else {
		goto L28
	}
L28:
	;
	v233 = v127
	goto L1
L29:
	;
	v233 = float64(1)
	goto L1
L30:
	;
	goto L31
L31:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v233 = base.F64_convert_i32_s(v132)
	goto L1
L32:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2)+8)) = int64(0)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v22)+100))
	v157 = F_get_sortgrouplist_exprs(m, v156, l3)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L13
	} else {
		goto L33
	}
L33:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v166 = int32(0)
	goto L34
L34:
	;
	v179 = int32(0)
	if v160 == v179 {
		v189 = v179
		goto L36
	} else {
		goto L37
	}
L36:
	;
	if v159 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v160)+4))
	if v183 <= v166 {
		v189 = int32(0)
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v160)+12))
	v189 = v185 + v166<<(uint(int32(2))%32)
	goto L36
L39:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v201
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v197+v166<<(uint(int32(2))%32))))
	v210 = F_estimate_num_groups(m, l0, v157, l1, v20+int32(8), int32(0))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L13
	} else {
		goto L44
	}
L40:
	;
	v199 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
	v233 = base.F64_add(v149, v199)
	goto L1
L41:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v159)+4))
	if base.B2i32(v189 == int32(0))|base.B2i32(v194 <= v166) != 0 {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v159)+12))
	if v197 != 0 {
		goto L39
	} else {
		goto L43
	}
L43:
	;
	goto L40
L44:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v206)+8)) = v210
	v213 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
	*(*float64)(unsafe.Add(mBase, uint32(l2)+8)) = base.F64_add(v210, v213)
	v166 = v166 + int32(1)
	goto L34
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
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
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
		v60 = v3
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	if v61 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L7:
	;
	v22 = int32(_a_F_get_returning_clause_1)
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_get_returning_clause[0])))
	if base.B2i32(v25 == int32(0))|base.B2i32(v25 != v28) != 0 {
		v46 = v25
		v47 = v28
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if v46-v47 == int32(0) {
		v60 = v3
		goto L6
	} else {
		goto L15
	}
L9:
	;
	goto L8
L10:
	;
	v31 = v19
	v32 = v22
	goto L11
L11:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+1)))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
	if v36 == int32(0) {
		v46 = v36
		v47 = v35
		goto L9
	} else {
		goto L13
	}
L12:
	;
	v46 = v36
	v47 = v35
	goto L9
L13:
	;
	v39 = int32(1)
	if v36 == v35 {
		v31 = v31 + v39
		v32 = v32 + v39
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v51 = F_quote_identifier(m, v19)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v51
	F_appendStringInfo(m, v12, int32(_a_F_get_returning_clause_2), v9+int32(16))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	v60 = int32(1)
	goto L6
L18:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	F_get_target_list(m, v106, l1)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L4
	} else {
		goto L37
	}
L19:
	;
	F_appendStringInfoChar(m, v12, int32(41))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L4
	} else {
		goto L36
	}
L20:
	;
	if v60 == int32(0) {
		goto L18
	} else {
		goto L35
	}
L21:
	;
	v64 = int32(_a_F_get_returning_clause_3)
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_get_returning_clause[1])))
	if base.B2i32(v67 == int32(0))|base.B2i32(v67 != v70) != 0 {
		v88 = v67
		v89 = v70
		goto L23
	} else {
		goto L24
	}
L22:
	;
	if v88-v89 == int32(0) {
		goto L20
	} else {
		goto L29
	}
L23:
	;
	goto L22
L24:
	;
	v73 = v61
	v74 = v64
	goto L25
L25:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+1)))
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+1)))
	if v78 == int32(0) {
		v88 = v78
		v89 = v77
		goto L23
	} else {
		goto L27
	}
L26:
	;
	v88 = v78
	v89 = v77
	goto L23
L27:
	;
	v81 = int32(1)
	if v78 == v77 {
		v73 = v73 + v81
		v74 = v74 + v81
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v93 = F_quote_identifier(m, v61)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L4
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v93
	if v60 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v98 = int32(_a_F_get_returning_clause_4)
	goto L33
L32:
	;
	v98 = int32(_a_F_get_returning_clause_5)
	goto L33
L33:
	;
	F_appendStringInfo(m, v12, v98, v9)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	goto L19
L35:
	;
	goto L19
L36:
	;
	goto L18
L37:
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
	var v10 int64
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
	var v36 int64
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
	var v58 int64
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
		v10 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+8)))
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
		v36 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_get_rolespec_tuple[0])))
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
		v58 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_get_rolespec_tuple[1])))
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
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
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
			v39 = v34
			v42 = int32(0)
			for {
				v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				if v42 < v48 {
					v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v54 = int32(*(*int16)(unsafe.Add(mBase, uint32(v50+v42<<(uint(int32(1))%32)))))
					v55 = v54
				} else {
					v55 = int32(0)
				}
				v56 = int32(1000)
				v57 = base.I32_div_s(v55, v56)
				v62 = v57&int32(_a_F_get_str_from_var_1)*int32(_a_F_get_str_from_var_2) + v55
				v65 = int32(0)
				if base.B2i32(v55 < v56)&base.B2i32(v42 <= v65) == v65 {
					v71 = v57 + int32(48)
					*(*uint8)(unsafe.Add(mBase, uint32(v39))) = uint8(v71)
					v77 = base.I32_div_s(base.I32_extend16_s(v62), int32(100))
					v93 = v39 + int32(1)
					v94 = v62 + v77&int32(_a_F_get_str_from_var_1)*int32(_a_F_get_str_from_var_3)
					v95 = v77
					v98 = v95 + int32(48)
					*(*uint8)(unsafe.Add(mBase, uint32(v93))) = uint8(v98)
					v104 = base.I32_div_s(base.I32_extend16_s(v94), int32(10))
					v120 = v93 + int32(1)
					v121 = v104&int32(_a_F_get_str_from_var_1)*int32(246) + v94
					v122 = v104
					v125 = v122 + int32(48)
					*(*uint8)(unsafe.Add(mBase, uint32(v120))) = uint8(v125)
					v129 = v120 + int32(1)
					v130 = v121
				} else {
					v83 = base.I32_extend16_s(v62)
					v84 = int32(100)
					v85 = base.I32_div_s(v83, v84)
					v90 = v62 + v85&int32(_a_F_get_str_from_var_1)*int32(_a_F_get_str_from_var_3)
					if v83 < v84 {
						v110 = base.I32_extend16_s(v90)
						v111 = int32(10)
						v112 = base.I32_div_s(v110, v111)
						v117 = v112&int32(_a_F_get_str_from_var_1)*int32(246) + v90
						if v110 < v111 {
							v129 = v39
							v130 = v117
						} else {
							v120 = v39
							v121 = v117
							v122 = v112
							v125 = v122 + int32(48)
							*(*uint8)(unsafe.Add(mBase, uint32(v120))) = uint8(v125)
							v129 = v120 + int32(1)
							v130 = v121
						}
					} else {
						v93 = v39
						v94 = v90
						v95 = v85
						v98 = v95 + int32(48)
						*(*uint8)(unsafe.Add(mBase, uint32(v93))) = uint8(v98)
						v104 = base.I32_div_s(base.I32_extend16_s(v94), int32(10))
						v120 = v93 + int32(1)
						v121 = v104&int32(_a_F_get_str_from_var_1)*int32(246) + v94
						v122 = v104
						v125 = v122 + int32(48)
						*(*uint8)(unsafe.Add(mBase, uint32(v120))) = uint8(v125)
						v129 = v120 + int32(1)
						v130 = v121
					}
				}
				v134 = v130 + int32(48)
				*(*uint8)(unsafe.Add(mBase, uint32(v129))) = uint8(v134)
				v136 = int32(1)
				v137 = v129 + v136
				v139 = v42 + v136
				v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if v139 <= v140 {
					v39 = v137
					v42 = v139
					continue
				} else {
					break
				}
				break
			}
			v149 = v137
			v152 = v139
		} else {
			v142 = int32(48)
			*(*uint8)(unsafe.Add(mBase, uint32(v34))) = uint8(v142)
			v144 = int32(1)
			v149 = v34 + v144
			v152 = v35 + v144
		}
		if int32(0) < v10 {
			v159 = int32(46)
			*(*uint8)(unsafe.Add(mBase, uint32(v149))) = uint8(v159)
			v163 = v149 + int32(1)
			v165 = v163
			v167 = int32(0)
			v168 = v152
			for {
				v173 = int32(0)
				if v168 < v173 {
					v184 = v173
				} else {
					v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					if v177 <= v168 {
						v184 = int32(0)
					} else {
						v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v183 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v179+v168<<(uint(int32(1))%32)))))
						v184 = v183
					}
				}
				v187 = base.I32_div_s(base.I32_extend16_s(v184), int32(1000))
				v188 = int32(48)
				v189 = v187 + v188
				*(*uint8)(unsafe.Add(mBase, uint32(v165))) = uint8(v189)
				v191 = int32(_a_F_get_str_from_var_1)
				v195 = v187&v191*int32(_a_F_get_str_from_var_2) + v184
				v198 = base.I32_div_s(base.I32_extend16_s(v195), int32(100))
				v200 = v198 + v188
				*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)) = uint8(v200)
				v206 = v198&v191*int32(_a_F_get_str_from_var_3) + v195
				v209 = base.I32_div_s(base.I32_extend16_s(v206), int32(10))
				v211 = v209 + v188
				*(*uint8)(unsafe.Add(mBase, uint32(v165)+2)) = uint8(v211)
				v217 = v209*int32(246) + v206 + v188
				*(*uint8)(unsafe.Add(mBase, uint32(v165)+3)) = uint8(v217)
				v221 = int32(4)
				v224 = v167 + v221
				if v224 < v10 {
					v165 = v165 + v221
					v167 = v224
					v168 = v168 + int32(1)
					continue
				} else {
					break
				}
				break
			}
			v236 = v163 + v10
		} else {
			v236 = v149
		}
		v237 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v236))) = uint8(v237)
		return v23
	}
}
func F_get_user_default_acl(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v28 int64
	_ = v28
	var v30 int64
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v40 int64
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int64
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
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
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	v4 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_get_user_default_acl[0]))
	if v17 == v4 {
		v218 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v14 + int32(16)
	return v218
L2:
	;
	switch l0 - int32(19) {
	case 0:
		goto L7
	default:
		v218 = v4
		goto L1
	case 3:
		goto L4
	case 18:
		goto L5
	case 19:
		goto L8
	case 23:
		v28 = int64(114)
		goto L3
	case 31:
		goto L6
	}
L3:
	;
	v30 = base.I64_extend_i32_u(l1)
	v32 = F_SearchSysCache3(m, int32(22), v30, int64(0), v28)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	v28 = int64(76)
	goto L3
L5:
	;
	v28 = int64(110)
	goto L3
L6:
	;
	v28 = int64(84)
	goto L3
L7:
	;
	v28 = int64(102)
	goto L3
L8:
	;
	v28 = int64(83)
	goto L3
L9:
	;
	return int32(0)
L10:
	;
	if v32 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v40 = F_SysCacheGetAttr(m, int32(22), v32, int32(5), v14+int32(14))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L9
	} else {
		goto L14
	}
L12:
	;
	v51 = v4
	goto L13
L13:
	;
	v55 = F_SearchSysCache3(m, int32(22), v30, base.I64_extend_i32_u(l2), v28)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L9
	} else {
		goto L20
	}
L14:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+14)))
	if v42 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v46 = F_pg_detoast_datum_copy(m, base.I32_wrap_i64(v40))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L9
	} else {
		goto L18
	}
L16:
	;
	v48 = v4
	goto L17
L17:
	;
	F_ReleaseCatCache(m, v32)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L9
	} else {
		goto L19
	}
L18:
	;
	v48 = v46
	goto L17
L19:
	;
	v51 = v48
	goto L13
L20:
	;
	if v55 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v61 = F_SysCacheGetAttr(m, int32(22), v55, int32(5), v14+int32(15))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L9
	} else {
		goto L24
	}
L22:
	;
	v72 = v4
	goto L23
L23:
	;
	v74 = int32(0)
	if v72|v51 == v74 {
		v218 = v74
		goto L1
	} else {
		goto L30
	}
L24:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+15)))
	if v63 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v67 = F_pg_detoast_datum_copy(m, base.I32_wrap_i64(v61))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L9
	} else {
		goto L28
	}
L26:
	;
	v69 = v4
	goto L27
L27:
	;
	F_ReleaseCatCache(m, v55)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L9
	} else {
		goto L29
	}
L28:
	;
	v69 = v67
	goto L27
L29:
	;
	v72 = v69
	goto L23
L30:
	;
	v78 = F_acldefault(m, l0, l1)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L9
	} else {
		goto L34
	}
L31:
	;
	F_aclitemsort(m, v161)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L9
	} else {
		goto L63
	}
L32:
	;
	v161 = v141
	goto L31
L33:
	;
	if v72 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L34:
	;
	if v51 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v80 = v51
	goto L37
L36:
	;
	v80 = v78
	goto L37
L37:
	;
	if v80 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
	if v81 != 0 {
		goto L33
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v82 = int32(0)
	if v72 == v82 {
		v141 = v82
		goto L32
	} else {
		goto L42
	}
L41:
	;
	goto L40
L42:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v72)+16))
	if v85 == int32(0) {
		v141 = v82
		goto L32
	} else {
		goto L43
	}
L43:
	;
	v88 = F_aclcopy(m, v72)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L9
	} else {
		goto L44
	}
L44:
	;
	v161 = v88
	goto L31
L45:
	;
	v92 = F_aclcopy(m, v80)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L9
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v72)+16))
	if v94 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v161 = v92
	goto L31
L49:
	;
	v97 = F_aclcopy(m, v80)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L9
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v99 = F_aclcopy(m, v80)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L9
	} else {
		goto L53
	}
L52:
	;
	v161 = v97
	goto L31
L53:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v72)+8))
	if v101 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	v111 = (v104<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L56
L55:
	;
	v111 = v101
	goto L56
L56:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v72)+16))
	if v112 <= int32(0) {
		v141 = v99
		goto L32
	} else {
		goto L57
	}
L57:
	;
	v119 = v99
	v120 = v72 + v111
	v121 = int32(0)
	goto L58
L58:
	;
	v130 = F_aclupdate(m, v119, v120, int32(1), l1, int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L9
	} else {
		goto L60
	}
L59:
	;
	v141 = v130
	goto L32
L60:
	;
	F_pfree(m, v119)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L9
	} else {
		goto L61
	}
L61:
	;
	v137 = v121 + int32(1)
	if v137 != v112 {
		v119 = v130
		v120 = v120 + int32(16)
		v121 = v137
		goto L58
	} else {
		goto L62
	}
L62:
	;
	goto L59
L63:
	;
	F_aclitemsort(m, v78)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L9
	} else {
		goto L64
	}
L64:
	;
	v166 = int32(0)
	if v161 != 0 {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	if v211 != 0 {
		goto L83
	} else {
		goto L84
	}
L66:
	;
	if v78 == int32(0) {
		v207 = v166
		goto L74
	} else {
		goto L75
	}
L67:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v161)+16))
	if v169 != 0 {
		goto L66
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	if v78 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	goto L69
L71:
	;
	v211 = int32(1)
	goto L65
L72:
	;
	goto L73
L73:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v78)+16))
	v211 = base.B2i32(v174 == int32(0))
	goto L65
L74:
	;
	v211 = v207
	goto L65
L75:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v78)+16))
	if v169 != v179 {
		v207 = v166
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v161)+8))
	if v181 != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v189 = v181
	goto L79
L78:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v161)+4))
	v189 = (v182<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L79
L79:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v78)+8))
	if v191 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v199 = v191
	goto L82
L81:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	v199 = (v192<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L82
L82:
	;
	v203 = F_memcmp(m, v189+v161, v199+v78, v169<<(uint(int32(4))%32))
	mBase = m.M
	v207 = base.B2i32(v203 == int32(0))
	goto L74
L83:
	;
	v212 = v166
	goto L85
L84:
	;
	v212 = v161
	goto L85
L85:
	;
	v218 = v212
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
	var v39 int32
	_ = v39
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
	var v81 int32
	_ = v81
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
	var v201 int32
	_ = v201
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
	var v285 int32
	_ = v285
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
	var v346 int64
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v354 int64
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
	v39 = v38
	v47 = v3
	goto L14
L14:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v48+v47<<(uint(int32(2))%32))))
	F_appendStringInfoString(m, v17, v39)
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
	v81 = v77
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
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v90+v81<<(uint(int32(2))%32))))
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
	v104 = v81 + int32(1)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	if v104 < v105 {
		v81 = v104
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
	v201 = v197
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
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v210+v201<<(uint(int32(2))%32))))
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
	v224 = v201 + int32(1)
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v186)+4))
	if v224 < v225 {
		v201 = v224
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
		v39 = int32(_a_F_get_with_clause_3)
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
	v285 = v281
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
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v294+v285<<(uint(int32(2))%32))))
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
	v308 = v285 + int32(1)
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v270)+4))
	if v308 < v309 {
		v285 = v308
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
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341)+32)))
	if v345 != 0 {
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v346 = *(*int64)(unsafe.Add(mBase, uint32(v341)+24))
	if v346 == int64(0) {
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
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349)+32)))
	if v353 != 0 {
		goto L96
	} else {
		goto L101
	}
L101:
	;
	v354 = *(*int64)(unsafe.Add(mBase, uint32(v349)+24))
	if v354 == int64(0) {
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
func F_get_xid_status(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int64
	_ = v25
	var v32 int64
	_ = v32
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int64
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int64
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v67 int64
	_ = v67
	var v75 int64
	_ = v75
	var v77 int32
	_ = v77
	var v80 int64
	_ = v80
	var v90 int64
	_ = v90
	var v91 int64
	_ = v91
	var v92 int64
	_ = v92
	var v99 int64
	_ = v99
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int64
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int64
	_ = v129
	var v138 int64
	_ = v138
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v262 int32
	_ = v262
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	if base.Ui32(l0) <= base.Ui32(int32(2)) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
	return int32(4)
L5:
	;
	goto L6
L6:
	;
	v19 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v20 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v22 = v21 - l0
	if int32(0) < v22 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	if base.Ui64(v90) <= base.Ui64(v91) {
		goto L29
	} else {
		goto L30
	}
L8:
	;
	if base.B2i32(base.Ui64(v20) <= base.Ui64(v32))&base.B2i32(base.Ui64(v32) < base.Ui64(v19)) != 0 {
		v90 = v19
		v91 = v32
		v92 = v20
		goto L7
	} else {
		goto L13
	}
L9:
	;
	v25 = int64(3)
	if base.Ui64(v19-v25) < base.Ui64(base.I64_extend_i32_u(v22)) {
		v32 = v25
		goto L8
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v32 = v19 - base.I64_extend_i32_s(v22)
	goto L8
L12:
	;
	goto L11
L13:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_get_xid_status[0]))
	v41 = F_LWLockAcquire(m, v37+int32(384), int32(1))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return int32(0)
L15:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_get_xid_status[1]))
	v47 = *(*int64)(unsafe.Add(mBase, uint32(v46)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v47
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v49
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_get_xid_status[0]))
	F_LWLockRelease(m, v52+int32(384))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v57 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v58 = base.I32_wrap_i64(v57)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v58
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if base.Ui32(v60) <= base.Ui32(int32(2)) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1)+16)) = v75
	v77 = v58 - l0
	if int32(0) < v77 {
		goto L25
	} else {
		goto L26
	}
L18:
	;
	v75 = base.I64_extend_i32_u(v60)
	goto L17
L19:
	;
	goto L20
L20:
	;
	v64 = v58 - v60
	if int32(0) < v64 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v67 = int64(3)
	if base.Ui64(v57-v67) < base.Ui64(base.I64_extend_i32_u(v64)) {
		v75 = v67
		goto L17
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v75 = v57 - base.I64_extend_i32_s(v64)
	goto L17
L24:
	;
	goto L23
L25:
	;
	v80 = int64(3)
	if base.Ui64(v57-v80) < base.Ui64(base.I64_extend_i32_u(v77)) {
		v90 = v57
		v91 = v80
		v92 = v75
		goto L7
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v90 = v57
	v91 = v57 - base.I64_extend_i32_s(v77)
	v92 = v75
	goto L7
L28:
	;
	goto L27
L29:
	;
	return int32(1)
L30:
	;
	goto L31
L31:
	;
	if base.Ui64(v91) < base.Ui64(v92) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	return int32(2)
L33:
	;
	goto L34
L34:
	;
	v99 = *(*int64)(unsafe.Add(mBase, uint32(l1)+56))
	if base.Ui64(v91) < base.Ui64(v99) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	return int32(3)
L36:
	;
	goto L37
L37:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v103 == l0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v105
	return int32(4)
L39:
	;
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
	v112 = *(*int32)(unsafe.Add(mBase, _c_F_get_xid_status[0]))
	v116 = F_LWLockAcquire(m, v112+int32(_a_F_get_xid_status_0), int32(1))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L14
	} else {
		goto L41
	}
L41:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_get_xid_status[1]))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+64))
	if base.Ui32(v120) <= base.Ui32(int32(2)) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	if base.Ui64(v138) <= base.Ui64(v91) {
		goto L50
	} else {
		goto L51
	}
L43:
	;
	v138 = base.I64_extend_i32_u(v120)
	goto L42
L44:
	;
	goto L45
L45:
	;
	v124 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v126 = v125 - v120
	if int32(0) < v126 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v129 = int64(3)
	if base.Ui64(v124-v129) < base.Ui64(base.I64_extend_i32_u(v126)) {
		v138 = v129
		goto L42
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v138 = v124 - base.I64_extend_i32_s(v126)
	goto L42
L49:
	;
	goto L48
L50:
	;
	if base.Ui32(l0) < base.Ui32(int32(3)) {
		goto L55
	} else {
		goto L56
	}
L51:
	;
	goto L52
L52:
	;
	v284 = *(*int32)(unsafe.Add(mBase, _c_F_get_xid_status[0]))
	F_LWLockRelease(m, v284+int32(_a_F_get_xid_status_0))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L14
	} else {
		goto L101
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v281
	goto L52
L54:
	;
	if v272 != 0 {
		v281 = int32(1)
		goto L53
	} else {
		goto L94
	}
L55:
	;
	v272 = int32(0)
	goto L54
L56:
	;
	goto L57
L57:
	;
	v152 = *(*int32)(unsafe.Add(mBase, _c_F_get_xid_status[2]))
	if v152 == l0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v272 = int32(1)
	goto L54
L59:
	;
	goto L60
L60:
	;
	v156 = *(*int32)(unsafe.Add(mBase, _c_F_get_xid_status[3]))
	if v156 <= int32(0) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v272 = v262
	goto L54
L62:
	;
	v160 = *(*int32)(unsafe.Add(mBase, _c_F_get_xid_status[4]))
	if v160 == int32(0) {
		v262 = int32(0)
		goto L61
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v230 = *(*int32)(unsafe.Add(mBase, _c_F_get_xid_status[5]))
	v232 = int32(0)
	v235 = v156 - int32(1)
	goto L84
L65:
	;
	v165 = v160
	goto L66
L66:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v165)+20))
	if v171 == int32(4) {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	v262 = int32(0)
	goto L61
L68:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v165)+80))
	if v225 != 0 {
		v165 = v225
		goto L66
	} else {
		goto L83
	}
L69:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	if v174 == int32(0) {
		goto L68
	} else {
		goto L70
	}
L70:
	;
	v177 = int32(1)
	if l0 == v174 {
		v262 = v177
		goto L61
	} else {
		goto L71
	}
L71:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v165)+52))
	v181 = v179 - int32(1)
	if v181 < int32(0) {
		goto L68
	} else {
		goto L72
	}
L72:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v165)+48))
	v187 = int32(0)
	v190 = v181
	goto L73
L73:
	;
	v195 = int32(2)
	v196 = base.I32_div_s(v190-v187, v195)
	v197 = v196 + v187
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v184+v197<<(uint(v195)%32))))
	if v201 == l0 {
		v262 = v177
		goto L61
	} else {
		goto L75
	}
L74:
	;
	goto L68
L75:
	;
	v210 = base.B2i32(v201-l0 < int32(0)) | base.B2i32(base.Ui32(v201) < base.Ui32(int32(3)))
	if v210 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v211 = v197 + int32(1)
	goto L78
L77:
	;
	v211 = v187
	goto L78
L78:
	;
	if v210 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v214 = v190
	goto L81
L80:
	;
	v214 = v197 - int32(1)
	goto L81
L81:
	;
	if v211 <= v214 {
		v187 = v211
		v190 = v214
		goto L73
	} else {
		goto L82
	}
L82:
	;
	goto L74
L83:
	;
	goto L67
L84:
	;
	v240 = int32(2)
	v241 = base.I32_div_s(v235-v232, v240)
	v242 = v241 + v232
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v230+v242<<(uint(v240)%32))))
	v247 = base.B2i32(v246 == l0)
	if v246 == l0 {
		v262 = v247
		goto L61
	} else {
		goto L86
	}
L85:
	;
	v262 = v247
	goto L61
L86:
	;
	v250 = base.B2i32(base.Ui32(v246) < base.Ui32(l0))
	if base.Ui32(v246) < base.Ui32(l0) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v251 = v242 + int32(1)
	goto L89
L88:
	;
	v251 = v232
	goto L89
L89:
	;
	if base.Ui32(v246) < base.Ui32(l0) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v254 = v235
	goto L92
L91:
	;
	v254 = v242 - int32(1)
	goto L92
L92:
	;
	if v251 <= v254 {
		v232 = v251
		v235 = v254
		goto L84
	} else {
		goto L93
	}
L93:
	;
	goto L85
L94:
	;
	v274 = F_TransactionIdIsInProgress(m, l0)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L14
	} else {
		goto L95
	}
L95:
	;
	if v274 != 0 {
		v281 = int32(2)
		goto L53
	} else {
		goto L96
	}
L96:
	;
	v278 = F_TransactionIdDidCommit(m, l0)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L14
	} else {
		goto L97
	}
L97:
	;
	if v278 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v280 = int32(0)
	goto L100
L99:
	;
	v280 = int32(3)
	goto L100
L100:
	;
	v281 = v280
	goto L53
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = l0
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v290
	return int32(4)
}
func F_getmissingattr(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v37 int32
	_ = v37
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int64
	_ = v64
	var v65 int64
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v93 int64
	_ = v93
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v112 int64
	_ = v112
	var v114 int32
	_ = v114
	var v115 int64
	_ = v115
	var v116 int32
	_ = v116
	var v121 int64
	_ = v121
	var v123 int32
	_ = v123
	var v132 int64
	_ = v132
	v8 = m.G0
	v10 = v8 + int32(-64)
	m.G0 = v10
	v13 = l1 - int32(1)
	v16 = l0 + v13<<(uint(int32(3))%32)
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+34)))
	if v17&int32(2) == int32(0) {
		v123 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v123)
		v132 = int64(0)
		m.G0 = v10 - int32(-64)
		return v132
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
		v26 = v23 + v13<<(uint(int32(4))%32)
		v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
		if v27 != int32(1) {
			v123 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v123)
			v132 = int64(0)
			m.G0 = v10 - int32(-64)
			return v132
		} else {
			v30 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v30)
			v33 = v16 + int32(28)
			v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+4)))
			if v34 != 0 {
				v35 = *(*int64)(unsafe.Add(mBase, uint32(v26)+8))
				v132 = v35
				m.G0 = v10 - int32(-64)
				return v132
			} else {
				v37 = *(*int32)(unsafe.Add(mBase, _c_F_getmissingattr[0]))
				if v37 == int32(0) {
					*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = int64(68719476752)
					*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = int32(24)
					*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = int32(25)
					v47 = *(*int32)(unsafe.Add(mBase, _c_F_getmissingattr[1]))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = v47
					v55 = F_hash_create(m, int32(_a_F_getmissingattr_0), int64(32), v8+int32(-48), int32(1224))
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_getmissingattr[0])) = v55
						v60 = v55
						v61 = int32(*(*int16)(unsafe.Add(mBase, uint32(v33)+2)))
						if int32(0) < v61 {
							v64 = *(*int64)(unsafe.Add(mBase, uint32(v26)+8))
							v91 = v61
							v93 = v64
						} else {
							v65 = *(*int64)(unsafe.Add(mBase, uint32(v26)+8))
							v66 = base.I32_wrap_i64(v65)
							v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
							if v67 == int32(1) {
								v71 = int32(18)
								v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+1)))
								if v73 == v71 {
									v76 = v71
								} else {
									v76 = int32(2)
								}
								if base.Ui32((v73-int32(1))&int32(255)) < base.Ui32(int32(3)) {
									v83 = int32(6)
								} else {
									v83 = v76
								}
								v91 = v83
								v93 = v65
							} else {
								if v67&int32(1) != 0 {
									v91 = int32(base.Ui32(v67) >> (uint(int32(1)) % 32))
									v93 = v65
								} else {
									v88 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
									v91 = int32(base.Ui32(v88) >> (uint(int32(2)) % 32))
									v93 = v65
								}
							}
						}
						*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v93
						*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v91
						v101 = F_hash_search(m, v60, v8+int32(-48), int32(1), v8+int32(-49))
						mBase = m.M
						v102 = m.ExcPending
						if v102 != 0 {
							return int64(0)
						} else {
							v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
							if v103 == int32(0) {
								v106 = int32(_a_F_getmissingattr_1)
								v107 = *(*int32)(unsafe.Add(mBase, _c_F_getmissingattr[2]))
								v110 = *(*int32)(unsafe.Add(mBase, _c_F_getmissingattr[1]))
								*(*int32)(unsafe.Add(mBase, _c_F_getmissingattr[2])) = v110
								v112 = *(*int64)(unsafe.Add(mBase, uint32(v26)+8))
								v114 = int32(*(*int16)(unsafe.Add(mBase, uint32(v33)+2)))
								v115 = F_datumCopy(m, v112, int32(0), v114)
								mBase = m.M
								v116 = m.ExcPending
								if v116 != 0 {
									return int64(0)
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v101)+8)) = v115
									*(*int32)(unsafe.Add(mBase, _c_F_getmissingattr[2])) = v107
									v121 = *(*int64)(unsafe.Add(mBase, uint32(v101)+8))
									v132 = v121
									m.G0 = v10 - int32(-64)
									return v132
								}
							} else {
								v121 = *(*int64)(unsafe.Add(mBase, uint32(v101)+8))
								v132 = v121
								m.G0 = v10 - int32(-64)
								return v132
							}
						}
					}
				} else {
					v60 = v37
					v61 = int32(*(*int16)(unsafe.Add(mBase, uint32(v33)+2)))
					if int32(0) < v61 {
						v64 = *(*int64)(unsafe.Add(mBase, uint32(v26)+8))
						v91 = v61
						v93 = v64
					} else {
						v65 = *(*int64)(unsafe.Add(mBase, uint32(v26)+8))
						v66 = base.I32_wrap_i64(v65)
						v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
						if v67 == int32(1) {
							v71 = int32(18)
							v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+1)))
							if v73 == v71 {
								v76 = v71
							} else {
								v76 = int32(2)
							}
							if base.Ui32((v73-int32(1))&int32(255)) < base.Ui32(int32(3)) {
								v83 = int32(6)
							} else {
								v83 = v76
							}
							v91 = v83
							v93 = v65
						} else {
							if v67&int32(1) != 0 {
								v91 = int32(base.Ui32(v67) >> (uint(int32(1)) % 32))
								v93 = v65
							} else {
								v88 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
								v91 = int32(base.Ui32(v88) >> (uint(int32(2)) % 32))
								v93 = v65
							}
						}
					}
					*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v93
					*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v91
					v101 = F_hash_search(m, v60, v8+int32(-48), int32(1), v8+int32(-49))
					mBase = m.M
					v102 = m.ExcPending
					if v102 != 0 {
						return int64(0)
					} else {
						v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
						if v103 == int32(0) {
							v106 = int32(_a_F_getmissingattr_1)
							v107 = *(*int32)(unsafe.Add(mBase, _c_F_getmissingattr[2]))
							v110 = *(*int32)(unsafe.Add(mBase, _c_F_getmissingattr[1]))
							*(*int32)(unsafe.Add(mBase, _c_F_getmissingattr[2])) = v110
							v112 = *(*int64)(unsafe.Add(mBase, uint32(v26)+8))
							v114 = int32(*(*int16)(unsafe.Add(mBase, uint32(v33)+2)))
							v115 = F_datumCopy(m, v112, int32(0), v114)
							mBase = m.M
							v116 = m.ExcPending
							if v116 != 0 {
								return int64(0)
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v101)+8)) = v115
								*(*int32)(unsafe.Add(mBase, _c_F_getmissingattr[2])) = v107
								v121 = *(*int64)(unsafe.Add(mBase, uint32(v101)+8))
								v132 = v121
								m.G0 = v10 - int32(-64)
								return v132
							}
						} else {
							v121 = *(*int64)(unsafe.Add(mBase, uint32(v101)+8))
							v132 = v121
							m.G0 = v10 - int32(-64)
							return v132
						}
					}
				}
			}
		}
	}
}
func F_getrusage(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(4)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(3)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(2)
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(1)
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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	v7 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l4))) = uint16(v7)
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v7)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v12 != 0 {
		*(*int32)(unsafe.Add(mBase, uint32(l3))) = v11
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v15 = F_strlen(m, v14)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v15
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v15 + v17
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v20 + int32(1)
		v27 = int32(2)
	} else {
		v27 = int32(0)
	}
	return v27
}
func F_ginbuildempty(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int64
	_ = v25
	var v27 int32
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
	var v70 int64
	_ = v70
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	v4 = m.G0
	v6 = v4 + int32(-64)
	m.G0 = v6
	*(*int64)(unsafe.Add(mBase, uint32(v6)+56)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+52)) = l0
	v11 = *(*int64)(unsafe.Add(mBase, uint32(v6)+52))
	*(*int64)(unsafe.Add(mBase, uint32(v6)+24)) = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v6)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = v13
	v20 = F_ExtendBufferedRel(m, v4+int32(-40), int32(3), int32(0), int32(9))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v6)+44)) = int64(0)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+40)) = l0
		v25 = *(*int64)(unsafe.Add(mBase, uint32(v6)+40))
		*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = v25
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v6)+48))
		*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v27
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
			v70 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v60)+64)) = v70
			*(*int64)(unsafe.Add(mBase, uint32(v60)+24)) = int64(-1)
			*(*int64)(unsafe.Add(mBase, uint32(v60)+32)) = v70
			*(*int64)(unsafe.Add(mBase, uint32(v60)+40)) = v70
			*(*int64)(unsafe.Add(mBase, uint32(v60)+48)) = v70
			*(*int32)(unsafe.Add(mBase, uint32(v60)+56)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v60)+72)) = int32(2)
			v84 = int32(80)
			*(*uint16)(unsafe.Add(mBase, uint32(v60)+12)) = uint16(v84)
			F_MarkBufferDirty(m, v20)
			mBase = m.M
			v87 = m.ExcPending
			if v87 != 0 {
				return
			} else {
				F_log_newpage_buffer(m, v20, int32(1))
				mBase = m.M
				v90 = m.ExcPending
				if v90 != 0 {
					return
				} else {
					v91 = int32(2)
					if v34 < int32(0) {
						v95 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuildempty[1]))
						v101 = *(*int32)(unsafe.Add(mBase, uint32(v95+(v34^int32(-1))<<(uint(int32(2))%32))))
						v109 = v101
					} else {
						v103 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuildempty[2]))
						v109 = v103 + v34<<(uint(int32(13))%32) + int32(-8192)
					}
					F_PageInit(m, v109, int32(_a_F_ginbuildempty_1), int32(8))
					mBase = m.M
					v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v109)+16)))
					v114 = v109 + v113
					*(*int32)(unsafe.Add(mBase, uint32(v114))) = int32(-1)
					*(*uint16)(unsafe.Add(mBase, uint32(v114)+6)) = uint16(v91)
					F_MarkBufferDirty(m, v34)
					mBase = m.M
					v119 = m.ExcPending
					if v119 != 0 {
						return
					} else {
						F_log_newpage_buffer(m, v34, int32(0))
						mBase = m.M
						v122 = m.ExcPending
						if v122 != 0 {
							return
						} else {
							v123 = int32(_a_F_ginbuildempty_0)
							v125 = *(*int32)(unsafe.Add(mBase, _c_F_ginbuildempty[0]))
							*(*int32)(unsafe.Add(mBase, _c_F_ginbuildempty[0])) = v125 - int32(1)
							F_UnlockReleaseBuffer(m, v20)
							mBase = m.M
							v130 = m.ExcPending
							if v130 != 0 {
								return
							} else {
								F_UnlockReleaseBuffer(m, v34)
								mBase = m.M
								v132 = m.ExcPending
								if v132 != 0 {
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
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v165 int32
	_ = v165
	var v182 int32
	_ = v182
	var v199 float64
	_ = v199
	var v200 int32
	_ = v200
	var v201 float64
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int64
	_ = v218
	var v228 float64
	_ = v228
	var v231 float64
	_ = v231
	var v234 int32
	_ = v234
	var v235 float64
	_ = v235
	var v237 int32
	_ = v237
	var v245 int32
	_ = v245
	var v248 int64
	_ = v248
	var v251 float64
	_ = v251
	var v252 int32
	_ = v252
	var v255 float64
	_ = v255
	var v256 float64
	_ = v256
	var v259 float64
	_ = v259
	var v261 float64
	_ = v261
	var v262 float64
	_ = v262
	var v264 float64
	_ = v264
	var v271 float64
	_ = v271
	var v274 float64
	_ = v274
	var v275 float64
	_ = v275
	var v278 float64
	_ = v278
	var v287 float64
	_ = v287
	var v288 float64
	_ = v288
	var v290 float64
	_ = v290
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v366 int32
	_ = v366
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v399 int32
	_ = v399
	var v423 float64
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v428 float64
	_ = v428
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v456 int32
	_ = v456
	var v471 int32
	_ = v471
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v513 int32
	_ = v513
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v556 float64
	_ = v556
	var v557 float64
	_ = v557
	var v560 float64
	_ = v560
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int64
	_ = v566
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
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
	var v595 float64
	_ = v595
	var v596 int32
	_ = v596
	var v597 float64
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
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v632 float64
	_ = v632
	var v633 int32
	_ = v633
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v659 float64
	_ = v659
	var v661 float64
	_ = v661
	var v665 float64
	_ = v665
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v677 int32
	_ = v677
	var v681 int64
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v696 float64
	_ = v696
	var v698 float64
	_ = v698
	var v700 float64
	_ = v700
	var v702 int32
	_ = v702
	var v704 float64
	_ = v704
	var v705 float64
	_ = v705
	var v706 float64
	_ = v706
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v713 float64
	_ = v713
	var v715 float64
	_ = v715
	var v719 float64
	_ = v719
	var v723 float64
	_ = v723
	var v726 float64
	_ = v726
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v770 int32
	_ = v770
	var v775 int32
	_ = v775
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v814 int32
	_ = v814
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v884 int32
	_ = v884
	var v899 int32
	_ = v899
	var v925 int32
	_ = v925
	var v929 int32
	_ = v929
	var v933 int32
	_ = v933
	var v970 float64
	_ = v970
	var v971 float64
	_ = v971
	var v972 int64
	_ = v972
	var v988 int32
	_ = v988
	var v1025 int32
	_ = v1025
	var v1039 float64
	_ = v1039
	var v1047 float64
	_ = v1047
	var v1048 int64
	_ = v1048
	var v1053 float64
	_ = v1053
	var v1058 float64
	_ = v1058
	var v1059 float64
	_ = v1059
	var v1062 float64
	_ = v1062
	var v1065 float64
	_ = v1065
	var v1066 float64
	_ = v1066
	var v1069 float64
	_ = v1069
	var v1074 float64
	_ = v1074
	var v1075 float64
	_ = v1075
	var v1077 float64
	_ = v1077
	var v1082 float64
	_ = v1082
	var v1086 float64
	_ = v1086
	var v1088 int32
	_ = v1088
	var v1089 float64
	_ = v1089
	var v1091 float64
	_ = v1091
	var v1098 float64
	_ = v1098
	var v1100 float64
	_ = v1100
	var v1104 float64
	_ = v1104
	var v1108 float64
	_ = v1108
	var v1110 float64
	_ = v1110
	var v1119 float64
	_ = v1119
	var v1121 float64
	_ = v1121
	var v1124 float64
	_ = v1124
	var v1128 int32
	_ = v1128
	var v1129 float64
	_ = v1129
	var v1130 float64
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1135 int32
	_ = v1135
	var v1138 int32
	_ = v1138
	var v1139 float64
	_ = v1139
	var v1140 float64
	_ = v1140
	var v1141 float64
	_ = v1141
	var v1143 int32
	_ = v1143
	var v1146 float64
	_ = v1146
	var v1147 float64
	_ = v1147
	var v1151 float64
	_ = v1151
	var v1152 float64
	_ = v1152
	var v1156 float64
	_ = v1156
	var v1160 float64
	_ = v1160
	var v1165 float64
	_ = v1165
	var v1175 float64
	_ = v1175
	var v1178 float64
	_ = v1178
	var v1183 float64
	_ = v1183
	var v1185 float64
	_ = v1185
	var v1186 int32
	_ = v1186
	var v1190 int32
	_ = v1190
	var v1193 int32
	_ = v1193
	var v1194 float64
	_ = v1194
	var v1195 float64
	_ = v1195
	var v1196 float64
	_ = v1196
	var v1198 int32
	_ = v1198
	var v1201 float64
	_ = v1201
	var v1202 float64
	_ = v1202
	var v1206 float64
	_ = v1206
	var v1207 float64
	_ = v1207
	var v1211 float64
	_ = v1211
	var v1215 float64
	_ = v1215
	var v1220 float64
	_ = v1220
	var v1230 float64
	_ = v1230
	var v1233 float64
	_ = v1233
	var v1238 float64
	_ = v1238
	var v1240 float64
	_ = v1240
	var v1242 float64
	_ = v1242
	var v1244 float64
	_ = v1244
	var v1246 float64
	_ = v1246
	var v1247 float64
	_ = v1247
	var v1249 float64
	_ = v1249
	var v1250 float64
	_ = v1250
	var v1252 float64
	_ = v1252
	var v1261 float64
	_ = v1261
	var v1264 float64
	_ = v1264
	var v1266 float64
	_ = v1266
	var v1271 float64
	_ = v1271
	var v1273 float64
	_ = v1273
	var v1274 float64
	_ = v1274
	var v1277 float64
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1282 int32
	_ = v1282
	var v1285 int32
	_ = v1285
	var v1286 float64
	_ = v1286
	var v1287 float64
	_ = v1287
	var v1288 float64
	_ = v1288
	var v1290 int32
	_ = v1290
	var v1293 float64
	_ = v1293
	var v1294 float64
	_ = v1294
	var v1298 float64
	_ = v1298
	var v1299 float64
	_ = v1299
	var v1303 float64
	_ = v1303
	var v1307 float64
	_ = v1307
	var v1312 float64
	_ = v1312
	var v1322 float64
	_ = v1322
	var v1325 float64
	_ = v1325
	var v1330 float64
	_ = v1330
	var v1332 float64
	_ = v1332
	var v1333 float64
	_ = v1333
	var v1334 float64
	_ = v1334
	var v1335 float64
	_ = v1335
	var v1336 float64
	_ = v1336
	var v1338 float64
	_ = v1338
	var v1342 float64
	_ = v1342
	var v1343 int32
	_ = v1343
	var v1345 float64
	_ = v1345
	var v1349 int32
	_ = v1349
	var v1351 float64
	_ = v1351
	var v1352 float64
	_ = v1352
	var v1358 float64
	_ = v1358
	var v1360 float64
	_ = v1360
	var v1362 float64
	_ = v1362
	var v1365 float64
	_ = v1365
	v9 = int32(0)
	v33 = m.G0
	v35 = v33 - int32(256)
	m.G0 = v35
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v38 == v9 {
		v182 = v9
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v199 = *(*float64)(unsafe.Add(mBase, uint32(v37)+24))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
	v201 = base.F64_convert_i32_u(v200)
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+105)))
	if v202 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L2:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	if v41 <= int32(0) {
		v182 = v9
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v54 = v9
	v55 = v41
	v59 = v9
	goto L4
L4:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v76+v54<<(uint(int32(2))%32))))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+8))
	if v81 == int32(0) {
		v143 = v55
		v147 = v59
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v182 = v147
	goto L1
L6:
	;
	v165 = v54 + int32(1)
	if v165 < v143 {
		v54 = v165
		v55 = v143
		v59 = v147
		goto L4
	} else {
		goto L14
	}
L7:
	;
	v84 = int32(0)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	if v85 <= v84 {
		v143 = v55
		v147 = v59
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v97 = v84
	v103 = v59
	goto L9
L9:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v81)+12))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v120+v97<<(uint(int32(2))%32))))
	v125 = F_lappend(m, v103, v124)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	v143 = v131
	v147 = v125
	goto L6
L11:
	;
	return
L12:
	;
	v128 = v97 + int32(1)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	if v128 < v129 {
		v97 = v128
		v103 = v125
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	goto L5
L15:
	;
	if base.F64_gt(v201, v228) != 0 {
		goto L22
	} else {
		goto L23
	}
L16:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v207 = F_index_open(m, v205, int32(0))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L11
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v218 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v35)+32)) = v218
	*(*int64)(unsafe.Add(mBase, uint32(v35)+24)) = v218
	*(*int64)(unsafe.Add(mBase, uint32(v35)+16)) = v218
	*(*int64)(unsafe.Add(mBase, uint32(v35)+8)) = v218
	v228 = float64(0)
	goto L15
L19:
	;
	F_ginGetStats(m, v207, v35+int32(8))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L11
	} else {
		goto L20
	}
L20:
	;
	F_relation_close(m, v207, int32(0))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L11
	} else {
		goto L21
	}
L21:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	v228 = base.F64_convert_i32_u(v216)
	goto L15
L22:
	;
	v231 = v228
	goto L24
L23:
	;
	v231 = float64(0)
	goto L24
L24:
	;
	if v200 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v37)+88))
	if v293 != 0 {
		goto L40
	} else {
		goto L41
	}
L26:
	;
	v271 = float64(10)
	if base.F64_gt(v201, v271) != 0 {
		goto L37
	} else {
		goto L38
	}
L27:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
	v235 = base.F64_convert_i32_u(v234)
	v237 = int32(0)
	if base.B2i32(base.F64_le(v235, v201) == v237)|base.B2i32(base.F64_lt(base.F64_mul(v201, float64(0.25)), v235) == v237) != 0 {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v35)+16))
	if v245 == int32(0) {
		goto L26
	} else {
		goto L29
	}
L29:
	;
	v248 = *(*int64)(unsafe.Add(mBase, uint32(v35)+24))
	if v248 <= int64(0) {
		goto L26
	} else {
		goto L30
	}
L30:
	;
	v251 = base.F64_div(v201, v235)
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
	v255 = base.F64_ceil(base.F64_mul(v251, base.F64_convert_i32_u(v252)))
	v256 = base.F64_sub(v201, v231)
	v259 = base.F64_ceil(base.F64_mul(v251, base.F64_convert_i32_u(v245)))
	if base.F64_gt(v256, v259) != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v261 = v259
	goto L33
L32:
	;
	v261 = v256
	goto L33
L33:
	;
	v262 = base.F64_sub(v256, v261)
	if base.F64_gt(v262, v255) != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v264 = v255
	goto L36
L35:
	;
	v264 = v262
	goto L36
L36:
	;
	v287 = v261
	v288 = v264
	v290 = base.F64_ceil(base.F64_mul(v251, base.F64_convert_i64_u(v248)))
	goto L25
L37:
	;
	v274 = v201
	goto L39
L38:
	;
	v274 = v271
	goto L39
L39:
	;
	v275 = base.F64_sub(v274, v231)
	v278 = base.F64_floor(base.F64_mul(v275, float64(0.9)))
	v287 = v278
	v288 = base.F64_sub(v275, v278)
	v290 = base.F64_floor(base.F64_mul(v278, float64(100)))
	goto L25
L40:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v293)+4))
	if v294 <= int32(0) {
		goto L44
	} else {
		goto L45
	}
L41:
	;
	v399 = v182
	goto L42
L42:
	;
	if base.F64_lt(v290, float64(1)) != 0 {
		goto L57
	} else {
		goto L58
	}
L43:
	;
	v388 = F_list_concat(m, v366, v182)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L11
	} else {
		goto L56
	}
L44:
	;
	v366 = int32(0)
	goto L43
L45:
	;
	goto L46
L46:
	;
	v298 = int32(0)
	v309 = v298
	v310 = v298
	goto L47
L47:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v293)+12))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v332+v309<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = v336
	*(*int32)(unsafe.Add(mBase, uint32(v35)+144)) = v336
	v342 = F_list_make1_impl(m, int32(1), v35+int32(4))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L11
	} else {
		goto L49
	}
L48:
	;
	v366 = v351
	goto L43
L49:
	;
	v345 = F_predicate_implied_by(m, v342, v182, int32(0))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L11
	} else {
		goto L50
	}
L50:
	;
	if v345 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v349 = F_list_concat(m, v310, v342)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L11
	} else {
		goto L54
	}
L52:
	;
	v351 = v310
	goto L53
L53:
	;
	v353 = v309 + int32(1)
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v293)+4))
	if v353 < v354 {
		v309 = v353
		v310 = v351
		goto L47
	} else {
		goto L55
	}
L54:
	;
	v351 = v349
	goto L53
L55:
	;
	goto L48
L56:
	;
	v399 = v388
	goto L42
L57:
	;
	v423 = float64(1)
	goto L59
L58:
	;
	v423 = v290
	goto L59
L59:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v424)+76))
	v426 = int32(0)
	v428 = F_clauselist_selectivity(m, l0, v399, v425, v426, v426)
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L11
	} else {
		goto L60
	}
L60:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l5))) = v428
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	F_get_tablespace_page_costs(m, v431, v35+int32(40), int32(0))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L11
	} else {
		goto L61
	}
L61:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l6))) = int64(0)
	v441 = int32(0)
	base.MemoryFill(m, v35+int32(48), v441, int32(88))
	*(*int64)(unsafe.Add(mBase, uint32(v35)+136)) = int64(4607182418800017408)
	v446 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v446 == v441 {
		goto L66
	} else {
		goto L67
	}
L62:
	;
	m.G0 = v35 + int32(256)
	return
L63:
	;
	v1048 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l3))) = v1048
	*(*int64)(unsafe.Add(mBase, uint32(l4))) = v1048
	v1053 = F_pow(m, v287, float64(0.15))
	mBase = m.M
	v1058 = base.F64_div(v1039, v423)
	v1059 = float64(1)
	if base.F64_lt(v1058, v1059) != 0 {
		goto L131
	} else {
		goto L132
	}
L64:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v35)+128)) = v423
	*(*float64)(unsafe.Add(mBase, uint32(v35)+120)) = v423
	*(*int64)(unsafe.Add(mBase, uint32(v35)+112)) = int64(0)
	v1025 = v988
	v1039 = float64(0)
	v1047 = v423
	goto L63
L65:
	;
	v972 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l3))) = v972
	*(*int64)(unsafe.Add(mBase, uint32(l4))) = v972
	*(*int64)(unsafe.Add(mBase, uint32(l5))) = v972
	goto L62
L66:
	;
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v37)+40))
	if int32(0) < v884 {
		goto L121
	} else {
		goto L122
	}
L67:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v446)+4))
	if v450 <= int32(0) {
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v456 = int32(1)
	v471 = v9
	goto L69
L69:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v446)+12))
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v487+v471<<(uint(int32(2))%32))))
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v491)+8))
	if v492 == int32(0) {
		v814 = v456
		goto L71
	} else {
		goto L72
	}
L70:
	;
	if v814 == int32(0) {
		goto L65
	} else {
		goto L120
	}
L71:
	;
	v846 = v471 + int32(1)
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v446)+4))
	if v846 < v847 {
		v456 = v814
		v471 = v846
		goto L69
	} else {
		goto L119
	}
L72:
	;
	v495 = int32(0)
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v492)+4))
	if v496 <= v495 {
		v814 = v456
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v513 = v495
	goto L74
L74:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v492)+12))
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v531+v513<<(uint(int32(2))%32))))
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v535)+4))
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v536)))
	switch v537 - int32(17) {
	case 0:
		goto L79
	default:
		goto L77
	case 3:
		goto L78
	}
L75:
	;
	v814 = v808
	goto L71
L76:
	;
	v808 = int32(1)
	v810 = v513 + v808
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v492)+4))
	if v810 < v811 {
		v513 = v810
		goto L74
	} else {
		goto L118
	}
L77:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L11
	} else {
		goto L115
	}
L78:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v536)+4))
	v572 = int32(*(*int16)(unsafe.Add(mBase, uint32(v491)+14)))
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v536)+28))
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v573)+12))
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v574)+4))
	v576 = F_estimate_expression_value(m, l0, v575)
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L11
	} else {
		goto L90
	}
L79:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v536)+4))
	v541 = int32(*(*int16)(unsafe.Add(mBase, uint32(v491)+14)))
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v536)+28))
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v542)+12))
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v543)+4))
	v545 = F_estimate_expression_value(m, l0, v544)
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L11
	} else {
		goto L80
	}
L80:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v545)))
	if v547 == int32(27) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v545)+4))
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v550)))
	v552 = v550
	v553 = v551
	goto L83
L82:
	;
	v552 = v545
	v553 = v547
	goto L83
L83:
	;
	if v553 != int32(7) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v556 = *(*float64)(unsafe.Add(mBase, uint32(v35)+120))
	v557 = float64(1)
	*(*float64)(unsafe.Add(mBase, uint32(v35)+120)) = base.F64_add(v556, v557)
	v560 = *(*float64)(unsafe.Add(mBase, uint32(v35)+128))
	*(*float64)(unsafe.Add(mBase, uint32(v35)+128)) = base.F64_add(v560, v557)
	goto L76
L85:
	;
	goto L86
L86:
	;
	v564 = int32(0)
	v565 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v552)+32)))
	if v565 != 0 {
		v814 = v564
		goto L71
	} else {
		goto L87
	}
L87:
	;
	v566 = *(*int64)(unsafe.Add(mBase, uint32(v552)+24))
	v569 = F_gincost_pattern(m, v37, v541, v540, v566, v35+int32(48))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L11
	} else {
		goto L88
	}
L88:
	;
	if v569 != 0 {
		goto L76
	} else {
		goto L89
	}
L89:
	;
	v814 = v564
	goto L71
L90:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v576)))
	if v578 == int32(27) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v576)+4))
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v581)))
	v583 = v581
	v584 = v582
	goto L93
L92:
	;
	v583 = v576
	v584 = v578
	goto L93
L93:
	;
	if v584 != int32(7) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v587 = *(*float64)(unsafe.Add(mBase, uint32(v35)+120))
	v588 = float64(1)
	*(*float64)(unsafe.Add(mBase, uint32(v35)+120)) = base.F64_add(v587, v588)
	v591 = *(*float64)(unsafe.Add(mBase, uint32(v35)+128))
	*(*float64)(unsafe.Add(mBase, uint32(v35)+128)) = base.F64_add(v591, v588)
	v595 = F_estimate_array_length(m, l0, v583)
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L11
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	v600 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v583)+32)))
	if v600 != 0 {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	v597 = *(*float64)(unsafe.Add(mBase, uint32(v35)+136))
	*(*float64)(unsafe.Add(mBase, uint32(v35)+136)) = base.F64_mul(v595, v597)
	goto L76
L98:
	;
	v814 = int32(0)
	goto L71
L99:
	;
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v583)+24))
	v602 = F_pg_detoast_datum(m, v601)
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L11
	} else {
		goto L100
	}
L100:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v602)+12))
	F_get_typlenbyvalalign(m, v604, v35+int32(254), v35+int32(253), v35+int32(252))
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L11
	} else {
		goto L101
	}
L101:
	;
	v614 = int32(*(*int16)(unsafe.Add(mBase, uint32(v35)+254)))
	v615 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+253)))
	v616 = int32(*(*int8)(unsafe.Add(mBase, uint32(v35)+252)))
	F_deconstruct_array(m, v602, v614, v615, v616, v35+int32(244), v35+int32(240), v35+int32(248))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L11
	} else {
		goto L102
	}
L102:
	;
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v35)+248))
	if v625 <= int32(0) {
		goto L98
	} else {
		goto L103
	}
L103:
	;
	v632 = float64(0)
	v633 = int32(0)
	v646 = v633
	v648 = v633
	v659 = v632
	v661 = v632
	v665 = v632
	goto L104
L104:
	;
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v35)+240))
	v671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v669+v646))))
	if v671 != 0 {
		v702 = v648
		v704 = v659
		v705 = v661
		v706 = v665
		goto L106
	} else {
		goto L107
	}
L105:
	;
	if v702 == int32(0) {
		goto L98
	} else {
		goto L114
	}
L106:
	;
	v708 = v646 + int32(1)
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v35)+248))
	if v708 < v709 {
		v646 = v708
		v648 = v702
		v659 = v704
		v661 = v705
		v665 = v706
		goto L104
	} else {
		goto L113
	}
L107:
	;
	v673 = v35 + int32(144)
	base.MemoryFill(m, v673, int32(0), int32(96))
	v677 = *(*int32)(unsafe.Add(mBase, uint32(v35)+244))
	v681 = *(*int64)(unsafe.Add(mBase, uint32(v677+v646<<(uint(int32(3))%32))))
	v682 = F_gincost_pattern(m, v37, v572, v571, v681, v673)
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L11
	} else {
		goto L108
	}
L108:
	;
	if v682 == int32(0) {
		v702 = v648
		v704 = v659
		v705 = v661
		v706 = v665
		goto L106
	} else {
		goto L109
	}
L109:
	;
	v686 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35+int32(144)+v572))))
	if v686 != int32(1) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v696 = *(*float64)(unsafe.Add(mBase, uint32(v35)+224))
	v698 = *(*float64)(unsafe.Add(mBase, uint32(v35)+216))
	v700 = *(*float64)(unsafe.Add(mBase, uint32(v35)+208))
	v702 = v648 + int32(1)
	v704 = base.F64_add(v659, v700)
	v705 = base.F64_add(v661, v696)
	v706 = base.F64_add(v665, v698)
	goto L106
L111:
	;
	v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v572+(v35+int32(176))))))
	if v689 != 0 {
		goto L110
	} else {
		goto L112
	}
L112:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v35)+224)) = v423
	*(*float64)(unsafe.Add(mBase, uint32(v35)+216)) = v423
	*(*int64)(unsafe.Add(mBase, uint32(v35)+208)) = int64(0)
	goto L110
L113:
	;
	goto L105
L114:
	;
	v713 = base.F64_convert_i32_s(v702)
	v715 = *(*float64)(unsafe.Add(mBase, uint32(v35)+112))
	*(*float64)(unsafe.Add(mBase, uint32(v35)+112)) = base.F64_add(base.F64_div(v704, v713), v715)
	v719 = *(*float64)(unsafe.Add(mBase, uint32(v35)+120))
	*(*float64)(unsafe.Add(mBase, uint32(v35)+120)) = base.F64_add(base.F64_div(v706, v713), v719)
	v723 = *(*float64)(unsafe.Add(mBase, uint32(v35)+128))
	*(*float64)(unsafe.Add(mBase, uint32(v35)+128)) = base.F64_add(base.F64_div(v705, v713), v723)
	v726 = *(*float64)(unsafe.Add(mBase, uint32(v35)+136))
	*(*float64)(unsafe.Add(mBase, uint32(v35)+136)) = base.F64_mul(v726, v713)
	goto L76
L115:
	;
	v766 = *(*int32)(unsafe.Add(mBase, uint32(v536)))
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = v766
	F_errmsg_internal(m, int32(_a_F_gincostestimate_0), v35)
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L11
	} else {
		goto L116
	}
L116:
	;
	F_errfinish(m, int32(_a_F_gincostestimate_1), int32(_a_F_gincostestimate_2), int32(_a_F_gincostestimate_3))
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L11
	} else {
		goto L117
	}
L117:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L118:
	;
	goto L75
L119:
	;
	goto L70
L120:
	;
	goto L66
L121:
	;
	v899 = int32(0)
	goto L124
L122:
	;
	goto L123
L123:
	;
	if v182 == int32(0) {
		v988 = int32(1)
		goto L64
	} else {
		goto L130
	}
L124:
	;
	v925 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35+int32(48)+v899))))
	if v925 != int32(1) {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	goto L123
L126:
	;
	v933 = v899 + int32(1)
	if v933 != v884 {
		v899 = v933
		goto L124
	} else {
		goto L129
	}
L127:
	;
	v929 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35+int32(80)+v899))))
	if v929 != 0 {
		goto L126
	} else {
		goto L128
	}
L128:
	;
	v988 = base.B2i32(v182 == int32(0))
	goto L64
L129:
	;
	goto L125
L130:
	;
	v970 = *(*float64)(unsafe.Add(mBase, uint32(v35)+112))
	v971 = *(*float64)(unsafe.Add(mBase, uint32(v35)+128))
	v1025 = int32(0)
	v1039 = v970
	v1047 = v971
	goto L63
L131:
	;
	v1062 = v1058
	goto L133
L132:
	;
	v1062 = v1059
	goto L133
L133:
	;
	v1065 = base.F64_add(base.F64_add(v231, base.F64_ceil(base.F64_mul(base.F64_nearest(v1053), v1047))), base.F64_ceil(base.F64_mul(v287, v1062)))
	v1066 = *(*float64)(unsafe.Add(mBase, uint32(v35)+136))
	if base.F64_gt(v423, float64(1)) != 0 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v1069 = F_log(m, v423)
	mBase = m.M
	v1074 = *(*float64)(unsafe.Add(mBase, _c_F_gincostestimate[0]))
	v1075 = base.F64_mul(base.F64_ceil(base.F64_div(v1069, float64(0.6931471805599453))), v1074)
	v1077 = *(*float64)(unsafe.Add(mBase, uint32(l3)))
	*(*float64)(unsafe.Add(mBase, uint32(l3))) = base.F64_add(base.F64_mul(v1075, v1047), v1077)
	v1082 = *(*float64)(unsafe.Add(mBase, uint32(l4)))
	*(*float64)(unsafe.Add(mBase, uint32(l4))) = base.F64_add(base.F64_mul(base.F64_mul(v1075, v1066), v1047), v1082)
	goto L136
L135:
	;
	goto L136
L136:
	;
	v1086 = float64(50)
	v1088 = int32(_a_F_gincostestimate_4)
	v1089 = *(*float64)(unsafe.Add(mBase, _c_F_gincostestimate[0]))
	v1091 = *(*float64)(unsafe.Add(mBase, uint32(l3)))
	*(*float64)(unsafe.Add(mBase, uint32(l3))) = base.F64_add(base.F64_mul(base.F64_mul(v1065, v1086), v1089), v1091)
	v1098 = *(*float64)(unsafe.Add(mBase, _c_F_gincostestimate[0]))
	v1100 = *(*float64)(unsafe.Add(mBase, uint32(l4)))
	*(*float64)(unsafe.Add(mBase, uint32(l4))) = base.F64_add(base.F64_mul(base.F64_mul(base.F64_mul(v1065, v1066), v1086), v1098), v1100)
	v1104 = *(*float64)(unsafe.Add(mBase, _c_F_gincostestimate[0]))
	v1108 = base.F64_ceil(base.F64_mul(v288, v1062))
	v1110 = *(*float64)(unsafe.Add(mBase, uint32(l3)))
	*(*float64)(unsafe.Add(mBase, uint32(l3))) = base.F64_add(base.F64_mul(base.F64_mul(v1104, v1086), v1108), v1110)
	v1119 = *(*float64)(unsafe.Add(mBase, _c_F_gincostestimate[0]))
	v1121 = *(*float64)(unsafe.Add(mBase, uint32(l4)))
	*(*float64)(unsafe.Add(mBase, uint32(l4))) = base.F64_add(base.F64_mul(base.F64_mul(base.F64_mul(v1108, base.F64_add(v1066, float64(-1))), v1086), v1119), v1121)
	v1124 = float64(1)
	v1128 = base.F64_gt(l2, v1124) | base.F64_gt(v1066, v1124)
	if v1128 != 0 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v1129 = base.F64_mul(l2, v1066)
	v1130 = base.F64_mul(v1065, v1129)
	v1131 = base.I32_trunc_sat_f64_u(v287)
	v1135 = int32(1)
	if base.Ui32(v1131) <= base.Ui32(v1135) {
		goto L141
	} else {
		goto L142
	}
L138:
	;
	v1240 = v1065
	v1242 = v1108
	goto L139
L139:
	;
	v1244 = *(*float64)(unsafe.Add(mBase, uint32(v35)+40))
	v1246 = *(*float64)(unsafe.Add(mBase, uint32(l3)))
	v1247 = base.F64_add(base.F64_mul(base.F64_add(v1242, v1240), v1244), v1246)
	*(*float64)(unsafe.Add(mBase, uint32(l3))) = v1247
	v1249 = *(*float64)(unsafe.Add(mBase, uint32(l5)))
	v1250 = *(*float64)(unsafe.Add(mBase, uint32(v35)+120))
	v1252 = *(*float64)(unsafe.Add(mBase, _c_F_gincostestimate[0]))
	*(*float64)(unsafe.Add(mBase, uint32(l3))) = base.F64_add(base.F64_mul(base.F64_mul(v1252, float64(50)), v1047), v1247)
	v1261 = base.F64_ceil(base.F64_mul(v1249, base.F64_div(v199, float64(2730))))
	v1264 = base.F64_ceil(base.F64_div(base.F64_mul(v288, v1250), v423))
	if base.F64_gt(v1261, v1264) != 0 {
		goto L176
	} else {
		goto L177
	}
L140:
	;
	v1185 = base.F64_mul(v1108, v1129)
	v1186 = base.I32_trunc_sat_f64_u(v288)
	v1190 = int32(1)
	if base.Ui32(v1186) <= base.Ui32(v1190) {
		goto L159
	} else {
		goto L160
	}
L141:
	;
	v1138 = v1135
	goto L143
L142:
	;
	v1138 = v1131
	goto L143
L143:
	;
	v1139 = base.F64_convert_i32_u(v1138)
	v1140 = base.F64_add(v1139, v1139)
	v1141 = float64(1)
	v1143 = *(*int32)(unsafe.Add(mBase, _c_F_gincostestimate[1]))
	v1146 = *(*float64)(unsafe.Add(mBase, uint32(l0)+304))
	v1147 = base.F64_add(v287, v1146)
	if base.F64_gt(v1147, v1141) != 0 {
		goto L145
	} else {
		goto L146
	}
L144:
	;
	v1183 = v1178
	goto L140
L145:
	;
	v1151 = v1147
	goto L147
L146:
	;
	v1151 = v1141
	goto L147
L147:
	;
	v1152 = base.F64_div(base.F64_mul(v1139, base.F64_convert_i32_s(v1143)), v1151)
	if base.F64_le(v1152, float64(1)) != 0 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v1156 = v1141
	goto L150
L149:
	;
	v1156 = base.F64_ceil(v1152)
	goto L150
L150:
	;
	if base.F64_le(v1139, v1156) != 0 {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v1160 = base.F64_div(base.F64_mul(v1130, v1140), base.F64_add(v1140, v1130))
	if base.F64_ge(v1160, v1139) != 0 {
		v1178 = v1139
		goto L144
	} else {
		goto L154
	}
L152:
	;
	goto L153
L153:
	;
	v1165 = base.F64_div(base.F64_mul(v1140, v1156), base.F64_sub(v1140, v1156))
	if base.F64_ge(v1165, v1130) != 0 {
		goto L155
	} else {
		goto L156
	}
L154:
	;
	v1183 = base.F64_ceil(v1160)
	goto L140
L155:
	;
	v1175 = base.F64_div(base.F64_mul(v1130, v1140), base.F64_add(v1140, v1130))
	goto L157
L156:
	;
	v1175 = base.F64_add(v1156, base.F64_div(base.F64_mul(base.F64_sub(v1139, v1156), base.F64_sub(v1130, v1165)), v1139))
	goto L157
L157:
	;
	v1178 = base.F64_ceil(v1175)
	goto L144
L158:
	;
	v1240 = base.F64_div(v1183, l2)
	v1242 = base.F64_div(v1238, l2)
	goto L139
L159:
	;
	v1193 = v1190
	goto L161
L160:
	;
	v1193 = v1186
	goto L161
L161:
	;
	v1194 = base.F64_convert_i32_u(v1193)
	v1195 = base.F64_add(v1194, v1194)
	v1196 = float64(1)
	v1198 = *(*int32)(unsafe.Add(mBase, _c_F_gincostestimate[1]))
	v1201 = *(*float64)(unsafe.Add(mBase, uint32(l0)+304))
	v1202 = base.F64_add(v288, v1201)
	if base.F64_gt(v1202, v1196) != 0 {
		goto L163
	} else {
		goto L164
	}
L162:
	;
	v1238 = v1233
	goto L158
L163:
	;
	v1206 = v1202
	goto L165
L164:
	;
	v1206 = v1196
	goto L165
L165:
	;
	v1207 = base.F64_div(base.F64_mul(v1194, base.F64_convert_i32_s(v1198)), v1206)
	if base.F64_le(v1207, float64(1)) != 0 {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v1211 = v1196
	goto L168
L167:
	;
	v1211 = base.F64_ceil(v1207)
	goto L168
L168:
	;
	if base.F64_le(v1194, v1211) != 0 {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v1215 = base.F64_div(base.F64_mul(v1185, v1195), base.F64_add(v1195, v1185))
	if base.F64_ge(v1215, v1194) != 0 {
		v1233 = v1194
		goto L162
	} else {
		goto L172
	}
L170:
	;
	goto L171
L171:
	;
	v1220 = base.F64_div(base.F64_mul(v1195, v1211), base.F64_sub(v1195, v1211))
	if base.F64_ge(v1220, v1185) != 0 {
		goto L173
	} else {
		goto L174
	}
L172:
	;
	v1238 = base.F64_ceil(v1215)
	goto L158
L173:
	;
	v1230 = base.F64_div(base.F64_mul(v1185, v1195), base.F64_add(v1195, v1185))
	goto L175
L174:
	;
	v1230 = base.F64_add(v1211, base.F64_div(base.F64_mul(base.F64_sub(v1194, v1211), base.F64_sub(v1185, v1220)), v1194))
	goto L175
L175:
	;
	v1233 = base.F64_ceil(v1230)
	goto L162
L176:
	;
	v1266 = v1261
	goto L178
L177:
	;
	v1266 = v1264
	goto L178
L178:
	;
	v1271 = *(*float64)(unsafe.Add(mBase, _c_F_gincostestimate[0]))
	v1273 = *(*float64)(unsafe.Add(mBase, uint32(l4)))
	v1274 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_mul(v1066, v1266), float64(50)), v1271), v1273)
	*(*float64)(unsafe.Add(mBase, uint32(l4))) = v1274
	if v1128 != 0 {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v1277 = base.F64_mul(base.F64_mul(l2, v1066), v1266)
	v1278 = base.I32_trunc_sat_f64_u(v288)
	v1282 = int32(1)
	if base.Ui32(v1278) <= base.Ui32(v1282) {
		goto L183
	} else {
		goto L184
	}
L180:
	;
	v1334 = v1266
	v1335 = v1244
	v1336 = v1274
	goto L181
L181:
	;
	v1338 = *(*float64)(unsafe.Add(mBase, uint32(l3)))
	*(*float64)(unsafe.Add(mBase, uint32(l4))) = base.F64_add(v1336, base.F64_add(base.F64_mul(v1334, v1335), v1338))
	v1342 = F_index_other_operands_eval_cost(m, l0, v182)
	mBase = m.M
	v1343 = m.ExcPending
	if v1343 != 0 {
		goto L11
	} else {
		goto L200
	}
L182:
	;
	v1332 = *(*float64)(unsafe.Add(mBase, uint32(v35)+40))
	v1333 = *(*float64)(unsafe.Add(mBase, uint32(l4)))
	v1334 = base.F64_div(v1330, l2)
	v1335 = v1332
	v1336 = v1333
	goto L181
L183:
	;
	v1285 = v1282
	goto L185
L184:
	;
	v1285 = v1278
	goto L185
L185:
	;
	v1286 = base.F64_convert_i32_u(v1285)
	v1287 = base.F64_add(v1286, v1286)
	v1288 = float64(1)
	v1290 = *(*int32)(unsafe.Add(mBase, _c_F_gincostestimate[1]))
	v1293 = *(*float64)(unsafe.Add(mBase, uint32(l0)+304))
	v1294 = base.F64_add(v288, v1293)
	if base.F64_gt(v1294, v1288) != 0 {
		goto L187
	} else {
		goto L188
	}
L186:
	;
	v1330 = v1325
	goto L182
L187:
	;
	v1298 = v1294
	goto L189
L188:
	;
	v1298 = v1288
	goto L189
L189:
	;
	v1299 = base.F64_div(base.F64_mul(v1286, base.F64_convert_i32_s(v1290)), v1298)
	if base.F64_le(v1299, float64(1)) != 0 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v1303 = v1288
	goto L192
L191:
	;
	v1303 = base.F64_ceil(v1299)
	goto L192
L192:
	;
	if base.F64_le(v1286, v1303) != 0 {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v1307 = base.F64_div(base.F64_mul(v1277, v1287), base.F64_add(v1287, v1277))
	if base.F64_ge(v1307, v1286) != 0 {
		v1325 = v1286
		goto L186
	} else {
		goto L196
	}
L194:
	;
	goto L195
L195:
	;
	v1312 = base.F64_div(base.F64_mul(v1287, v1303), base.F64_sub(v1287, v1303))
	if base.F64_ge(v1312, v1277) != 0 {
		goto L197
	} else {
		goto L198
	}
L196:
	;
	v1330 = base.F64_ceil(v1307)
	goto L182
L197:
	;
	v1322 = base.F64_div(base.F64_mul(v1277, v1287), base.F64_add(v1287, v1277))
	goto L199
L198:
	;
	v1322 = base.F64_add(v1303, base.F64_div(base.F64_mul(base.F64_sub(v1286, v1303), base.F64_sub(v1277, v1312)), v1286))
	goto L199
L199:
	;
	v1325 = base.F64_ceil(v1322)
	goto L186
L200:
	;
	v1345 = *(*float64)(unsafe.Add(mBase, _c_F_gincostestimate[0]))
	if v1025 == int32(0) {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v1349 = *(*int32)(unsafe.Add(mBase, uint32(v182)+4))
	v1351 = base.F64_convert_i32_s(v1349)
	goto L203
L202:
	;
	v1351 = float64(0)
	goto L203
L203:
	;
	v1352 = *(*float64)(unsafe.Add(mBase, uint32(l3)))
	*(*float64)(unsafe.Add(mBase, uint32(l3))) = base.F64_add(v1342, v1352)
	v1358 = *(*float64)(unsafe.Add(mBase, uint32(l4)))
	v1360 = base.F64_add(base.F64_mul(base.F64_mul(v1047, v1066), base.F64_mul(v1345, v1351)), base.F64_add(v1342, v1358))
	*(*float64)(unsafe.Add(mBase, uint32(l4))) = v1360
	v1362 = *(*float64)(unsafe.Add(mBase, uint32(l5)))
	v1365 = *(*float64)(unsafe.Add(mBase, _c_F_gincostestimate[2]))
	*(*float64)(unsafe.Add(mBase, uint32(l4))) = base.F64_add(base.F64_mul(base.F64_mul(v199, v1362), v1365), v1360)
	*(*float64)(unsafe.Add(mBase, uint32(l7))) = v1334
	goto L62
}
func F_gininsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
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
	var v90 int64
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
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v130 int64
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int64
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v181 int32
	_ = v181
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v199 int64
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v270 int32
	_ = v270
	var v289 int64
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v414 int64
	_ = v414
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v502 int32
	_ = v502
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
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
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v598 int32
	_ = v598
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v658 int32
	_ = v658
	var v662 int32
	_ = v662
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v683 int32
	_ = v683
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v693 int64
	_ = v693
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v714 int32
	_ = v714
	var v718 int32
	_ = v718
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v753 int64
	_ = v753
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v767 int32
	_ = v767
	var v769 int32
	_ = v769
	var v771 int32
	_ = v771
	var v781 int32
	_ = v781
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v795 int32
	_ = v795
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v835 int32
	_ = v835
	var v843 int32
	_ = v843
	var v845 int32
	_ = v845
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v854 int32
	_ = v854
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v867 int32
	_ = v867
	var v876 int32
	_ = v876
	var v879 int32
	_ = v879
	var v880 int64
	_ = v880
	var v882 int64
	_ = v882
	var v884 int64
	_ = v884
	var v886 int64
	_ = v886
	var v888 int64
	_ = v888
	var v890 int64
	_ = v890
	var v892 int64
	_ = v892
	var v897 int32
	_ = v897
	var v902 int32
	_ = v902
	var v905 int64
	_ = v905
	var v906 int32
	_ = v906
	var v908 int64
	_ = v908
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v919 int32
	_ = v919
	var v923 int32
	_ = v923
	var v941 int32
	_ = v941
	var v944 int32
	_ = v944
	var v949 int32
	_ = v949
	var v968 int32
	_ = v968
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1000 int32
	_ = v1000
	var v1005 int32
	_ = v1005
	var v1011 int32
	_ = v1011
	var v1016 int32
	_ = v1016
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1050 int32
	_ = v1050
	var v1055 int32
	_ = v1055
	var v1079 int32
	_ = v1079
	v9 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(32)
	m.G0 = v23
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_gininsert[0]))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l7)+136))
	if v27 == v9 {
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
	v1079 = m.ExcPending
	if v1079 != 0 {
		goto L4
	} else {
		goto L202
	}
L9:
	;
	v165 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+16)) = v165
	*(*int64)(unsafe.Add(mBase, uint32(v23)+8)) = v165
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v169)))
	if int32(0) < v170 {
		goto L28
	} else {
		goto L29
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
	v90 = *(*int64)(unsafe.Add(mBase, uint32(l1+v70<<(uint(int32(3))%32))))
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
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
	if v84 < v163 {
		v70 = v84
		goto L13
	} else {
		goto L27
	}
L19:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_gininsert[1]))
	if v124 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L18
L21:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L4
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v130 = *(*int64)(unsafe.Add(mBase, uint32(v97+v110<<(uint(int32(3))%32))))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v133 = int32(*(*int8)(unsafe.Add(mBase, uint32(v131+v110))))
	F_ginEntryInsert(m, v43, v86, v130, v133, l3, int32(1), int32(0))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L4
	} else {
		goto L25
	}
L24:
	;
	goto L23
L25:
	;
	v139 = v110 + int32(1)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
	if v139 < v140 {
		v110 = v139
		goto L19
	} else {
		goto L26
	}
L26:
	;
	goto L20
L27:
	;
	goto L14
L28:
	;
	v181 = int32(0)
	goto L31
L29:
	;
	goto L30
L30:
	;
	v382 = int32(0)
	v386 = m.G0
	v388 = v386 - int32(96)
	m.G0 = v388
	v391 = v23 + int32(8)
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v391)+4))
	if v392 == v382 {
		goto L68
	} else {
		goto L69
	}
L31:
	;
	v194 = int32(8)
	v195 = v23 + v194
	v199 = *(*int64)(unsafe.Add(mBase, uint32(l1+v181<<(uint(int32(3))%32))))
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v181))))
	v202 = m.G0
	v204 = v202 - int32(16)
	m.G0 = v204
	v207 = v181 + int32(1)
	v209 = v207 & int32(_a_F_gininsert_5)
	v214 = F_ginExtractEntries(m, v43, v209, v199, v201, v204+v194, v204+int32(12))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L4
	} else {
		goto L33
	}
L32:
	;
	goto L30
L33:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v204)+8))
	if v216 < int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v359)))
	if v207 < v360 {
		v181 = v207
		goto L31
	} else {
		goto L65
	}
L35:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L4
	} else {
		goto L62
	}
L36:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v195)+4))
	v220 = v219 + v216
	if base.Ui32(int32(268435456)) <= base.Ui32(v220) {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
	if v223 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v204)+8))
	if int32(0) < v262 {
		goto L55
	} else {
		goto L56
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v195))) = v259
	goto L38
L40:
	;
	v228 = int32(16)
	if base.Ui32(v216) <= base.Ui32(v228) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	goto L42
L42:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v195)+8))
	if base.Ui32(v220) <= base.Ui32(v243) {
		goto L38
	} else {
		goto L50
	}
L43:
	;
	v231 = v228
	goto L45
L44:
	;
	v231 = v216
	goto L45
L45:
	;
	if v231&(v231-int32(1)) != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v238 = int32(1) << (uint(int32(32)-base.I32_clz(v231)) % 32)
	goto L48
L47:
	;
	v238 = v231
	goto L48
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v195)+8)) = v238
	v241 = F_palloc_mul(m, int32(4), v238)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	v259 = v241
	goto L39
L50:
	;
	v245 = int32(1)
	if v220&(v220-v245) != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v253 = v245 << (uint(int32(32)-base.I32_clz(v220)) % 32)
	goto L53
L52:
	;
	v253 = v220
	goto L53
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v195)+8)) = v253
	v256 = F_repalloc_mul(m, v223, int32(4), v253)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	v259 = v256
	goto L39
L55:
	;
	v270 = int32(0)
	goto L58
L56:
	;
	goto L57
L57:
	;
	m.G0 = v204 + int32(16)
	goto L34
L58:
	;
	v289 = *(*int64)(unsafe.Add(mBase, uint32(v214+v270<<(uint(int32(3))%32))))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v204)+12))
	v292 = int32(*(*int8)(unsafe.Add(mBase, uint32(v290+v270))))
	v293 = int32(0)
	v297 = F_GinFormTuple(m, v43, v209, v289, v292, v293, v293, v293, int32(1))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L4
	} else {
		goto L60
	}
L59:
	;
	goto L57
L60:
	;
	v299 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v297)+4)) = uint16(v299)
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v297))) = v301
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v195)+4))
	v304 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v195)+4)) = v303 + v304
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
	*(*int32)(unsafe.Add(mBase, uint32(v307+v303<<(uint(int32(2))%32)))) = v297
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v195)+12))
	v313 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v297)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v195)+12)) = v312 + v313&int32(_a_F_gininsert_6)
	v319 = v270 + v304
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v204)+8))
	if v319 < v320 {
		v270 = v319
		goto L58
	} else {
		goto L61
	}
L61:
	;
	goto L59
L62:
	;
	F_errmsg_internal(m, int32(_a_F_gininsert_7), int32(0))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L4
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_gininsert_8), int32(504), int32(_a_F_gininsert_9))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L4
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L65:
	;
	goto L32
L66:
	;
	goto L8
L67:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1043 = m.ExcPending
	if v1043 != 0 {
		goto L4
	} else {
		goto L199
	}
L68:
	;
	m.G0 = v388 + int32(96)
	goto L66
L69:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v396)+48))
	v398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v397)+118)))
	if v398 != int32(112) {
		v411 = int32(0)
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v396)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v388)+16)) = v412
	v414 = *(*int64)(unsafe.Add(mBase, uint32(v396)))
	*(*int64)(unsafe.Add(mBase, uint32(v388)+8)) = v414
	v416 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v388)+88)) = v416
	*(*int64)(unsafe.Add(mBase, uint32(v388)+80)) = int64(-1)
	v421 = F_ReadBuffer(m, v396, v416)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L4
	} else {
		goto L75
	}
L71:
	;
	v403 = *(*int32)(unsafe.Add(mBase, _c_F_gininsert[2]))
	if int32(0) < v403 {
		v411 = int32(1)
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v396)+32))
	if v407 != 0 {
		v411 = int32(0)
		goto L70
	} else {
		goto L73
	}
L73:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v396)+40))
	v411 = base.B2i32(v408 == int32(0))
	goto L70
L74:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v391)+12))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v391)+4))
	if base.Ui32(v441+v442<<(uint(int32(2))%32)) <= base.Ui32(int32(_a_F_gininsert_10)) {
		goto L85
	} else {
		goto L86
	}
L75:
	;
	if v421 < int32(0) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v426 = *(*int32)(unsafe.Add(mBase, _c_F_gininsert[3]))
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v426+(v421^int32(-1))<<(uint(int32(2))%32))))
	v440 = v432
	goto L74
L77:
	;
	goto L78
L78:
	;
	v434 = *(*int32)(unsafe.Add(mBase, _c_F_gininsert[4]))
	v440 = v434 + v421<<(uint(int32(13))%32) + int32(-8192)
	goto L74
L79:
	;
	v989 = *(*int32)(unsafe.Add(mBase, uint32(v396)+180))
	if v989 != 0 {
		goto L192
	} else {
		goto L193
	}
L80:
	;
	F_UnlockReleaseBuffer(m, v949)
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L4
	} else {
		goto L190
	}
L81:
	;
	v941 = int32(80)
	*(*uint16)(unsafe.Add(mBase, uint32(v440)+12)) = uint16(v941)
	F_MarkBufferDirty(m, v421)
	mBase = m.M
	v944 = m.ExcPending
	if v944 != 0 {
		goto L4
	} else {
		goto L188
	}
L82:
	;
	v913 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v732)+14)))
	v914 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v732)+12)))
	v915 = v913 - v914
	v916 = int32(0)
	if v916 < v915 {
		goto L185
	} else {
		goto L186
	}
L83:
	;
	v876 = int32(80)
	*(*uint16)(unsafe.Add(mBase, uint32(v440)+12)) = uint16(v876)
	F_MarkBufferDirty(m, v421)
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L4
	} else {
		goto L179
	}
L84:
	;
	v705 = int32(0)
	F_CheckForSerializableConflictIn(m, v396, v705, v705)
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L4
	} else {
		goto L145
	}
L85:
	;
	F_LockBufferInternal(m, v421, int32(3))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L4
	} else {
		goto L88
	}
L86:
	;
	v464 = v442
	goto L87
L87:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	v467 = int32(0)
	if v467 < v464 {
		goto L94
	} else {
		goto L95
	}
L88:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v440)+24))
	if v451 != int32(-1) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v440)+32))
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v391)+12))
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v391)+4))
	if base.Ui32(v455+v456<<(uint(int32(2))%32)) <= base.Ui32(v454) {
		goto L84
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	F_UnlockBuffer(m, v421)
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L4
	} else {
		goto L93
	}
L92:
	;
	goto L91
L93:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v391)+4))
	v464 = v463
	goto L87
L94:
	;
	v471 = v467
	v472 = v382
	v473 = v382
	v475 = v382
	v478 = v382
	v484 = int32(0)
	v487 = v9
	goto L97
L95:
	;
	v575 = v467
	v577 = v382
	v588 = int32(1)
	v591 = v9
	goto L96
L96:
	;
	if v575 < int32(0) {
		goto L124
	} else {
		goto L125
	}
L97:
	;
	if v471 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	v575 = v570
	v577 = v547
	v588 = v548 + int32(1)
	v591 = v551
	goto L96
L99:
	;
	v493 = F_GinNewBuffer(m, v396)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L4
	} else {
		goto L102
	}
L100:
	;
	v546 = v471
	v547 = v473
	v548 = v475
	v549 = v478
	v550 = v484
	v551 = v487
	goto L101
L101:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v466+v472<<(uint(int32(2))%32))))
	v556 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v555)+6)))
	v565 = v550 + (v556&int32(_a_F_gininsert_6)+int32(7))&int32(_a_F_gininsert_11) + int32(4)
	v567 = base.B2i32(base.Ui32(v565) < base.Ui32(int32(_a_F_gininsert_12)))
	if base.Ui32(v565) < base.Ui32(int32(_a_F_gininsert_12)) {
		goto L116
	} else {
		goto L117
	}
L102:
	;
	if v478 != 0 {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	v546 = v493
	v547 = v472
	v548 = v543
	v549 = v493
	v550 = v544
	v551 = v545
	goto L101
L104:
	;
	if v493 < int32(0) {
		goto L108
	} else {
		goto L109
	}
L105:
	;
	goto L106
L106:
	;
	v523 = int32(0)
	if v493 < v523 {
		goto L113
	} else {
		goto L114
	}
L107:
	;
	v518 = F_writeListPage(m, v396, v478, v466+v473<<(uint(int32(2))%32), v472-v473, v517)
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L4
	} else {
		goto L111
	}
L108:
	;
	v502 = *(*int32)(unsafe.Add(mBase, _c_F_gininsert[5]))
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v502+(v493^int32(-1))*int32(56))+16))
	v517 = v508
	goto L107
L109:
	;
	goto L110
L110:
	;
	v510 = *(*int32)(unsafe.Add(mBase, _c_F_gininsert[6]))
	v511 = int32(56)
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v510+v493*v511-v511)+16))
	v517 = v516
	goto L107
L111:
	;
	v543 = v475 + int32(1)
	v544 = int32(0)
	v545 = v487
	goto L103
L112:
	;
	v543 = v475
	v544 = v523
	v545 = v542
	goto L103
L113:
	;
	v527 = *(*int32)(unsafe.Add(mBase, _c_F_gininsert[5]))
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v527+(v493^int32(-1))*int32(56))+16))
	v542 = v533
	goto L112
L114:
	;
	goto L115
L115:
	;
	v535 = *(*int32)(unsafe.Add(mBase, _c_F_gininsert[6]))
	v536 = int32(56)
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v535+v493*v536-v536)+16))
	v542 = v541
	goto L112
L116:
	;
	v568 = v565
	goto L118
L117:
	;
	v568 = v550
	goto L118
L118:
	;
	if base.Ui32(v565) < base.Ui32(int32(_a_F_gininsert_12)) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v570 = v546
	goto L121
L120:
	;
	v570 = int32(0)
	goto L121
L121:
	;
	v571 = v472 + v567
	if v571 < v464 {
		v471 = v570
		v472 = v571
		v473 = v547
		v475 = v548
		v478 = v549
		v484 = v568
		v487 = v551
		goto L97
	} else {
		goto L122
	}
L122:
	;
	goto L98
L123:
	;
	v619 = F_writeListPage(m, v396, v575, v466+v577<<(uint(int32(2))%32), v464-v577, int32(-1))
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L4
	} else {
		goto L127
	}
L124:
	;
	v598 = *(*int32)(unsafe.Add(mBase, _c_F_gininsert[5]))
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v598+(v575^int32(-1))*int32(56))+16))
	v613 = v604
	goto L123
L125:
	;
	goto L126
L126:
	;
	v606 = *(*int32)(unsafe.Add(mBase, _c_F_gininsert[6]))
	v607 = int32(56)
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v606+v575*v607-v607)+16))
	v613 = v612
	goto L123
L127:
	;
	F_LockBufferInternal(m, v421, int32(3))
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L4
	} else {
		goto L128
	}
L128:
	;
	v624 = int32(0)
	F_CheckForSerializableConflictIn(m, v396, v624, v624)
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L4
	} else {
		goto L129
	}
L129:
	;
	v629 = v440 + int32(24)
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v440)+24))
	if v630 == int32(-1) {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v633 = int32(0)
	v634 = int32(_a_F_gininsert_13)
	v636 = *(*int32)(unsafe.Add(mBase, _c_F_gininsert[7]))
	*(*int32)(unsafe.Add(mBase, _c_F_gininsert[7])) = v636 + int32(1)
	*(*int64)(unsafe.Add(mBase, uint32(v440)+40)) = int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v440)+36)) = v588
	*(*int32)(unsafe.Add(mBase, uint32(v440)+32)) = v619
	*(*int32)(unsafe.Add(mBase, uint32(v440)+28)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v440)+24)) = v591
	if v411 == v633 {
		v923 = v633
		goto L81
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v440)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v388)+84)) = v591
	*(*int32)(unsafe.Add(mBase, uint32(v388)+80)) = v651
	v654 = F_ReadBuffer(m, v396, v651)
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L4
	} else {
		goto L135
	}
L133:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L4
	} else {
		goto L134
	}
L134:
	;
	v857 = v629
	v858 = v633
	v867 = int32(0)
	goto L83
L135:
	;
	F_LockBufferInternal(m, v654, int32(3))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L4
	} else {
		goto L136
	}
L136:
	;
	if v654 < int32(0) {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	v677 = int32(_a_F_gininsert_13)
	v679 = *(*int32)(unsafe.Add(mBase, _c_F_gininsert[7]))
	*(*int32)(unsafe.Add(mBase, _c_F_gininsert[7])) = v679 + int32(1)
	v683 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v676)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v676+v683))) = v591
	F_MarkBufferDirty(m, v654)
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L4
	} else {
		goto L141
	}
L138:
	;
	v662 = *(*int32)(unsafe.Add(mBase, _c_F_gininsert[3]))
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v662+(v654^int32(-1))<<(uint(int32(2))%32))))
	v676 = v668
	goto L137
L139:
	;
	goto L140
L140:
	;
	v670 = *(*int32)(unsafe.Add(mBase, _c_F_gininsert[4]))
	v676 = v670 + v654<<(uint(int32(13))%32) + int32(-8192)
	goto L137
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v440)+32)) = v619
	*(*int32)(unsafe.Add(mBase, uint32(v440)+28)) = v613
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v440)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v440)+36)) = v690 + v588
	v693 = *(*int64)(unsafe.Add(mBase, uint32(v440)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v440)+40)) = v693 + int64(1)
	if v411 == int32(0) {
		v923 = v654
		goto L81
	} else {
		goto L142
	}
L142:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L4
	} else {
		goto L143
	}
L143:
	;
	F_XLogRegisterBuffer(m, int32(1), v654, int32(8))
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L4
	} else {
		goto L144
	}
L144:
	;
	v857 = v629
	v858 = v654
	v867 = v676
	goto L83
L145:
	;
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v440)+28))
	v710 = F_ReadBuffer(m, v396, v709)
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L4
	} else {
		goto L146
	}
L146:
	;
	F_LockBufferInternal(m, v710, int32(3))
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L4
	} else {
		goto L147
	}
L147:
	;
	if v710 < int32(0) {
		goto L149
	} else {
		goto L150
	}
L148:
	;
	v733 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v732)+12)))
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v391)+12))
	v735 = F_palloc(m, v734)
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L4
	} else {
		goto L152
	}
L149:
	;
	v718 = *(*int32)(unsafe.Add(mBase, _c_F_gininsert[3]))
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v718+(v710^int32(-1))<<(uint(int32(2))%32))))
	v732 = v724
	goto L148
L150:
	;
	goto L151
L151:
	;
	v726 = *(*int32)(unsafe.Add(mBase, _c_F_gininsert[4]))
	v732 = v726 + v710<<(uint(int32(13))%32) + int32(-8192)
	goto L148
L152:
	;
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v391)+4))
	v738 = int32(_a_F_gininsert_13)
	v740 = *(*int32)(unsafe.Add(mBase, _c_F_gininsert[7]))
	*(*int32)(unsafe.Add(mBase, _c_F_gininsert[7])) = v740 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v388)+88)) = v737
	if v411 != 0 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L4
	} else {
		goto L156
	}
L154:
	;
	goto L155
L155:
	;
	v747 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v732)+16)))
	v748 = v732 + v747
	v749 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v748)+4)))
	v751 = v749 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v748)+4)) = uint16(v751)
	v753 = *(*int64)(unsafe.Add(mBase, uint32(v440)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v440)+40)) = v753 + int64(1)
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v391)+4))
	if v757 != 0 {
		goto L157
	} else {
		goto L158
	}
L156:
	;
	goto L155
L157:
	;
	v758 = int32(1)
	if base.Ui32(v733) < base.Ui32(int32(25)) {
		goto L160
	} else {
		goto L161
	}
L158:
	;
	goto L159
L159:
	;
	F_MarkBufferDirty(m, v710)
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L4
	} else {
		goto L171
	}
L160:
	;
	v767 = v758
	goto L162
L161:
	;
	v767 = int32(base.Ui32(v733+int32(_a_F_gininsert_14))>>(uint(int32(2))%32)) + v758
	goto L162
L162:
	;
	v769 = v767
	v771 = v735
	v781 = v9
	goto L163
L163:
	;
	v789 = v781 << (uint(int32(2)) % 32)
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v789+v790)))
	v793 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v792)+6)))
	v795 = v793 & int32(_a_F_gininsert_6)
	v799 = F_PageAddItemExtended(m, v732, v792, v795, v769&int32(_a_F_gininsert_5), int32(0))
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L4
	} else {
		goto L165
	}
L164:
	;
	goto L159
L165:
	;
	if v799 == int32(0) {
		goto L67
	} else {
		goto L166
	}
L166:
	;
	if v795 != 0 {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v803 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	v805 = *(*int32)(unsafe.Add(mBase, uint32(v803+v789)))
	base.MemoryCopy(m, v771, v805, v795)
	goto L169
L168:
	;
	goto L169
L169:
	;
	v807 = int32(1)
	v811 = v781 + v807
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v391)+4))
	if base.Ui32(v811) < base.Ui32(v812) {
		v769 = v769 + v807
		v771 = v771 + v795
		v781 = v811
		goto L163
	} else {
		goto L170
	}
L170:
	;
	goto L164
L171:
	;
	if v411 == int32(0) {
		goto L82
	} else {
		goto L172
	}
L172:
	;
	F_XLogRegisterBuffer(m, int32(1), v710, int32(8))
	mBase = m.M
	v843 = m.ExcPending
	if v843 != 0 {
		goto L4
	} else {
		goto L173
	}
L173:
	;
	v845 = *(*int32)(unsafe.Add(mBase, uint32(v391)+12))
	F_XLogRegisterBufData(m, int32(1), v735, v845)
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
		goto L4
	} else {
		goto L174
	}
L174:
	;
	v848 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v732)+14)))
	v849 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v732)+12)))
	v850 = v848 - v849
	v851 = int32(0)
	if v851 < v850 {
		goto L176
	} else {
		goto L177
	}
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v440)+32)) = v854
	v857 = v440 + int32(24)
	v858 = v710
	v867 = v732
	goto L83
L176:
	;
	v854 = v850
	goto L178
L177:
	;
	v854 = v851
	goto L178
L178:
	;
	goto L175
L179:
	;
	v880 = *(*int64)(unsafe.Add(mBase, uint32(v857)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v388)+72)) = v880
	v882 = *(*int64)(unsafe.Add(mBase, uint32(v857)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v388)+64)) = v882
	v884 = *(*int64)(unsafe.Add(mBase, uint32(v857)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v388)+56)) = v884
	v886 = *(*int64)(unsafe.Add(mBase, uint32(v857)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v388)+48)) = v886
	v888 = *(*int64)(unsafe.Add(mBase, uint32(v857)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v388)+40)) = v888
	v890 = *(*int64)(unsafe.Add(mBase, uint32(v857)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v388)+32)) = v890
	v892 = *(*int64)(unsafe.Add(mBase, uint32(v857)))
	*(*int64)(unsafe.Add(mBase, uint32(v388)+24)) = v892
	F_XLogRegisterBuffer(m, int32(0), v421, int32(14))
	mBase = m.M
	v897 = m.ExcPending
	if v897 != 0 {
		goto L4
	} else {
		goto L180
	}
L180:
	;
	F_XLogRegisterData(m, v388+int32(8), int32(88))
	mBase = m.M
	v902 = m.ExcPending
	if v902 != 0 {
		goto L4
	} else {
		goto L181
	}
L181:
	;
	v905 = F_XLogInsert(m, int32(13), int32(96))
	mBase = m.M
	v906 = m.ExcPending
	if v906 != 0 {
		goto L4
	} else {
		goto L182
	}
L182:
	;
	v908 = base.I64_rotl(v905, int64(32))
	*(*int64)(unsafe.Add(mBase, uint32(v440))) = v908
	if v858 == int32(0) {
		goto L79
	} else {
		goto L183
	}
L183:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v867))) = v908
	v949 = v858
	goto L80
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v440)+32)) = v919
	v923 = v710
	goto L81
L185:
	;
	v919 = v915
	goto L187
L186:
	;
	v919 = v916
	goto L187
L187:
	;
	goto L184
L188:
	;
	if v923 == int32(0) {
		goto L79
	} else {
		goto L189
	}
L189:
	;
	v949 = v923
	goto L80
L190:
	;
	goto L79
L191:
	;
	v997 = *(*int32)(unsafe.Add(mBase, uint32(v440)+36))
	v998 = int32(_a_F_gininsert_13)
	v1000 = *(*int32)(unsafe.Add(mBase, _c_F_gininsert[7]))
	*(*int32)(unsafe.Add(mBase, _c_F_gininsert[7])) = v1000 - int32(1)
	F_UnlockReleaseBuffer(m, v421)
	mBase = m.M
	v1005 = m.ExcPending
	if v1005 != 0 {
		goto L4
	} else {
		goto L196
	}
L192:
	;
	v990 = *(*int32)(unsafe.Add(mBase, uint32(v989)+8))
	if v990 != int32(-1) {
		v996 = v990
		goto L191
	} else {
		goto L195
	}
L193:
	;
	goto L194
L194:
	;
	v995 = *(*int32)(unsafe.Add(mBase, _c_F_gininsert[8]))
	v996 = v995
	goto L191
L195:
	;
	goto L194
L196:
	;
	if base.Ui32(v997*int32(_a_F_gininsert_10)) <= base.Ui32(v996<<(uint(int32(10))%32)) {
		goto L68
	} else {
		goto L197
	}
L197:
	;
	v1011 = int32(0)
	F_ginInsertCleanup(m, v43, v1011, int32(1), v1011, v1011)
	mBase = m.M
	v1016 = m.ExcPending
	if v1016 != 0 {
		goto L4
	} else {
		goto L198
	}
L198:
	;
	goto L68
L199:
	;
	v1044 = *(*int32)(unsafe.Add(mBase, uint32(v396)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v388))) = v1044 + int32(4)
	F_errmsg_internal(m, int32(_a_F_gininsert_15), v388)
	mBase = m.M
	v1050 = m.ExcPending
	if v1050 != 0 {
		goto L4
	} else {
		goto L200
	}
L200:
	;
	F_errfinish(m, int32(_a_F_gininsert_8), int32(391), int32(_a_F_gininsert_16))
	mBase = m.M
	v1055 = m.ExcPending
	if v1055 != 0 {
		goto L4
	} else {
		goto L201
	}
L201:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L202:
	;
	m.G0 = v23 + int32(32)
	return int32(0)
}
func F_gintuple_get_key(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int64
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int64
	_ = v26
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v37 int64
	_ = v37
	var v38 int32
	_ = v38
	var v39 int64
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
	var v57 int32
	_ = v57
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v10 == int32(1) {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v17 = F_index_getattr_1(m, l1, int32(1), v14, v8+int32(14))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int64(0)
		} else {
			v39 = v17
			v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+14)))
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
				v57 = v55
			} else {
				v57 = int32(0)
			}
			*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v57)
			m.G0 = v8 + int32(16)
			return v39
		}
	} else {
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v26 = F_index_getattr_1(m, l1, int32(1), v23, v8+int32(15))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int64(0)
		} else {
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l0+base.I32_wrap_i64(v26)&int32(_a_F_gintuple_get_key_0)<<(uint(int32(2))%32))+8))
			v37 = F_index_getattr_1(m, l1, int32(2), v34, v8+int32(14))
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int64(0)
			} else {
				v39 = v37
				v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+14)))
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
					v57 = v55
				} else {
					v57 = int32(0)
				}
				*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v57)
				m.G0 = v8 + int32(16)
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
	var v60 int32
	_ = v60
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
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v223 int32
	_ = v223
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v401 int32
	_ = v401
	var v409 int32
	_ = v409
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
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v530 int32
	_ = v530
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v545 int32
	_ = v545
	var v550 int32
	_ = v550
	var v558 int32
	_ = v558
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
	var v580 int32
	_ = v580
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v591 int32
	_ = v591
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v619 int64
	_ = v619
	var v621 int64
	_ = v621
	var v623 int64
	_ = v623
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v649 int32
	_ = v649
	var v668 int32
	_ = v668
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v685 int32
	_ = v685
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v721 int32
	_ = v721
	var v726 int32
	_ = v726
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v763 int32
	_ = v763
	var v769 int32
	_ = v769
	var v772 int32
	_ = v772
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v803 int32
	_ = v803
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v832 int32
	_ = v832
	var v835 int32
	_ = v835
	var v842 int32
	_ = v842
	var v843 float32
	_ = v843
	var v844 int32
	_ = v844
	var v847 int32
	_ = v847
	var v848 float32
	_ = v848
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v870 int32
	_ = v870
	var v872 int32
	_ = v872
	var v874 int32
	_ = v874
	var v879 int32
	_ = v879
	var v884 int32
	_ = v884
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v950 int32
	_ = v950
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v987 int32
	_ = v987
	var v990 int32
	_ = v990
	var v994 int32
	_ = v994
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1020 int32
	_ = v1020
	var v1023 int32
	_ = v1023
	var v1027 int32
	_ = v1027
	var v1033 int32
	_ = v1033
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1046 int32
	_ = v1046
	var v1052 int32
	_ = v1052
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1063 int32
	_ = v1063
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1074 int32
	_ = v1074
	var v1078 int32
	_ = v1078
	var v1084 int32
	_ = v1084
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1097 int32
	_ = v1097
	var v1103 int32
	_ = v1103
	var v1105 int32
	_ = v1105
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1116 int32
	_ = v1116
	var v1133 int32
	_ = v1133
	var v1154 int32
	_ = v1154
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1161 int32
	_ = v1161
	var v1164 int32
	_ = v1164
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1199 int32
	_ = v1199
	var v1201 int32
	_ = v1201
	var v1203 int32
	_ = v1203
	var v1206 int32
	_ = v1206
	var v1207 int32
	_ = v1207
	var v1238 int32
	_ = v1238
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1242 int32
	_ = v1242
	var v1266 int32
	_ = v1266
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
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v26)+68))
	if v328 != 0 {
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
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v51+(l1^int32(-1))*int32(56))+16))
	v66 = v57
	goto L5
L7:
	;
	goto L8
L8:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_gistbufferinginserttuples[1]))
	v60 = int32(56)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v59+l1*v60-v60)+16))
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
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v84+int32(20)+v129<<(uint(int32(2))%32))))
	v154 = v84 + v151&int32(_a_F_gistbufferinginserttuples_5)
	v155 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v154))))
	v158 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v154)+2)))
	v159 = v155<<(uint(int32(16))%32) | v158
	v160 = F_ReadBuffer(m, v147, v159)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L2
	} else {
		goto L25
	}
L24:
	;
	goto L1
L25:
	;
	F_LockBufferInternal(m, v160, int32(1))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L2
	} else {
		goto L26
	}
L26:
	;
	if v160 < int32(0) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	if v160 < int32(0) {
		goto L33
	} else {
		goto L34
	}
L28:
	;
	v168 = *(*int32)(unsafe.Add(mBase, _c_F_gistbufferinginserttuples[0]))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v168+(v160^int32(-1))*int32(56))+16))
	v183 = v174
	goto L27
L29:
	;
	goto L30
L30:
	;
	v176 = *(*int32)(unsafe.Add(mBase, _c_F_gistbufferinginserttuples[1]))
	v177 = int32(56)
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v176+v160*v177-v177)+16))
	v183 = v182
	goto L27
L31:
	;
	F_UnlockReleaseBuffer(m, v160)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L2
	} else {
		goto L42
	}
L32:
	;
	v202 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v201)+12)))
	if base.Ui32(v202) < base.Ui32(int32(25)) {
		goto L31
	} else {
		goto L36
	}
L33:
	;
	v187 = *(*int32)(unsafe.Add(mBase, _c_F_gistbufferinginserttuples[2]))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v187+(v160^int32(-1))<<(uint(int32(2))%32))))
	v201 = v193
	goto L32
L34:
	;
	goto L35
L35:
	;
	v195 = *(*int32)(unsafe.Add(mBase, _c_F_gistbufferinginserttuples[3]))
	v201 = v195 + v160<<(uint(int32(13))%32) + int32(-8192)
	goto L32
L36:
	;
	v206 = v202 + int32(_a_F_gistbufferinginserttuples_3)
	if v206&int32(_a_F_gistbufferinginserttuples_4) == int32(0) {
		goto L31
	} else {
		goto L37
	}
L37:
	;
	v223 = int32(1)
	goto L38
L38:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v201+int32(20)+v223<<(uint(int32(2))%32))))
	v247 = v201 + v244&int32(_a_F_gistbufferinginserttuples_5)
	v248 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v247))))
	v251 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v247)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+76)) = v248<<(uint(int32(16))%32) | v251
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v260 = F_hash_search(m, v254, v26+int32(76), int32(1), v26+int32(75))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L2
	} else {
		goto L40
	}
L39:
	;
	goto L31
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v260)+4)) = v183
	if v223 != int32(base.Ui32(v206)>>(uint(int32(2))%32))&int32(_a_F_gistbufferinginserttuples_6) {
		v223 = v223 + int32(1)
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+76)) = v159
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v298 = F_hash_search(m, v292, v26+int32(76), int32(1), v26+int32(75))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L2
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v298)+4)) = int32(0)
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
	v1266 = *(*int32)(unsafe.Add(mBase, uint32(v26)+64))
	m.G0 = v26 + int32(80)
	return v1266
L46:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v575 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v26)+68))
	v577 = int32(0)
	v578 = m.G0
	v580 = v578 - int32(976)
	m.G0 = v580
	if l2 == v577 {
		goto L94
	} else {
		goto L95
	}
L47:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L2
	} else {
		goto L91
	}
L48:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L2
	} else {
		goto L88
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
	v519 = m.ExcPending
	if v519 != 0 {
		goto L2
	} else {
		goto L87
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
	v332 = *(*int32)(unsafe.Add(mBase, _c_F_gistbufferinginserttuples[0]))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v332+(l1^int32(-1))*int32(56))+16))
	v347 = v338
	goto L52
L54:
	;
	goto L55
L55:
	;
	v340 = *(*int32)(unsafe.Add(mBase, _c_F_gistbufferinginserttuples[1]))
	v341 = int32(56)
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v340+l1*v341-v341)+16))
	v347 = v346
	goto L52
L56:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v367 = F_ReadBuffer(m, v366, v365)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L2
	} else {
		goto L64
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+76)) = v347
	v351 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v357 = F_hash_search(m, v351, v26+int32(76), int32(0), v26+int32(75))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
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
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+75)))
	if v359 == int32(0) {
		goto L48
	} else {
		goto L61
	}
L61:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v357)+4))
	v365 = v362
	goto L56
L62:
	;
	v365 = l6
	goto L56
L63:
	;
	F_LockBufferInternal(m, v367, int32(3))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L2
	} else {
		goto L68
	}
L64:
	;
	if v367 < int32(0) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v372 = *(*int32)(unsafe.Add(mBase, _c_F_gistbufferinginserttuples[2]))
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v372+(v367^int32(-1))<<(uint(int32(2))%32))))
	v386 = v378
	goto L63
L66:
	;
	goto L67
L67:
	;
	v380 = *(*int32)(unsafe.Add(mBase, _c_F_gistbufferinginserttuples[3]))
	v386 = v380 + v367<<(uint(int32(13))%32) + int32(-8192)
	goto L63
L68:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_gistcheckpage(m, v390, v367)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L2
	} else {
		goto L69
	}
L69:
	;
	v401 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v386)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v401) {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L2
	} else {
		goto L84
	}
L71:
	;
	v444 = int32(1)
	goto L80
L72:
	;
	v409 = int32(base.Ui32(v401+int32(_a_F_gistbufferinginserttuples_3)) >> (uint(int32(2)) % 32))
	goto L74
L73:
	;
	v409 = int32(0)
	goto L74
L74:
	;
	if base.B2i32(l6 == int32(-1))|base.B2i32(v365 != l6)|base.B2i32(base.Ui32(v409&int32(_a_F_gistbufferinginserttuples_6)) <= base.Ui32((l7-int32(1))&int32(_a_F_gistbufferinginserttuples_6))) == int32(0) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v386+l7<<(uint(int32(2))%32))+20))
	v422 = v386 + v419&int32(_a_F_gistbufferinginserttuples_5)
	v423 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v422))))
	v426 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v422)+2)))
	if v423<<(uint(int32(16))%32)|v426 != v347 {
		goto L71
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	if v409&int32(_a_F_gistbufferinginserttuples_6) == int32(0) {
		goto L70
	} else {
		goto L79
	}
L78:
	;
	v558 = l7
	goto L46
L79:
	;
	goto L71
L80:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v386+int32(20)+v444&int32(_a_F_gistbufferinginserttuples_6)<<(uint(int32(2))%32))))
	v468 = v386 + v465&int32(_a_F_gistbufferinginserttuples_5)
	v469 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v468))))
	v472 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v468)+2)))
	if v469<<(uint(int32(16))%32)|v472 == v347 {
		v558 = v444
		goto L46
	} else {
		goto L82
	}
L81:
	;
	goto L70
L82:
	;
	v476 = v444 + int32(1)
	v477 = int32(_a_F_gistbufferinginserttuples_6)
	if base.Ui32(v476&v477) <= base.Ui32(v409&v477) {
		v444 = v476
		goto L80
	} else {
		goto L83
	}
L83:
	;
	goto L81
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v347
	F_errmsg_internal(m, int32(_a_F_gistbufferinginserttuples_7), v26)
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L2
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(_a_F_gistbufferinginserttuples_1), int32(1289), int32(_a_F_gistbufferinginserttuples_8))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L2
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L87:
	;
	goto L45
L88:
	;
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v26)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v524
	F_errmsg_internal(m, int32(_a_F_gistbufferinginserttuples_9), v26+int32(16))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L2
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(_a_F_gistbufferinginserttuples_1), int32(1578), int32(_a_F_gistbufferinginserttuples_10))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L2
	} else {
		goto L90
	}
L90:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v347
	F_errmsg_internal(m, int32(_a_F_gistbufferinginserttuples_11), v26+int32(32))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L2
	} else {
		goto L92
	}
L92:
	;
	F_errfinish(m, int32(_a_F_gistbufferinginserttuples_1), int32(1245), int32(_a_F_gistbufferinginserttuples_8))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L2
	} else {
		goto L93
	}
L93:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L94:
	;
	m.G0 = v580 + int32(976)
	v978 = *(*int32)(unsafe.Add(mBase, uint32(v26)+68))
	if v978 != 0 {
		goto L156
	} else {
		goto L157
	}
L95:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v28)+32))
	v585 = base.I32_rem_s(l2, v584)
	if v585 != 0 {
		goto L94
	} else {
		goto L96
	}
L96:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v28)+60))
	if l2 == v586 {
		goto L94
	} else {
		goto L97
	}
L97:
	;
	if l1 < int32(0) {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v580)+968)) = v606
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v28)+24))
	v614 = F_hash_search(m, v608, v580+int32(968), int32(0), v580+int32(975))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L2
	} else {
		goto L102
	}
L99:
	;
	v591 = *(*int32)(unsafe.Add(mBase, _c_F_gistbufferinginserttuples[0]))
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v591+(l1^int32(-1))*int32(56))+16))
	v606 = v597
	goto L98
L100:
	;
	goto L101
L101:
	;
	v599 = *(*int32)(unsafe.Add(mBase, _c_F_gistbufferinginserttuples[1]))
	v600 = int32(56)
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v599+l1*v600-v600)+16))
	v606 = v605
	goto L98
L102:
	;
	v616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v580)+975)))
	if v616 != int32(1) {
		goto L94
	} else {
		goto L103
	}
L103:
	;
	v619 = *(*int64)(unsafe.Add(mBase, uint32(v614)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v580)+152)) = v619
	v621 = *(*int64)(unsafe.Add(mBase, uint32(v614)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v580)+144)) = v621
	v623 = *(*int64)(unsafe.Add(mBase, uint32(v614)))
	*(*int64)(unsafe.Add(mBase, uint32(v580)+136)) = v623
	v625 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v580)+153)) = uint8(v625)
	v627 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v614)+12)) = v627
	*(*int64)(unsafe.Add(mBase, uint32(v614)+4)) = int64(-4294967296)
	if v576 == v627 {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	v736 = F_gistPopItupFromNodeBuffer(m, v28, v580+int32(136), v580+int32(964))
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L2
	} else {
		goto L120
	}
L105:
	;
	v635 = F_palloc_mul(m, int32(808), int32(0))
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L2
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v576)+4))
	v639 = F_palloc_mul(m, int32(808), v638)
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L2
	} else {
		goto L109
	}
L108:
	;
	v721 = v635
	v726 = v577
	goto L104
L109:
	;
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v576)+4))
	if v641 <= int32(0) {
		v721 = v639
		v726 = v638
		goto L104
	} else {
		goto L110
	}
L110:
	;
	v649 = int32(0)
	goto L111
L111:
	;
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v576)+12))
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v668+v649<<(uint(int32(2))%32))))
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v672)+4))
	v676 = v639 + v649*int32(808)
	F_gistDeCompressAtt(m, v574, v575, v673, v676, v676+int32(768))
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L2
	} else {
		goto L113
	}
L112:
	;
	v721 = v639
	v726 = v638
	goto L104
L113:
	;
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v672)))
	if v681 < int32(0) {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	v701 = F_gistGetNodeBuffer(m, v28, v700, l2)
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L2
	} else {
		goto L118
	}
L115:
	;
	v685 = *(*int32)(unsafe.Add(mBase, _c_F_gistbufferinginserttuples[0]))
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v685+(v681^int32(-1))*int32(56))+16))
	v700 = v691
	goto L114
L116:
	;
	goto L117
L117:
	;
	v693 = *(*int32)(unsafe.Add(mBase, _c_F_gistbufferinginserttuples[1]))
	v694 = int32(56)
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v693+v681*v694-v694)+16))
	v700 = v699
	goto L114
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v676)+800)) = v672
	*(*int32)(unsafe.Add(mBase, uint32(v676)+804)) = v701
	v706 = v649 + int32(1)
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v576)+4))
	if v706 < v707 {
		v649 = v706
		goto L111
	} else {
		goto L119
	}
L119:
	;
	goto L112
L120:
	;
	if v736 != 0 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	goto L124
L122:
	;
	goto L123
L123:
	;
	F_pfree(m, v721)
	mBase = m.M
	v950 = m.ExcPending
	if v950 != 0 {
		goto L2
	} else {
		goto L155
	}
L124:
	;
	v763 = *(*int32)(unsafe.Add(mBase, uint32(v580)+964))
	F_gistDeCompressAtt(m, v574, v575, v763, v580+int32(192), v580+int32(160))
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L2
	} else {
		goto L126
	}
L125:
	;
	goto L123
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v580))) = int32(-1082130432)
	v772 = int32(0)
	if v726 <= int32(0) {
		v884 = v772
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v906 = v721 + v884*int32(808)
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v906)+804))
	F_gistPushItupToNodeBuffer(m, v28, v907, v763)
	mBase = m.M
	v909 = m.ExcPending
	if v909 != 0 {
		goto L2
	} else {
		goto L147
	}
L128:
	;
	v775 = v772
	v777 = v772
	goto L129
L129:
	;
	v797 = *(*int32)(unsafe.Add(mBase, uint32(v575)+192))
	v798 = int32(*(*int16)(unsafe.Add(mBase, uint32(v797)+10)))
	if v798 <= int32(0) {
		v884 = v777
		goto L127
	} else {
		goto L131
	}
L130:
	;
	v884 = v874
	goto L127
L131:
	;
	v803 = v721 + v775*int32(808)
	v811 = v777
	v812 = int32(1)
	v813 = int32(0)
	goto L133
L132:
	;
	v879 = v775 + int32(1)
	if v879 != v726 {
		v775 = v879
		v777 = v874
		goto L129
	} else {
		goto L146
	}
L133:
	;
	v832 = v813 * int32(24)
	v835 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v813+(v803+int32(768))))))
	v842 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v580+int32(160)+v813))))
	v843 = F_gistpenalty(m, v574, v813, v803+v832, v835, v580+int32(192)+v832, v842)
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L2
	} else {
		goto L135
	}
L134:
	;
	if v870 != 0 {
		v884 = v864
		goto L127
	} else {
		goto L145
	}
L135:
	;
	v847 = v580 + v813<<(uint(int32(2))%32)
	v848 = *(*float32)(unsafe.Add(mBase, uint32(v847)))
	if base.F32_lt(v848, float32(0))|base.F32_lt(v843, v848) != 0 {
		goto L137
	} else {
		goto L138
	}
L136:
	;
	v870 = base.B2i32(base.F32_gt(v843, float32(0)) == int32(0)) & v812
	v872 = v813 + int32(1)
	if v872 < v865 {
		v811 = v864
		v812 = v870
		v813 = v872
		goto L133
	} else {
		goto L144
	}
L137:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v847))) = v843
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v575)+192))
	v855 = int32(*(*int16)(unsafe.Add(mBase, uint32(v854)+10)))
	if v855-int32(1) <= v813 {
		goto L140
	} else {
		goto L141
	}
L138:
	;
	goto L139
L139:
	;
	if base.F32_ne(v843, v848) != 0 {
		v874 = v811
		goto L132
	} else {
		goto L143
	}
L140:
	;
	v864 = v775
	v865 = v855
	goto L136
L141:
	;
	goto L142
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v847)+4)) = int32(-1082130432)
	v864 = v775
	v865 = v855
	goto L136
L143:
	;
	v862 = *(*int32)(unsafe.Add(mBase, uint32(v575)+192))
	v863 = int32(*(*int16)(unsafe.Add(mBase, uint32(v862)+10)))
	v864 = v811
	v865 = v863
	goto L136
L144:
	;
	goto L134
L145:
	;
	v874 = v864
	goto L132
L146:
	;
	goto L130
L147:
	;
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v906)+800))
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v910)+4))
	v912 = F_gistgetadjusted(m, v575, v911, v763, v574)
	mBase = m.M
	v913 = m.ExcPending
	if v913 != 0 {
		goto L2
	} else {
		goto L148
	}
L148:
	;
	if v912 != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	F_gistDeCompressAtt(m, v574, v575, v912, v906, v906+int32(768))
	mBase = m.M
	v917 = m.ExcPending
	if v917 != 0 {
		goto L2
	} else {
		goto L152
	}
L150:
	;
	goto L151
L151:
	;
	v924 = F_gistPopItupFromNodeBuffer(m, v28, v580+int32(136), v580+int32(964))
	mBase = m.M
	v925 = m.ExcPending
	if v925 != 0 {
		goto L2
	} else {
		goto L153
	}
L152:
	;
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v906)+800))
	*(*int32)(unsafe.Add(mBase, uint32(v918)+4)) = v912
	goto L151
L153:
	;
	if v924 != 0 {
		goto L124
	} else {
		goto L154
	}
L154:
	;
	goto L125
L155:
	;
	goto L94
L156:
	;
	v979 = *(*int32)(unsafe.Add(mBase, uint32(v978)+4))
	v981 = v979
	goto L158
L157:
	;
	v981 = int32(0)
	goto L158
L158:
	;
	v982 = F_palloc_mul(m, int32(4), v981)
	mBase = m.M
	v983 = m.ExcPending
	if v983 != 0 {
		goto L2
	} else {
		goto L159
	}
L159:
	;
	v984 = *(*int32)(unsafe.Add(mBase, uint32(v26)+68))
	if v984 == int32(0) {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v1238 = F_gistbufferinginserttuples(m, l0, v367, l2+int32(1), v982, v981, v558&int32(_a_F_gistbufferinginserttuples_6), int32(-1), int32(0))
	mBase = m.M
	v1239 = m.ExcPending
	if v1239 != 0 {
		goto L2
	} else {
		goto L193
	}
L161:
	;
	v987 = *(*int32)(unsafe.Add(mBase, uint32(v984)+4))
	if v987 <= int32(0) {
		goto L160
	} else {
		goto L162
	}
L162:
	;
	v990 = int32(0)
	v994 = v990
	goto L163
L163:
	;
	v1017 = v994 << (uint(int32(2)) % 32)
	v1018 = *(*int32)(unsafe.Add(mBase, uint32(v984)+12))
	v1020 = *(*int32)(unsafe.Add(mBase, uint32(v1017+v1018)))
	if base.B2i32(v990 < l2) == int32(0) {
		goto L165
	} else {
		goto L166
	}
L164:
	;
	goto L160
L165:
	;
	v1199 = *(*int32)(unsafe.Add(mBase, uint32(v1020)))
	F_UnlockReleaseBuffer(m, v1199)
	mBase = m.M
	v1201 = m.ExcPending
	if v1201 != 0 {
		goto L2
	} else {
		goto L191
	}
L166:
	;
	v1023 = *(*int32)(unsafe.Add(mBase, uint32(v1020)))
	if v1023 < int32(0) {
		goto L168
	} else {
		goto L169
	}
L167:
	;
	if v367 < int32(0) {
		goto L172
	} else {
		goto L173
	}
L168:
	;
	v1027 = *(*int32)(unsafe.Add(mBase, _c_F_gistbufferinginserttuples[0]))
	v1033 = *(*int32)(unsafe.Add(mBase, uint32(v1027+(v1023^int32(-1))*int32(56))+16))
	v1042 = v1033
	goto L167
L169:
	;
	goto L170
L170:
	;
	v1035 = *(*int32)(unsafe.Add(mBase, _c_F_gistbufferinginserttuples[1]))
	v1036 = int32(56)
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(v1035+v1023*v1036-v1036)+16))
	v1042 = v1041
	goto L167
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+76)) = v1042
	v1063 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v1069 = F_hash_search(m, v1063, v26+int32(76), int32(1), v26+int32(75))
	mBase = m.M
	v1070 = m.ExcPending
	if v1070 != 0 {
		goto L2
	} else {
		goto L175
	}
L172:
	;
	v1046 = *(*int32)(unsafe.Add(mBase, _c_F_gistbufferinginserttuples[0]))
	v1052 = *(*int32)(unsafe.Add(mBase, uint32(v1046+(v367^int32(-1))*int32(56))+16))
	v1061 = v1052
	goto L171
L173:
	;
	goto L174
L174:
	;
	v1054 = *(*int32)(unsafe.Add(mBase, _c_F_gistbufferinginserttuples[1]))
	v1055 = int32(56)
	v1060 = *(*int32)(unsafe.Add(mBase, uint32(v1054+v367*v1055-v1055)+16))
	v1061 = v1060
	goto L171
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1069)+4)) = v1061
	if l2 == int32(1) {
		goto L165
	} else {
		goto L176
	}
L176:
	;
	v1074 = *(*int32)(unsafe.Add(mBase, uint32(v1020)))
	if v1074 < int32(0) {
		goto L178
	} else {
		goto L179
	}
L177:
	;
	if v1074 < int32(0) {
		goto L182
	} else {
		goto L183
	}
L178:
	;
	v1078 = *(*int32)(unsafe.Add(mBase, _c_F_gistbufferinginserttuples[0]))
	v1084 = *(*int32)(unsafe.Add(mBase, uint32(v1078+(v1074^int32(-1))*int32(56))+16))
	v1093 = v1084
	goto L177
L179:
	;
	goto L180
L180:
	;
	v1086 = *(*int32)(unsafe.Add(mBase, _c_F_gistbufferinginserttuples[1]))
	v1087 = int32(56)
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(v1086+v1074*v1087-v1087)+16))
	v1093 = v1092
	goto L177
L181:
	;
	v1112 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1111)+12)))
	if base.Ui32(v1112) < base.Ui32(int32(25)) {
		goto L165
	} else {
		goto L185
	}
L182:
	;
	v1097 = *(*int32)(unsafe.Add(mBase, _c_F_gistbufferinginserttuples[2]))
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(v1097+(v1074^int32(-1))<<(uint(int32(2))%32))))
	v1111 = v1103
	goto L181
L183:
	;
	goto L184
L184:
	;
	v1105 = *(*int32)(unsafe.Add(mBase, _c_F_gistbufferinginserttuples[3]))
	v1111 = v1105 + v1074<<(uint(int32(13))%32) + int32(-8192)
	goto L181
L185:
	;
	v1116 = v1112 + int32(_a_F_gistbufferinginserttuples_3)
	if v1116&int32(_a_F_gistbufferinginserttuples_4) == int32(0) {
		goto L165
	} else {
		goto L186
	}
L186:
	;
	v1133 = int32(1)
	goto L187
L187:
	;
	v1154 = *(*int32)(unsafe.Add(mBase, uint32(v1111+int32(20)+v1133<<(uint(int32(2))%32))))
	v1157 = v1111 + v1154&int32(_a_F_gistbufferinginserttuples_5)
	v1158 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1157))))
	v1161 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1157)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+76)) = v1158<<(uint(int32(16))%32) | v1161
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v1170 = F_hash_search(m, v1164, v26+int32(76), int32(1), v26+int32(75))
	mBase = m.M
	v1171 = m.ExcPending
	if v1171 != 0 {
		goto L2
	} else {
		goto L189
	}
L188:
	;
	goto L165
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1170)+4)) = v1093
	if v1133 != int32(base.Ui32(v1116)>>(uint(int32(2))%32))&int32(_a_F_gistbufferinginserttuples_6) {
		v1133 = v1133 + int32(1)
		goto L187
	} else {
		goto L190
	}
L190:
	;
	goto L188
L191:
	;
	v1203 = *(*int32)(unsafe.Add(mBase, uint32(v1020)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v982+v1017))) = v1203
	v1206 = v994 + int32(1)
	v1207 = *(*int32)(unsafe.Add(mBase, uint32(v984)+4))
	if v1206 < v1207 {
		v994 = v1206
		goto L163
	} else {
		goto L192
	}
L192:
	;
	goto L164
L193:
	;
	v1240 = *(*int32)(unsafe.Add(mBase, uint32(v26)+68))
	F_list_free_deep(m, v1240)
	mBase = m.M
	v1242 = m.ExcPending
	if v1242 != 0 {
		goto L2
	} else {
		goto L194
	}
L194:
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
func F_gisthandler(m *base.Module, l0 int32) int64 {
	return int64(829504)
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
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
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
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v94 int32
	_ = v94
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v288 int32
	_ = v288
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v333 int32
	_ = v333
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int64
	_ = v356
	var v358 int32
	_ = v358
	var v360 int64
	_ = v360
	var v362 int64
	_ = v362
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v397 int32
	_ = v397
	var v411 int32
	_ = v411
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v427 int32
	_ = v427
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v476 int32
	_ = v476
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int64
	_ = v513
	var v515 int32
	_ = v515
	var v517 int64
	_ = v517
	var v519 int64
	_ = v519
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v545 int32
	_ = v545
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v577 int32
	_ = v577
	var v582 int32
	_ = v582
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+40)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v16)+28)) = int64(-4294967296)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	if v21 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v39 != int32(1) {
		goto L9
	} else {
		goto L10
	}
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	if v24 == v26 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v32 = F_AllocSetContextCreateInternal(m, v26, int32(_a_F_gistrescan_0), int32(0), int32(_a_F_gistrescan_1), int32(_a_F_gistrescan_2))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	F_MemoryContextReset(m, v24)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L6
	} else {
		goto L8
	}
L6:
	;
	return
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v32
	goto L1
L8:
	;
	goto L1
L9:
	;
	v257 = int32(_a_F_gistrescan_3)
	v258 = *(*int32)(unsafe.Add(mBase, _c_F_gistrescan[0]))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	*(*int32)(unsafe.Add(mBase, _c_F_gistrescan[0])) = v260
	v263 = F_pairingheap_allocate(m, int32(121), l0)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L6
	} else {
		goto L47
	}
L10:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v42 != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+192))
	v45 = int32(*(*int16)(unsafe.Add(mBase, uint32(v44)+10)))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+48))
	v47 = int32(*(*int16)(unsafe.Add(mBase, uint32(v46)+120)))
	v48 = F_CreateTemplateTupleDesc(m, v47)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	*(*int32)(unsafe.Add(mBase, uint32(v50)+16)) = v48
	if int32(0) < v45 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v55 = int32(1)
	v60 = v55
	goto L16
L14:
	;
	v94 = int32(1)
	goto L15
L15:
	;
	if v94 <= v47 {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
	v72 = int32(0)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+212))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v74+v60<<(uint(int32(2))%32)-int32(4))))
	F_TupleDescInitEntry(m, v70, base.I32_extend16_s(v60), v72, v80, int32(-1), v72)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L6
	} else {
		goto L18
	}
L17:
	;
	v94 = v45 + v55
	goto L15
L18:
	;
	if base.B2i32(v60 == v45) == int32(0) {
		v60 = v60 + int32(1)
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v106 = v94
	goto L23
L21:
	;
	goto L22
L22:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)+16))
	v147 = int32(0)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
	if v147 < v156 {
		goto L28
	} else {
		goto L29
	}
L23:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+16))
	v115 = int32(0)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v113)+8))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v116+v117<<(uint(int32(3))%32)+v106*int32(100)-int32(4))))
	F_TupleDescInitEntry(m, v114, v106, v115, v126, int32(-1), v115)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L6
	} else {
		goto L25
	}
L24:
	;
	goto L22
L25:
	;
	if v106 != v47 {
		v106 = v106 + int32(1)
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v234)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v235
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v237)))
	v243 = F_AllocSetContextCreateInternal(m, v238, int32(_a_F_gistrescan_4), int32(0), int32(_a_F_gistrescan_1), int32(_a_F_gistrescan_2))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L6
	} else {
		goto L46
	}
L28:
	;
	v160 = v146 + int32(28)
	v167 = v147
	v168 = v156
	v170 = v147
	goto L32
L29:
	;
	v224 = v147
	v231 = v156
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v146)+20)) = v231
	*(*int32)(unsafe.Add(mBase, uint32(v146)+16)) = v224
	goto L27
L31:
	;
	v224 = v218
	v231 = v197
	goto L30
L32:
	;
	v176 = v160 + v156<<(uint(int32(3))%32) + v167*int32(100)
	v179 = v160 + v167<<(uint(int32(3))%32)
	if v156 != v168 {
		v197 = v168
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v218 = v156
	goto L31
L34:
	;
	v198 = int32(*(*int16)(unsafe.Add(mBase, uint32(v179)+2)))
	if v198 <= int32(0) {
		v218 = v167
		goto L31
	} else {
		goto L42
	}
L35:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+7)))
	if v181 != int32(118) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v197 = v167
	goto L34
L37:
	;
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+4)))
	if v184 != int32(1) {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+6)))
	if v187&int32(6) != 0 {
		goto L36
	} else {
		goto L39
	}
L39:
	;
	v190 = int32(*(*int16)(unsafe.Add(mBase, uint32(v179)+2)))
	if v190 <= int32(0) {
		goto L36
	} else {
		goto L40
	}
L40:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176)+90)))
	if v193 != int32(118) {
		v197 = v156
		goto L34
	} else {
		goto L41
	}
L41:
	;
	goto L36
L42:
	;
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176)+90)))
	if v201 == int32(118) {
		v218 = v167
		goto L31
	} else {
		goto L43
	}
L43:
	;
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+5)))
	v210 = (v170 + v204 - int32(1)) & (int32(0) - v204)
	if int32(_a_F_gistrescan_5) < v210 {
		v218 = v167
		goto L31
	} else {
		goto L44
	}
L44:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v179))) = uint16(v210)
	v216 = v167 + int32(1)
	if v216 != v156 {
		v167 = v216
		v168 = v197
		v170 = v210 + v198
		goto L32
	} else {
		goto L45
	}
L45:
	;
	goto L33
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+uint32(_c_F_gistrescan[1]))) = v243
	goto L9
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v263
	*(*int32)(unsafe.Add(mBase, _c_F_gistrescan[0])) = v258
	v268 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+17)) = uint8(v268)
	if l1 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	if l3 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L49:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v272 <= int32(0) {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	if v21 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v322 = v314 * int32(56)
	if v322 != 0 {
		goto L60
	} else {
		goto L61
	}
L52:
	;
	v314 = v272
	v316 = int32(0)
	goto L51
L53:
	;
	goto L54
L54:
	;
	v280 = F_palloc(m, v272<<(uint(int32(2))%32))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L6
	} else {
		goto L55
	}
L55:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v282 <= int32(0) {
		v314 = v282
		v316 = v280
		goto L51
	} else {
		goto L56
	}
L56:
	;
	v288 = int32(0)
	goto L57
L57:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v300+v288*int32(56))+32))
	*(*int32)(unsafe.Add(mBase, uint32(v280+v288<<(uint(int32(2))%32)))) = v304
	v307 = v288 + int32(1)
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v307 < v308 {
		v288 = v307
		goto L57
	} else {
		goto L59
	}
L58:
	;
	v314 = v308
	v316 = v280
	goto L51
L59:
	;
	goto L58
L60:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	base.MemoryCopy(m, v323, l1, v322)
	goto L62
L61:
	;
	goto L62
L62:
	;
	v325 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+16)) = uint8(v325)
	v327 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if int32(0) < v327 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v333 = int32(0)
	goto L66
L64:
	;
	goto L65
L65:
	;
	if v21 == int32(0) {
		goto L48
	} else {
		goto L76
	}
L66:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v345 = v342 + v333*int32(56)
	v347 = v345 + int32(16)
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v349 = int32(*(*int16)(unsafe.Add(mBase, uint32(v345)+4)))
	v354 = v348 + v349*int32(28) - int32(8)
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v348)))
	v356 = *(*int64)(unsafe.Add(mBase, uint32(v354)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v347)+16)) = v356
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v354)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v347)+24)) = v358
	v360 = *(*int64)(unsafe.Add(mBase, uint32(v354)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v347)+8)) = v360
	v362 = *(*int64)(unsafe.Add(mBase, uint32(v354)))
	*(*int64)(unsafe.Add(mBase, uint32(v347))) = v362
	*(*int32)(unsafe.Add(mBase, uint32(v347)+20)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v347)+16)) = int32(0)
	goto L68
L67:
	;
	goto L65
L68:
	;
	if v21 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v316+v333<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v345)+32)) = v370
	goto L71
L70:
	;
	goto L71
L71:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v345)))
	if v372&int32(193) == int32(1) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v377 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+16)) = uint8(v377)
	goto L74
L73:
	;
	goto L74
L74:
	;
	v380 = v333 + int32(1)
	v381 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v380 < v381 {
		v333 = v380
		goto L66
	} else {
		goto L75
	}
L75:
	;
	goto L67
L76:
	;
	F_pfree(m, v316)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L6
	} else {
		goto L77
	}
L77:
	;
	goto L48
L78:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L6
	} else {
		goto L109
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = int32(0)
	m.G0 = v14 + int32(16)
	return
L80:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v411 <= int32(0) {
		goto L79
	} else {
		goto L81
	}
L81:
	;
	if v21 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	v461 = v453 * int32(56)
	if v461 != 0 {
		goto L91
	} else {
		goto L92
	}
L83:
	;
	v450 = int32(0)
	v453 = v411
	goto L82
L84:
	;
	goto L85
L85:
	;
	v419 = F_palloc(m, v411<<(uint(int32(2))%32))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L6
	} else {
		goto L86
	}
L86:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v421 <= int32(0) {
		v450 = v419
		v453 = v421
		goto L82
	} else {
		goto L87
	}
L87:
	;
	v427 = int32(0)
	goto L88
L88:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v439+v427*int32(56))+32))
	*(*int32)(unsafe.Add(mBase, uint32(v419+v427<<(uint(int32(2))%32)))) = v443
	v446 = v427 + int32(1)
	v447 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v446 < v447 {
		v427 = v446
		goto L88
	} else {
		goto L90
	}
L89:
	;
	v450 = v419
	v453 = v447
	goto L82
L90:
	;
	goto L89
L91:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	base.MemoryCopy(m, v462, l3, v461)
	goto L93
L92:
	;
	goto L93
L93:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v467 = F_palloc(m, v464<<(uint(int32(2))%32))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L6
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v467
	v470 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if int32(0) < v470 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v476 = int32(0)
	goto L98
L96:
	;
	goto L97
L97:
	;
	if v21 == int32(0) {
		goto L79
	} else {
		goto L107
	}
L98:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v486 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v489 = v486 + v476*int32(56)
	v490 = int32(*(*int16)(unsafe.Add(mBase, uint32(v489)+4)))
	v493 = v485 + v490*int32(28)
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v493)+uint32(_c_F_gistrescan[2])))
	if v496 == int32(0) {
		goto L78
	} else {
		goto L100
	}
L99:
	;
	goto L97
L100:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v489)+20))
	v500 = F_get_func_rettype(m, v499)
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L6
	} else {
		goto L101
	}
L101:
	;
	v503 = v476 << (uint(int32(2)) % 32)
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v503+v504))) = v500
	v508 = v489 + int32(16)
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v511)))
	v513 = *(*int64)(unsafe.Add(mBase, uint32(v493)+uint32(_c_F_gistrescan[3])))
	*(*int64)(unsafe.Add(mBase, uint32(v508)+16)) = v513
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v493)+uint32(_c_F_gistrescan[4])))
	*(*int32)(unsafe.Add(mBase, uint32(v508)+24)) = v515
	v517 = *(*int64)(unsafe.Add(mBase, uint32(v493)+uint32(_c_F_gistrescan[5])))
	*(*int64)(unsafe.Add(mBase, uint32(v508)+8)) = v517
	v519 = *(*int64)(unsafe.Add(mBase, uint32(v493)+uint32(_c_F_gistrescan[6])))
	*(*int64)(unsafe.Add(mBase, uint32(v508))) = v519
	*(*int32)(unsafe.Add(mBase, uint32(v508)+20)) = v512
	*(*int32)(unsafe.Add(mBase, uint32(v508)+16)) = int32(0)
	goto L102
L102:
	;
	if v21 != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v450+v503)))
	*(*int32)(unsafe.Add(mBase, uint32(v489)+32)) = v525
	goto L105
L104:
	;
	goto L105
L105:
	;
	v528 = v476 + int32(1)
	v529 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v528 < v529 {
		v476 = v528
		goto L98
	} else {
		goto L106
	}
L106:
	;
	goto L99
L107:
	;
	F_pfree(m, v450)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L6
	} else {
		goto L108
	}
L108:
	;
	goto L79
L109:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v566)+48))
	v568 = int32(*(*int16)(unsafe.Add(mBase, uint32(v489)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v568
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v567 + int32(4)
	F_errmsg_internal(m, int32(_a_F_gistrescan_6), v14)
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L6
	} else {
		goto L110
	}
L110:
	;
	F_errfinish(m, int32(_a_F_gistrescan_7), int32(317), int32(_a_F_gistrescan_8))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L6
	} else {
		goto L111
	}
L111:
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
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int64
	_ = v201
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v266 int32
	_ = v266
	var v269 int64
	_ = v269
	var v271 int64
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v312 int64
	_ = v312
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v343 int32
	_ = v343
	var v350 int32
	_ = v350
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v422 int64
	_ = v422
	var v423 int32
	_ = v423
	var v426 int64
	_ = v426
	var v427 int32
	_ = v427
	var v430 int64
	_ = v430
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v438 float64
	_ = v438
	var v442 int32
	_ = v442
	var v456 int32
	_ = v456
	var v480 int32
	_ = v480
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v535 int32
	_ = v535
	var v563 int32
	_ = v563
	var v567 int32
	_ = v567
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v595 int32
	_ = v595
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v634 float64
	_ = v634
	var v669 int32
	_ = v669
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v689 int64
	_ = v689
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v696 int32
	_ = v696
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v710 int32
	_ = v710
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v723 int32
	_ = v723
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v784 int32
	_ = v784
	var v785 int64
	_ = v785
	var v787 int64
	_ = v787
	var v790 int64
	_ = v790
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v796 int64
	_ = v796
	var v799 int64
	_ = v799
	var v801 int32
	_ = v801
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v815 int32
	_ = v815
	var v819 int32
	_ = v819
	var v829 int32
	_ = v829
	var v842 int64
	_ = v842
	var v845 int64
	_ = v845
	var v849 int32
	_ = v849
	var v852 int64
	_ = v852
	var v853 int64
	_ = v853
	var v855 int64
	_ = v855
	var v859 int64
	_ = v859
	var v861 int64
	_ = v861
	var v865 int64
	_ = v865
	var v867 int64
	_ = v867
	var v871 int64
	_ = v871
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v875 int64
	_ = v875
	var v877 int32
	_ = v877
	var v883 int32
	_ = v883
	var v906 int64
	_ = v906
	var v909 int64
	_ = v909
	var v913 int32
	_ = v913
	var v916 int32
	_ = v916
	var v936 int64
	_ = v936
	var v939 int64
	_ = v939
	var v947 int64
	_ = v947
	var v949 int32
	_ = v949
	var v953 int32
	_ = v953
	var v1018 int32
	_ = v1018
	var v1021 int32
	_ = v1021
	var v1025 int32
	_ = v1025
	var v1028 int32
	_ = v1028
	var v1031 int32
	_ = v1031
	var v1034 int32
	_ = v1034
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1047 int32
	_ = v1047
	var v1074 int32
	_ = v1074
	var v1077 int32
	_ = v1077
	var v1081 int64
	_ = v1081
	var v1114 int32
	_ = v1114
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1131 int32
	_ = v1131
	var v1137 int32
	_ = v1137
	var v1139 int32
	_ = v1139
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1149 int32
	_ = v1149
	var v1155 int32
	_ = v1155
	var v1156 int32
	_ = v1156
	var v1164 int32
	_ = v1164
	var v1166 int32
	_ = v1166
	var v1170 int32
	_ = v1170
	var v1172 int32
	_ = v1172
	var v1183 int32
	_ = v1183
	var v1190 int32
	_ = v1190
	var v1213 int32
	_ = v1213
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1222 int64
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1227 int64
	_ = v1227
	var v1230 int32
	_ = v1230
	var v1235 int32
	_ = v1235
	var v1240 int32
	_ = v1240
	var v1264 int32
	_ = v1264
	var v1265 int32
	_ = v1265
	var v1269 int64
	_ = v1269
	var v1270 int32
	_ = v1270
	var v1271 int32
	_ = v1271
	var v1274 int32
	_ = v1274
	var v1281 int64
	_ = v1281
	var v1283 int32
	_ = v1283
	var v1286 int32
	_ = v1286
	var v1288 int32
	_ = v1288
	var v1307 int32
	_ = v1307
	var v1312 int32
	_ = v1312
	var v1321 int32
	_ = v1321
	var v1330 int32
	_ = v1330
	var v1335 int32
	_ = v1335
	var v1359 int32
	_ = v1359
	var v1360 int32
	_ = v1360
	var v1364 int64
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1366 int32
	_ = v1366
	var v1369 int32
	_ = v1369
	var v1371 int32
	_ = v1371
	var v1374 int32
	_ = v1374
	var v1379 int32
	_ = v1379
	var v1403 int32
	_ = v1403
	var v1412 int32
	_ = v1412
	var v1416 int32
	_ = v1416
	var v1421 int32
	_ = v1421
	var v1426 int32
	_ = v1426
	var v1450 int32
	_ = v1450
	var v1451 int32
	_ = v1451
	var v1455 int64
	_ = v1455
	var v1456 int32
	_ = v1456
	var v1457 int32
	_ = v1457
	var v1460 int32
	_ = v1460
	var v1467 int32
	_ = v1467
	var v1470 int64
	_ = v1470
	var v1474 int64
	_ = v1474
	var v1482 int32
	_ = v1482
	var v1483 int32
	_ = v1483
	var v1490 int32
	_ = v1490
	var v1493 int32
	_ = v1493
	var v1494 int64
	_ = v1494
	var v1495 int64
	_ = v1495
	var v1508 int32
	_ = v1508
	var v1528 int64
	_ = v1528
	var v1529 int64
	_ = v1529
	var v1533 int64
	_ = v1533
	var v1537 int32
	_ = v1537
	var v1551 int32
	_ = v1551
	var v1600 int32
	_ = v1600
	var v1603 int32
	_ = v1603
	var v1615 int32
	_ = v1615
	var v1618 int32
	_ = v1618
	var v1624 int32
	_ = v1624
	var v1629 int32
	_ = v1629
	var v1637 int32
	_ = v1637
	var v1656 int32
	_ = v1656
	var v1665 int32
	_ = v1665
	var v1675 int32
	_ = v1675
	var v1681 int32
	_ = v1681
	var v1683 int32
	_ = v1683
	var v1684 int32
	_ = v1684
	var v1685 int32
	_ = v1685
	var v1688 int32
	_ = v1688
	var v1690 int32
	_ = v1690
	var v1693 int32
	_ = v1693
	var v1702 int32
	_ = v1702
	var v1704 int32
	_ = v1704
	var v1706 int32
	_ = v1706
	var v1710 int32
	_ = v1710
	var v1711 int32
	_ = v1711
	var v1715 int32
	_ = v1715
	var v1721 int32
	_ = v1721
	var v1723 int32
	_ = v1723
	var v1729 int32
	_ = v1729
	var v1730 int32
	_ = v1730
	var v1732 int32
	_ = v1732
	var v1737 int32
	_ = v1737
	var v1745 int32
	_ = v1745
	var v1746 int32
	_ = v1746
	var v1749 int32
	_ = v1749
	var v1751 int32
	_ = v1751
	var v1754 int32
	_ = v1754
	var v1762 int32
	_ = v1762
	var v1763 int32
	_ = v1763
	var v1764 int32
	_ = v1764
	var v1767 int32
	_ = v1767
	var v1775 int32
	_ = v1775
	var v1779 int32
	_ = v1779
	var v1785 int32
	_ = v1785
	var v1787 int32
	_ = v1787
	var v1788 int32
	_ = v1788
	var v1793 int32
	_ = v1793
	var v1794 int32
	_ = v1794
	var v1797 int32
	_ = v1797
	var v1798 int32
	_ = v1798
	var v1801 int32
	_ = v1801
	var v1804 int64
	_ = v1804
	var v1805 int32
	_ = v1805
	var v1806 int32
	_ = v1806
	var v1808 int32
	_ = v1808
	var v1813 int32
	_ = v1813
	var v1814 int32
	_ = v1814
	var v1815 int32
	_ = v1815
	var v1816 int32
	_ = v1816
	var v1818 int32
	_ = v1818
	var v1821 int32
	_ = v1821
	var v1823 int32
	_ = v1823
	var v1824 int32
	_ = v1824
	var v1827 int32
	_ = v1827
	var v1832 int32
	_ = v1832
	var v1834 int32
	_ = v1834
	var v1835 int32
	_ = v1835
	var v1836 int32
	_ = v1836
	var v1837 int32
	_ = v1837
	var v1841 int32
	_ = v1841
	var v1844 int32
	_ = v1844
	var v1845 int32
	_ = v1845
	var v1846 int32
	_ = v1846
	var v1848 int32
	_ = v1848
	var v1853 int32
	_ = v1853
	var v1856 int32
	_ = v1856
	var v1860 int32
	_ = v1860
	var v1864 int32
	_ = v1864
	var v1867 int64
	_ = v1867
	var v1868 int32
	_ = v1868
	var v1872 int64
	_ = v1872
	var v1873 int32
	_ = v1873
	var v1875 int64
	_ = v1875
	var v1877 int64
	_ = v1877
	var v1880 int32
	_ = v1880
	var v1882 int32
	_ = v1882
	var v1883 int32
	_ = v1883
	var v1890 int32
	_ = v1890
	var v1893 int32
	_ = v1893
	var v1895 int32
	_ = v1895
	var v1897 int32
	_ = v1897
	var v1904 int32
	_ = v1904
	var v1930 int32
	_ = v1930
	var v1951 int32
	_ = v1951
	var v1975 int32
	_ = v1975
	var v1986 int32
	_ = v1986
	var v1993 int32
	_ = v1993
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
	v74 = F_XLogGetFakeLSN(m, v35)
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
	v93 = F_read_stream_begin_relation(m, int32(13), v87, v35, v84, int32(3), v33+int32(8), v84)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
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
	F_read_stream_end(m, v93)
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L1
	} else {
		goto L136
	}
L21:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	if base.Ui32(v140) < base.Ui32(v139) {
		goto L29
	} else {
		goto L30
	}
L22:
	;
	v128 = F_RelationGetNumberOfBlocksInFork(m, v35, int32(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
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
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L26
	}
L25:
	;
	v139 = v128
	goto L21
L26:
	;
	v134 = F_RelationGetNumberOfBlocksInFork(m, v35, int32(0))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_UnlockRelationForExtension(m, v35, int32(7))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v139 = v134
	goto L21
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+12)) = v139
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
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L35
	}
L33:
	;
	F_read_stream_reset(m, v93)
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L1
	} else {
		goto L135
	}
L34:
	;
	goto L33
L35:
	;
	v177 = F_read_stream_next_buffer(m, v93, int32(0))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	if v177 == int32(0) {
		goto L34
	} else {
		goto L37
	}
L37:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v177 < int32(0) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v201 = base.I64_extend_i32_u(v200)
	v216 = v177
	v221 = v200
	goto L42
L39:
	;
	v185 = *(*int32)(unsafe.Add(mBase, _c_F_gistvacuumscan[2]))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v185+(v177^int32(-1))*int32(56))+16))
	v200 = v191
	goto L38
L40:
	;
	goto L41
L41:
	;
	v193 = *(*int32)(unsafe.Add(mBase, _c_F_gistvacuumscan[3]))
	v194 = int32(56)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v193+v177*v194-v194)+16))
	v200 = v199
	goto L38
L42:
	;
	F_LockBufferInternal(m, v216, int32(3))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	if v216 < int32(0) {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	if v274 != 0 {
		goto L56
	} else {
		goto L57
	}
L46:
	;
	v254 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v253)+14)))
	if v254 == int32(0) {
		v274 = int32(1)
		goto L45
	} else {
		goto L50
	}
L47:
	;
	v239 = *(*int32)(unsafe.Add(mBase, _c_F_gistvacuumscan[4]))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v239+(v216^int32(-1))<<(uint(int32(2))%32))))
	v253 = v245
	goto L46
L48:
	;
	goto L49
L49:
	;
	v247 = *(*int32)(unsafe.Add(mBase, _c_F_gistvacuumscan[5]))
	v253 = v247 + v216<<(uint(int32(13))%32) + int32(-8192)
	goto L46
L50:
	;
	v257 = int32(0)
	v258 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v253)+16)))
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253+v258)+12)))
	if v260&int32(2) == v257 {
		v274 = v257
		goto L45
	} else {
		goto L51
	}
L51:
	;
	v266 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v253)+12)))
	if base.Ui32(int32(32)) <= base.Ui32(v266) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v269 = *(*int64)(unsafe.Add(mBase, uint32(v253)+24))
	v271 = v269
	goto L54
L53:
	;
	v271 = int64(3)
	goto L54
L54:
	;
	v272 = F_GlobalVisCheckRemovableFullXid(m, int32(0), v271)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v274 = v272
	goto L45
L56:
	;
	F_RecordFreeIndexPage(m, v181, v221)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v287 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v253)+16)))
	v288 = v253 + v287
	v289 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v288)+12)))
	if v289&int32(2) != 0 {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v278 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v277 + v278
	v281 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v281 + v278
	F_UnlockReleaseBuffer(m, v216)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	goto L32
L61:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v292 + int32(1)
	F_UnlockReleaseBuffer(m, v216)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L1
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	if v289&int32(1) != 0 {
		goto L67
	} else {
		goto L68
	}
L64:
	;
	goto L32
L65:
	;
	F_UnlockReleaseBuffer(m, v216)
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L1
	} else {
		goto L131
	}
L66:
	;
	v634 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = base.F64_add(v634, base.F64_convert_i32_u(v480))
	goto L65
L67:
	;
	v300 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v253)+12)))
	if v289&int32(8) == int32(0) {
		goto L71
	} else {
		goto L72
	}
L68:
	;
	goto L69
L69:
	;
	v514 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v253)+12)))
	if base.Ui32(v514) < base.Ui32(int32(25)) {
		goto L112
	} else {
		goto L113
	}
L70:
	;
	if base.Ui32(v300) < base.Ui32(int32(25)) {
		goto L81
	} else {
		goto L82
	}
L71:
	;
	v312 = *(*int64)(unsafe.Add(mBase, uint32(v288)))
	if base.Ui64(base.I64_rotl(v312, int64(32))) <= base.Ui64(v76) {
		v325 = int32(-1)
		goto L70
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v316 = int32(-1)
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v288)+8))
	if base.Ui32(v200) <= base.Ui32(v318) {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	goto L73
L75:
	;
	v320 = v316
	goto L77
L76:
	;
	v320 = v318
	goto L77
L77:
	;
	if v318 == int32(-1) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v323 = v316
	goto L80
L79:
	;
	v323 = v320
	goto L80
L80:
	;
	v325 = v323
	goto L70
L81:
	;
	v327 = int32(0)
	goto L83
L82:
	;
	v327 = int32(base.Ui32(v300+int32(_a_F_gistvacuumscan_4)) >> (uint(int32(2)) % 32))
	goto L83
L83:
	;
	if l2 == int32(0) {
		v456 = v327
		goto L85
	} else {
		goto L86
	}
L84:
	;
	if v200 != v221 {
		goto L65
	} else {
		goto L110
	}
L85:
	;
	v480 = v456 & int32(_a_F_gistvacuumscan_5)
	if v480 != 0 {
		goto L66
	} else {
		goto L109
	}
L86:
	;
	v331 = v327 & int32(_a_F_gistvacuumscan_5)
	if v331 == int32(0) {
		goto L84
	} else {
		goto L87
	}
L87:
	;
	v343 = int32(0)
	v350 = int32(1)
	goto L88
L88:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v253+int32(20)+v350&int32(_a_F_gistvacuumscan_5)<<(uint(int32(2))%32))))
	v377 = m.T0[l2].(func(*base.Module, int32, int32) int32)(m, v253+v373&int32(_a_F_gistvacuumscan_6), l3)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L1
	} else {
		goto L90
	}
L89:
	;
	if v387 <= int32(0) {
		v456 = v327
		goto L85
	} else {
		goto L95
	}
L90:
	;
	if v377 != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v381 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v33+int32(16)+v343<<(uint(v381)%32)))) = uint16(v350)
	v387 = v343 + v381
	goto L93
L92:
	;
	v387 = v343
	goto L93
L93:
	;
	v389 = v350 + int32(1)
	if base.Ui32(v389&int32(_a_F_gistvacuumscan_5)) <= base.Ui32(v331) {
		v343 = v387
		v350 = v389
		goto L88
	} else {
		goto L94
	}
L94:
	;
	goto L89
L95:
	;
	v395 = int32(_a_F_gistvacuumscan_7)
	v397 = *(*int32)(unsafe.Add(mBase, _c_F_gistvacuumscan[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_gistvacuumscan[6])) = v397 + int32(1)
	F_MarkBufferDirty(m, v216)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	F_PageIndexMultiDelete(m, v253, v33+int32(16), v387)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v181)+48))
	v408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v407)+118)))
	if v408 != int32(112) {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v253))) = v430
	v432 = int32(_a_F_gistvacuumscan_7)
	v434 = *(*int32)(unsafe.Add(mBase, _c_F_gistvacuumscan[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_gistvacuumscan[6])) = v434 - int32(1)
	v438 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*float64)(unsafe.Add(mBase, uint32(l1)+16)) = base.F64_add(v438, base.F64_convert_i32_u(v387))
	v442 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v253)+12)))
	if base.Ui32(v442) < base.Ui32(int32(25)) {
		goto L84
	} else {
		goto L108
	}
L99:
	;
	v426 = F_XLogGetFakeLSN(m, v181)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L1
	} else {
		goto L107
	}
L100:
	;
	v412 = *(*int32)(unsafe.Add(mBase, _c_F_gistvacuumscan[1]))
	if v412 <= int32(0) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v181)+32))
	if v415 != 0 {
		goto L99
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	v419 = int32(0)
	v422 = F_gistXLogUpdate(m, v216, v33+int32(16), v387, v419, v419, v419)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L1
	} else {
		goto L106
	}
L104:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v181)+40))
	if v416 != 0 {
		goto L99
	} else {
		goto L105
	}
L105:
	;
	goto L103
L106:
	;
	v430 = base.I64_rotl(v422, int64(32))
	goto L98
L107:
	;
	v430 = base.I64_rotl(v426, int64(32))
	goto L98
L108:
	;
	v456 = int32(base.Ui32(v442+int32(_a_F_gistvacuumscan_4)) >> (uint(int32(2)) % 32))
	goto L85
L109:
	;
	goto L84
L110:
	;
	F_intset_add_member(m, v58, v201)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	goto L65
L112:
	;
	if v200 == v221 {
		goto L126
	} else {
		goto L127
	}
L113:
	;
	v518 = v514 + int32(_a_F_gistvacuumscan_4)
	if v518&int32(_a_F_gistvacuumscan_8) == int32(0) {
		goto L112
	} else {
		goto L114
	}
L114:
	;
	v535 = int32(1)
	goto L115
L115:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v253+int32(20)+v535<<(uint(int32(2))%32))))
	v567 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v253+v563&int32(_a_F_gistvacuumscan_6))+4)))
	if v567 != int32(_a_F_gistvacuumscan_9) {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	goto L112
L117:
	;
	if v535 != int32(base.Ui32(v518)>>(uint(int32(2))%32))&int32(_a_F_gistvacuumscan_5) {
		v535 = v535 + int32(1)
		goto L115
	} else {
		goto L125
	}
L118:
	;
	v572 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	if v572 == int32(0) {
		goto L117
	} else {
		goto L120
	}
L120:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v181)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v576 + int32(4)
	F_errmsg(m, int32(_a_F_gistvacuumscan_10), v33)
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	v585 = F_errdetail(m, int32(_a_F_gistvacuumscan_11), int32(0))
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	F_errhint(m, int32(_a_F_gistvacuumscan_12), int32(0))
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	F_errfinish(m, int32(_a_F_gistvacuumscan_13), int32(465), int32(_a_F_gistvacuumscan_14))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
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
	F_intset_add_member(m, v56, v201)
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L1
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	F_UnlockReleaseBuffer(m, v216)
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
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
	if v325 == int32(-1) {
		goto L32
	} else {
		goto L132
	}
L132:
	;
	F_vacuum_delay_point(m, int32(0))
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	v675 = int32(0)
	v677 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v678 = F_ReadBufferExtended(m, v181, v675, v325, v675, v677)
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	v216 = v678
	v221 = v325
	goto L42
L135:
	;
	goto L19
L136:
	;
	v684 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v684 != 0 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	F_FreeSpaceMapVacuum(m, v35)
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L1
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v139
	v688 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v689 = *(*int64)(unsafe.Add(mBase, uint32(v58)+16))
	v690 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+3948)) = uint8(v690)
	v692 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v56)+3968)) = v692
	*(*int64)(unsafe.Add(mBase, uint32(v56)+3956)) = int64(0)
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v56)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+3964)) = v696
	*(*int32)(unsafe.Add(mBase, uint32(v56)+3952)) = v56 + int32(3976)
	v701 = base.I32_wrap_i64(v689)
	if v701 == v692 {
		v1975 = v33
		v1986 = v50
		goto L141
	} else {
		goto L142
	}
L140:
	;
	goto L139
L141:
	;
	F_MemoryContextDelete(m, v1986)
	mBase = m.M
	v1993 = m.ExcPending
	if v1993 != 0 {
		goto L1
	} else {
		goto L320
	}
L142:
	;
	v704 = l0
	v705 = l1
	v710 = v56
	v717 = v33
	v718 = v688
	v723 = v701
	v726 = v58
	v728 = v50
	goto L143
L143:
	;
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v710)+3960))
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v710)+3956))
	if v736 < v737 {
		v1047 = v736
		goto L146
	} else {
		goto L147
	}
L144:
	;
	v1975 = v717
	v1986 = v728
	goto L141
L145:
	;
	if v1114 == int32(0) {
		v1975 = v717
		v1986 = v728
		goto L141
	} else {
		goto L179
	}
L146:
	;
	v1074 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v710)+3960)) = v1047 + v1074
	v1077 = *(*int32)(unsafe.Add(mBase, uint32(v710)+3952))
	v1081 = *(*int64)(unsafe.Add(mBase, uint32(v1077+v1047<<(uint(int32(3))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v717)+uint32(_c_F_gistvacuumscan[7]))) = v1081
	v1114 = v1074
	goto L145
L147:
	;
	v742 = v710 + int32(3984)
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v710)+3964))
	v748 = v745
	v749 = v736
	v751 = v737
	goto L148
L148:
	;
	if v748 != 0 {
		goto L151
	} else {
		goto L152
	}
L149:
	;
	v1047 = v1040
	goto L146
L150:
	;
	if v1041 <= v1040 {
		v748 = v1039
		v749 = v1040
		v751 = v1041
		goto L148
	} else {
		goto L178
	}
L151:
	;
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v710)+3968))
	v777 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v748)+2)))
	if v776 < v777 {
		goto L154
	} else {
		goto L155
	}
L152:
	;
	goto L153
L153:
	;
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(v710)+3952))
	if v710+int32(3976) == v1025 {
		goto L175
	} else {
		goto L176
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v710)+3968)) = v776 + int32(1)
	v784 = v748 + v776<<(uint(int32(4))%32)
	v785 = *(*int64)(unsafe.Add(mBase, uint32(v784)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v710)+3976)) = v785
	v787 = *(*int64)(unsafe.Add(mBase, uint32(v784)+16))
	if v787 != int64(1152921504606846975) {
		goto L157
	} else {
		goto L158
	}
L155:
	;
	goto L156
L156:
	;
	v1021 = *(*int32)(unsafe.Add(mBase, uint32(v748)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v710)+3968)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v710)+3964)) = v1021
	v1039 = v1021
	v1040 = v749
	v1041 = v751
	goto L150
L157:
	;
	v790 = int64(-1)
	v793 = base.I32_wrap_i64(int64(base.Ui64(v787) >> (uint(int64(60)) % 64)))
	v794 = int32(1)
	v795 = v793 << (uint(v794) % 32)
	v796 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v795)+uint32(_c_F_gistvacuumscan[8]))))
	v799 = v790<<(uint(v796)%64) ^ v790
	v801 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v795)+uint32(_c_F_gistvacuumscan[9]))))
	if base.Ui32(v801) <= base.Ui32(v794) {
		goto L160
	} else {
		goto L161
	}
L158:
	;
	v1018 = int32(1)
	goto L159
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v710)+3956)) = v1018
	v1047 = int32(0)
	goto L146
L160:
	;
	v804 = v794
	goto L162
L161:
	;
	v804 = v801
	goto L162
L162:
	;
	v806 = v804 & int32(3)
	if base.Ui32(v793-int32(13)) < base.Ui32(int32(4)) {
		goto L165
	} else {
		goto L166
	}
L163:
	;
	v1018 = v801 + int32(1)
	goto L159
L164:
	;
	v913 = v883
	v916 = int32(0)
	v936 = v906
	v939 = v909
	goto L172
L165:
	;
	v883 = int32(0)
	v906 = v787
	v909 = v785
	goto L164
L166:
	;
	goto L167
L167:
	;
	v815 = int32(0)
	v819 = v815
	v829 = v815
	v842 = v787
	v845 = v785
	goto L168
L168:
	;
	v849 = v742 + v819<<(uint(int32(3))%32)
	v852 = int64(1)
	v853 = v845 + v842&v799 + v852
	*(*int64)(unsafe.Add(mBase, uint32(v849))) = v853
	v855 = int64(base.Ui64(v842) >> (uint(v796) % 64))
	v859 = v853 + v855&v799 + v852
	*(*int64)(unsafe.Add(mBase, uint32(v849)+8)) = v859
	v861 = int64(base.Ui64(v855) >> (uint(v796) % 64))
	v865 = v859 + v861&v799 + v852
	*(*int64)(unsafe.Add(mBase, uint32(v849)+16)) = v865
	v867 = int64(base.Ui64(v861) >> (uint(v796) % 64))
	v871 = v865 + v867&v799 + v852
	*(*int64)(unsafe.Add(mBase, uint32(v849)+24)) = v871
	v873 = int32(4)
	v874 = v819 + v873
	v875 = int64(base.Ui64(v867) >> (uint(v796) % 64))
	v877 = v829 + v873
	if v877 != v804&int32(252) {
		v819 = v874
		v829 = v877
		v842 = v875
		v845 = v871
		goto L168
	} else {
		goto L170
	}
L169:
	;
	if v806 == int32(0) {
		goto L163
	} else {
		goto L171
	}
L170:
	;
	goto L169
L171:
	;
	v883 = v874
	v906 = v875
	v909 = v871
	goto L164
L172:
	;
	v947 = v939 + v936&v799 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v742+v913<<(uint(int32(3))%32)))) = v947
	v949 = int32(1)
	v953 = v916 + v949
	if v953 != v806 {
		v913 = v913 + v949
		v916 = v953
		v936 = int64(base.Ui64(v936) >> (uint(v796) % 64))
		v939 = v947
		goto L172
	} else {
		goto L174
	}
L173:
	;
	goto L163
L174:
	;
	goto L173
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v710)+3952)) = v710 + int32(88)
	v1028 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v710)+3960)) = v1028
	v1031 = *(*int32)(unsafe.Add(mBase, uint32(v710)+3944))
	*(*int32)(unsafe.Add(mBase, uint32(v710)+3956)) = v1031
	v1039 = v1028
	v1040 = v1028
	v1041 = v1031
	goto L150
L176:
	;
	goto L177
L177:
	;
	v1034 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v710)+3948)) = uint8(v1034)
	*(*int64)(unsafe.Add(mBase, uint32(v717)+uint32(_c_F_gistvacuumscan[7]))) = int64(0)
	v1114 = v1034
	goto L145
L178:
	;
	goto L149
L179:
	;
	v1117 = int32(0)
	v1118 = *(*int32)(unsafe.Add(mBase, uint32(v717)+uint32(_c_F_gistvacuumscan[7])))
	v1120 = *(*int32)(unsafe.Add(mBase, uint32(v704)+24))
	v1121 = F_ReadBufferExtended(m, v718, v1117, v1118, v1117, v1120)
	mBase = m.M
	v1122 = m.ExcPending
	if v1122 != 0 {
		goto L1
	} else {
		goto L180
	}
L180:
	;
	F_LockBufferInternal(m, v1121, int32(1))
	mBase = m.M
	v1125 = m.ExcPending
	if v1125 != 0 {
		goto L1
	} else {
		goto L181
	}
L181:
	;
	v1126 = int32(0)
	v1127 = base.B2i32(v1126 <= v1121)
	if v1127 == v1126 {
		goto L185
	} else {
		goto L186
	}
L182:
	;
	if v1951 != 0 {
		v723 = v1951
		goto L143
	} else {
		goto L319
	}
L183:
	;
	v1156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1145)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v1156) {
		goto L194
	} else {
		goto L195
	}
L184:
	;
	v1146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1145)+14)))
	if v1146 != 0 {
		goto L188
	} else {
		goto L189
	}
L185:
	;
	v1131 = *(*int32)(unsafe.Add(mBase, _c_F_gistvacuumscan[4]))
	v1137 = *(*int32)(unsafe.Add(mBase, uint32(v1131+(v1121^int32(-1))<<(uint(int32(2))%32))))
	v1145 = v1137
	goto L184
L186:
	;
	goto L187
L187:
	;
	v1139 = *(*int32)(unsafe.Add(mBase, _c_F_gistvacuumscan[5]))
	v1145 = v1139 + v1121<<(uint(int32(13))%32) + int32(-8192)
	goto L184
L188:
	;
	v1147 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1145)+16)))
	v1149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1145+v1147)+12)))
	if v1149&int32(3) == int32(0) {
		goto L183
	} else {
		goto L191
	}
L189:
	;
	goto L190
L190:
	;
	F_UnlockReleaseBuffer(m, v1121)
	mBase = m.M
	v1155 = m.ExcPending
	if v1155 != 0 {
		goto L1
	} else {
		goto L192
	}
L191:
	;
	goto L190
L192:
	;
	v1951 = v723
	goto L182
L193:
	;
	F_ReleaseBuffer(m, v1121)
	mBase = m.M
	v1930 = m.ExcPending
	if v1930 != 0 {
		goto L1
	} else {
		goto L318
	}
L194:
	;
	v1164 = int32(base.Ui32(v1156+int32(_a_F_gistvacuumscan_4)) >> (uint(int32(2)) % 32))
	goto L196
L195:
	;
	v1164 = int32(0)
	goto L196
L196:
	;
	v1166 = v1164 & int32(_a_F_gistvacuumscan_5)
	if base.Ui32(v1166) <= base.Ui32(int32(1)) {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	F_UnlockBuffer(m, v1121)
	mBase = m.M
	v1170 = m.ExcPending
	if v1170 != 0 {
		goto L1
	} else {
		goto L200
	}
L198:
	;
	goto L199
L199:
	;
	v1172 = int32(1)
	v1183 = int32(0)
	v1190 = v1172
	goto L201
L200:
	;
	v1904 = int32(0)
	goto L193
L201:
	;
	v1213 = *(*int32)(unsafe.Add(mBase, uint32(v1145+int32(20)+v1190&int32(_a_F_gistvacuumscan_5)<<(uint(int32(2))%32))))
	v1216 = v1145 + v1213&int32(_a_F_gistvacuumscan_6)
	v1217 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1216))))
	v1220 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1216)+2)))
	v1221 = v1217<<(uint(int32(16))%32) | v1220
	v1222 = base.I64_extend_i32_u(v1221)
	v1223 = int32(0)
	v1224 = *(*int32)(unsafe.Add(mBase, uint32(v726)+3944))
	if v1224 <= v1223 {
		goto L204
	} else {
		goto L205
	}
L202:
	;
	F_UnlockBuffer(m, v1121)
	mBase = m.M
	v1624 = m.ExcPending
	if v1624 != 0 {
		goto L1
	} else {
		goto L263
	}
L203:
	;
	if v1600 != 0 {
		goto L259
	} else {
		goto L260
	}
L204:
	;
	v1283 = *(*int32)(unsafe.Add(mBase, uint32(v726)+36))
	if v1283 == int32(0) {
		v1551 = v1223
		goto L217
	} else {
		goto L218
	}
L205:
	;
	v1227 = *(*int64)(unsafe.Add(mBase, uint32(v726)+88))
	if base.Ui64(v1222) < base.Ui64(v1227) {
		goto L204
	} else {
		goto L206
	}
L206:
	;
	v1230 = v726 + int32(88)
	v1235 = v1224
	v1240 = int32(0)
	goto L207
L207:
	;
	v1264 = base.I32_div_s(v1235-v1240, int32(2))
	v1265 = v1264 + v1240
	v1269 = *(*int64)(unsafe.Add(mBase, uint32(v1230+v1265<<(uint(int32(3))%32))))
	v1270 = base.B2i32(base.Ui64(v1269) < base.Ui64(v1222))
	if base.Ui64(v1269) < base.Ui64(v1222) {
		goto L209
	} else {
		goto L210
	}
L208:
	;
	if v1224 <= v1274 {
		v1600 = int32(0)
		goto L203
	} else {
		goto L216
	}
L209:
	;
	v1271 = v1235
	goto L211
L210:
	;
	v1271 = v1265
	goto L211
L211:
	;
	if base.Ui64(v1269) < base.Ui64(v1222) {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	v1274 = v1265 + int32(1)
	goto L214
L213:
	;
	v1274 = v1240
	goto L214
L214:
	;
	if v1274 < v1271 {
		v1235 = v1271
		v1240 = v1274
		goto L207
	} else {
		goto L215
	}
L215:
	;
	goto L208
L216:
	;
	v1281 = *(*int64)(unsafe.Add(mBase, uint32(v1230+v1274<<(uint(int32(3))%32))))
	v1600 = base.B2i32(v1281 == v1222)
	goto L203
L217:
	;
	v1600 = v1551
	goto L203
L218:
	;
	v1286 = *(*int32)(unsafe.Add(mBase, uint32(v726)+32))
	v1288 = v1286 - int32(1)
	if int32(0) < v1288 {
		goto L219
	} else {
		goto L220
	}
L219:
	;
	v1307 = v1288
	v1312 = v1283
	goto L222
L220:
	;
	v1403 = v1283
	goto L221
L221:
	;
	v1412 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1403)+2)))
	if v1412 == int32(0) {
		v1551 = v1223
		goto L217
	} else {
		goto L236
	}
L222:
	;
	v1321 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1312)+2)))
	if v1321 == int32(0) {
		v1551 = v1223
		goto L217
	} else {
		goto L224
	}
L223:
	;
	v1403 = v1379
	goto L221
L224:
	;
	v1330 = v1321
	v1335 = int32(0)
	goto L225
L225:
	;
	v1359 = base.I32_div_s(v1330-v1335, int32(2))
	v1360 = v1359 + v1335
	v1364 = *(*int64)(unsafe.Add(mBase, uint32(v1312+int32(8)+v1360<<(uint(int32(3))%32))))
	v1365 = base.B2i32(base.Ui64(v1222) < base.Ui64(v1364))
	if base.Ui64(v1222) < base.Ui64(v1364) {
		goto L227
	} else {
		goto L228
	}
L226:
	;
	v1371 = int32(0)
	if v1369 == v1371 {
		v1600 = v1371
		goto L203
	} else {
		goto L234
	}
L227:
	;
	v1366 = v1360
	goto L229
L228:
	;
	v1366 = v1330
	goto L229
L229:
	;
	if base.Ui64(v1222) < base.Ui64(v1364) {
		goto L230
	} else {
		goto L231
	}
L230:
	;
	v1369 = v1335
	goto L232
L231:
	;
	v1369 = v1360 + int32(1)
	goto L232
L232:
	;
	if v1369 < v1366 {
		v1330 = v1366
		v1335 = v1369
		goto L225
	} else {
		goto L233
	}
L233:
	;
	goto L226
L234:
	;
	v1374 = int32(1)
	v1379 = *(*int32)(unsafe.Add(mBase, uint32(v1312+v1369<<(uint(int32(2))%32))+516))
	if v1374 < v1307 {
		v1307 = v1307 - v1374
		v1312 = v1379
		goto L222
	} else {
		goto L235
	}
L235:
	;
	goto L223
L236:
	;
	v1416 = v1403 + int32(8)
	v1421 = v1412
	v1426 = int32(0)
	goto L237
L237:
	;
	v1450 = base.I32_div_s(v1421-v1426, int32(2))
	v1451 = v1450 + v1426
	v1455 = *(*int64)(unsafe.Add(mBase, uint32(v1416+v1451<<(uint(int32(4))%32))))
	v1456 = base.B2i32(base.Ui64(v1222) < base.Ui64(v1455))
	if base.Ui64(v1222) < base.Ui64(v1455) {
		goto L239
	} else {
		goto L240
	}
L238:
	;
	if v1460 == int32(0) {
		v1551 = v1223
		goto L217
	} else {
		goto L246
	}
L239:
	;
	v1457 = v1451
	goto L241
L240:
	;
	v1457 = v1421
	goto L241
L241:
	;
	if base.Ui64(v1222) < base.Ui64(v1455) {
		goto L242
	} else {
		goto L243
	}
L242:
	;
	v1460 = v1426
	goto L244
L243:
	;
	v1460 = v1451 + int32(1)
	goto L244
L244:
	;
	if v1460 < v1457 {
		v1421 = v1457
		v1426 = v1460
		goto L237
	} else {
		goto L245
	}
L245:
	;
	goto L238
L246:
	;
	v1467 = v1416 + v1460<<(uint(int32(4))%32)
	v1470 = *(*int64)(unsafe.Add(mBase, uint32(v1467-int32(16))))
	if v1470 == v1222 {
		v1600 = int32(1)
		goto L203
	} else {
		goto L247
	}
L247:
	;
	v1474 = *(*int64)(unsafe.Add(mBase, uint32(v1467-int32(8))))
	if v1474 == int64(1152921504606846975) {
		v1551 = v1223
		goto L217
	} else {
		goto L248
	}
L248:
	;
	v1482 = base.I32_wrap_i64(int64(base.Ui64(v1474)>>(uint(int64(60))%64))) << (uint(int32(1)) % 32)
	v1483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1482)+uint32(_c_F_gistvacuumscan[9]))))
	if base.Ui64(v1474) <= base.Ui64(int64(2305843009213693951)) {
		v1600 = base.B2i32(base.Ui64(v1222-v1470) <= base.Ui64(base.I64_extend_i32_u(v1483)&int64(255)))
		goto L203
	} else {
		goto L249
	}
L249:
	;
	v1490 = int32(1)
	if base.Ui32(v1483) <= base.Ui32(v1490) {
		goto L250
	} else {
		goto L251
	}
L250:
	;
	v1493 = v1490
	goto L252
L251:
	;
	v1493 = v1483
	goto L252
L252:
	;
	v1494 = int64(-1)
	v1495 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1482)+uint32(_c_F_gistvacuumscan[8]))))
	v1508 = int32(0)
	v1528 = v1474
	v1529 = v1470
	goto L253
L253:
	;
	v1533 = v1529 + (v1494<<(uint(v1495)%64)^v1494)&v1528 + int64(1)
	if base.Ui64(v1533) < base.Ui64(v1222) {
		goto L255
	} else {
		goto L256
	}
L254:
	;
	v1551 = base.B2i32(v1222 == v1533)
	goto L217
L255:
	;
	v1537 = v1508 + int32(1)
	if v1537 != v1493 {
		v1508 = v1537
		v1528 = int64(base.Ui64(v1528) >> (uint(v1495) % 64))
		v1529 = v1533
		goto L253
	} else {
		goto L258
	}
L256:
	;
	goto L257
L257:
	;
	goto L254
L258:
	;
	v1551 = v1223
	goto L217
L259:
	;
	v1603 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v717+int32(_a_F_gistvacuumscan_15)+v1183<<(uint(v1603)%32)))) = uint16(v1190)
	*(*int32)(unsafe.Add(mBase, uint32(v717+int32(16)+v1183<<(uint(int32(2))%32)))) = v1221
	v1615 = v1183 + v1603
	goto L261
L260:
	;
	v1615 = v1183
	goto L261
L261:
	;
	v1618 = v1190 + int32(1)
	if base.B2i32(v1615 < v1166-v1172)&base.B2i32(base.Ui32(v1618&int32(_a_F_gistvacuumscan_5)) <= base.Ui32(v1166)) != 0 {
		v1183 = v1615
		v1190 = v1618
		goto L201
	} else {
		goto L262
	}
L262:
	;
	goto L202
L263:
	;
	if v1615 <= int32(0) {
		v1904 = v1615
		goto L193
	} else {
		goto L264
	}
L264:
	;
	v1629 = int32(0)
	v1637 = v1629
	v1656 = v1629
	goto L265
L265:
	;
	v1665 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1145)+12)))
	if base.B2i32(base.Ui32(int32(25)) <= base.Ui32(v1665))&base.B2i32((v1665+int32(_a_F_gistvacuumscan_4))&int32(_a_F_gistvacuumscan_8) == int32(4)) != 0 {
		v1904 = v1615
		goto L193
	} else {
		goto L267
	}
L266:
	;
	v1904 = v1615
	goto L193
L267:
	;
	v1675 = int32(0)
	v1681 = *(*int32)(unsafe.Add(mBase, uint32(v717+int32(16)+v1637<<(uint(int32(2))%32))))
	v1683 = *(*int32)(unsafe.Add(mBase, uint32(v704)+24))
	v1684 = F_ReadBufferExtended(m, v718, v1675, v1681, v1675, v1683)
	mBase = m.M
	v1685 = m.ExcPending
	if v1685 != 0 {
		goto L1
	} else {
		goto L268
	}
L268:
	;
	F_LockBufferInternal(m, v1684, int32(3))
	mBase = m.M
	v1688 = m.ExcPending
	if v1688 != 0 {
		goto L1
	} else {
		goto L269
	}
L269:
	;
	F_gistcheckpage(m, v718, v1684)
	mBase = m.M
	v1690 = m.ExcPending
	if v1690 != 0 {
		goto L1
	} else {
		goto L270
	}
L270:
	;
	F_LockBufferInternal(m, v1121, int32(3))
	mBase = m.M
	v1693 = m.ExcPending
	if v1693 != 0 {
		goto L1
	} else {
		goto L271
	}
L271:
	;
	if v1127 == int32(0) {
		goto L273
	} else {
		goto L274
	}
L272:
	;
	v1711 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v717+int32(_a_F_gistvacuumscan_15)+v1637<<(uint(int32(1))%32)))))
	if v1684 < int32(0) {
		goto L278
	} else {
		goto L279
	}
L273:
	;
	v1702 = *(*int32)(unsafe.Add(mBase, _c_F_gistvacuumscan[4]))
	v1704 = *(*int32)(unsafe.Add(mBase, uint32(v1702+(v1121^int32(-1))<<(uint(int32(2))%32))))
	v1710 = v1704
	goto L272
L274:
	;
	goto L275
L275:
	;
	v1706 = *(*int32)(unsafe.Add(mBase, _c_F_gistvacuumscan[5]))
	v1710 = v1706 + v1121<<(uint(int32(13))%32) + int32(-8192)
	goto L272
L276:
	;
	F_UnlockBuffer(m, v1121)
	mBase = m.M
	v1893 = m.ExcPending
	if v1893 != 0 {
		goto L1
	} else {
		goto L315
	}
L277:
	;
	v1730 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1729)+16)))
	v1732 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1730+v1729)+12)))
	if v1732&int32(9) != int32(1) {
		v1890 = v1656
		goto L276
	} else {
		goto L281
	}
L278:
	;
	v1715 = *(*int32)(unsafe.Add(mBase, _c_F_gistvacuumscan[4]))
	v1721 = *(*int32)(unsafe.Add(mBase, uint32(v1715+(v1684^int32(-1))<<(uint(int32(2))%32))))
	v1729 = v1721
	goto L277
L279:
	;
	goto L280
L280:
	;
	v1723 = *(*int32)(unsafe.Add(mBase, _c_F_gistvacuumscan[5]))
	v1729 = v1723 + v1684<<(uint(int32(13))%32) + int32(-8192)
	goto L277
L281:
	;
	v1737 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1729)+12)))
	if (v1737+int32(_a_F_gistvacuumscan_4))&int32(_a_F_gistvacuumscan_8) != 0 {
		goto L282
	} else {
		goto L283
	}
L282:
	;
	v1745 = base.B2i32(base.Ui32(int32(25)) <= base.Ui32(v1737))
	goto L284
L283:
	;
	v1745 = int32(0)
	goto L284
L284:
	;
	if v1745 != 0 {
		v1890 = v1656
		goto L276
	} else {
		goto L285
	}
L285:
	;
	v1746 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1710)+14)))
	if v1746 == int32(0) {
		v1890 = v1656
		goto L276
	} else {
		goto L286
	}
L286:
	;
	v1749 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1710)+16)))
	v1751 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1710+v1749)+12)))
	if v1751&int32(3) != 0 {
		v1890 = v1656
		goto L276
	} else {
		goto L287
	}
L287:
	;
	v1754 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1710)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v1754) {
		goto L288
	} else {
		goto L289
	}
L288:
	;
	v1762 = int32(base.Ui32(v1754+int32(_a_F_gistvacuumscan_4)) >> (uint(int32(2)) % 32))
	goto L290
L289:
	;
	v1762 = int32(0)
	goto L290
L290:
	;
	v1763 = int32(_a_F_gistvacuumscan_5)
	v1764 = v1762 & v1763
	v1767 = (v1711 - v1656) & v1763
	if base.B2i32(base.Ui32(v1764) < base.Ui32(v1767))|base.B2i32(base.Ui32(v1764) < base.Ui32(int32(2))) != 0 {
		v1890 = v1656
		goto L276
	} else {
		goto L291
	}
L291:
	;
	v1775 = *(*int32)(unsafe.Add(mBase, uint32(v1710+v1767<<(uint(int32(2))%32))+20))
	if v1684 < int32(0) {
		goto L293
	} else {
		goto L294
	}
L292:
	;
	v1797 = v1710 + v1775&int32(_a_F_gistvacuumscan_6)
	v1798 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1797))))
	v1801 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1797)+2)))
	if v1794 != v1798<<(uint(int32(16))%32)|v1801 {
		v1890 = v1656
		goto L276
	} else {
		goto L296
	}
L293:
	;
	v1779 = *(*int32)(unsafe.Add(mBase, _c_F_gistvacuumscan[2]))
	v1785 = *(*int32)(unsafe.Add(mBase, uint32(v1779+(v1684^int32(-1))*int32(56))+16))
	v1794 = v1785
	goto L292
L294:
	;
	goto L295
L295:
	;
	v1787 = *(*int32)(unsafe.Add(mBase, _c_F_gistvacuumscan[3]))
	v1788 = int32(56)
	v1793 = *(*int32)(unsafe.Add(mBase, uint32(v1787+v1684*v1788-v1788)+16))
	v1794 = v1793
	goto L292
L296:
	;
	v1804 = F_ReadNextFullTransactionId(m)
	mBase = m.M
	v1805 = m.ExcPending
	if v1805 != 0 {
		goto L1
	} else {
		goto L297
	}
L297:
	;
	v1806 = int32(_a_F_gistvacuumscan_7)
	v1808 = *(*int32)(unsafe.Add(mBase, _c_F_gistvacuumscan[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_gistvacuumscan[6])) = v1808 + int32(1)
	F_MarkBufferDirty(m, v1684)
	mBase = m.M
	v1813 = m.ExcPending
	if v1813 != 0 {
		goto L1
	} else {
		goto L298
	}
L298:
	;
	v1814 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1729)+16)))
	v1815 = v1729 + v1814
	v1816 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1815)+12)))
	v1818 = v1816 | int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v1815)+12)) = uint16(v1818)
	*(*int64)(unsafe.Add(mBase, uint32(v1729)+24)) = v1804
	v1821 = int32(32)
	*(*uint16)(unsafe.Add(mBase, uint32(v1729)+12)) = uint16(v1821)
	v1823 = *(*int32)(unsafe.Add(mBase, uint32(v705)+24))
	v1824 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v705)+24)) = v1823 + v1824
	v1827 = *(*int32)(unsafe.Add(mBase, uint32(v705)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v705)+28)) = v1827 + v1824
	F_MarkBufferDirty(m, v1121)
	mBase = m.M
	v1832 = m.ExcPending
	if v1832 != 0 {
		goto L1
	} else {
		goto L299
	}
L299:
	;
	F_PageIndexTupleDelete(m, v1710, v1767)
	mBase = m.M
	v1834 = m.ExcPending
	if v1834 != 0 {
		goto L1
	} else {
		goto L300
	}
L300:
	;
	v1835 = *(*int32)(unsafe.Add(mBase, uint32(v704)))
	v1836 = *(*int32)(unsafe.Add(mBase, uint32(v1835)+48))
	v1837 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1836)+118)))
	if v1837 != int32(112) {
		goto L302
	} else {
		goto L303
	}
L301:
	;
	v1877 = base.I64_rotl(v1875, int64(32))
	*(*int64)(unsafe.Add(mBase, uint32(v1710))) = v1877
	*(*int64)(unsafe.Add(mBase, uint32(v1729))) = v1877
	v1880 = int32(_a_F_gistvacuumscan_7)
	v1882 = *(*int32)(unsafe.Add(mBase, _c_F_gistvacuumscan[6]))
	v1883 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_gistvacuumscan[6])) = v1882 - v1883
	v1890 = v1656 + v1883
	goto L276
L302:
	;
	v1872 = F_XLogGetFakeLSN(m, v1835)
	mBase = m.M
	v1873 = m.ExcPending
	if v1873 != 0 {
		goto L1
	} else {
		goto L314
	}
L303:
	;
	v1841 = *(*int32)(unsafe.Add(mBase, _c_F_gistvacuumscan[1]))
	if v1841 <= int32(0) {
		goto L304
	} else {
		goto L305
	}
L304:
	;
	v1844 = *(*int32)(unsafe.Add(mBase, uint32(v1835)+32))
	if v1844 != 0 {
		goto L302
	} else {
		goto L307
	}
L305:
	;
	goto L306
L306:
	;
	v1846 = m.G0
	v1848 = v1846 - int32(16)
	m.G0 = v1848
	*(*uint16)(unsafe.Add(mBase, uint32(v1848)+8)) = uint16(v1767)
	*(*int64)(unsafe.Add(mBase, uint32(v1848))) = v1804
	F_XLogBeginInsert(m)
	mBase = m.M
	v1853 = m.ExcPending
	if v1853 != 0 {
		goto L1
	} else {
		goto L309
	}
L307:
	;
	v1845 = *(*int32)(unsafe.Add(mBase, uint32(v1835)+40))
	if v1845 != 0 {
		goto L302
	} else {
		goto L308
	}
L308:
	;
	goto L306
L309:
	;
	F_XLogRegisterData(m, v1848, int32(10))
	mBase = m.M
	v1856 = m.ExcPending
	if v1856 != 0 {
		goto L1
	} else {
		goto L310
	}
L310:
	;
	F_XLogRegisterBuffer(m, int32(0), v1684, int32(8))
	mBase = m.M
	v1860 = m.ExcPending
	if v1860 != 0 {
		goto L1
	} else {
		goto L311
	}
L311:
	;
	F_XLogRegisterBuffer(m, int32(1), v1121, int32(8))
	mBase = m.M
	v1864 = m.ExcPending
	if v1864 != 0 {
		goto L1
	} else {
		goto L312
	}
L312:
	;
	v1867 = F_XLogInsert(m, int32(14), int32(96))
	mBase = m.M
	v1868 = m.ExcPending
	if v1868 != 0 {
		goto L1
	} else {
		goto L313
	}
L313:
	;
	m.G0 = v1848 + int32(16)
	v1875 = v1867
	goto L301
L314:
	;
	v1875 = v1872
	goto L301
L315:
	;
	F_UnlockReleaseBuffer(m, v1684)
	mBase = m.M
	v1895 = m.ExcPending
	if v1895 != 0 {
		goto L1
	} else {
		goto L316
	}
L316:
	;
	v1897 = v1637 + int32(1)
	if v1897 != v1615 {
		v1637 = v1897
		v1656 = v1890
		goto L265
	} else {
		goto L317
	}
L317:
	;
	goto L266
L318:
	;
	v1951 = v723 - v1904
	goto L182
L319:
	;
	goto L144
L320:
	;
	m.G0 = v1975 + int32(_a_F_gistvacuumscan_0)
	return
}
func F_gtsquery_compress(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v33 int64
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int64
	_ = v41
	var v44 int64
	_ = v44
	var v45 int32
	_ = v45
	var v49 int64
	_ = v49
	var v52 int64
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v65 int32
	_ = v65
	var v69 int64
	_ = v69
	var v73 int64
	_ = v73
	var v82 int64
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+18)))
	if v4 != int32(1) {
		return base.I64_extend_i32_u(v3)
	} else {
		v10 = F_palloc(m, int32(24))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int64(0)
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
			v15 = int64(0)
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
			if v19 <= int32(0) {
				v82 = int64(0)
			} else {
				v24 = v14 + int32(8)
				if v19 != int32(1) {
					v32 = v24
					v33 = v15
					v34 = int32(0)
					for {
						v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
						if v37 == int32(1) {
							v41 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v32)+4)))
							v44 = int64(1)<<(uint(v41)%64) | v33
						} else {
							v44 = v33
						}
						v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+12)))
						if v45 == int32(1) {
							v49 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v32)+16)))
							v52 = int64(1)<<(uint(v49)%64) | v44
						} else {
							v52 = v44
						}
						v54 = v32 + int32(24)
						v56 = v34 + int32(2)
						if v56 != v19&int32(2147483646) {
							v32 = v54
							v33 = v52
							v34 = v56
							continue
						} else {
							break
						}
						break
					}
					if v19&int32(1) == int32(0) {
						v73 = v52
					} else {
						v60 = v54
						v61 = v52
						v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
						if v65 != int32(1) {
							v73 = v61
						} else {
							v69 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v60)+4)))
							v73 = int64(1)<<(uint(v69)%64) | v61
						}
					}
				} else {
					v60 = v24
					v61 = v15
					v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
					if v65 != int32(1) {
						v73 = v61
					} else {
						v69 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v60)+4)))
						v73 = int64(1)<<(uint(v69)%64) | v61
					}
				}
				v82 = v73
			}
			*(*int64)(unsafe.Add(mBase, uint32(v10))) = v82
			v84 = *(*int32)(unsafe.Add(mBase, uint32(v3)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v84
			v86 = *(*int32)(unsafe.Add(mBase, uint32(v3)+12))
			*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v86
			v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3)+16)))
			v89 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v10)+18)) = uint8(v89)
			*(*uint16)(unsafe.Add(mBase, uint32(v10)+16)) = uint16(v88)
			return base.I64_extend_i32_u(v10)
		}
	}
}
func F_gtsquery_consistent(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int64
	_ = v38
	var v41 int64
	_ = v41
	var v42 int32
	_ = v42
	var v46 int64
	_ = v46
	var v49 int64
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int64
	_ = v58
	var v62 int32
	_ = v62
	var v66 int64
	_ = v66
	var v70 int64
	_ = v70
	var v79 int64
	_ = v79
	var v80 int32
	_ = v80
	var v85 int64
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v99 int64
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	v2 = int64(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
	v9 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+56)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v16 <= int32(0) {
		v79 = int64(0)
	} else {
		v21 = v11 + int32(8)
		if v16 != int32(1) {
			v29 = v21
			v30 = v2
			v31 = int32(0)
			for {
				v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
				if v34 == int32(1) {
					v38 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v29)+4)))
					v41 = int64(1)<<(uint(v38)%64) | v30
				} else {
					v41 = v30
				}
				v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+12)))
				if v42 == int32(1) {
					v46 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v29)+16)))
					v49 = int64(1)<<(uint(v46)%64) | v41
				} else {
					v49 = v41
				}
				v51 = v29 + int32(24)
				v53 = v31 + int32(2)
				if v53 != v16&int32(2147483646) {
					v29 = v51
					v30 = v49
					v31 = v53
					continue
				} else {
					break
				}
				break
			}
			if v16&int32(1) == int32(0) {
				v70 = v49
			} else {
				v57 = v51
				v58 = v49
				v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
				if v62 != int32(1) {
					v70 = v58
				} else {
					v66 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v57)+4)))
					v70 = int64(1)<<(uint(v66)%64) | v58
				}
			}
		} else {
			v57 = v21
			v58 = v2
			v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
			if v62 != int32(1) {
				v70 = v58
			} else {
				v66 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v57)+4)))
				v70 = int64(1)<<(uint(v66)%64) | v58
			}
		}
		v79 = v70
	}
	v80 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v10))) = uint8(v80)
	switch v9 - int32(7) {
	case 0:
		v85 = v8 & v79
		v86 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
		v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v86)+16)))
		v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86+v87)+12)))
		if v89&int32(1) != 0 {
			return base.I64_extend_i32_u(base.B2i32(v85 == v79))
		} else {
			return base.I64_extend_i32_u(base.B2i32(v85 != int64(0)))
		}
	case 1:
		v99 = v8 & v79
		v100 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
		v101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100)+16)))
		v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100+v101)+12)))
		if v103&int32(1) != 0 {
			return base.I64_extend_i32_u(base.B2i32(v8 == v99))
		} else {
			v111 = base.B2i32(v99 != int64(0))
			return base.I64_extend_i32_u(v111)
		}
	default:
		v111 = int32(0)
		return base.I64_extend_i32_u(v111)
	}
}
