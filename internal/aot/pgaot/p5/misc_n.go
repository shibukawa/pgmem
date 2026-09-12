package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_NonFiniteTimestampTzPart(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) float64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v64 float64
	_ = v64
	var v65 float64
	_ = v65
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	switch l0 {
	case 0, 17:
		switch l1 - int32(4) {
		case 0, 14, 15, 16, 17, 18, 19, 20, 25, 26, 28, 29, 30, 31, 33:
			v65 = float64(0)
			m.G0 = v10 + int32(32)
			return v65
		default:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return float64(0)
			} else {
				F_errcode(m, int32(1088))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return float64(0)
				} else {
					if l4 != 0 {
						v47 = int32(1184)
					} else {
						v47 = int32(1114)
					}
					v48 = F_format_type_be(m, v47)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return float64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v48
						*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l2
						F_errmsg(m, int32(182600), v10+int32(16))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return float64(0)
						} else {
							F_errfinish(m, int32(479085), int32(5490), int32(80044))
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return float64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		case 7, 21, 22, 23, 24, 27, 32:
			if l3 != 0 {
				v64 = math.Float64frombits(uint64(0xfff0000000000000))
			} else {
				v64 = math.Float64frombits(uint64(0x7ff0000000000000))
			}
			v65 = v64
			m.G0 = v10 + int32(32)
			return v65
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return float64(0)
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return float64(0)
			} else {
				if l4 != 0 {
					v23 = int32(1184)
				} else {
					v23 = int32(1114)
				}
				v24 = F_format_type_be(m, v23)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return float64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v24
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = l2
					F_errmsg(m, int32(182563), v10)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return float64(0)
					} else {
						F_errfinish(m, int32(479085), int32(5450), int32(80044))
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return float64(0)
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
func F_NotifyMyFrontEnd(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _consts[281]))
	if v11 == int32(2) {
		F_pq_beginmessage(m, v8+int32(16), int32(65))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			F_enlargeStringInfo(m, v8+int32(16), int32(4))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
				v27 = int32(24)
				v29 = int32(65280)
				v31 = int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(v24+v25))) = l2<<(uint(v27)%32) | l2&v29<<(uint(v31)%32) | (int32(base.Ui32(l2)>>(uint(v31)%32))&v29 | int32(base.Ui32(l2)>>(uint(v27)%32)))
				*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v24 + int32(4)
				F_pq_sendstring(m, v8+int32(16), l0)
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return
				} else {
					F_pq_sendstring(m, v8+int32(16), l1)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return
					} else {
						F_pq_endmessage(m, v8+int32(16))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return
						} else {
							m.G0 = v8 + int32(32)
							return
						}
					}
				}
			}
		}
	} else {
		v60 = F_errstart(m, int32(17), int32(0))
		mBase = m.M
		v61 = m.ExcPending
		if v61 != 0 {
			return
		} else {
			if v60 == int32(0) {
				m.G0 = v8 + int32(32)
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
				F_errmsg_internal(m, int32(688824), v8)
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return
				} else {
					F_errfinish(m, int32(483029), int32(2377), int32(414349))
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return
					} else {
						m.G0 = v8 + int32(32)
						return
					}
				}
			}
		}
	}
}
func F_namefastcmp_c(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
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
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	goto L3
L1:
	;
	return v40 - v41
L3:
	;
	goto L4
L4:
	;
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v10 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v11 = l0
	v12 = l1
	v13 = int32(64)
	v14 = v10
	goto L9
L6:
	;
	v36 = l1
	v40 = int32(0)
	goto L7
L7:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	goto L1
L8:
	;
	v36 = v31
	v40 = v33
	goto L7
L9:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v14 != v16 {
		v31 = v12
		v33 = v14
		goto L8
	} else {
		goto L11
	}
L10:
	;
	v31 = v25
	v33 = int32(0)
	goto L8
L11:
	;
	if v16 == int32(0) {
		v31 = v12
		v33 = v14
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v21 = v13 - int32(1)
	if v21 == int32(0) {
		v31 = v12
		v33 = v14
		goto L8
	} else {
		goto L13
	}
L13:
	;
	v24 = int32(1)
	v25 = v12 + v24
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
	if v26 != 0 {
		v11 = v11 + v24
		v12 = v25
		v13 = v21
		v14 = v26
		goto L9
	} else {
		goto L14
	}
L14:
	;
	goto L10
}
func F_nameletext(m *base.Module, l0 int32) int32 {
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
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = F_DirectFunctionCall2Coll(m, int32(1557), v3, v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 <= int32(0))
	}
}
func F_namelt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v6 == int32(950) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return int32(base.Ui32(v172) >> (uint(int32(31)) % 32))
L2:
	;
	goto L7
L3:
	;
	goto L4
L4:
	;
	if v5&int32(3) == int32(0) {
		v77 = v5
		goto L21
	} else {
		goto L22
	}
L5:
	;
	v172 = v45 - v46
	goto L1
L7:
	;
	goto L8
L8:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v15 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v16 = v5
	v17 = v4
	v18 = int32(64)
	v19 = v15
	goto L13
L10:
	;
	v41 = v4
	v45 = int32(0)
	goto L11
L11:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	goto L5
L12:
	;
	v41 = v36
	v45 = v38
	goto L11
L13:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v19 != v21 {
		v36 = v17
		v38 = v19
		goto L12
	} else {
		goto L15
	}
