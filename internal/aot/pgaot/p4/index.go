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
		v12 = F_ExecScan(m, l0, int32(769), int32(770))
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
			v12 = F_ExecScan(m, l0, int32(769), int32(770))
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
				v12 = F_ExecScan(m, l0, int32(769), int32(770))
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
			v15 = int32(773)
		} else {
			v15 = int32(774)
		}
		v17 = F_ExecScan(m, l0, v15, int32(775))
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
				v15 = int32(773)
			} else {
				v15 = int32(774)
			}
			v17 = F_ExecScan(m, l0, v15, int32(775))
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
					v15 = int32(773)
				} else {
					v15 = int32(774)
				}
				v17 = F_ExecScan(m, l0, v15, int32(775))
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
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = F_OidFunctionCall0Coll(m, l0)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = base.I32_wrap_i64(v8)
		if v12 != 0 {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
			if v13 == int32(444) {
				m.G0 = v6 + int32(16)
				return v12
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
					F_errmsg_internal(m, int32(_a_F_GetIndexAmRoutine_0), v6)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_GetIndexAmRoutine_1), int32(43), int32(_a_F_GetIndexAmRoutine_2))
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
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
				F_errmsg_internal(m, int32(_a_F_GetIndexAmRoutine_0), v6)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_GetIndexAmRoutine_1), int32(43), int32(_a_F_GetIndexAmRoutine_2))
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
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = F_SearchSysCache1(m, int32(34), base.I64_extend_i32_u(l0))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if v12 == int32(0) {
			if l1 != 0 {
				v38 = int32(0)
				m.G0 = v8 + int32(16)
				return v38
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
					F_errmsg_internal(m, int32(_a_F_IndexGetRelation_0), v8)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_IndexGetRelation_1), int32(3716), int32(_a_F_IndexGetRelation_2))
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
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
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
			v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+22)))
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v31+v32)+4))
			F_ReleaseCatCache(m, v12)
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int32(0)
			} else {
				v38 = v34
				m.G0 = v8 + int32(16)
				return v38
			}
		}
	}
}
func F_IndexSupportsBackwardScan(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
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
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v10 = F_SearchSysCache1(m, int32(57), base.I64_extend_i32_u(l0))
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
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
				F_errmsg_internal(m, int32(_a_F_IndexSupportsBackwardScan_0), v6)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_IndexSupportsBackwardScan_1), int32(614), int32(_a_F_IndexSupportsBackwardScan_2))
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
				F_ReleaseCatCache(m, v10)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					m.G0 = v6 + int32(16)
					return v36
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
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
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
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
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
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
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
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+228))
	if v55 == int32(0) {
		v90 = v53
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
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
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
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
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
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
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
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
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
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
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191)+107)))
	if v202 != int32(1) {
		goto L55
	} else {
		goto L62
	}
