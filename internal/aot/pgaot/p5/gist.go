package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_gistFetchTuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v41 int32
	_ = v41
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	v4 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(176)
	m.G0 = v17
	v19 = int32(4536272)
	v20 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+192))
	v25 = int32(*(*int16)(unsafe.Add(mBase, uint32(v24)+10)))
	if v4 < v25 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v41 = v4
	goto L4
L2:
	;
	v127 = v4
	goto L3
L3:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	if v127 < v137 {
		goto L23
	} else {
		goto L24
	}
L4:
	;
	v51 = v41 + int32(1)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v53 = v17 + v41
	v54 = F_index_getattr_2(m, l2, v51, v52, v53)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v127 = v51
	goto L3
L6:
	;
	return int32(0)
L7:
	;
	v59 = v41 * int32(28)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(7192)+v59)))
	if v61 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l1)+192))
	v120 = int32(*(*int16)(unsafe.Add(mBase, uint32(v119)+10)))
	if v51 < v120 {
		v41 = v51
		goto L4
	} else {
		goto L22
	}
L9:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v62 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v59+(l0+int32(1816)))))
	if v95 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	v65 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+174)) = uint8(v65)
	*(*uint16)(unsafe.Add(mBase, uint32(v17)+172)) = uint16(v65)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+168)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v17)+164)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v17)+160)) = v54
	v74 = v41 << (uint(int32(2)) % 32)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v74+(l0+int32(8084)))))
	v83 = F_FunctionCall1Coll(m, v59+(l0+int32(7188)), v80, v17+int32(160))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L6
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17+int32(32)+v41<<(uint(int32(2))%32)))) = int32(0)
	goto L8
L15:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	*(*int32)(unsafe.Add(mBase, uint32(v74+(v17+int32(32))))) = v85
	goto L8
L16:
	;
	v102 = v17 + int32(32) + v41<<(uint(int32(2))%32)
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v103 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v109 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v53))) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v17+int32(32)+v41<<(uint(int32(2))%32)))) = int32(0)
	goto L8
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v102))) = v54
	goto L8
L20:
	;
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v102))) = int32(0)
	goto L8
L22:
	;
	goto L5
L23:
	;
	v144 = v127
	goto L26
L24:
	;
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v20
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v187 = F_heap_form_tuple(m, v184, v17+int32(32), v17)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L6
	} else {
		goto L30
	}
L26:
	;
	v159 = v144 + int32(1)
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v162 = F_index_getattr_2(m, l2, v159, v160, v17+v144)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L6
	} else {
		goto L28
	}
L27:
	;
	goto L25
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17+int32(32)+v144<<(uint(int32(2))%32)))) = v162
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	if v159 < v166 {
		v144 = v159
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	m.G0 = v17 + int32(176)
	return v187
}
func F_gistGetNodeBuffer(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v19 = F_hash_search(m, v13, v10+int32(12), int32(1), v10+int32(11))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+11)))
		if v23 == int32(0) {
			v26 = int32(4536272)
			v27 = *(*int32)(unsafe.Add(mBase, _consts[10]))
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*int32)(unsafe.Add(mBase, _consts[10])) = v29
			*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = l2
			v32 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(v19)+16)) = uint16(v32)
			*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v32
			*(*int64)(unsafe.Add(mBase, uint32(v19)+4)) = int64(-4294967296)
			v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			if v38 <= l2 {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				v42 = l2 + int32(1)
				v45 = F_repalloc(m, v40, v42<<(uint(int32(2))%32))
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v45
					v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
					if v48 <= l2 {
						v51 = v48
						for {
							v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							*(*int32)(unsafe.Add(mBase, uint32(v57+v51<<(uint(int32(2))%32)))) = int32(0)
							v64 = v51 + int32(1)
							if v64 <= l2 {
								v51 = v64
								continue
							} else {
								break
							}
							break
						}
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v42
					v82 = l2 << (uint(int32(2)) % 32)
					v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					v85 = *(*int32)(unsafe.Add(mBase, uint32(v82+v83)))
					v86 = F_lcons(m, v19, v85)
					mBase = m.M
					v87 = m.ExcPending
					if v87 != 0 {
						return int32(0)
					} else {
						v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						*(*int32)(unsafe.Add(mBase, uint32(v88+v82))) = v86
						*(*int32)(unsafe.Add(mBase, _consts[10])) = v27
						m.G0 = v10 + int32(16)
						return v19
					}
				}
			} else {
				v82 = l2 << (uint(int32(2)) % 32)
				v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				v85 = *(*int32)(unsafe.Add(mBase, uint32(v82+v83)))
				v86 = F_lcons(m, v19, v85)
				mBase = m.M
				v87 = m.ExcPending
				if v87 != 0 {
					return int32(0)
				} else {
					v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					*(*int32)(unsafe.Add(mBase, uint32(v88+v82))) = v86
					*(*int32)(unsafe.Add(mBase, _consts[10])) = v27
					m.G0 = v10 + int32(16)
					return v19
				}
			}
		} else {
			m.G0 = v10 + int32(16)
			return v19
		}
	}
}
func F_gistLoadNodeBuffer(m *base.Module, l0 int32, l1 int32) {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v11 != 0 {
		m.G0 = v9 + int32(16)
		return
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		if v12 <= int32(0) {
			m.G0 = v9 + int32(16)
			return
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v17 = F_MemoryContextAllocZero(m, v15, int32(8192))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v17))) = int64(35154307317759)
				*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v17
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v25 = F_BufFileSeekBlock(m, v22, base.I64_extend_i32_s(v23))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					if v25 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v23
							F_errmsg_internal(m, int32(395632), v9)
							mBase = m.M
							v98 = m.ExcPending
							if v98 != 0 {
								return
							} else {
								F_errfinish(m, int32(504795), int32(753), int32(323863))
								mBase = m.M
								v103 = m.ExcPending
								if v103 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						F_BufFileReadExact(m, v22, v17, int32(8192))
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							if v31 < v32 {
								v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v45 = v31
								v46 = v34
								*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v45 + int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v46+v45<<(uint(int32(2))%32)))) = v30
								v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+17)))
								if v54 == int32(0) {
									v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
									v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
									if v57 < v58 {
										v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
										v71 = v57
										v72 = v60
										*(*int32)(unsafe.Add(mBase, uint32(v72+v71<<(uint(int32(2))%32)))) = l1
										v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v77 + int32(1)
										*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(-1)
										m.G0 = v9 + int32(16)
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v58 << (uint(int32(1)) % 32)
										v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
										v67 = F_repalloc(m, v64, v58<<(uint(int32(3))%32))
										mBase = m.M
										v68 = m.ExcPending
										if v68 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v67
											v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
											v71 = v70
											v72 = v67
											*(*int32)(unsafe.Add(mBase, uint32(v72+v71<<(uint(int32(2))%32)))) = l1
											v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v77 + int32(1)
											*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(-1)
											m.G0 = v9 + int32(16)
											return
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(-1)
									m.G0 = v9 + int32(16)
									return
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v32 << (uint(int32(1)) % 32)
								v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v41 = F_repalloc(m, v38, v32<<(uint(int32(3))%32))
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v41
									v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									v45 = v44
									v46 = v41
									*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v45 + int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v46+v45<<(uint(int32(2))%32)))) = v30
									v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+17)))
									if v54 == int32(0) {
										v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
										v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
										if v57 < v58 {
											v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
											v71 = v57
											v72 = v60
											*(*int32)(unsafe.Add(mBase, uint32(v72+v71<<(uint(int32(2))%32)))) = l1
											v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v77 + int32(1)
											*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(-1)
											m.G0 = v9 + int32(16)
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v58 << (uint(int32(1)) % 32)
											v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
											v67 = F_repalloc(m, v64, v58<<(uint(int32(3))%32))
											mBase = m.M
											v68 = m.ExcPending
											if v68 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v67
												v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
												v71 = v70
												v72 = v67
												*(*int32)(unsafe.Add(mBase, uint32(v72+v71<<(uint(int32(2))%32)))) = l1
												v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v77 + int32(1)
												*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(-1)
												m.G0 = v9 + int32(16)
												return
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(-1)
										m.G0 = v9 + int32(16)
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
func F_gistSplitByKey(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v257 int32
	_ = v257
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v322 int32
	_ = v322
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v519 int32
	_ = v519
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v564 int32
	_ = v564
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v705 int32
	_ = v705
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v726 int32
	_ = v726
	var v735 int32
	_ = v735
	var v744 int32
	_ = v744
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v756 float32
	_ = v756
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v762 float32
	_ = v762
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v769 int32
	_ = v769
	var v773 float32
	_ = v773
	var v774 int32
	_ = v774
	var v777 int32
	_ = v777
	var v781 float32
	_ = v781
	var v782 int32
	_ = v782
	var v786 int32
	_ = v786
	var v790 float32
	_ = v790
	var v791 int32
	_ = v791
	var v794 int32
	_ = v794
	var v798 float32
	_ = v798
	var v799 int32
	_ = v799
	var v805 int32
	_ = v805
	var v808 int64
	_ = v808
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v818 int32
	_ = v818
	var v837 int32
	_ = v837
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v871 int32
	_ = v871
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v886 int32
	_ = v886
	var v888 int32
	_ = v888
	var v900 int32
	_ = v900
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v909 int32
	_ = v909
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v925 int32
	_ = v925
	var v927 int32
	_ = v927
	var v938 int32
	_ = v938
	var v940 int32
	_ = v940
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v968 int32
	_ = v968
	var v973 float32
	_ = v973
	var v974 int32
	_ = v974
	var v977 int32
	_ = v977
	var v979 int32
	_ = v979
	var v983 int32
	_ = v983
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v997 int32
	_ = v997
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1027 int32
	_ = v1027
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1068 int32
	_ = v1068
	var v1073 float32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1077 int32
	_ = v1077
	var v1079 int32
	_ = v1079
	var v1083 int32
	_ = v1083
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1132 int32
	_ = v1132
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1151 int32
	_ = v1151
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1169 int32
	_ = v1169
	var v1177 int32
	_ = v1177
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1181 int32
	_ = v1181
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1192 int32
	_ = v1192
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1229 int32
	_ = v1229
	var v1231 int32
	_ = v1231
	var v1236 int32
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1238 int32
	_ = v1238
	var v1246 int32
	_ = v1246
	var v1249 int32
	_ = v1249
	var v1251 int32
	_ = v1251
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1280 int32
	_ = v1280
	var v1292 int32
	_ = v1292
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1299 int32
	_ = v1299
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1317 int32
	_ = v1317
	var v1325 int32
	_ = v1325
	var v1326 int32
	_ = v1326
	var v1327 int32
	_ = v1327
	var v1329 int32
	_ = v1329
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1338 int32
	_ = v1338
	var v1340 int32
	_ = v1340
	var v1352 int32
	_ = v1352
	var v1353 int32
	_ = v1353
	var v1354 int32
	_ = v1354
	var v1377 int32
	_ = v1377
	var v1379 int32
	_ = v1379
	var v1383 int32
	_ = v1383
	var v1385 int32
	_ = v1385
	var v1396 int32
	_ = v1396
	var v1398 int32
	_ = v1398
	var v1418 int32
	_ = v1418
	var v1424 int32
	_ = v1424
	var v1426 int32
	_ = v1426
	var v1427 int32
	_ = v1427
	var v1430 int32
	_ = v1430
	var v1431 int32
	_ = v1431
	var v1434 int32
	_ = v1434
	var v1442 int32
	_ = v1442
	var v1444 int32
	_ = v1444
	var v1466 int32
	_ = v1466
	var v1468 int32
	_ = v1468
	var v1470 int32
	_ = v1470
	var v1479 int32
	_ = v1479
	var v1481 int32
	_ = v1481
	var v1507 int32
	_ = v1507
	var v1513 int32
	_ = v1513
	var v1514 int32
	_ = v1514
	var v1515 int32
	_ = v1515
	var v1535 int32
	_ = v1535
	var v1550 int32
	_ = v1550
	var v1552 int32
	_ = v1552
	var v1553 int32
	_ = v1553
	var v1564 int32
	_ = v1564
	var v1569 int32
	_ = v1569
	var v1572 int32
	_ = v1572
	var v1573 int32
	_ = v1573
	var v1574 float32
	_ = v1574
	var v1575 int32
	_ = v1575
	var v1577 int32
	_ = v1577
	var v1578 int32
	_ = v1578
	var v1589 int32
	_ = v1589
	var v1590 int32
	_ = v1590
	var v1591 float32
	_ = v1591
	var v1592 int32
	_ = v1592
	var v1599 int32
	_ = v1599
	var v1600 int32
	_ = v1600
	var v1601 int32
	_ = v1601
	var v1633 int32
	_ = v1633
	var v1637 int32
	_ = v1637
	var v1641 int32
	_ = v1641
	var v1642 int32
	_ = v1642
	var v1643 int32
	_ = v1643
	var v1644 int32
	_ = v1644
	var v1645 int32
	_ = v1645
	var v1648 int32
	_ = v1648
	var v1655 int32
	_ = v1655
	var v1664 int32
	_ = v1664
	var v1666 int32
	_ = v1666
	var v1667 int32
	_ = v1667
	var v1687 int32
	_ = v1687
	var v1688 int32
	_ = v1688
	var v1689 int32
	_ = v1689
	var v1691 int32
	_ = v1691
	var v1694 int32
	_ = v1694
	var v1700 int32
	_ = v1700
	var v1702 int32
	_ = v1702
	var v1708 int32
	_ = v1708
	var v1710 int32
	_ = v1710
	var v1711 int32
	_ = v1711
	var v1713 int32
	_ = v1713
	var v1716 int32
	_ = v1716
	var v1722 int32
	_ = v1722
	var v1724 int32
	_ = v1724
	var v1730 int32
	_ = v1730
	var v1732 int32
	_ = v1732
	var v1741 int32
	_ = v1741
	var v1744 int32
	_ = v1744
	var v1766 int32
	_ = v1766
	var v1767 int32
	_ = v1767
	var v1768 int32
	_ = v1768
	var v1770 int32
	_ = v1770
	var v1773 int32
	_ = v1773
	var v1779 int32
	_ = v1779
	var v1781 int32
	_ = v1781
	var v1797 int32
	_ = v1797
	var v1817 int64
	_ = v1817
	var v1818 int32
	_ = v1818
	var v1819 int64
	_ = v1819
	var v1820 int32
	_ = v1820
	var v1821 int32
	_ = v1821
	var v1822 int32
	_ = v1822
	var v1823 int32
	_ = v1823
	var v1824 int32
	_ = v1824
	var v1826 int32
	_ = v1826
	var v1827 int32
	_ = v1827
	var v1828 int32
	_ = v1828
	var v1829 int32
	_ = v1829
	var v1830 int32
	_ = v1830
	var v1831 int32
	_ = v1831
	var v1832 int32
	_ = v1832
	var v1834 int32
	_ = v1834
	var v1835 int32
	_ = v1835
	var v1836 int32
	_ = v1836
	var v1838 int32
	_ = v1838
	var v1840 int32
	_ = v1840
	var v1841 int32
	_ = v1841
	var v1853 int32
	_ = v1853
	var v1855 int32
	_ = v1855
	var v1874 int32
	_ = v1874
	var v1877 int32
	_ = v1877
	var v1881 int32
	_ = v1881
	var v1885 int32
	_ = v1885
	var v1888 int32
	_ = v1888
	var v1890 int32
	_ = v1890
	var v1891 int32
	_ = v1891
	var v1902 int32
	_ = v1902
	var v1923 int32
	_ = v1923
	var v1924 int32
	_ = v1924
	var v1934 int32
	_ = v1934
	var v1938 int32
	_ = v1938
	var v1957 int32
	_ = v1957
	var v1960 int32
	_ = v1960
	var v1964 int32
	_ = v1964
	var v1968 int32
	_ = v1968
	var v1971 int32
	_ = v1971
	var v1973 int32
	_ = v1973
	var v1974 int32
	_ = v1974
	var v1983 int32
	_ = v1983
	var v2053 int32
	_ = v2053
	var v2073 int32
	_ = v2073
	var v2074 int32
	_ = v2074
	var v2075 int32
	_ = v2075
	var v2076 int32
	_ = v2076
	var v2077 int32
	_ = v2077
	var v2144 int32
	_ = v2144
	var v2145 int32
	_ = v2145
	var v2151 int32
	_ = v2151
	v31 = m.G0
	v33 = v31 - int32(608)
	m.G0 = v33
	v36 = l3 + int32(1)
	v37 = int32(4)
	v41 = F_palloc(m, v36<<(uint(v37)%32)|v37)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41))) = v36
	v45 = l3 << (uint(int32(1)) % 32)
	v46 = F_palloc(m, v45)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if int32(0) < l3 {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	m.G0 = v33 + int32(608)
	return
L5:
	;
	v2144 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v2145 = *(*int32)(unsafe.Add(mBase, uint32(v2144)))
	if v2145 < int32(2) {
		goto L4
	} else {
		goto L286
	}
L6:
	;
	if l6 != 0 {
		goto L4
	} else {
		goto L285
	}
L7:
	;
	v401 = l5 + int32(160)
	v402 = v401 + l6
	v403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v402))))
	v404 = int32(1)
	v405 = v403 ^ v404
	*(*uint8)(unsafe.Add(mBase, uint32(l5)+12)) = uint8(v405)
	v408 = l5 + int32(320)
	v409 = v408 + l6
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v409))))
	v412 = v410 ^ v404
	*(*uint8)(unsafe.Add(mBase, uint32(l5)+28)) = uint8(v412)
	v415 = l5 + int32(32)
	v417 = l6 << (uint(int32(2)) % 32)
	v418 = v415 + v417
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v418)))
	*(*int32)(unsafe.Add(mBase, uint32(l5)+8)) = v419
	v422 = l5 + int32(192)
	v423 = v422 + v417
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v423)))
	*(*int32)(unsafe.Add(mBase, uint32(l5)+24)) = v424
	v431 = l4 + v417
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v431)+uint32(_consts[32])))
	v435 = F_FunctionCall2Coll(m, l4+l6*int32(28)+int32(4500), v434, v41, l5)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L1
	} else {
		goto L63
	}