L14:
	;
	v36 = v30
	v38 = int32(0)
	goto L12
L15:
	;
	if v21 == int32(0) {
		v36 = v17
		v38 = v19
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v26 = v18 - int32(1)
	if v26 == int32(0) {
		v36 = v17
		v38 = v19
		goto L12
	} else {
		goto L17
	}
L17:
	;
	v29 = int32(1)
	v30 = v17 + v29
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	if v31 != 0 {
		v16 = v16 + v29
		v17 = v30
		v18 = v26
		v19 = v31
		goto L13
	} else {
		goto L18
	}
L18:
	;
	goto L14
L19:
	;
	if v4&int32(3) == int32(0) {
		v134 = v4
		goto L38
	} else {
		goto L39
	}
L20:
	;
	v110 = v102 - v5
	goto L19
L21:
	;
	v81 = v77
	goto L30
L22:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v61 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v110 = int32(0)
	goto L19
L24:
	;
	goto L25
L25:
	;
	v66 = v5
	goto L26
L26:
	;
	v70 = v66 + int32(1)
	if v70&int32(3) == int32(0) {
		v77 = v70
		goto L21
	} else {
		goto L28
	}
L27:
	;
	v102 = v70
	goto L20
L28:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	if v75 != 0 {
		v66 = v70
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v90 = int32(-2139062144)
	if (int32(16843008)-v87|v87)&v90 == v90 {
		v81 = v81 + int32(4)
		goto L30
	} else {
		goto L32
	}
L31:
	;
	v96 = v81
	goto L33
L32:
	;
	goto L31
L33:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	if v100 != 0 {
		v96 = v96 + int32(1)
		goto L33
	} else {
		goto L35
	}
L34:
	;
	v102 = v96
	goto L20
L35:
	;
	goto L34
L36:
	;
	v168 = F_varstr_cmp(m, v5, v110, v4, v167, v6)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L53
	} else {
		goto L54
	}
L37:
	;
	v167 = v159 - v4
	goto L36
L38:
	;
	v138 = v134
	goto L47
L39:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
	if v118 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v167 = int32(0)
	goto L36
L41:
	;
	goto L42
L42:
	;
	v123 = v4
	goto L43
L43:
	;
	v127 = v123 + int32(1)
	if v127&int32(3) == int32(0) {
		v134 = v127
		goto L38
	} else {
		goto L45
	}
L44:
	;
	v159 = v127
	goto L37
L45:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	if v132 != 0 {
		v123 = v127
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	v147 = int32(-2139062144)
	if (int32(16843008)-v144|v144)&v147 == v147 {
		v138 = v138 + int32(4)
		goto L47
	} else {
		goto L49
	}
L48:
	;
	v153 = v138
	goto L50
L49:
	;
	goto L48
L50:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153))))
	if v157 != 0 {
		v153 = v153 + int32(1)
		goto L50
	} else {
		goto L52
	}
L51:
	;
	v159 = v153
	goto L37
L52:
	;
	goto L51
L53:
	;
	return int32(0)
L54:
	;
	v172 = v168
	goto L1
}
func F_namerecv(m *base.Module, l0 int32) int32 {
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
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
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
	var v49 int32
	_ = v49
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v14 = F_pq_getmsgtext(m, v8, v9-v10, v6+int32(12))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
		if int32(64) <= v18 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(34103428))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(315944), int32(0))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(64)
						F_errdetail(m, int32(554422), v6)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(481997), int32(95), int32(34479))
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
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
			v43 = F_palloc0(m, int32(64))
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int32(0)
			} else {
				v45 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
				if v45 != 0 {
					v46 = F__emscripten_memcpy_bulkmem(m, v43, v14, v45)
					mBase = m.M
					v47 = v46
				} else {
					v47 = v43
				}
				F_pfree(m, v14)
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int32(0)
				} else {
					m.G0 = v6 + int32(16)
					return v47
				}
			}
		}
	}
}
func F_networkjoinsel_semi(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) float64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v20 float64
	_ = v20
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 float32
	_ = v36
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
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
	var v59 float64
	_ = v59
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v97 float64
	_ = v97
	var v104 int32
	_ = v104
	var v105 float32
	_ = v105
	var v108 float32
	_ = v108
	var v111 float32
	_ = v111
	var v114 float32
	_ = v114
	var v116 float64
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v144 float64
	_ = v144
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v174 float64
	_ = v174
	var v179 int32
	_ = v179
	var v183 float32
	_ = v183
	var v185 float64
	_ = v185
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int64
	_ = v193
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v235 float64
	_ = v235
	var v237 float64
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 float32
	_ = v244
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 float64
	_ = v265
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v290 int32
	_ = v290
	var v301 float64
	_ = v301
	var v311 int32
	_ = v311
	var v312 float32
	_ = v312
	var v315 float32
	_ = v315
	var v318 float32
	_ = v318
	var v321 float32
	_ = v321
	var v323 float64
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v348 float64
	_ = v348
	var v361 int32
	_ = v361
	var v368 int32
	_ = v368
	var v378 float64
	_ = v378
	var v386 int32
	_ = v386
	var v390 float32
	_ = v390
	var v392 float64
	_ = v392
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v400 int64
	_ = v400
	var v424 int32
	_ = v424
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v438 float64
	_ = v438
	var v444 float64
	_ = v444
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v456 int32
	_ = v456
	var v462 float64
	_ = v462
	var v465 float64
	_ = v465
	var v466 int32
	_ = v466
	var v472 int32
	_ = v472
	var v488 int32
	_ = v488
	var v501 float64
	_ = v501
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v517 float32
	_ = v517
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v594 float64
	_ = v594
	var v595 int32
	_ = v595
	var v600 float64
	_ = v600
	var v601 float64
	_ = v601
	var v604 float64
	_ = v604
	var v634 float64
	_ = v634
	var v636 float64
	_ = v636
	var v638 int32
	_ = v638
	var v660 float64
	_ = v660
	var v667 int32
	_ = v667
	var v672 int32
	_ = v672
	var v674 float64
	_ = v674
	var v684 int32
	_ = v684
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v708 float64
	_ = v708
	var v716 int32
	_ = v716
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v725 int32
	_ = v725
	var v729 int32
	_ = v729
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v766 int32
	_ = v766
	var v798 float64
	_ = v798
	var v799 int32
	_ = v799
	var v804 float64
	_ = v804
	var v805 float64
	_ = v805
	var v808 float64
	_ = v808
	var v838 float64
	_ = v838
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v841 float64
	_ = v841
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v874 float64
	_ = v874
	var v891 float64
	_ = v891
	var v892 float64
	_ = v892
	var v898 float64
	_ = v898
	var v902 int32
	_ = v902
	var v906 int32
	_ = v906
	var v910 int32
	_ = v910
	var v914 int32
	_ = v914
	v5 = int32(0)
	v20 = float64(0)
	v28 = m.G0
	v30 = v28 - int32(192)
	m.G0 = v30
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v32 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v240 != 0 {
		goto L29
	} else {
		goto L30
	}
