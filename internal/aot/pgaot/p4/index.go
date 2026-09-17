package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_ExecIndexOnlyScan(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v2 == int32(0) {
		v12 = F_ExecScan(m, l0, int32(723), int32(724))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			return v12
		}
	} else {
		v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+144)))
		if v5 != 0 {
			v12 = F_ExecScan(m, l0, int32(723), int32(724))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				return v12
			}
		} else {
			F_ExecReScan(m, l0)
			mBase = m.M
			v9 = m.ExcPending
			if v9 != 0 {
				return int32(0)
			} else {
				v12 = F_ExecScan(m, l0, int32(723), int32(724))
				mBase = m.M
				v13 = m.ExcPending
				if v13 != 0 {
					return int32(0)
				} else {
					return v12
				}
			}
		}
	}
}
func F_ExecIndexScan(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v2 == int32(0) {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
		if int32(0) < v12 {
			v15 = int32(727)
		} else {
			v15 = int32(728)
		}
		v17 = F_ExecScan(m, l0, v15, int32(729))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			return v17
		}
	} else {
		v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+148)))
		if v5 != 0 {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
			if int32(0) < v12 {
				v15 = int32(727)
			} else {
				v15 = int32(728)
			}
			v17 = F_ExecScan(m, l0, v15, int32(729))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				return v17
			}
		} else {
			F_ExecReScan(m, l0)
			mBase = m.M
			v9 = m.ExcPending
			if v9 != 0 {
				return int32(0)
			} else {
				v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
				if int32(0) < v12 {
					v15 = int32(727)
				} else {
					v15 = int32(728)
				}
				v17 = F_ExecScan(m, l0, v15, int32(729))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					return v17
				}
			}
		}
	}
}
func F_GetIndexAmRoutine(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = F_OidFunctionCall0Coll(m, l0)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		if v8 != 0 {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
			if v12 == int32(438) {
				m.G0 = v6 + int32(16)
				return v8
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
					F_errmsg_internal(m, int32(_a_F_GetIndexAmRoutine_0), v6)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_GetIndexAmRoutine_1), int32(43), int32(_a_F_GetIndexAmRoutine_2))
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
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
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
				F_errmsg_internal(m, int32(_a_F_GetIndexAmRoutine_0), v6)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_GetIndexAmRoutine_1), int32(43), int32(_a_F_GetIndexAmRoutine_2))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
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
func F_IndexGetRelation(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v37 int32
	_ = v37
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = F_SearchSysCache1(m, int32(34), l0)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if v11 == int32(0) {
			if l1 != 0 {
				v37 = int32(0)
				m.G0 = v8 + int32(16)
				return v37
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
					F_errmsg_internal(m, int32(_a_F_IndexGetRelation_0), v8)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_IndexGetRelation_1), int32(3594), int32(_a_F_IndexGetRelation_2))
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
			}
		} else {
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+22)))
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v30+v31)+4))
			F_ReleaseCatCache(m, v11)
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				v37 = v33
				m.G0 = v8 + int32(16)
				return v37
			}
		}
	}
}
func F_IndexSupportsBackwardScan(m *base.Module, l0 int32) int32 {
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
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = F_SearchSysCache1(m, int32(57), l0)
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
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
				F_errmsg_internal(m, int32(_a_F_IndexSupportsBackwardScan_0), v7)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_IndexSupportsBackwardScan_1), int32(613), int32(_a_F_IndexSupportsBackwardScan_2))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+22)))
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v29+v30)+84))
			v34 = F_GetIndexAmRoutineByAmId(m, v32, int32(0))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+15)))
				F_pfree(m, v34)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					F_ReleaseCatCache(m, v10)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						m.G0 = v7 + int32(16)
						return v36
					}
				}
			}
		}
	}
}
func F_RecordFreeIndexPage(m *base.Module, l0 int32, l1 int32) {
	var v5 int32
	_ = v5
	F_RecordPageWithFreeSpace(m, l0, l1, int32(_a_F_RecordFreeIndexPage_0))
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
func F_SetIndexStorageProperties(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
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
	var v54 int32
	_ = v54
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
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
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	v5 = l4
	v7 = l6
	v15 = F_RelationGetIndexList(m, l0)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	return
L3:
	;
	if v15 == int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v19 <= int32(0) {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v34 = int32(0)
	goto L6
L6:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38+v34<<(uint(int32(2))%32))))
	v43 = F_index_open(m, v42, l7)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L2
	} else {
		goto L9
	}
L7:
	;
	goto L1
L8:
	;
	F_relation_close(m, v43, l7)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L2
	} else {
		goto L32
	}
L9:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v43)+192))
	v46 = int32(*(*int16)(unsafe.Add(mBase, uint32(v45)+8)))
	if v46 <= int32(0) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v54 = int32(0)
	goto L11
L11:
	;
	v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45+int32(48)+v54<<(uint(int32(1))%32)))))
	if l2&int32(_a_F_SetIndexStorageProperties_0) != v69 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v74 = int32(_a_F_SetIndexStorageProperties_0)
	v77 = v54&v74 + int32(1)
	if v77&v74 != v77 {
		goto L8
	} else {
		goto L17
	}
L13:
	;
	v72 = v54 + int32(1)
	if v72 != v46 {
		v54 = v72
		goto L11
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	goto L12
L16:
	;
	goto L8
L17:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v43)+56))
	v83 = F_SearchSysCacheCopyAttNum(m, v81, base.I32_extend16_s(v77))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L2
	} else {
		goto L18
	}
L18:
	;
	if v83 == int32(0) {
		goto L8
	} else {
		goto L19
	}
L19:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v83)+16))
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+22)))
	v89 = v87 + v88
	if l3 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v89)+84)) = uint8(v5)
	goto L22
L21:
	;
	goto L22
L22:
	;
	if l5 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v89)+85)) = uint8(v7)
	goto L25
L24:
	;
	goto L25
L25:
	;
	F_CatalogTupleUpdate(m, l1, v83+int32(4), v83)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L2
	} else {
		goto L26
	}
L26:
	;
	v97 = *(*int32)(unsafe.Add(mBase, _c_F_SetIndexStorageProperties[0]))
	if v97 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v100 = int32(*(*int16)(unsafe.Add(mBase, uint32(v89)+74)))
	v101 = int32(0)
	F_RunObjectPostAlterHook(m, int32(1259), v99, v100, v101, v101)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L2
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	F_pfree(m, v83)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L2
	} else {
		goto L31
	}
L30:
	;
	goto L29
L31:
	;
	goto L8
L32:
	;
	v124 = v34 + int32(1)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v124 < v125 {
		v34 = v124
		goto L6
	} else {
		goto L33
	}
L33:
	;
	goto L7
}
func F_check_index_predicates(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
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
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	v3 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+108))
	if v16 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v14 + int32(16)
	return
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v19 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v24 = v3
	v26 = v3
	goto L4
L4:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33+v24<<(uint(int32(2))%32))))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+184))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+96)) = v38
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v37)+88))
	v43 = base.B2i32(v40 != int32(0)) | v26
	v45 = v24 + int32(1)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v45 < v46 {
		v24 = v45
		v26 = v43
		goto L4
	} else {
		goto L6
	}
