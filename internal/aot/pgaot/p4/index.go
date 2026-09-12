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
		v12 = F_ExecScan(m, l0, int32(722), int32(723))
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
			v12 = F_ExecScan(m, l0, int32(722), int32(723))
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
				v12 = F_ExecScan(m, l0, int32(722), int32(723))
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
			v15 = int32(726)
		} else {
			v15 = int32(727)
		}
		v17 = F_ExecScan(m, l0, v15, int32(728))
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
				v15 = int32(726)
			} else {
				v15 = int32(727)
			}
			v17 = F_ExecScan(m, l0, v15, int32(728))
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
					v15 = int32(726)
				} else {
					v15 = int32(727)
				}
				v17 = F_ExecScan(m, l0, v15, int32(728))
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
					F_errmsg_internal(m, int32(101795), v6)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(476056), int32(43), int32(356759))
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
				F_errmsg_internal(m, int32(101795), v6)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(476056), int32(43), int32(356759))
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
					F_errmsg_internal(m, int32(37833), v8)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(471121), int32(3594), int32(252400))
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
				F_errmsg_internal(m, int32(43981), v7)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(476084), int32(613), int32(272145))
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
	F_RecordPageWithFreeSpace(m, l0, l1, int32(8191))
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
	if l2&int32(65535) != v69 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v74 = int32(65535)
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
	v97 = *(*int32)(unsafe.Add(mBase, _consts[380]))
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
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
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
	var v92 int32
	_ = v92
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
	var v217 int32
	_ = v217
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
		v92 = v53
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
	v58 = int32(0)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	if v59 <= v58 {
		v92 = v53
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v64 = v58
	v67 = v53
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
	v92 = v82
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
	v80 = F_lappend(m, v67, v77)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L8
	} else {
		goto L19
	}
L17:
	;
	v82 = v67
	goto L18
L18:
	;
	v84 = v64 + int32(1)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	if v84 < v85 {
		v64 = v84
		v67 = v82
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
	v119 = v92
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
	v117 = F_list_concat(m, v92, v115)
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
	v217 = v210
	goto L65
L65:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v207)+12))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v225+v217<<(uint(int32(2))%32))))
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
	v254 = v217 + int32(1)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v207)+4))
	if v254 < v255 {
		v217 = v254
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
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
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
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
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
	var v247 int32
	_ = v247
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
	var v278 int32
	_ = v278
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
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
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
	var v339 int32
	_ = v339
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v400 int32
	_ = v400
	var v409 int32
	_ = v409
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v424 float64
	_ = v424
	var v426 float64
	_ = v426
	var v428 float64
	_ = v428
	var v429 float64
	_ = v429
	var v430 float64
	_ = v430
	var v438 float64
	_ = v438
	var v442 float64
	_ = v442
	var v443 float64
	_ = v443
	var v444 int32
	_ = v444
	var v450 int32
	_ = v450
	var v453 float64
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v459 float64
	_ = v459
	var v460 float64
	_ = v460
	var v461 float64
	_ = v461
	var v463 int32
	_ = v463
	var v466 float64
	_ = v466
	var v467 int32
	_ = v467
	var v469 float64
	_ = v469
	var v473 float64
	_ = v473
	var v474 float64
	_ = v474
	var v478 float64
	_ = v478
	var v479 int32
	_ = v479
	var v482 float64
	_ = v482
	var v487 float64
	_ = v487
	var v497 float64
	_ = v497
	var v501 float64
	_ = v501
	var v505 float64
	_ = v505
	var v509 float64
	_ = v509
	var v510 float64
	_ = v510
	var v511 float64
	_ = v511
	var v515 float64
	_ = v515
	var v518 float64
	_ = v518
	var v523 float64
	_ = v523
	var v533 float64
	_ = v533
	var v535 float64
	_ = v535
	var v543 float64
	_ = v543
	var v547 float64
	_ = v547
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v555 float64
	_ = v555
	var v556 float64
	_ = v556
	var v557 float64
	_ = v557
	var v559 int32
	_ = v559
	var v562 float64
	_ = v562
	var v563 int32
	_ = v563
	var v565 float64
	_ = v565
	var v569 float64
	_ = v569
	var v570 float64
	_ = v570
	var v574 float64
	_ = v574
	var v578 float64
	_ = v578
	var v583 float64
	_ = v583
	var v593 float64
	_ = v593
	var v596 float64
	_ = v596
	var v598 float64
	_ = v598
	var v601 float64
	_ = v601
	var v603 int32
	_ = v603
	var v607 float64
	_ = v607
	var v611 float64
	_ = v611
	var v612 float64
	_ = v612
	var v613 float64
	_ = v613
	var v617 float64
	_ = v617
	var v621 float64
	_ = v621
	var v633 float64
	_ = v633
	var v636 float64
	_ = v636
	var v637 float64
	_ = v637
	var v639 float64
	_ = v639
	var v649 float64
	_ = v649
	var v650 float64
	_ = v650
	var v652 int32
	_ = v652
	var v655 int32
	_ = v655
	var v658 int32
	_ = v658
	var v663 int32
	_ = v663
	var v672 int32
	_ = v672
	var v675 int32
	_ = v675
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v700 int32
	_ = v700
	var v707 int32
	_ = v707
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v731 int32
	_ = v731
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v739 int32
	_ = v739
	var v744 int32
	_ = v744
	var v747 int32
	_ = v747
	var v754 int32
	_ = v754
	var v758 int32
	_ = v758
	var v761 float64
	_ = v761
	var v766 float64
	_ = v766
	var v767 int64
	_ = v767
	var v781 int32
	_ = v781
	var v799 int32
	_ = v799
	var v809 int32
	_ = v809
	var v813 int32
	_ = v813
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v822 float64
	_ = v822
	var v823 float64
	_ = v823
	var v828 float64
	_ = v828
	var v848 float64
	_ = v848
	var v850 int32
	_ = v850
	var v851 float64
	_ = v851
	var v852 float64
	_ = v852
	var v855 float64
	_ = v855
	var v860 float64
	_ = v860
	var v862 float64
	_ = v862
	var v863 float64
	_ = v863
	var v864 int32
	_ = v864
	var v867 float64
	_ = v867
	var v869 int32
	_ = v869
	var v875 float64
	_ = v875
	var v879 float64
	_ = v879
	var v881 float64
	_ = v881
	var v883 float64
	_ = v883
	var v884 float64
	_ = v884
	var v892 float64
	_ = v892
	var v896 float64
	_ = v896
	var v899 float64
	_ = v899
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
	v409 = int32(*(*uint8)(unsafe.Add(mBase, _consts[623])))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v409 ^ int32(1)
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v30)+116))
	m.T0[v421].(func(*base.Module, int32, int32, float64, int32, int32, int32, int32, int32))(m, l1, l0, l2, v27+int32(48), v27+int32(40), v27+int32(32), v27+int32(24), v27)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L34
	} else {
		goto L99
	}
