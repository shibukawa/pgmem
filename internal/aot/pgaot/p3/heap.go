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
		v14 = *(*int32)(unsafe.Add(mBase, _consts[33]))
		if v14 == l3 {
			v17 = *(*int64)(unsafe.Add(mBase, _consts[34]))
			v26 = v17
			m.G0 = v11 + int32(16)
			if int32(0) <= l1 {
				v33 = *(*int32)(unsafe.Add(mBase, _consts[7]))
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
					v33 = *(*int32)(unsafe.Add(mBase, _consts[7]))
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
						v33 = *(*int32)(unsafe.Add(mBase, _consts[7]))
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
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
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
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v215 int32
	_ = v215
	var v232 int32
	_ = v232
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
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
	if v215 < v21 {
		goto L61
	} else {
		goto L62
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
	v215 = int32(0)
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
	v46 = v5
	v47 = v5
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
	v215 = v27
	goto L1
L10:
	;
	v206 = v41 + int32(1)
	if v206 != v27 {
		v41 = v206
		v46 = v204
		v47 = v201
		goto L8
	} else {
		goto L60
	}
L11:
	;
	v79 = l1 + int32(20) + v41<<(uint(int32(4))%32)
	v81 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v41+l3))) = uint8(v81)
	if v46&int32(1) == v81 {
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
	v201 = v47
	v204 = v74
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
	if v103 != int32(65535) {
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
	if v98 != int32(65535) {
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
		v196 = v160
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
	F_errmsg_internal(m, int32(504261), v19)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(341792), int32(70), int32(73498))
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
	v201 = v196 + v128
	v204 = v129 | v162
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
	v193 = F_strlen(m, v130)
	mBase = m.M
	v196 = v193 + int32(1)
	goto L40
L45:
	;
	v170 = int32(6)
	v172 = int32(18)
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+1)))
	if v174 == v172 {
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
		goto L57
	} else {
		goto L58
	}
L48:
	;
	v177 = v172
	goto L50
L49:
	;
	v177 = int32(2)
	goto L50
L50:
	;
	if v174&int32(254) == int32(2) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v182 = v170
	goto L53
L52:
	;
	v182 = v177
	goto L53
L53:
	;
	if v174 == int32(1) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v185 = v170
	goto L56
L55:
	;
	v185 = v182
	goto L56
L56:
	;
	v196 = v185
	goto L40
L57:
	;
	v196 = int32(base.Ui32(v167) >> (uint(int32(1)) % 32))
	goto L40
L58:
	;
	goto L59
L59:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	v196 = int32(base.Ui32(v190) >> (uint(int32(2)) % 32))
	goto L40
L60:
	;
	goto L9
L61:
	;
	v232 = v215
	goto L64
L62:
	;
	goto L63
L63:
	;
	m.G0 = v19 + int32(16)
	return
L64:
	;
	v246 = v232 + int32(1)
	v247 = F_getmissingattr(m, l1, v246, l3+v232)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L36
	} else {
		goto L66
	}
L65:
	;
	goto L63
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2+v232<<(uint(int32(2))%32)))) = v247
	if v246 != v21 {
		v232 = v246
		goto L64
	} else {
		goto L67
	}