L8:
	;
	v288 = l5 + l6
	v289 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v288)+320)) = uint8(v289)
	*(*uint8)(unsafe.Add(mBase, uint32(v288)+160)) = uint8(v289)
	v294 = l6 + v289
	v295 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v295)))
	if v294 < v296 {
		goto L49
	} else {
		goto L50
	}
L9:
	;
	v50 = int32(4)
	v53 = v41 + v50
	v56 = int32(1)
	v70 = v56
	v78 = int32(0)
	goto L12
L10:
	;
	goto L11
L11:
	;
	if l3 != 0 {
		goto L7
	} else {
		goto L48
	}
L12:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l2-v50+v70<<(uint(int32(2))%32))))
	v101 = int32(*(*int16)(unsafe.Add(mBase, uint32(v100)+6)))
	if int32(0) <= v101 {
		goto L18
	} else {
		goto L19
	}
L13:
	;
	if l3 == v177 {
		goto L8
	} else {
		goto L35
	}
L14:
	;
	v180 = v70 + int32(1)
	if v180 <= l3 {
		v70 = v180
		v78 = v177
		goto L12
	} else {
		goto L34
	}
L15:
	;
	F_gistdentryinit(m, l4, l6, v53+v70<<(uint(int32(4))%32), int32(0), l0, l1, v70&int32(65535), int32(1))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L33
	}
L16:
	;
	F_gistdentryinit(m, l4, l6, v53+v70<<(uint(int32(4))%32), v153, l0, l1, v70&int32(65535), int32(0))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L32
	}
L17:
	;
	v147 = F_nocache_index_getattr(m, v100, l6+v56, v93)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L31
	}
L18:
	;
	v108 = v93 + l6<<(uint(int32(4))%32) + int32(20)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	if v109 < int32(0) {
		goto L17
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100+l6>>(uint(int32(3))%32))+8)))
	if v56<<(uint(l6&int32(7))%32)&v141 == int32(0) {
		goto L15
	} else {
		goto L30
	}
L21:
	;
	v114 = v100 + v109 + int32(8)
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+6)))
	if v115 != int32(1) {
		v153 = v114
		goto L16
	} else {
		goto L22
	}
L22:
	;
	v118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v108)+4)))
	switch v118 - int32(1) {
	case 0:
		goto L26
	case 1:
		goto L25
	default:
		goto L23
	case 3:
		goto L24
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L27
	}
L24:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	v153 = v123
	goto L16
L25:
	;
	v122 = int32(*(*int16)(unsafe.Add(mBase, uint32(v114))))
	v153 = v122
	goto L16
L26:
	;
	v121 = int32(*(*int8)(unsafe.Add(mBase, uint32(v114))))
	v153 = v121
	goto L16
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+16)) = base.I32_extend16_s(v118)
	F_errmsg_internal(m, int32(493552), v33+int32(16))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(333863), int32(70), int32(69245))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L30:
	;
	goto L17
L31:
	;
	v153 = v147
	goto L16
L32:
	;
	v177 = v78
	goto L14
L33:
	;
	v168 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v46+v78<<(uint(v168)%32)))) = uint16(v70)
	v177 = v78 + v168
	goto L14
L34:
	;
	goto L13
L35:
	;
	if v177 <= int32(0) {
		goto L7
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5)+20)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(l5)+16)) = v46
	v188 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l5+l6)+320)) = uint8(v188)
	v191 = F_palloc(m, v45)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v191
	v204 = int32(1)
	v206 = int32(0)
	goto L38
L38:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l5)+20))
	if v227 <= v206 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	if l6 != 0 {
		goto L4
	} else {
		goto L45
	}
L40:
	;
	if l3 != v204 {
		v204 = v204 + int32(1)
		v206 = v245
		goto L38
	} else {
		goto L44
	}
L41:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	v237 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l5)+4)) = v236 + v237
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	*(*uint16)(unsafe.Add(mBase, uint32(v240+v236<<(uint(v237)%32)))) = uint16(v204)
	v245 = v206
	goto L40
L42:
	;
	v232 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v46+v206<<(uint(int32(1))%32)))))
	if v204 != v232 {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v245 = v206 + int32(1)
	goto L40
L44:
	;
	goto L39
L45:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v250)))
	if v251 != int32(1) {
		goto L5
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5)+352)) = int32(0)
	F_gistunionsubkey(m, l4, l2, l5)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	goto L5
L48:
	;
	goto L8
L49:
	;
	F_gistSplitByKey(m, l0, l1, l2, l3, l4, l5, v294)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v300 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l5)+20)) = v300
	*(*int32)(unsafe.Add(mBase, uint32(l5)+4)) = v300
	v304 = F_palloc(m, v45)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L53
	}
L52:
	;
	goto L6
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v304
	v307 = F_palloc(m, v45)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5)+16)) = v307
	if l3 <= int32(0) {
		goto L6
	} else {
		goto L55
	}
L55:
	;
	v312 = int32(1)
	v322 = v312
	goto L56
L56:
	;
	if base.Ui32(v322) < base.Ui32(int32(base.Ui32(l3)>>(uint(v312)%32))) {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	goto L6
L58:
	;
	if base.B2i32(l3 == v322) == int32(0) {
		v322 = v322 + int32(1)
		goto L56
	} else {
		goto L62
	}
L59:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(l5)+20))
	v347 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l5)+20)) = v346 + v347
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l5)+16))
	*(*uint16)(unsafe.Add(mBase, uint32(v350+v346<<(uint(v347)%32)))) = uint16(v322)
	goto L58
L60:
	;
	goto L61
L61:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	v356 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l5)+4)) = v355 + v356
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	*(*uint16)(unsafe.Add(mBase, uint32(v359+v355<<(uint(v356)%32)))) = uint16(v322)
	goto L58
L62:
	;
	goto L57
L63:
	;
	v438 = l5 + int32(12)
	v439 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v439 != 0 {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	v697 = l5 + int32(24)
	v699 = l5 + int32(8)
	v701 = l5 + int32(28)
	v702 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438))))
	if v702 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L65:
	;
	v640 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v645 = v640 + v439<<(uint(int32(1))%32) - int32(2)
	v646 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v645))))
	if v646 == int32(0) {
		goto L104
	} else {
		goto L105
	}
L66:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(l5)+20))
	if v440 != 0 {
		goto L65
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v444 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L1
	} else {
		goto L70
	}
L69:
	;
	goto L68
L70:
	;
	if v444 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L1
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v469 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v402))))
	v470 = int32(1)
	v471 = v469 ^ v470
	*(*uint8)(unsafe.Add(mBase, uint32(l5)+12)) = uint8(v471)
	v473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v409))))
	v475 = v473 ^ v470
	*(*uint8)(unsafe.Add(mBase, uint32(l5)+28)) = uint8(v475)
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v418)))
	*(*int32)(unsafe.Add(mBase, uint32(l5)+8)) = v477
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v423)))
	*(*int32)(unsafe.Add(mBase, uint32(l5)+24)) = v479
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	v485 = (v481 - v470) & int32(65535)
	v489 = v485<<(uint(v470)%32) + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v489
	v491 = F_palloc(m, v489)
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L1
	} else {
		goto L78
	}
L74:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = l6 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v449 + int32(4)
	F_errmsg(m, int32(464571), v33)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	F_errhint(m, int32(659739), int32(0))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(504252), int32(448), int32(104393))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	goto L73
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v491
	v494 = F_palloc(m, v489)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v496 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l5)+20)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(l5)+16)) = v494
	*(*int32)(unsafe.Add(mBase, uint32(l5)+4)) = v496
	v502 = v481 & int32(65535)
	if v502 != int32(1) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v505 = int32(2)
	if base.Ui32(v502) <= base.Ui32(v505) {
		goto L83
	} else {
		goto L84
	}
L81:
	;
	goto L82
L82:
	;
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	v597 = int32(4)
	v601 = F_palloc(m, v596<<(uint(v597)%32)|v597)
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L1
	} else {
		goto L93
	}
L83:
	;
	v508 = v505
	goto L85
L84:
	;
	v508 = v502
	goto L85
L85:
	;
	v509 = int32(1)
	v519 = v509
	goto L86
L86:
	;
	if base.Ui32(v519) <= base.Ui32(int32(base.Ui32(v485)>>(uint(v509)%32))) {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	goto L82
L88:
	;
	v564 = v519 + int32(1)
	if v564 != v508 {
		v519 = v564
		goto L86
	} else {
		goto L92
	}
L89:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v544 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	v545 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v543+v544<<(uint(v545)%32)))) = uint16(v519)
	v549 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l5)+4)) = v549 + v545
	goto L88
L90:
	;
	goto L91
L91:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(l5)+16))
	v554 = *(*int32)(unsafe.Add(mBase, uint32(l5)+20))
	v555 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v553+v554<<(uint(v555)%32)))) = uint16(v519)
	v559 = *(*int32)(unsafe.Add(mBase, uint32(l5)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l5)+20)) = v559 + v555
	goto L88
L92:
	;
	goto L87
L93:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v601))) = v603
	v605 = int32(4)
	v606 = v601 + v605
	v608 = v41 + int32(20)
	v610 = v603 << (uint(v605) % 32)
	if v610 != 0 {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v617 = l4 + l6*int32(28) + int32(916)
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v431)+uint32(_consts[32])))
	v621 = F_FunctionCall2Coll(m, v617, v618, v601, v33+int32(96))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L1
	} else {
		goto L98
	}
L95:
	;
	v611 = F__emscripten_memcpy_bulkmem(m, v606, v608, v610)
	mBase = m.M
	v612 = v611
	goto L97
L96:
	;
	v612 = v606
	goto L97
L97:
	;
	goto L94
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5)+8)) = v621
	v624 = *(*int32)(unsafe.Add(mBase, uint32(l5)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v601))) = v624
	v626 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	v627 = int32(4)
	v631 = v624 << (uint(v627) % 32)
	if v631 != 0 {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v431)+uint32(_consts[32])))
	v637 = F_FunctionCall2Coll(m, v617, v634, v601, v33+int32(96))
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L1
	} else {
		goto L103
	}
L100:
	;
	v632 = F__emscripten_memcpy_bulkmem(m, v612, v608+v626<<(uint(v627)%32), v631)
	mBase = m.M
	goto L102
L101:
	;
	goto L102
L102:
	;
	goto L99
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5)+24)) = v637
	goto L64
L104:
	;
	v649 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41))))
	v651 = v649 - int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v645))) = uint16(v651)
	v653 = *(*int32)(unsafe.Add(mBase, uint32(l5)+20))
	v654 = v653
	goto L106
L105:
	;
	v654 = v440
	goto L106
L106:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(l5)+16))
	v660 = v655 + v654<<(uint(int32(1))%32) - int32(2)
	v661 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v660))))
	if v661 != 0 {
		goto L64
	} else {
		goto L107
	}
L107:
	;
	v662 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41))))
	v664 = v662 - int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v660))) = uint16(v664)
	goto L64
L108:
	;
	v867 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v418))) = v867
	v869 = *(*int32)(unsafe.Add(mBase, uint32(l5)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v423))) = v869
	v871 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v402))) = uint8(v871)
	*(*uint8)(unsafe.Add(mBase, uint32(v409))) = uint8(v871)
	*(*int32)(unsafe.Add(mBase, uint32(l5)+352)) = v871
	v878 = l6 + int32(1)
	v879 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v879)))
	if v880 <= v878 {
		goto L6
	} else {
		goto L139
	}
L109:
	;
	v705 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v701))))
	if v705 != int32(1) {
		goto L108
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v423)))
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v418)))
	v710 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+110)) = uint8(v710)
	*(*uint16)(unsafe.Add(mBase, uint32(v33)+108)) = uint16(v710)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v710
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v709
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+78)) = uint8(v710)
	*(*uint16)(unsafe.Add(mBase, uint32(v33)+76)) = uint16(v710)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+72)) = v710
	*(*int32)(unsafe.Add(mBase, uint32(v33)+68)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v33)+64)) = v708
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v699)))
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+62)) = uint8(v710)
	*(*uint16)(unsafe.Add(mBase, uint32(v33)+60)) = uint16(v710)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+56)) = v710
	*(*int32)(unsafe.Add(mBase, uint32(v33)+52)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v33)+48)) = v726
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v697)))
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+42)) = uint8(v710)
	*(*uint16)(unsafe.Add(mBase, uint32(v33)+40)) = uint16(v710)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+36)) = v710
	*(*int32)(unsafe.Add(mBase, uint32(v33)+32)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v33)+28)) = v735
	if v702 != 0 {
		goto L115
	} else {
		goto L116
	}
L112:
	;
	goto L111
L113:
	;
	if v805 == int32(0) {
		goto L128
	} else {
		goto L129
	}
L114:
	;
	v769 = int32(0)
	v773 = F_gistpenalty(m, l4, l6, v33+int32(96), v769, v33+int32(48), v769)
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L1
	} else {
		goto L124
	}
L115:
	;
	v744 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v701))))
	if v744 == int32(1) {
		goto L114
	} else {
		goto L118
	}
L116:
	;
	v751 = v33 - int32(-64)
	goto L117
L117:
	;
	v752 = int32(0)
	v756 = F_gistpenalty(m, l4, l6, v751, v752, v33+int32(48), v752)
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L1
	} else {
		goto L119
	}
L118:
	;
	v751 = v33 + int32(96)
	goto L117
L119:
	;
	v758 = int32(0)
	v762 = F_gistpenalty(m, l4, l6, v751, v758, v33+int32(28), v758)
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	if base.F32_lt(v756, v762) != 0 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v765 = v438
	goto L123
L122:
	;
	v765 = v701
	goto L123
L123:
	;
	v766 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v765))))
	v805 = v766
	goto L113