L2:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+22)))
	v36 = *(*float32)(unsafe.Add(mBase, uint32(v33+v34)+8))
	v42 = F_get_attstatsslot(m, v30+int32(128), v32, int32(1), int32(0), int32(3))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v191 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+160)) = v191
	v193 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v30)+152)) = v193
	*(*int64)(unsafe.Add(mBase, uint32(v30)+144)) = v193
	*(*int64)(unsafe.Add(mBase, uint32(v30)+136)) = v193
	*(*int64)(unsafe.Add(mBase, uint32(v30)+56)) = v193
	*(*int64)(unsafe.Add(mBase, uint32(v30-int32(-64)))) = v193
	*(*int64)(unsafe.Add(mBase, uint32(v30)+72)) = v193
	*(*int32)(unsafe.Add(mBase, uint32(v30)+80)) = v191
	*(*int64)(unsafe.Add(mBase, uint32(v30)+128)) = v193
	*(*int64)(unsafe.Add(mBase, uint32(v30)+48)) = v193
	v224 = v5
	v226 = v5
	v229 = v5
	v235 = v20
	v237 = v20
	goto L1
L5:
	;
	return float64(0)
L6:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v52 = F_get_attstatsslot(m, v30+int32(48), v48, int32(2), int32(0), int32(1))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v54 = int32(1024)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v30)+144))
	if v54 <= v55 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v58 = v54
	goto L10
L9:
	;
	v58 = v55
	goto L10
L10:
	;
	v59 = base.F64_promote_f32(v36)
	if v42 == int32(0) {
		v224 = v58
		v226 = v5
		v229 = v52
		v235 = v20
		v237 = v59
		goto L1
	} else {
		goto L11
	}
L11:
	;
	if v55 <= int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v224 = v58
	v226 = int32(1)
	v229 = v52
	v235 = v20
	v237 = v59
	goto L1
L13:
	;
	goto L14
L14:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v30)+148))
	v67 = v58 & int32(3)
	if v55 < int32(4) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	if v67 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L16:
	;
	v124 = int32(0)
	v144 = v20
	goto L15
L17:
	;
	goto L18
L18:
	;
	v77 = int32(0)
	v82 = v5
	v97 = v20
	goto L19
L19:
	;
	v104 = v65 + v77<<(uint(int32(2))%32)
	v105 = *(*float32)(unsafe.Add(mBase, uint32(v104)))
	v108 = *(*float32)(unsafe.Add(mBase, uint32(v104)+4))
	v111 = *(*float32)(unsafe.Add(mBase, uint32(v104)+8))
	v114 = *(*float32)(unsafe.Add(mBase, uint32(v104)+12))
	v116 = base.F64_add(base.F64_add(base.F64_add(base.F64_add(v97, base.F64_promote_f32(v105)), base.F64_promote_f32(v108)), base.F64_promote_f32(v111)), base.F64_promote_f32(v114))
	v117 = int32(4)
	v118 = v77 + v117
	v120 = v82 + v117
	if v120 != v58&int32(2044) {
		v77 = v118
		v82 = v120
		v97 = v116
		goto L19
	} else {
		goto L21
	}
L20:
	;
	v124 = v118
	v144 = v116
	goto L15
L21:
	;
	goto L20
L22:
	;
	v224 = v58
	v226 = int32(1)
	v229 = v52
	v235 = v144
	v237 = v59
	goto L1
L23:
	;
	goto L24
L24:
	;
	v154 = v124
	v158 = int32(0)
	v174 = v144
	goto L25
L25:
	;
	v179 = int32(1)
	v183 = *(*float32)(unsafe.Add(mBase, uint32(v65+v154<<(uint(int32(2))%32))))
	v185 = base.F64_add(v174, base.F64_promote_f32(v183))
	v189 = v158 + v179
	if v189 != v67 {
		v154 = v154 + v179
		v158 = v189
		v174 = v185
		goto L25
	} else {
		goto L27
	}