L2:
	;
	v33 = *(*float64)(unsafe.Add(mBase, uint32(v32)+8))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v33
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v30)+96))
	if v36 == int32(0) {
		goto L6
	} else {
		goto L7
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
		goto L69
	} else {
		goto L70
	}
L5:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v152)+16))
	if v159 != 0 {
		goto L37
	} else {
		goto L38
	}
L6:
	;
	v152 = v32
	v154 = v35
	v156 = v15
	goto L5
L7:
	;
	goto L8
L8:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	if v39 <= int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v152 = v32
	v154 = v35
	v156 = v15
	goto L5
L10:
	;
	goto L11
L11:
	;
	v56 = v15
	v63 = v15
	goto L12
L12:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v66+v56<<(uint(int32(2))%32))))
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+10)))
	if v71 != 0 {
		v128 = v63
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v152 = v134
	v154 = v133
	v156 = v128
	goto L5
L14:
	;
	v130 = v56 + int32(1)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	if v130 < v131 {
		v56 = v130
		v63 = v128
		goto L12
	} else {
		goto L36
	}
L15:
	;
	if v35 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	if v122 != 0 {
		v128 = v63
		goto L14
	} else {
		goto L33
	}
L17:
	;
	goto L16
L18:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v77 <= int32(0) {
		v122 = int32(0)
		goto L17
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v122 = int32(0)
	goto L17
L21:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v70)+60))
	v81 = int32(0)
	if v81 < v77 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v84 = v77
	goto L24
L23:
	;
	v84 = v81
	goto L24
L24:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
	v88 = int32(0)
	goto L25
L25:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v85+v88<<(uint(int32(2))%32))))
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+12)))
	if v98 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	goto L20
L27:
	;
	v109 = v88 + int32(1)
	if v109 != v84 {
		v88 = v109
		goto L25
	} else {
		goto L32
	}
L28:
	;
	v99 = int32(1)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
	if v70 == v100 {
		v122 = v99
		goto L17
	} else {
		goto L29
	}
L29:
	;
	if v80 == int32(0) {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v100)+60))
	if v104 == v80 {
		v122 = v99
		goto L17
	} else {
		goto L31
	}
L31:
	;
	goto L27
L32:
	;
	goto L26
L33:
	;
	v126 = F_lappend(m, v63, v70)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	return
L35:
	;
	v128 = v126
	goto L14
L36:
	;
	goto L13
L37:
	;
	v160 = int32(0)
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v159)+4))
	if v161 <= v160 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	v278 = v15
	goto L39
L39:
	;
	v282 = F_list_concat(m, v156, v278)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L34
	} else {
		goto L68
	}
L40:
	;
	v165 = F_list_concat(m, v156, int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L34
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v181 = v160
	v187 = v15
	goto L44
L43:
	;
	v400 = v165
	goto L1
L44:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v159)+12))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v191+v181<<(uint(int32(2))%32))))
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195)+10)))
	if v196 != 0 {
		v253 = v187
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v278 = v253
	goto L39
L46:
	;
	v255 = v181 + int32(1)
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v159)+4))
	if v255 < v256 {
		v181 = v255
		v187 = v253
		goto L44
	} else {
		goto L67
	}
