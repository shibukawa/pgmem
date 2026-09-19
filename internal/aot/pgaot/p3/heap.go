package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_HeapTupleSetHintBits(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int64
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int64
	_ = v25
	var v26 int64
	_ = v26
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int64
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	if l3 == int32(0) {
		v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
		v56 = v55 | l2
		*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)) = uint16(v56)
		F_MarkBufferDirtyHint(m, l1, int32(1))
		mBase = m.M
		v60 = m.ExcPending
		if v60 != 0 {
			return
		} else {
			return
		}
	} else {
		v9 = m.G0
		v11 = v9 - int32(16)
		m.G0 = v11
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSetHintBits[0]))
		if v14 == l3 {
			v17 = *(*int64)(unsafe.Add(mBase, _c_F_HeapTupleSetHintBits[1]))
			v26 = v17
			m.G0 = v11 + int32(16)
			if int32(0) <= l1 {
				v33 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSetHintBits[2]))
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v33+l1<<(uint(int32(6))%32)-int32(40))))
				v43 = int32(base.Ui32(v39) >> (uint(int32(31)) % 32))
			} else {
				v43 = int32(0)
			}
			if v43 == int32(0) {
				v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
				v56 = v55 | l2
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)) = uint16(v56)
				F_MarkBufferDirtyHint(m, l1, int32(1))
				mBase = m.M
				v60 = m.ExcPending
				if v60 != 0 {
					return
				} else {
					return
				}
			} else {
				v46 = F_XLogNeedsFlush(m, v26)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return
				} else {
					if v46 == int32(0) {
						v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
						v56 = v55 | l2
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)) = uint16(v56)
						F_MarkBufferDirtyHint(m, l1, int32(1))
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return
						} else {
							return
						}
					} else {
						v50 = F_BufferGetLSNAtomic(m, l1)
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return
						} else {
							if base.Ui64(v50) < base.Ui64(v26) {
								return
							} else {
								v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
								v56 = v55 | l2
								*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)) = uint16(v56)
								F_MarkBufferDirtyHint(m, l1, int32(1))
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				}
			}
		} else {
			if base.Ui32(l3) < base.Ui32(int32(3)) {
				v26 = int64(0)
				m.G0 = v11 + int32(16)
				if int32(0) <= l1 {
					v33 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSetHintBits[2]))
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v33+l1<<(uint(int32(6))%32)-int32(40))))
					v43 = int32(base.Ui32(v39) >> (uint(int32(31)) % 32))
				} else {
					v43 = int32(0)
				}
				if v43 == int32(0) {
					v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
					v56 = v55 | l2
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)) = uint16(v56)
					F_MarkBufferDirtyHint(m, l1, int32(1))
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return
					} else {
						return
					}
				} else {
					v46 = F_XLogNeedsFlush(m, v26)
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return
					} else {
						if v46 == int32(0) {
							v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
							v56 = v55 | l2
							*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)) = uint16(v56)
							F_MarkBufferDirtyHint(m, l1, int32(1))
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return
							} else {
								return
							}
						} else {
							v50 = F_BufferGetLSNAtomic(m, l1)
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return
							} else {
								if base.Ui64(v50) < base.Ui64(v26) {
									return
								} else {
									v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
									v56 = v55 | l2
									*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)) = uint16(v56)
									F_MarkBufferDirtyHint(m, l1, int32(1))
									mBase = m.M
									v60 = m.ExcPending
									if v60 != 0 {
										return
									} else {
										return
									}
								}
							}
						}
					}
				}
			} else {
				v23 = F_TransactionIdGetStatus(m, l3, v11+int32(8))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					v25 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
					v26 = v25
					m.G0 = v11 + int32(16)
					if int32(0) <= l1 {
						v33 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSetHintBits[2]))
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v33+l1<<(uint(int32(6))%32)-int32(40))))
						v43 = int32(base.Ui32(v39) >> (uint(int32(31)) % 32))
					} else {
						v43 = int32(0)
					}
					if v43 == int32(0) {
						v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
						v56 = v55 | l2
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)) = uint16(v56)
						F_MarkBufferDirtyHint(m, l1, int32(1))
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return
						} else {
							return
						}
					} else {
						v46 = F_XLogNeedsFlush(m, v26)
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return
						} else {
							if v46 == int32(0) {
								v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
								v56 = v55 | l2
								*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)) = uint16(v56)
								F_MarkBufferDirtyHint(m, l1, int32(1))
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return
								} else {
									return
								}
							} else {
								v50 = F_BufferGetLSNAtomic(m, l1)
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return
								} else {
									if base.Ui64(v50) < base.Ui64(v26) {
										return
									} else {
										v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
										v56 = v55 | l2
										*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)) = uint16(v56)
										F_MarkBufferDirtyHint(m, l1, int32(1))
										mBase = m.M
										v60 = m.ExcPending
										if v60 != 0 {
											return
										} else {
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
}
func F_heap_deform_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v62 int32
	_ = v62
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v214 int32
	_ = v214
	var v231 int32
	_ = v231
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	v5 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+18)))
	v25 = v23 & int32(2047)
	if v21 < v25 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v214 < v21 {
		goto L58
	} else {
		goto L59
	}
L2:
	;
	v27 = v21
	goto L4
L3:
	;
	v27 = v25
	goto L4
L4:
	;
	if v27 <= int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v214 = int32(0)
	goto L1
L6:
	;
	goto L7
L7:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+22)))
	v34 = v22 + v33
	v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+20)))
	v41 = int32(0)
	v47 = v5
	v49 = v5
	goto L8
L8:
	;
	if v35&int32(1) == int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v214 = v27
	goto L1