L26:
	;
	v224 = v58
	v226 = v179
	v229 = v52
	v235 = v185
	v237 = v59
	goto L1
L27:
	;
	goto L26
L28:
	;
	v446 = F_get_opcode(m, l0)
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L5
	} else {
		goto L54
	}
L29:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v240)+16))
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241)+22)))
	v244 = *(*float32)(unsafe.Add(mBase, uint32(v241+v242)+8))
	v250 = F_get_attstatsslot(m, v30+int32(88), v240, int32(1), int32(0), int32(3))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L5
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v398 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+120)) = v398
	v400 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v30)+112)) = v400
	*(*int64)(unsafe.Add(mBase, uint32(v30)+104)) = v400
	*(*int64)(unsafe.Add(mBase, uint32(v30)+96)) = v400
	*(*int64)(unsafe.Add(mBase, uint32(v30)+16)) = v400
	*(*int64)(unsafe.Add(mBase, uint32(v30)+24)) = v400
	*(*int64)(unsafe.Add(mBase, uint32(v30)+32)) = v400
	*(*int32)(unsafe.Add(mBase, uint32(v30)+40)) = v398
	*(*int64)(unsafe.Add(mBase, uint32(v30)+88)) = v400
	*(*int64)(unsafe.Add(mBase, uint32(v30)+8)) = v400
	v424 = v398
	v431 = v5
	v433 = v5
	v438 = v20
	v444 = v20
	goto L28
L32:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v258 = F_get_attstatsslot(m, v30+int32(8), v254, int32(2), int32(0), int32(1))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L5
	} else {
		goto L33
	}
L33:
	;
	v260 = int32(1024)
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v30)+104))
	if v260 <= v261 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v264 = v260
	goto L36
L35:
	;
	v264 = v261
	goto L36
L36:
	;
	v265 = base.F64_promote_f32(v244)
	if v250 == int32(0) {
		v424 = v264
		v431 = v5
		v433 = v258
		v438 = v20
		v444 = v265
		goto L28
	} else {
		goto L37
	}
L37:
	;
	if v261 <= int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v424 = v264
	v431 = int32(1)
	v433 = v258
	v438 = v20
	v444 = v265
	goto L28
L39:
	;
	goto L40
L40:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v30)+108))
	v273 = v264 & int32(3)
	if v261 < int32(4) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	if v273 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L42:
	;
	v331 = int32(0)
	v348 = v20
	goto L41
L43:
	;
	goto L44
L44:
	;
	v280 = int32(0)
	v284 = v280
	v290 = v280
	v301 = v20
	goto L45
L45:
	;
	v311 = v271 + v284<<(uint(int32(2))%32)
	v312 = *(*float32)(unsafe.Add(mBase, uint32(v311)))
	v315 = *(*float32)(unsafe.Add(mBase, uint32(v311)+4))
	v318 = *(*float32)(unsafe.Add(mBase, uint32(v311)+8))
	v321 = *(*float32)(unsafe.Add(mBase, uint32(v311)+12))
	v323 = base.F64_add(base.F64_add(base.F64_add(base.F64_add(v301, base.F64_promote_f32(v312)), base.F64_promote_f32(v315)), base.F64_promote_f32(v318)), base.F64_promote_f32(v321))
	v324 = int32(4)
	v325 = v284 + v324
	v327 = v290 + v324
	if v327 != v264&int32(2044) {
		v284 = v325
		v290 = v327
		v301 = v323
		goto L45
	} else {
		goto L47
	}
L46:
	;
	v331 = v325
	v348 = v323
	goto L41
L47:
	;
	goto L46
L48:
	;
	v424 = v264
	v431 = int32(1)
	v433 = v258
	v438 = v348
	v444 = v265
	goto L28
L49:
	;
	goto L50
L50:
	;
	v361 = v331
	v368 = int32(0)
	v378 = v348
	goto L51
L51:
	;
	v386 = int32(1)
	v390 = *(*float32)(unsafe.Add(mBase, uint32(v271+v361<<(uint(int32(2))%32))))
	v392 = base.F64_add(v378, base.F64_promote_f32(v390))
	v396 = v368 + v386
	if v396 != v273 {
		v361 = v361 + v386
		v368 = v396
		v378 = v392
		goto L51
	} else {
		goto L53
	}
L52:
	;
	v424 = v264
	v431 = v386
	v433 = v258
	v438 = v392
	v444 = v265
	goto L28
L53:
	;
	goto L52
L54:
	;
	F_fmgr_info(m, v446, v30+int32(164))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L5
	} else {
		goto L55
	}
L55:
	;
	if v433 == int32(0) {
		v465 = float64(0)
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v466 = v431 | v433
	if v226&v466 == int32(0) {
		v660 = v20
		goto L59
	} else {
		goto L60
	}
L57:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v456 == int32(0) {
		v465 = float64(0)
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v462 = *(*float64)(unsafe.Add(mBase, uint32(v456)+16))
	v465 = base.F64_mul(base.F64_sub(base.F64_sub(float64(1), v444), v438), v462)
	goto L56
L59:
	;
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v30)+64))
	if v466&(v229&base.B2i32(int32(2) < v667)) != 0 {
		goto L82
	} else {
		goto L83
	}