L62:
	;
	v205 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v191)+96)) = v205
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
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
	var v409 int64
	_ = v409
	var v414 int64
	_ = v414
	var v417 int64
	_ = v417
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v434 float64
	_ = v434
	var v436 float64
	_ = v436
	var v438 float64
	_ = v438
	var v439 float64
	_ = v439
	var v440 float64
	_ = v440
	var v449 float64
	_ = v449
	var v453 float64
	_ = v453
	var v454 float64
	_ = v454
	var v455 int32
	_ = v455
	var v461 int32
	_ = v461
	var v464 float64
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v470 float64
	_ = v470
	var v471 float64
	_ = v471
	var v472 float64
	_ = v472
	var v474 int32
	_ = v474
	var v477 float64
	_ = v477
	var v478 int32
	_ = v478
	var v480 float64
	_ = v480
	var v484 float64
	_ = v484
	var v485 float64
	_ = v485
	var v489 float64
	_ = v489
	var v490 int32
	_ = v490
	var v493 float64
	_ = v493
	var v498 float64
	_ = v498
	var v508 float64
	_ = v508
	var v512 float64
	_ = v512
	var v516 float64
	_ = v516
	var v520 float64
	_ = v520
	var v521 float64
	_ = v521
	var v522 float64
	_ = v522
	var v526 float64
	_ = v526
	var v529 float64
	_ = v529
	var v534 float64
	_ = v534
	var v544 float64
	_ = v544
	var v546 float64
	_ = v546
	var v554 float64
	_ = v554
	var v558 float64
	_ = v558
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v566 float64
	_ = v566
	var v567 float64
	_ = v567
	var v568 float64
	_ = v568
	var v570 int32
	_ = v570
	var v573 float64
	_ = v573
	var v574 int32
	_ = v574
	var v576 float64
	_ = v576
	var v580 float64
	_ = v580
	var v581 float64
	_ = v581
	var v585 float64
	_ = v585
	var v589 float64
	_ = v589
	var v594 float64
	_ = v594
	var v604 float64
	_ = v604
	var v607 float64
	_ = v607
	var v609 float64
	_ = v609
	var v612 float64
	_ = v612
	var v614 int32
	_ = v614
	var v618 float64
	_ = v618
	var v622 float64
	_ = v622
	var v623 float64
	_ = v623
	var v624 float64
	_ = v624
	var v628 float64
	_ = v628
	var v632 float64
	_ = v632
	var v644 float64
	_ = v644
	var v647 float64
	_ = v647
	var v649 float64
	_ = v649
	var v650 float64
	_ = v650
	var v660 float64
	_ = v660
	var v661 float64
	_ = v661
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	var v669 int32
	_ = v669
	var v674 int32
	_ = v674
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v702 int32
	_ = v702
	var v705 int32
	_ = v705
	var v711 int32
	_ = v711
	var v718 int32
	_ = v718
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v750 int32
	_ = v750
	var v755 int32
	_ = v755
	var v758 int32
	_ = v758
	var v765 int32
	_ = v765
	var v769 int32
	_ = v769
	var v772 float64
	_ = v772
	var v777 float64
	_ = v777
	var v778 int64
	_ = v778
	var v792 int32
	_ = v792
	var v810 int32
	_ = v810
	var v820 int32
	_ = v820
	var v824 int32
	_ = v824
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v833 float64
	_ = v833
	var v834 float64
	_ = v834
	var v840 float64
	_ = v840
	var v859 float64
	_ = v859
	var v860 int32
	_ = v860
	var v861 float64
	_ = v861
	var v862 float64
	_ = v862
	var v865 float64
	_ = v865
	var v870 float64
	_ = v870
	var v872 float64
	_ = v872
	var v873 float64
	_ = v873
	var v874 int32
	_ = v874
	var v877 float64
	_ = v877
	var v879 int32
	_ = v879
	var v885 float64
	_ = v885
	var v889 float64
	_ = v889
	var v892 float64
	_ = v892
	var v893 float64
	_ = v893
	var v894 float64
	_ = v894
	var v903 float64
	_ = v903
	var v907 float64
	_ = v907
	var v912 float64
	_ = v912
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
	v409 = *(*int64)(unsafe.Add(mBase, uint32(v31)+32))
	if v29 == int32(346) {
		goto L87
	} else {
		goto L88
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
	v414 = int64(-5)
	goto L89
L88:
	;
	v414 = int64(-3)
	goto L89
L89:
	;
	if l3 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v417 = int64(-1)
	goto L92
L91:
	;
	v417 = int64(-262145)
	goto L92
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = base.B2i32(v409|v414&v417 != int64(-1))
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v30)+116))
	m.T0[v431].(func(*base.Module, int32, int32, float64, int32, int32, int32, int32, int32))(m, l1, l0, l2, v27+int32(48), v27+int32(40), v27+int32(32), v27+int32(24), v27)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L30
	} else {
		goto L93
	}
L93:
	;
	v434 = *(*float64)(unsafe.Add(mBase, uint32(v27)+40))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+96)) = v434
	v436 = *(*float64)(unsafe.Add(mBase, uint32(v27)+32))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+104)) = v436
	v438 = float64(1e+100)
	v439 = *(*float64)(unsafe.Add(mBase, uint32(v31)+128))
	v440 = base.F64_mul(v436, v439)
	if base.F64_gt(v440, v438)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v440)&int64(9223372036854775807))) != 0 {
		v453 = v438
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v454 = *(*float64)(unsafe.Add(mBase, uint32(v27)+48))
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v31)+80))
	F_get_tablespace_page_costs(m, v455, v27+int32(8), v27+int32(16))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L30
	} else {
		goto L97
	}
L95:
	;
	v449 = float64(1)
	if base.F64_le(v440, v449) != 0 {
		v453 = v449
		goto L94
	} else {
		goto L96
	}
L96:
	;
	v453 = base.F64_nearest(v440)
	goto L94
L97:
	;
	if base.F64_gt(l2, float64(1)) != 0 {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	if l3 != 0 {
		goto L161
	} else {
		goto L162
	}
L99:
	;
	v464 = base.F64_mul(l2, v453)
	v465 = int32(1)
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v31)+124))
	if base.Ui32(v466) <= base.Ui32(v465) {
		goto L102
	} else {
		goto L103
	}
L100:
	;
	goto L101
L101:
	;
	v561 = int32(1)
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v31)+124))
	if base.Ui32(v562) <= base.Ui32(v561) {
		goto L133
	} else {
		goto L134
	}