L67:
	;
	goto L65
}
func F_heap_getattr_2(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if int32(0) < l1 {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+18)))
		if base.Ui32(v15&int32(2047)) < base.Ui32(l1) {
			v19 = F_getmissingattr(m, l2, l1, l3)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v93 = v19
				m.G0 = v10 + int32(16)
				return v93
			}
		} else {
			v23 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v23)
			v25 = int32(1)
			v26 = l1 - v25
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+20)))
			if v28&v25 == v23 {
				v37 = l2 + v26<<(uint(int32(4))%32) + int32(20)
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
				if int32(0) <= v38 {
					v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+22)))
					v43 = v27 + v41 + v38
					v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+6)))
					if v44 != int32(1) {
						v93 = v43
						m.G0 = v10 + int32(16)
						return v93
					} else {
						v47 = int32(*(*int16)(unsafe.Add(mBase, uint32(v37)+4)))
						switch v47&int32(65535) - int32(1) {
						case 0:
							v52 = int32(*(*int8)(unsafe.Add(mBase, uint32(v43))))
							v93 = v52
							m.G0 = v10 + int32(16)
							return v93
						case 1:
							v53 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43))))
							v93 = v53
							m.G0 = v10 + int32(16)
							return v93
						default:
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10))) = v47
								F_errmsg_internal(m, int32(504261), v10)
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(341795), int32(70), int32(73498))
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						case 3:
							v54 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
							v93 = v54
							m.G0 = v10 + int32(16)
							return v93
						}
					}
				} else {
					v68 = F_nocachegetattr(m, l0, l1, l2)
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return int32(0)
					} else {
						v93 = v68
						m.G0 = v10 + int32(16)
						return v93
					}
				}
			} else {
				v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+int32(base.Ui32(v26)>>(uint(int32(3))%32)))+23)))
				if int32(base.Ui32(v73)>>(uint(v26&int32(7))%32))&int32(1) == int32(0) {
					v81 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v81)
					v93 = int32(0)
					m.G0 = v10 + int32(16)
					return v93
				} else {
					v84 = F_nocachegetattr(m, l0, l1, l2)
					mBase = m.M
					v85 = m.ExcPending
					if v85 != 0 {
						return int32(0)
					} else {
						v93 = v84
						m.G0 = v10 + int32(16)
						return v93
					}
				}
			}
		}
	} else {
		v86 = F_heap_getsysattr(m, l0, l1, l3)
		mBase = m.M
		v87 = m.ExcPending
		if v87 != 0 {
			return int32(0)
		} else {
			v93 = v86
			m.G0 = v10 + int32(16)
			return v93
		}
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
					switch v38&int32(65535) - int32(1) {
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
							F_errmsg_internal(m, int32(504261), v8)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(341795), int32(70), int32(73498))
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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	if base.Ui32(l0) <= base.Ui32(int32(207)) {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(int32(base.Ui32(l0)>>(uint(int32(2))%32))&int32(60))+uint32(_consts[79])))
		v12 = v11
	} else {
		v12 = int32(0)
	}
	return v12
}
func F_heap_lock_updated_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v93 int32
	_ = v93
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
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
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v566 int32
	_ = v566
	var v570 int32
	_ = v570
	var v573 int64
	_ = v573
	var v574 int32
	_ = v574
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v622 int32
	_ = v622
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v707 int32
	_ = v707
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v714 int32
	_ = v714
	var v724 int32
	_ = v724
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v744 int32
	_ = v744
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v799 int32
	_ = v799
	var v818 int32
	_ = v818
	var v820 int32
	_ = v820
	var v838 int32
	_ = v838
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	v7 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(48)
	m.G0 = v22
	v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+4)))
	if v24 != int32(65533) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v22 + int32(48)
	return v844
L2:
	;
	F_MultiXactIdSetOldestMember(m)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3))))
	v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+2)))
	if v27&v28 != int32(65535) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v844 = int32(0)
	goto L1
L5:
	;
	return int32(0)
L6:
	;
	if l1&int32(4096) == int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v103 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v103
	v106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3))))
	v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+2)))
	v108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v103
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+22)) = uint16(v103)
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+36)) = uint16(v108)
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+34)) = uint16(v107)
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+32)) = uint16(v106)
	v122 = F_heap_fetch(m, l0, int32(4211808), v22+int32(28), v22+int32(24), v103)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L5
	} else {
		goto L21
	}
L8:
	;
	v93 = l2
	goto L7
L9:
	;
	goto L10