L60:
	;
	if v224 <= int32(0) {
		v660 = v20
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v472 = int32(0)
	v488 = v472
	v501 = v20
	goto L62
L62:
	;
	v509 = v488 << (uint(int32(2)) % 32)
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v30)+140))
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v509+v510)))
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v30)+24))
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v30)+148))
	v517 = *(*float32)(unsafe.Add(mBase, uint32(v515+v509)))
	if v431&base.B2i32(v472 < v424) == int32(0) {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	v660 = v636
	goto L59
L64:
	;
	v636 = base.F64_add(base.F64_mul(base.F64_promote_f32(v517), v634), v501)
	v638 = v488 + int32(1)
	if v638 != v224 {
		v488 = v638
		v501 = v636
		goto L62
	} else {
		goto L81
	}
L65:
	;
	if v433&base.F64_gt(v465, float64(0)) == int32(0) {
		goto L74
	} else {
		goto L75
	}
L66:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v30)+100))
	v525 = int32(0)
	goto L67
L67:
	;
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v521+v525<<(uint(int32(2))%32))))
	v557 = F_FunctionCall2Coll(m, v30+int32(164), int32(0), v512, v556)
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L5
	} else {
		goto L69
	}
L68:
	;
	v634 = float64(1)
	goto L64
L69:
	;
	if v557 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v562 = v525 + int32(1)
	if v424 != v562 {
		v525 = v562
		goto L67
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	goto L68
L73:
	;
	goto L65
L74:
	;
	v634 = float64(0)
	goto L64
L75:
	;
	v594 = F_inet_hist_value_sel(m, v514, v513, v512, v472-l1)
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L5
	} else {
		goto L76
	}
L76:
	;
	if base.F64_gt(v594, float64(0)) == int32(0) {
		goto L74
	} else {
		goto L77
	}
L77:
	;
	v600 = float64(1)
	v601 = base.F64_mul(v465, v594)
	if base.F64_gt(v601, v600) != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v604 = v600
	goto L80
L79:
	;
	v604 = v601
	goto L80
L80:
	;
	v634 = v604
	goto L64
L81:
	;
	goto L63
L82:
	;
	v672 = int32(0)
	v674 = float64(0)
	v684 = int32(1)
	v696 = v684
	v697 = v672
	v708 = v674
	goto L85
L83:
	;
	v874 = v660
	goto L84
L84:
	;
	if (v226|v229)&v466&int32(1) == int32(0) {
		goto L105
	} else {
		goto L106
	}
L85:
	;
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v30)+60))
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v716+v696<<(uint(int32(2))%32))))
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v30)+24))
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
	if v431&base.B2i32(v672 < v424) == int32(0) {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	v874 = base.F64_add(v660, base.F64_div(base.F64_mul(base.F64_sub(base.F64_sub(float64(1), v237), v235), v841), base.F64_convert_i32_s(v840)))
	goto L84
L87:
	;
	v839 = int32(1)
	v840 = v697 + v839
	v841 = base.F64_add(v708, v838)
	v842 = int32(base.Ui32(v667-int32(3))>>(uint(int32(10))%32)) + v684 + v696
	v843 = *(*int32)(unsafe.Add(mBase, uint32(v30)+64))
	if v842 < v843-v839 {
		v696 = v842
		v697 = v840
		v708 = v841
		goto L85
	} else {
		goto L104
	}
L88:
	;
	if v433&base.F64_gt(v465, v674) == int32(0) {
		goto L97
	} else {
		goto L98
	}
L89:
	;
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v30)+100))
	v729 = int32(0)
	goto L90
L90:
	;
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v725+v729<<(uint(int32(2))%32))))
	v761 = F_FunctionCall2Coll(m, v30+int32(164), int32(0), v720, v760)
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L5
	} else {
		goto L92
	}
L91:
	;
	v838 = float64(1)
	goto L87
L92:
	;
	if v761 == int32(0) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v766 = v729 + int32(1)
	if v424 != v766 {
		v729 = v766
		goto L90
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	goto L91
L96:
	;
	goto L88
L97:
	;
	v838 = float64(0)
	goto L87
L98:
	;
	v798 = F_inet_hist_value_sel(m, v722, v721, v720, v672-l1)
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L5
	} else {
		goto L99
	}
L99:
	;
	if base.F64_gt(v798, float64(0)) == int32(0) {
		goto L97
	} else {
		goto L100
	}
L100:
	;
	v804 = float64(1)
	v805 = base.F64_mul(v465, v798)
	if base.F64_gt(v805, v804) != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v808 = v804
	goto L103
L102:
	;
	v808 = v805
	goto L103
L103:
	;
	v838 = v808
	goto L87
L104:
	;
	goto L86
L105:
	;
	if l0 == int32(3552) {
		goto L108
	} else {
		goto L109
	}
L106:
	;
	v898 = v874
	goto L107
L107:
	;
	F_free_attstatsslot(m, v30+int32(128))
	mBase = m.M
	v902 = m.ExcPending
	if v902 != 0 {
		goto L5
	} else {
		goto L111
	}
L108:
	;
	v891 = float64(0.01)
	goto L110
L109:
	;
	v891 = float64(0.005)
	goto L110
L110:
	;
	v892 = float64(1)
	v898 = base.F64_mul(v891, base.F64_mul(base.F64_sub(v892, v237), base.F64_sub(v892, v444)))
	goto L107