L47:
	;
	if v154 != 0 {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	if v247 != 0 {
		v253 = v187
		goto L46
	} else {
		goto L65
	}
L49:
	;
	goto L48
L50:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v154)+4))
	if v202 <= int32(0) {
		v247 = int32(0)
		goto L49
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v247 = int32(0)
	goto L49
L53:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v195)+60))
	v206 = int32(0)
	if v206 < v202 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v209 = v202
	goto L56
L55:
	;
	v209 = v206
	goto L56
L56:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v154)+12))
	v213 = int32(0)
	goto L57
L57:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v210+v213<<(uint(int32(2))%32))))
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+12)))
	if v223 != 0 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	goto L52
L59:
	;
	v234 = v213 + int32(1)
	if v234 != v209 {
		v213 = v234
		goto L57
	} else {
		goto L64
	}
L60:
	;
	v224 = int32(1)
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
	if v195 == v225 {
		v247 = v224
		goto L49
	} else {
		goto L61
	}
L61:
	;
	if v205 == int32(0) {
		goto L59
	} else {
		goto L62
	}
L62:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v225)+60))
	if v229 == v205 {
		v247 = v224
		goto L49
	} else {
		goto L63
	}
L63:
	;
	goto L59
L64:
	;
	goto L58
L65:
	;
	v251 = F_lappend(m, v187, v195)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L34
	} else {
		goto L66
	}
L66:
	;
	v253 = v251
	goto L46
L67:
	;
	goto L45
L68:
	;
	v400 = v282
	goto L1
L69:
	;
	v400 = v15
	goto L1
L70:
	;
	goto L71
L71:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v286)+4))
	if v289 <= int32(0) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v400 = v15
	goto L1
L73:
	;
	goto L74
L74:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v307 = v15
	v309 = v15
	goto L75
L75:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v286)+12))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v317+v307<<(uint(int32(2))%32))))
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v321)+10)))
	if v322 != 0 {
		v379 = v309
		goto L77
	} else {
		goto L78
	}
L76:
	;
	v400 = v379
	goto L1
L77:
	;
	v381 = v307 + int32(1)
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v286)+4))
	if v381 < v382 {
		v307 = v381
		v309 = v379
		goto L75
	} else {
		goto L98
	}
L78:
	;
	if v292 != 0 {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	if v373 != 0 {
		v379 = v309
		goto L77
	} else {
		goto L96
	}
L80:
	;
	goto L79
L81:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v292)+4))
	if v328 <= int32(0) {
		v373 = int32(0)
		goto L80
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v373 = int32(0)
	goto L80
L84:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v321)+60))
	v332 = int32(0)
	if v332 < v328 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v335 = v328
	goto L87
L86:
	;
	v335 = v332
	goto L87
L87:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v292)+12))
	v339 = int32(0)
	goto L88
L88:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v336+v339<<(uint(int32(2))%32))))
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v348)+12)))
	if v349 != 0 {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	goto L83
L90:
	;
	v360 = v339 + int32(1)
	if v360 != v335 {
		v339 = v360
		goto L88
	} else {
		goto L95
	}
L91:
	;
	v350 = int32(1)
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v348)+4))
	if v321 == v351 {
		v373 = v350
		goto L80
	} else {
		goto L92
	}
L92:
	;
	if v331 == int32(0) {
		goto L90
	} else {
		goto L93
	}
L93:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v351)+60))
	if v355 == v331 {
		v373 = v350
		goto L80
	} else {
		goto L94
	}
L94:
	;
	goto L90
L95:
	;
	goto L89
L96:
	;
	v377 = F_lappend(m, v309, v321)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L34
	} else {
		goto L97
	}
L97:
	;
	v379 = v377
	goto L77
L98:
	;
	goto L76
L99:
	;
	v424 = *(*float64)(unsafe.Add(mBase, uint32(v27)+40))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+96)) = v424
	v426 = *(*float64)(unsafe.Add(mBase, uint32(v27)+32))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+104)) = v426
	v428 = float64(1e+100)
	v429 = *(*float64)(unsafe.Add(mBase, uint32(v31)+120))
	v430 = base.F64_mul(v426, v429)
	if base.F64_gt(v430, v428) != 0 {
		v442 = v428
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v443 = *(*float64)(unsafe.Add(mBase, uint32(v27)+48))
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v31)+72))
	F_get_tablespace_page_costs(m, v444, v27+int32(8), v27+int32(16))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L34
	} else {
		goto L104
	}
L101:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v430)&int64(9223372036854775807)) {
		v442 = v428
		goto L100
	} else {
		goto L102
	}
L102:
	;
	v438 = float64(1)
	if base.F64_le(v430, v438) != 0 {
		v442 = v438
		goto L100
	} else {
		goto L103
	}
L103:
	;
	v442 = base.F64_nearest(v430)
	goto L100
L104:
	;
	if base.F64_gt(l2, float64(1)) != 0 {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	if l3 != 0 {
		goto L168
	} else {
		goto L169
	}
L106:
	;
	v453 = base.F64_mul(l2, v442)
	v454 = int32(1)
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v31)+116))
	if base.Ui32(v455) <= base.Ui32(v454) {
		goto L109
	} else {
		goto L110
	}
L107:
	;
	goto L108