L10:
	;
	v44 = F_GetMultiXactIdMembers(m, l2, v22+int32(28), int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	if v44 <= int32(0) {
		v93 = v7
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	v52 = int32(0)
	goto L15
L13:
	;
	F_pfree(m, v49)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L5
	} else {
		goto L19
	}
L14:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	v81 = v79
	goto L13
L15:
	;
	v71 = v49 + v52<<(uint(int32(3))%32)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	if base.Ui32(int32(4)) <= base.Ui32(v72) {
		goto L14
	} else {
		goto L17
	}
L16:
	;
	v81 = int32(0)
	goto L13
L17:
	;
	v76 = v52 + int32(1)
	if v76 != v44 {
		v52 = v76
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v93 = v81
	goto L7
L20:
	;
	v838 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	if v838 == int32(0) {
		v844 = v820
		goto L1
	} else {
		goto L178
	}
L21:
	;
	if v122 == int32(0) {
		v820 = v103
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v127 = v22 + int32(32)
	v140 = v93
	v141 = v7
	v146 = v106<<(uint(int32(16))%32) | v107
	goto L24
L23:
	;
	F_UnlockReleaseBuffer(m, v150)
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L5
	} else {
		goto L177
	}
L24:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	v152 = v150 << (uint(int32(13)) % 32)
	v156 = (v150 ^ int32(-1)) << (uint(int32(2)) % 32)
	goto L28
L25:
	;
	v799 = int32(0)
	goto L23
L26:
	;
	goto L25
L27:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v22)+44))
	v632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v631)+21)))
	if v632&int32(8) != 0 {
		goto L26
	} else {
		goto L145
	}
L28:
	;
	v177 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	if v177 != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	switch v591 - int32(8) {
	case 0:
		v799 = v590
		goto L23
	default:
		v844 = v590
		goto L1
	case 4:
		v622 = v141
		goto L27
	}
L30:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L5
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v180 = int32(0)
	v181 = base.B2i32(v180 <= v150)
	if v181 == v180 {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	goto L32
L34:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v22)+44))
	v230 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v229)+20)))
	v232 = v230 & int32(768)
	if v140 != 0 {
		goto L51
	} else {
		goto L52
	}
L35:
	;
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+10)))
	if v194&int32(4) == int32(0) {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	v185 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v185+v156)))
	v193 = v187
	goto L35
L37:
	;
	goto L38
L38:
	;
	v189 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v193 = v189 + v152 + int32(-8192)
	goto L35
L39:
	;
	F_LockBuffer(m, v150, int32(2))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L5
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	F_visibilitymap_pin(m, l0, v146, v22+int32(12))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L5
	} else {
		goto L49
	}
L42:
	;
	if v181 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+10)))
	if v214&int32(4) == int32(0) {
		goto L34
	} else {
		goto L47
	}
L44:
	;
	v205 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v205+v156)))
	v213 = v207
	goto L43
L45:
	;
	goto L46
L46:
	;
	v209 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v213 = v209 + v152 + int32(-8192)
	goto L43
L47:
	;
	F_LockBuffer(m, v150, int32(0))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L5
	} else {
		goto L48
	}
L48:
	;
	goto L41
L49:
	;
	F_LockBuffer(m, v150, int32(2))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L5
	} else {
		goto L50
	}
L50:
	;
	goto L34
L51:
	;
	if v232 != int32(768) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	goto L53
L53:
	;
	if v232 != int32(768) {
		goto L58
	} else {
		goto L59
	}
L54:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
	v237 = v236
	goto L56
L55:
	;
	v237 = int32(2)
	goto L56
L56:
	;
	if v237 != v140 {
		goto L26
	} else {
		goto L57
	}
L57:
	;
	goto L53
L58:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
	v244 = v243
	goto L60
L59:
	;
	v244 = int32(2)
	goto L60
L60:
	;
	v245 = F_TransactionIdDidAbort(m, v244)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L5
	} else {
		goto L61
	}
L61:
	;
	if v245 != 0 {
		goto L26
	} else {
		goto L62
	}
L62:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v22)+44))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v247)+4))
	v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v247)+18)))
	v250 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v247)+20)))
	if v250&int32(2048) == int32(0) {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	if v591 == int32(5) {
		goto L28
	} else {
		goto L144
	}
L64:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	F_pfree(m, v586)
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L5
	} else {
		goto L143
	}