L5:
	;
	if v43&int32(1) == int32(0) {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	goto L5
L7:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)+184))
	v53 = F_list_copy(m, v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return
L9:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+212))
	if v55 == int32(0) {
		v90 = v53
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v99 == int32(2) {
		goto L22
	} else {
		goto L23
	}
L11:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	if v58 <= int32(0) {
		v90 = v53
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v64 = int32(0)
	v65 = v53
	goto L13
L13:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v55)+12))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v73+v64<<(uint(int32(2))%32))))
	v78 = F_join_clause_is_movable_to(m, v77, l1)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L8
	} else {
		goto L15
	}
L14:
	;
	v90 = v82
	goto L10
L15:
	;
	if v78 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v80 = F_lappend(m, v65, v77)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L8
	} else {
		goto L19
	}
L17:
	;
	v82 = v65
	goto L18
L18:
	;
	v84 = v64 + int32(1)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	if v84 < v85 {
		v64 = v84
		v65 = v82
		goto L13
	} else {
		goto L20
	}
L19:
	;
	v82 = v80
	goto L18
L20:
	;
	goto L14
L21:
	;
	v106 = F_bms_difference(m, v98, v105)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L8
	} else {
		goto L26
	}
L22:
	;
	v102 = F_find_childrel_parents(m, l0, l1)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L8
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v105 = v104
	goto L21
L25:
	;
	v105 = v102
	goto L21
L26:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v109 = F_bms_del_members(m, v106, v108)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L8
	} else {
		goto L27
	}
L27:
	;
	if v109 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v112 = F_bms_union(m, v111, v109)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L8
	} else {
		goto L31
	}
L29:
	;
	v119 = v90
	goto L30
L30:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v123 = F_bms_is_member(m, v121, v122)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L8
	} else {
		goto L34
	}
L31:
	;
	v115 = F_generate_join_implied_equalities(m, l0, v112, v109, l1, int32(0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L8
	} else {
		goto L32
	}
L32:
	;
	v117 = F_list_concat(m, v90, v115)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L8
	} else {
		goto L33
	}
L33:
	;
	v119 = v117
	goto L30
L34:
	;
	if v123 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	if v127 != 0 {
		goto L40
	} else {
		goto L41
	}
L36:
	;
	v168 = int32(1)
	goto L37
L37:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l1)+108))
	if v169 == int32(0) {
		goto L1
	} else {
		goto L51
	}
L38:
	;
	v168 = base.B2i32(v164 != int32(0))
	goto L37
L39:
	;
	goto L38
L40:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
	if v132 <= int32(0) {
		v164 = int32(0)
		goto L39
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v164 = int32(0)
	goto L39
L43:
	;
	v135 = int32(0)
	if v135 < v132 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v138 = v132
	goto L46
L45:
	;
	v138 = v135
	goto L46
L46:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v127)+12))
	v141 = int32(0)
	goto L47
L47:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v139+v141<<(uint(int32(2))%32))))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v149)+4))
	if v150 == v128 {
		v164 = v149
		goto L39
	} else {
		goto L49
	}
L48:
	;
	goto L42
L49:
	;
	v153 = v141 + int32(1)
	if v153 != v138 {
		v141 = v153
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v169)+4))
	if v172 <= int32(0) {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v176 = int32(0)
	goto L53
L53:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v169)+12))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v187+v176<<(uint(int32(2))%32))))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v191)+88))
	if v192 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	goto L1
L55:
	;
	v269 = v176 + int32(1)
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v169)+4))
	if v269 < v270 {
		v176 = v269
		goto L53
	} else {
		goto L77
	}
L56:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191)+100)))
	if v195 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v199 = F_predicate_implied_by(m, v192, v119, int32(0))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L8
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	if v168 != 0 {
		goto L55
	} else {
		goto L61
	}
L60:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v191)+100)) = uint8(v199)
	goto L59
L61:
	;
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191)+106)))
	if v202 != int32(1) {
		goto L55
	} else {
		goto L62
	}
L62:
	;
	v205 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v191)+96)) = v205
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l1)+184))
	if v207 == v205 {
		goto L55
	} else {
		goto L63
	}
L63:
	;
	v210 = int32(0)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v207)+4))
	if v211 <= v210 {
		goto L55
	} else {
		goto L64
	}
L64:
	;
	v222 = v210
	goto L65
L65:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v207)+12))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v225+v222<<(uint(int32(2))%32))))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v229)+4))
	v231 = F_contain_mutable_functions(m, v230)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L8
	} else {
		goto L68
	}
L66:
	;
	goto L55
L67:
	;
	v254 = v222 + int32(1)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v207)+4))
	if v254 < v255 {
		v222 = v254
		goto L65
	} else {
		goto L76
	}
L68:
	;
	if v231 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v229)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v235
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v235
	v241 = F_list_make1_impl(m, int32(1), v14+int32(8))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L8
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v191)+96))
	v249 = F_lappend(m, v248, v229)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L8
	} else {
		goto L75
	}
L72:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v191)+88))
	v245 = F_predicate_implied_by(m, v241, v243, int32(0))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L8
	} else {
		goto L73
	}
L73:
	;
	if v245 != 0 {
		goto L67
	} else {
		goto L74
	}
L74:
	;
	goto L71
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v191)+96)) = v249
	goto L67
L76:
	;
	goto L66