L124:
	;
	v777 = int32(0)
	v781 = F_gistpenalty(m, l4, l6, v33-int32(-64), v777, v33+int32(28), v777)
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	v786 = int32(0)
	v790 = F_gistpenalty(m, l4, l6, v33+int32(96), v786, v33+int32(28), v786)
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	v794 = int32(0)
	v798 = F_gistpenalty(m, l4, l6, v33-int32(-64), v794, v33+int32(48), v794)
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	v805 = base.B2i32(base.F32_gt(base.F32_add(v773, v781), base.F32_add(v790, v798)) == int32(0))
	goto L113
L128:
	;
	v808 = *(*int64)(unsafe.Add(mBase, uint32(l5)+16))
	v809 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	*(*int32)(unsafe.Add(mBase, uint32(l5)+16)) = v809
	v811 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l5)+4)))
	*(*int64)(unsafe.Add(mBase, uint32(l5))) = v808
	*(*int32)(unsafe.Add(mBase, uint32(l5)+20)) = v811
	v814 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	v815 = *(*int32)(unsafe.Add(mBase, uint32(l5)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l5)+8)) = v815
	*(*int32)(unsafe.Add(mBase, uint32(l5)+24)) = v814
	v818 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+62)) = uint8(v818)
	*(*uint16)(unsafe.Add(mBase, uint32(v33)+60)) = uint16(v818)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+56)) = v818
	*(*int32)(unsafe.Add(mBase, uint32(v33)+52)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v33)+48)) = v815
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+42)) = uint8(v818)
	*(*uint16)(unsafe.Add(mBase, uint32(v33)+40)) = uint16(v818)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+36)) = v818
	*(*int32)(unsafe.Add(mBase, uint32(v33)+32)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v33)+28)) = v814
	goto L130
L129:
	;
	goto L130
L130:
	;
	v837 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438))))
	if v837 == int32(1) {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	F_gistMakeUnionKey(m, l4, l6, v33+int32(96), v33+int32(48), v699, v33+int32(47))
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
		goto L1
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	v848 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v701))))
	if v848 == int32(1) {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	goto L133
L135:
	;
	F_gistMakeUnionKey(m, l4, l6, v33-int32(-64), v33+int32(28), v697, v33+int32(47))
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L1
	} else {
		goto L138
	}
L136:
	;
	goto L137
L137:
	;
	v859 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v701))) = uint8(v859)
	*(*uint8)(unsafe.Add(mBase, uint32(v438))) = uint8(v859)
	goto L108
L138:
	;
	goto L137
L139:
	;
	v882 = *(*int32)(unsafe.Add(mBase, uint32(v699)))
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v697)))
	v884 = m.G0
	v886 = v884 - int32(16)
	m.G0 = v886
	v888 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v886)+15)) = uint8(v888)
	v900 = *(*int32)(unsafe.Add(mBase, uint32(l4+l6<<(uint(int32(2))%32))+uint32(_consts[32])))
	v903 = F_FunctionCall3Coll(m, l4+l6*int32(28)+int32(5396), v900, v882, v883, v886+int32(15))
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L1
	} else {
		goto L141
	}
L140:
	;
	v2074 = *(*int32)(unsafe.Add(mBase, uint32(v2053)))
	v2075 = l5 + v2073
	v2076 = *(*int32)(unsafe.Add(mBase, uint32(v2075)))
	v2077 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2075))) = v2076 + v2077
	*(*uint16)(unsafe.Add(mBase, uint32(v2074+v2076<<(uint(v2077)%32)))) = uint16(v1479)
	goto L6
L141:
	;
	v905 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v886)+15)))
	m.G0 = v886 + int32(16)
	if v905 != 0 {
		goto L143
	} else {
		goto L144
	}
L142:
	;
	v2053 = l5
	v2073 = int32(4)
	goto L140
L143:
	;
	v1633 = *(*int32)(unsafe.Add(mBase, uint32(l5)+352))
	if v1633 == int32(0) {
		goto L239
	} else {
		goto L240
	}
L144:
	;
	v909 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	v912 = F_palloc0(m, v909+int32(1))
	mBase = m.M
	v913 = m.ExcPending
	if v913 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5)+352)) = v912
	v915 = *(*int32)(unsafe.Add(mBase, uint32(l5)+24))
	v916 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+110)) = uint8(v916)
	*(*uint16)(unsafe.Add(mBase, uint32(v33)+108)) = uint16(v916)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v916
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v915
	v925 = v41 + int32(4)
	v927 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v916 < v927 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v938 = int32(0)
	v940 = v916
	goto L149
L147:
	;
	v997 = v916
	goto L148
L148:
	;
	v1018 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	v1019 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+110)) = uint8(v1019)
	*(*uint16)(unsafe.Add(mBase, uint32(v33)+108)) = uint16(v1019)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1019
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v1018
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(l5)+20))
	if v1019 < v1027 {
		goto L156
	} else {
		goto L157
	}
L149:
	;
	v963 = int32(0)
	v964 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v968 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v964+v938<<(uint(int32(1))%32)))))
	v973 = F_gistpenalty(m, l4, l6, v33+int32(96), v963, v925+v968<<(uint(int32(4))%32), v963)
	mBase = m.M
	v974 = m.ExcPending
	if v974 != 0 {
		goto L1
	} else {
		goto L151
	}
L150:
	;
	v997 = v983
	goto L148
L151:
	;
	if base.F32_eq(v973, float32(0)) != 0 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v977 = *(*int32)(unsafe.Add(mBase, uint32(l5)+352))
	v979 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v977+v968))) = uint8(v979)
	v983 = v940 + v979
	goto L154
L153:
	;
	v983 = v940
	goto L154
L154:
	;
	v985 = v938 + int32(1)
	v986 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v985 < v986 {
		v938 = v985
		v940 = v983
		goto L149
	} else {
		goto L155
	}
L155:
	;
	goto L150
L156:
	;
	v1038 = int32(0)
	v1040 = v997
	goto L159
L157:
	;
	v1097 = v997
	v1098 = v1027
	goto L158
L158:
	;
	if v1097 <= int32(0) {
		goto L6
	} else {
		goto L166
	}
L159:
	;
	v1063 = int32(0)
	v1064 = *(*int32)(unsafe.Add(mBase, uint32(l5)+16))
	v1068 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1064+v1038<<(uint(int32(1))%32)))))
	v1073 = F_gistpenalty(m, l4, l6, v33+int32(96), v1063, v925+v1068<<(uint(int32(4))%32), v1063)
	mBase = m.M
	v1074 = m.ExcPending
	if v1074 != 0 {
		goto L1
	} else {
		goto L161
	}
L160:
	;
	v1097 = v1083
	v1098 = v1086
	goto L158
L161:
	;
	if base.F32_eq(v1073, float32(0)) != 0 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v1077 = *(*int32)(unsafe.Add(mBase, uint32(l5)+352))
	v1079 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1077+v1068))) = uint8(v1079)
	v1083 = v1040 + v1079
	goto L164
L163:
	;
	v1083 = v1040
	goto L164
L164:
	;
	v1085 = v1038 + int32(1)
	v1086 = *(*int32)(unsafe.Add(mBase, uint32(l5)+20))
	if v1085 < v1086 {
		v1038 = v1085
		v1040 = v1083
		goto L159
	} else {
		goto L165
	}
L165:
	;
	goto L160
L166:
	;
	v1120 = *(*int32)(unsafe.Add(mBase, uint32(l5)+352))
	v1121 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if int32(0) < v1121 {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v1124 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v1125 = int32(1)
	if v1121 == v1125 {
		goto L171
	} else {
		goto L172
	}
L168:
	;
	v1246 = v1120
	v1249 = v1098
	v1251 = v1121
	goto L169
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5)+4)) = v1251
	if int32(0) < v1249 {
		goto L190
	} else {
		goto L191
	}
L170:
	;
	if v1121&v1125 == int32(0) {
		v1236 = v1206
		goto L185
	} else {
		goto L186
	}
L171:
	;
	v1204 = v1124
	v1205 = int32(0)
	v1206 = v1121
	goto L170
L172:
	;
	goto L173
L173:
	;
	v1132 = int32(0)
	v1144 = v1124
	v1145 = v1132
	v1146 = v1121
	v1151 = v1132
	goto L174
L174:
	;
	v1166 = v1124 + v1145<<(uint(int32(1))%32)
	v1167 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1166))))
	v1169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1120+v1167))))
	if v1169 == int32(0) {
		goto L177
	} else {
		goto L178
	}
L175:
	;
	v1204 = v1187
	v1205 = v1190
	v1206 = v1188
	goto L170
L176:
	;
	v1179 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1166)+2)))
	v1181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1120+v1179))))
	if v1181 != 0 {
		goto L181
	} else {
		goto L182
	}
L177:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1144))) = uint16(v1167)
	v1177 = v1144 + int32(2)
	v1178 = v1146
	goto L176
L178:
	;
	goto L179
L179:
	;
	v1177 = v1144
	v1178 = v1146 - int32(1)
	goto L176
L180:
	;
	v1189 = int32(2)
	v1190 = v1145 + v1189
	v1192 = v1151 + v1189
	if v1192 != v1121&int32(2147483646) {
		v1144 = v1187
		v1145 = v1190
		v1146 = v1188
		v1151 = v1192
		goto L174
	} else {
		goto L184
	}
L181:
	;
	v1187 = v1177
	v1188 = v1178 - int32(1)
	goto L180
L182:
	;
	goto L183
L183:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1177))) = uint16(v1179)
	v1187 = v1177 + int32(2)
	v1188 = v1178
	goto L180
L184:
	;
	goto L175
L185:
	;
	v1237 = *(*int32)(unsafe.Add(mBase, uint32(l5)+352))
	v1238 = *(*int32)(unsafe.Add(mBase, uint32(l5)+20))
	v1246 = v1237
	v1249 = v1238
	v1251 = v1236
	goto L169
L186:
	;
	v1229 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1124+v1205<<(uint(int32(1))%32)))))
	v1231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1120+v1229))))
	if v1231 != 0 {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v1236 = v1206 - int32(1)
	goto L185
L188:
	;
	goto L189
L189:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1204))) = uint16(v1229)
	v1236 = v1206
	goto L185
L190:
	;
	v1272 = *(*int32)(unsafe.Add(mBase, uint32(l5)+16))
	v1273 = int32(1)
	if v1249 == v1273 {
		goto L194
	} else {
		goto L195
	}
L191:
	;
	v1396 = v1249
	v1398 = v1251
	goto L192
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5)+20)) = v1396
	if v1396 != 0 {
		goto L213
	} else {
		goto L214
	}
L193:
	;
	if v1249&v1273 == int32(0) {
		v1383 = v1352
		goto L208
	} else {
		goto L209
	}
L194:
	;
	v1352 = v1249
	v1353 = int32(0)
	v1354 = v1272
	goto L193
L195:
	;
	goto L196
L196:
	;
	v1280 = int32(0)
	v1292 = v1249
	v1293 = v1280
	v1294 = v1272
	v1299 = v1280
	goto L197
L197:
	;
	v1314 = v1272 + v1293<<(uint(int32(1))%32)
	v1315 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1314))))
	v1317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1246+v1315))))
	if v1317 == int32(0) {
		goto L200
	} else {
		goto L201
	}
L198:
	;
	v1352 = v1335
	v1353 = v1338
	v1354 = v1336
	goto L193
L199:
	;
	v1327 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1314)+2)))
	v1329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1246+v1327))))
	if v1329 != 0 {
		goto L204
	} else {
		goto L205
	}
L200:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1294))) = uint16(v1315)
	v1325 = v1292
	v1326 = v1294 + int32(2)
	goto L199
L201:
	;
	goto L202
L202:
	;
	v1325 = v1292 - int32(1)
	v1326 = v1294
	goto L199
L203:
	;
	v1337 = int32(2)
	v1338 = v1293 + v1337
	v1340 = v1299 + v1337
	if v1340 != v1249&int32(2147483646) {
		v1292 = v1335
		v1293 = v1338
		v1294 = v1336
		v1299 = v1340
		goto L197
	} else {
		goto L207
	}
L204:
	;
	v1335 = v1325 - int32(1)
	v1336 = v1326
	goto L203
L205:
	;
	goto L206
L206:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1326))) = uint16(v1327)
	v1335 = v1325
	v1336 = v1326 + int32(2)
	goto L203
L207:
	;
	goto L198
L208:
	;
	v1385 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	v1396 = v1383
	v1398 = v1385
	goto L192
L209:
	;
	v1377 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1272+v1353<<(uint(int32(1))%32)))))
	v1379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1246+v1377))))
	if v1379 != 0 {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	v1383 = v1352 - int32(1)
	goto L208
L211:
	;
	goto L212
L212:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1354))) = uint16(v1377)
	v1383 = v1352
	goto L208
L213:
	;
	v1418 = v1398
	goto L215
L214:
	;
	v1418 = int32(0)
	goto L215
L215:
	;
	if v1418 == int32(0) {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5)+352)) = int32(0)
	F_gistSplitByKey(m, l0, l1, l2, l3, l4, l5, v878)
	mBase = m.M
	v1424 = m.ExcPending
	if v1424 != 0 {
		goto L1
	} else {
		goto L219
	}
L217:
	;
	goto L218
L218:
	;
	F_gistunionsubkey(m, l4, l2, l5)
	mBase = m.M
	v1426 = m.ExcPending
	if v1426 != 0 {
		goto L1
	} else {
		goto L220
	}
L219:
	;
	goto L6
L220:
	;
	v1427 = int32(1)
	if v1097 != v1427 {
		goto L143
	} else {
		goto L221
	}
L221:
	;
	v1430 = int32(1)
	v1431 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	if v1431 < int32(2) {
		v1479 = v1427
		v1481 = v1430
		goto L222
	} else {
		goto L223
	}
L222:
	;
	v1507 = *(*int32)(unsafe.Add(mBase, uint32(l2+v1481<<(uint(int32(2))%32)-int32(4))))
	F_gistDeCompressAtt(m, l4, l0, v1507, v33+int32(96), v33-int32(-64))
	mBase = m.M
	v1513 = m.ExcPending
	if v1513 != 0 {
		goto L1
	} else {
		goto L228
	}
L223:
	;
	v1434 = *(*int32)(unsafe.Add(mBase, uint32(l5)+352))
	v1442 = v1427
	v1444 = v1430
	goto L224
L224:
	;
	v1466 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1434+v1444))))
	if v1466 != 0 {
		v1479 = v1442
		v1481 = v1444
		goto L222
	} else {
		goto L226
	}
L225:
	;
	v1479 = v1468
	v1481 = v1470
	goto L222
L226:
	;
	v1468 = v1442 + int32(1)
	v1470 = v1468 & int32(65535)
	if base.Ui32(v1470) < base.Ui32(v1431) {
		v1442 = v1468
		v1444 = v1470
		goto L224
	} else {
		goto L227
	}
L227:
	;
	goto L225
L228:
	;
	v1514 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v1515 = *(*int32)(unsafe.Add(mBase, uint32(v1514)))
	if v1515 <= v878 {
		goto L142
	} else {
		goto L229
	}
L229:
	;
	v1535 = v878
	goto L230
L230:
	;
	v1550 = v1535 << (uint(int32(2)) % 32)
	v1552 = *(*int32)(unsafe.Add(mBase, uint32(v415+v1550)))
	v1553 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+62)) = uint8(v1553)
	*(*uint16)(unsafe.Add(mBase, uint32(v33)+60)) = uint16(v1553)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+56)) = v1553
	*(*int32)(unsafe.Add(mBase, uint32(v33)+52)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v33)+48)) = v1552
	v1564 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1535+v401))))
	v1569 = v33 + int32(96) + v1535<<(uint(int32(4))%32)
	v1572 = v33 - int32(-64) + v1535
	v1573 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1572))))
	v1574 = F_gistpenalty(m, l4, v1535, v33+int32(48), v1564, v1569, v1573)
	mBase = m.M
	v1575 = m.ExcPending
	if v1575 != 0 {
		goto L1
	} else {
		goto L232
	}