L108:
	;
	v550 = int32(1)
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v31)+116))
	if base.Ui32(v551) <= base.Ui32(v550) {
		goto L140
	} else {
		goto L141
	}
L109:
	;
	v458 = v454
	goto L111
L110:
	;
	v458 = v455
	goto L111
L111:
	;
	v459 = base.F64_convert_i32_u(v458)
	v460 = base.F64_add(v459, v459)
	v461 = float64(1)
	v463 = *(*int32)(unsafe.Add(mBase, _consts[79]))
	v466 = *(*float64)(unsafe.Add(mBase, uint32(l1)+288))
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
	v469 = base.F64_add(v466, base.F64_convert_i32_u(v467))
	if base.F64_gt(v469, v461) != 0 {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	if v29 == int32(342) {
		goto L126
	} else {
		goto L127
	}
L113:
	;
	v473 = v469
	goto L115
L114:
	;
	v473 = v461
	goto L115
L115:
	;
	v474 = base.F64_div(base.F64_mul(v459, base.F64_convert_i32_s(v463)), v473)
	if base.F64_le(v474, float64(1)) != 0 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v478 = v461
	goto L118
L117:
	;
	v478 = base.F64_ceil(v474)
	goto L118
L118:
	;
	v479 = base.F64_ge(v478, v459)
	if v479 != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v482 = base.F64_div(base.F64_mul(v453, v460), base.F64_add(v460, v453))
	if base.F64_ge(v482, v459) != 0 {
		v501 = v459
		goto L112
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v487 = base.F64_div(base.F64_mul(v460, v478), base.F64_sub(v460, v478))
	if base.F64_ge(v487, v453) != 0 {
		goto L123
	} else {
		goto L124
	}
L122:
	;
	v501 = base.F64_ceil(v482)
	goto L112
L123:
	;
	v497 = base.F64_div(base.F64_mul(v453, v460), base.F64_add(v460, v453))
	goto L125
L124:
	;
	v497 = base.F64_add(v478, base.F64_div(base.F64_mul(base.F64_sub(v459, v478), base.F64_sub(v453, v487)), v459))
	goto L125
L125:
	;
	v501 = base.F64_ceil(v497)
	goto L112
L126:
	;
	v505 = *(*float64)(unsafe.Add(mBase, uint32(v31)+128))
	v509 = base.F64_ceil(base.F64_mul(v501, base.F64_sub(float64(1), v505)))
	goto L128
L127:
	;
	v509 = v501
	goto L128
L128:
	;
	v510 = *(*float64)(unsafe.Add(mBase, uint32(v27)+8))
	v511 = *(*float64)(unsafe.Add(mBase, uint32(v27)+32))
	v515 = base.F64_mul(l2, base.F64_ceil(base.F64_mul(v511, base.F64_convert_i32_u(v455))))
	if v479 != 0 {
		goto L130
	} else {
		goto L131
	}
L129:
	;
	if v29 == int32(342) {
		goto L137
	} else {
		goto L138
	}
L130:
	;
	v518 = base.F64_div(base.F64_mul(v460, v515), base.F64_add(v460, v515))
	if base.F64_ge(v518, v459) != 0 {
		v535 = v459
		goto L129
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	v523 = base.F64_div(base.F64_mul(v460, v478), base.F64_sub(v460, v478))
	if base.F64_ge(v523, v515) != 0 {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	v535 = base.F64_ceil(v518)
	goto L129
L134:
	;
	v533 = base.F64_div(base.F64_mul(v460, v515), base.F64_add(v460, v515))
	goto L136
L135:
	;
	v533 = base.F64_add(v478, base.F64_div(base.F64_mul(base.F64_sub(v459, v478), base.F64_sub(v515, v523)), v459))
	goto L136
L136:
	;
	v535 = base.F64_ceil(v533)
	goto L129
L137:
	;
	v543 = *(*float64)(unsafe.Add(mBase, uint32(v31)+128))
	v547 = base.F64_ceil(base.F64_mul(v535, base.F64_sub(float64(1), v543)))
	goto L139
L138:
	;
	v547 = v535
	goto L139
L139:
	;
	v636 = base.F64_div(base.F64_mul(v510, v547), l2)
	v637 = v509
	v639 = base.F64_div(base.F64_mul(v509, v510), l2)
	goto L105
L140:
	;
	v554 = v550
	goto L142
L141:
	;
	v554 = v551
	goto L142
L142:
	;
	v555 = base.F64_convert_i32_u(v554)
	v556 = base.F64_add(v555, v555)
	v557 = float64(1)
	v559 = *(*int32)(unsafe.Add(mBase, _consts[79]))
	v562 = *(*float64)(unsafe.Add(mBase, uint32(l1)+288))
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
	v565 = base.F64_add(v562, base.F64_convert_i32_u(v563))
	if base.F64_gt(v565, v557) != 0 {
		goto L144
	} else {
		goto L145
	}
L143:
	;
	v598 = *(*float64)(unsafe.Add(mBase, uint32(v27)+32))
	v601 = base.F64_ceil(base.F64_mul(v598, base.F64_convert_i32_u(v551)))
	v603 = base.B2i32(v29 != int32(342))
	if v603 == int32(0) {
		goto L157
	} else {
		goto L158
	}
L144:
	;
	v569 = v565
	goto L146
L145:
	;
	v569 = v557
	goto L146
L146:
	;
	v570 = base.F64_div(base.F64_mul(v555, base.F64_convert_i32_s(v559)), v569)
	if base.F64_le(v570, float64(1)) != 0 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v574 = v557
	goto L149
L148:
	;
	v574 = base.F64_ceil(v570)
	goto L149
L149:
	;
	if base.F64_le(v555, v574) != 0 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v578 = base.F64_div(base.F64_mul(v442, v556), base.F64_add(v556, v442))
	if base.F64_ge(v578, v555) != 0 {
		v596 = v555
		goto L143
	} else {
		goto L153
	}
L151:
	;
	goto L152
L152:
	;
	v583 = base.F64_div(base.F64_mul(v556, v574), base.F64_sub(v556, v574))
	if base.F64_ge(v583, v442) != 0 {
		goto L154
	} else {
		goto L155
	}
L153:
	;
	v596 = base.F64_ceil(v578)
	goto L143
L154:
	;
	v593 = base.F64_div(base.F64_mul(v442, v556), base.F64_add(v556, v442))
	goto L156
L155:
	;
	v593 = base.F64_add(v574, base.F64_div(base.F64_mul(base.F64_sub(v555, v574), base.F64_sub(v442, v583)), v555))
	goto L156
L156:
	;
	v596 = base.F64_ceil(v593)
	goto L143
L157:
	;
	v607 = *(*float64)(unsafe.Add(mBase, uint32(v31)+128))
	v611 = base.F64_ceil(base.F64_mul(v596, base.F64_sub(float64(1), v607)))
	goto L159
L158:
	;
	v611 = v596
	goto L159
L159:
	;
	v612 = *(*float64)(unsafe.Add(mBase, uint32(v27)+8))
	v613 = base.F64_mul(v611, v612)
	if v603 == int32(0) {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v617 = *(*float64)(unsafe.Add(mBase, uint32(v31)+128))
	v621 = base.F64_ceil(base.F64_mul(v601, base.F64_sub(float64(1), v617)))
	goto L162
L161:
	;
	v621 = v601
	goto L162
L162:
	;
	if base.F64_gt(v621, float64(0)) == int32(0) {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v636 = float64(0)
	v637 = v611
	v639 = v613
	goto L105
L164:
	;
	goto L165
L165:
	;
	if base.F64_gt(v621, float64(1)) == int32(0) {
		v636 = v612
		v637 = v611
		v639 = v613
		goto L105
	} else {
		goto L166
	}
L166:
	;
	v633 = *(*float64)(unsafe.Add(mBase, uint32(v27)+16))
	v636 = base.F64_add(base.F64_mul(base.F64_add(v621, float64(-1)), v633), v612)
	v637 = v611
	v639 = v613
	goto L105
L167:
	;
	m.G0 = v27 + int32(80)
	return
L168:
	;
	if v29 == int32(342) {
		goto L171
	} else {
		goto L172
	}
L169:
	;
	goto L170
L170:
	;
	v761 = float64(0)
	v766 = *(*float64)(unsafe.Add(mBase, uint32(v27)+24))
	v767 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v27)+72)) = v767
	*(*int64)(unsafe.Add(mBase, uint32(v27)+64)) = v767
	*(*int32)(unsafe.Add(mBase, uint32(v27)+56)) = l1
	if v400 == int32(0) {
		v828 = v761
		v848 = v761
		goto L215
	} else {
		goto L216
	}
L171:
	;
	v649 = float64(-1)
	goto L173
L172:
	;
	v649 = v637
	goto L173
L173:
	;
	v650 = *(*float64)(unsafe.Add(mBase, uint32(v27)))
	v652 = *(*int32)(unsafe.Add(mBase, _consts[605]))
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v31)+148))
	if v655 != int32(-1) {
		v744 = v655
		goto L176
	} else {
		goto L177
	}
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v754
	if v754 <= int32(0) {
		goto L167
	} else {
		goto L214
	}