L10:
	;
	v205 = v41 + int32(1)
	if v205 != v27 {
		v41 = v205
		v47 = v200
		v49 = v203
		goto L8
	} else {
		goto L57
	}
L11:
	;
	v79 = l1 + int32(20) + v41<<(uint(int32(4))%32)
	v81 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v41+l3))) = uint8(v81)
	if v49&int32(1) == v81 {
		goto L18
	} else {
		goto L19
	}
L12:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+int32(23)+int32(base.Ui32(v41)>>(uint(int32(3))%32))))))
	if int32(base.Ui32(v62)>>(uint(v41&int32(7))%32))&int32(1) != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2+v41<<(uint(int32(2))%32)))) = int32(0)
	v74 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v41+l3))) = uint8(v74)
	v200 = v47
	v203 = v74
	goto L10
L14:
	;
	v130 = v128 + v34
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+6)))
	if v134 == int32(1) {
		goto L29
	} else {
		goto L30
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v79))) = v97
	v128 = v97
	v129 = v87
	goto L14
L16:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+12)))
	v120 = int32(1)
	v128 = (v47 + v118 - v120) & (int32(0) - v118)
	v129 = v120
	goto L14
L17:
	;
	v108 = int32(1)
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47+v34))))
	if v110 != 0 {
		v128 = v47
		v129 = v108
		goto L14
	} else {
		goto L27
	}
L18:
	;
	v87 = int32(0)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	if v87 <= v88 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	v103 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79)+4)))
	if v103 != int32(_a_F_heap_deform_tuple_0) {
		goto L16
	} else {
		goto L26
	}
L21:
	;
	v128 = v88
	v129 = v87
	goto L14
L22:
	;
	goto L23
L23:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+12)))
	v97 = (v47 + v91 - int32(1)) & (int32(0) - v91)
	v98 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79)+4)))
	if v98 != int32(_a_F_heap_deform_tuple_0) {
		goto L15
	} else {
		goto L24
	}
L24:
	;
	if v97 != v47 {
		goto L17
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v79))) = v47
	v128 = v47
	v129 = v87
	goto L14
L26:
	;
	goto L17
L27:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+12)))
	v128 = (v47 + v111 - int32(1)) & (int32(0) - v111)
	v129 = v108
	goto L14
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2+v41<<(uint(int32(2))%32)))) = v158
	v160 = int32(*(*int16)(unsafe.Add(mBase, uint32(v79)+4)))
	v161 = int32(0)
	v162 = base.B2i32(v160 <= v161)
	if v162 == v161 {
		v195 = v160
		goto L40
	} else {
		goto L41
	}
L29:
	;
	v137 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79)+4)))
	switch v137 - int32(1) {
	case 0:
		goto L35
	case 1:
		goto L34
	default:
		goto L32
	case 3:
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	v158 = v130
	goto L28
L32:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	v158 = v142
	goto L28
L34:
	;
	v141 = int32(*(*int16)(unsafe.Add(mBase, uint32(v130))))
	v158 = v141
	goto L28
L35:
	;
	v140 = int32(*(*int8)(unsafe.Add(mBase, uint32(v130))))
	v158 = v140
	goto L28
L36:
	;
	return
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = base.I32_extend16_s(v137)
	F_errmsg_internal(m, int32(_a_F_heap_deform_tuple_1), v19)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(_a_F_heap_deform_tuple_2), int32(70), int32(_a_F_heap_deform_tuple_3))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L36
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L40:
	;
	v200 = v195 + v128
	v203 = v129 | v162
	goto L10
L41:
	;
	if v160 == int32(-1) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130))))
	if v167 == int32(1) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	goto L44
L44:
	;
	v191 = F_strlen(m, v130)
	mBase = m.M
	v195 = v191 + int32(1)
	goto L40
L45:
	;
	v171 = int32(18)
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+1)))
	if v173 == v171 {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	goto L47
L47:
	;
	if v167&int32(1) != 0 {
		goto L54
	} else {
		goto L55
	}
L48:
	;
	v176 = v171
	goto L50
L49:
	;
	v176 = int32(2)
	goto L50
L50:
	;
	if base.Ui32((v173-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v183 = int32(6)
	goto L53
L52:
	;
	v183 = v176
	goto L53
L53:
	;
	v195 = v183
	goto L40
L54:
	;
	v195 = int32(base.Ui32(v167) >> (uint(int32(1)) % 32))
	goto L40
L55:
	;
	goto L56
L56:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	v195 = int32(base.Ui32(v188) >> (uint(int32(2)) % 32))
	goto L40
L57:
	;
	goto L9
L58:
	;
	v231 = v214
	goto L61
L59:
	;
	goto L60
L60:
	;
	m.G0 = v19 + int32(16)
	return
L61:
	;
	v245 = v231 + int32(1)
	v246 = F_getmissingattr(m, l1, v245, l3+v231)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L36
	} else {
		goto L63
	}
L62:
	;
	goto L60
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2+v231<<(uint(int32(2))%32)))) = v246
	if v245 != v21 {
		v231 = v245
		goto L61
	} else {
		goto L64
	}