L231:
	;
	goto L142
L232:
	;
	v1577 = *(*int32)(unsafe.Add(mBase, uint32(v1550+v422)))
	v1578 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+62)) = uint8(v1578)
	*(*uint16)(unsafe.Add(mBase, uint32(v33)+60)) = uint16(v1578)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+56)) = v1578
	*(*int32)(unsafe.Add(mBase, uint32(v33)+52)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v33)+48)) = v1577
	v1589 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1535+v408))))
	v1590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1572))))
	v1591 = F_gistpenalty(m, l4, v1535, v33+int32(48), v1589, v1569, v1590)
	mBase = m.M
	v1592 = m.ExcPending
	if v1592 != 0 {
		goto L1
	} else {
		goto L233
	}
L233:
	;
	if base.F32_ne(v1591, v1574) != 0 {
		goto L234
	} else {
		goto L235
	}
L234:
	;
	if base.F32_gt(v1574, v1591) == int32(0) {
		goto L142
	} else {
		goto L237
	}
L235:
	;
	goto L236
L236:
	;
	v1599 = v1535 + int32(1)
	v1600 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v1601 = *(*int32)(unsafe.Add(mBase, uint32(v1600)))
	if v1599 < v1601 {
		v1535 = v1599
		goto L230
	} else {
		goto L238
	}
L237:
	;
	v2053 = l5 + int32(16)
	v2073 = int32(20)
	goto L140
L238:
	;
	goto L231
L239:
	;
	F_gistSplitByKey(m, l0, l1, l2, l3, l4, l5, v878)
	mBase = m.M
	v1637 = m.ExcPending
	if v1637 != 0 {
		goto L1
	} else {
		goto L242
	}
L240:
	;
	goto L241
L241:
	;
	v1641 = F_palloc(m, l3<<(uint(int32(2))%32))
	mBase = m.M
	v1642 = m.ExcPending
	if v1642 != 0 {
		goto L1
	} else {
		goto L243
	}
L242:
	;
	goto L6
L243:
	;
	v1643 = F_palloc(m, v45)
	mBase = m.M
	v1644 = m.ExcPending
	if v1644 != 0 {
		goto L1
	} else {
		goto L244
	}
L244:
	;
	v1645 = int32(0)
	if l3 <= v1645 {
		v1797 = v1645
		goto L245
	} else {
		goto L246
	}
L245:
	;
	v1817 = *(*int64)(unsafe.Add(mBase, uint32(l5)+24))
	v1818 = *(*int32)(unsafe.Add(mBase, uint32(l5)+20))
	v1819 = *(*int64)(unsafe.Add(mBase, uint32(l5)+8))
	v1820 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	v1821 = F_palloc(m, v45)
	mBase = m.M
	v1822 = m.ExcPending
	if v1822 != 0 {
		goto L1
	} else {
		goto L262
	}
L246:
	;
	v1648 = int32(1)
	if l3 == v1648 {
		goto L248
	} else {
		goto L249
	}
L247:
	;
	if l3&v1648 == int32(0) {
		v1797 = v1744
		goto L245
	} else {
		goto L260
	}
L248:
	;
	v1741 = int32(0)
	v1744 = v1645
	goto L247
L249:
	;
	goto L250
L250:
	;
	v1655 = int32(0)
	v1664 = v1655
	v1666 = v1655
	v1667 = v1645
	goto L251
L251:
	;
	v1687 = int32(1)
	v1688 = v1664 | v1687
	v1689 = *(*int32)(unsafe.Add(mBase, uint32(l5)+352))
	v1691 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1688+v1689))))
	if v1691 == v1687 {
		goto L253
	} else {
		goto L254
	}
L252:
	;
	v1741 = v1710
	v1744 = v1730
	goto L247
L253:
	;
	v1694 = int32(2)
	v1700 = *(*int32)(unsafe.Add(mBase, uint32(l2+v1664<<(uint(v1694)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1641+v1667<<(uint(v1694)%32)))) = v1700
	v1702 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1643+v1667<<(uint(v1702)%32)))) = uint16(v1688)
	v1708 = v1667 + v1702
	goto L255
L254:
	;
	v1708 = v1667
	goto L255
L255:
	;
	v1710 = v1664 + int32(2)
	v1711 = *(*int32)(unsafe.Add(mBase, uint32(l5)+352))
	v1713 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1710+v1711))))
	if v1713 == int32(1) {
		goto L256
	} else {
		goto L257
	}
L256:
	;
	v1716 = int32(2)
	v1722 = *(*int32)(unsafe.Add(mBase, uint32(l2+v1688<<(uint(v1716)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1641+v1708<<(uint(v1716)%32)))) = v1722
	v1724 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1643+v1708<<(uint(v1724)%32)))) = uint16(v1710)
	v1730 = v1708 + v1724
	goto L258
L257:
	;
	v1730 = v1708
	goto L258
L258:
	;
	v1732 = v1666 + int32(2)
	if v1732 != l3&int32(2147483646) {
		v1664 = v1710
		v1666 = v1732
		v1667 = v1730
		goto L251
	} else {
		goto L259
	}
L259:
	;
	goto L252
L260:
	;
	v1766 = int32(1)
	v1767 = v1741 + v1766
	v1768 = *(*int32)(unsafe.Add(mBase, uint32(l5)+352))
	v1770 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1767+v1768))))
	if v1770 != v1766 {
		v1797 = v1744
		goto L245
	} else {
		goto L261
	}
L261:
	;
	v1773 = int32(2)
	v1779 = *(*int32)(unsafe.Add(mBase, uint32(l2+v1741<<(uint(v1773)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1641+v1744<<(uint(v1773)%32)))) = v1779
	v1781 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1643+v1744<<(uint(v1781)%32)))) = uint16(v1767)
	v1797 = v1744 + v1781
	goto L245
L262:
	;
	v1823 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v1824 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	v1826 = v1824 << (uint(int32(1)) % 32)
	if v1826 != 0 {
		goto L264
	} else {
		goto L265
	}
L263:
	;
	v1829 = F_palloc(m, v45)
	mBase = m.M
	v1830 = m.ExcPending
	if v1830 != 0 {
		goto L1
	} else {
		goto L267
	}
L264:
	;
	v1827 = F__emscripten_memcpy_bulkmem(m, v1821, v1823, v1826)
	mBase = m.M
	v1828 = v1827
	goto L266
L265:
	;
	v1828 = v1821
	goto L266
L266:
	;
	goto L263
L267:
	;
	v1831 = *(*int32)(unsafe.Add(mBase, uint32(l5)+16))
	v1832 = *(*int32)(unsafe.Add(mBase, uint32(l5)+20))
	v1834 = v1832 << (uint(int32(1)) % 32)
	if v1834 != 0 {
		goto L269
	} else {
		goto L270
	}
L268:
	;
	F_gistSplitByKey(m, l0, l1, v1641, v1797, l4, l5, v878)
	mBase = m.M
	v1838 = m.ExcPending
	if v1838 != 0 {
		goto L1
	} else {
		goto L272
	}
L269:
	;
	v1835 = F__emscripten_memcpy_bulkmem(m, v1829, v1831, v1834)
	mBase = m.M
	v1836 = v1835
	goto L271
L270:
	;
	v1836 = v1829
	goto L271
L271:
	;
	goto L268
L272:
	;
	v1840 = v1643 - int32(2)
	v1841 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if int32(0) < v1841 {
		goto L273
	} else {
		goto L274
	}
L273:
	;
	v1853 = v1820
	v1855 = int32(0)
	goto L276
L274:
	;
	v1902 = v1820
	goto L275
L275:
	;
	v1923 = int32(0)
	v1924 = *(*int32)(unsafe.Add(mBase, uint32(l5)+20))
	if v1923 < v1924 {
		goto L279
	} else {
		goto L280
	}
L276:
	;
	v1874 = int32(1)
	v1877 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v1881 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1877+v1855<<(uint(v1874)%32)))))
	v1885 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1840+v1881<<(uint(v1874)%32)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1828+v1853<<(uint(v1874)%32)))) = uint16(v1885)
	v1888 = v1853 + v1874
	v1890 = v1855 + v1874
	v1891 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v1890 < v1891 {
		v1853 = v1888
		v1855 = v1890
		goto L276
	} else {
		goto L278
	}
L277:
	;
	v1902 = v1888
	goto L275
L278:
	;
	goto L277
L279:
	;
	v1934 = v1818
	v1938 = v1923
	goto L282
L280:
	;
	v1983 = v1818
	goto L281
L281:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l5)+24)) = v1817
	*(*int32)(unsafe.Add(mBase, uint32(l5)+20)) = v1983
	*(*int32)(unsafe.Add(mBase, uint32(l5)+16)) = v1836
	*(*int64)(unsafe.Add(mBase, uint32(l5)+8)) = v1819
	*(*int32)(unsafe.Add(mBase, uint32(l5)+4)) = v1902
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v1828
	goto L6
L282:
	;
	v1957 = int32(1)
	v1960 = *(*int32)(unsafe.Add(mBase, uint32(l5)+16))
	v1964 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1960+v1938<<(uint(v1957)%32)))))
	v1968 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1840+v1964<<(uint(v1957)%32)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1836+v1934<<(uint(v1957)%32)))) = uint16(v1968)
	v1971 = v1934 + v1957
	v1973 = v1938 + v1957
	v1974 = *(*int32)(unsafe.Add(mBase, uint32(l5)+20))
	if v1973 < v1974 {
		v1934 = v1971
		v1938 = v1973
		goto L282
	} else {
		goto L284
	}
L283:
	;
	v1983 = v1971
	goto L281
L284:
	;
	goto L283
L285:
	;
	goto L5
L286:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5)+352)) = int32(0)
	F_gistunionsubkey(m, l4, l2, l5)
	mBase = m.M
	v2151 = m.ExcPending
	if v2151 != 0 {
		goto L1
	} else {
		goto L287
	}