L175:
	;
	goto L174
L176:
	;
	if v744 < v652 {
		goto L211
	} else {
		goto L212
	}
L177:
	;
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v658 != 0 {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v675 = int32(0)
	if base.F64_ge(v649, float64(0)) == v675 {
		v707 = v675
		goto L186
	} else {
		goto L187
	}
L179:
	;
	if base.F64_ge(v649, float64(0)) != 0 {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v663 = *(*int32)(unsafe.Add(mBase, _consts[606]))
	if base.F64_lt(v649, base.F64_convert_i32_s(v663)) != 0 {
		v754 = int32(0)
		goto L175
	} else {
		goto L183
	}
L181:
	;
	goto L182
L182:
	;
	if base.F64_ge(v650, float64(0)) == int32(0) {
		goto L178
	} else {
		goto L184
	}
L183:
	;
	goto L182
L184:
	;
	v672 = *(*int32)(unsafe.Add(mBase, _consts[624]))
	if base.F64_lt(v650, base.F64_convert_i32_s(v672)) != 0 {
		v754 = int32(0)
		goto L175
	} else {
		goto L185
	}
L185:
	;
	goto L178
L186:
	;
	if base.F64_ge(v650, float64(0)) == int32(0) {
		v744 = v707
		goto L176
	} else {
		goto L195
	}
L187:
	;
	v680 = int32(1)
	v682 = *(*int32)(unsafe.Add(mBase, _consts[606]))
	if v682 <= v680 {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v685 = v680
	goto L190
L189:
	;
	v685 = v682
	goto L190
L190:
	;
	v687 = v685
	v691 = int32(1)
	goto L191
L191:
	;
	v694 = v687 * int32(3)
	if base.F64_ge(v649, base.F64_convert_i32_u(v694)) == int32(0) {
		v707 = v691
		goto L186
	} else {
		goto L193
	}
L192:
	;
	v707 = v700
	goto L186
L193:
	;
	v700 = v691 + int32(1)
	if v694 < int32(715827883) {
		v687 = v694
		v691 = v700
		goto L191
	} else {
		goto L194
	}
L194:
	;
	goto L192
L195:
	;
	v713 = int32(1)
	v715 = *(*int32)(unsafe.Add(mBase, _consts[624]))
	if v715 <= v713 {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	v718 = v713
	goto L198
L197:
	;
	v718 = v715
	goto L198
L198:
	;
	v720 = v718
	v725 = int32(1)
	goto L199
L199:
	;
	v727 = v720 * int32(3)
	if base.F64_le(base.F64_convert_i32_u(v727), v650) != 0 {
		goto L201
	} else {
		goto L202
	}
L200:
	;
	if v707 < v734 {
		goto L205
	} else {
		goto L206
	}
L201:
	;
	v731 = v725 + int32(1)
	if v727 < int32(715827883) {
		v720 = v727
		v725 = v731
		goto L199
	} else {
		goto L204
	}
L202:
	;
	v734 = v725
	goto L203
L203:
	;
	goto L200
L204:
	;
	v734 = v731
	goto L203
L205:
	;
	v736 = v707
	goto L207
L206:
	;
	v736 = v734
	goto L207
L207:
	;
	if int32(0) < v707 {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	v739 = v736
	goto L210
L209:
	;
	v739 = v734
	goto L210
L210:
	;
	v744 = v739
	goto L176
L211:
	;
	v747 = v744
	goto L213
L212:
	;
	v747 = v652
	goto L213
L213:
	;
	v754 = v747
	goto L175
L214:
	;
	v758 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)) = uint8(v758)
	goto L170
L215:
	;
	v850 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v851 = *(*float64)(unsafe.Add(mBase, uint32(v850)+24))
	v852 = *(*float64)(unsafe.Add(mBase, uint32(l0)+32))
	v855 = *(*float64)(unsafe.Add(mBase, _consts[601]))
	v860 = base.F64_add(base.F64_mul(v851, v852), base.F64_add(base.F64_mul(base.F64_add(v828, v855), v442), float64(0)))
	v862 = *(*float64)(unsafe.Add(mBase, uint32(v850)+16))
	v863 = base.F64_add(base.F64_add(base.F64_add(v443, v761), v848), v862)
	v864 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if int32(0) < v864 {
		goto L222
	} else {
		goto L223
	}
L216:
	;
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v400)+4))
	if v781 <= int32(0) {
		v828 = v761
		v848 = float64(0)
		goto L215
	} else {
		goto L217
	}