L65:
	;
	if v250&int32(4096) != 0 {
		goto L69
	} else {
		goto L70
	}
L66:
	;
	v429 = v249
	goto L67
L67:
	;
	F_compute_new_xmax_infomask(m, v248, v250, v429&int32(65535), l4, l5, int32(0), v22+int32(16), v22+int32(22), v22+int32(20))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L5
	} else {
		goto L118
	}
L68:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v22)+44))
	v427 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v426)+18)))
	v429 = v427
	goto L67
L69:
	;
	v267 = F_GetMultiXactIdMembers(m, v248, v22, int32(base.Ui32(v250&int32(128))>>(uint(int32(7))%32))|base.B2i32(v250&int32(4176) == int32(64)))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L5
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v348 = int32(0)
	if base.B2i32(v250&int32(128) == v348)&base.B2i32(v250&int32(80) != int32(64)) == v348 {
		goto L95
	} else {
		goto L96
	}
L72:
	;
	if int32(0) < v267 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v273 = int32(0)
	goto L76
L74:
	;
	goto L75
L75:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	if v341 == int32(0) {
		goto L68
	} else {
		goto L91
	}
L76:
	;
	v291 = v273 << (uint(int32(3)) % 32)
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v293 = v291 + v292
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v293)+4))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v293)))
	v300 = F_test_lockmode_for_conflict(m, v294, v295, l5, v22+int32(28), v22+int32(11))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L5
	} else {
		goto L78
	}
L77:
	;
	goto L75
L78:
	;
	if v300 == int32(2) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v585 = int32(12)
	goto L64
L80:
	;
	goto L81
L81:
	;
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+11)))
	if v305 == int32(1) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	F_LockBuffer(m, v150, int32(0))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L5
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	if v300 != 0 {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v311+v291)))
	F_XactLockTableWait(m, v313, l0, v127, int32(4))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L5
	} else {
		goto L86
	}
L86:
	;
	v585 = int32(5)
	goto L64
L87:
	;
	v585 = int32(8)
	goto L64
L88:
	;
	goto L89
L89:
	;
	v320 = v273 + int32(1)
	if v320 != v267 {
		v273 = v320
		goto L76
	} else {
		goto L90
	}
L90:
	;
	goto L77
L91:
	;
	F_pfree(m, v341)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L5
	} else {
		goto L92
	}
L92:
	;
	goto L68
L93:
	;
	v393 = F_test_lockmode_for_conflict(m, v388, v248, l5, v22+int32(28), v22+int32(11))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L5
	} else {
		goto L110
	}
L94:
	;
	v388 = int32(1)
	goto L93
L95:
	;
	switch int32(base.Ui32(v250)>>(uint(int32(4))%32))&int32(5) - int32(1) {
	case 0:
		v388 = int32(0)
		goto L93
	case 1, 2:
		goto L99
	case 3:
		goto L100
	case 4:
		goto L94
	default:
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	if v249&int32(8192) != 0 {
		goto L107
	} else {
		goto L108
	}
L98:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L5
	} else {
		goto L104
	}
L99:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L100:
	;
	if v249&int32(8192) != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v368 = int32(3)
	goto L103
L102:
	;
	v368 = int32(2)
	goto L103
L103:
	;
	v388 = v368
	goto L93
L104:
	;
	F_errmsg_internal(m, int32(401605), int32(0))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L5
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(520051), int32(5904), int32(513070))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
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
	v386 = int32(5)
	goto L109
L108:
	;
	v386 = int32(4)
	goto L109
L109:
	;
	v388 = v386
	goto L93
L110:
	;
	if v393 == int32(2) {
		v622 = v141
		goto L27
	} else {
		goto L111
	}
L111:
	;
	v397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+11)))
	if v397 == int32(1) {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	F_LockBuffer(m, v150, int32(0))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L5
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	if v393 != 0 {
		v799 = v393
		goto L23
	} else {
		goto L117
	}
L115:
	;
	F_XactLockTableWait(m, v248, l0, v127, int32(4))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L5
	} else {
		goto L116
	}