L77:
	;
	goto L54
}
func F_cost_index(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 float64
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v181 int32
	_ = v181
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 float64
	_ = v284
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v403 int32
	_ = v403
	var v410 int32
	_ = v410
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v425 float64
	_ = v425
	var v427 float64
	_ = v427
	var v429 float64
	_ = v429
	var v430 float64
	_ = v430
	var v431 float64
	_ = v431
	var v440 float64
	_ = v440
	var v444 float64
	_ = v444
	var v445 float64
	_ = v445
	var v446 int32
	_ = v446
	var v452 int32
	_ = v452
	var v455 float64
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v461 float64
	_ = v461
	var v462 float64
	_ = v462
	var v463 float64
	_ = v463
	var v465 int32
	_ = v465
	var v468 float64
	_ = v468
	var v469 int32
	_ = v469
	var v471 float64
	_ = v471
	var v475 float64
	_ = v475
	var v476 float64
	_ = v476
	var v480 float64
	_ = v480
	var v481 int32
	_ = v481
	var v484 float64
	_ = v484
	var v489 float64
	_ = v489
	var v499 float64
	_ = v499
	var v503 float64
	_ = v503
	var v507 float64
	_ = v507
	var v511 float64
	_ = v511
	var v512 float64
	_ = v512
	var v513 float64
	_ = v513
	var v517 float64
	_ = v517
	var v520 float64
	_ = v520
	var v525 float64
	_ = v525
	var v535 float64
	_ = v535
	var v537 float64
	_ = v537
	var v545 float64
	_ = v545
	var v549 float64
	_ = v549
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v557 float64
	_ = v557
	var v558 float64
	_ = v558
	var v559 float64
	_ = v559
	var v561 int32
	_ = v561
	var v564 float64
	_ = v564
	var v565 int32
	_ = v565
	var v567 float64
	_ = v567
	var v571 float64
	_ = v571
	var v572 float64
	_ = v572
	var v576 float64
	_ = v576
	var v580 float64
	_ = v580
	var v585 float64
	_ = v585
	var v595 float64
	_ = v595
	var v598 float64
	_ = v598
	var v600 float64
	_ = v600
	var v603 float64
	_ = v603
	var v605 int32
	_ = v605
	var v609 float64
	_ = v609
	var v613 float64
	_ = v613
	var v614 float64
	_ = v614
	var v615 float64
	_ = v615
	var v619 float64
	_ = v619
	var v623 float64
	_ = v623
	var v635 float64
	_ = v635
	var v638 float64
	_ = v638
	var v640 float64
	_ = v640
	var v641 float64
	_ = v641
	var v651 float64
	_ = v651
	var v652 float64
	_ = v652
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v660 int32
	_ = v660
	var v665 int32
	_ = v665
	var v674 int32
	_ = v674
	var v677 int32
	_ = v677
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v693 int32
	_ = v693
	var v696 int32
	_ = v696
	var v702 int32
	_ = v702
	var v709 int32
	_ = v709
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v741 int32
	_ = v741
	var v746 int32
	_ = v746
	var v749 int32
	_ = v749
	var v756 int32
	_ = v756
	var v760 int32
	_ = v760
	var v763 float64
	_ = v763
	var v768 float64
	_ = v768
	var v769 int64
	_ = v769
	var v783 int32
	_ = v783
	var v801 int32
	_ = v801
	var v811 int32
	_ = v811
	var v815 int32
	_ = v815
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v824 float64
	_ = v824
	var v825 float64
	_ = v825
	var v831 float64
	_ = v831
	var v850 float64
	_ = v850
	var v851 int32
	_ = v851
	var v852 float64
	_ = v852
	var v853 float64
	_ = v853
	var v856 float64
	_ = v856
	var v861 float64
	_ = v861
	var v863 float64
	_ = v863
	var v864 float64
	_ = v864
	var v865 int32
	_ = v865
	var v868 float64
	_ = v868
	var v870 int32
	_ = v870
	var v876 float64
	_ = v876
	var v880 float64
	_ = v880
	var v883 float64
	_ = v883
	var v884 float64
	_ = v884
	var v885 float64
	_ = v885
	var v894 float64
	_ = v894
	var v898 float64
	_ = v898
	var v903 float64
	_ = v903
	v15 = int32(0)
	v25 = m.G0
	v27 = v25 - int32(80)
	m.G0 = v27
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v32 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v410 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_cost_index[0])))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v410 ^ int32(1)
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v30)+116))
	m.T0[v422].(func(*base.Module, int32, int32, float64, int32, int32, int32, int32, int32))(m, l1, l0, l2, v27+int32(48), v27+int32(40), v27+int32(32), v27+int32(24), v27)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L30
	} else {
		goto L87
	}
L2:
	;
	v33 = *(*float64)(unsafe.Add(mBase, uint32(v32)+8))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v33
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v30)+96))
	if v36 == int32(0) {
		v150 = v32
		v154 = v35
		v158 = v15
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v284 = *(*float64)(unsafe.Add(mBase, uint32(v31)+16))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v284
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v30)+96))
	if v286 == int32(0) {
		v403 = v15
		goto L1
	} else {
		goto L61
	}
L5:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v150)+16))
	if v160 == int32(0) {
		v281 = v15
		goto L33
	} else {
		goto L34
	}
L6:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	if v39 <= int32(0) {
		v150 = v32
		v154 = v35
		v158 = v15
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v57 = int32(0)
	v65 = v15
	goto L8
L8:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v67+v57<<(uint(int32(2))%32))))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+10)))
	if v72 != 0 {
		v129 = v65
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v150 = v135
	v154 = v134
	v158 = v129
	goto L5
L10:
	;
	v131 = v57 + int32(1)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	if v131 < v132 {
		v57 = v131
		v65 = v129
		goto L8
	} else {
		goto L32
	}
L11:
	;
	if v35 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	if v125 != 0 {
		v129 = v65
		goto L10
	} else {
		goto L29
	}
L13:
	;
	goto L12
L14:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v78 <= int32(0) {
		v125 = int32(0)
		goto L13
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v125 = int32(0)
	goto L13
L17:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v71)+60))
	v82 = int32(0)
	if v82 < v78 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v85 = v78
	goto L20
L19:
	;
	v85 = v82
	goto L20
L20:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
	v89 = int32(0)
	goto L21
L21:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v86+v89<<(uint(int32(2))%32))))
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+12)))
	if v99 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L16
L23:
	;
	v110 = v89 + int32(1)
	if v110 != v85 {
		v89 = v110
		goto L21
	} else {
		goto L28
	}
L24:
	;
	v100 = int32(1)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	if v71 == v101 {
		v125 = v100
		goto L13
	} else {
		goto L25
	}
L25:
	;
	if v81 == int32(0) {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v101)+60))
	if v105 == v81 {
		v125 = v100
		goto L13
	} else {
		goto L27
	}
L27:
	;
	goto L23
L28:
	;
	goto L22
L29:
	;
	v127 = F_lappend(m, v65, v71)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	return
L31:
	;
	v129 = v127
	goto L10
L32:
	;
	goto L9
L33:
	;
	v282 = F_list_concat(m, v158, v281)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L30
	} else {
		goto L60
	}
L34:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v160)+4))
	if v163 <= int32(0) {
		v281 = v15
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v181 = int32(0)
	v190 = v15
	goto L36
L36:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v160)+12))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v191+v181<<(uint(int32(2))%32))))
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195)+10)))
	if v196 != 0 {
		v253 = v190
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v281 = v253
	goto L33
L38:
	;
	v255 = v181 + int32(1)
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v160)+4))
	if v255 < v256 {
		v181 = v255
		v190 = v253
		goto L36
	} else {
		goto L59
	}
L39:
	;
	if v154 != 0 {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	if v249 != 0 {
		v253 = v190
		goto L38
	} else {
		goto L57
	}
L41:
	;
	goto L40
L42:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v154)+4))
	if v202 <= int32(0) {
		v249 = int32(0)
		goto L41
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v249 = int32(0)
	goto L41
L45:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v195)+60))
	v206 = int32(0)
	if v206 < v202 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v209 = v202
	goto L48
L47:
	;
	v209 = v206
	goto L48
L48:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v154)+12))
	v213 = int32(0)
	goto L49
L49:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v210+v213<<(uint(int32(2))%32))))
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+12)))
	if v223 != 0 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	goto L44
L51:
	;
	v234 = v213 + int32(1)
	if v234 != v209 {
		v213 = v234
		goto L49
	} else {
		goto L56
	}
L52:
	;
	v224 = int32(1)
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
	if v195 == v225 {
		v249 = v224
		goto L41
	} else {
		goto L53
	}
L53:
	;
	if v205 == int32(0) {
		goto L51
	} else {
		goto L54
	}
L54:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v225)+60))
	if v229 == v205 {
		v249 = v224
		goto L41
	} else {
		goto L55
	}
L55:
	;
	goto L51
L56:
	;
	goto L50
L57:
	;
	v251 = F_lappend(m, v190, v195)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L30
	} else {
		goto L58
	}
L58:
	;
	v253 = v251
	goto L38
L59:
	;
	goto L37
L60:
	;
	v403 = v282
	goto L1