L217:
	;
	v799 = int32(0)
	goto L218
L218:
	;
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v400)+12))
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v809+v799<<(uint(int32(2))%32))))
	v816 = F_cost_qual_eval_walker(m, v813, v27+int32(56))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L34
	} else {
		goto L220
	}
L219:
	;
	v822 = *(*float64)(unsafe.Add(mBase, uint32(v27)+72))
	v823 = *(*float64)(unsafe.Add(mBase, uint32(v27)+64))
	v828 = v822
	v848 = v823
	goto L215
L220:
	;
	v819 = v799 + int32(1)
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v400)+4))
	if v819 < v820 {
		v799 = v819
		goto L218
	} else {
		goto L221
	}
L221:
	;
	goto L219
L222:
	;
	v867 = base.F64_convert_i32_u(v864)
	v869 = int32(*(*uint8)(unsafe.Add(mBase, _consts[625])))
	if v869 == int32(1) {
		goto L225
	} else {
		goto L226
	}
L223:
	;
	v899 = v860
	goto L224
L224:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = v863
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = base.F64_add(v863, base.F64_add(base.F64_add(base.F64_add(base.F64_sub(v424, v443), v761), base.F64_add(base.F64_mul(base.F64_mul(v766, v766), base.F64_sub(v636, v639)), v639)), v899))
	goto L167
L225:
	;
	v875 = base.F64_add(base.F64_mul(v867, float64(-0.3)), float64(1))
	if base.F64_gt(v875, float64(0)) != 0 {
		goto L228
	} else {
		goto L229
	}
L226:
	;
	v881 = v867
	goto L227