L111:
	;
	F_free_attstatsslot(m, v30+int32(88))
	mBase = m.M
	v906 = m.ExcPending
	if v906 != 0 {
		goto L5
	} else {
		goto L112
	}
L112:
	;
	F_free_attstatsslot(m, v30+int32(48))
	mBase = m.M
	v910 = m.ExcPending
	if v910 != 0 {
		goto L5
	} else {
		goto L113
	}
L113:
	;
	F_free_attstatsslot(m, v30+int32(8))
	mBase = m.M
	v914 = m.ExcPending
	if v914 != 0 {
		goto L5
	} else {
		goto L114
	}
L114:
	;
	m.G0 = v30 + int32(192)
	return v898
}
func F_newnfa(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int64
	_ = v22
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v281 int32
	_ = v281
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v334 int32
	_ = v334
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v361 int32
	_ = v361
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v387 int32
	_ = v387
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	v8 = F_palloc_extended(m, int32(84), int32(2))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v8 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v16 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v22 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v8)+76)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v8)+52)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v8)+72)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+56)) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+80)) = l2
	*(*int64)(unsafe.Add(mBase, uint32(v8)+64)) = int64(-4294967296)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+40)) = v22
	*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = v22
	*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v22
	v41 = F_newstate(m, v8)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L9
	}
L6:
	;
	v18 = v16
	goto L8
L7:
	;
	v18 = int32(12)
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v18
	return int32(0)
L9:
	;
	if v41 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v43 = int32(64)
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+4)) = uint8(v43)
	goto L12
L11:
	;
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v41
	v46 = F_newstate(m, v8)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	if v46 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v48 = int32(62)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+4)) = uint8(v48)
	goto L16
L15:
	;
	goto L16
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v46
	v51 = F_newstate(m, v8)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v51
	v54 = F_newstate(m, v8)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v54
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v57 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = v399
	F_pfree(m, v8)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L1
	} else {
		goto L160
	}
L20:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v8)+36))
	if v58 != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v8)+52))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	F_rainbow(m, v8, v108, int32(-1), v110, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L37
	}
L23:
	;
	v61 = v58
	goto L26
L24:
	;
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = int32(0)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v8)+40))
	if v83 != 0 {
		goto L30
	} else {
		goto L31
	}
L26:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v8)+76))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+136))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v65)+136)) = v66 + v67*int32(-36) - int32(8)
	F_pfree(m, v61)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L28
	}
L27:
	;
	goto L25
L28:
	;
	if v64 != 0 {
		v61 = v64
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v86 = v83
	goto L33
L31:
	;
	goto L32
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = int32(-1)
	v399 = int32(0)
	goto L19
L33:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v8)+76))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+136))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v90)+136)) = v91 + v92*int32(-40) - int32(8)
	F_pfree(m, v86)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L35
	}
L34:
	;
	goto L32
L35:
	;
	if v89 != 0 {
		v86 = v89
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v117 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	if v117 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v115)+12))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v114)+8))
	if v120 <= v121 {
		goto L44
	} else {
		goto L45
	}
L41:
	;
	goto L40
L42:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v174 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	if v174 != 0 {
		goto L64
	} else {
		goto L65
	}
L43:
	;
	F_createarc(m, v8, int32(94), int32(1), v115, v114)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L63
	}
L44:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v115)+20))
	if v123 == int32(0) {
		goto L43
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v114)+16))
	if v140 == int32(0) {
		goto L43
	} else {
		goto L55
	}
L47:
	;
	v128 = v123
	goto L48
L48:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v128)+12))
	if v131 != v114 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	goto L43
L50:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v128)+16))
	if v139 != 0 {
		v128 = v139
		goto L48
	} else {
		goto L54
	}
L51:
	;
	v133 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v128)+4)))
	if v133 != int32(1) {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
	if v136 == int32(94) {
		goto L42
	} else {
		goto L53
	}
L53:
	;
	goto L50
L54:
	;
	goto L49
L55:
	;
	v145 = v140
	goto L56
L56:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v145)+8))
	if v148 != v115 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	goto L43
L58:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v145)+24))
	if v156 != 0 {
		v145 = v156
		goto L56
	} else {
		goto L62
	}
L59:
	;
	v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v145)+4)))
	if v150 != int32(1) {
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
	if v153 == int32(94) {
		goto L42
	} else {
		goto L61
	}
L61:
	;
	goto L58
L62:
	;
	goto L57
L63:
	;
	goto L42
L64:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v172)+12))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v171)+8))
	if v177 <= v178 {
		goto L70
	} else {
		goto L71
	}
L67:
	;
	goto L66
L68:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v8)+52))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	F_rainbow(m, v8, v224, int32(-1), v226, v227)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L90
	}
L69:
	;
	F_createarc(m, v8, int32(94), int32(0), v172, v171)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L89
	}
L70:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v172)+20))
	if v180 == int32(0) {
		goto L69
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v171)+16))
	if v195 == int32(0) {
		goto L69
	} else {
		goto L81
	}
L73:
	;
	v185 = v180
	goto L74
L74:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v185)+12))
	if v188 != v171 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	goto L69
L76:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v185)+16))
	if v194 != 0 {
		v185 = v194
		goto L74
	} else {
		goto L80
	}
L77:
	;
	v190 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v185)+4)))
	if v190 != 0 {
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v185)))
	if v191 == int32(94) {
		goto L68
	} else {
		goto L79
	}
L79:
	;
	goto L76
L80:
	;
	goto L75