L116:
	;
	v590 = v393
	v591 = int32(5)
	goto L63
L117:
	;
	goto L68
L118:
	;
	if v181 == int32(0) {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	v470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v469)+10)))
	if v470&int32(4) != 0 {
		goto L123
	} else {
		goto L124
	}
L120:
	;
	v461 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v461+v156)))
	v469 = v463
	goto L119
L121:
	;
	goto L122
L122:
	;
	v465 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v469 = v465 + v152 + int32(-8192)
	goto L119
L123:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v475 = F_visibilitymap_clear(m, v146, v473, int32(2))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L5
	} else {
		goto L126
	}
L124:
	;
	v478 = v141
	goto L125
L125:
	;
	v480 = int32(4548548)
	v482 = *(*int32)(unsafe.Add(mBase, _consts[17]))
	*(*int32)(unsafe.Add(mBase, _consts[17])) = v482 + int32(1)
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v22)+44))
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v486)+4)) = v487
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v22)+44))
	v490 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v489)+20)))
	v492 = v490 & int32(58159)
	*(*uint16)(unsafe.Add(mBase, uint32(v489)+20)) = uint16(v492)
	v494 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v489)+18)))
	v496 = v494 & int32(57343)
	*(*uint16)(unsafe.Add(mBase, uint32(v489)+18)) = uint16(v496)
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v22)+44))
	v499 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+22)))
	v500 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v498)+20)))
	v501 = v499 | v500
	*(*uint16)(unsafe.Add(mBase, uint32(v498)+20)) = uint16(v501)
	v503 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+20)))
	v504 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v498)+18)))
	v505 = v503 | v504
	*(*uint16)(unsafe.Add(mBase, uint32(v498)+18)) = uint16(v505)
	F_MarkBufferDirty(m, v150)
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L5
	} else {
		goto L127
	}
L126:
	;
	v478 = v475 | v141
	goto L125
L127:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v509)+118)))
	if v510 != int32(112) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v579 = int32(4548548)
	v581 = *(*int32)(unsafe.Add(mBase, _consts[17]))
	*(*int32)(unsafe.Add(mBase, _consts[17])) = v581 - int32(1)
	v622 = v478
	goto L27
L129:
	;
	v514 = *(*int32)(unsafe.Add(mBase, _consts[18]))
	if v514 <= int32(0) {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v517 != 0 {
		goto L128
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	if v181 == int32(0) {
		goto L136
	} else {
		goto L137
	}
L133:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v518 != 0 {
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
	v532 = m.ExcPending
	if v532 != 0 {
		goto L5
	} else {
		goto L139
	}
L136:
	;
	v522 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v522+v156)))
	v530 = v524
	goto L135
L137:
	;
	goto L138
L138:
	;
	v526 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v530 = v526 + v152 + int32(-8192)
	goto L135
L139:
	;
	F_XLogRegisterBuffer(m, int32(0), v150, int32(8))
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L5
	} else {
		goto L140
	}
L140:
	;
	v537 = int32(1)
	v538 = v478 & v537
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+7)) = uint8(v538)
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v487
	v541 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+36)))
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+4)) = uint16(v541)
	v549 = int32(8)
	v551 = int32(4)
	v566 = int32(base.Ui32(v503)>>(uint(int32(9))%32))&int32(16) | (int32(base.Ui32(v499)>>(uint(v537)%32))&v549 | (int32(base.Ui32(v499)>>(uint(v551)%32))&v551 | (int32(base.Ui32(v499)>>(uint(int32(12))%32))&v537 | int32(base.Ui32(v499)>>(uint(int32(6))%32))&int32(2))))
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+6)) = uint8(v566)
	F_XLogRegisterData(m, v22, v549)
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L5
	} else {
		goto L141
	}
L141:
	;
	v573 = F_XLogInsert(m, int32(9), int32(96))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L5
	} else {
		goto L142
	}
L142:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v530))) = base.I64_rotr(v573, int64(32))
	goto L128
