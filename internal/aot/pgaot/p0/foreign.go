package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_DropForeignKeyConstraintTriggers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	v7 = m.G0
	v9 = v7 + int32(-64)
	m.G0 = v9
	F_ScanKeyInit(m, v7+int32(-48), int32(11), int32(3), int32(184), l1)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v19 = int32(1)
	v24 = F_systable_beginscan(m, l0, int32(2699), v19, int32(0), v19, v7+int32(-48))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v26 = F_systable_getnext(m, v24)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if v26 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v29 = v26
	goto L8
L6:
	;
	goto L7
L7:
	;
	F_systable_endscan(m, v24)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L23
	}
L8:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+22)))
	v36 = v34 + v35
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+84))
	if v37 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L7
L10:
	;
	v70 = F_systable_getnext(m, v24)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L21
	}
L11:
	;
	v40 = int32(0)
	if base.B2i32(l3 == v40)|base.B2i32(v37 == l3) == v40 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	if l2 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	if v46 != l2 {
		goto L10
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v51 = F_deleteDependencyRecordsFor(m, int32(2620), v49, int32(0))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	goto L15
L17:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = int32(2620)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v58 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v57
	F_performDeletion(m, v7+int32(-60), v58, v58)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	goto L10
L21:
	;
	if v70 != 0 {
		v29 = v70
		goto L8
	} else {
		goto L22
	}
L22:
	;
	goto L9
L23:
	;
	m.G0 = v9 - int32(-64)
	return
}
func F_ExecForeignScan(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+156))
	if v3 != 0 {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+80))
		if v6 != int32(1) {
			v15 = int32(0)
			return v15
		} else {
			v11 = F_ExecScan(m, l0, int32(708), int32(709))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				v15 = v11
				return v15
			}
		}
	} else {
		v11 = F_ExecScan(m, l0, int32(708), int32(709))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v15 = v11
			return v15
		}
	}
}
func F_GetForeignServerExtended(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = F_SearchSysCache1(m, int32(32), l0)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if v12 == int32(0) {
			if l1&int32(1) != 0 {
				v89 = int32(0)
				m.G0 = v9 + int32(16)
				return v89
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
					F_errmsg_internal(m, int32(42034), v9)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(479868), int32(137), int32(445787))
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
		} else {
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
			v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+22)))
			v36 = F_palloc(m, int32(28))
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v36))) = l0
				v39 = v33 + v34
				v42 = F_pstrdup(m, v39+int32(4))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v36)+12)) = v42
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v39)+68))
					*(*int32)(unsafe.Add(mBase, uint32(v36)+8)) = v45
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v39)+72))
					*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = v47
					v53 = F_SysCacheGetAttr(m, int32(32), v12, int32(5), v9+int32(15))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
						if v55 != 0 {
							v59 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = v59
							v65 = F_SysCacheGetAttr(m, int32(32), v12, int32(6), v9+int32(15))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
								if v67 != 0 {
									v71 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v36)+20)) = v71
									v77 = F_SysCacheGetAttr(m, int32(32), v12, int32(8), v9+int32(15))
									mBase = m.M
									v78 = m.ExcPending
									if v78 != 0 {
										return int32(0)
									} else {
										v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
										if v79 != 0 {
											v83 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v36)+24)) = v83
											F_ReleaseCatCache(m, v12)
											mBase = m.M
											v86 = m.ExcPending
											if v86 != 0 {
												return int32(0)
											} else {
												v89 = v36
												m.G0 = v9 + int32(16)
												return v89
											}
										} else {
											v81 = F_untransformRelOptions(m, v77)
											mBase = m.M
											v82 = m.ExcPending
											if v82 != 0 {
												return int32(0)
											} else {
												v83 = v81
												*(*int32)(unsafe.Add(mBase, uint32(v36)+24)) = v83
												F_ReleaseCatCache(m, v12)
												mBase = m.M
												v86 = m.ExcPending
												if v86 != 0 {
													return int32(0)
												} else {
													v89 = v36
													m.G0 = v9 + int32(16)
													return v89
												}
											}
										}
									}
								} else {
									v69 = F_text_to_cstring(m, v65)
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return int32(0)
									} else {
										v71 = v69
										*(*int32)(unsafe.Add(mBase, uint32(v36)+20)) = v71
										v77 = F_SysCacheGetAttr(m, int32(32), v12, int32(8), v9+int32(15))
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return int32(0)
										} else {
											v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
											if v79 != 0 {
												v83 = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(v36)+24)) = v83
												F_ReleaseCatCache(m, v12)
												mBase = m.M
												v86 = m.ExcPending
												if v86 != 0 {
													return int32(0)
												} else {
													v89 = v36
													m.G0 = v9 + int32(16)
													return v89
												}
											} else {
												v81 = F_untransformRelOptions(m, v77)
												mBase = m.M
												v82 = m.ExcPending
												if v82 != 0 {
													return int32(0)
												} else {
													v83 = v81
													*(*int32)(unsafe.Add(mBase, uint32(v36)+24)) = v83
													F_ReleaseCatCache(m, v12)
													mBase = m.M
													v86 = m.ExcPending
													if v86 != 0 {
														return int32(0)
													} else {
														v89 = v36
														m.G0 = v9 + int32(16)
														return v89
													}
												}
											}
										}
									}
								}
							}
						} else {
							v57 = F_text_to_cstring(m, v53)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								v59 = v57
								*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = v59
								v65 = F_SysCacheGetAttr(m, int32(32), v12, int32(6), v9+int32(15))
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int32(0)
								} else {
									v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
									if v67 != 0 {
										v71 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v36)+20)) = v71
										v77 = F_SysCacheGetAttr(m, int32(32), v12, int32(8), v9+int32(15))
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return int32(0)
										} else {
											v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
											if v79 != 0 {
												v83 = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(v36)+24)) = v83
												F_ReleaseCatCache(m, v12)
												mBase = m.M
												v86 = m.ExcPending
												if v86 != 0 {
													return int32(0)
												} else {
													v89 = v36
													m.G0 = v9 + int32(16)
													return v89
												}
											} else {
												v81 = F_untransformRelOptions(m, v77)
												mBase = m.M
												v82 = m.ExcPending
												if v82 != 0 {
													return int32(0)
												} else {
													v83 = v81
													*(*int32)(unsafe.Add(mBase, uint32(v36)+24)) = v83
													F_ReleaseCatCache(m, v12)
													mBase = m.M
													v86 = m.ExcPending
													if v86 != 0 {
														return int32(0)
													} else {
														v89 = v36
														m.G0 = v9 + int32(16)
														return v89
													}
												}
											}
										}
									} else {
										v69 = F_text_to_cstring(m, v65)
										mBase = m.M
										v70 = m.ExcPending
										if v70 != 0 {
											return int32(0)
										} else {
											v71 = v69
											*(*int32)(unsafe.Add(mBase, uint32(v36)+20)) = v71
											v77 = F_SysCacheGetAttr(m, int32(32), v12, int32(8), v9+int32(15))
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return int32(0)
											} else {
												v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
												if v79 != 0 {
													v83 = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(v36)+24)) = v83
													F_ReleaseCatCache(m, v12)
													mBase = m.M
													v86 = m.ExcPending
													if v86 != 0 {
														return int32(0)
													} else {
														v89 = v36
														m.G0 = v9 + int32(16)
														return v89
													}
												} else {
													v81 = F_untransformRelOptions(m, v77)
													mBase = m.M
													v82 = m.ExcPending
													if v82 != 0 {
														return int32(0)
													} else {
														v83 = v81
														*(*int32)(unsafe.Add(mBase, uint32(v36)+24)) = v83
														F_ReleaseCatCache(m, v12)
														mBase = m.M
														v86 = m.ExcPending
														if v86 != 0 {
															return int32(0)
														} else {
															v89 = v36
															m.G0 = v9 + int32(16)
															return v89
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
}
func F_createForeignKeyActionTriggers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	v10 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(32)
	m.G0 = v16
	v19 = F_palloc0(m, int32(52))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return
	} else {
		v21 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = v21
		*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v21
		*(*int64)(unsafe.Add(mBase, uint32(v19)+32)) = int64(0)
		*(*int32)(unsafe.Add(mBase, uint32(v19)+26)) = int32(524288)
		v29 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v19)+24)) = uint8(v29)
		*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v21
		*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v21
		*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = int32(490114)
		v37 = int32(256)
		*(*uint16)(unsafe.Add(mBase, uint32(v19)+4)) = uint16(v37)
		*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(181)
		v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+88)))
		switch v41 - int32(97) {
		case 0:
			v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+12)))
			*(*uint8)(unsafe.Add(mBase, uint32(v19)+44)) = uint8(v70)
			v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+13)))
			v74 = v72
			v75 = int32(296907)
			*(*uint8)(unsafe.Add(mBase, uint32(v19)+45)) = uint8(v74)
			v77 = F_SystemFuncName(m, v75)
			mBase = m.M
			v78 = m.ExcPending
			if v78 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v77
				v80 = int32(0)
				F_CreateTrigger(m, v16+int32(20), v19, v80, l1, l0, l3, l4, l5, int32(1))
				mBase = m.M
				v86 = m.ExcPending
				if v86 != 0 {
					return
				} else {
					v87 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
					*(*int32)(unsafe.Add(mBase, uint32(l7))) = v87
					F_CommandCounterIncrement(m)
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return
					} else {
						v92 = F_palloc0(m, int32(52))
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
							return
						} else {
							v94 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v92)+48)) = v94
							*(*int32)(unsafe.Add(mBase, uint32(v92)+40)) = v94
							*(*int64)(unsafe.Add(mBase, uint32(v92)+32)) = int64(0)
							*(*int32)(unsafe.Add(mBase, uint32(v92)+26)) = int32(1048576)
							v102 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v92)+24)) = uint8(v102)
							*(*int32)(unsafe.Add(mBase, uint32(v92)+20)) = v94
							*(*int32)(unsafe.Add(mBase, uint32(v92)+12)) = v94
							*(*int32)(unsafe.Add(mBase, uint32(v92)+8)) = int32(490114)
							v110 = int32(256)
							*(*uint16)(unsafe.Add(mBase, uint32(v92)+4)) = uint16(v110)
							*(*int32)(unsafe.Add(mBase, uint32(v92))) = int32(181)
							v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+87)))
							switch v114 - int32(97) {
							case 0:
								v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+12)))
								*(*uint8)(unsafe.Add(mBase, uint32(v92)+44)) = uint8(v145)
								v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+13)))
								v149 = v147
								v150 = int32(407117)
								*(*uint8)(unsafe.Add(mBase, uint32(v92)+45)) = uint8(v149)
								v152 = F_SystemFuncName(m, v150)
								mBase = m.M
								v153 = m.ExcPending
								if v153 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v92)+16)) = v152
									F_CreateTrigger(m, v16+int32(20), v92, int32(0), l1, l0, l3, l4, l6, int32(1))
									mBase = m.M
									v160 = m.ExcPending
									if v160 != 0 {
										return
									} else {
										v161 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
										*(*int32)(unsafe.Add(mBase, uint32(l8))) = v161
										m.G0 = v16 + int32(32)
										return
									}
								}
							default:
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v132 = m.ExcPending
								if v132 != 0 {
									return
								} else {
									v133 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2)+87)))
									*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v133
									F_errmsg_internal(m, int32(468696), v16+int32(16))
									mBase = m.M
									v139 = m.ExcPending
									if v139 != 0 {
										return
									} else {
										F_errfinish(m, int32(477727), int32(13972), int32(127860))
										mBase = m.M
										v144 = m.ExcPending
										if v144 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							case 2:
								v120 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v92)+44)) = uint8(v120)
								v149 = v80
								v150 = int32(407176)
								*(*uint8)(unsafe.Add(mBase, uint32(v92)+45)) = uint8(v149)
								v152 = F_SystemFuncName(m, v150)
								mBase = m.M
								v153 = m.ExcPending
								if v153 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v92)+16)) = v152
									F_CreateTrigger(m, v16+int32(20), v92, int32(0), l1, l0, l3, l4, l6, int32(1))
									mBase = m.M
									v160 = m.ExcPending
									if v160 != 0 {
										return
									} else {
										v161 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
										*(*int32)(unsafe.Add(mBase, uint32(l8))) = v161
										m.G0 = v16 + int32(32)
										return
									}
								}
							case 3:
								v126 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v92)+44)) = uint8(v126)
								v149 = v80
								v150 = int32(407073)
								*(*uint8)(unsafe.Add(mBase, uint32(v92)+45)) = uint8(v149)
								v152 = F_SystemFuncName(m, v150)
								mBase = m.M
								v153 = m.ExcPending
								if v153 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v92)+16)) = v152
									F_CreateTrigger(m, v16+int32(20), v92, int32(0), l1, l0, l3, l4, l6, int32(1))
									mBase = m.M
									v160 = m.ExcPending
									if v160 != 0 {
										return
									} else {
										v161 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
										*(*int32)(unsafe.Add(mBase, uint32(l8))) = v161
										m.G0 = v16 + int32(32)
										return
									}
								}
							case 13:
								v123 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v92)+44)) = uint8(v123)
								v149 = v80
								v150 = int32(407138)
								*(*uint8)(unsafe.Add(mBase, uint32(v92)+45)) = uint8(v149)
								v152 = F_SystemFuncName(m, v150)
								mBase = m.M
								v153 = m.ExcPending
								if v153 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v92)+16)) = v152
									F_CreateTrigger(m, v16+int32(20), v92, int32(0), l1, l0, l3, l4, l6, int32(1))
									mBase = m.M
									v160 = m.ExcPending
									if v160 != 0 {
										return
									} else {
										v161 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
										*(*int32)(unsafe.Add(mBase, uint32(l8))) = v161
										m.G0 = v16 + int32(32)
										return
									}
								}
							case 17:
								v117 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v92)+44)) = uint8(v117)
								v149 = v80
								v150 = int32(407096)
								*(*uint8)(unsafe.Add(mBase, uint32(v92)+45)) = uint8(v149)
								v152 = F_SystemFuncName(m, v150)
								mBase = m.M
								v153 = m.ExcPending
								if v153 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v92)+16)) = v152
									F_CreateTrigger(m, v16+int32(20), v92, int32(0), l1, l0, l3, l4, l6, int32(1))
									mBase = m.M
									v160 = m.ExcPending
									if v160 != 0 {
										return
									} else {
										v161 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
										*(*int32)(unsafe.Add(mBase, uint32(l8))) = v161
										m.G0 = v16 + int32(32)
										return
									}
								}
							}
						}
					}
				}
			}
		default:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return
			} else {
				v60 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2)+88)))
				*(*int32)(unsafe.Add(mBase, uint32(v16))) = v60
				F_errmsg_internal(m, int32(468696), v16)
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return
				} else {
					F_errfinish(m, int32(477727), int32(13912), int32(127860))
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		case 2:
			v47 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v19)+44)) = uint8(v47)
			v74 = v10
			v75 = int32(296948)
			*(*uint8)(unsafe.Add(mBase, uint32(v19)+45)) = uint8(v74)
			v77 = F_SystemFuncName(m, v75)
			mBase = m.M
			v78 = m.ExcPending
			if v78 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v77
				v80 = int32(0)
				F_CreateTrigger(m, v16+int32(20), v19, v80, l1, l0, l3, l4, l5, int32(1))
				mBase = m.M
				v86 = m.ExcPending
				if v86 != 0 {
					return
				} else {
					v87 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
					*(*int32)(unsafe.Add(mBase, uint32(l7))) = v87
					F_CommandCounterIncrement(m)
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return
					} else {
						v92 = F_palloc0(m, int32(52))
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
							return
						} else {
							v94 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v92)+48)) = v94
							*(*int32)(unsafe.Add(mBase, uint32(v92)+40)) = v94
							*(*int64)(unsafe.Add(mBase, uint32(v92)+32)) = int64(0)
							*(*int32)(unsafe.Add(mBase, uint32(v92)+26)) = int32(1048576)
							v102 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v92)+24)) = uint8(v102)
							*(*int32)(unsafe.Add(mBase, uint32(v92)+20)) = v94
							*(*int32)(unsafe.Add(mBase, uint32(v92)+12)) = v94
							*(*int32)(unsafe.Add(mBase, uint32(v92)+8)) = int32(490114)
							v110 = int32(256)
							*(*uint16)(unsafe.Add(mBase, uint32(v92)+4)) = uint16(v110)
							*(*int32)(unsafe.Add(mBase, uint32(v92))) = int32(181)
							v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+87)))
							switch v114 - int32(97) {
							case 0:
								v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+12)))
								*(*uint8)(unsafe.Add(mBase, uint32(v92)+44)) = uint8(v145)
								v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+13)))
								v149 = v147
								v150 = int32(407117)
								*(*uint8)(unsafe.Add(mBase, uint32(v92)+45)) = uint8(v149)
								v152 = F_SystemFuncName(m, v150)
								mBase = m.M
								v153 = m.ExcPending
								if v153 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v92)+16)) = v152
									F_CreateTrigger(m, v16+int32(20), v92, int32(0), l1, l0, l3, l4, l6, int32(1))
									mBase = m.M
									v160 = m.ExcPending
									if v160 != 0 {
										return
									} else {
										v161 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
										*(*int32)(unsafe.Add(mBase, uint32(l8))) = v161
										m.G0 = v16 + int32(32)
										return
									}
								}
							default:
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v132 = m.ExcPending
								if v132 != 0 {
									return
								} else {
									v133 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2)+87)))
									*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v133
									F_errmsg_internal(m, int32(468696), v16+int32(16))
									mBase = m.M
									v139 = m.ExcPending
									if v139 != 0 {
										return
									} else {
										F_errfinish(m, int32(477727), int32(13972), int32(127860))
										mBase = m.M
										v144 = m.ExcPending
										if v144 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							case 2:
								v120 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v92)+44)) = uint8(v120)
								v149 = v80
								v150 = int32(407176)
								*(*uint8)(unsafe.Add(mBase, uint32(v92)+45)) = uint8(v149)
								v152 = F_SystemFuncName(m, v150)
								mBase = m.M
								v153 = m.ExcPending
								if v153 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v92)+16)) = v152
									F_CreateTrigger(m, v16+int32(20), v92, int32(0), l1, l0, l3, l4, l6, int32(1))
									mBase = m.M
									v160 = m.ExcPending
									if v160 != 0 {
										return
									} else {
										v161 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
										*(*int32)(unsafe.Add(mBase, uint32(l8))) = v161
										m.G0 = v16 + int32(32)
										return
									}
								}
							case 3:
								v126 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v92)+44)) = uint8(v126)
								v149 = v80
								v150 = int32(407073)
								*(*uint8)(unsafe.Add(mBase, uint32(v92)+45)) = uint8(v149)
								v152 = F_SystemFuncName(m, v150)
								mBase = m.M
								v153 = m.ExcPending
								if v153 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v92)+16)) = v152
									F_CreateTrigger(m, v16+int32(20), v92, int32(0), l1, l0, l3, l4, l6, int32(1))
									mBase = m.M
									v160 = m.ExcPending
									if v160 != 0 {
										return
									} else {
										v161 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
										*(*int32)(unsafe.Add(mBase, uint32(l8))) = v161
										m.G0 = v16 + int32(32)
										return
									}
								}
							case 13:
								v123 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v92)+44)) = uint8(v123)
								v149 = v80
								v150 = int32(407138)
								*(*uint8)(unsafe.Add(mBase, uint32(v92)+45)) = uint8(v149)
								v152 = F_SystemFuncName(m, v150)
								mBase = m.M
								v153 = m.ExcPending
								if v153 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v92)+16)) = v152
									F_CreateTrigger(m, v16+int32(20), v92, int32(0), l1, l0, l3, l4, l6, int32(1))
									mBase = m.M
									v160 = m.ExcPending
									if v160 != 0 {
										return
									} else {
										v161 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
										*(*int32)(unsafe.Add(mBase, uint32(l8))) = v161
										m.G0 = v16 + int32(32)
										return
									}
								}
							case 17:
								v117 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v92)+44)) = uint8(v117)
								v149 = v80
								v150 = int32(407096)
								*(*uint8)(unsafe.Add(mBase, uint32(v92)+45)) = uint8(v149)
								v152 = F_SystemFuncName(m, v150)
								mBase = m.M
								v153 = m.ExcPending
								if v153 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v92)+16)) = v152
									F_CreateTrigger(m, v16+int32(20), v92, int32(0), l1, l0, l3, l4, l6, int32(1))
									mBase = m.M
									v160 = m.ExcPending
									if v160 != 0 {
										return
									} else {
										v161 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
										*(*int32)(unsafe.Add(mBase, uint32(l8))) = v161
										m.G0 = v16 + int32(32)
										return
									}
								}
							}
						}
					}
				}
			}
		case 3:
			v53 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v19)+44)) = uint8(v53)
			v74 = v10
			v75 = int32(296863)
			*(*uint8)(unsafe.Add(mBase, uint32(v19)+45)) = uint8(v74)
			v77 = F_SystemFuncName(m, v75)
			mBase = m.M
			v78 = m.ExcPending
			if v78 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v77
				v80 = int32(0)
				F_CreateTrigger(m, v16+int32(20), v19, v80, l1, l0, l3, l4, l5, int32(1))
				mBase = m.M
				v86 = m.ExcPending
				if v86 != 0 {
					return
				} else {
					v87 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
					*(*int32)(unsafe.Add(mBase, uint32(l7))) = v87
					F_CommandCounterIncrement(m)
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return
					} else {
						v92 = F_palloc0(m, int32(52))
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
							return
						} else {
							v94 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v92)+48)) = v94
							*(*int32)(unsafe.Add(mBase, uint32(v92)+40)) = v94
							*(*int64)(unsafe.Add(mBase, uint32(v92)+32)) = int64(0)
							*(*int32)(unsafe.Add(mBase, uint32(v92)+26)) = int32(1048576)
							v102 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v92)+24)) = uint8(v102)
							*(*int32)(unsafe.Add(mBase, uint32(v92)+20)) = v94
							*(*int32)(unsafe.Add(mBase, uint32(v92)+12)) = v94
							*(*int32)(unsafe.Add(mBase, uint32(v92)+8)) = int32(490114)
							v110 = int32(256)
							*(*uint16)(unsafe.Add(mBase, uint32(v92)+4)) = uint16(v110)
							*(*int32)(unsafe.Add(mBase, uint32(v92))) = int32(181)
							v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+87)))
							switch v114 - int32(97) {
							case 0:
								v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+12)))
								*(*uint8)(unsafe.Add(mBase, uint32(v92)+44)) = uint8(v145)
								v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+13)))
								v149 = v147
								v150 = int32(407117)
								*(*uint8)(unsafe.Add(mBase, uint32(v92)+45)) = uint8(v149)
								v152 = F_SystemFuncName(m, v150)
								mBase = m.M
								v153 = m.ExcPending
								if v153 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v92)+16)) = v152
									F_CreateTrigger(m, v16+int32(20), v92, int32(0), l1, l0, l3, l4, l6, int32(1))
									mBase = m.M
									v160 = m.ExcPending
									if v160 != 0 {
										return
									} else {
										v161 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
										*(*int32)(unsafe.Add(mBase, uint32(l8))) = v161
										m.G0 = v16 + int32(32)
										return
									}
								}
							default:
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v132 = m.ExcPending
								if v132 != 0 {
									return
								} else {
									v133 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2)+87)))
									*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v133
									F_errmsg_internal(m, int32(468696), v16+int32(16))
									mBase = m.M
									v139 = m.ExcPending
									if v139 != 0 {
										return
									} else {
										F_errfinish(m, int32(477727), int32(13972), int32(127860))
										mBase = m.M
										v144 = m.ExcPending
										if v144 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							case 2:
								v120 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v92)+44)) = uint8(v120)
								v149 = v80
								v150 = int32(407176)
								*(*uint8)(unsafe.Add(mBase, uint32(v92)+45)) = uint8(v149)
								v152 = F_SystemFuncName(m, v150)
								mBase = m.M
								v153 = m.ExcPending
								if v153 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v92)+16)) = v152
									F_CreateTrigger(m, v16+int32(20), v92, int32(0), l1, l0, l3, l4, l6, int32(1))
									mBase = m.M
									v160 = m.ExcPending
									if v160 != 0 {
										return
									} else {
										v161 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
										*(*int32)(unsafe.Add(mBase, uint32(l8))) = v161
										m.G0 = v16 + int32(32)
										return
									}
								}
							case 3:
								v126 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v92)+44)) = uint8(v126)
								v149 = v80
								v150 = int32(407073)
								*(*uint8)(unsafe.Add(mBase, uint32(v92)+45)) = uint8(v149)
								v152 = F_SystemFuncName(m, v150)
								mBase = m.M
								v153 = m.ExcPending
								if v153 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v92)+16)) = v152
									F_CreateTrigger(m, v16+int32(20), v92, int32(0), l1, l0, l3, l4, l6, int32(1))
									mBase = m.M
									v160 = m.ExcPending
									if v160 != 0 {
										return
									} else {
										v161 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
										*(*int32)(unsafe.Add(mBase, uint32(l8))) = v161
										m.G0 = v16 + int32(32)
										return
									}
								}
							case 13:
								v123 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v92)+44)) = uint8(v123)
								v149 = v80
								v150 = int32(407138)
								*(*uint8)(unsafe.Add(mBase, uint32(v92)+45)) = uint8(v149)
								v152 = F_SystemFuncName(m, v150)
								mBase = m.M
								v153 = m.ExcPending
								if v153 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v92)+16)) = v152
									F_CreateTrigger(m, v16+int32(20), v92, int32(0), l1, l0, l3, l4, l6, int32(1))
									mBase = m.M
									v160 = m.ExcPending
									if v160 != 0 {
										return
									} else {
										v161 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
										*(*int32)(unsafe.Add(mBase, uint32(l8))) = v161
										m.G0 = v16 + int32(32)
										return
									}
								}
							case 17:
								v117 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v92)+44)) = uint8(v117)
								v149 = v80
								v150 = int32(407096)
								*(*uint8)(unsafe.Add(mBase, uint32(v92)+45)) = uint8(v149)
								v152 = F_SystemFuncName(m, v150)
								mBase = m.M
								v153 = m.ExcPending
								if v153 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v92)+16)) = v152
									F_CreateTrigger(m, v16+int32(20), v92, int32(0), l1, l0, l3, l4, l6, int32(1))
									mBase = m.M
									v160 = m.ExcPending
									if v160 != 0 {
										return
									} else {
										v161 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
										*(*int32)(unsafe.Add(mBase, uint32(l8))) = v161
										m.G0 = v16 + int32(32)
										return
									}
								}
							}
						}
					}
				}
			}
		case 13:
			v50 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v19)+44)) = uint8(v50)
			v74 = v10
			v75 = int32(296928)
			*(*uint8)(unsafe.Add(mBase, uint32(v19)+45)) = uint8(v74)
			v77 = F_SystemFuncName(m, v75)
			mBase = m.M
			v78 = m.ExcPending
			if v78 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v77
				v80 = int32(0)
				F_CreateTrigger(m, v16+int32(20), v19, v80, l1, l0, l3, l4, l5, int32(1))
				mBase = m.M
				v86 = m.ExcPending
				if v86 != 0 {
					return
				} else {
					v87 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
					*(*int32)(unsafe.Add(mBase, uint32(l7))) = v87
					F_CommandCounterIncrement(m)
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return
					} else {
						v92 = F_palloc0(m, int32(52))
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
							return
						} else {
							v94 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v92)+48)) = v94
							*(*int32)(unsafe.Add(mBase, uint32(v92)+40)) = v94
							*(*int64)(unsafe.Add(mBase, uint32(v92)+32)) = int64(0)
							*(*int32)(unsafe.Add(mBase, uint32(v92)+26)) = int32(1048576)
							v102 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v92)+24)) = uint8(v102)
							*(*int32)(unsafe.Add(mBase, uint32(v92)+20)) = v94
							*(*int32)(unsafe.Add(mBase, uint32(v92)+12)) = v94
							*(*int32)(unsafe.Add(mBase, uint32(v92)+8)) = int32(490114)
							v110 = int32(256)
							*(*uint16)(unsafe.Add(mBase, uint32(v92)+4)) = uint16(v110)
							*(*int32)(unsafe.Add(mBase, uint32(v92))) = int32(181)
							v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+87)))
							switch v114 - int32(97) {
							case 0:
								v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+12)))
								*(*uint8)(unsafe.Add(mBase, uint32(v92)+44)) = uint8(v145)
								v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+13)))
								v149 = v147
								v150 = int32(407117)
								*(*uint8)(unsafe.Add(mBase, uint32(v92)+45)) = uint8(v149)
								v152 = F_SystemFuncName(m, v150)
								mBase = m.M
								v153 = m.ExcPending
								if v153 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v92)+16)) = v152
									F_CreateTrigger(m, v16+int32(20), v92, int32(0), l1, l0, l3, l4, l6, int32(1))
									mBase = m.M
									v160 = m.ExcPending
									if v160 != 0 {
										return
									} else {
										v161 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
										*(*int32)(unsafe.Add(mBase, uint32(l8))) = v161
										m.G0 = v16 + int32(32)
										return
									}
								}
							default:
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v132 = m.ExcPending
								if v132 != 0 {
									return
								} else {
									v133 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2)+87)))
									*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v133
									F_errmsg_internal(m, int32(468696), v16+int32(16))
									mBase = m.M
									v139 = m.ExcPending
									if v139 != 0 {
										return
									} else {
										F_errfinish(m, int32(477727), int32(13972), int32(127860))
										mBase = m.M
										v144 = m.ExcPending
										if v144 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							case 2:
								v120 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v92)+44)) = uint8(v120)
								v149 = v80
								v150 = int32(407176)
								*(*uint8)(unsafe.Add(mBase, uint32(v92)+45)) = uint8(v149)
								v152 = F_SystemFuncName(m, v150)
								mBase = m.M
								v153 = m.ExcPending
								if v153 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v92)+16)) = v152
									F_CreateTrigger(m, v16+int32(20), v92, int32(0), l1, l0, l3, l4, l6, int32(1))
									mBase = m.M
									v160 = m.ExcPending
									if v160 != 0 {
										return
									} else {
										v161 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
										*(*int32)(unsafe.Add(mBase, uint32(l8))) = v161
										m.G0 = v16 + int32(32)
										return
									}
								}
							case 3:
								v126 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v92)+44)) = uint8(v126)
								v149 = v80
								v150 = int32(407073)
								*(*uint8)(unsafe.Add(mBase, uint32(v92)+45)) = uint8(v149)
								v152 = F_SystemFuncName(m, v150)
								mBase = m.M
								v153 = m.ExcPending
								if v153 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v92)+16)) = v152
									F_CreateTrigger(m, v16+int32(20), v92, int32(0), l1, l0, l3, l4, l6, int32(1))
									mBase = m.M
									v160 = m.ExcPending
									if v160 != 0 {
										return
									} else {
										v161 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
										*(*int32)(unsafe.Add(mBase, uint32(l8))) = v161
										m.G0 = v16 + int32(32)
										return
									}
								}
							case 13:
								v123 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v92)+44)) = uint8(v123)
								v149 = v80
								v150 = int32(407138)
								*(*uint8)(unsafe.Add(mBase, uint32(v92)+45)) = uint8(v149)
								v152 = F_SystemFuncName(m, v150)
								mBase = m.M
								v153 = m.ExcPending
								if v153 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v92)+16)) = v152
									F_CreateTrigger(m, v16+int32(20), v92, int32(0), l1, l0, l3, l4, l6, int32(1))
									mBase = m.M
									v160 = m.ExcPending
									if v160 != 0 {
										return
									} else {
										v161 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
										*(*int32)(unsafe.Add(mBase, uint32(l8))) = v161
										m.G0 = v16 + int32(32)
										return
									}
								}
							case 17:
								v117 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v92)+44)) = uint8(v117)
								v149 = v80
								v150 = int32(407096)
								*(*uint8)(unsafe.Add(mBase, uint32(v92)+45)) = uint8(v149)
								v152 = F_SystemFuncName(m, v150)
								mBase = m.M
								v153 = m.ExcPending
								if v153 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v92)+16)) = v152
									F_CreateTrigger(m, v16+int32(20), v92, int32(0), l1, l0, l3, l4, l6, int32(1))
									mBase = m.M
									v160 = m.ExcPending
									if v160 != 0 {
										return
									} else {
										v161 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
										*(*int32)(unsafe.Add(mBase, uint32(l8))) = v161
										m.G0 = v16 + int32(32)
										return
									}
								}
							}
						}
					}
				}
			}
		case 17:
			v44 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v19)+44)) = uint8(v44)
			v74 = v10
			v75 = int32(296886)
			*(*uint8)(unsafe.Add(mBase, uint32(v19)+45)) = uint8(v74)
			v77 = F_SystemFuncName(m, v75)
			mBase = m.M
			v78 = m.ExcPending
			if v78 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v77
				v80 = int32(0)
				F_CreateTrigger(m, v16+int32(20), v19, v80, l1, l0, l3, l4, l5, int32(1))
				mBase = m.M
				v86 = m.ExcPending
				if v86 != 0 {
					return
				} else {
					v87 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
					*(*int32)(unsafe.Add(mBase, uint32(l7))) = v87
					F_CommandCounterIncrement(m)
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return
					} else {
						v92 = F_palloc0(m, int32(52))
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
							return
						} else {
							v94 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v92)+48)) = v94
							*(*int32)(unsafe.Add(mBase, uint32(v92)+40)) = v94
							*(*int64)(unsafe.Add(mBase, uint32(v92)+32)) = int64(0)
							*(*int32)(unsafe.Add(mBase, uint32(v92)+26)) = int32(1048576)
							v102 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v92)+24)) = uint8(v102)
							*(*int32)(unsafe.Add(mBase, uint32(v92)+20)) = v94
							*(*int32)(unsafe.Add(mBase, uint32(v92)+12)) = v94
							*(*int32)(unsafe.Add(mBase, uint32(v92)+8)) = int32(490114)
							v110 = int32(256)
							*(*uint16)(unsafe.Add(mBase, uint32(v92)+4)) = uint16(v110)
							*(*int32)(unsafe.Add(mBase, uint32(v92))) = int32(181)
							v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+87)))
							switch v114 - int32(97) {
							case 0:
								v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+12)))
								*(*uint8)(unsafe.Add(mBase, uint32(v92)+44)) = uint8(v145)
								v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+13)))
								v149 = v147
								v150 = int32(407117)
								*(*uint8)(unsafe.Add(mBase, uint32(v92)+45)) = uint8(v149)
								v152 = F_SystemFuncName(m, v150)
								mBase = m.M
								v153 = m.ExcPending
								if v153 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v92)+16)) = v152
									F_CreateTrigger(m, v16+int32(20), v92, int32(0), l1, l0, l3, l4, l6, int32(1))
									mBase = m.M
									v160 = m.ExcPending
									if v160 != 0 {
										return
									} else {
										v161 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
										*(*int32)(unsafe.Add(mBase, uint32(l8))) = v161
										m.G0 = v16 + int32(32)
										return
									}
								}
							default:
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v132 = m.ExcPending
								if v132 != 0 {
									return
								} else {
									v133 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2)+87)))
									*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v133
									F_errmsg_internal(m, int32(468696), v16+int32(16))
									mBase = m.M
									v139 = m.ExcPending
									if v139 != 0 {
										return
									} else {
										F_errfinish(m, int32(477727), int32(13972), int32(127860))
										mBase = m.M
										v144 = m.ExcPending
										if v144 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							case 2:
								v120 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v92)+44)) = uint8(v120)
								v149 = v80
								v150 = int32(407176)
								*(*uint8)(unsafe.Add(mBase, uint32(v92)+45)) = uint8(v149)
								v152 = F_SystemFuncName(m, v150)
								mBase = m.M
								v153 = m.ExcPending
								if v153 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v92)+16)) = v152
									F_CreateTrigger(m, v16+int32(20), v92, int32(0), l1, l0, l3, l4, l6, int32(1))
									mBase = m.M
									v160 = m.ExcPending
									if v160 != 0 {
										return
									} else {
										v161 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
										*(*int32)(unsafe.Add(mBase, uint32(l8))) = v161
										m.G0 = v16 + int32(32)
										return
									}
								}
							case 3:
								v126 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v92)+44)) = uint8(v126)
								v149 = v80
								v150 = int32(407073)
								*(*uint8)(unsafe.Add(mBase, uint32(v92)+45)) = uint8(v149)
								v152 = F_SystemFuncName(m, v150)
								mBase = m.M
								v153 = m.ExcPending
								if v153 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v92)+16)) = v152
									F_CreateTrigger(m, v16+int32(20), v92, int32(0), l1, l0, l3, l4, l6, int32(1))
									mBase = m.M
									v160 = m.ExcPending
									if v160 != 0 {
										return
									} else {
										v161 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
										*(*int32)(unsafe.Add(mBase, uint32(l8))) = v161
										m.G0 = v16 + int32(32)
										return
									}
								}
							case 13:
								v123 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v92)+44)) = uint8(v123)
								v149 = v80
								v150 = int32(407138)
								*(*uint8)(unsafe.Add(mBase, uint32(v92)+45)) = uint8(v149)
								v152 = F_SystemFuncName(m, v150)
								mBase = m.M
								v153 = m.ExcPending
								if v153 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v92)+16)) = v152
									F_CreateTrigger(m, v16+int32(20), v92, int32(0), l1, l0, l3, l4, l6, int32(1))
									mBase = m.M
									v160 = m.ExcPending
									if v160 != 0 {
										return
									} else {
										v161 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
										*(*int32)(unsafe.Add(mBase, uint32(l8))) = v161
										m.G0 = v16 + int32(32)
										return
									}
								}
							case 17:
								v117 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v92)+44)) = uint8(v117)
								v149 = v80
								v150 = int32(407096)
								*(*uint8)(unsafe.Add(mBase, uint32(v92)+45)) = uint8(v149)
								v152 = F_SystemFuncName(m, v150)
								mBase = m.M
								v153 = m.ExcPending
								if v153 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v92)+16)) = v152
									F_CreateTrigger(m, v16+int32(20), v92, int32(0), l1, l0, l3, l4, l6, int32(1))
									mBase = m.M
									v160 = m.ExcPending
									if v160 != 0 {
										return
									} else {
										v161 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
										*(*int32)(unsafe.Add(mBase, uint32(l8))) = v161
										m.G0 = v16 + int32(32)
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
}
func F_get_foreign_server_oid(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	v3 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v13 = F_GetSysCacheOid(m, int32(31), l0, v3, v3, v3)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		if l1 != 0 {
			m.G0 = v7 + int32(16)
			return v13
		} else {
			if v13 != 0 {
				m.G0 = v7 + int32(16)
				return v13
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
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
						F_errmsg(m, int32(68756), v7)
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(479868), int32(714), int32(418803))
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
	}
}
func F_has_foreign_data_wrapper_privilege_id_id(m *base.Module, l0 int32) int32 {
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
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)) = uint8(v18)
		v22 = F_convert_any_priv_string(m, v14, int32(1624288))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v26 = F_object_aclcheck_ext(m, int32(2328), v11, v12, v22, v9+int32(15))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
				if v28 == int32(1) {
					v31 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v31)
					v35 = int32(0)
				} else {
					v35 = base.B2i32(v26 == int32(0))
				}
				m.G0 = v9 + int32(16)
				return v35
			}
		}
	}
}
func F_has_foreign_data_wrapper_privilege_id_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
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
	var v18 int32
	_ = v18
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v11 = F_pg_detoast_datum_packed(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v14 = F_text_to_cstring(m, v6)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				v17 = F_get_foreign_data_wrapper_oid(m, v14, int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					v20 = F_convert_any_priv_string(m, v11, int32(1624288))
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						v22 = F_object_aclcheck(m, int32(2328), v17, v4, v20)
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return int32(0)
						} else {
							return base.B2i32(v22 == int32(0))
						}
					}
				}
			}
		}
	}
}
func F_has_foreign_data_wrapper_privilege_name_id(m *base.Module, l0 int32) int32 {
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
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int64
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)) = uint8(v18)
		v21 = F_get_role_oid_or_public(m, v12)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v24 = F_convert_any_priv_string(m, v14, int32(1624288))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v28 = F_object_aclcheck_ext(m, int32(2328), v11, v21, v24, v9+int32(15))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
					if v30 == int32(1) {
						v33 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v33)
						v37 = int32(0)
					} else {
						v37 = base.B2i32(v28 == int32(0))
					}
					m.G0 = v9 + int32(16)
					return v37
				}
			}
		}
	}
}