L227:
	;
	v883 = float64(1e+100)
	v884 = base.F64_div(v852, v881)
	if base.F64_gt(v884, v883) != 0 {
		v896 = v883
		goto L231
	} else {
		goto L232
	}
L228:
	;
	v879 = v875
	goto L230
L229:
	;
	v879 = math.Float64frombits(uint64(0x8000000000000000))
	goto L230
L230:
	;
	v881 = base.F64_add(v879, v867)
	goto L227
L231:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v896
	v899 = base.F64_div(v860, v881)
	goto L224
L232:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v884)&int64(9223372036854775807)) {
		v896 = v883
		goto L231
	} else {
		goto L233
	}
L233:
	;
	v892 = float64(1)
	if base.F64_le(v884, v892) != 0 {
		v896 = v892
		goto L231
	} else {
		goto L234
	}
L234:
	;
	v896 = base.F64_nearest(v884)
	goto L231
}
func F_index_am_handler_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(209287)
			F_errmsg(m, int32(182861), v4)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(472671), int32(371), int32(266552))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
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
func F_index_form_tuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v5 = *(*int32)(unsafe.Add(mBase, _consts[28]))
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
	F_errmsg_internal(m, int32(43981), v10)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(474666), int32(190), int32(236848))
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
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(127710)
			*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v18 + int32(4)
			F_errmsg_internal(m, int32(653155), v6)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				F_errfinish(m, int32(475543), int32(415), int32(127720))
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
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v17 = *(*int32)(unsafe.Add(mBase, _consts[113]))
	if v17 != v13 {
		v20 = *(*int32)(unsafe.Add(mBase, _consts[114]))
		v21 = F_list_member_ptr(m, v20, v13)
		mBase = m.M
		v22 = v21
	} else {
		v22 = int32(1)
	}
	if v22 == int32(0) {
		v26 = F_EstimateSnapshotSpace(m, l3)
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int32(0)
		} else {
			v30 = F_add_size(m, int32(32), v26)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				v35 = (v30 + int32(7)) & int32(-8)
				if l4 != 0 {
					v40 = F_add_size(m, v35, l6<<(uint(int32(3))%32)+int32(8))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						v46 = (v40 + int32(7)) & int32(-8)
						if l5 == int32(0) {
							v57 = v46
							m.G0 = v11 + int32(16)
							return v57
						} else {
							v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
							v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+120))
							if v50 == int32(0) {
								v57 = v46
								m.G0 = v11 + int32(16)
								return v57
							} else {
								v53 = m.T0[v50].(func(*base.Module, int32, int32, int32) int32)(m, l0, l1, l2)
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									v55 = F_add_size(m, v46, v53)
									mBase = m.M
									v56 = m.ExcPending
									if v56 != 0 {
										return int32(0)
									} else {
										v57 = v55
										m.G0 = v11 + int32(16)
										return v57
									}
								}
							}
						}
					}
				} else {
					v46 = v35
					if l5 == int32(0) {
						v57 = v46
						m.G0 = v11 + int32(16)
						return v57
					} else {
						v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
						v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+120))
						if v50 == int32(0) {
							v57 = v46
							m.G0 = v11 + int32(16)
							return v57
						} else {
							v53 = m.T0[v50].(func(*base.Module, int32, int32, int32) int32)(m, l0, l1, l2)
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								v55 = F_add_size(m, v46, v53)
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return int32(0)
								} else {
									v57 = v55
									m.G0 = v11 + int32(16)
									return v57
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
		v66 = m.ExcPending
		if v66 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v69 = m.ExcPending
			if v69 != 0 {
				return int32(0)
			} else {
				v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = v70 + int32(4)
				F_errmsg(m, int32(419294), v11)
				mBase = m.M
				v76 = m.ExcPending
				if v76 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(475543), int32(469), int32(338733))
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
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
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
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
		goto L10
	case 1:
		goto L11
	default:
		goto L9
	}
L5:
	;
	goto L3
L6:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	*(*uint8)(unsafe.Add(mBase, uint32(v90+v15))) = uint8(v89)
	v94 = v15 + int32(1)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v94 < v95 {
		v15 = v94
		goto L4
	} else {
		goto L27
	}
L7:
	;
	v89 = int32(1)
	goto L6
L8:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v82+v19))) = int32(0)
	goto L7
L9:
	;
	v64 = int32(1)
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+84)))
	if v65 != v64 {
		v89 = v64
		goto L6
	} else {
		goto L23
	}
L10:
	;
	if l2 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L11:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v15))))
	if v26 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v29+v19)))
	F_pfree(m, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	if l2 == int32(0) {
		goto L8
	} else {
		goto L17
	}
L15:
	;
	return
L16:
	;
	goto L14
L17:
	;
	v38 = l2 + v15<<(uint(int32(4))%32)
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+8)))
	if v39 != 0 {
		goto L8
	} else {
		goto L18
	}
L18:
	;
	v40 = *(*float64)(unsafe.Add(mBase, uint32(v38)))
	v41 = F_Float8GetDatum(m, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L15
	} else {
		goto L19
	}
L19:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v43+v19))) = v41
	v89 = int32(0)
	goto L6