L81:
	;
	v200 = v195
	goto L82
L82:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v200)+8))
	if v203 != v172 {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	goto L69
L84:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v200)+24))
	if v209 != 0 {
		v200 = v209
		goto L82
	} else {
		goto L88
	}
L85:
	;
	v205 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v200)+4)))
	if v205 != 0 {
		goto L84
	} else {
		goto L86
	}
L86:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v200)))
	if v206 == int32(94) {
		goto L68
	} else {
		goto L87
	}
L87:
	;
	goto L84
L88:
	;
	goto L83
L89:
	;
	goto L68
L90:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v233 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	if v233 != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v231)+12))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v230)+8))
	if v236 <= v237 {
		goto L97
	} else {
		goto L98
	}
L94:
	;
	goto L93
L95:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v290 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	if v290 != 0 {
		goto L117
	} else {
		goto L118
	}
L96:
	;
	F_createarc(m, v8, int32(36), int32(1), v231, v230)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L116
	}
L97:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v231)+20))
	if v239 == int32(0) {
		goto L96
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v230)+16))
	if v256 == int32(0) {
		goto L96
	} else {
		goto L108
	}
L100:
	;
	v244 = v239
	goto L101
L101:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v244)+12))
	if v247 != v230 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	goto L96
L103:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v244)+16))
	if v255 != 0 {
		v244 = v255
		goto L101
	} else {
		goto L107
	}
L104:
	;
	v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v244)+4)))
	if v249 != int32(1) {
		goto L103
	} else {
		goto L105
	}
L105:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v244)))
	if v252 == int32(36) {
		goto L95
	} else {
		goto L106
	}
L106:
	;
	goto L103
L107:
	;
	goto L102
L108:
	;
	v261 = v256
	goto L109
L109:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v261)+8))
	if v264 != v231 {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	goto L96
L111:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v261)+24))
	if v272 != 0 {
		v261 = v272
		goto L109
	} else {
		goto L115
	}
L112:
	;
	v266 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v261)+4)))
	if v266 != int32(1) {
		goto L111
	} else {
		goto L113
	}
L113:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v261)))
	if v269 == int32(36) {
		goto L95
	} else {
		goto L114
	}
L114:
	;
	goto L111
L115:
	;
	goto L110
L116:
	;
	goto L95
L117:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v288)+12))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v287)+8))
	if v293 <= v294 {
		goto L123
	} else {
		goto L124
	}
L120:
	;
	goto L119
L121:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v340 == int32(0) {
		goto L143
	} else {
		goto L144
	}
L122:
	;
	F_createarc(m, v8, int32(36), int32(0), v288, v287)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L1
	} else {
		goto L142
	}
L123:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v288)+20))
	if v296 == int32(0) {
		goto L122
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v287)+16))
	if v311 == int32(0) {
		goto L122
	} else {
		goto L134
	}
L126:
	;
	v301 = v296
	goto L127
L127:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v301)+12))
	if v304 != v287 {
		goto L129
	} else {
		goto L130
	}
L128:
	;
	goto L122
L129:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v301)+16))
	if v310 != 0 {
		v301 = v310
		goto L127
	} else {
		goto L133
	}
L130:
	;
	v306 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v301)+4)))
	if v306 != 0 {
		goto L129
	} else {
		goto L131
	}
L131:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v301)))
	if v307 == int32(36) {
		goto L121
	} else {
		goto L132
	}
L132:
	;
	goto L129
L133:
	;
	goto L128
L134:
	;
	v316 = v311
	goto L135
L135:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v316)+8))
	if v319 != v288 {
		goto L137
	} else {
		goto L138
	}
L136:
	;
	goto L122
L137:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v316)+24))
	if v325 != 0 {
		v316 = v325
		goto L135
	} else {
		goto L141
	}
L138:
	;
	v321 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v316)+4)))
	if v321 != 0 {
		goto L137
	} else {
		goto L139
	}
L139:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v316)))
	if v322 == int32(36) {
		goto L121
	} else {
		goto L140
	}
L140:
	;
	goto L137
L141:
	;
	goto L136
L142:
	;
	goto L121
L143:
	;
	return v8
L144:
	;
	goto L145
L145:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v8)+36))
	if v344 != 0 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v347 = v344
	goto L149
L147:
	;
	goto L148
L148:
	;
	v367 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v367
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v8)+40))
	if v370 != 0 {
		goto L153
	} else {
		goto L154
	}
L149:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v347)))
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v8)+76))
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v351)+136))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v347)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v351)+136)) = v352 + v353*int32(-36) - int32(8)
	F_pfree(m, v347)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L1
	} else {
		goto L151
	}
L150:
	;
	goto L148
L151:
	;
	if v350 != 0 {
		v347 = v350
		goto L149
	} else {
		goto L152
	}
L152:
	;
	goto L150
L153:
	;
	v373 = v370
	goto L156
L154:
	;
	goto L155
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = int32(-1)
	v399 = v367
	goto L19
L156:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v373)))
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v8)+76))
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v377)+136))
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v373)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v377)+136)) = v378 + v379*int32(-40) - int32(8)
	F_pfree(m, v373)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L1
	} else {
		goto L158
	}
L157:
	;
	goto L155
L158:
	;
	if v376 != 0 {
		v373 = v376
		goto L156
	} else {
		goto L159
	}
L159:
	;
	goto L157