L143:
	;
	v590 = v300
	v591 = v585
	goto L63
L144:
	;
	goto L29
L145:
	;
	v635 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v631)+16)))
	if v635 == int32(65533) {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v638 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v631)+12)))
	v639 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v631)+14)))
	if v638&v639 == int32(65535) {
		goto L26
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	v644 = v631 + int32(12)
	v645 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v127)+2)))
	v646 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v127))))
	v647 = int32(16)
	v650 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v644)+2)))
	v651 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v644))))
	if v645|v646<<(uint(v647)%32) == v650|v651<<(uint(v647)%32) {
		goto L152
	} else {
		goto L153
	}
L149:
	;
	goto L148
L150:
	;
	if v661 != 0 {
		goto L26
	} else {
		goto L156
	}
L151:
	;
	goto L150
L152:
	;
	v657 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v127)+4)))
	v658 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v644)+4)))
	if v657 == v658 {
		v661 = int32(1)
		goto L151
	} else {
		goto L155
	}
L153:
	;
	goto L154
L154:
	;
	v661 = int32(0)
	goto L151
L155:
	;
	goto L154
L156:
	;
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v22)+44))
	v664 = F_HeapTupleHeaderIsOnlyLocked(m, v663)
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L5
	} else {
		goto L157
	}
L157:
	;
	if v664 != 0 {
		v799 = int32(0)
		goto L23
	} else {
		goto L158
	}
L158:
	;
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v22)+44))
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v666)+4))
	v668 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v666)+20)))
	if v668&int32(6272) != int32(4096) {
		goto L160
	} else {
		goto L161
	}
L159:
	;
	v754 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v737)+12)))
	v755 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v737)+14)))
	v756 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v737)+16)))
	F_UnlockReleaseBuffer(m, v150)
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L5
	} else {
		goto L174
	}
L160:
	;
	v737 = v666
	v744 = v667
	goto L159
L161:
	;
	goto L162
L162:
	;
	v673 = int32(0)
	v675 = F_GetMultiXactIdMembers(m, v667, v22, v673)
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L5
	} else {
		goto L163
	}
L163:
	;
	if int32(0) < v675 {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v683 = int32(0)
	goto L169
L165:
	;
	v724 = v673
	goto L166
L166:
	;
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v22)+44))
	v737 = v734
	v744 = v724
	goto L159
L167:
	;
	F_pfree(m, v680)
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L5
	} else {
		goto L173
	}
L168:
	;
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v702)))
	v712 = v710
	goto L167
L169:
	;
	v702 = v680 + v683<<(uint(int32(3))%32)
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v702)+4))
	if base.Ui32(int32(4)) <= base.Ui32(v703) {
		goto L168
	} else {
		goto L171
	}
L170:
	;
	v712 = int32(0)
	goto L167
L171:
	;
	v707 = v683 + int32(1)
	if v707 != v675 {
		v683 = v707
		goto L169
	} else {
		goto L172
	}
L172:
	;
	goto L170
L173:
	;
	v724 = v712
	goto L166
L174:
	;
	v759 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v759
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+22)) = uint16(v759)
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+36)) = uint16(v756)
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+34)) = uint16(v755)
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+32)) = uint16(v754)
	v776 = F_heap_fetch(m, l0, int32(4211808), v22+int32(28), v22+int32(24), v759)
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
		goto L5
	} else {
		goto L175
	}
L175:
	;
	if v776 != 0 {
		v140 = v744
		v141 = v622
		v146 = v755 | v754<<(uint(int32(16))%32)
		goto L24
	} else {
		goto L176
	}
L176:
	;
	v820 = v759
	goto L20
L177:
	;
	v820 = v799
	goto L20
L178:
	;
	F_ReleaseBuffer(m, v838)
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L5
	} else {
		goto L179
	}