L64:
	;
	goto L62
}
func F_heap_getattr_2(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = Fn13924(m, l0, l1, l2, l3, int32(_a_F_heap_getattr_2_0))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_heap_getattr_4(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+18)))
	if base.Ui32(v11&int32(2047)) <= base.Ui32(int32(2)) {
		v17 = F_getmissingattr(m, l1, int32(3), l2)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v77 = v17
			m.G0 = v8 + int32(16)
			return v77
		}
	} else {
		v21 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v21)
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+20)))
		if v24&int32(1) == v21 {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
			if int32(0) <= v29 {
				v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+22)))
				v34 = v23 + v32 + v29
				v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+58)))
				if v35 != int32(1) {
					v77 = v34
					m.G0 = v8 + int32(16)
					return v77
				} else {
					v38 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+56)))
					switch v38&int32(_a_F_heap_getattr_4_0) - int32(1) {
					case 0:
						v43 = int32(*(*int8)(unsafe.Add(mBase, uint32(v34))))
						v77 = v43
						m.G0 = v8 + int32(16)
						return v77
					case 1:
						v44 = int32(*(*int16)(unsafe.Add(mBase, uint32(v34))))
						v77 = v44
						m.G0 = v8 + int32(16)
						return v77
					default:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v38
							F_errmsg_internal(m, int32(_a_F_heap_getattr_4_1), v8)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_heap_getattr_4_2), int32(70), int32(_a_F_heap_getattr_4_3))
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					case 3:
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
						v77 = v45
						m.G0 = v8 + int32(16)
						return v77
					}
				}
			} else {
				v60 = F_nocachegetattr(m, l0, int32(3), l1)
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return int32(0)
				} else {
					v77 = v60
					m.G0 = v8 + int32(16)
					return v77
				}
			}
		} else {
			v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+23)))
			if v62&int32(4) == int32(0) {
				v67 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v67)
				v77 = int32(0)
				m.G0 = v8 + int32(16)
				return v77
			} else {
				v71 = F_nocachegetattr(m, l0, int32(3), l1)
				mBase = m.M
				v72 = m.ExcPending
				if v72 != 0 {
					return int32(0)
				} else {
					v77 = v71
					m.G0 = v8 + int32(16)
					return v77
				}
			}
		}
	}
}
func F_heap_getnextslot_tidrange(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
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
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
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
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int64
	_ = v125
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	v10 = l0 + int32(68)
	v12 = l0 + int32(22)
	v14 = l0 + int32(16)
	goto L2
L1:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+272))
	if v114 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L2:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
	if v25&int32(1) != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	return int32(0)
L4:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v34 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L5:
	;
	F_heapgettup_pagemode(m, l0, l1, v24, v23)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	F_heapgettup(m, l0, l1, v24, v23)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L8
	} else {
		goto L10
	}
L8:
	;
	return int32(0)
L9:
	;
	goto L4
L10:
	;
	goto L4
L11:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	m.T0[v38].(func(*base.Module, int32))(m, l2)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L8
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+2)))
	v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10))))
	v48 = int32(16)
	v50 = v46 | v47<<(uint(v48)%32)
	v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+2)))
	v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14))))
	v55 = v51 | v52<<(uint(v48)%32)
	if base.Ui32(v50) < base.Ui32(v55) {
		v66 = int32(-1)
		goto L16
	} else {
		goto L17
	}
L14:
	;
	return int32(0)
L15:
	;
	if v66 < int32(0) {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	goto L15
L17:
	;
	if base.Ui32(v55) < base.Ui32(v50) {
		v66 = int32(1)
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+4)))
	v61 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+4)))
	if base.Ui32(v60) < base.Ui32(v61) {
		v66 = int32(-1)
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v66 = base.B2i32(base.Ui32(v61) < base.Ui32(v60))
	goto L16
L20:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	m.T0[v70].(func(*base.Module, int32))(m, l2)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L8
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+2)))
	v81 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10))))
	v82 = int32(16)
	v84 = v80 | v81<<(uint(v82)%32)
	v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+2)))
	v86 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12))))
	v89 = v85 | v86<<(uint(v82)%32)
	if base.Ui32(v84) < base.Ui32(v89) {
		v100 = int32(-1)
		goto L26
	} else {
		goto L27
	}
L23:
	;
	if l1 != int32(-1) {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	return int32(0)
L25:
	;
	if v100 <= int32(0) {
		goto L1
	} else {
		goto L30
	}
L26:
	;
	goto L25
L27:
	;
	if base.Ui32(v89) < base.Ui32(v84) {
		v100 = int32(1)
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+4)))
	v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+4)))
	if base.Ui32(v94) < base.Ui32(v95) {
		v100 = int32(-1)
		goto L26
	} else {
		goto L29
	}
L29:
	;
	v100 = base.B2i32(base.Ui32(v95) < base.Ui32(v94))
	goto L26
L30:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+12))
	m.T0[v104].(func(*base.Module, int32))(m, l2)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L8
	} else {
		goto L31
	}
L31:
	;
	if l1 != int32(1) {
		goto L2
	} else {
		goto L32
	}
L32:
	;
	goto L3
L33:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_ExecStoreBufferHeapTuple(m, l0-int32(-64), l2, v130)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L8
	} else {
		goto L39
	}
L34:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113)+268)))
	if v117 != int32(1) {
		goto L33
	} else {
		goto L37
	}
L35:
	;
	v124 = v114
	goto L36
L36:
	;
	v125 = *(*int64)(unsafe.Add(mBase, uint32(v124)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v124)+24)) = v125 + int64(1)
	goto L33
L37:
	;
	F_pgstat_assoc_relation(m, v113)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L8
	} else {
		goto L38
	}
L38:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+272))
	v124 = v123
	goto L36