L102:
	;
	v469 = v465
	goto L104
L103:
	;
	v469 = v466
	goto L104
L104:
	;
	v470 = base.F64_convert_i32_u(v469)
	v471 = base.F64_add(v470, v470)
	v472 = float64(1)
	v474 = *(*int32)(unsafe.Add(mBase, _c_F_cost_index[0]))
	v477 = *(*float64)(unsafe.Add(mBase, uint32(l1)+304))
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
	v480 = base.F64_add(v477, base.F64_convert_i32_u(v478))
	if base.F64_gt(v480, v472) != 0 {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	if v29 == int32(346) {
		goto L119
	} else {
		goto L120
	}
L106:
	;
	v484 = v480
	goto L108
L107:
	;
	v484 = v472
	goto L108
L108:
	;
	v485 = base.F64_div(base.F64_mul(v470, base.F64_convert_i32_s(v474)), v484)
	if base.F64_le(v485, float64(1)) != 0 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v489 = v472
	goto L111
L110:
	;
	v489 = base.F64_ceil(v485)
	goto L111
L111:
	;
	v490 = base.F64_ge(v489, v470)
	if v490 != 0 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v493 = base.F64_div(base.F64_mul(v464, v471), base.F64_add(v471, v464))
	if base.F64_ge(v493, v470) != 0 {
		v512 = v470
		goto L105
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	v498 = base.F64_div(base.F64_mul(v471, v489), base.F64_sub(v471, v489))
	if base.F64_ge(v498, v464) != 0 {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	v512 = base.F64_ceil(v493)
	goto L105
L116:
	;
	v508 = base.F64_div(base.F64_mul(v464, v471), base.F64_add(v471, v464))
	goto L118
L117:
	;
	v508 = base.F64_add(v489, base.F64_div(base.F64_mul(base.F64_sub(v470, v489), base.F64_sub(v464, v498)), v470))
	goto L118
L118:
	;
	v512 = base.F64_ceil(v508)
	goto L105
L119:
	;
	v516 = *(*float64)(unsafe.Add(mBase, uint32(v31)+136))
	v520 = base.F64_ceil(base.F64_mul(v512, base.F64_sub(float64(1), v516)))
	goto L121
L120:
	;
	v520 = v512
	goto L121
L121:
	;
	v521 = *(*float64)(unsafe.Add(mBase, uint32(v27)+8))
	v522 = *(*float64)(unsafe.Add(mBase, uint32(v27)+32))
	v526 = base.F64_mul(l2, base.F64_ceil(base.F64_mul(v522, base.F64_convert_i32_u(v466))))
	if v490 != 0 {
		goto L123
	} else {
		goto L124
	}
L122:
	;
	if v29 == int32(346) {
		goto L130
	} else {
		goto L131
	}
L123:
	;
	v529 = base.F64_div(base.F64_mul(v471, v526), base.F64_add(v471, v526))
	if base.F64_ge(v529, v470) != 0 {
		v546 = v470
		goto L122
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	v534 = base.F64_div(base.F64_mul(v471, v489), base.F64_sub(v471, v489))
	if base.F64_ge(v534, v526) != 0 {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	v546 = base.F64_ceil(v529)
	goto L122
L127:
	;
	v544 = base.F64_div(base.F64_mul(v471, v526), base.F64_add(v471, v526))
	goto L129
L128:
	;
	v544 = base.F64_add(v489, base.F64_div(base.F64_mul(base.F64_sub(v470, v489), base.F64_sub(v526, v534)), v470))
	goto L129
L129:
	;
	v546 = base.F64_ceil(v544)
	goto L122
L130:
	;
	v554 = *(*float64)(unsafe.Add(mBase, uint32(v31)+136))
	v558 = base.F64_ceil(base.F64_mul(v546, base.F64_sub(float64(1), v554)))
	goto L132
L131:
	;
	v558 = v546
	goto L132
L132:
	;
	v647 = base.F64_div(base.F64_mul(v521, v558), l2)
	v649 = v520
	v650 = base.F64_div(base.F64_mul(v520, v521), l2)
	goto L98
L133:
	;
	v565 = v561
	goto L135
L134:
	;
	v565 = v562
	goto L135
L135:
	;
	v566 = base.F64_convert_i32_u(v565)
	v567 = base.F64_add(v566, v566)
	v568 = float64(1)
	v570 = *(*int32)(unsafe.Add(mBase, _c_F_cost_index[0]))
	v573 = *(*float64)(unsafe.Add(mBase, uint32(l1)+304))
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
	v576 = base.F64_add(v573, base.F64_convert_i32_u(v574))
	if base.F64_gt(v576, v568) != 0 {
		goto L137
	} else {
		goto L138
	}
L136:
	;
	v609 = *(*float64)(unsafe.Add(mBase, uint32(v27)+32))
	v612 = base.F64_ceil(base.F64_mul(v609, base.F64_convert_i32_u(v562)))
	v614 = base.B2i32(v29 != int32(346))
	if v614 == int32(0) {
		goto L150
	} else {
		goto L151
	}
L137:
	;
	v580 = v576
	goto L139
L138:
	;
	v580 = v568
	goto L139
L139:
	;
	v581 = base.F64_div(base.F64_mul(v566, base.F64_convert_i32_s(v570)), v580)
	if base.F64_le(v581, float64(1)) != 0 {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v585 = v568
	goto L142
L141:
	;
	v585 = base.F64_ceil(v581)
	goto L142
L142:
	;
	if base.F64_le(v566, v585) != 0 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v589 = base.F64_div(base.F64_mul(v453, v567), base.F64_add(v567, v453))
	if base.F64_ge(v589, v566) != 0 {
		v607 = v566
		goto L136
	} else {
		goto L146
	}
L144:
	;
	goto L145
L145:
	;
	v594 = base.F64_div(base.F64_mul(v567, v585), base.F64_sub(v567, v585))
	if base.F64_ge(v594, v453) != 0 {
		goto L147
	} else {
		goto L148
	}
L146:
	;
	v607 = base.F64_ceil(v589)
	goto L136
L147:
	;
	v604 = base.F64_div(base.F64_mul(v453, v567), base.F64_add(v567, v453))
	goto L149
L148:
	;
	v604 = base.F64_add(v585, base.F64_div(base.F64_mul(base.F64_sub(v566, v585), base.F64_sub(v453, v594)), v566))
	goto L149
L149:
	;
	v607 = base.F64_ceil(v604)
	goto L136
L150:
	;
	v618 = *(*float64)(unsafe.Add(mBase, uint32(v31)+136))
	v622 = base.F64_ceil(base.F64_mul(v607, base.F64_sub(float64(1), v618)))
	goto L152
L151:
	;
	v622 = v607
	goto L152
L152:
	;
	v623 = *(*float64)(unsafe.Add(mBase, uint32(v27)+8))
	v624 = base.F64_mul(v622, v623)
	if v614 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v628 = *(*float64)(unsafe.Add(mBase, uint32(v31)+136))
	v632 = base.F64_ceil(base.F64_mul(v612, base.F64_sub(float64(1), v628)))
	goto L155
L154:
	;
	v632 = v612
	goto L155
L155:
	;
	if base.F64_gt(v632, float64(0)) == int32(0) {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v647 = float64(0)
	v649 = v622
	v650 = v624
	goto L98
L157:
	;
	goto L158
L158:
	;
	if base.F64_gt(v632, float64(1)) == int32(0) {
		v647 = v623
		v649 = v622
		v650 = v624
		goto L98
	} else {
		goto L159
	}
L159:
	;
	v644 = *(*float64)(unsafe.Add(mBase, uint32(v27)+16))
	v647 = base.F64_add(base.F64_mul(base.F64_add(v632, float64(-1)), v644), v623)
	v649 = v622
	v650 = v624
	goto L98
L160:
	;
	m.G0 = v27 + int32(80)
	return
L161:
	;
	if v29 == int32(346) {
		goto L164
	} else {
		goto L165
	}
L162:
	;
	goto L163
L163:
	;
	v772 = float64(0)
	v777 = *(*float64)(unsafe.Add(mBase, uint32(v27)+24))
	v778 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v27)+64)) = v778
	*(*int32)(unsafe.Add(mBase, uint32(v27)+56)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v27)+72)) = v778
	if v403 == int32(0) {
		v840 = v772
		v859 = v772
		goto L208
	} else {
		goto L209
	}
L164:
	;
	v660 = float64(-1)
	goto L166
L165:
	;
	v660 = v649
	goto L166
L166:
	;
	v661 = *(*float64)(unsafe.Add(mBase, uint32(v27)))
	v663 = *(*int32)(unsafe.Add(mBase, _c_F_cost_index[1]))
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v31)+156))
	if v666 != int32(-1) {
		v755 = v666
		goto L169
	} else {
		goto L170
	}
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v765
	if v765 <= int32(0) {
		goto L160
	} else {
		goto L207
	}