L61:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v286)+4))
	if v289 <= int32(0) {
		v403 = v15
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v308 = int32(0)
	v312 = v15
	goto L63
L63:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v286)+12))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v318+v308<<(uint(int32(2))%32))))
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v322)+10)))
	if v323 != 0 {
		v380 = v312
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v403 = v380
	goto L1
L65:
	;
	v382 = v308 + int32(1)
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v286)+4))
	if v382 < v383 {
		v308 = v382
		v312 = v380
		goto L63
	} else {
		goto L86
	}
L66:
	;
	if v292 != 0 {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	if v376 != 0 {
		v380 = v312
		goto L65
	} else {
		goto L84
	}
L68:
	;
	goto L67
L69:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v292)+4))
	if v329 <= int32(0) {
		v376 = int32(0)
		goto L68
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v376 = int32(0)
	goto L68
L72:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v322)+60))
	v333 = int32(0)
	if v333 < v329 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v336 = v329
	goto L75
L74:
	;
	v336 = v333
	goto L75
L75:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v292)+12))
	v340 = int32(0)
	goto L76
L76:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v337+v340<<(uint(int32(2))%32))))
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349)+12)))
	if v350 != 0 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	goto L71
L78:
	;
	v361 = v340 + int32(1)
	if v361 != v336 {
		v340 = v361
		goto L76
	} else {
		goto L83
	}
L79:
	;
	v351 = int32(1)
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v349)+4))
	if v322 == v352 {
		v376 = v351
		goto L68
	} else {
		goto L80
	}
L80:
	;
	if v332 == int32(0) {
		goto L78
	} else {
		goto L81
	}
L81:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v352)+60))
	if v356 == v332 {
		v376 = v351
		goto L68
	} else {
		goto L82
	}
L82:
	;
	goto L78
L83:
	;
	goto L77
L84:
	;
	v378 = F_lappend(m, v312, v322)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L30
	} else {
		goto L85
	}
L85:
	;
	v380 = v378
	goto L65
L86:
	;
	goto L64
L87:
	;
	v425 = *(*float64)(unsafe.Add(mBase, uint32(v27)+40))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+96)) = v425
	v427 = *(*float64)(unsafe.Add(mBase, uint32(v27)+32))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+104)) = v427
	v429 = float64(1e+100)
	v430 = *(*float64)(unsafe.Add(mBase, uint32(v31)+120))
	v431 = base.F64_mul(v427, v430)
	if base.F64_gt(v431, v429)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v431)&int64(9223372036854775807))) != 0 {
		v444 = v429
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v445 = *(*float64)(unsafe.Add(mBase, uint32(v27)+48))
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v31)+72))
	F_get_tablespace_page_costs(m, v446, v27+int32(8), v27+int32(16))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L30
	} else {
		goto L91
	}
L89:
	;
	v440 = float64(1)
	if base.F64_le(v431, v440) != 0 {
		v444 = v440
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v444 = base.F64_nearest(v431)
	goto L88
L91:
	;
	if base.F64_gt(l2, float64(1)) != 0 {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	if l3 != 0 {
		goto L155
	} else {
		goto L156
	}
L93:
	;
	v455 = base.F64_mul(l2, v444)
	v456 = int32(1)
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v31)+116))
	if base.Ui32(v457) <= base.Ui32(v456) {
		goto L96
	} else {
		goto L97
	}
L94:
	;
	goto L95
L95:
	;
	v552 = int32(1)
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v31)+116))
	if base.Ui32(v553) <= base.Ui32(v552) {
		goto L127
	} else {
		goto L128
	}
L96:
	;
	v460 = v456
	goto L98
L97:
	;
	v460 = v457
	goto L98