L39:
	;
	return int32(1)
}
func F_heap_identify(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	if base.Ui32(l0) <= base.Ui32(int32(207)) {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(int32(base.Ui32(l0)>>(uint(int32(2))%32))&int32(60))+uint32(_c_F_heap_identify[0])))
		v10 = v8
	} else {
		v10 = int32(0)
	}
	return v10
}
func F_heap_lock_updated_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v90 int32
	_ = v90
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v550 int32
	_ = v550
	var v554 int32
	_ = v554
	var v557 int64
	_ = v557
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v703 int32
	_ = v703
	var v713 int32
	_ = v713
	var v716 int32
	_ = v716
	var v722 int32
	_ = v722
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v794 int32
	_ = v794
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v831 int32
	_ = v831
	var v835 int32
	_ = v835
	var v837 int32
	_ = v837
	v7 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(48)
	m.G0 = v21
	v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+4)))
	if v23 != int32(_a_F_heap_lock_updated_tuple_0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v21 + int32(48)
	return v837
L2:
	;
	F_MultiXactIdSetOldestMember(m)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3))))
	v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+2)))
	if v26&v27 != int32(_a_F_heap_lock_updated_tuple_1) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v837 = int32(0)
	goto L1
L5:
	;
	return int32(0)
L6:
	;
	if l1&int32(_a_F_heap_lock_updated_tuple_2) == int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v100 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = v100
	v103 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3))))
	v104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+2)))
	v105 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v100
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+22)) = uint16(v100)
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+36)) = uint16(v105)
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+34)) = uint16(v104)
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+32)) = uint16(v103)
	v119 = F_heap_fetch(m, l0, int32(_a_F_heap_lock_updated_tuple_3), v21+int32(28), v21+int32(24), v100)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L5
	} else {
		goto L21
	}
L8:
	;
	v90 = l2
	goto L7
L9:
	;
	goto L10
L10:
	;
	v43 = F_GetMultiXactIdMembers(m, l2, v21+int32(28), int32(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	if v43 <= int32(0) {
		v90 = v7
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
	v51 = int32(0)
	goto L15
L13:
	;
	F_pfree(m, v48)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L5
	} else {
		goto L19
	}
L14:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v79 = v77
	goto L13
L15:
	;
	v69 = v48 + v51<<(uint(int32(3))%32)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	if base.Ui32(int32(4)) <= base.Ui32(v70) {
		goto L14
	} else {
		goto L17
	}
L16:
	;
	v79 = int32(0)
	goto L13
L17:
	;
	v74 = v51 + int32(1)
	if v74 != v43 {
		v51 = v74
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v90 = v79
	goto L7
L20:
	;
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	if v831 == int32(0) {
		v837 = v814
		goto L1
	} else {
		goto L178
	}
L21:
	;
	if v119 == int32(0) {
		v814 = v100
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v124 = v21 + int32(32)
	v136 = v90
	v141 = v103<<(uint(int32(16))%32) | v104
	v143 = v7
	goto L25
L23:
	;
	F_UnlockReleaseBuffer(m, v146)
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L5
	} else {
		goto L177
	}
L24:
	;
	v794 = int32(0)
	goto L23
L25:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
	v148 = v146 << (uint(int32(13)) % 32)
	v152 = (v146 ^ int32(-1)) << (uint(int32(2)) % 32)
	goto L29
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L27:
	;
	goto L26
L28:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v21)+44))
	v613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v612)+21)))
	if v613&int32(8) != 0 {
		goto L24
	} else {
		goto L145
	}
L29:
	;
	v172 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_updated_tuple[0]))
	if v172 != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	switch v574 - int32(9) {
	case 0, 1, 2:
		goto L27
	case 3:
		v609 = v143
		goto L28
	default:
		v794 = v573
		goto L23
	}
L31:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L5
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v175 = int32(0)
	v176 = base.B2i32(v175 <= v146)
	if v176 == v175 {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	goto L33
L35:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v21)+44))
	v225 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v224)+20)))
	v227 = v225 & int32(768)
	if v136 != 0 {
		goto L52
	} else {
		goto L53
	}
L36:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+10)))
	if v189&int32(4) == int32(0) {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	v180 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_updated_tuple[1]))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v180+v152)))
	v188 = v182
	goto L36
L38:
	;
	goto L39
L39:
	;
	v184 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_updated_tuple[2]))
	v188 = v184 + v148 + int32(-8192)
	goto L36
L40:
	;
	F_LockBuffer(m, v146, int32(2))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L5
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	F_visibilitymap_pin(m, l0, v141, v21+int32(12))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L5
	} else {
		goto L50
	}
L43:
	;
	if v176 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+10)))
	if v209&int32(4) == int32(0) {
		goto L35
	} else {
		goto L48
	}
L45:
	;
	v200 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_updated_tuple[1]))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v200+v152)))
	v208 = v202
	goto L44
L46:
	;
	goto L47
L47:
	;
	v204 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_updated_tuple[2]))
	v208 = v204 + v148 + int32(-8192)
	goto L44
L48:
	;
	F_LockBuffer(m, v146, int32(0))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L5
	} else {
		goto L49
	}
L49:
	;
	goto L42
L50:
	;
	F_LockBuffer(m, v146, int32(2))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L5
	} else {
		goto L51
	}
L51:
	;
	goto L35
L52:
	;
	if v227 != int32(768) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	goto L54
L54:
	;
	if v227 != int32(768) {
		goto L59
	} else {
		goto L60
	}
L55:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
	v232 = v230
	goto L57
L56:
	;
	v232 = int32(2)
	goto L57
L57:
	;
	if v232 != v136 {
		goto L24
	} else {
		goto L58
	}
L58:
	;
	goto L54
L59:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
	v238 = v236
	goto L61
L60:
	;
	v238 = int32(2)
	goto L61
L61:
	;
	v239 = F_TransactionIdDidAbort(m, v238)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L5
	} else {
		goto L62
	}