L168:
	;
	goto L167
L169:
	;
	if v755 < v663 {
		goto L204
	} else {
		goto L205
	}
L170:
	;
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v669 != 0 {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v686 = int32(0)
	if base.F64_ge(v660, float64(0)) == v686 {
		v718 = v686
		goto L179
	} else {
		goto L180
	}
L172:
	;
	if base.F64_ge(v660, float64(0)) != 0 {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v674 = *(*int32)(unsafe.Add(mBase, _c_F_cost_index[2]))
	if base.F64_lt(v660, base.F64_convert_i32_s(v674)) != 0 {
		v765 = int32(0)
		goto L168
	} else {
		goto L176
	}
L174:
	;
	goto L175
L175:
	;
	if base.F64_ge(v661, float64(0)) == int32(0) {
		goto L171
	} else {
		goto L177
	}
L176:
	;
	goto L175
L177:
	;
	v683 = *(*int32)(unsafe.Add(mBase, _c_F_cost_index[3]))
	if base.F64_lt(v661, base.F64_convert_i32_s(v683)) != 0 {
		v765 = int32(0)
		goto L168
	} else {
		goto L178
	}
L178:
	;
	goto L171
L179:
	;
	if base.F64_ge(v661, float64(0)) == int32(0) {
		v755 = v718
		goto L169
	} else {
		goto L188
	}
L180:
	;
	v691 = int32(1)
	v693 = *(*int32)(unsafe.Add(mBase, _c_F_cost_index[2]))
	if v693 <= v691 {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v696 = v691
	goto L183
L182:
	;
	v696 = v693
	goto L183
L183:
	;
	v698 = v696
	v702 = int32(1)
	goto L184
L184:
	;
	v705 = v698 * int32(3)
	if base.F64_ge(v660, base.F64_convert_i32_u(v705)) == int32(0) {
		v718 = v702
		goto L179
	} else {
		goto L186
	}
L185:
	;
	v718 = v711
	goto L179
L186:
	;
	v711 = v702 + int32(1)
	if v705 < int32(715827883) {
		v698 = v705
		v702 = v711
		goto L184
	} else {
		goto L187
	}
L187:
	;
	goto L185
L188:
	;
	v724 = int32(1)
	v726 = *(*int32)(unsafe.Add(mBase, _c_F_cost_index[3]))
	if v726 <= v724 {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v729 = v724
	goto L191
L190:
	;
	v729 = v726
	goto L191
L191:
	;
	v731 = v729
	v736 = int32(1)
	goto L192
L192:
	;
	v738 = v731 * int32(3)
	if base.F64_le(base.F64_convert_i32_u(v738), v661) != 0 {
		goto L194
	} else {
		goto L195
	}
L193:
	;
	if v718 < v745 {
		goto L198
	} else {
		goto L199
	}
L194:
	;
	v742 = v736 + int32(1)
	if v738 < int32(715827883) {
		v731 = v738
		v736 = v742
		goto L192
	} else {
		goto L197
	}
L195:
	;
	v745 = v736
	goto L196
L196:
	;
	goto L193
L197:
	;
	v745 = v742
	goto L196
L198:
	;
	v747 = v718
	goto L200
L199:
	;
	v747 = v745
	goto L200
L200:
	;
	if int32(0) < v718 {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v750 = v747
	goto L203
L202:
	;
	v750 = v745
	goto L203
L203:
	;
	v755 = v750
	goto L169
L204:
	;
	v758 = v755
	goto L206
L205:
	;
	v758 = v663
	goto L206
L206:
	;
	v765 = v758
	goto L168
L207:
	;
	v769 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)) = uint8(v769)
	goto L163
L208:
	;
	v860 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v861 = *(*float64)(unsafe.Add(mBase, uint32(v860)+24))
	v862 = *(*float64)(unsafe.Add(mBase, uint32(l0)+32))
	v865 = *(*float64)(unsafe.Add(mBase, _c_F_cost_index[4]))
	v870 = base.F64_add(base.F64_mul(v861, v862), base.F64_add(base.F64_mul(base.F64_add(v840, v865), v453), float64(0)))
	v872 = *(*float64)(unsafe.Add(mBase, uint32(v860)+16))
	v873 = base.F64_add(base.F64_add(base.F64_add(v454, v772), v859), v872)
	v874 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if int32(0) < v874 {
		goto L215
	} else {
		goto L216
	}
L209:
	;
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v403)+4))
	if v792 <= int32(0) {
		v840 = v772
		v859 = float64(0)
		goto L208
	} else {
		goto L210
	}