L179:
	;
	v844 = v820
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
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	v2 = l1
	v3 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)) = uint16(v3)
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(0)
	v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)))
	v17 = v15 & int32(65528)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)) = uint16(v17)
	F_mask_unused_space(m, l0)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return
	} else {
		v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
		if base.Ui32(v21) < base.Ui32(int32(25)) {
		} else {
			if (v21+int32(262120))&int32(262140) == int32(0) {
			} else {
				v31 = int32(base.Ui32(v2) >> (uint(int32(16)) % 32))
				v39 = int32(1)
				for {
					v49 = v39&int32(65535)<<(uint(int32(2))%32) + (l0 + int32(24)) - int32(4)
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
					v53 = l0 + v50&int32(32767)
					if v50&int32(98304) == int32(32768) {
						*(*int32)(unsafe.Add(mBase, uint32(v53)+8)) = int32(0)
						v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53)+20)))
						v63 = int32(768)
						if v60&v63 == v63 {
							v67 = int32(-3073)
						} else {
							v67 = int32(15)
						}
						v68 = v60 & v67
						*(*uint16)(unsafe.Add(mBase, uint32(v53)+20)) = uint16(v68)
						v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53)+16)))
						if v70 == int32(65534) {
							*(*uint16)(unsafe.Add(mBase, uint32(v53)+16)) = uint16(v39)
							*(*uint16)(unsafe.Add(mBase, uint32(v53)+14)) = uint16(v2)
							*(*uint16)(unsafe.Add(mBase, uint32(v53)+12)) = uint16(v31)
						} else {
						}
						v76 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
						v77 = v76
					} else {
						v77 = v50
					}
					if base.Ui32(v77) < base.Ui32(int32(131072)) {
					} else {
						v81 = int32(base.Ui32(v77) >> (uint(int32(17)) % 32))
						v86 = (v81+int32(7))&int32(65528) - v81
						if v86 <= int32(0) {
						} else {
							v92 = F__emscripten_memset_bulkmem(m, v53+v81, base.I32_extend8_s(int32(0)), v86)
							mBase = m.M
						}
					}
					v96 = v39 + int32(1)
					v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
					if base.Ui32(int32(25)) <= base.Ui32(v99) {
						v107 = int32(base.Ui32(v99+int32(262120)) >> (uint(int32(2)) % 32))
					} else {
						v107 = int32(0)
					}
					if base.Ui32(v96&int32(65535)) <= base.Ui32(v107&int32(65535)) {
						v39 = v96
						continue
					} else {
						break
					}
					break
				}
			}
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
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v110 int32
	_ = v110
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v189 int32
	_ = v189
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int64
	_ = v246
	var v255 int32
	_ = v255
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
	v31 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1])))
	if v31 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v14+(l1^int32(-1))<<(uint(int32(2))%32))))
	v28 = v20
	goto L1
L3:
	;
	goto L4
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _consts[10]))
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
	v36 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+316))
	v39 = base.B2i32(v37 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _consts[1])) = uint8(v39)
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
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v45<<(uint(int32(2))%32))+uint32(_consts[48])))
	goto L12
L12:
	;
	v51 = F_GlobalVisTestIsRemovableXid(m, v50, v42)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return
L14:
	;
	if v51 == int32(0) {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v56 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	v63 = base.I32_div_s(int32(819200)-v58<<(uint(int32(13))%32), int32(100))
	v65 = v63
	goto L18
L17:
	;
	v65 = int32(0)
	goto L18
L18:
	;
	if base.Ui32(v65) <= base.Ui32(int32(819)) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v68 = int32(819)
	goto L21
L20:
	;
	v68 = v65
	goto L21
L21:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+10)))
	if v69&int32(2) == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v77 = int32(4)
	v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+14)))
	v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+12)))
	v80 = v78 - v79
	if v80 <= v77 {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	goto L24
L24:
	;
	v144 = F_ConditionalLockBufferForCleanup(m, l1)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L13
	} else {
		goto L45
	}
L25:
	;
	if base.Ui32(v68) <= base.Ui32(v142) {
		goto L5
	} else {
		goto L44
	}
L26:
	;
	v83 = v77
	goto L28
L27:
	;
	v83 = v80
	goto L28