L160:
	;
	return v399
}
func F_nextval_oid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_nextval_internal(m, v2, int32(1))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = F_Int64GetDatum(m, v4)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v8
		}
	}
}
func F_nlikesel(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
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
	var v33 float64
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = F_get_negator(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if v11 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(201796), int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(476144), int32(773), int32(295907))
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
			v30 = int32(0)
			v33 = F_patternsel_common(m, v9, v11, v30, v7, v6, v8, v30, int32(1))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				v35 = F_Float8GetDatum(m, v33)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					return v35
				}
			}
		}
	}
}
func F_nonword(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
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
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v123 int32
	_ = v123
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v12 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	if l1 == int32(97) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	v17 = int32(36)
	goto L8
L7:
	;
	v17 = int32(94)
	goto L8
L8:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v18 <= v19 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v76 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	if v76 != 0 {
		goto L31
	} else {
		goto L32
	}
L10:
	;
	F_createarc(m, v10, v17, int32(1), l2, l3)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L4
	} else {
		goto L30
	}
L11:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v21 == int32(0) {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v39 == int32(0) {
		goto L10
	} else {
		goto L22
	}
L14:
	;
	v28 = v21
	goto L15
L15:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	if v31 != l3 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L10
L17:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	if v38 != 0 {
		v28 = v38
		goto L15
	} else {
		goto L21
	}
L18:
	;
	v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+4)))
	if v33 != int32(1) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v36 == v17 {
		goto L9
	} else {
		goto L20
	}
L20:
	;
	goto L17
L21:
	;
	goto L16
L22:
	;
	v46 = v39
	goto L23
L23:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
	if v49 != l2 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	goto L10
L25:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v46)+24))
	if v56 != 0 {
		v46 = v56
		goto L23
	} else {
		goto L29
	}
L26:
	;
	v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v46)+4)))
	if v51 != int32(1) {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	if v54 == v17 {
		goto L9
	} else {
		goto L28
	}
L28:
	;
	goto L25
L29:
	;
	goto L24
L30:
	;
	goto L9
L31:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L4
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v79 <= v80 {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	goto L33
L35:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	F_colorcomplement(m, v131, v132, l1, v133, l2, l3)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L4
	} else {
		goto L57
	}
L36:
	;
	F_createarc(m, v74, v17, int32(0), l2, l3)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L4
	} else {
		goto L56
	}
L37:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v82 == int32(0) {
		goto L36
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v98 == int32(0) {
		goto L36
	} else {
		goto L48
	}
L40:
	;
	v89 = v82
	goto L41
L41:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	if v92 != l3 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	goto L36
L43:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
	if v97 != 0 {
		v89 = v97
		goto L41
	} else {
		goto L47
	}
L44:
	;
	v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89)+4)))
	if v94 != 0 {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	if v95 == v17 {
		goto L35
	} else {
		goto L46
	}
L46:
	;
	goto L43
L47:
	;
	goto L42
L48:
	;
	v105 = v98
	goto L49
L49:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v105)+8))
	if v108 != l2 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	goto L36
L51:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v105)+24))
	if v113 != 0 {
		v105 = v113
		goto L49
	} else {
		goto L55
	}
L52:
	;
	v110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v105)+4)))
	if v110 != 0 {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	if v111 == v17 {
		goto L35
	} else {
		goto L54
	}
L54:
	;
	goto L51
L55:
	;
	goto L50
L56:
	;
	goto L35
L57:
	;
	return
}
func F_notification_match(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
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
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v81 int32
	_ = v81
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5))))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v8 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7))))
	if v6 != v8 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(1)
L2:
	;
	v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5)+2)))
	v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+2)))
	if v10 != v11 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v13 = int32(4)
	v14 = v5 + v13
	v16 = v7 + v13
	v19 = v6 + v10 + int32(2)
	if base.Ui32(v13) <= base.Ui32(v19) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	if v81 != 0 {
		goto L1
	} else {
		goto L22
	}
L5:
	;
	v81 = int32(0)
	goto L4
L6:
	;
	v55 = v50
	v56 = v51
	v57 = v52
	goto L16
L7:
	;
	if (v14|v16)&int32(3) != 0 {
		v50 = v14
		v51 = v16
		v52 = v19
		goto L6
	} else {
		goto L10
	}
L8:
	;
	v43 = v14
	v44 = v16
	v45 = v19
	goto L9
L9:
	;
	if v45 == int32(0) {
		goto L5
	} else {
		goto L15
	}
L10:
	;
	v27 = v14
	v28 = v16
	v29 = v19
	goto L11
L11:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v32 != v33 {
		v50 = v27
		v51 = v28
		v52 = v29
		goto L6
	} else {
		goto L13
	}
L12:
	;
	v43 = v38
	v44 = v36
	v45 = v40
	goto L9
L13:
	;
	v35 = int32(4)
	v36 = v28 + v35
	v38 = v27 + v35
	v40 = v29 - v35
	if base.Ui32(int32(3)) < base.Ui32(v40) {
		v27 = v38
		v28 = v36
		v29 = v40
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v50 = v43
	v51 = v44
	v52 = v45
	goto L6
L16:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	if v60 == v61 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v81 = v60 - v61
	goto L4
L18:
	;
	v63 = int32(1)
	v68 = v57 - v63
	if v68 != 0 {
		v55 = v55 + v63
		v56 = v56 + v63
		v57 = v68
		goto L16
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	goto L17
L21:
	;
	goto L5
L22:
	;
	return int32(0)
}