L210:
	;
	v810 = int32(0)
	goto L211
L211:
	;
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v403)+12))
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v820+v810<<(uint(int32(2))%32))))
	v827 = F_cost_qual_eval_walker(m, v824, v27+int32(56))
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		goto L30
	} else {
		goto L213
	}
L212:
	;
	v833 = *(*float64)(unsafe.Add(mBase, uint32(v27)+72))
	v834 = *(*float64)(unsafe.Add(mBase, uint32(v27)+64))
	v840 = v833
	v859 = v834
	goto L208
L213:
	;
	v830 = v810 + int32(1)
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v403)+4))
	if v830 < v831 {
		v810 = v830
		goto L211
	} else {
		goto L214
	}
L214:
	;
	goto L212
L215:
	;
	v877 = base.F64_convert_i32_u(v874)
	v879 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_cost_index[5])))
	if v879 == int32(1) {
		goto L219
	} else {
		goto L220
	}
L216:
	;
	v912 = v870
	goto L217
L217:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = v873
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = base.F64_add(v873, base.F64_add(base.F64_add(base.F64_add(base.F64_sub(v434, v454), v772), base.F64_add(base.F64_mul(base.F64_mul(v777, v777), base.F64_sub(v647, v650)), v650)), v912))
	goto L160
L218:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v907
	v912 = base.F64_div(v870, v892)
	goto L217