L287:
	;
	goto L4
}
func F_gist_bbox_zorder_abbrev_convert(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 float64
	_ = v6
	var v8 int64
	_ = v8
	var v10 float64
	_ = v10
	var v12 int32
	_ = v12
	var v21 int32
	_ = v21
	var v24 int64
	_ = v24
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v37 int64
	_ = v37
	var v38 int64
	_ = v38
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v43 int64
	_ = v43
	var v46 int64
	_ = v46
	var v47 int64
	_ = v47
	var v48 int64
	_ = v48
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v53 int64
	_ = v53
	var v56 int64
	_ = v56
	var v57 int64
	_ = v57
	var v60 int64
	_ = v60
	var v69 int64
	_ = v69
	var v74 int64
	_ = v74
	var v79 int64
	_ = v79
	var v84 int64
	_ = v84
	v6 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = int64(4294967295)
	v10 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
	v12 = base.I32_reinterpret_f32(base.F32_demote_f64(v10))
	if base.Ui32(v12&int32(2147483647)) <= base.Ui32(int32(2139095040)) {
		if v12 < int32(0) {
			v21 = int32(-1)
		} else {
			v21 = int32(-2147483648)
		}
		v24 = base.I64_extend_i32_u(v12 ^ v21)
	} else {
		v24 = v8
	}
	v25 = base.I32_reinterpret_f32(base.F32_demote_f64(v6))
	if base.Ui32(v25&int32(2147483647)) <= base.Ui32(int32(2139095040)) {
		if v25 < int32(0) {
			v34 = int32(-1)
		} else {
			v34 = int32(-2147483648)
		}
		v37 = base.I64_extend_i32_u(v25 ^ v34)
	} else {
		v37 = v8
	}
	v38 = int64(16)
	v41 = int64(281470681808895)
	v42 = (v37<<(uint(v38)%64) | v37) & v41
	v43 = int64(8)
	v46 = int64(71777214294589695)
	v47 = (v42<<(uint(v43)%64) | v42) & v46
	v48 = int64(4)
	v51 = int64(1085102592571150095)
	v52 = (v47<<(uint(v48)%64) | v47) & v51
	v53 = int64(2)
	v56 = int64(3689348814741910323)
	v57 = (v52<<(uint(v53)%64) | v52) & v56
	v60 = int64(1)
	v69 = (v24<<(uint(v38)%64) | v24) & v41
	v74 = (v69<<(uint(v43)%64) | v69) & v46
	v79 = (v74<<(uint(v48)%64) | v74) & v51
	v84 = (v79<<(uint(v53)%64) | v79) & v56
	return base.I32_wrap_i64(int64(base.Ui64((v57<<(uint(v53)%64)|v57<<(uint(v60)%64))&int64(-6148914694099828736)|(v84<<(uint(v60)%64)|v84)&int64(6148914689804861440)) >> (uint(int64(32)) % 64)))
}
func F_gist_box_penalty(m *base.Module, l0 int32) int32 {
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
	var v7 int32
	_ = v7
	var v8 float64
	_ = v8
	var v11 int32
	_ = v11
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v8 = F_box_penalty(m, v5, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		*(*float32)(unsafe.Add(mBase, uint32(v3))) = base.F32_demote_f64(v8)
		return v3
	}
}
func F_gist_box_picksplit(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
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
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v93 int32
	_ = v93
	var v96 int64
	_ = v96
	var v98 int64
	_ = v98
	var v100 int64
	_ = v100
	var v102 int64
	_ = v102
	var v104 float64
	_ = v104
	var v110 float64
	_ = v110
	var v121 float64
	_ = v121
	var v127 float64
	_ = v127
	var v133 float64
	_ = v133
	var v135 float64
	_ = v135
	var v138 float64
	_ = v138
	var v144 float64
	_ = v144
	var v155 float64
	_ = v155
	var v161 float64
	_ = v161
	var v167 float64
	_ = v167
	var v169 float64
	_ = v169
	var v176 int32
	_ = v176
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v243 int32
	_ = v243
	var v251 int32
	_ = v251
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v281 float64
	_ = v281
	var v283 float64
	_ = v283
	var v287 int32
	_ = v287
	var v290 float64
	_ = v290
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v307 float64
	_ = v307
	var v308 float64
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 float64
	_ = v312
	var v313 float64
	_ = v313
	var v321 int32
	_ = v321
	var v340 int32
	_ = v340
	var v341 float64
	_ = v341
	var v368 int32
	_ = v368
	var v369 float64
	_ = v369
	var v372 int64
	_ = v372
	var v385 float64
	_ = v385
	var v391 float64
	_ = v391
	var v393 float64
	_ = v393
	var v394 float64
	_ = v394
	var v397 int32
	_ = v397
	var v413 int32
	_ = v413
	var v434 float64
	_ = v434
	var v445 int32
	_ = v445
	var v457 int32
	_ = v457
	var v476 int32
	_ = v476
	var v504 float64
	_ = v504
	var v505 float64
	_ = v505
	var v506 int32
	_ = v506
	var v507 float64
	_ = v507
	var v508 float64
	_ = v508
	var v520 int32
	_ = v520
	var v535 int32
	_ = v535
	var v537 float64
	_ = v537
	var v563 int32
	_ = v563
	var v564 float64
	_ = v564
	var v567 int64
	_ = v567
	var v575 float64
	_ = v575
	var v582 float64
	_ = v582
	var v588 float64
	_ = v588
	var v589 float64
	_ = v589
	var v604 int32
	_ = v604
	var v628 float64
	_ = v628
	var v657 int32
	_ = v657
	var v671 int32
	_ = v671
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v684 int32
	_ = v684
	var v688 int32
	_ = v688
	var v715 int32
	_ = v715
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v726 int32
	_ = v726
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v747 int32
	_ = v747
	var v751 int32
	_ = v751
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v760 int32
	_ = v760
	var v780 int32
	_ = v780
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v793 int64
	_ = v793
	var v795 int64
	_ = v795
	var v797 int64
	_ = v797
	var v799 int64
	_ = v799
	var v801 float64
	_ = v801
	var v807 float64
	_ = v807
	var v818 float64
	_ = v818
	var v824 float64
	_ = v824
	var v835 float64
	_ = v835
	var v841 float64
	_ = v841
	var v852 float64
	_ = v852
	var v858 float64
	_ = v858
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v886 int64
	_ = v886
	var v888 int64
	_ = v888
	var v890 int64
	_ = v890
	var v892 int64
	_ = v892
	var v894 float64
	_ = v894
	var v900 float64
	_ = v900
	var v911 float64
	_ = v911
	var v917 float64
	_ = v917
	var v928 float64
	_ = v928
	var v934 float64
	_ = v934
	var v945 float64
	_ = v945
	var v951 float64
	_ = v951
	var v961 int32
	_ = v961
	var v964 int32
	_ = v964
	var v968 int32
	_ = v968
	var v971 int32
	_ = v971
	var v973 int32
	_ = v973
	var v975 int32
	_ = v975
	var v977 int32
	_ = v977
	var v986 int32
	_ = v986
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1026 float64
	_ = v1026
	var v1028 int64
	_ = v1028
	var v1030 float64
	_ = v1030
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1053 int32
	_ = v1053
	var v1070 int32
	_ = v1070
	var v1072 float64
	_ = v1072
	var v1075 int32
	_ = v1075
	var v1080 float64
	_ = v1080
	var v1107 int32
	_ = v1107
	var v1110 float64
	_ = v1110
	var v1116 float64
	_ = v1116
	var v1127 float64
	_ = v1127
	var v1133 float64
	_ = v1133
	var v1144 float64
	_ = v1144
	var v1150 float64
	_ = v1150
	var v1161 float64
	_ = v1161
	var v1167 float64
	_ = v1167
	var v1177 int64
	_ = v1177
	var v1179 int64
	_ = v1179
	var v1181 int64
	_ = v1181
	var v1183 int64
	_ = v1183
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1191 int32
	_ = v1191
	var v1196 int32
	_ = v1196
	var v1199 float64
	_ = v1199
	var v1205 float64
	_ = v1205
	var v1216 float64
	_ = v1216
	var v1222 float64
	_ = v1222
	var v1233 float64
	_ = v1233
	var v1239 float64
	_ = v1239
	var v1250 float64
	_ = v1250
	var v1256 float64
	_ = v1256
	var v1266 int64
	_ = v1266
	var v1268 int64
	_ = v1268
	var v1270 int64
	_ = v1270
	var v1272 int64
	_ = v1272
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1280 int32
	_ = v1280
	var v1288 int32
	_ = v1288
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1301 float64
	_ = v1301
	var v1305 int32
	_ = v1305
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1316 int32
	_ = v1316
	var v1317 int32
	_ = v1317
	var v1336 int32
	_ = v1336
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1343 int32
	_ = v1343
	var v1344 float64
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1346 float64
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1349 float64
	_ = v1349
	var v1360 int32
	_ = v1360
	var v1362 int32
	_ = v1362
	var v1367 int32
	_ = v1367
	var v1368 int32
	_ = v1368
	var v1377 int32
	_ = v1377
	var v1380 int32
	_ = v1380
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
	var v1405 int32
	_ = v1405
	var v1410 float64
	_ = v1410
	var v1416 float64
	_ = v1416
	var v1427 float64
	_ = v1427
	var v1433 float64
	_ = v1433
	var v1444 float64
	_ = v1444
	var v1450 float64
	_ = v1450
	var v1461 float64
	_ = v1461
	var v1467 float64
	_ = v1467
	var v1477 int64
	_ = v1477
	var v1479 int64
	_ = v1479
	var v1481 int64
	_ = v1481
	var v1483 int64
	_ = v1483
	var v1487 int32
	_ = v1487
	var v1488 int32
	_ = v1488
	var v1489 int32
	_ = v1489
	var v1492 int32
	_ = v1492
	var v1497 int32
	_ = v1497
	var v1502 float64
	_ = v1502
	var v1508 float64
	_ = v1508
	var v1519 float64
	_ = v1519
	var v1525 float64
	_ = v1525
	var v1536 float64
	_ = v1536
	var v1542 float64
	_ = v1542
	var v1553 float64
	_ = v1553
	var v1559 float64
	_ = v1559
	var v1569 int64
	_ = v1569
	var v1571 int64
	_ = v1571
	var v1573 int64
	_ = v1573
	var v1575 int64
	_ = v1575
	var v1579 int32
	_ = v1579
	var v1580 int32
	_ = v1580
	var v1581 int32
	_ = v1581
	var v1584 int32
	_ = v1584
	var v1589 float64
	_ = v1589
	var v1590 int32
	_ = v1590
	var v1591 float64
	_ = v1591
	var v1592 int32
	_ = v1592
	var v1594 int32
	_ = v1594
	var v1597 float64
	_ = v1597
	var v1603 float64
	_ = v1603
	var v1614 float64
	_ = v1614
	var v1620 float64
	_ = v1620
	var v1631 float64
	_ = v1631
	var v1637 float64
	_ = v1637
	var v1648 float64
	_ = v1648
	var v1654 float64
	_ = v1654
	var v1664 int64
	_ = v1664
	var v1666 int64
	_ = v1666
	var v1668 int64
	_ = v1668
	var v1670 int64
	_ = v1670
	var v1674 int32
	_ = v1674
	var v1675 int32
	_ = v1675
	var v1676 int32
	_ = v1676
	var v1679 int32
	_ = v1679
	var v1684 int32
	_ = v1684
	var v1687 float64
	_ = v1687
	var v1693 float64
	_ = v1693
	var v1704 float64
	_ = v1704
	var v1710 float64
	_ = v1710
	var v1721 float64
	_ = v1721
	var v1727 float64
	_ = v1727
	var v1738 float64
	_ = v1738
	var v1744 float64
	_ = v1744
	var v1754 int64
	_ = v1754
	var v1756 int64
	_ = v1756
	var v1758 int64
	_ = v1758
	var v1760 int64
	_ = v1760
	var v1764 int32
	_ = v1764
	var v1765 int32
	_ = v1765
	var v1766 int32
	_ = v1766
	var v1769 int32
	_ = v1769
	var v1780 int32
	_ = v1780
	var v1782 int32
	_ = v1782
	var v1811 int32
	_ = v1811
	var v1843 int32
	_ = v1843
	v6 = int32(0)
	v27 = m.G0
	v29 = v27 - int32(96)
	m.G0 = v29
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v38 = F__emscripten_memset_bulkmem(m, v29+int32(8), base.I32_extend8_s(v6), int32(88))
	mBase = m.M
	goto L1
L1:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v40 = int32(65535)
	v41 = v39 + v40
	v43 = v41 & v40
	*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = v43
	v46 = v43 - int32(1)
	v48 = v43 << (uint(int32(4)) % 32)
	v49 = F_palloc(m, v48)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	return int32(0)
L3:
	;
	v53 = F_palloc(m, v48)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v56 = v39 & int32(65535)
	if v56 != int32(1) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v62 = v29 + int32(16)
	v69 = int32(1)
	goto L8
L6:
	;
	goto L7
L7:
	;
	v204 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+48)) = uint8(v204)
	v206 = int32(4)
	v207 = v32 + v206
	v209 = v46 << (uint(v206) % 32)
	v211 = int32(8)
	v229 = v204
	v232 = v6
	goto L43
L8:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v32+int32(4)+v69<<(uint(int32(4))%32))))
	if v69 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L7
L10:
	;
	v176 = (v69 + int32(1)) & int32(65535)
	if base.Ui32(v176) <= base.Ui32(v43) {
		v69 = v176
		goto L8
	} else {
		goto L42
	}
L11:
	;
	v96 = *(*int64)(unsafe.Add(mBase, uint32(v93)))
	*(*int64)(unsafe.Add(mBase, uint32(v62))) = v96
	v98 = *(*int64)(unsafe.Add(mBase, uint32(v93)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v62)+24)) = v98
	v100 = *(*int64)(unsafe.Add(mBase, uint32(v93)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v62)+16)) = v100
	v102 = *(*int64)(unsafe.Add(mBase, uint32(v93)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v62)+8)) = v102
	goto L10
L12:
	;
	goto L13
L13:
	;
	v104 = *(*float64)(unsafe.Add(mBase, uint32(v29)+16))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v104)&int64(9223372036854775807)) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v121 = *(*float64)(unsafe.Add(mBase, uint32(v93)+16))
	if base.Ui64(base.I64_reinterpret_f64(v121)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L20
	} else {
		goto L21
	}
L15:
	;
	v110 = *(*float64)(unsafe.Add(mBase, uint32(v93)))
	if base.F64_lt(v104, v110) == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	if base.Ui64(base.I64_reinterpret_f64(v110)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L14
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v29)+16)) = v110
	goto L14
L19:
	;
	goto L18
L20:
	;
	v127 = *(*float64)(unsafe.Add(mBase, uint32(v29)+32))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v127)&int64(9223372036854775807)) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	v138 = *(*float64)(unsafe.Add(mBase, uint32(v29)+24))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v138)&int64(9223372036854775807)) {
		goto L29
	} else {
		goto L30
	}
L23:
	;
	v133 = v121
	goto L25
L24:
	;
	v133 = v127
	goto L25
L25:
	;
	if base.F64_lt(v121, v127) != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v135 = v121
	goto L28
L27:
	;
	v135 = v133
	goto L28
L28:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v29)+32)) = v135
	goto L22
L29:
	;
	v155 = *(*float64)(unsafe.Add(mBase, uint32(v93)+24))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v155)&int64(9223372036854775807)) {
		goto L10
	} else {
		goto L35
	}
L30:
	;
	v144 = *(*float64)(unsafe.Add(mBase, uint32(v93)+8))
	if base.F64_lt(v138, v144) == int32(0) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	if base.Ui64(base.I64_reinterpret_f64(v144)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L29
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v29)+24)) = v144
	goto L29
L34:
	;
	goto L33
L35:
	;
	v161 = *(*float64)(unsafe.Add(mBase, uint32(v29)+40))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v161)&int64(9223372036854775807)) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v167 = v155
	goto L38
L37:
	;
	v167 = v161
	goto L38
L38:
	;
	if base.F64_lt(v155, v161) != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v169 = v155
	goto L41
L40:
	;
	v169 = v167
	goto L41
L41:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v29)+40)) = v169
	goto L10
L42:
	;
	goto L9
L43:
	;
	v243 = int32(1)
	if v56 != v243 {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	v719 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+48)))
	if v719 == int32(1) {
		goto L138
	} else {
		goto L139
	}
L45:
	;
	v715 = int32(1)
	if v229&v715 != 0 {
		v229 = int32(0)
		v232 = v715
		goto L43
	} else {
		goto L135
	}
L46:
	;
	v251 = v243
	goto L49
L47:
	;
	goto L48
L48:
	;
	if v48 != 0 {
		goto L130
	} else {
		goto L131
	}
L49:
	;
	v273 = v251 << (uint(int32(4)) % 32)
	v274 = v49 + v273
	v276 = v274 - int32(16)
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v273+v207)))
	if v229&int32(1) != 0 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	if v48 != 0 {
		goto L57
	} else {
		goto L58
	}
L51:
	;
	v290 = *(*float64)(unsafe.Add(mBase, uint32(v287)))
	*(*float64)(unsafe.Add(mBase, uint32(v274-int32(8)))) = v290
	v295 = (v251 + int32(1)) & int32(65535)
	if base.Ui32(v295) <= base.Ui32(v43) {
		v251 = v295
		goto L49
	} else {
		goto L55
	}
L52:
	;
	v281 = *(*float64)(unsafe.Add(mBase, uint32(v278)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v276))) = v281
	v287 = v278
	goto L51
L53:
	;
	goto L54
L54:
	;
	v283 = *(*float64)(unsafe.Add(mBase, uint32(v278)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v276))) = v283
	v287 = v278 + int32(8)
	goto L51
L55:
	;
	goto L50
L56:
	;
	F_pg_qsort(m, v49, v43, int32(16), int32(105))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L2
	} else {
		goto L60
	}
L57:
	;
	v297 = F__emscripten_memcpy_bulkmem(m, v53, v49, v48)
	mBase = m.M
	v298 = v297
	goto L59
L58:
	;
	v298 = v53
	goto L59
L59:
	;
	goto L56
L60:
	;
	F_pg_qsort(m, v298, v43, int32(16), int32(106))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L2
	} else {
		goto L61
	}
L61:
	;
	v307 = *(*float64)(unsafe.Add(mBase, uint32(v49)))
	v308 = *(*float64)(unsafe.Add(mBase, uint32(v298)))
	v309 = int32(0)
	v311 = v309
	v312 = v308
	v313 = v307
	v321 = v309
	goto L62
L62:
	;
	v340 = v311
	v341 = v312
	goto L65
L63:
	;
	v504 = *(*float64)(unsafe.Add(mBase, uint32(v209+v49+v211)))
	v505 = *(*float64)(unsafe.Add(mBase, uint32(v53+v209+v211)))
	v506 = v46
	v507 = v505
	v508 = v504
	v520 = v46
	goto L97
L64:
	;
	goto L63
L65:
	;
	v368 = v49 + v340<<(uint(int32(4))%32)
	v369 = *(*float64)(unsafe.Add(mBase, uint32(v368)))
	v372 = base.I64_reinterpret_f64(v369) & int64(9223372036854775807)
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v313)&int64(9223372036854775807)) {
		goto L69
	} else {
		goto L70
	}
L66:
	;
	if v43 <= v321 {
		v457 = v321
		goto L85
	} else {
		goto L86
	}
L67:
	;
	goto L66
L68:
	;
	if base.Ui64(base.I64_reinterpret_f64(v341)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L75
	} else {
		goto L76
	}
L69:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v372) {
		goto L68
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	if base.F64_ne(v369, v313) != 0 {
		goto L67
	} else {
		goto L73
	}
L72:
	;
	goto L67
L73:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v372) {
		goto L67
	} else {
		goto L74
	}
L74:
	;
	goto L68
L75:
	;
	v385 = *(*float64)(unsafe.Add(mBase, uint32(v368)+8))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v385)&int64(9223372036854775807)) {
		goto L78
	} else {
		goto L79
	}
L76:
	;
	v394 = v341
	goto L77
L77:
	;
	v397 = v340 + int32(1)
	if v397 < v43 {
		v340 = v397
		v341 = v394
		goto L65
	} else {
		goto L84
	}
L78:
	;
	v391 = v385
	goto L80
L79:
	;
	v391 = v341
	goto L80
L80:
	;
	if base.F64_lt(v341, v385) != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v393 = v385
	goto L83
L82:
	;
	v393 = v391
	goto L83
L83:
	;
	v394 = v393
	goto L77
L84:
	;
	goto L64
L85:
	;
	F_g_box_consider_split(m, v29+int32(8), v232, v369, v340, v341, v457)
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L2
	} else {
		goto L95
	}
L86:
	;
	v413 = v321
	goto L87
L87:
	;
	if base.Ui64(base.I64_reinterpret_f64(v341)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v457 = v43
	goto L85
L89:
	;
	v434 = *(*float64)(unsafe.Add(mBase, uint32(v298+v413<<(uint(int32(4))%32))+8))
	if base.F64_le(v434, v341) == int32(0) {
		v457 = v413
		goto L85
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v445 = v413 + int32(1)
	if v445 != v43 {
		v413 = v445
		goto L87
	} else {
		goto L94
	}
L92:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v434)&int64(9223372036854775807)) {
		v457 = v413
		goto L85
	} else {
		goto L93
	}
L93:
	;
	goto L91
L94:
	;
	goto L88
L95:
	;
	if v340 < v43 {
		v311 = v340
		v312 = v341
		v313 = v369
		v321 = v457
		goto L62
	} else {
		goto L96
	}
L96:
	;
	goto L64
L97:
	;
	v535 = v506
	v537 = v508
	goto L99
L98:
	;
	goto L45