L98:
	;
	v461 = base.F64_convert_i32_u(v460)
	v462 = base.F64_add(v461, v461)
	v463 = float64(1)
	v465 = *(*int32)(unsafe.Add(mBase, _c_F_cost_index[1]))
	v468 = *(*float64)(unsafe.Add(mBase, uint32(l1)+288))
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
	v471 = base.F64_add(v468, base.F64_convert_i32_u(v469))
	if base.F64_gt(v471, v463) != 0 {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	if v29 == int32(342) {
		goto L113
	} else {
		goto L114
	}
L100:
	;
	v475 = v471
	goto L102
L101:
	;
	v475 = v463
	goto L102
L102:
	;
	v476 = base.F64_div(base.F64_mul(v461, base.F64_convert_i32_s(v465)), v475)
	if base.F64_le(v476, float64(1)) != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v480 = v463
	goto L105
L104:
	;
	v480 = base.F64_ceil(v476)
	goto L105
L105:
	;
	v481 = base.F64_ge(v480, v461)
	if v481 != 0 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v484 = base.F64_div(base.F64_mul(v455, v462), base.F64_add(v462, v455))
	if base.F64_ge(v484, v461) != 0 {
		v503 = v461
		goto L99
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	v489 = base.F64_div(base.F64_mul(v462, v480), base.F64_sub(v462, v480))
	if base.F64_ge(v489, v455) != 0 {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	v503 = base.F64_ceil(v484)
	goto L99
L110:
	;
	v499 = base.F64_div(base.F64_mul(v455, v462), base.F64_add(v462, v455))
	goto L112
L111:
	;
	v499 = base.F64_add(v480, base.F64_div(base.F64_mul(base.F64_sub(v461, v480), base.F64_sub(v455, v489)), v461))
	goto L112
L112:
	;
	v503 = base.F64_ceil(v499)
	goto L99
L113:
	;
	v507 = *(*float64)(unsafe.Add(mBase, uint32(v31)+128))
	v511 = base.F64_ceil(base.F64_mul(v503, base.F64_sub(float64(1), v507)))
	goto L115
L114:
	;
	v511 = v503
	goto L115
L115:
	;
	v512 = *(*float64)(unsafe.Add(mBase, uint32(v27)+8))
	v513 = *(*float64)(unsafe.Add(mBase, uint32(v27)+32))
	v517 = base.F64_mul(l2, base.F64_ceil(base.F64_mul(v513, base.F64_convert_i32_u(v457))))
	if v481 != 0 {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	if v29 == int32(342) {
		goto L124
	} else {
		goto L125
	}
L117:
	;
	v520 = base.F64_div(base.F64_mul(v462, v517), base.F64_add(v462, v517))
	if base.F64_ge(v520, v461) != 0 {
		v537 = v461
		goto L116
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v525 = base.F64_div(base.F64_mul(v462, v480), base.F64_sub(v462, v480))
	if base.F64_ge(v525, v517) != 0 {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	v537 = base.F64_ceil(v520)
	goto L116
L121:
	;
	v535 = base.F64_div(base.F64_mul(v462, v517), base.F64_add(v462, v517))
	goto L123
L122:
	;
	v535 = base.F64_add(v480, base.F64_div(base.F64_mul(base.F64_sub(v461, v480), base.F64_sub(v517, v525)), v461))
	goto L123
L123:
	;
	v537 = base.F64_ceil(v535)
	goto L116
L124:
	;
	v545 = *(*float64)(unsafe.Add(mBase, uint32(v31)+128))
	v549 = base.F64_ceil(base.F64_mul(v537, base.F64_sub(float64(1), v545)))
	goto L126
L125:
	;
	v549 = v537
	goto L126
L126:
	;
	v638 = base.F64_div(base.F64_mul(v512, v549), l2)
	v640 = v511
	v641 = base.F64_div(base.F64_mul(v511, v512), l2)
	goto L92
L127:
	;
	v556 = v552
	goto L129
L128:
	;
	v556 = v553
	goto L129
L129:
	;
	v557 = base.F64_convert_i32_u(v556)
	v558 = base.F64_add(v557, v557)
	v559 = float64(1)
	v561 = *(*int32)(unsafe.Add(mBase, _c_F_cost_index[1]))
	v564 = *(*float64)(unsafe.Add(mBase, uint32(l1)+288))
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
	v567 = base.F64_add(v564, base.F64_convert_i32_u(v565))
	if base.F64_gt(v567, v559) != 0 {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	v600 = *(*float64)(unsafe.Add(mBase, uint32(v27)+32))
	v603 = base.F64_ceil(base.F64_mul(v600, base.F64_convert_i32_u(v553)))
	v605 = base.B2i32(v29 != int32(342))
	if v605 == int32(0) {
		goto L144
	} else {
		goto L145
	}
L131:
	;
	v571 = v567
	goto L133
L132:
	;
	v571 = v559
	goto L133
L133:
	;
	v572 = base.F64_div(base.F64_mul(v557, base.F64_convert_i32_s(v561)), v571)
	if base.F64_le(v572, float64(1)) != 0 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v576 = v559
	goto L136
L135:
	;
	v576 = base.F64_ceil(v572)
	goto L136
L136:
	;
	if base.F64_le(v557, v576) != 0 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v580 = base.F64_div(base.F64_mul(v444, v558), base.F64_add(v558, v444))
	if base.F64_ge(v580, v557) != 0 {
		v598 = v557
		goto L130
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	v585 = base.F64_div(base.F64_mul(v558, v576), base.F64_sub(v558, v576))
	if base.F64_ge(v585, v444) != 0 {
		goto L141
	} else {
		goto L142
	}
L140:
	;
	v598 = base.F64_ceil(v580)
	goto L130
L141:
	;
	v595 = base.F64_div(base.F64_mul(v444, v558), base.F64_add(v558, v444))
	goto L143
L142:
	;
	v595 = base.F64_add(v576, base.F64_div(base.F64_mul(base.F64_sub(v557, v576), base.F64_sub(v444, v585)), v557))
	goto L143
L143:
	;
	v598 = base.F64_ceil(v595)
	goto L130
L144:
	;
	v609 = *(*float64)(unsafe.Add(mBase, uint32(v31)+128))
	v613 = base.F64_ceil(base.F64_mul(v598, base.F64_sub(float64(1), v609)))
	goto L146
L145:
	;
	v613 = v598
	goto L146
L146:
	;
	v614 = *(*float64)(unsafe.Add(mBase, uint32(v27)+8))
	v615 = base.F64_mul(v613, v614)
	if v605 == int32(0) {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v619 = *(*float64)(unsafe.Add(mBase, uint32(v31)+128))
	v623 = base.F64_ceil(base.F64_mul(v603, base.F64_sub(float64(1), v619)))
	goto L149
L148:
	;
	v623 = v603
	goto L149
L149:
	;
	if base.F64_gt(v623, float64(0)) == int32(0) {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v638 = float64(0)
	v640 = v613
	v641 = v615
	goto L92
L151:
	;
	goto L152
L152:
	;
	if base.F64_gt(v623, float64(1)) == int32(0) {
		v638 = v614
		v640 = v613
		v641 = v615
		goto L92
	} else {
		goto L153
	}
L153:
	;
	v635 = *(*float64)(unsafe.Add(mBase, uint32(v27)+16))
	v638 = base.F64_add(base.F64_mul(base.F64_add(v623, float64(-1)), v635), v614)
	v640 = v613
	v641 = v615
	goto L92
L154:
	;
	m.G0 = v27 + int32(80)
	return
L155:
	;
	if v29 == int32(342) {
		goto L158
	} else {
		goto L159
	}
L156:
	;
	goto L157
L157:
	;
	v763 = float64(0)
	v768 = *(*float64)(unsafe.Add(mBase, uint32(v27)+24))
	v769 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v27)+64)) = v769
	*(*int32)(unsafe.Add(mBase, uint32(v27)+56)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v27)+72)) = v769
	if v403 == int32(0) {
		v831 = v763
		v850 = v763
		goto L202
	} else {
		goto L203
	}
L158:
	;
	v651 = float64(-1)
	goto L160
L159:
	;
	v651 = v640
	goto L160
L160:
	;
	v652 = *(*float64)(unsafe.Add(mBase, uint32(v27)))
	v654 = *(*int32)(unsafe.Add(mBase, _c_F_cost_index[2]))
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v31)+148))
	if v657 != int32(-1) {
		v746 = v657
		goto L163
	} else {
		goto L164
	}
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v756
	if v756 <= int32(0) {
		goto L154
	} else {
		goto L201
	}
L162:
	;
	goto L161
L163:
	;
	if v746 < v654 {
		goto L198
	} else {
		goto L199
	}
L164:
	;
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v660 != 0 {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v677 = int32(0)
	if base.F64_ge(v651, float64(0)) == v677 {
		v709 = v677
		goto L173
	} else {
		goto L174
	}
L166:
	;
	if base.F64_ge(v651, float64(0)) != 0 {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v665 = *(*int32)(unsafe.Add(mBase, _c_F_cost_index[3]))
	if base.F64_lt(v651, base.F64_convert_i32_s(v665)) != 0 {
		v756 = int32(0)
		goto L162
	} else {
		goto L170
	}
L168:
	;
	goto L169
L169:
	;
	if base.F64_ge(v652, float64(0)) == int32(0) {
		goto L165
	} else {
		goto L171
	}
L170:
	;
	goto L169
L171:
	;
	v674 = *(*int32)(unsafe.Add(mBase, _c_F_cost_index[4]))
	if base.F64_lt(v652, base.F64_convert_i32_s(v674)) != 0 {
		v756 = int32(0)
		goto L162
	} else {
		goto L172
	}
L172:
	;
	goto L165
L173:
	;
	if base.F64_ge(v652, float64(0)) == int32(0) {
		v746 = v709
		goto L163
	} else {
		goto L182
	}
L174:
	;
	v682 = int32(1)
	v684 = *(*int32)(unsafe.Add(mBase, _c_F_cost_index[3]))
	if v684 <= v682 {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v687 = v682
	goto L177
L176:
	;
	v687 = v684
	goto L177
L177:
	;
	v689 = v687
	v693 = int32(1)
	goto L178
L178:
	;
	v696 = v689 * int32(3)
	if base.F64_ge(v651, base.F64_convert_i32_u(v696)) == int32(0) {
		v709 = v693
		goto L173
	} else {
		goto L180
	}
L179:
	;
	v709 = v702
	goto L173
L180:
	;
	v702 = v693 + int32(1)
	if v696 < int32(715827883) {
		v689 = v696
		v693 = v702
		goto L178
	} else {
		goto L181
	}
L181:
	;
	goto L179
L182:
	;
	v715 = int32(1)
	v717 = *(*int32)(unsafe.Add(mBase, _c_F_cost_index[4]))
	if v717 <= v715 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v720 = v715
	goto L185
L184:
	;
	v720 = v717
	goto L185
L185:
	;
	v722 = v720
	v727 = int32(1)
	goto L186
L186:
	;
	v729 = v722 * int32(3)
	if base.F64_le(base.F64_convert_i32_u(v729), v652) != 0 {
		goto L188
	} else {
		goto L189
	}
L187:
	;
	if v709 < v736 {
		goto L192
	} else {
		goto L193
	}
L188:
	;
	v733 = v727 + int32(1)
	if v729 < int32(715827883) {
		v722 = v729
		v727 = v733
		goto L186
	} else {
		goto L191
	}
L189:
	;
	v736 = v727
	goto L190
L190:
	;
	goto L187
L191:
	;
	v736 = v733
	goto L190
L192:
	;
	v738 = v709
	goto L194
L193:
	;
	v738 = v736
	goto L194
L194:
	;
	if int32(0) < v709 {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v741 = v738
	goto L197
L196:
	;
	v741 = v736
	goto L197
L197:
	;
	v746 = v741
	goto L163
L198:
	;
	v749 = v746
	goto L200
L199:
	;
	v749 = v654
	goto L200
L200:
	;
	v756 = v749
	goto L162
L201:
	;
	v760 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)) = uint8(v760)
	goto L157
L202:
	;
	v851 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v852 = *(*float64)(unsafe.Add(mBase, uint32(v851)+24))
	v853 = *(*float64)(unsafe.Add(mBase, uint32(l0)+32))
	v856 = *(*float64)(unsafe.Add(mBase, _c_F_cost_index[5]))
	v861 = base.F64_add(base.F64_mul(v852, v853), base.F64_add(base.F64_mul(base.F64_add(v831, v856), v444), float64(0)))
	v863 = *(*float64)(unsafe.Add(mBase, uint32(v851)+16))
	v864 = base.F64_add(base.F64_add(base.F64_add(v445, v763), v850), v863)
	v865 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if int32(0) < v865 {
		goto L209
	} else {
		goto L210
	}
L203:
	;
	v783 = *(*int32)(unsafe.Add(mBase, uint32(v403)+4))
	if v783 <= int32(0) {
		v831 = v763
		v850 = float64(0)
		goto L202
	} else {
		goto L204
	}
L204:
	;
	v801 = int32(0)
	goto L205
L205:
	;
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v403)+12))
	v815 = *(*int32)(unsafe.Add(mBase, uint32(v811+v801<<(uint(int32(2))%32))))
	v818 = F_cost_qual_eval_walker(m, v815, v27+int32(56))
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L30
	} else {
		goto L207
	}