L219:
	;
	v885 = base.F64_add(base.F64_mul(v877, float64(-0.3)), float64(1))
	if base.F64_gt(v885, float64(0)) != 0 {
		goto L222
	} else {
		goto L223
	}
L220:
	;
	v892 = v877
	goto L221
L221:
	;
	v893 = float64(1e+100)
	v894 = base.F64_div(v862, v892)
	if base.F64_gt(v894, v893)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v894)&int64(9223372036854775807))) != 0 {
		v907 = v893
		goto L218
	} else {
		goto L225
	}
L222:
	;
	v889 = v885
	goto L224
L223:
	;
	v889 = math.Float64frombits(uint64(0x8000000000000000))
	goto L224
L224:
	;
	v892 = base.F64_add(v889, v877)
	goto L221
L225:
	;
	v903 = float64(1)
	if base.F64_le(v894, v903) != 0 {
		v907 = v903
		goto L218
	} else {
		goto L226
	}
L226:
	;
	v907 = base.F64_nearest(v894)
	goto L218
}
func F_index_am_handler_in(m *base.Module, l0 int32) int64 {
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14235(m, l0, int32(_a_F_index_am_handler_in_0), int32(371), int32(_a_F_index_am_handler_in_1), int32(_a_F_index_am_handler_in_2), int32(_a_F_index_am_handler_in_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
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
	v80 = m.ExcPending
	if v80 != 0 {
		goto L4
	} else {
		goto L20
	}
L2:
	;
	F_list_free(m, v12)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L4
	} else {
		goto L19
	}
L3:
	;
	v67 = int32(0)
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
	v36 = F_SearchSysCache1(m, int32(57), base.I64_extend_i32_u(v34))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L4
	} else {
		goto L10
	}
L9:
	;
	goto L3
L10:
	;
	if v36 == int32(0) {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+22)))
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40+v41)+131)))
	F_ReleaseCatCache(m, v36)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	if v43 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v49 = F_get_partition_parent(m, v34, int32(0))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L4
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v53 = v22 + int32(1)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v53 < v54 {
		v22 = v53
		goto L8
	} else {
		goto L18
	}
L16:
	;
	if v49 == l1 {
		v67 = v34
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
	return v67
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v34
	F_errmsg_internal(m, int32(_a_F_index_get_partition_0), v10)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(_a_F_index_get_partition_1), int32(190), int32(_a_F_index_get_partition_2))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
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
				F_errfinish(m, int32(_a_F_index_markpos_2), int32(427), int32(_a_F_index_markpos_3))
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
func F_index_parallelscan_estimate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_index_parallelscan_estimate[0]))
	if v13 != v11 {
		v16 = *(*int32)(unsafe.Add(mBase, _c_F_index_parallelscan_estimate[1]))
		v17 = F_list_member_ptr(m, v16, v11)
		mBase = m.M
		v19 = v17
	} else {
		v19 = int32(1)
	}
	if v19 == int32(0) {
		v23 = F_EstimateSnapshotSpace(m, l3)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			v27 = F_add_size(m, int32(28), v23)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				v32 = (v27 + int32(7)) & int32(-8)
				v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+120))
				if v34 != 0 {
					v35 = m.T0[v34].(func(*base.Module, int32, int32, int32) int32)(m, l0, l1, l2)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						v37 = F_add_size(m, v32, v35)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							v39 = v37
							m.G0 = v9 + int32(16)
							return v39
						}
					}
				} else {
					v39 = v32
					m.G0 = v9 + int32(16)
					return v39
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v47 = m.ExcPending
		if v47 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return int32(0)
			} else {
				v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = v51 + int32(4)
				F_errmsg(m, int32(_a_F_index_parallelscan_estimate_0), v9)
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_index_parallelscan_estimate_1), int32(475), int32(_a_F_index_parallelscan_estimate_2))
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
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
func F_index_reloptions(m *base.Module, l0 int32, l1 int64) {
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	if base.I32_wrap_i64(l1) == int32(0) {
		return
	} else {
		v7 = m.T0[l0].(func(*base.Module, int64, int32) int32)(m, l1, int32(1))
		v8 = m.ExcPending
		if v8 != 0 {
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
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int64
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int64
	_ = v34
	var v35 int32
	_ = v35
	var v36 int64
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 float64
	_ = v43
	var v50 int64
	_ = v50
	var v52 int32
	_ = v52
	var v54 int64
	_ = v54
	var v55 int32
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
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	v4 = l3
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+84)) = uint8(v4)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if int32(0) < v9 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v16 = int32(0)
	goto L4
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1+v16<<(uint(int32(2))%32))))
	switch v23 - int32(700) {
	case 0:
		goto L10
	case 1:
		goto L11
	default:
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	*(*uint8)(unsafe.Add(mBase, uint32(v80+v16))) = uint8(v77)
	v84 = v16 + int32(1)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v84 < v85 {
		v16 = v84
		goto L4
	} else {
		goto L21
	}
L7:
	;
	v60 = int32(1)
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+84)))
	if v61 != v60 {
		v77 = v60
		goto L6
	} else {
		goto L16
	}