L99:
	;
	v563 = v298 + v535<<(uint(int32(4))%32)
	v564 = *(*float64)(unsafe.Add(mBase, uint32(v563)+8))
	v567 = base.I64_reinterpret_f64(v564) & int64(9223372036854775807)
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v507)&int64(9223372036854775807)) {
		goto L103
	} else {
		goto L104
	}
L100:
	;
	if v520 < int32(0) {
		v657 = v520
		goto L119
	} else {
		goto L120
	}
L101:
	;
	goto L100
L102:
	;
	v575 = *(*float64)(unsafe.Add(mBase, uint32(v563)))
	if base.Ui64(base.I64_reinterpret_f64(v575)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L109
	} else {
		goto L110
	}
L103:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v567) {
		goto L102
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	if base.F64_ne(v564, v507) != 0 {
		goto L101
	} else {
		goto L107
	}
L106:
	;
	goto L101
L107:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v567) {
		goto L101
	} else {
		goto L108
	}
L108:
	;
	goto L102
L109:
	;
	if base.F64_lt(v575, v537) != 0 {
		goto L112
	} else {
		goto L113
	}
L110:
	;
	v589 = v537
	goto L111
L111:
	;
	if int32(0) < v535 {
		v535 = v535 - int32(1)
		v537 = v589
		goto L99
	} else {
		goto L118
	}
L112:
	;
	v582 = v575
	goto L114
L113:
	;
	v582 = v537
	goto L114
L114:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v537)&int64(9223372036854775807)) {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v588 = v575
	goto L117
L116:
	;
	v588 = v582
	goto L117
L117:
	;
	v589 = v588
	goto L111
L118:
	;
	goto L45
L119:
	;
	v671 = int32(1)
	F_g_box_consider_split(m, v29+int32(8), v232, v537, v657+v671, v564, v535+v671)
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L2
	} else {
		goto L127
	}
L120:
	;
	v604 = v520
	goto L121
L121:
	;
	v628 = *(*float64)(unsafe.Add(mBase, uint32(v49+v604<<(uint(int32(4))%32))))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v628)&int64(9223372036854775807)) {
		goto L123
	} else {
		goto L124
	}
L122:
	;
	v657 = int32(-1)
	goto L119
L123:
	;
	if int32(0) < v604 {
		v604 = v604 - int32(1)
		goto L121
	} else {
		goto L126
	}
L124:
	;
	if base.B2i32(base.Ui64(base.I64_reinterpret_f64(v537)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)))&base.F64_le(v537, v628) != 0 {
		goto L123
	} else {
		goto L125
	}
L125:
	;
	v657 = v604
	goto L119
L126:
	;
	goto L122
L127:
	;
	if int32(0) <= v535 {
		v506 = v535
		v507 = v564
		v508 = v537
		v520 = v657
		goto L97
	} else {
		goto L128
	}
L128:
	;
	goto L98
L129:
	;
	F_pg_qsort(m, v49, v43, int32(16), int32(105))
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L2
	} else {
		goto L133
	}
L130:
	;
	v679 = F__emscripten_memcpy_bulkmem(m, v53, v49, v48)
	mBase = m.M
	v680 = v679
	goto L132
L131:
	;
	v680 = v53
	goto L132
L132:
	;
	goto L129
L133:
	;
	F_pg_qsort(m, v680, v43, int32(16), int32(106))
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L2
	} else {
		goto L134
	}
L134:
	;
	goto L45
L135:
	;
	goto L44
L136:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v1843 = m.ExcPending
	if v1843 != 0 {
		goto L2
	} else {
		goto L426
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+24)) = v1811
	m.G0 = v29 + int32(96)
	return v31
L138:
	;
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v723 = int32(65535)
	v726 = (v722 + v723) & v723
	v730 = v726<<(uint(int32(1))%32) + int32(4)
	v731 = F_palloc(m, v730)
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L2
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	v1005 = v43 << (uint(int32(1)) % 32)
	v1006 = F_palloc(m, v1005)
	mBase = m.M
	v1007 = m.ExcPending
	if v1007 != 0 {
		goto L2
	} else {
		goto L209
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = v731
	v734 = int32(0)
	v735 = F_palloc(m, v730)
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L2
	} else {
		goto L142
	}
L142:
	;
	v737 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = v737
	*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = v735
	*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v737
	if v722&int32(65535) != int32(1) {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v747 = int32(1)
	v751 = v734
	v757 = v747
	v758 = v747
	v760 = v737
	goto L146
L144:
	;
	v977 = v734
	v986 = v737
	goto L145
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = v986
	v1811 = v977
	goto L137
L146:
	;
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v207+v757<<(uint(int32(4))%32))))
	if base.Ui32(v757) <= base.Ui32(int32(base.Ui32(v726)>>(uint(v747)%32))) {
		goto L149
	} else {
		goto L150
	}
L147:
	;
	v977 = v968
	v986 = v971
	goto L145
L148:
	;
	v973 = v758 + int32(1)
	v975 = v973 & int32(65535)
	if base.Ui32(v975) <= base.Ui32(v726) {
		v751 = v968
		v757 = v975
		v758 = v973
		v760 = v971
		goto L146
	} else {
		goto L208
	}
L149:
	;
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v783 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	*(*uint16)(unsafe.Add(mBase, uint32(v782+v783<<(uint(int32(1))%32)))) = uint16(v758)
	if v760 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L150:
	;
	goto L151
L151:
	;
	v875 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
	*(*uint16)(unsafe.Add(mBase, uint32(v875+v876<<(uint(int32(1))%32)))) = uint16(v758)
	if v751 == int32(0) {
		goto L181
	} else {
		goto L182
	}
L152:
	;
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v871 + int32(1)
	v968 = v751
	v971 = v870
	goto L148
L153:
	;
	v791 = F_palloc(m, int32(32))
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L2
	} else {
		goto L156
	}
L154:
	;
	goto L155
L155:
	;
	v801 = *(*float64)(unsafe.Add(mBase, uint32(v760)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v801)&int64(9223372036854775807)) {
		goto L157
	} else {
		goto L158
	}
L156:
	;
	v793 = *(*int64)(unsafe.Add(mBase, uint32(v780)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v791)+24)) = v793
	v795 = *(*int64)(unsafe.Add(mBase, uint32(v780)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v791)+16)) = v795
	v797 = *(*int64)(unsafe.Add(mBase, uint32(v780)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v791)+8)) = v797
	v799 = *(*int64)(unsafe.Add(mBase, uint32(v780)))
	*(*int64)(unsafe.Add(mBase, uint32(v791))) = v799
	v870 = v791
	goto L152
L157:
	;
	v818 = *(*float64)(unsafe.Add(mBase, uint32(v780)+16))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v818)&int64(9223372036854775807)) {
		goto L163
	} else {
		goto L164
	}
L158:
	;
	v807 = *(*float64)(unsafe.Add(mBase, uint32(v780)))
	if base.F64_lt(v801, v807) == int32(0) {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	if base.Ui64(base.I64_reinterpret_f64(v807)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L157
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v760))) = v807
	goto L157
L162:
	;
	goto L161
L163:
	;
	v835 = *(*float64)(unsafe.Add(mBase, uint32(v760)+8))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v835)&int64(9223372036854775807)) {
		goto L169
	} else {
		goto L170
	}
L164:
	;
	v824 = *(*float64)(unsafe.Add(mBase, uint32(v760)+16))
	if base.F64_gt(v824, v818) == int32(0) {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	if base.Ui64(base.I64_reinterpret_f64(v824)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L163
	} else {
		goto L168
	}
L166:
	;
	goto L167
L167:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v760)+16)) = v818
	goto L163
L168:
	;
	goto L167
L169:
	;
	v852 = *(*float64)(unsafe.Add(mBase, uint32(v780)+24))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v852)&int64(9223372036854775807)) {
		v870 = v760
		goto L152
	} else {
		goto L175
	}
L170:
	;
	v841 = *(*float64)(unsafe.Add(mBase, uint32(v780)+8))
	if base.F64_lt(v835, v841) == int32(0) {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	if base.Ui64(base.I64_reinterpret_f64(v841)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L169
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v760)+8)) = v841
	goto L169
L174:
	;
	goto L173
L175:
	;
	v858 = *(*float64)(unsafe.Add(mBase, uint32(v760)+24))
	if base.F64_gt(v858, v852) == int32(0) {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	if base.Ui64(base.I64_reinterpret_f64(v858)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		v870 = v760
		goto L152
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v760)+24)) = v852
	v870 = v760
	goto L152
L179:
	;
	goto L178
L180:
	;
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = v964 + int32(1)
	v968 = v961
	v971 = v760
	goto L148
L181:
	;
	v884 = F_palloc(m, int32(32))
	mBase = m.M
	v885 = m.ExcPending
	if v885 != 0 {
		goto L2
	} else {
		goto L184
	}
L182:
	;
	goto L183
L183:
	;
	v894 = *(*float64)(unsafe.Add(mBase, uint32(v751)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v894)&int64(9223372036854775807)) {
		goto L185
	} else {
		goto L186
	}
L184:
	;
	v886 = *(*int64)(unsafe.Add(mBase, uint32(v780)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v884)+24)) = v886
	v888 = *(*int64)(unsafe.Add(mBase, uint32(v780)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v884)+16)) = v888
	v890 = *(*int64)(unsafe.Add(mBase, uint32(v780)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v884)+8)) = v890
	v892 = *(*int64)(unsafe.Add(mBase, uint32(v780)))
	*(*int64)(unsafe.Add(mBase, uint32(v884))) = v892
	v961 = v884
	goto L180
L185:
	;
	v911 = *(*float64)(unsafe.Add(mBase, uint32(v780)+16))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v911)&int64(9223372036854775807)) {
		goto L191
	} else {
		goto L192
	}
L186:
	;
	v900 = *(*float64)(unsafe.Add(mBase, uint32(v780)))
	if base.F64_lt(v894, v900) == int32(0) {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	if base.Ui64(base.I64_reinterpret_f64(v900)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L185
	} else {
		goto L190
	}
L188:
	;
	goto L189
L189:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v751))) = v900
	goto L185
L190:
	;
	goto L189
L191:
	;
	v928 = *(*float64)(unsafe.Add(mBase, uint32(v751)+8))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v928)&int64(9223372036854775807)) {
		goto L197
	} else {
		goto L198
	}
L192:
	;
	v917 = *(*float64)(unsafe.Add(mBase, uint32(v751)+16))
	if base.F64_gt(v917, v911) == int32(0) {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	if base.Ui64(base.I64_reinterpret_f64(v917)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L191
	} else {
		goto L196
	}
L194:
	;
	goto L195
L195:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v751)+16)) = v911
	goto L191
L196:
	;
	goto L195
L197:
	;
	v945 = *(*float64)(unsafe.Add(mBase, uint32(v780)+24))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v945)&int64(9223372036854775807)) {
		v961 = v751
		goto L180
	} else {
		goto L203
	}
L198:
	;
	v934 = *(*float64)(unsafe.Add(mBase, uint32(v780)+8))
	if base.F64_lt(v928, v934) == int32(0) {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	if base.Ui64(base.I64_reinterpret_f64(v934)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L197
	} else {
		goto L202
	}
L200:
	;
	goto L201
L201:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v751)+8)) = v934
	goto L197
L202:
	;
	goto L201
L203:
	;
	v951 = *(*float64)(unsafe.Add(mBase, uint32(v751)+24))
	if base.F64_gt(v951, v945) == int32(0) {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	if base.Ui64(base.I64_reinterpret_f64(v951)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		v961 = v751
		goto L180
	} else {
		goto L207
	}
L205:
	;
	goto L206
L206:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v751)+24)) = v945
	v961 = v751
	goto L180
L207:
	;
	goto L206
L208:
	;
	goto L147
L209:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = v1006
	v1009 = F_palloc(m, v1005)
	mBase = m.M
	v1010 = m.ExcPending
	if v1010 != 0 {
		goto L2
	} else {
		goto L210
	}
L210:
	;
	v1011 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = v1011
	*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v1011
	*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = v1009
	v1017 = F_palloc0(m, int32(32))
	mBase = m.M
	v1018 = m.ExcPending
	if v1018 != 0 {
		goto L2
	} else {
		goto L211
	}
L211:
	;
	v1020 = F_palloc0(m, int32(32))
	mBase = m.M
	v1021 = m.ExcPending
	if v1021 != 0 {
		goto L2
	} else {
		goto L212
	}
L212:
	;
	v1022 = F_palloc(m, v48)
	mBase = m.M
	v1023 = m.ExcPending
	if v1023 != 0 {
		goto L2
	} else {
		goto L213
	}
L213:
	;
	if v56 == int32(1) {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = v1017
	v1811 = v1020
	goto L137
L215:
	;
	v1026 = *(*float64)(unsafe.Add(mBase, uint32(v29)+64))
	v1028 = int64(9223372036854775807)
	v1030 = *(*float64)(unsafe.Add(mBase, uint32(v29)+56))
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(v29)+80))
	if v1036 != 0 {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	v1037 = int32(24)
	goto L218
L217:
	;
	v1037 = int32(16)
	goto L218
L218:
	;
	v1038 = int32(1)
	v1050 = int32(0)
	v1051 = v1038
	v1053 = v1038
	goto L219
L219:
	;
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(v207+v1053<<(uint(int32(4))%32))))
	v1072 = *(*float64)(unsafe.Add(mBase, uint32(v1070+v1037)))
	if base.Ui64(base.I64_reinterpret_f64(v1030)&v1028) <= base.Ui64(int64(9218868437227405312)) {
		goto L223
	} else {
		goto L224
	}
L220:
	;
	if v1288 <= int32(0) {
		goto L214
	} else {
		goto L288
	}
L221:
	;
	v1290 = v1051 + int32(1)
	v1291 = int32(65535)
	v1292 = v1290 & v1291
	if base.Ui32(v1292) <= base.Ui32(v41&v1291) {
		v1050 = v1288
		v1051 = v1290
		v1053 = v1292
		goto L219
	} else {
		goto L287
	}
L222:
	;
	v1196 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
	if int32(0) < v1196 {
		goto L261
	} else {
		goto L262
	}
L223:
	;
	v1075 = int32(0)
	v1080 = *(*float64)(unsafe.Add(mBase, uint32(v1070+base.B2i32(v1036 != v1075)<<(uint(int32(3))%32))))
	if base.F64_le(v1080, v1030) == v1075 {
		goto L222
	} else {
		goto L226
	}
L224:
	;
	goto L225
L225:
	;
	if base.Ui64(base.I64_reinterpret_f64(v1072)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L229
	} else {
		goto L230
	}
L226:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1080)&int64(9223372036854775807)) {
		goto L222
	} else {
		goto L227
	}
L227:
	;
	goto L225
L228:
	;
	v1107 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if int32(0) < v1107 {
		goto L234
	} else {
		goto L235
	}
L229:
	;
	if base.B2i32(base.Ui64(base.I64_reinterpret_f64(v1026)&v1028) < base.Ui64(int64(9218868437227405313)))&base.F64_ge(v1072, v1026) == int32(0) {
		goto L228
	} else {
		goto L232
	}
L230:
	;
	goto L231
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1022+v1050<<(uint(int32(4))%32)))) = v1053
	v1288 = v1050 + int32(1)
	goto L221
L232:
	;
	goto L231
L233:
	;
	v1187 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v1188 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v1187 + v1188
	v1191 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1191+v1187<<(uint(v1188)%32)))) = uint16(v1051)
	v1288 = v1050
	goto L221
L234:
	;
	v1110 = *(*float64)(unsafe.Add(mBase, uint32(v1017)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1110)&int64(9223372036854775807)) {
		goto L237
	} else {
		goto L238
	}
L235:
	;
	goto L236