L62:
	;
	if v239 != 0 {
		goto L24
	} else {
		goto L63
	}
L63:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v21)+44))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v241)+4))
	v243 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v241)+18)))
	v244 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v241)+20)))
	if v244&int32(2048) == int32(0) {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	if v574 == int32(5) {
		goto L29
	} else {
		goto L144
	}
L65:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	F_pfree(m, v569)
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L5
	} else {
		goto L143
	}
L66:
	;
	if v244&int32(_a_F_heap_lock_updated_tuple_2) != 0 {
		goto L70
	} else {
		goto L71
	}
L67:
	;
	v415 = v243
	goto L68
L68:
	;
	F_compute_new_xmax_infomask(m, v242, v244, v415&int32(_a_F_heap_lock_updated_tuple_1), l4, l5, int32(0), v21+int32(16), v21+int32(22), v21+int32(20))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L5
	} else {
		goto L118
	}
L69:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v21)+44))
	v413 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v412)+18)))
	v415 = v413
	goto L68
L70:
	;
	v256 = F_GetMultiXactIdMembers(m, v242, v21, int32(base.Ui32(v244&int32(128))>>(uint(int32(7))%32)))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L5
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v335 = int32(0)
	if base.B2i32(v244&int32(128) == v335)&base.B2i32(v244&int32(80) != int32(64)) == v335 {
		goto L96
	} else {
		goto L97
	}
L73:
	;
	if int32(0) < v256 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v262 = int32(0)
	goto L77
L75:
	;
	goto L76
L76:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if v328 == int32(0) {
		goto L69
	} else {
		goto L92
	}
L77:
	;
	v279 = v262 << (uint(int32(3)) % 32)
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v281 = v279 + v280
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v281)+4))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v281)))
	v288 = F_test_lockmode_for_conflict(m, v282, v283, l5, v21+int32(28), v21+int32(11))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L5
	} else {
		goto L79
	}
L78:
	;
	goto L76
L79:
	;
	if v288 == int32(2) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v568 = int32(12)
	goto L65
L81:
	;
	goto L82
L82:
	;
	v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+11)))
	if v293 == int32(1) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	F_LockBuffer(m, v146, int32(0))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L5
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	if v288 != 0 {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v299+v279)))
	F_XactLockTableWait(m, v301, l0, v124, int32(4))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L5
	} else {
		goto L87
	}
L87:
	;
	v568 = int32(5)
	goto L65
L88:
	;
	v568 = int32(8)
	goto L65
L89:
	;
	goto L90
L90:
	;
	v308 = v262 + int32(1)
	if v308 != v256 {
		v262 = v308
		goto L77
	} else {
		goto L91
	}
L91:
	;
	goto L78
L92:
	;
	F_pfree(m, v328)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L5
	} else {
		goto L93
	}
L93:
	;
	goto L69
L94:
	;
	v380 = F_test_lockmode_for_conflict(m, v375, v242, l5, v21+int32(28), v21+int32(11))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L5
	} else {
		goto L110
	}
L95:
	;
	v375 = int32(1)
	goto L94
L96:
	;
	switch int32(base.Ui32(v244)>>(uint(int32(4))%32))&int32(5) - int32(1) {
	case 0:
		v375 = int32(0)
		goto L94
	case 1, 2:
		goto L27
	case 3:
		goto L100
	case 4:
		goto L95
	default:
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	if v243&int32(_a_F_heap_lock_updated_tuple_4) != 0 {
		goto L107
	} else {
		goto L108
	}
L99:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L5
	} else {
		goto L104
	}
L100:
	;
	if v243&int32(_a_F_heap_lock_updated_tuple_4) != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v355 = int32(3)
	goto L103
L102:
	;
	v355 = int32(2)
	goto L103
L103:
	;
	v375 = v355
	goto L94
L104:
	;
	F_errmsg_internal(m, int32(_a_F_heap_lock_updated_tuple_5), int32(0))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L5
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(_a_F_heap_lock_updated_tuple_6), int32(_a_F_heap_lock_updated_tuple_7), int32(_a_F_heap_lock_updated_tuple_8))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L5
	} else {
		goto L106
	}
L106:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L107:
	;
	v373 = int32(5)
	goto L109
L108:
	;
	v373 = int32(4)
	goto L109
L109:
	;
	v375 = v373
	goto L94
L110:
	;
	if v380 == int32(2) {
		v609 = v143
		goto L28
	} else {
		goto L111
	}
L111:
	;
	v384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+11)))
	if v384 == int32(1) {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	F_LockBuffer(m, v146, int32(0))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L5
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	if v380 != 0 {
		v794 = v380
		goto L23
	} else {
		goto L117
	}
L115:
	;
	F_XactLockTableWait(m, v242, l0, v124, int32(4))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L5
	} else {
		goto L116
	}
L116:
	;
	v573 = v380
	v574 = int32(5)
	goto L64
L117:
	;
	goto L69
L118:
	;
	if v176 == int32(0) {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	v455 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v454)+10)))
	if v455&int32(4) != 0 {
		goto L123
	} else {
		goto L124
	}
L120:
	;
	v446 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_updated_tuple[1]))
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v446+v152)))
	v454 = v448
	goto L119
L121:
	;
	goto L122
L122:
	;
	v450 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_updated_tuple[2]))
	v454 = v450 + v148 + int32(-8192)
	goto L119
L123:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v460 = F_visibilitymap_clear(m, v141, v458, int32(2))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L5
	} else {
		goto L126
	}
L124:
	;
	v463 = v143
	goto L125