L8:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	*(*int64)(unsafe.Add(mBase, uint32(v55+v16<<(uint(int32(3))%32)))) = v54
	v77 = v52
	goto L6
L9:
	;
	v52 = int32(0)
	v54 = v50
	goto L8
L10:
	;
	v35 = int32(1)
	v36 = int64(0)
	if l2 == int32(0) {
		v52 = v35
		v54 = v36
		goto L8
	} else {
		goto L14
	}
L11:
	;
	v26 = int32(1)
	v27 = int64(0)
	if l2 == int32(0) {
		v52 = v26
		v54 = v27
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v32 = l2 + v16<<(uint(int32(4))%32)
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+8)))
	if v33 != 0 {
		v52 = v26
		v54 = v27
		goto L8
	} else {
		goto L13
	}
L13:
	;
	v34 = *(*int64)(unsafe.Add(mBase, uint32(v32)))
	v50 = v34
	goto L9
L14:
	;
	v41 = l2 + v16<<(uint(int32(4))%32)
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+8)))
	if v42 != 0 {
		v52 = v35
		v54 = v36
		goto L8
	} else {
		goto L15
	}
L15:
	;
	v43 = *(*float64)(unsafe.Add(mBase, uint32(v41)))
	v50 = base.I64_extend_i32_s(base.I32_reinterpret_f32(base.F32_demote_f64(v43)))
	goto L9
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	return
L18:
	;
	F_errmsg_internal(m, int32(_a_F_index_store_float8_orderby_distances_0), int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(_a_F_index_store_float8_orderby_distances_1), int32(1003), int32(_a_F_index_store_float8_orderby_distances_2))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L21:
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
	var v95 int32
	_ = v95
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
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 float32
	_ = v118
	var v119 float32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
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
	F_ScanKeyInit(m, v89, int32(1), int32(3), int32(184), base.I64_extend_i32_u(v29))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L13
	} else {
		goto L18
	}
L18:
	;
	F_systable_inplace_update_begin(m, v86, int32(2662), v89, v13+int32(12), v13+int32(8))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L13
	} else {
		goto L19
	}
L19:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v103 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+16))
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+22)))
	v106 = v104 + v105
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+116)))
	if v2 != v107 {
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
	v167 = m.ExcPending
	if v167 != 0 {
		goto L13
	} else {
		goto L47
	}
L23:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v106)+116)) = uint8(v2)
	goto L25
L24:
	;
	goto L25
L25:
	;
	v110 = base.B2i32(v2 != v107)
	if v83 == int32(0) {
		v134 = v110
		goto L29
	} else {
		goto L30
	}
L26:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	F_pfree(m, v155)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L13
	} else {
		goto L45
	}
L27:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	F_systable_inplace_update_cancel(m, v146)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L13
	} else {
		goto L43
	}
L28:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	F_systable_inplace_update_finish(m, v142, v143)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L13
	} else {
		goto L42
	}
L29:
	;
	if v134 == int32(0) {
		goto L27
	} else {
		goto L41
	}
L30:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v106)+96))
	if v82 != v113 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v106)+96)) = v82
	v117 = int32(1)
	goto L33
L32:
	;
	v117 = v110
	goto L33