L236:
	;
	v1177 = *(*int64)(unsafe.Add(mBase, uint32(v1070)))
	*(*int64)(unsafe.Add(mBase, uint32(v1017))) = v1177
	v1179 = *(*int64)(unsafe.Add(mBase, uint32(v1070)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1017)+24)) = v1179
	v1181 = *(*int64)(unsafe.Add(mBase, uint32(v1070)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1017)+16)) = v1181
	v1183 = *(*int64)(unsafe.Add(mBase, uint32(v1070)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1017)+8)) = v1183
	goto L233
L237:
	;
	v1127 = *(*float64)(unsafe.Add(mBase, uint32(v1070)+16))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1127)&int64(9223372036854775807)) {
		goto L243
	} else {
		goto L244
	}
L238:
	;
	v1116 = *(*float64)(unsafe.Add(mBase, uint32(v1070)))
	if base.F64_lt(v1110, v1116) == int32(0) {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	if base.Ui64(base.I64_reinterpret_f64(v1116)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L237
	} else {
		goto L242
	}
L240:
	;
	goto L241
L241:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1017))) = v1116
	goto L237
L242:
	;
	goto L241
L243:
	;
	v1144 = *(*float64)(unsafe.Add(mBase, uint32(v1017)+8))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1144)&int64(9223372036854775807)) {
		goto L249
	} else {
		goto L250
	}
L244:
	;
	v1133 = *(*float64)(unsafe.Add(mBase, uint32(v1017)+16))
	if base.F64_gt(v1133, v1127) == int32(0) {
		goto L245
	} else {
		goto L246
	}
L245:
	;
	if base.Ui64(base.I64_reinterpret_f64(v1133)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L243
	} else {
		goto L248
	}
L246:
	;
	goto L247
L247:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1017)+16)) = v1127
	goto L243
L248:
	;
	goto L247
L249:
	;
	v1161 = *(*float64)(unsafe.Add(mBase, uint32(v1070)+24))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1161)&int64(9223372036854775807)) {
		goto L233
	} else {
		goto L255
	}
L250:
	;
	v1150 = *(*float64)(unsafe.Add(mBase, uint32(v1070)+8))
	if base.F64_lt(v1144, v1150) == int32(0) {
		goto L251
	} else {
		goto L252
	}
L251:
	;
	if base.Ui64(base.I64_reinterpret_f64(v1150)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L249
	} else {
		goto L254
	}
L252:
	;
	goto L253
L253:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1017)+8)) = v1150
	goto L249
L254:
	;
	goto L253
L255:
	;
	v1167 = *(*float64)(unsafe.Add(mBase, uint32(v1017)+24))
	if base.F64_gt(v1167, v1161) == int32(0) {
		goto L256
	} else {
		goto L257
	}
L256:
	;
	if base.Ui64(base.I64_reinterpret_f64(v1167)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L233
	} else {
		goto L259
	}
L257:
	;
	goto L258
L258:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1017)+24)) = v1161
	goto L233
L259:
	;
	goto L258
L260:
	;
	v1276 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
	v1277 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = v1276 + v1277
	v1280 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
	*(*uint16)(unsafe.Add(mBase, uint32(v1280+v1276<<(uint(v1277)%32)))) = uint16(v1051)
	v1288 = v1050
	goto L221
L261:
	;
	v1199 = *(*float64)(unsafe.Add(mBase, uint32(v1020)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1199)&int64(9223372036854775807)) {
		goto L264
	} else {
		goto L265
	}
L262:
	;
	goto L263
L263:
	;
	v1266 = *(*int64)(unsafe.Add(mBase, uint32(v1070)))
	*(*int64)(unsafe.Add(mBase, uint32(v1020))) = v1266
	v1268 = *(*int64)(unsafe.Add(mBase, uint32(v1070)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1020)+24)) = v1268
	v1270 = *(*int64)(unsafe.Add(mBase, uint32(v1070)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1020)+16)) = v1270
	v1272 = *(*int64)(unsafe.Add(mBase, uint32(v1070)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1020)+8)) = v1272
	goto L260
L264:
	;
	v1216 = *(*float64)(unsafe.Add(mBase, uint32(v1070)+16))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1216)&int64(9223372036854775807)) {
		goto L270
	} else {
		goto L271
	}
L265:
	;
	v1205 = *(*float64)(unsafe.Add(mBase, uint32(v1070)))
	if base.F64_lt(v1199, v1205) == int32(0) {
		goto L266
	} else {
		goto L267
	}
L266:
	;
	if base.Ui64(base.I64_reinterpret_f64(v1205)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L264
	} else {
		goto L269
	}
L267:
	;
	goto L268
L268:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1020))) = v1205
	goto L264
L269:
	;
	goto L268
L270:
	;
	v1233 = *(*float64)(unsafe.Add(mBase, uint32(v1020)+8))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1233)&int64(9223372036854775807)) {
		goto L276
	} else {
		goto L277
	}
L271:
	;
	v1222 = *(*float64)(unsafe.Add(mBase, uint32(v1020)+16))
	if base.F64_gt(v1222, v1216) == int32(0) {
		goto L272
	} else {
		goto L273
	}
L272:
	;
	if base.Ui64(base.I64_reinterpret_f64(v1222)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L270
	} else {
		goto L275
	}
L273:
	;
	goto L274
L274:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1020)+16)) = v1216
	goto L270
L275:
	;
	goto L274
L276:
	;
	v1250 = *(*float64)(unsafe.Add(mBase, uint32(v1070)+24))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1250)&int64(9223372036854775807)) {
		goto L260
	} else {
		goto L282
	}
L277:
	;
	v1239 = *(*float64)(unsafe.Add(mBase, uint32(v1070)+8))
	if base.F64_lt(v1233, v1239) == int32(0) {
		goto L278
	} else {
		goto L279
	}
L278:
	;
	if base.Ui64(base.I64_reinterpret_f64(v1239)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L276
	} else {
		goto L281
	}
L279:
	;
	goto L280
L280:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1020)+8)) = v1239
	goto L276
L281:
	;
	goto L280
L282:
	;
	v1256 = *(*float64)(unsafe.Add(mBase, uint32(v1020)+24))
	if base.F64_gt(v1256, v1250) == int32(0) {
		goto L283
	} else {
		goto L284
	}
L283:
	;
	if base.Ui64(base.I64_reinterpret_f64(v1256)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L260
	} else {
		goto L286
	}
L284:
	;
	goto L285
L285:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1020)+24)) = v1250
	goto L260
L286:
	;
	goto L285
L287:
	;
	goto L220
L288:
	;
	v1301 = base.F64_ceil(base.F64_mul(base.F64_convert_i32_u(v43), float64(0.3)))
	if base.F64_lt(base.F64_abs(v1301), float64(2.147483648e+09)) != 0 {
		goto L290
	} else {
		goto L291
	}
L289:
	;
	v1308 = int32(0)
	v1316 = v1308
	v1317 = v1308
	goto L293
L290:
	;
	v1305 = base.I32_trunc_f64_s(v1301)
	v1307 = v1305
	goto L289
L291:
	;
	goto L292
L292:
	;
	v1307 = int32(-2147483648)
	goto L289
L293:
	;
	v1336 = int32(4)
	v1338 = v1022 + v1316<<(uint(v1336)%32)
	v1339 = *(*int32)(unsafe.Add(mBase, uint32(v1338)))
	v1343 = *(*int32)(unsafe.Add(mBase, uint32(v207+v1339<<(uint(v1336)%32))))
	v1344 = F_box_penalty(m, v1017, v1343)
	mBase = m.M
	v1345 = m.ExcPending
	if v1345 != 0 {
		goto L2
	} else {
		goto L296
	}
L294:
	;
	F_pg_qsort(m, v1022, v1288, int32(16), int32(107))
	mBase = m.M
	v1367 = m.ExcPending
	if v1367 != 0 {
		goto L2
	} else {
		goto L302
	}
L295:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1338)+8)) = v1349
	v1360 = v1317 + int32(1)
	v1362 = v1360 & int32(65535)
	if base.Ui32(v1362) < base.Ui32(v1288) {
		v1316 = v1362
		v1317 = v1360
		goto L293
	} else {
		goto L301
	}
L296:
	;
	v1346 = F_box_penalty(m, v1020, v1343)
	mBase = m.M
	v1347 = m.ExcPending
	if v1347 != 0 {
		goto L2
	} else {
		goto L297
	}
L297:
	;
	v1349 = base.F64_abs(base.F64_sub(v1344, v1346))
	if base.F64_ne(v1349, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L295
	} else {
		goto L298
	}
L298:
	;
	if base.F64_eq(base.F64_abs(v1344), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L295
	} else {
		goto L299
	}
L299:
	;
	if base.F64_ne(base.F64_abs(v1346), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L136
	} else {
		goto L300
	}
L300:
	;
	goto L295
L301:
	;
	goto L294
L302:
	;
	v1368 = int32(0)
	v1377 = v1368
	v1380 = v1368
	goto L303
L303:
	;
	v1396 = int32(4)
	v1398 = v1022 + v1377<<(uint(v1396)%32)
	v1399 = *(*int32)(unsafe.Add(mBase, uint32(v1398)))
	v1403 = *(*int32)(unsafe.Add(mBase, uint32(v207+v1399<<(uint(v1396)%32))))
	v1404 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v1405 = v1288 - v1377
	if v1404+v1405 <= v1307 {
		goto L306
	} else {
		goto L307
	}
L304:
	;
	goto L214
L305:
	;
	v1780 = v1380 + int32(1)
	v1782 = v1780 & int32(65535)
	if base.Ui32(v1782) < base.Ui32(v1288) {
		v1377 = v1782
		v1380 = v1780
		goto L303
	} else {
		goto L425
	}
L306:
	;
	if int32(0) < v1404 {
		goto L310
	} else {
		goto L311
	}
L307:
	;
	goto L308
L308:
	;
	v1497 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
	if v1497+v1405 <= v1307 {
		goto L336
	} else {
		goto L337
	}
L309:
	;
	v1487 = *(*int32)(unsafe.Add(mBase, uint32(v1398)))
	v1488 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v1489 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v1488 + v1489
	v1492 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1492+v1488<<(uint(v1489)%32)))) = uint16(v1487)
	goto L305
L310:
	;
	v1410 = *(*float64)(unsafe.Add(mBase, uint32(v1017)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1410)&int64(9223372036854775807)) {
		goto L313
	} else {
		goto L314
	}
L311:
	;
	goto L312
L312:
	;
	v1477 = *(*int64)(unsafe.Add(mBase, uint32(v1403)))
	*(*int64)(unsafe.Add(mBase, uint32(v1017))) = v1477
	v1479 = *(*int64)(unsafe.Add(mBase, uint32(v1403)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1017)+24)) = v1479
	v1481 = *(*int64)(unsafe.Add(mBase, uint32(v1403)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1017)+16)) = v1481
	v1483 = *(*int64)(unsafe.Add(mBase, uint32(v1403)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1017)+8)) = v1483
	goto L309
L313:
	;
	v1427 = *(*float64)(unsafe.Add(mBase, uint32(v1403)+16))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1427)&int64(9223372036854775807)) {
		goto L319
	} else {
		goto L320
	}
L314:
	;
	v1416 = *(*float64)(unsafe.Add(mBase, uint32(v1403)))
	if base.F64_lt(v1410, v1416) == int32(0) {
		goto L315
	} else {
		goto L316
	}
L315:
	;
	if base.Ui64(base.I64_reinterpret_f64(v1416)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L313
	} else {
		goto L318
	}
L316:
	;
	goto L317
L317:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1017))) = v1416
	goto L313
L318:
	;
	goto L317
L319:
	;
	v1444 = *(*float64)(unsafe.Add(mBase, uint32(v1017)+8))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1444)&int64(9223372036854775807)) {
		goto L325
	} else {
		goto L326
	}
L320:
	;
	v1433 = *(*float64)(unsafe.Add(mBase, uint32(v1017)+16))
	if base.F64_gt(v1433, v1427) == int32(0) {
		goto L321
	} else {
		goto L322
	}
L321:
	;
	if base.Ui64(base.I64_reinterpret_f64(v1433)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L319
	} else {
		goto L324
	}
L322:
	;
	goto L323
L323:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1017)+16)) = v1427
	goto L319
L324:
	;
	goto L323
L325:
	;
	v1461 = *(*float64)(unsafe.Add(mBase, uint32(v1403)+24))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1461)&int64(9223372036854775807)) {
		goto L309
	} else {
		goto L331
	}
L326:
	;
	v1450 = *(*float64)(unsafe.Add(mBase, uint32(v1403)+8))
	if base.F64_lt(v1444, v1450) == int32(0) {
		goto L327
	} else {
		goto L328
	}
L327:
	;
	if base.Ui64(base.I64_reinterpret_f64(v1450)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L325
	} else {
		goto L330
	}
L328:
	;
	goto L329
L329:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1017)+8)) = v1450
	goto L325
L330:
	;
	goto L329
L331:
	;
	v1467 = *(*float64)(unsafe.Add(mBase, uint32(v1017)+24))
	if base.F64_gt(v1467, v1461) == int32(0) {
		goto L332
	} else {
		goto L333
	}
L332:
	;
	if base.Ui64(base.I64_reinterpret_f64(v1467)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L309
	} else {
		goto L335
	}
L333:
	;
	goto L334
L334:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1017)+24)) = v1461
	goto L309
L335:
	;
	goto L334
L336:
	;
	if int32(0) < v1497 {
		goto L340
	} else {
		goto L341
	}
L337:
	;
	goto L338
L338:
	;
	v1589 = F_box_penalty(m, v1017, v1403)
	mBase = m.M
	v1590 = m.ExcPending
	if v1590 != 0 {
		goto L2
	} else {
		goto L366
	}
L339:
	;
	v1579 = *(*int32)(unsafe.Add(mBase, uint32(v1398)))
	v1580 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
	v1581 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = v1580 + v1581
	v1584 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
	*(*uint16)(unsafe.Add(mBase, uint32(v1584+v1580<<(uint(v1581)%32)))) = uint16(v1579)
	goto L305
L340:
	;
	v1502 = *(*float64)(unsafe.Add(mBase, uint32(v1020)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1502)&int64(9223372036854775807)) {
		goto L343
	} else {
		goto L344
	}
L341:
	;
	goto L342
L342:
	;
	v1569 = *(*int64)(unsafe.Add(mBase, uint32(v1403)))
	*(*int64)(unsafe.Add(mBase, uint32(v1020))) = v1569
	v1571 = *(*int64)(unsafe.Add(mBase, uint32(v1403)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1020)+24)) = v1571
	v1573 = *(*int64)(unsafe.Add(mBase, uint32(v1403)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1020)+16)) = v1573
	v1575 = *(*int64)(unsafe.Add(mBase, uint32(v1403)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1020)+8)) = v1575
	goto L339
L343:
	;
	v1519 = *(*float64)(unsafe.Add(mBase, uint32(v1403)+16))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1519)&int64(9223372036854775807)) {
		goto L349
	} else {
		goto L350
	}
L344:
	;
	v1508 = *(*float64)(unsafe.Add(mBase, uint32(v1403)))
	if base.F64_lt(v1502, v1508) == int32(0) {
		goto L345
	} else {
		goto L346
	}
L345:
	;
	if base.Ui64(base.I64_reinterpret_f64(v1508)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L343
	} else {
		goto L348
	}
L346:
	;
	goto L347
L347:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1020))) = v1508
	goto L343
L348:
	;
	goto L347
L349:
	;
	v1536 = *(*float64)(unsafe.Add(mBase, uint32(v1020)+8))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1536)&int64(9223372036854775807)) {
		goto L355
	} else {
		goto L356
	}
L350:
	;
	v1525 = *(*float64)(unsafe.Add(mBase, uint32(v1020)+16))
	if base.F64_gt(v1525, v1519) == int32(0) {
		goto L351
	} else {
		goto L352
	}
L351:
	;
	if base.Ui64(base.I64_reinterpret_f64(v1525)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L349
	} else {
		goto L354
	}
L352:
	;
	goto L353
L353:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1020)+16)) = v1519
	goto L349
L354:
	;
	goto L353