L28:
	;
	v85 = v83 - int32(4)
	if v85 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v142 = int32(0)
	goto L25
L30:
	;
	goto L31
L31:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v79) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v142 = v85
	goto L25
L33:
	;
	v96 = int32(base.Ui32(v79+int32(262120)) >> (uint(int32(2)) % 32))
	goto L35
L34:
	;
	v96 = int32(0)
	goto L35
L35:
	;
	if base.Ui32(v96&int32(65535)) < base.Ui32(int32(291)) {
		goto L32
	} else {
		goto L36
	}
L36:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+10)))
	if v101&int32(1) == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v142 = int32(0)
	goto L25
L38:
	;
	goto L39
L39:
	;
	v110 = int32(1)
	goto L40
L40:
	;
	v121 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v110&int32(65535)<<(uint(int32(2))%32)+(v28+int32(24))-int32(3)))))
	if v121&int32(384) == int32(0) {
		goto L32
	} else {
		goto L42
	}
L41:
	;
	v142 = int32(0)
	goto L25
L42:
	;
	v127 = v110 + int32(1)
	v128 = int32(65535)
	if base.Ui32(v127&v128) <= base.Ui32(v96&v128) {
		v110 = v127
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
	if v144 == int32(0) {
		goto L5
	} else {
		goto L46
	}
L46:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+10)))
	if v148&int32(2) == int32(0) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	F_LockBuffer(m, l1, int32(0))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L13
	} else {
		goto L79
	}
L48:
	;
	v156 = int32(4)
	v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+14)))
	v158 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+12)))
	v159 = v157 - v158
	if v159 <= v156 {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	goto L50
L50:
	;
	v223 = int32(0)
	F_heap_page_prune_and_freeze(m, l0, l1, v50, v223, v223, v9, v223, v9+int32(622), v223, v223)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L13
	} else {
		goto L71
	}
L51:
	;
	if base.Ui32(v68) <= base.Ui32(v221) {
		goto L47
	} else {
		goto L70
	}
L52:
	;
	v162 = v156
	goto L54
L53:
	;
	v162 = v159
	goto L54
L54:
	;
	v164 = v162 - int32(4)
	if v164 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v221 = int32(0)
	goto L51
L56:
	;
	goto L57
L57:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v158) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v221 = v164
	goto L51
L59:
	;
	v175 = int32(base.Ui32(v158+int32(262120)) >> (uint(int32(2)) % 32))
	goto L61
L60:
	;
	v175 = int32(0)
	goto L61
L61:
	;
	if base.Ui32(v175&int32(65535)) < base.Ui32(int32(291)) {
		goto L58
	} else {
		goto L62
	}
L62:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+10)))
	if v180&int32(1) == int32(0) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v221 = int32(0)
	goto L51
L64:
	;
	goto L65
L65:
	;
	v189 = int32(1)
	goto L66
L66:
	;
	v200 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v189&int32(65535)<<(uint(int32(2))%32)+(v28+int32(24))-int32(3)))))
	if v200&int32(384) == int32(0) {
		goto L58
	} else {
		goto L68
	}
L67:
	;
	v221 = int32(0)
	goto L51
L68:
	;
	v206 = v189 + int32(1)
	v207 = int32(65535)
	if base.Ui32(v206&v207) <= base.Ui32(v175&v207) {
		v189 = v206
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
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v232 <= v233 {
		goto L47
	} else {
		goto L72
	}
L72:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v236 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	goto L47
L74:
	;
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+268)))
	if v239 != int32(1) {
		goto L73
	} else {
		goto L77
	}
L75:
	;
	v245 = v236
	goto L76
L76:
	;
	v246 = *(*int64)(unsafe.Add(mBase, uint32(v245)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v245)+96)) = v246 - base.I64_extend_i32_s(v232-v233)
	goto L73
L77:
	;
	F_pgstat_assoc_relation(m, l0)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L13
	} else {
		goto L78
	}
L78:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	v245 = v244
	goto L76
L79:
	;
	goto L5
}