L33:
	;
	v118 = base.F32_demote_f64(v28)
	v119 = *(*float32)(unsafe.Add(mBase, uint32(v106)+100))
	if base.F32_ne(v118, v119) != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v106)+100)) = v118
	v123 = int32(1)
	goto L36
L35:
	;
	v123 = v117
	goto L36
L36:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+76))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v106)+104))
	if v124 != v125 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v106)+104)) = v124
	v129 = int32(1)
	goto L39
L38:
	;
	v129 = v123
	goto L39
L39:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v13)+72))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v106)+108))
	if v130 == v131 {
		v134 = v129
		goto L29
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v106)+108)) = v130
	goto L28
L41:
	;
	goto L28
L42:
	;
	goto L26
L43:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	F_CacheInvalidateRelcacheByTuple(m, v149)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
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
	v160 = m.ExcPending
	if v160 != 0 {
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
	v171 = m.ExcPending
	if v171 != 0 {
		goto L13
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_index_update_stats_1), int32(3037), int32(_a_F_index_update_stats_2))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
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
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	v7 = F_strip_noop_phvs(m, l0)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L6
	} else {
		goto L48
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L6
	} else {
		goto L45
	}
L3:
	;
	return int32(0)
L4:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l2)+84))
	if v57 != 0 {
		goto L21
	} else {
		goto L22
	}
L5:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33+l1<<(uint(int32(2))%32))))
	if v37 == int32(0) {
		v51 = v11
		v56 = v33
		goto L4
	} else {
		goto L16
	}
L6:
	;
	return int32(0)
L7:
	;
	if v7 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v11 = v7
	goto L11
L9:
	;
	goto L10
L10:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v27+l1<<(uint(int32(2))%32))))
	if v31 != 0 {
		goto L3
	} else {
		goto L15
	}
L11:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	if v17 != int32(27) {
		goto L5
	} else {
		goto L13
	}
L12:
	;
	goto L10
L13:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v20 != 0 {
		v11 = v20
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v51 = int32(0)
	v56 = v27
	goto L4
L16:
	;
	if v17 != int32(6) {
		goto L3
	} else {
		goto L17
	}
L17:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+76))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v43 != v44 {
		goto L3
	} else {
		goto L18
	}
L18:
	;
	v46 = int32(*(*int16)(unsafe.Add(mBase, uint32(v11)+8)))
	if v37 != v46 {
		goto L3
	} else {
		goto L19
	}
L19:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	if v48 != 0 {
		goto L3
	} else {
		goto L20
	}
L20:
	;
	return int32(1)
L21:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	v60 = v58
	goto L23
L22:
	;
	v60 = int32(0)
	goto L23
L23:
	;
	if int32(0) < l1 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v66 = v60
	v67 = int32(0)
	goto L27
L25:
	;
	v94 = v60
	goto L26
L26:
	;
	if v94 == int32(0) {
		goto L1
	} else {
		goto L37
	}
L27:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v56+v67<<(uint(int32(2))%32))))
	if v73 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v94 = v88
	goto L26
L29:
	;
	if v66 == int32(0) {
		goto L2
	} else {
		goto L32
	}
L30:
	;
	v88 = v66
	goto L31
L31:
	;
	v90 = v67 + int32(1)
	if v90 != l1 {
		v66 = v88
		v67 = v90
		goto L27
	} else {
		goto L36
	}
L32:
	;
	v79 = v66 + int32(4)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if base.Ui32(v79) < base.Ui32(v81+v82<<(uint(int32(2))%32)) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v87 = v79
	goto L35
L34:
	;
	v87 = int32(0)
	goto L35
L35:
	;
	v88 = v87
	goto L31
L36:
	;
	goto L28
L37:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	if v100 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v109 = F_equal(m, v108, v51)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L6
	} else {
		goto L43
	}
L39:
	;
	v108 = int32(0)
	goto L38
L40:
	;
	goto L41
L41:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	if v104 != int32(27) {
		v108 = v100
		goto L38
	} else {
		goto L42
	}
L42:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
	v108 = v107
	goto L38
L43:
	;
	if v109 == int32(0) {
		goto L3
	} else {
		goto L44
	}
L44:
	;
	return int32(1)
L45:
	;
	F_errmsg_internal(m, int32(_a_F_match_index_to_operand_0), int32(0))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L6
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_match_index_to_operand_1), int32(_a_F_match_index_to_operand_2), int32(_a_F_match_index_to_operand_3))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L6
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L48:
	;
	F_errmsg_internal(m, int32(_a_F_match_index_to_operand_0), int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L6
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(_a_F_match_index_to_operand_1), int32(_a_F_match_index_to_operand_4), int32(_a_F_match_index_to_operand_3))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L6
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