L206:
	;
	v824 = *(*float64)(unsafe.Add(mBase, uint32(v27)+72))
	v825 = *(*float64)(unsafe.Add(mBase, uint32(v27)+64))
	v831 = v824
	v850 = v825
	goto L202
L207:
	;
	v821 = v801 + int32(1)
	v822 = *(*int32)(unsafe.Add(mBase, uint32(v403)+4))
	if v821 < v822 {
		v801 = v821
		goto L205
	} else {
		goto L208
	}
L208:
	;
	goto L206
L209:
	;
	v868 = base.F64_convert_i32_u(v865)
	v870 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_cost_index[6])))
	if v870 == int32(1) {
		goto L213
	} else {
		goto L214
	}
L210:
	;
	v903 = v861
	goto L211
L211:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = v864
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = base.F64_add(v864, base.F64_add(base.F64_add(base.F64_add(base.F64_sub(v425, v445), v763), base.F64_add(base.F64_mul(base.F64_mul(v768, v768), base.F64_sub(v638, v641)), v641)), v903))
	goto L154
L212:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v898
	v903 = base.F64_div(v861, v883)
	goto L211
L213:
	;
	v876 = base.F64_add(base.F64_mul(v868, float64(-0.3)), float64(1))
	if base.F64_gt(v876, float64(0)) != 0 {
		goto L216
	} else {
		goto L217
	}
L214:
	;
	v883 = v868
	goto L215
L215:
	;
	v884 = float64(1e+100)
	v885 = base.F64_div(v853, v883)
	if base.F64_gt(v885, v884)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v885)&int64(9223372036854775807))) != 0 {
		v898 = v884
		goto L212
	} else {
		goto L219
	}
L216:
	;
	v880 = v876
	goto L218
L217:
	;
	v880 = math.Float64frombits(uint64(0x8000000000000000))
	goto L218
L218:
	;
	v883 = base.F64_add(v880, v868)
	goto L215
L219:
	;
	v894 = float64(1)
	if base.F64_le(v885, v894) != 0 {
		v898 = v894
		goto L212
	} else {
		goto L220
	}
L220:
	;
	v898 = base.F64_nearest(v885)
	goto L212
}
func F_index_am_handler_in(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13852(m, l0, int32(_a_F_index_am_handler_in_0), int32(371), int32(_a_F_index_am_handler_in_1), int32(_a_F_index_am_handler_in_2), int32(_a_F_index_am_handler_in_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_index_form_tuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_index_form_tuple[0]))
	v6 = F_index_form_tuple_context(m, l0, l1, l2, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_index_get_partition(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = F_RelationGetIndexList(m, l0)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L4
	} else {
		goto L20
	}
L2:
	;
	F_list_free(m, v12)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L4
	} else {
		goto L19
	}
L3:
	;
	v66 = int32(0)
	goto L2
L4:
	;
	return int32(0)
L5:
	;
	if v12 == int32(0) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v18 = int32(0)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v19 <= v18 {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v22 = v18
	goto L8
L8:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v30+v22<<(uint(int32(2))%32))))
	v35 = F_SearchSysCache1(m, int32(57), v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L10
	}
L9:
	;
	goto L3