L125:
	;
	v464 = int32(_a_F_heap_lock_updated_tuple_9)
	v466 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_updated_tuple[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_lock_updated_tuple[3])) = v466 + int32(1)
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v21)+44))
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v470)+4)) = v471
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v21)+44))
	v474 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v473)+20)))
	v476 = v474 & int32(_a_F_heap_lock_updated_tuple_10)
	*(*uint16)(unsafe.Add(mBase, uint32(v473)+20)) = uint16(v476)
	v478 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v473)+18)))
	v480 = v478 & int32(_a_F_heap_lock_updated_tuple_11)
	*(*uint16)(unsafe.Add(mBase, uint32(v473)+18)) = uint16(v480)
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v21)+44))
	v483 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+22)))
	v484 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v482)+20)))
	v485 = v483 | v484
	*(*uint16)(unsafe.Add(mBase, uint32(v482)+20)) = uint16(v485)
	v487 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+20)))
	v488 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v482)+18)))
	v489 = v487 | v488
	*(*uint16)(unsafe.Add(mBase, uint32(v482)+18)) = uint16(v489)
	F_MarkBufferDirty(m, v146)
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L5
	} else {
		goto L127
	}
L126:
	;
	v463 = v460 | v143
	goto L125
L127:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v494 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+118)))
	if v494 != int32(112) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v562 = int32(_a_F_heap_lock_updated_tuple_9)
	v564 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_updated_tuple[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_lock_updated_tuple[3])) = v564 - int32(1)
	v609 = v463
	goto L28
L129:
	;
	v498 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_updated_tuple[4]))
	if v498 <= int32(0) {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v501 != 0 {
		goto L128
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	if v176 == int32(0) {
		goto L136
	} else {
		goto L137
	}
L133:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v502 != 0 {
		goto L128
	} else {
		goto L134
	}
L134:
	;
	goto L132
L135:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L5
	} else {
		goto L139
	}
L136:
	;
	v506 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_updated_tuple[1]))
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v506+v152)))
	v514 = v508
	goto L135
L137:
	;
	goto L138
L138:
	;
	v510 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_updated_tuple[2]))
	v514 = v510 + v148 + int32(-8192)
	goto L135
L139:
	;
	F_XLogRegisterBuffer(m, int32(0), v146, int32(8))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L5
	} else {
		goto L140
	}
L140:
	;
	v521 = int32(1)
	v522 = v463 & v521
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+7)) = uint8(v522)
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v471
	v525 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+36)))
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+4)) = uint16(v525)
	v533 = int32(8)
	v535 = int32(4)
	v550 = int32(base.Ui32(v487)>>(uint(int32(9))%32))&int32(16) | (int32(base.Ui32(v483)>>(uint(v521)%32))&v533 | (int32(base.Ui32(v483)>>(uint(v535)%32))&v535 | (int32(base.Ui32(v483)>>(uint(int32(12))%32))&v521 | int32(base.Ui32(v483)>>(uint(int32(6))%32))&int32(2))))
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+6)) = uint8(v550)
	F_XLogRegisterData(m, v21, v533)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L5
	} else {
		goto L141
	}
L141:
	;
	v557 = F_XLogInsert(m, int32(9), int32(96))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L5
	} else {
		goto L142
	}
L142:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v514))) = base.I64_rotr(v557, int64(32))
	goto L128
L143:
	;
	v573 = v288
	v574 = v568
	goto L64
L144:
	;
	goto L30
L145:
	;
	v616 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v612)+16)))
	if v616 == int32(_a_F_heap_lock_updated_tuple_0) {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v619 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v612)+12)))
	v620 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v612)+14)))
	if v619&v620 == int32(_a_F_heap_lock_updated_tuple_1) {
		goto L24
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	v625 = v612 + int32(12)
	v626 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v124)+2)))
	v627 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v124))))
	v628 = int32(16)
	v631 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v625)+2)))
	v632 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v625))))
	if v626|v627<<(uint(v628)%32) == v631|v632<<(uint(v628)%32) {
		goto L152
	} else {
		goto L153
	}
L149:
	;
	goto L148
L150:
	;
	if v642 != 0 {
		goto L24
	} else {
		goto L156
	}
L151:
	;
	goto L150
L152:
	;
	v638 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v124)+4)))
	v639 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v625)+4)))
	if v638 == v639 {
		v642 = int32(1)
		goto L151
	} else {
		goto L155
	}
L153:
	;
	goto L154
L154:
	;
	v642 = int32(0)
	goto L151
L155:
	;
	goto L154
L156:
	;
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v21)+44))
	v645 = F_HeapTupleHeaderIsOnlyLocked(m, v644)
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L5
	} else {
		goto L157
	}
L157:
	;
	if v645 != 0 {
		v794 = int32(0)
		goto L23
	} else {
		goto L158
	}
L158:
	;
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v21)+44))
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v647)+4))
	v649 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v647)+20)))
	if v649&int32(_a_F_heap_lock_updated_tuple_12) != int32(_a_F_heap_lock_updated_tuple_2) {
		goto L160
	} else {
		goto L161
	}
L159:
	;
	v732 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v716)+12)))
	v733 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v716)+14)))
	v734 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v716)+16)))
	F_UnlockReleaseBuffer(m, v146)
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L5
	} else {
		goto L174
	}
L160:
	;
	v716 = v647
	v722 = v648
	goto L159
L161:
	;
	goto L162
L162:
	;
	v654 = int32(0)
	v656 = F_GetMultiXactIdMembers(m, v648, v21, v654)
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L5
	} else {
		goto L163
	}
L163:
	;
	if int32(0) < v656 {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v664 = int32(0)
	goto L169
L165:
	;
	v703 = v654
	goto L166
L166:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v21)+44))
	v716 = v713
	v722 = v703
	goto L159