L355:
	;
	v1553 = *(*float64)(unsafe.Add(mBase, uint32(v1403)+24))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1553)&int64(9223372036854775807)) {
		goto L339
	} else {
		goto L361
	}
L356:
	;
	v1542 = *(*float64)(unsafe.Add(mBase, uint32(v1403)+8))
	if base.F64_lt(v1536, v1542) == int32(0) {
		goto L357
	} else {
		goto L358
	}
L357:
	;
	if base.Ui64(base.I64_reinterpret_f64(v1542)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L355
	} else {
		goto L360
	}
L358:
	;
	goto L359
L359:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1020)+8)) = v1542
	goto L355
L360:
	;
	goto L359
L361:
	;
	v1559 = *(*float64)(unsafe.Add(mBase, uint32(v1020)+24))
	if base.F64_gt(v1559, v1553) == int32(0) {
		goto L362
	} else {
		goto L363
	}
L362:
	;
	if base.Ui64(base.I64_reinterpret_f64(v1559)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L339
	} else {
		goto L365
	}
L363:
	;
	goto L364
L364:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1020)+24)) = v1553
	goto L339
L365:
	;
	goto L364
L366:
	;
	v1591 = F_box_penalty(m, v1020, v1403)
	mBase = m.M
	v1592 = m.ExcPending
	if v1592 != 0 {
		goto L2
	} else {
		goto L367
	}
L367:
	;
	if base.F64_lt(v1589, v1591) != 0 {
		goto L368
	} else {
		goto L369
	}
L368:
	;
	v1594 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if int32(0) < v1594 {
		goto L372
	} else {
		goto L373
	}
L369:
	;
	goto L370
L370:
	;
	v1684 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
	if int32(0) < v1684 {
		goto L399
	} else {
		goto L400
	}
L371:
	;
	v1674 = *(*int32)(unsafe.Add(mBase, uint32(v1398)))
	v1675 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v1676 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v1675 + v1676
	v1679 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1679+v1675<<(uint(v1676)%32)))) = uint16(v1674)
	goto L305
L372:
	;
	v1597 = *(*float64)(unsafe.Add(mBase, uint32(v1017)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1597)&int64(9223372036854775807)) {
		goto L375
	} else {
		goto L376
	}
L373:
	;
	goto L374
L374:
	;
	v1664 = *(*int64)(unsafe.Add(mBase, uint32(v1403)))
	*(*int64)(unsafe.Add(mBase, uint32(v1017))) = v1664
	v1666 = *(*int64)(unsafe.Add(mBase, uint32(v1403)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1017)+24)) = v1666
	v1668 = *(*int64)(unsafe.Add(mBase, uint32(v1403)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1017)+16)) = v1668
	v1670 = *(*int64)(unsafe.Add(mBase, uint32(v1403)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1017)+8)) = v1670
	goto L371
L375:
	;
	v1614 = *(*float64)(unsafe.Add(mBase, uint32(v1403)+16))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1614)&int64(9223372036854775807)) {
		goto L381
	} else {
		goto L382
	}
L376:
	;
	v1603 = *(*float64)(unsafe.Add(mBase, uint32(v1403)))
	if base.F64_lt(v1597, v1603) == int32(0) {
		goto L377
	} else {
		goto L378
	}
L377:
	;
	if base.Ui64(base.I64_reinterpret_f64(v1603)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L375
	} else {
		goto L380
	}
L378:
	;
	goto L379
L379:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1017))) = v1603
	goto L375
L380:
	;
	goto L379
L381:
	;
	v1631 = *(*float64)(unsafe.Add(mBase, uint32(v1017)+8))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1631)&int64(9223372036854775807)) {
		goto L387
	} else {
		goto L388
	}
L382:
	;
	v1620 = *(*float64)(unsafe.Add(mBase, uint32(v1017)+16))
	if base.F64_gt(v1620, v1614) == int32(0) {
		goto L383
	} else {
		goto L384
	}
L383:
	;
	if base.Ui64(base.I64_reinterpret_f64(v1620)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L381
	} else {
		goto L386
	}
L384:
	;
	goto L385
L385:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1017)+16)) = v1614
	goto L381
L386:
	;
	goto L385
L387:
	;
	v1648 = *(*float64)(unsafe.Add(mBase, uint32(v1403)+24))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1648)&int64(9223372036854775807)) {
		goto L371
	} else {
		goto L393
	}
L388:
	;
	v1637 = *(*float64)(unsafe.Add(mBase, uint32(v1403)+8))
	if base.F64_lt(v1631, v1637) == int32(0) {
		goto L389
	} else {
		goto L390
	}
L389:
	;
	if base.Ui64(base.I64_reinterpret_f64(v1637)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L387
	} else {
		goto L392
	}
L390:
	;
	goto L391
L391:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1017)+8)) = v1637
	goto L387
L392:
	;
	goto L391
L393:
	;
	v1654 = *(*float64)(unsafe.Add(mBase, uint32(v1017)+24))
	if base.F64_gt(v1654, v1648) == int32(0) {
		goto L394
	} else {
		goto L395
	}
L394:
	;
	if base.Ui64(base.I64_reinterpret_f64(v1654)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L371
	} else {
		goto L397
	}
L395:
	;
	goto L396
L396:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1017)+24)) = v1648
	goto L371
L397:
	;
	goto L396
L398:
	;
	v1764 = *(*int32)(unsafe.Add(mBase, uint32(v1398)))
	v1765 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
	v1766 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = v1765 + v1766
	v1769 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
	*(*uint16)(unsafe.Add(mBase, uint32(v1769+v1765<<(uint(v1766)%32)))) = uint16(v1764)
	goto L305
L399:
	;
	v1687 = *(*float64)(unsafe.Add(mBase, uint32(v1020)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1687)&int64(9223372036854775807)) {
		goto L402
	} else {
		goto L403
	}
L400:
	;
	goto L401
L401:
	;
	v1754 = *(*int64)(unsafe.Add(mBase, uint32(v1403)))
	*(*int64)(unsafe.Add(mBase, uint32(v1020))) = v1754
	v1756 = *(*int64)(unsafe.Add(mBase, uint32(v1403)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1020)+24)) = v1756
	v1758 = *(*int64)(unsafe.Add(mBase, uint32(v1403)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1020)+16)) = v1758
	v1760 = *(*int64)(unsafe.Add(mBase, uint32(v1403)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1020)+8)) = v1760
	goto L398
L402:
	;
	v1704 = *(*float64)(unsafe.Add(mBase, uint32(v1403)+16))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1704)&int64(9223372036854775807)) {
		goto L408
	} else {
		goto L409
	}
L403:
	;
	v1693 = *(*float64)(unsafe.Add(mBase, uint32(v1403)))
	if base.F64_lt(v1687, v1693) == int32(0) {
		goto L404
	} else {
		goto L405
	}
L404:
	;
	if base.Ui64(base.I64_reinterpret_f64(v1693)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L402
	} else {
		goto L407
	}
L405:
	;
	goto L406
L406:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1020))) = v1693
	goto L402
L407:
	;
	goto L406
L408:
	;
	v1721 = *(*float64)(unsafe.Add(mBase, uint32(v1020)+8))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1721)&int64(9223372036854775807)) {
		goto L414
	} else {
		goto L415
	}
L409:
	;
	v1710 = *(*float64)(unsafe.Add(mBase, uint32(v1020)+16))
	if base.F64_gt(v1710, v1704) == int32(0) {
		goto L410
	} else {
		goto L411
	}
L410:
	;
	if base.Ui64(base.I64_reinterpret_f64(v1710)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L408
	} else {
		goto L413
	}
L411:
	;
	goto L412
L412:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1020)+16)) = v1704
	goto L408
L413:
	;
	goto L412
L414:
	;
	v1738 = *(*float64)(unsafe.Add(mBase, uint32(v1403)+24))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1738)&int64(9223372036854775807)) {
		goto L398
	} else {
		goto L420
	}
L415:
	;
	v1727 = *(*float64)(unsafe.Add(mBase, uint32(v1403)+8))
	if base.F64_lt(v1721, v1727) == int32(0) {
		goto L416
	} else {
		goto L417
	}
L416:
	;
	if base.Ui64(base.I64_reinterpret_f64(v1727)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L414
	} else {
		goto L419
	}
L417:
	;
	goto L418
L418:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1020)+8)) = v1727
	goto L414
L419:
	;
	goto L418
L420:
	;
	v1744 = *(*float64)(unsafe.Add(mBase, uint32(v1020)+24))
	if base.F64_gt(v1744, v1738) == int32(0) {
		goto L421
	} else {
		goto L422
	}
L421:
	;
	if base.Ui64(base.I64_reinterpret_f64(v1744)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L398
	} else {
		goto L424
	}
L422:
	;
	goto L423
L423:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1020)+24)) = v1738
	goto L398
L424:
	;
	goto L423
L425:
	;
	goto L304
L426:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_gist_circle_consistent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 float64
	_ = v25
	var v26 float64
	_ = v26
	var v27 float64
	_ = v27
	var v38 float64
	_ = v38
	var v49 float64
	_ = v49
	var v50 float64
	_ = v50
	var v61 float64
	_ = v61
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v89 int32
	_ = v89
	v5 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v17 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v16))) = uint8(v17)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if v20 == v5 {
		v78 = v5
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L21
	} else {
		goto L23
	}
L2:
	;
	m.G0 = v11 + int32(32)
	return v78
L3:
	;
	if v14 == int32(0) {
		v78 = v5
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v25 = *(*float64)(unsafe.Add(mBase, uint32(v14)))
	v26 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
	v27 = base.F64_add(v25, v26)
	if base.F64_ne(base.F64_abs(v27), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v11))) = v27
	v38 = base.F64_sub(v25, v26)
	if base.F64_ne(base.F64_abs(v38), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	if base.F64_eq(base.F64_abs(v25), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	if base.F64_ne(base.F64_abs(v26), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	goto L5
L9:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v38
	v49 = *(*float64)(unsafe.Add(mBase, uint32(v14)+8))
	v50 = base.F64_add(v49, v26)
	if base.F64_ne(base.F64_abs(v50), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	if base.F64_eq(base.F64_abs(v26), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	if base.F64_ne(base.F64_abs(v25), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	goto L9
L13:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v50
	v61 = base.F64_sub(v49, v26)
	if base.F64_ne(base.F64_abs(v61), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	if base.F64_eq(base.F64_abs(v49), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	if base.F64_ne(base.F64_abs(v26), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	goto L13
L17:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v61
	v74 = F_rtree_internal_consistent(m, v20, v11, v13&int32(65535))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	if base.F64_eq(base.F64_abs(v26), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	if base.F64_ne(base.F64_abs(v49), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	goto L17
L21:
	;
	return int32(0)
L22:
	;
	v78 = v74
	goto L2
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_gist_desc(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v27 int32
	_ = v27
	var v30 int64
	_ = v30
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int64
	_ = v60
	var v61 int32
	_ = v61
	var v65 int64
	_ = v65
	var v69 int32
	_ = v69
	v9 = m.G0
	v11 = v9 - int32(80)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+64))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+48)))
	switch int32(base.Ui32(v15)>>(uint(int32(4))%32)) - int32(1) {
	case 0:
		v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+6)))
		v40 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
		v41 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+4)))
		*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v41
		*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v40
		if v39 != 0 {
			v46 = int32(84)
		} else {
			v46 = int32(70)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = v46
		F_appendStringInfo(m, l0, int32(513361), v11+int32(48))
		mBase = m.M
		v52 = m.ExcPending
		if v52 != 0 {
			return
		} else {
			m.G0 = v11 + int32(80)
			return
		}
	case 1:
		v20 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
		v21 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
		v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+24)))
		v23 = *(*int64)(unsafe.Add(mBase, uint32(v14)+16))
		*(*uint32)(unsafe.Add(mBase, uint32(v11)+36)) = uint32(v23)
		if v22 != 0 {
			v27 = int32(84)
		} else {
			v27 = int32(70)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v27
		v30 = int64(base.Ui64(v23) >> (uint(int64(32)) % 64))
		*(*uint32)(unsafe.Add(mBase, uint32(v11)+32)) = uint32(v30)
		*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v21
		*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v20
		F_appendStringInfo(m, l0, int32(513236), v11+int32(16))
		mBase = m.M
		v38 = m.ExcPending
		if v38 != 0 {
			return
		} else {
			m.G0 = v11 + int32(80)
			return
		}
	case 2:
		v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+18)))
		*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v53
		F_appendStringInfo(m, l0, int32(174953), v11-int32(-64))
		mBase = m.M
		v59 = m.ExcPending
		if v59 != 0 {
			return
		} else {
			m.G0 = v11 + int32(80)
			return
		}
	default:
		m.G0 = v11 + int32(80)
		return
	case 5:
		v60 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
		v61 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+8)))
		*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v61
		*(*uint32)(unsafe.Add(mBase, uint32(v11)+4)) = uint32(v60)
		v65 = int64(base.Ui64(v60) >> (uint(int64(32)) % 64))
		*(*uint32)(unsafe.Add(mBase, uint32(v11))) = uint32(v65)
		F_appendStringInfo(m, l0, int32(48404), v11)
		mBase = m.M
		v69 = m.ExcPending
		if v69 != 0 {
			return
		} else {
			m.G0 = v11 + int32(80)
			return
		}
	}
}
func F_gist_point_compress(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v25 int64
	_ = v25
	var v27 int64
	_ = v27
	var v28 int64
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+14)))
	if v8 != int32(1) {
		return v7
	} else {
		v13 = F_palloc(m, int32(32))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
			v19 = F_palloc(m, int32(16))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				v22 = v17 + int32(8)
				v23 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
				*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v23
				v25 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
				*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v25
				v27 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
				v28 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
				*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v28
				*(*int64)(unsafe.Add(mBase, uint32(v13))) = v27
				*(*int32)(unsafe.Add(mBase, uint32(v19))) = v13
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v32
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v34
				v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+12)))
				v37 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v19)+14)) = uint8(v37)
				*(*uint16)(unsafe.Add(mBase, uint32(v19)+12)) = uint16(v36)
				return v19
			}
		}
	}
}
func F_gist_point_fetch(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 float64
	_ = v15
	var v17 float64
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v8 = F_palloc(m, int32(16))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v13 = F_palloc(m, int32(16))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v15 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
			*(*float64)(unsafe.Add(mBase, uint32(v13))) = v15
			v17 = *(*float64)(unsafe.Add(mBase, uint32(v6)+8))
			*(*float64)(unsafe.Add(mBase, uint32(v13)+8)) = v17
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = v13
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v20
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v22
			v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5)+12)))
			v25 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v8)+14)) = uint8(v25)
			*(*uint16)(unsafe.Add(mBase, uint32(v8)+12)) = uint16(v24)
			return v8
		}
	}
}
func F_gist_poly_compress(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v19 int64
	_ = v19
	var v21 int64
	_ = v21
	var v23 int64
	_ = v23
	var v26 int32
	_ = v26
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
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+14)))
	if v5 != int32(1) {
		return v4
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		v10 = F_pg_detoast_datum(m, v9)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v15 = F_palloc(m, int32(32))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				v17 = *(*int64)(unsafe.Add(mBase, uint32(v10)+32))
				*(*int64)(unsafe.Add(mBase, uint32(v15)+24)) = v17
				v19 = *(*int64)(unsafe.Add(mBase, uint32(v10)+24))
				*(*int64)(unsafe.Add(mBase, uint32(v15)+16)) = v19
				v21 = *(*int64)(unsafe.Add(mBase, uint32(v10)+16))
				*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = v21
				v23 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v15))) = v23
				v26 = F_palloc(m, int32(16))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v26))) = v15
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v29
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v26)+8)) = v31
					v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4)+12)))
					v34 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v26)+14)) = uint8(v34)
					*(*uint16)(unsafe.Add(mBase, uint32(v26)+12)) = uint16(v33)
					return v26
				}
			}
		}
	}
}