L10:
	;
	if v35 == int32(0) {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v35)+16))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+22)))
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39+v40)+131)))
	F_ReleaseCatCache(m, v35)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	if v42 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v48 = F_get_partition_parent(m, v34, int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L4
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v52 = v22 + int32(1)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v52 < v53 {
		v22 = v52
		goto L8
	} else {
		goto L18
	}
L16:
	;
	if v48 == l1 {
		v66 = v34
		goto L2
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	goto L9
L19:
	;
	m.G0 = v10 + int32(16)
	return v66
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v34
	F_errmsg_internal(m, int32(_a_F_index_get_partition_0), v10)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(_a_F_index_get_partition_1), int32(190), int32(_a_F_index_get_partition_2))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_index_getprocid(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v6 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5)+6)))
	v10 = int32(2)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v4+v6*(l1-int32(1))<<(uint(v10)%32)+l2<<(uint(v10)%32)-int32(4))))
	return v18
}
func F_index_markpos(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+204))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+112))
	if v10 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+48))
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(_a_F_index_markpos_0)
			*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v18 + int32(4)
			F_errmsg_internal(m, int32(_a_F_index_markpos_1), v6)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_index_markpos_2), int32(415), int32(_a_F_index_markpos_3))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		m.T0[v10].(func(*base.Module, int32))(m, l0)
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return
		} else {
			m.G0 = v6 + int32(16)
			return
		}
	}
}
func F_index_parallelrescan(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v3 != 0 {
		v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
		v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+188))
		v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+48))
		m.T0[v6].(func(*base.Module, int32))(m, v3)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+204))
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+128))
			if v11 != 0 {
				m.T0[v11].(func(*base.Module, int32))(m, l0)
				mBase = m.M
				v13 = m.ExcPending
				if v13 != 0 {
					return
				} else {
					return
				}
			} else {
				return
			}
		}
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+204))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+128))
		if v11 != 0 {
			m.T0[v11].(func(*base.Module, int32))(m, l0)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				return
			}
		} else {
			return
		}
	}
}
func F_index_parallelscan_estimate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_index_parallelscan_estimate[0]))
	if v15 != v13 {
		v18 = *(*int32)(unsafe.Add(mBase, _c_F_index_parallelscan_estimate[1]))
		v19 = F_list_member_ptr(m, v18, v13)
		mBase = m.M
		v21 = v19
	} else {
		v21 = int32(1)
	}
	if v21 == int32(0) {
		v25 = F_EstimateSnapshotSpace(m, l3)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			v29 = F_add_size(m, int32(32), v25)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				v34 = (v29 + int32(7)) & int32(-8)
				if l4 != 0 {
					v39 = F_add_size(m, v34, l6<<(uint(int32(3))%32)+int32(8))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						v45 = (v39 + int32(7)) & int32(-8)
						if l5 == int32(0) {
							v56 = v45
							m.G0 = v11 + int32(16)
							return v56
						} else {
							v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
							v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+120))
							if v49 == int32(0) {
								v56 = v45
								m.G0 = v11 + int32(16)
								return v56
							} else {
								v52 = m.T0[v49].(func(*base.Module, int32, int32, int32) int32)(m, l0, l1, l2)
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									v54 = F_add_size(m, v45, v52)
									mBase = m.M
									v55 = m.ExcPending
									if v55 != 0 {
										return int32(0)
									} else {
										v56 = v54
										m.G0 = v11 + int32(16)
										return v56
									}
								}
							}
						}
					}
				} else {
					v45 = v34
					if l5 == int32(0) {
						v56 = v45
						m.G0 = v11 + int32(16)
						return v56
					} else {
						v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+120))
						if v49 == int32(0) {
							v56 = v45
							m.G0 = v11 + int32(16)
							return v56
						} else {
							v52 = m.T0[v49].(func(*base.Module, int32, int32, int32) int32)(m, l0, l1, l2)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								v54 = F_add_size(m, v45, v52)
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return int32(0)
								} else {
									v56 = v54
									m.G0 = v11 + int32(16)
									return v56
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
		v65 = m.ExcPending
		if v65 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v68 = m.ExcPending
			if v68 != 0 {
				return int32(0)
			} else {
				v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = v69 + int32(4)
				F_errmsg(m, int32(_a_F_index_parallelscan_estimate_0), v11)
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_index_parallelscan_estimate_1), int32(469), int32(_a_F_index_parallelscan_estimate_2))
					mBase = m.M
					v80 = m.ExcPending
					if v80 != 0 {
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
func F_index_reloptions(m *base.Module, l0 int32, l1 int32) {
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	if l1 == int32(0) {
		return
	} else {
		v6 = m.T0[l0].(func(*base.Module, int32, int32) int32)(m, l1, int32(1))
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			return
		}
	}
}
func F_index_store_float8_orderby_distances(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 float64
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 float64
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	v4 = l3
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+84)) = uint8(v4)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if int32(0) < v8 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v15 = int32(0)
	goto L4
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v19 = v15 << (uint(int32(2)) % 32)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1+v19)))
	switch v21 - int32(700) {
	case 0:
		goto L9
	case 1:
		goto L10
	default:
		goto L8
	}
L5:
	;
	goto L3
L6:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	*(*uint8)(unsafe.Add(mBase, uint32(v85+v15))) = uint8(v84)
	v89 = v15 + int32(1)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v89 < v90 {
		v15 = v89
		goto L4
	} else {
		goto L26
	}
L7:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v78+v19))) = int32(0)
	v84 = int32(1)
	goto L6
L8:
	;
	v60 = int32(1)
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+84)))
	if v61 != v60 {
		v84 = v60
		goto L6
	} else {
		goto L22
	}
L9:
	;
	if l2 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L10:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v15))))
	if v26 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v29+v19)))
	F_pfree(m, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	if l2 == int32(0) {
		goto L7
	} else {
		goto L16
	}
L14:
	;
	return
L15:
	;
	goto L13
L16:
	;
	v38 = l2 + v15<<(uint(int32(4))%32)
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+8)))
	if v39 != 0 {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	v40 = *(*float64)(unsafe.Add(mBase, uint32(v38)))
	v41 = F_Float8GetDatum(m, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v43+v19))) = v41
	v84 = int32(0)
	goto L6
L19:
	;
	goto L7
L20:
	;
	v51 = l2 + v15<<(uint(int32(4))%32)
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+8)))
	if v52 != 0 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v55 = *(*float64)(unsafe.Add(mBase, uint32(v51)))
	*(*float32)(unsafe.Add(mBase, uint32(v53+v19))) = base.F32_demote_f64(v55)
	v84 = int32(0)
	goto L6
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L14
	} else {
		goto L23
	}
L23:
	;
	F_errmsg_internal(m, int32(_a_F_index_store_float8_orderby_distances_0), int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L14
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(_a_F_index_store_float8_orderby_distances_1), int32(1030), int32(_a_F_index_store_float8_orderby_distances_2))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L14
	} else {
		goto L25
	}
L25:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L26:
	;
	goto L5
}
func F_index_update_stats(m *base.Module, l0 int32, l1 int32, l2 float64) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v22 float32
	_ = v22
	var v28 float64
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 float32
	_ = v117
	var v118 float32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	v2 = l1
	v4 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(80)
	m.G0 = v13
	*(*int32)(unsafe.Add(mBase, uint32(v13)+76)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(v13)+72)) = v4
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if base.F64_ne(l2, float64(0)) != 0 {
		v28 = l2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v31 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_index_update_stats[0])))
	v32 = int32(1)
	v36 = (v31 ^ v32) & base.F64_ge(v28, float64(0))
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+119)))
	v39 = v37 - int32(109)
	if base.B2i32(base.Ui32(int32(7)) < base.Ui32(v39))|base.B2i32(v32<<(uint(v39)%32)&int32(161) == int32(0)) != 0 {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	v22 = *(*float32)(unsafe.Add(mBase, uint32(v19)+100))
	if base.F32_lt(v22, float32(0)) == int32(0) {
		v28 = l2
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v28 = float64(-1)
	goto L1
L4:
	;
	v86 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L13
	} else {
		goto L17
	}
L5:
	;
	v68 = F_RelationGetNumberOfBlocksInFork(m, l0, int32(0))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L13
	} else {
		goto L14
	}
L6:
	;
	if v36 == int32(0) {
		v82 = v4
		v83 = v4
		goto L4
	} else {
		goto L12
	}
L7:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_index_update_stats[1])))
	v52 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_index_update_stats[2])))
	goto L8
L8:
	;
	if v50&v52&int32(1) == int32(0) {
		v82 = v4
		v83 = v4
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v58 == int32(0) {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+16)))
	if v36&v61 != 0 {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v82 = v4
	v83 = v4
	goto L4
L12:
	;
	goto L5
L13:
	;
	return
L14:
	;
	v70 = int32(1)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+119)))
	if v72 == int32(105) {
		v82 = v68
		v83 = v70
		goto L4
	} else {
		goto L15
	}
