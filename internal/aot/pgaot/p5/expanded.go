package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_expanded_record_fetch_tupdesc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v4 == int32(0) {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v9 = F_lookup_rowtype_tupdesc(m, v7, v8)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
			if int32(0) <= v13 {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
				if v16 == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = int32(1308)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = l0
					v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v24 = l0 + int32(108)
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
					*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v25
					v27 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v22)+4)) = uint8(v27)
					*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v24
				} else {
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v9
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
				v33 = v31 + int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v33
				if v33 < int32(0) {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
					v43 = F_assign_record_type_identifier(m, v41, v42)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v43
						v46 = v9
						return v46
					}
				} else {
					F_DecrTupleDescRefCount(m, v9)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						v41 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
						v42 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
						v43 = F_assign_record_type_identifier(m, v41, v42)
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v43
							v46 = v9
							return v46
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v9
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
				v43 = F_assign_record_type_identifier(m, v41, v42)
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v43
					v46 = v9
					return v46
				}
			}
		}
	} else {
		v46 = v4
		return v46
	}
}
func F_expanded_record_lookup_field(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v7 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v10 = F_expanded_record_fetch_tupdesc(m, l0)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v14 = v7
	goto L3
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if int32(0) < v15 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	return int32(0)
L5:
	;
	v14 = v10
	goto L3
L6:
	;
	v103 = int32(*(*int16)(unsafe.Add(mBase, uint32(v100)+74)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v103
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v100)+68))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v105
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v100)+76))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v107
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v100)+96))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v109
	return int32(1)
L7:
	;
	v21 = int32(0)
	v24 = v15
	goto L10
L8:
	;
	goto L9
L9:
	;
	v65 = F_strcmp(m, int32(797468), l1)
	mBase = m.M
	if v65 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L10:
	;
	v27 = int32(4)
	v32 = v14 + int32(20) + v24<<(uint(v27)%32) + v21*int32(100)
	v34 = v32 + v27
	if v34|l1 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L9
L12:
	;
	if v48 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L13:
	;
	v40 = int32(-1)
	goto L15
L14:
	;
	v40 = int32(0)
	goto L15
L15:
	;
	if v34 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v41 = int32(1)
	goto L18
L17:
	;
	v41 = v40
	goto L18
L18:
	;
	if v34 == int32(0) {
		v48 = v41
		goto L19
	} else {
		goto L20
	}
L19:
	;
	goto L12