L20:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v60+v19))) = int32(0)
	goto L7
L21:
	;
	v51 = l2 + v15<<(uint(int32(4))%32)
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+8)))
	if v52 != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v55 = *(*float64)(unsafe.Add(mBase, uint32(v51)))
	*(*float32)(unsafe.Add(mBase, uint32(v53+v19))) = base.F32_demote_f64(v55)
	v89 = int32(0)
	goto L6
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L15
	} else {
		goto L24
	}
L24:
	;
	F_errmsg_internal(m, int32(10932), int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L15
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(475543), int32(1030), int32(162256))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L15
	} else {
		goto L26
	}
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L27:
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
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
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
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
	v31 = int32(*(*uint8)(unsafe.Add(mBase, _consts[288])))
	v36 = (v31 ^ int32(1)) & base.F64_ge(v28, float64(0))
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+119)))
	v39 = v37 - int32(109)
	if base.Ui32(int32(7)) < base.Ui32(v39) {
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
	v85 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L14
	} else {
		goto L18
	}
L5:
	;
	v67 = F_RelationGetNumberOfBlocksInFork(m, l0, int32(0))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L14
	} else {
		goto L15
	}
L6:
	;
	if v36 == int32(0) {
		v81 = v4
		v82 = v4
		goto L4
	} else {
		goto L13
	}
L7:
	;
	if int32(1)<<(uint(v39)%32)&int32(161) == int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, _consts[374])))
	v51 = int32(*(*uint8)(unsafe.Add(mBase, _consts[375])))
	goto L9
L9:
	;
	if v49&v51&int32(1) == int32(0) {
		v81 = v4
		v82 = v4
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v57 == int32(0) {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+16)))
	if v36&v60 != 0 {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	v81 = v4
	v82 = v4
	goto L4
L13:
	;
	goto L5
L14:
	;
	return
L15:
	;
	v69 = int32(1)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+119)))
	if v71 == int32(105) {
		v81 = v69
		v82 = v67
		goto L4
	} else {
		goto L16
	}
L16:
	;
	F_visibilitymap_count(m, l0, v13+int32(76), v13+int32(72))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v81 = v69
	v82 = v67
	goto L4
L18:
	;
	F_ScanKeyInit(m, v13+int32(16), int32(1), int32(3), int32(184), v29)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	F_systable_inplace_update_begin(m, v85, int32(2662), v13+int32(16), v13+int32(12), v13+int32(8))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L14
	} else {
		goto L20
	}
L20:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v103 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+16))
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+22)))
	v106 = v104 + v105
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+116)))
	if v2 != v107 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L14
	} else {
		goto L48
	}
L24:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v106)+116)) = uint8(v2)
	goto L26
L25:
	;
	goto L26
L26:
	;
	v110 = base.B2i32(v2 != v107)
	if v81 == int32(0) {
		v134 = v110
		goto L30
	} else {
		goto L31
	}
L27:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	F_pfree(m, v155)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L14
	} else {
		goto L46
	}
L28:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	F_systable_inplace_update_cancel(m, v146)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L14
	} else {
		goto L44
	}
L29:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	F_systable_inplace_update_finish(m, v142, v143)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L14
	} else {
		goto L43
	}
L30:
	;
	if v134 == int32(0) {
		goto L28
	} else {
		goto L42
	}
L31:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v106)+96))
	if v82 != v113 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v106)+96)) = v82
	v117 = int32(1)
	goto L34
L33:
	;
	v117 = v110
	goto L34
L34:
	;
	v118 = base.F32_demote_f64(v28)
	v119 = *(*float32)(unsafe.Add(mBase, uint32(v106)+100))
	if base.F32_ne(v118, v119) != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v106)+100)) = v118
	v123 = int32(1)
	goto L37
L36:
	;
	v123 = v117
	goto L37
L37:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+76))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v106)+104))
	if v124 != v125 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v106)+104)) = v124
	v129 = int32(1)
	goto L40
L39:
	;
	v129 = v123
	goto L40
L40:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v13)+72))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v106)+108))
	if v130 == v131 {
		v134 = v129
		goto L30
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v106)+108)) = v130
	goto L29
L42:
	;
	goto L29
L43:
	;
	goto L27
L44:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	F_CacheInvalidateRelcacheByTuple(m, v149)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L14
	} else {
		goto L45
	}
L45:
	;
	goto L27
L46:
	;
	F_sequence_close(m, v85, int32(3))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L14
	} else {
		goto L47
	}
L47:
	;
	m.G0 = v13 + int32(80)
	return
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v29
	F_errmsg_internal(m, int32(43870), v13)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L14
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(471121), int32(2918), int32(118441))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L14
	} else {
		goto L50
	}
L50:
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
	v17 = F_expression_tree_walker_impl(m, l0, int32(824), int32(0))
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
	v36 = F_expression_tree_mutator_impl(m, v23, int32(825), int32(0))
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
	F_errmsg_internal(m, int32(136165), int32(0))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L14
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(476148), int32(4468), int32(408464))
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
	F_errmsg_internal(m, int32(136165), int32(0))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L14
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(476148), int32(4473), int32(408464))
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