L167:
	;
	F_pfree(m, v661)
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L5
	} else {
		goto L173
	}
L168:
	;
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v682)))
	v692 = v690
	goto L167
L169:
	;
	v682 = v661 + v664<<(uint(int32(3))%32)
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v682)+4))
	if base.Ui32(int32(4)) <= base.Ui32(v683) {
		goto L168
	} else {
		goto L171
	}
L170:
	;
	v692 = int32(0)
	goto L167
L171:
	;
	v687 = v664 + int32(1)
	if v687 != v656 {
		v664 = v687
		goto L169
	} else {
		goto L172
	}
L172:
	;
	goto L170
L173:
	;
	v703 = v692
	goto L166
L174:
	;
	v737 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v737
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+22)) = uint16(v737)
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+36)) = uint16(v734)
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+34)) = uint16(v733)
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+32)) = uint16(v732)
	v754 = F_heap_fetch(m, l0, int32(_a_F_heap_lock_updated_tuple_3), v21+int32(28), v21+int32(24), v737)
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L5
	} else {
		goto L175
	}
L175:
	;
	if v754 != 0 {
		v136 = v722
		v141 = v733 | v732<<(uint(int32(16))%32)
		v143 = v609
		goto L25
	} else {
		goto L176
	}
L176:
	;
	v814 = v737
	goto L20
L177:
	;
	v814 = v794
	goto L20
L178:
	;
	F_ReleaseBuffer(m, v831)
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L5
	} else {
		goto L179
	}
L179:
	;
	v837 = v814
	goto L1
}
func F_heap_mask(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	v2 = l1
	v3 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)) = uint16(v3)
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(0)
	v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)))
	v17 = v15 & int32(_a_F_heap_mask_0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)) = uint16(v17)
	F_mask_unused_space(m, l0)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return
	} else {
		v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
		v28 = int32(0)
		if base.B2i32(base.Ui32(v21) < base.Ui32(int32(25)))|base.B2i32((v21+int32(_a_F_heap_mask_1))&int32(_a_F_heap_mask_2) == v28) == v28 {
			v34 = int32(base.Ui32(v2) >> (uint(int32(16)) % 32))
			v43 = int32(1)
			for {
				v50 = l0 + int32(20) + v43&int32(_a_F_heap_mask_3)<<(uint(int32(2))%32)
				v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
				v54 = l0 + v51&int32(_a_F_heap_mask_4)
				if v51&int32(_a_F_heap_mask_5) != int32(_a_F_heap_mask_6) {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v54)+8)) = int32(0)
					v61 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v54)+20)))
					v64 = int32(768)
					if v61&v64 == v64 {
						v68 = int32(-3073)
					} else {
						v68 = int32(15)
					}
					v69 = v61 & v68
					*(*uint16)(unsafe.Add(mBase, uint32(v54)+20)) = uint16(v69)
					v71 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v54)+16)))
					if v71 != int32(_a_F_heap_mask_7) {
					} else {
						*(*uint16)(unsafe.Add(mBase, uint32(v54)+16)) = uint16(v43)
						*(*uint16)(unsafe.Add(mBase, uint32(v54)+14)) = uint16(v2)
						*(*uint16)(unsafe.Add(mBase, uint32(v54)+12)) = uint16(v34)
					}
				}
				v78 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
				v80 = int32(base.Ui32(v78) >> (uint(int32(17)) % 32))
				if v80 == int32(0) {
				} else {
					v87 = (v80+int32(7))&int32(_a_F_heap_mask_0) - v80
					v88 = int32(0)
					if base.B2i32(v87 <= v88)|base.B2i32(v87 == v88) != 0 {
					} else {
						base.MemoryFill(m, v54+v80, int32(0), v87)
					}
				}
				v98 = v43 + int32(1)
				v101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
				if base.Ui32(int32(25)) <= base.Ui32(v101) {
					v109 = int32(base.Ui32(v101+int32(_a_F_heap_mask_1)) >> (uint(int32(2)) % 32))
				} else {
					v109 = int32(0)
				}
				if base.Ui32(v98&int32(_a_F_heap_mask_3)) <= base.Ui32(v109&int32(_a_F_heap_mask_3)) {
					v43 = v98
					continue
				} else {
					break
				}
				break
			}
		} else {
		}
		return
	}
}
func F_heap_page_prune_opt(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
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
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v109 int32
	_ = v109
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v186 int32
	_ = v186
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int64
	_ = v241
	var v250 int32
	_ = v250
	v7 = m.G0
	v9 = v7 - int32(624)
	m.G0 = v9
	if l1 < int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_page_prune_opt[0])))
	if v31 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_heap_page_prune_opt[1]))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v14+(l1^int32(-1))<<(uint(int32(2))%32))))
	v28 = v20
	goto L1
L3:
	;
	goto L4
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_heap_page_prune_opt[2]))
	v28 = v22 + l1<<(uint(int32(13))%32) + int32(-8192)
	goto L1
L5:
	;
	m.G0 = v9 + int32(624)
	return
L6:
	;
	if v41 != 0 {
		goto L5
	} else {
		goto L10
	}
L7:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_heap_page_prune_opt[3]))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+316))
	v39 = base.B2i32(v37 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_heap_page_prune_opt[0])) = uint8(v39)
	v41 = v39
	goto L9
L8:
	;
	v41 = int32(0)
	goto L9
L9:
	;
	goto L6
L10:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v28)+20))
	if v42 == int32(0) {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v45 = F_GlobalVisHorizonKindForRel(m, l0)
	mBase = m.M
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v45<<(uint(int32(2))%32))+uint32(_c_F_heap_page_prune_opt[4])))
	goto L12