L20:
	;
	if l1 == int32(0) {
		v48 = v41
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v47 = F_strncmp(m, v34, l1, int32(64))
	mBase = m.M
	v48 = v47
	goto L19
L22:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+91)))
	if v51 != int32(1) {
		v100 = v32
		goto L6
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v55 = v21 + int32(1)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v55 < v56 {
		v21 = v55
		v24 = v56
		goto L10
	} else {
		goto L26
	}
L25:
	;
	goto L24
L26:
	;
	goto L11
L27:
	;
	if v94 != 0 {
		v100 = v94
		goto L6
	} else {
		goto L46
	}
L28:
	;
	v94 = int32(797464)
	goto L27
L29:
	;
	goto L30
L30:
	;
	v70 = F_strcmp(m, int32(797568), l1)
	mBase = m.M
	if v70 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v94 = int32(797564)
	goto L27
L32:
	;
	goto L33
L33:
	;
	v75 = F_strcmp(m, int32(797668), l1)
	mBase = m.M
	if v75 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v94 = int32(797664)
	goto L27
L35:
	;
	goto L36
L36:
	;
	v80 = F_strcmp(m, int32(797768), l1)
	mBase = m.M
	if v80 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v94 = int32(797764)
	goto L27
L38:
	;
	goto L39
L39:
	;
	v85 = F_strcmp(m, int32(797868), l1)
	mBase = m.M
	if v85 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v94 = int32(797864)
	goto L27
L41:
	;
	goto L42
L42:
	;
	v92 = F_strcmp(m, int32(797968), l1)
	mBase = m.M
	if v92 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v93 = int32(0)
	goto L45
L44:
	;
	v93 = int32(797964)
	goto L45
L45:
	;
	v94 = v93
	goto L27
L46:
	;
	return int32(0)
}
func F_make_expanded_record_from_typeid(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
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
	var v29 int32
	_ = v29
	var v32 int64
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int64
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int64
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l0 != int32(2249) {
		v15 = F_lookup_type_cache(m, l0, int32(4352))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+13)))
			if v19 == int32(100) {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v15)+300))
				v25 = F_lookup_type_cache(m, v23, int32(256))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = v25
					v28 = int32(64)
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v27)+188))
					if v29 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v124 = m.ExcPending
						if v124 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(151027844))
							mBase = m.M
							v127 = m.ExcPending
							if v127 != 0 {
								return int32(0)
							} else {
								v128 = F_format_type_be(m, l0)
								mBase = m.M
								v129 = m.ExcPending
								if v129 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10))) = v128
									F_errmsg(m, int32(365444), v10)
									mBase = m.M
									v133 = m.ExcPending
									if v133 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(522459), int32(100), int32(455608))
										mBase = m.M
										v138 = m.ExcPending
										if v138 != 0 {
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
						v32 = *(*int64)(unsafe.Add(mBase, uint32(v27)+192))
						v40 = v29
						v41 = v28
						v42 = v32
						v47 = F_AllocSetContextCreateInternal(m, l2, int32(440424), int32(0), int32(8192), int32(8388608))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							v49 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
							v54 = F_MemoryContextAlloc(m, v47, v49*int32(5)+int32(120))
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int32(0)
							} else {
								v59 = F__emscripten_memset_bulkmem(m, v54, base.I32_extend8_s(int32(0)), int32(120))
								mBase = m.M
								v61 = int32(513)
								*(*uint16)(unsafe.Add(mBase, uint32(v59)+18)) = uint16(v61)
								v63 = int32(769)
								*(*uint16)(unsafe.Add(mBase, uint32(v59)+12)) = uint16(v63)
								*(*int32)(unsafe.Add(mBase, uint32(v59)+8)) = v47
								*(*int32)(unsafe.Add(mBase, uint32(v59)+4)) = int32(1694320)
								*(*int32)(unsafe.Add(mBase, uint32(v59))) = int32(-1)
								*(*int32)(unsafe.Add(mBase, uint32(v59)+20)) = v59
								*(*int32)(unsafe.Add(mBase, uint32(v59)+14)) = v59
								v72 = v59 + int32(120)
								*(*int32)(unsafe.Add(mBase, uint32(v59)+56)) = v72
								*(*int32)(unsafe.Add(mBase, uint32(v59)+24)) = int32(1384727874)
								v76 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
								*(*int32)(unsafe.Add(mBase, uint32(v59)+60)) = v72 + v76<<(uint(int32(2))%32)
								v81 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
								*(*int32)(unsafe.Add(mBase, uint32(v59)+32)) = l0
								*(*int32)(unsafe.Add(mBase, uint32(v59)+64)) = v81
								v84 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v59)+36)) = v84
								v86 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v59)+48)) = v42
								*(*int32)(unsafe.Add(mBase, uint32(v59)+40)) = v86
								*(*int32)(unsafe.Add(mBase, uint32(v59)+28)) = v41
								v90 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
								if int32(0) <= v90 {
									*(*int32)(unsafe.Add(mBase, uint32(v59)+108)) = int32(1308)
									*(*int32)(unsafe.Add(mBase, uint32(v59)+112)) = v59
									v96 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
									v98 = v59 + int32(108)
									v99 = *(*int32)(unsafe.Add(mBase, uint32(v96)+40))
									*(*int32)(unsafe.Add(mBase, uint32(v98)+8)) = v99
									v101 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v96)+4)) = uint8(v101)
									*(*int32)(unsafe.Add(mBase, uint32(v96)+40)) = v98
									*(*int32)(unsafe.Add(mBase, uint32(v59)+44)) = v40
									v105 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
									v107 = v105 + int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v40)+12)) = v107
									if l0 != int32(2249) {
										m.G0 = v10 + int32(16)
										return v59
									} else {
										if v107 < int32(0) {
											m.G0 = v10 + int32(16)
											return v59
										} else {
											F_DecrTupleDescRefCount(m, v40)
											mBase = m.M
											v114 = m.ExcPending
											if v114 != 0 {
												return int32(0)
											} else {
												m.G0 = v10 + int32(16)
												return v59
											}
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v59)+44)) = v40
									m.G0 = v10 + int32(16)
									return v59
								}
							}
						}
					}
				}
			} else {
				v27 = v15
				v28 = v4
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v27)+188))
				if v29 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v124 = m.ExcPending
					if v124 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(151027844))
						mBase = m.M
						v127 = m.ExcPending
						if v127 != 0 {
							return int32(0)
						} else {
							v128 = F_format_type_be(m, l0)
							mBase = m.M
							v129 = m.ExcPending
							if v129 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10))) = v128
								F_errmsg(m, int32(365444), v10)
								mBase = m.M
								v133 = m.ExcPending
								if v133 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(522459), int32(100), int32(455608))
									mBase = m.M
									v138 = m.ExcPending
									if v138 != 0 {
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
					v32 = *(*int64)(unsafe.Add(mBase, uint32(v27)+192))
					v40 = v29
					v41 = v28
					v42 = v32
					v47 = F_AllocSetContextCreateInternal(m, l2, int32(440424), int32(0), int32(8192), int32(8388608))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
						v54 = F_MemoryContextAlloc(m, v47, v49*int32(5)+int32(120))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							v59 = F__emscripten_memset_bulkmem(m, v54, base.I32_extend8_s(int32(0)), int32(120))
							mBase = m.M
							v61 = int32(513)
							*(*uint16)(unsafe.Add(mBase, uint32(v59)+18)) = uint16(v61)
							v63 = int32(769)
							*(*uint16)(unsafe.Add(mBase, uint32(v59)+12)) = uint16(v63)
							*(*int32)(unsafe.Add(mBase, uint32(v59)+8)) = v47
							*(*int32)(unsafe.Add(mBase, uint32(v59)+4)) = int32(1694320)
							*(*int32)(unsafe.Add(mBase, uint32(v59))) = int32(-1)
							*(*int32)(unsafe.Add(mBase, uint32(v59)+20)) = v59
							*(*int32)(unsafe.Add(mBase, uint32(v59)+14)) = v59
							v72 = v59 + int32(120)
							*(*int32)(unsafe.Add(mBase, uint32(v59)+56)) = v72
							*(*int32)(unsafe.Add(mBase, uint32(v59)+24)) = int32(1384727874)
							v76 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
							*(*int32)(unsafe.Add(mBase, uint32(v59)+60)) = v72 + v76<<(uint(int32(2))%32)
							v81 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
							*(*int32)(unsafe.Add(mBase, uint32(v59)+32)) = l0
							*(*int32)(unsafe.Add(mBase, uint32(v59)+64)) = v81
							v84 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v59)+36)) = v84
							v86 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v59)+48)) = v42
							*(*int32)(unsafe.Add(mBase, uint32(v59)+40)) = v86
							*(*int32)(unsafe.Add(mBase, uint32(v59)+28)) = v41
							v90 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
							if int32(0) <= v90 {
								*(*int32)(unsafe.Add(mBase, uint32(v59)+108)) = int32(1308)
								*(*int32)(unsafe.Add(mBase, uint32(v59)+112)) = v59
								v96 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
								v98 = v59 + int32(108)
								v99 = *(*int32)(unsafe.Add(mBase, uint32(v96)+40))
								*(*int32)(unsafe.Add(mBase, uint32(v98)+8)) = v99
								v101 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v96)+4)) = uint8(v101)
								*(*int32)(unsafe.Add(mBase, uint32(v96)+40)) = v98
								*(*int32)(unsafe.Add(mBase, uint32(v59)+44)) = v40
								v105 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
								v107 = v105 + int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v40)+12)) = v107
								if l0 != int32(2249) {
									m.G0 = v10 + int32(16)
									return v59
								} else {
									if v107 < int32(0) {
										m.G0 = v10 + int32(16)
										return v59
									} else {
										F_DecrTupleDescRefCount(m, v40)
										mBase = m.M
										v114 = m.ExcPending
										if v114 != 0 {
											return int32(0)
										} else {
											m.G0 = v10 + int32(16)
											return v59
										}
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v59)+44)) = v40
								m.G0 = v10 + int32(16)
								return v59
							}
						}
					}
				}
			}
		}
	} else {
		v34 = F_lookup_rowtype_tupdesc(m, int32(2249), l1)
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return int32(0)
		} else {
			v37 = F_assign_record_type_identifier(m, int32(2249), l1)
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int32(0)
			} else {
				v40 = v34
				v41 = v4
				v42 = v37
				v47 = F_AllocSetContextCreateInternal(m, l2, int32(440424), int32(0), int32(8192), int32(8388608))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return int32(0)
				} else {
					v49 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
					v54 = F_MemoryContextAlloc(m, v47, v49*int32(5)+int32(120))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						v59 = F__emscripten_memset_bulkmem(m, v54, base.I32_extend8_s(int32(0)), int32(120))
						mBase = m.M
						v61 = int32(513)
						*(*uint16)(unsafe.Add(mBase, uint32(v59)+18)) = uint16(v61)
						v63 = int32(769)
						*(*uint16)(unsafe.Add(mBase, uint32(v59)+12)) = uint16(v63)
						*(*int32)(unsafe.Add(mBase, uint32(v59)+8)) = v47
						*(*int32)(unsafe.Add(mBase, uint32(v59)+4)) = int32(1694320)
						*(*int32)(unsafe.Add(mBase, uint32(v59))) = int32(-1)
						*(*int32)(unsafe.Add(mBase, uint32(v59)+20)) = v59
						*(*int32)(unsafe.Add(mBase, uint32(v59)+14)) = v59
						v72 = v59 + int32(120)
						*(*int32)(unsafe.Add(mBase, uint32(v59)+56)) = v72
						*(*int32)(unsafe.Add(mBase, uint32(v59)+24)) = int32(1384727874)
						v76 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
						*(*int32)(unsafe.Add(mBase, uint32(v59)+60)) = v72 + v76<<(uint(int32(2))%32)
						v81 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
						*(*int32)(unsafe.Add(mBase, uint32(v59)+32)) = l0
						*(*int32)(unsafe.Add(mBase, uint32(v59)+64)) = v81
						v84 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v59)+36)) = v84
						v86 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v59)+48)) = v42
						*(*int32)(unsafe.Add(mBase, uint32(v59)+40)) = v86
						*(*int32)(unsafe.Add(mBase, uint32(v59)+28)) = v41
						v90 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
						if int32(0) <= v90 {
							*(*int32)(unsafe.Add(mBase, uint32(v59)+108)) = int32(1308)
							*(*int32)(unsafe.Add(mBase, uint32(v59)+112)) = v59
							v96 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
							v98 = v59 + int32(108)
							v99 = *(*int32)(unsafe.Add(mBase, uint32(v96)+40))
							*(*int32)(unsafe.Add(mBase, uint32(v98)+8)) = v99
							v101 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v96)+4)) = uint8(v101)
							*(*int32)(unsafe.Add(mBase, uint32(v96)+40)) = v98
							*(*int32)(unsafe.Add(mBase, uint32(v59)+44)) = v40
							v105 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
							v107 = v105 + int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v40)+12)) = v107
							if l0 != int32(2249) {
								m.G0 = v10 + int32(16)
								return v59
							} else {
								if v107 < int32(0) {
									m.G0 = v10 + int32(16)
									return v59
								} else {
									F_DecrTupleDescRefCount(m, v40)
									mBase = m.M
									v114 = m.ExcPending
									if v114 != 0 {
										return int32(0)
									} else {
										m.G0 = v10 + int32(16)
										return v59
									}
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v59)+44)) = v40
							m.G0 = v10 + int32(16)
							return v59
						}
					}
				}
			}
		}
	}
}