L15:
	;
	F_visibilitymap_count(m, l0, v13+int32(76), v13+int32(72))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v82 = v68
	v83 = v70
	goto L4
L17:
	;
	v89 = v13 + int32(16)
	F_ScanKeyInit(m, v89, int32(1), int32(3), int32(184), v29)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L13
	} else {
		goto L18
	}
L18:
	;
	F_systable_inplace_update_begin(m, v86, int32(2662), v89, v13+int32(12), v13+int32(8))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L13
	} else {
		goto L19
	}
L19:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v102 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+16))
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+22)))
	v105 = v103 + v104
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+116)))
	if v2 != v106 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L13
	} else {
		goto L47
	}
L23:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v105)+116)) = uint8(v2)
	goto L25
L24:
	;
	goto L25
L25:
	;
	v109 = base.B2i32(v2 != v106)
	if v83 == int32(0) {
		v133 = v109
		goto L29
	} else {
		goto L30
	}
L26:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	F_pfree(m, v154)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L13
	} else {
		goto L45
	}
L27:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	F_systable_inplace_update_cancel(m, v145)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L13
	} else {
		goto L43
	}
L28:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	F_systable_inplace_update_finish(m, v141, v142)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L13
	} else {
		goto L42
	}
L29:
	;
	if v133 == int32(0) {
		goto L27
	} else {
		goto L41
	}
L30:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v105)+96))
	if v82 != v112 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105)+96)) = v82
	v116 = int32(1)
	goto L33
L32:
	;
	v116 = v109
	goto L33
L33:
	;
	v117 = base.F32_demote_f64(v28)
	v118 = *(*float32)(unsafe.Add(mBase, uint32(v105)+100))
	if base.F32_ne(v117, v118) != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v105)+100)) = v117
	v122 = int32(1)
	goto L36
L35:
	;
	v122 = v116
	goto L36
L36:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v13)+76))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v105)+104))
	if v123 != v124 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105)+104)) = v123
	v128 = int32(1)
	goto L39
L38:
	;
	v128 = v122
	goto L39
L39:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v13)+72))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v105)+108))
	if v129 == v130 {
		v133 = v128
		goto L29
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105)+108)) = v129
	goto L28
L41:
	;
	goto L28
L42:
	;
	goto L26
L43:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	F_CacheInvalidateRelcacheByTuple(m, v148)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L13
	} else {
		goto L44
	}
L44:
	;
	goto L26
L45:
	;
	F_relation_close(m, v86, int32(3))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L13
	} else {
		goto L46
	}
L46:
	;
	m.G0 = v13 + int32(80)
	return
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v29
	F_errmsg_internal(m, int32(_a_F_index_update_stats_0), v13)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L13
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_index_update_stats_1), int32(2918), int32(_a_F_index_update_stats_2))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L13
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_match_index_to_operand(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
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
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	if l0 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L14
	} else {
		goto L62
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L14
	} else {
		goto L59
	}
L3:
	;
	return int32(0)
L4:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l2)+84))
	if v92 != 0 {
		goto L35
	} else {
		goto L36
	}
L5:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v68+l1<<(uint(int32(2))%32))))
	if v72 == int32(0) {
		v86 = v46
		v91 = v68
		goto L4
	} else {
		goto L30
	}
L6:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v62+l1<<(uint(int32(2))%32))))
	if v66 != 0 {
		goto L3
	} else {
		goto L29
	}
L7:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v9 == int32(319) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v46 = v40
	goto L25
L9:
	;
	v23 = l0
	goto L17
L10:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v12 == int32(0) {
		goto L9
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v17 = F_expression_tree_walker_impl(m, l0, int32(825), int32(0))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L12
L14:
	;
	return int32(0)
L15:
	;
	if v17 == int32(0) {
		v40 = l0
		goto L8
	} else {
		goto L16
	}
L16:
	;
	goto L9
L17:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v29 != int32(319) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v36 = F_expression_tree_mutator_impl(m, v23, int32(826), int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L14
	} else {
		goto L23
	}
L19:
	;
	goto L18
L20:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	if v32 != 0 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v33 != 0 {
		v23 = v33
		goto L17
	} else {
		goto L22
	}
L22:
	;
	goto L6
L23:
	;
	if v36 == int32(0) {
		goto L6
	} else {
		goto L24
	}
L24:
	;
	v40 = v36
	goto L8
L25:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	if v52 != int32(27) {
		goto L5
	} else {
		goto L27
	}
L26:
	;
	goto L6
L27:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v55 != 0 {
		v46 = v55
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v86 = int32(0)
	v91 = v62
	goto L4
L30:
	;
	if v52 != int32(6) {
		goto L3
	} else {
		goto L31
	}
L31:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+68))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v78 != v79 {
		goto L3
	} else {
		goto L32
	}
L32:
	;
	v81 = int32(*(*int16)(unsafe.Add(mBase, uint32(v46)+8)))
	if v72 != v81 {
		goto L3
	} else {
		goto L33
	}
L33:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v46)+24))
	if v83 != 0 {
		goto L3
	} else {
		goto L34
	}
L34:
	;
	return int32(1)
L35:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+12))
	v95 = v93
	goto L37
L36:
	;
	v95 = int32(0)
	goto L37
L37:
	;
	if int32(0) < l1 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v101 = v95
	v102 = int32(0)
	goto L41
L39:
	;
	v129 = v95
	goto L40
L40:
	;
	if v129 == int32(0) {
		goto L1
	} else {
		goto L51
	}
L41:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v91+v102<<(uint(int32(2))%32))))
	if v108 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v129 = v123
	goto L40
L43:
	;
	if v101 == int32(0) {
		goto L2
	} else {
		goto L46
	}
L44:
	;
	v123 = v101
	goto L45
L45:
	;
	v125 = v102 + int32(1)
	if v125 != l1 {
		v101 = v123
		v102 = v125
		goto L41
	} else {
		goto L50
	}
L46:
	;
	v114 = v101 + int32(4)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v92)+12))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
	if base.Ui32(v114) < base.Ui32(v116+v117<<(uint(int32(2))%32)) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v122 = v114
	goto L49
L48:
	;
	v122 = int32(0)
	goto L49
L49:
	;
	v123 = v122
	goto L45
L50:
	;
	goto L42
L51:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	if v135 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v144 = F_equal(m, v143, v86)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L14
	} else {
		goto L57
	}
L53:
	;
	v143 = int32(0)
	goto L52
L54:
	;
	goto L55
L55:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v135)))
	if v139 != int32(27) {
		v143 = v135
		goto L52
	} else {
		goto L56
	}
L56:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v135)+4))
	v143 = v142
	goto L52
L57:
	;
	if v144 == int32(0) {
		goto L3
	} else {
		goto L58
	}
L58:
	;
	return int32(1)
L59:
	;
	F_errmsg_internal(m, int32(_a_F_match_index_to_operand_0), int32(0))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L14
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_match_index_to_operand_1), int32(_a_F_match_index_to_operand_2), int32(_a_F_match_index_to_operand_3))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L14
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L62:
	;
	F_errmsg_internal(m, int32(_a_F_match_index_to_operand_0), int32(0))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L14
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_match_index_to_operand_1), int32(_a_F_match_index_to_operand_4), int32(_a_F_match_index_to_operand_3))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L14
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