L12:
	;
	v49 = F_GlobalVisTestIsRemovableXid(m, v48, v42)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return
L14:
	;
	if v49 == int32(0) {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v53 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v54 = int32(819)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	v61 = base.I32_div_s(int32(_a_F_heap_page_prune_opt_0)-v56<<(uint(int32(13))%32), int32(100))
	if base.Ui32(v61) <= base.Ui32(v54) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v67 = int32(819)
	goto L18
L18:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+10)))
	if v68&int32(2) == int32(0) {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	v64 = v54
	goto L21
L20:
	;
	v64 = v61
	goto L21
L21:
	;
	v67 = v64
	goto L18
L22:
	;
	v76 = int32(4)
	v77 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+14)))
	v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+12)))
	v79 = v77 - v78
	if v79 <= v76 {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	goto L24
L24:
	;
	v141 = F_ConditionalLockBufferForCleanup(m, l1)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L13
	} else {
		goto L45
	}
L25:
	;
	if base.Ui32(v67) <= base.Ui32(v139) {
		goto L5
	} else {
		goto L44
	}
L26:
	;
	v82 = v76
	goto L28
L27:
	;
	v82 = v79
	goto L28
L28:
	;
	v84 = v82 - int32(4)
	if v84 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v139 = int32(0)
	goto L25
L30:
	;
	goto L31
L31:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v78) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v139 = v84
	goto L25
L33:
	;
	v95 = int32(base.Ui32(v78+int32(_a_F_heap_page_prune_opt_1)) >> (uint(int32(2)) % 32))
	goto L35
L34:
	;
	v95 = int32(0)
	goto L35
L35:
	;
	if base.Ui32(v95&int32(_a_F_heap_page_prune_opt_2)) < base.Ui32(int32(291)) {
		goto L32
	} else {
		goto L36
	}
L36:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+10)))
	if v100&int32(1) == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v139 = int32(0)
	goto L25
L38:
	;
	goto L39
L39:
	;
	v109 = int32(1)
	goto L40
L40:
	;
	v118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28+int32(20)+v109&int32(_a_F_heap_page_prune_opt_2)<<(uint(int32(2))%32))+1)))
	if v118&int32(384) == int32(0) {
		goto L32
	} else {
		goto L42
	}
L41:
	;
	v139 = int32(0)
	goto L25
L42:
	;
	v124 = v109 + int32(1)
	v125 = int32(_a_F_heap_page_prune_opt_2)
	if base.Ui32(v124&v125) <= base.Ui32(v95&v125) {
		v109 = v124
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	goto L24
L45:
	;
	if v141 == int32(0) {
		goto L5
	} else {
		goto L46
	}
L46:
	;
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+10)))
	if v145&int32(2) == int32(0) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	F_LockBuffer(m, l1, int32(0))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L13
	} else {
		goto L79
	}
L48:
	;
	v153 = int32(4)
	v154 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+14)))
	v155 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+12)))
	v156 = v154 - v155
	if v156 <= v153 {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	goto L50
L50:
	;
	v218 = int32(0)
	F_heap_page_prune_and_freeze(m, l0, l1, v48, v218, v218, v9, v218, v9+int32(622), v218, v218)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L13
	} else {
		goto L71
	}
L51:
	;
	if base.Ui32(v67) <= base.Ui32(v216) {
		goto L47
	} else {
		goto L70
	}
L52:
	;
	v159 = v153
	goto L54
L53:
	;
	v159 = v156
	goto L54
L54:
	;
	v161 = v159 - int32(4)
	if v161 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v216 = int32(0)
	goto L51
L56:
	;
	goto L57
L57:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v155) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v216 = v161
	goto L51
L59:
	;
	v172 = int32(base.Ui32(v155+int32(_a_F_heap_page_prune_opt_1)) >> (uint(int32(2)) % 32))
	goto L61
L60:
	;
	v172 = int32(0)
	goto L61
L61:
	;
	if base.Ui32(v172&int32(_a_F_heap_page_prune_opt_2)) < base.Ui32(int32(291)) {
		goto L58
	} else {
		goto L62
	}
L62:
	;
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+10)))
	if v177&int32(1) == int32(0) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v216 = int32(0)
	goto L51
L64:
	;
	goto L65
L65:
	;
	v186 = int32(1)
	goto L66
L66:
	;
	v195 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28+int32(20)+v186&int32(_a_F_heap_page_prune_opt_2)<<(uint(int32(2))%32))+1)))
	if v195&int32(384) == int32(0) {
		goto L58
	} else {
		goto L68
	}
L67:
	;
	v216 = int32(0)
	goto L51
L68:
	;
	v201 = v186 + int32(1)
	v202 = int32(_a_F_heap_page_prune_opt_2)
	if base.Ui32(v201&v202) <= base.Ui32(v172&v202) {
		v186 = v201
		goto L66
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	goto L50
L71:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v227 <= v228 {
		goto L47
	} else {
		goto L72
	}
L72:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v231 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	goto L47
L74:
	;
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+268)))
	if v234 != int32(1) {
		goto L73
	} else {
		goto L77
	}
L75:
	;
	v240 = v231
	goto L76
L76:
	;
	v241 = *(*int64)(unsafe.Add(mBase, uint32(v240)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v240)+96)) = v241 - base.I64_extend_i32_s(v227-v228)
	goto L73
L77:
	;
	F_pgstat_assoc_relation(m, l0)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L13
	} else {
		goto L78
	}
L78:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	v240 = v239
	goto L76
L79:
	;
	goto L5
}
