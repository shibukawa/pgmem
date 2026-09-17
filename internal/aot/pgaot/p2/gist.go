package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_gistPopItupFromNodeBuffer(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
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
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v85 int32
	_ = v85
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v13 <= int32(0) {
		m.G0 = v11 + int32(16)
		return base.B2i32(int32(0) < v13)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		if v16 == int32(0) {
			F_gistLoadNodeBuffer(m, l0, l1)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				v24 = v23
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
				v26 = v24 + v25
				v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+14)))
				v29 = v27 & int32(_a_F_gistPopItupFromNodeBuffer_0)
				v30 = F_palloc(m, v29)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v30
					if v29 != 0 {
						base.MemoryCopy(m, v30, v26+int32(8), v29)
					} else {
					}
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v36 + (v29+int32(7))&int32(_a_F_gistPopItupFromNodeBuffer_1)
					v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
					if v44 != int32(_a_F_gistPopItupFromNodeBuffer_2) {
						m.G0 = v11 + int32(16)
						return base.B2i32(int32(0) < v13)
					} else {
						v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v47 - int32(1)
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
						if v51 != int32(-1) {
							v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v56 = F_BufFileSeekBlock(m, v54, base.I64_extend_i32_s(v51))
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int32(0)
							} else {
								if v56 != 0 {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v102 = m.ExcPending
									if v102 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v11))) = v51
										F_errmsg_internal(m, int32(_a_F_gistPopItupFromNodeBuffer_3), v11)
										mBase = m.M
										v106 = m.ExcPending
										if v106 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_gistPopItupFromNodeBuffer_4), int32(753), int32(_a_F_gistPopItupFromNodeBuffer_5))
											mBase = m.M
											v111 = m.ExcPending
											if v111 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									F_BufFileReadExact(m, v54, v43, int32(_a_F_gistPopItupFromNodeBuffer_6))
									mBase = m.M
									v60 = m.ExcPending
									if v60 != 0 {
										return int32(0)
									} else {
										v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
										v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
										if v61 < v62 {
											v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											v75 = v61
											v76 = v64
											*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v75 + int32(1)
											*(*int32)(unsafe.Add(mBase, uint32(v76+v75<<(uint(int32(2))%32)))) = v51
											m.G0 = v11 + int32(16)
											return base.B2i32(int32(0) < v13)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v62 << (uint(int32(1)) % 32)
											v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											v71 = F_repalloc(m, v68, v62<<(uint(int32(3))%32))
											mBase = m.M
											v72 = m.ExcPending
											if v72 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v71
												v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
												v75 = v74
												v76 = v71
												*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v75 + int32(1)
												*(*int32)(unsafe.Add(mBase, uint32(v76+v75<<(uint(int32(2))%32)))) = v51
												m.G0 = v11 + int32(16)
												return base.B2i32(int32(0) < v13)
											}
										}
									}
								}
							}
						} else {
							F_pfree(m, v43)
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = int32(0)
								m.G0 = v11 + int32(16)
								return base.B2i32(int32(0) < v13)
							}
						}
					}
				}
			}
		} else {
			v24 = v16
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
			v26 = v24 + v25
			v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+14)))
			v29 = v27 & int32(_a_F_gistPopItupFromNodeBuffer_0)
			v30 = F_palloc(m, v29)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v30
				if v29 != 0 {
					base.MemoryCopy(m, v30, v26+int32(8), v29)
				} else {
				}
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v36 + (v29+int32(7))&int32(_a_F_gistPopItupFromNodeBuffer_1)
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
				if v44 != int32(_a_F_gistPopItupFromNodeBuffer_2) {
					m.G0 = v11 + int32(16)
					return base.B2i32(int32(0) < v13)
				} else {
					v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v47 - int32(1)
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
					if v51 != int32(-1) {
						v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v56 = F_BufFileSeekBlock(m, v54, base.I64_extend_i32_s(v51))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int32(0)
						} else {
							if v56 != 0 {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v102 = m.ExcPending
								if v102 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v11))) = v51
									F_errmsg_internal(m, int32(_a_F_gistPopItupFromNodeBuffer_3), v11)
									mBase = m.M
									v106 = m.ExcPending
									if v106 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_gistPopItupFromNodeBuffer_4), int32(753), int32(_a_F_gistPopItupFromNodeBuffer_5))
										mBase = m.M
										v111 = m.ExcPending
										if v111 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								F_BufFileReadExact(m, v54, v43, int32(_a_F_gistPopItupFromNodeBuffer_6))
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return int32(0)
								} else {
									v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									if v61 < v62 {
										v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										v75 = v61
										v76 = v64
										*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v75 + int32(1)
										*(*int32)(unsafe.Add(mBase, uint32(v76+v75<<(uint(int32(2))%32)))) = v51
										m.G0 = v11 + int32(16)
										return base.B2i32(int32(0) < v13)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v62 << (uint(int32(1)) % 32)
										v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										v71 = F_repalloc(m, v68, v62<<(uint(int32(3))%32))
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v71
											v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
											v75 = v74
											v76 = v71
											*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v75 + int32(1)
											*(*int32)(unsafe.Add(mBase, uint32(v76+v75<<(uint(int32(2))%32)))) = v51
											m.G0 = v11 + int32(16)
											return base.B2i32(int32(0) < v13)
										}
									}
								}
							}
						}
					} else {
						F_pfree(m, v43)
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = int32(0)
							m.G0 = v11 + int32(16)
							return base.B2i32(int32(0) < v13)
						}
					}
				}
			}
		}
	}
}
func F_gistProcessEmptyingQueue(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
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
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v162 int32
	_ = v162
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v18 = v16
	goto L4
L2:
	;
	goto L3
L3:
	;
	m.G0 = v13 + int32(16)
	return
L4:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v29 = F_list_delete_first(m, v18)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	return
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v29
	v32 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+16)) = uint8(v32)
	v35 = m.G0
	v37 = v35 - int32(16)
	m.G0 = v37
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
	if v32 < v39 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v125 = F_gistPopItupFromNodeBuffer(m, v15, v28, v13+int32(12))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L6
	} else {
		goto L31
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L6
	} else {
		goto L27
	}
L10:
	;
	v43 = v39
	v48 = v32
	goto L13
L11:
	;
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = int32(0)
	m.G0 = v37 + int32(16)
	goto L8
L13:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v15)+48))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v52+v48<<(uint(int32(2))%32))))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	if v57 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L12
L15:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	if int32(0) < v58 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v89 = v43
	goto L17
L17:
	;
	v93 = v48 + int32(1)
	if v93 < v89 {
		v43 = v89
		v48 = v93
		goto L13
	} else {
		goto L26
	}
L18:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v77 = F_BufFileSeekBlock(m, v75, base.I64_extend_i32_s(v73))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L6
	} else {
		goto L22
	}
L19:
	;
	v62 = v58 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v62
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v64+v62<<(uint(int32(2))%32))))
	v73 = v68
	goto L18
L20:
	;
	goto L21
L21:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v69 + int32(1)
	v73 = v69
	goto L18
L22:
	;
	if v77 != 0 {
		goto L9
	} else {
		goto L23
	}
L23:
	;
	F_BufFileWrite(m, v75, v74, int32(_a_F_gistProcessEmptyingQueue_0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L6
	} else {
		goto L24
	}
L24:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	F_pfree(m, v82)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L6
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+8)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v56)+12)) = int32(0)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
	v89 = v88
	goto L17
L26:
	;
	goto L14
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37))) = v73
	F_errmsg_internal(m, int32(_a_F_gistProcessEmptyingQueue_1), v37)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L6
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(_a_F_gistProcessEmptyingQueue_2), int32(761), int32(_a_F_gistProcessEmptyingQueue_3))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L6
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
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	if v162 != 0 {
		v18 = v162
		goto L4
	} else {
		goto L40
	}
L31:
	;
	if v125 == int32(0) {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	goto L33
L33:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v28)+20))
	v142 = F_gistProcessItup(m, l0, v139, v140, v141)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L6
	} else {
		goto L35
	}
L34:
	;
	goto L30
L35:
	;
	if v142 != 0 {
		goto L30
	} else {
		goto L36
	}
L36:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)+4))
	F_MemoryContextReset(m, v145)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L6
	} else {
		goto L37
	}
L37:
	;
	v150 = F_gistPopItupFromNodeBuffer(m, v15, v28, v13+int32(12))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L6
	} else {
		goto L38
	}
L38:
	;
	if v150 != 0 {
		goto L33
	} else {
		goto L39
	}
L39:
	;
	goto L34
L40:
	;
	goto L5
}
func F_gistPushItupToNodeBuffer(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = int32(_a_F_gistPushItupToNodeBuffer_0)
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_gistPushItupToNodeBuffer[0]))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, _c_F_gistPushItupToNodeBuffer[0])) = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v54 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L2:
	;
	v20 = F_MemoryContextAllocZero(m, v16, int32(_a_F_gistPushItupToNodeBuffer_6))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = int64(35154307317759)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v20
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+17)))
	if v27 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v28 < v29 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43+v42<<(uint(int32(2))%32)))) = l1
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v48 + int32(1)
	goto L1
L7:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v42 = v28
	v43 = v31
	goto L6
L8:
	;
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v29 << (uint(int32(1)) % 32)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v38 = F_repalloc(m, v35, v29<<(uint(int32(3))%32))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v38
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v42 = v41
	v43 = v38
	goto L6
L11:
	;
	F_gistLoadNodeBuffer(m, l0, l1)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L3
	} else {
		goto L14
	}
L12:
	;
	v60 = v54
	goto L13
L13:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+6)))
	v64 = v62 & int32(_a_F_gistPushItupToNodeBuffer_1)
	if base.Ui32(v61) < base.Ui32((v64+int32(7))&int32(_a_F_gistPushItupToNodeBuffer_2)) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v60 = v59
	goto L13
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L3
	} else {
		goto L33
	}
L16:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if int32(0) < v70 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v108 = v60
	v109 = v64
	v110 = v61
	goto L18
L18:
	;
	v115 = v110 - (v109+int32(7))&int32(_a_F_gistPushItupToNodeBuffer_2)
	*(*int32)(unsafe.Add(mBase, uint32(v108)+4)) = v115
	if v109 != 0 {
		goto L26
	} else {
		goto L27
	}
L19:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v89 = F_BufFileSeekBlock(m, v87, base.I64_extend_i32_s(v85))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L3
	} else {
		goto L23
	}
L20:
	;
	v74 = v70 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v74
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v76+v74<<(uint(int32(2))%32))))
	v85 = v80
	goto L19
L21:
	;
	goto L22
L22:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v81 + int32(1)
	v85 = v81
	goto L19
L23:
	;
	if v89 != 0 {
		goto L15
	} else {
		goto L24
	}
L24:
	;
	F_BufFileWrite(m, v87, v86, int32(_a_F_gistPushItupToNodeBuffer_6))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L3
	} else {
		goto L25
	}
L25:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v94)+4)) = int32(_a_F_gistPushItupToNodeBuffer_7)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = v85
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v99 + int32(1)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	v105 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+6)))
	v108 = v103
	v109 = v105 & int32(_a_F_gistPushItupToNodeBuffer_1)
	v110 = v104
	goto L18
L26:
	;
	base.MemoryCopy(m, v108+v115+int32(8), l2, v109)
	goto L28
L27:
	;
	goto L28
L28:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v124 = base.I32_div_s(v122, int32(2))
	if v121 <= v124 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_gistPushItupToNodeBuffer[0])) = v14
	m.G0 = v11 + int32(16)
	return
L30:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	if v126 != 0 {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v128 = F_lcons(m, l1, v127)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L3
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v128
	v131 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)) = uint8(v131)
	goto L29
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v85
	F_errmsg_internal(m, int32(_a_F_gistPushItupToNodeBuffer_3), v11)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L3
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(_a_F_gistPushItupToNodeBuffer_4), int32(761), int32(_a_F_gistPushItupToNodeBuffer_5))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L3
	} else {
		goto L35
	}
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_gist_bbox_zorder_abbrev_abort(m *base.Module, l0 int32, l1 int32) int32 {
	return int32(0)
}
func F_gist_page_items(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v195 int32
	_ = v195
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v441 int32
	_ = v441
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v494 int32
	_ = v494
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v512 int32
	_ = v512
	v15 = m.G0
	v17 = v15 - int32(256)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = F_pg_detoast_datum(m, v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v26 = F_superuser(m)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L1
	} else {
		goto L107
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L1
	} else {
		goto L103
	}
L5:
	;
	if v26 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	F_InitMaterializedSRF(m, l0, int32(0))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L1
	} else {
		goto L99
	}
L9:
	;
	v32 = F_index_open(m, v25, int32(1))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v32)+48))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+84))
	if v35 != int32(783) {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v38 = F_verify_gist_page(m, v20)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	m.G0 = v17 + int32(256)
	return int32(0)
L13:
	;
	v40 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+14)))
	if v40 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	F_relation_close(m, v32, int32(1))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v48 = int32(1)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v32)+52))
	v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+16)))
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38+v50)+12)))
	if v52&v48 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v46 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v46)
	goto L12
L18:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v32)+192))
	v59 = int32(*(*int16)(unsafe.Add(mBase, uint32(v58)+10)))
	v60 = F_CreateTupleDescTruncatedCopy(m, v49, v59)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	v62 = v48
	v63 = v49
	goto L20
L20:
	;
	v64 = int32(0)
	v66 = int32(1)
	v67 = int32(2)
	if v62&v66 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v62 = int32(3)
	v63 = v60
	goto L20
L22:
	;
	v77 = int32(7)
	goto L24
L23:
	;
	v77 = v67
	goto L24
L24:
	;
	v79 = F_pg_get_indexdef_worker(m, v25, v64, v64, v66, int32(base.Ui32(v62&v67)>>(uint(v66)%32)), v64, v64, v77, int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v81 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+16)))
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38+v81)+12)))
	if v83&int32(2) != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	F_relation_close(m, v32, int32(1))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L1
	} else {
		goto L98
	}
L27:
	;
	v88 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+12)))
	if base.Ui32(v101) < base.Ui32(int32(25)) {
		goto L26
	} else {
		goto L34
	}
L30:
	;
	if v88 == int32(0) {
		goto L26
	} else {
		goto L31
	}
L31:
	;
	F_errmsg_internal(m, int32(_a_F_gist_page_items_0), int32(0))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(_a_F_gist_page_items_1), int32(255), int32(_a_F_gist_page_items_2))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	goto L26
L34:
	;
	v105 = v101 + int32(_a_F_gist_page_items_3)
	if v105&int32(_a_F_gist_page_items_4) == int32(0) {
		goto L26
	} else {
		goto L35
	}
L35:
	;
	v116 = int32(1)
	v118 = v116
	v129 = v116
	goto L36
L36:
	;
	v134 = v38 + int32(20) + v118<<(uint(int32(2))%32)
	if v134 == int32(0) {
		goto L3
	} else {
		goto L38
	}
L37:
	;
	goto L26
L38:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	v140 = v38 + v137&int32(_a_F_gist_page_items_5)
	F_index_deform_tuple(m, v140, v63, v17+int32(80), v17+int32(48))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v147 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+220)) = uint8(v147)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+216)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v17)+228)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v17)+224)) = v118
	v154 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v140)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+232)) = v154 & int32(_a_F_gist_page_items_6)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	v159 = int32(_a_F_gist_page_items_7)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+236)) = base.B2i32(v158&v159 == v159)
	if v79 == v147 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+220)) = uint8(v396)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+240)) = v400
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	F_tuplestore_putvalues(m, v412, v413, v17+int32(224), v17+int32(216))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L1
	} else {
		goto L96
	}
L41:
	;
	v396 = int32(1)
	v400 = v147
	goto L40
L42:
	;
	goto L43
L43:
	;
	v168 = v17 + int32(32)
	F_initStringInfo(m, v168)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v79
	F_appendStringInfo(m, v168, int32(_a_F_gist_page_items_8), v17)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	if int32(0) < v175 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v178 = v175
	v182 = v147
	goto L49
L47:
	;
	goto L48
L48:
	;
	F_appendStringInfoChar(m, v17+int32(32), int32(41))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L1
	} else {
		goto L94
	}
L49:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17+int32(48)+v182))))
	if v195 != 0 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	goto L48
L51:
	;
	v219 = int32(_a_F_gist_page_items_9)
	goto L53
L52:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v63+v178<<(uint(int32(4))%32)+v182*int32(100))+88))
	F_getTypeOutputInfo(m, v203, v17+int32(28), v17+int32(27))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L54
	}
L53:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v32)+192))
	v223 = int32(*(*int16)(unsafe.Add(mBase, uint32(v222)+10)))
	if v223 == v182 {
		goto L57
	} else {
		goto L58
	}
L54:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v17+int32(80)+v182<<(uint(int32(2))%32))))
	v217 = F_OidOutputFunctionCall(m, v210, v216)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v219 = v217
	goto L53
L56:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219))))
	v235 = v219
	v236 = v232
	goto L64
L57:
	;
	v229 = int32(_a_F_gist_page_items_10)
	goto L59
L58:
	;
	if v182 == int32(0) {
		goto L56
	} else {
		goto L60
	}
L59:
	;
	F_appendStringInfoString(m, v17+int32(32), v229)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L61
	}
L60:
	;
	v229 = int32(_a_F_gist_page_items_11)
	goto L59
L61:
	;
	goto L56
L62:
	;
	v277 = v219
	goto L74
L63:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v17)+40))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v17)+36))
	if v251 <= v252+int32(1) {
		goto L69
	} else {
		goto L70
	}
L64:
	;
	switch v236 {
	case 0:
		goto L66
	default:
		goto L67
	case 9, 10, 11, 12, 13, 32, 34, 40, 41, 44, 92:
		goto L63
	}
L65:
	;
	if v232 != 0 {
		v276 = int32(0)
		goto L62
	} else {
		goto L68
	}
L66:
	;
	goto L65
L67:
	;
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235)+1)))
	v235 = v235 + int32(1)
	v236 = v247
	goto L64
L68:
	;
	goto L63
L69:
	;
	F_appendStringInfoChar(m, v17+int32(32), int32(34))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	v264 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v262+v252))) = uint8(v264)
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v17)+36))
	v267 = int32(1)
	v268 = v266 + v267
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v268
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	v272 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v270+v268))) = uint8(v272)
	v276 = v267
	goto L62
L72:
	;
	v276 = int32(1)
	goto L62
L73:
	;
	v370 = v182 + int32(1)
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	if v370 < v371 {
		v178 = v371
		v182 = v370
		goto L49
	} else {
		goto L93
	}
L74:
	;
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v277))))
	v292 = base.I32_extend8_s(v291)
	if base.B2i32(v291 == int32(34))|base.B2i32(v291 == int32(92)) == int32(0) {
		goto L78
	} else {
		goto L79
	}
L75:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	v358 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v356+v303))) = uint8(v358)
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v17)+36))
	v362 = v360 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v362
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	v366 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v364+v362))) = uint8(v366)
	goto L73
L76:
	;
	goto L75
L77:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v17)+40))
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v17)+36))
	if v333 <= v334+int32(1) {
		goto L89
	} else {
		goto L90
	}
L78:
	;
	if v291 != 0 {
		goto L77
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v17)+40))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v17)+36))
	if v312 <= v313+int32(1) {
		goto L85
	} else {
		goto L86
	}
L81:
	;
	if v276 == int32(0) {
		goto L73
	} else {
		goto L82
	}
L82:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v17)+40))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v17)+36))
	if v303+int32(1) < v302 {
		goto L76
	} else {
		goto L83
	}
L83:
	;
	F_appendStringInfoChar(m, v17+int32(32), int32(34))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	goto L73
L85:
	;
	F_appendStringInfoChar(m, v17+int32(32), v292)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v321+v313))) = uint8(v292)
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v17)+36))
	v326 = v324 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v326
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	v330 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v328+v326))) = uint8(v330)
	goto L77
L88:
	;
	goto L77
L89:
	;
	F_appendStringInfoChar(m, v17+int32(32), v292)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L1
	} else {
		goto L92
	}
L90:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v342+v334))) = uint8(v292)
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v17)+36))
	v347 = v345 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v347
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	v351 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v349+v347))) = uint8(v351)
	goto L91
L91:
	;
	v277 = v277 + int32(1)
	goto L74
L92:
	;
	goto L91
L93:
	;
	goto L50
L94:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	v394 = F_cstring_to_text(m, v393)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	v396 = int32(0)
	v400 = v394
	goto L40
L96:
	;
	v421 = v129 + int32(1)
	v423 = v421 & int32(_a_F_gist_page_items_12)
	if base.Ui32(v423) <= base.Ui32(int32(base.Ui32(v105)>>(uint(int32(2))%32))&int32(_a_F_gist_page_items_12)) {
		v118 = v423
		v129 = v421
		goto L36
	} else {
		goto L97
	}
L97:
	;
	goto L37
L98:
	;
	goto L12
L99:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	F_errmsg(m, int32(_a_F_gist_page_items_13), int32(0))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	F_errfinish(m, int32(_a_F_gist_page_items_1), int32(211), int32(_a_F_gist_page_items_2))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L103:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v32)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = int32(_a_F_gist_page_items_14)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v484 + int32(4)
	F_errmsg(m, int32(_a_F_gist_page_items_15), v17+int32(16))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(_a_F_gist_page_items_1), int32(222), int32(_a_F_gist_page_items_2))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L1
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
	F_errmsg_internal(m, int32(_a_F_gist_page_items_16), int32(0))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(_a_F_gist_page_items_1), int32(275), int32(_a_F_gist_page_items_2))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_gist_page_items_bytea(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = F_pg_detoast_datum(m, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v22 = F_superuser(m)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L36
	}
L4:
	;
	if v22 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	F_InitMaterializedSRF(m, l0, int32(0))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L32
	}
L8:
	;
	v27 = F_verify_gist_page(m, v17)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	m.G0 = v14 + int32(48)
	return int32(0)
L10:
	;
	v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27)+14)))
	if v29 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v32 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v32)
	goto L9
L12:
	;
	goto L13
L13:
	;
	v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27)+16)))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+v34)+12)))
	if v36&int32(2) != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v41 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27)+12)))
	if base.Ui32(v54) < base.Ui32(int32(25)) {
		goto L9
	} else {
		goto L21
	}
L17:
	;
	if v41 == int32(0) {
		goto L9
	} else {
		goto L18
	}
L18:
	;
	F_errmsg_internal(m, int32(_a_F_gist_page_items_bytea_0), int32(0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(_a_F_gist_page_items_bytea_1), int32(152), int32(_a_F_gist_page_items_bytea_2))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	goto L9
L21:
	;
	v58 = v54 + int32(_a_F_gist_page_items_bytea_3)
	if v58&int32(_a_F_gist_page_items_bytea_4) == int32(0) {
		goto L9
	} else {
		goto L22
	}
L22:
	;
	v69 = int32(1)
	v71 = v69
	v76 = v69
	goto L23
L23:
	;
	v84 = v27 + int32(20) + v71<<(uint(int32(2))%32)
	if v84 == int32(0) {
		goto L3
	} else {
		goto L25
	}
L24:
	;
	goto L9
L25:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v90 = v27 + v87&int32(_a_F_gist_page_items_bytea_5)
	v91 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v90)+6)))
	v92 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+12)) = uint8(v92)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v92
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v71
	v99 = v91 & int32(_a_F_gist_page_items_bytea_6)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v99
	v102 = v99 + int32(4)
	v103 = F_palloc(m, v102)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v103))) = v102 << (uint(int32(2)) % 32)
	if v99 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	base.MemoryCopy(m, v103+int32(4), v90, v99)
	goto L29
L28:
	;
	goto L29
L29:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v103
	v113 = int32(_a_F_gist_page_items_bytea_7)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = base.B2i32(v111&v113 == v113)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
	F_tuplestore_putvalues(m, v118, v119, v14+int32(16), v14+int32(8))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v127 = v76 + int32(1)
	v129 = v127 & int32(_a_F_gist_page_items_bytea_8)
	if base.Ui32(v129) <= base.Ui32(int32(base.Ui32(v58)>>(uint(int32(2))%32))&int32(_a_F_gist_page_items_bytea_8)) {
		v71 = v129
		v76 = v127
		goto L23
	} else {
		goto L31
	}
L31:
	;
	goto L24
L32:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_errmsg(m, int32(_a_F_gist_page_items_bytea_9), int32(0))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(_a_F_gist_page_items_bytea_1), int32(141), int32(_a_F_gist_page_items_bytea_2))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L36:
	;
	F_errmsg_internal(m, int32(_a_F_gist_page_items_bytea_10), int32(0))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(_a_F_gist_page_items_bytea_1), int32(170), int32(_a_F_gist_page_items_bytea_2))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_gist_page_opaque_info(m *base.Module, l0 int32) int32 {
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
	var v14 int32
	_ = v14
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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v107 int64
	_ = v107
	var v108 int64
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int64
	_ = v117
	var v120 int64
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int64
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	v6 = m.G0
	v8 = v6 - int32(112)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v15 = F_superuser(m)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L49
	}
L4:
	;
	if v15 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v17 = F_verify_gist_page(m, v11)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L45
	}
L8:
	;
	m.G0 = v8 + int32(112)
	return v148
L9:
	;
	v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+14)))
	if v19 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v22 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v22)
	v148 = int32(0)
	goto L8
L11:
	;
	goto L12
L12:
	;
	v25 = int32(0)
	v29 = F_get_call_result_type(m, l0, v25, v8+int32(108))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	if v29 != int32(1) {
		goto L3
	} else {
		goto L14
	}
L14:
	;
	v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+16)))
	v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17+v33)+12)))
	if v35&int32(1) != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v41 = F_cstring_to_text(m, int32(_a_F_gist_page_opaque_info_3))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	v45 = v8
	v46 = v25
	goto L17
L17:
	;
	if v35&int32(2) != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v41
	v45 = v8 | int32(4)
	v46 = int32(1)
	goto L17
L19:
	;
	v50 = F_cstring_to_text(m, int32(_a_F_gist_page_opaque_info_4))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	v55 = v46
	goto L21
L21:
	;
	if v35&int32(4) != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45))) = v50
	v55 = v46 + int32(1)
	goto L21
L23:
	;
	v62 = F_cstring_to_text(m, int32(_a_F_gist_page_opaque_info_5))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	v67 = v55
	goto L25
L25:
	;
	if v35&int32(8) != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8|v55<<(uint(int32(2))%32)))) = v62
	v67 = v55 + int32(1)
	goto L25
L27:
	;
	v74 = F_cstring_to_text(m, int32(_a_F_gist_page_opaque_info_6))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	v79 = v67
	goto L29
L29:
	;
	if v35&int32(16) != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8+v67<<(uint(int32(2))%32)))) = v74
	v79 = v67 + int32(1)
	goto L29
L31:
	;
	v86 = F_cstring_to_text(m, int32(_a_F_gist_page_opaque_info_7))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L34
	}
L32:
	;
	v91 = v79
	goto L33
L33:
	;
	v93 = v35 & int32(_a_F_gist_page_opaque_info_8)
	if v93 != 0 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8+v79<<(uint(int32(2))%32)))) = v86
	v91 = v79 + int32(1)
	goto L33
L35:
	;
	v99 = F_DirectFunctionCall1Coll(m, int32(2895), int32(0), v93)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	v104 = v91
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+76)) = int32(0)
	v107 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v17)+4)))
	v108 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v17))))
	v112 = F_Int64GetDatum(m, v107|v108<<(uint(int64(32))%64))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L39
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8+v91<<(uint(int32(2))%32)))) = v99
	v104 = v91 + int32(1)
	goto L37
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+80)) = v112
	v115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+16)))
	v116 = v17 + v115
	v117 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v116))))
	v120 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v116)+4)))
	v122 = F_Int64GetDatum(m, v117<<(uint(int64(32))%64)|v120)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+84)) = v122
	v125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+16)))
	v127 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v17+v125)+8)))
	v128 = F_Int64GetDatum(m, v127)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+88)) = v128
	v132 = F_construct_array_builtin(m, v8, v104, int32(25))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+92)) = v132
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v8)+108))
	v140 = F_heap_form_tuple(m, v135, v8+int32(80), v8+int32(76))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v140)+16))
	v143 = F_HeapTupleHeaderGetDatum(m, v142)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v148 = v143
	goto L8
L45:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_errmsg(m, int32(_a_F_gist_page_opaque_info_9), int32(0))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_gist_page_opaque_info_1), int32(86), int32(_a_F_gist_page_opaque_info_2))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L49:
	;
	F_errmsg_internal(m, int32(_a_F_gist_page_opaque_info_0), int32(0))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(_a_F_gist_page_opaque_info_1), int32(95), int32(_a_F_gist_page_opaque_info_2))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_gist_point_sortsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v15 int32
	_ = v15
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+20)))
	if v5 == int32(1) {
		*(*int32)(unsafe.Add(mBase, uint32(v4)+32)) = int32(113)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+28)) = int32(114)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+24)) = int32(115)
		v15 = int32(116)
	} else {
		v15 = int32(113)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v4)+16)) = v15
	return int32(0)
}
func F_gist_redo(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int64
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
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
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int64
	_ = v257
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int64
	_ = v273
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v325 int64
	_ = v325
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v343 int64
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v347 int64
	_ = v347
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int64
	_ = v355
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v376 int32
	_ = v376
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v473 int32
	_ = v473
	var v481 int32
	_ = v481
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v521 int32
	_ = v521
	var v527 int32
	_ = v527
	var v545 int32
	_ = v545
	var v551 int32
	_ = v551
	var v558 int32
	_ = v558
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v579 int32
	_ = v579
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v614 int32
	_ = v614
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v649 int32
	_ = v649
	var v651 int64
	_ = v651
	var v653 int32
	_ = v653
	var v655 int64
	_ = v655
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v671 int32
	_ = v671
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v687 int64
	_ = v687
	var v689 int32
	_ = v689
	var v691 int64
	_ = v691
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v700 int32
	_ = v700
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v719 int32
	_ = v719
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v737 int32
	_ = v737
	var v752 int32
	_ = v752
	var v755 int32
	_ = v755
	var v760 int32
	_ = v760
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v764 int64
	_ = v764
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v772 int64
	_ = v772
	var v773 int32
	_ = v773
	var v777 int32
	_ = v777
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v796 int32
	_ = v796
	var v799 int32
	_ = v799
	var v803 int64
	_ = v803
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v818 int32
	_ = v818
	var v822 int32
	_ = v822
	var v828 int32
	_ = v828
	var v830 int32
	_ = v830
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v839 int32
	_ = v839
	var v843 int32
	_ = v843
	var v845 int32
	_ = v845
	var v847 int32
	_ = v847
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v854 int32
	_ = v854
	var v882 int32
	_ = v882
	var v884 int32
	_ = v884
	var v891 int32
	_ = v891
	var v897 int32
	_ = v897
	var v902 int32
	_ = v902
	var v906 int32
	_ = v906
	var v912 int32
	_ = v912
	var v917 int32
	_ = v917
	var v921 int32
	_ = v921
	var v925 int32
	_ = v925
	var v930 int32
	_ = v930
	v2 = int32(0)
	v25 = m.G0
	v27 = v25 - int32(96)
	m.G0 = v27
	v29 = int32(_a_F_gist_redo_0)
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_gist_redo[0]))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+48)))
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_gist_redo[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_gist_redo[0])) = v35
	v38 = v32 & int32(240)
	switch int32(base.Ui32(v38) >> (uint(int32(4)) % 32)) {
	case 0:
		goto L9
	case 1:
		goto L8
	case 2:
		goto L7
	case 3:
		goto L6
	default:
		goto L1
	case 6:
		goto L5
	case 7:
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v921 = m.ExcPending
	if v921 != 0 {
		goto L10
	} else {
		goto L173
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v906 = m.ExcPending
	if v906 != 0 {
		goto L10
	} else {
		goto L170
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L10
	} else {
		goto L167
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_gist_redo[0])) = v30
	v882 = *(*int32)(unsafe.Add(mBase, _c_F_gist_redo[1]))
	F_MemoryContextReset(m, v882)
	mBase = m.M
	v884 = m.ExcPending
	if v884 != 0 {
		goto L10
	} else {
		goto L166
	}
L5:
	;
	v763 = *(*int32)(unsafe.Add(mBase, uint32(v31)+64))
	v764 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v768 = F_XLogReadBufferForRedo(m, l0, int32(0), v27+int32(92))
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L10
	} else {
		goto L141
	}
L6:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v31)+64))
	v354 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v353)+18)))
	if v354 != 0 {
		goto L75
	} else {
		goto L76
	}
L7:
	;
	v339 = *(*int32)(unsafe.Add(mBase, _c_F_gist_redo[2]))
	if base.Ui32(v339) < base.Ui32(int32(2)) {
		goto L4
	} else {
		goto L73
	}
L8:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v31)+64))
	v257 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v259 = *(*int32)(unsafe.Add(mBase, _c_F_gist_redo[2]))
	if base.Ui32(int32(2)) <= base.Ui32(v259) {
		goto L56
	} else {
		goto L57
	}
L9:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v31)+64))
	v42 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v46 = F_XLogReadBufferForRedo(m, l0, int32(0), v27+int32(76))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return
L11:
	;
	if v46 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v50 = int32(0)
	v52 = v27 + int32(92)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+72))
	if v55 < v50 {
		v77 = v50
		goto L16
	} else {
		goto L17
	}
L13:
	;
	goto L14
L14:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v241)+72))
	if v242 <= int32(0) {
		goto L50
	} else {
		goto L51
	}
L15:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v27)+76))
	if v81 < int32(0) {
		goto L27
	} else {
		goto L28
	}
L16:
	;
	v80 = v77
	goto L15
L17:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54+int32(0))+76)))
	if v60 != int32(1) {
		v77 = v50
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v64 = v54 + int32(76)
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+43)))
	if v65 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	if v52 == int32(0) {
		v77 = v50
		goto L16
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if v52 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v70 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v70
	v80 = v70
	goto L15
L23:
	;
	v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v64)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v73
	goto L25
L24:
	;
	goto L25
L25:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v64)+44))
	v77 = v75
	goto L16
L26:
	;
	v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41))))
	switch v100 {
	case 0:
		v130 = v80
		goto L30
	case 1:
		goto L32
	default:
		goto L31
	}
L27:
	;
	v85 = *(*int32)(unsafe.Add(mBase, _c_F_gist_redo[3]))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v85+(v81^int32(-1))<<(uint(int32(2))%32))))
	v99 = v91
	goto L26
L28:
	;
	goto L29
L29:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_gist_redo[4]))
	v99 = v93 + v81<<(uint(int32(13))%32) + int32(-8192)
	goto L26
L30:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v27)+92))
	if base.Ui32(v130-v80) < base.Ui32(v133) {
		goto L38
	} else {
		goto L39
	}
L31:
	;
	F_PageIndexMultiDelete(m, v99, v80, v100)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L10
	} else {
		goto L36
	}
L32:
	;
	v101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+2)))
	if v101 != int32(1) {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v80))))
	v106 = v80 + int32(2)
	v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v80)+8)))
	v109 = v107 & int32(_a_F_gist_redo_1)
	v110 = F_PageIndexTupleOverwrite(m, v99, v104, v106, v109)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L10
	} else {
		goto L34
	}
L34:
	;
	if v110 == int32(0) {
		goto L3
	} else {
		goto L35
	}
L35:
	;
	v130 = v109 + v106
	goto L30
L36:
	;
	v117 = int32(1)
	v119 = v80 + v100<<(uint(v117)%32)
	v120 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+16)))
	v121 = v99 + v120
	v122 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v121)+12)))
	if v122&v117 == int32(0) {
		v130 = v119
		goto L30
	} else {
		goto L37
	}
L37:
	;
	v128 = v122 | int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v121)+12)) = uint16(v128)
	v130 = v119
	goto L30
L38:
	;
	v136 = int32(1)
	v137 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+12)))
	if base.Ui32(v137) < base.Ui32(int32(25)) {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v99))) = base.I64_rotr(v42, int64(32))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v27)+76))
	F_MarkBufferDirty(m, v214)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L10
	} else {
		goto L49
	}
L41:
	;
	v146 = v136
	goto L43
L42:
	;
	v146 = int32(base.Ui32(v137+int32(_a_F_gist_redo_2))>>(uint(int32(2))%32)) + v136
	goto L43
L43:
	;
	v148 = v130
	v154 = v146
	goto L44
L44:
	;
	v171 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v148)+6)))
	v173 = v171 & int32(_a_F_gist_redo_1)
	v177 = F_PageAddItemExtended(m, v99, v148, v173, v154&int32(_a_F_gist_redo_3), int32(0))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L10
	} else {
		goto L46
	}
L45:
	;
	goto L40
L46:
	;
	if v177 == int32(0) {
		goto L2
	} else {
		goto L47
	}
L47:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v27)+92))
	v184 = v148 + v173
	if base.Ui32(v184-v80) < base.Ui32(v183) {
		v148 = v184
		v154 = v154 + int32(1)
		goto L44
	} else {
		goto L48
	}
L48:
	;
	goto L45
L49:
	;
	goto L14
L50:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v27)+76))
	if v251 == int32(0) {
		goto L4
	} else {
		goto L54
	}
L51:
	;
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241)+128)))
	if v245 != int32(1) {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	F_gistRedoClearFollowRight(m, l0, int32(1))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L10
	} else {
		goto L53
	}
L53:
	;
	goto L50
L54:
	;
	F_UnlockReleaseBuffer(m, v251)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L10
	} else {
		goto L55
	}
L55:
	;
	goto L4
L56:
	;
	v262 = int32(0)
	F_XLogRecGetBlockTag(m, l0, v262, v27+int32(76), v262, v262)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L10
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v283 = F_XLogReadBufferForRedo(m, l0, int32(0), v27+int32(76))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L10
	} else {
		goto L61
	}
L59:
	;
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v256)+6)))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v27)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+56)) = v271
	v273 = *(*int64)(unsafe.Add(mBase, uint32(v27)+76))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+48)) = v273
	F_ResolveRecoveryConflictWithSnapshot(m, v270, v269, v27+int32(48))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L10
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	if v283 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v27)+76))
	if v289 < int32(0) {
		goto L66
	} else {
		goto L67
	}
L63:
	;
	goto L64
L64:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v27)+76))
	if v333 == int32(0) {
		goto L4
	} else {
		goto L71
	}
L65:
	;
	v308 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v256)+4)))
	F_PageIndexMultiDelete(m, v307, v256+int32(8), v308)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L10
	} else {
		goto L69
	}
L66:
	;
	v293 = *(*int32)(unsafe.Add(mBase, _c_F_gist_redo[3]))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v293+(v289^int32(-1))<<(uint(int32(2))%32))))
	v307 = v299
	goto L65
L67:
	;
	goto L68
L68:
	;
	v301 = *(*int32)(unsafe.Add(mBase, _c_F_gist_redo[4]))
	v307 = v301 + v289<<(uint(int32(13))%32) + int32(-8192)
	goto L65
L69:
	;
	v311 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v307)+16)))
	v312 = v307 + v311
	v313 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v312)+12)))
	v315 = v313 & int32(_a_F_gist_redo_4)
	*(*uint16)(unsafe.Add(mBase, uint32(v312)+12)) = uint16(v315)
	v317 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v307)+16)))
	v318 = v307 + v317
	v319 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v318)+12)))
	v321 = v319 | int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v318)+12)) = uint16(v321)
	*(*uint32)(unsafe.Add(mBase, uint32(v307)+4)) = uint32(v257)
	v325 = int64(base.Ui64(v257) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v307))) = uint32(v325)
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v27)+76))
	F_MarkBufferDirty(m, v327)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L10
	} else {
		goto L70
	}
L70:
	;
	goto L64
L71:
	;
	F_UnlockReleaseBuffer(m, v333)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L10
	} else {
		goto L72
	}
L72:
	;
	goto L4
L73:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v31)+64))
	v343 = *(*int64)(unsafe.Add(mBase, uint32(v342)+16))
	v344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v342)+24)))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v342)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+72)) = v345
	v347 = *(*int64)(unsafe.Add(mBase, uint32(v342)))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+64)) = v347
	F_ResolveRecoveryConflictWithSnapshotFullXid(m, v343, v344, v27-int32(-64))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L10
	} else {
		goto L74
	}
L74:
	;
	goto L4
L75:
	;
	v355 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v365 = v2
	v369 = v2
	v376 = v2
	goto L78
L76:
	;
	v729 = v31
	v737 = v2
	goto L77
L77:
	;
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v729)+72))
	if v752 < int32(0) {
		goto L136
	} else {
		goto L137
	}
L78:
	;
	v385 = v365 + int32(1)
	v387 = v385 & int32(255)
	v388 = int32(0)
	F_XLogRecGetBlockTag(m, l0, v387, v388, v388, v27+int32(92))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L10
	} else {
		goto L80
	}
L79:
	;
	v727 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v729 = v727
	v737 = v724
	goto L77
L80:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v27)+92))
	v395 = F_XLogInitBufferForRedo(m, l0, v387)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L10
	} else {
		goto L82
	}
L81:
	;
	v415 = int32(0)
	v418 = v27 + int32(76)
	v420 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v420)+72))
	if v421 < v387 {
		v443 = v415
		goto L87
	} else {
		goto L88
	}
L82:
	;
	if v395 < int32(0) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v400 = *(*int32)(unsafe.Add(mBase, _c_F_gist_redo[3]))
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v400+(v395^int32(-1))<<(uint(int32(2))%32))))
	v414 = v406
	goto L81
L84:
	;
	goto L85
L85:
	;
	v408 = *(*int32)(unsafe.Add(mBase, _c_F_gist_redo[4]))
	v414 = v408 + v395<<(uint(int32(13))%32) + int32(-8192)
	goto L81
L86:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v446)))
	v450 = F_palloc(m, v447<<(uint(int32(2))%32))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L10
	} else {
		goto L97
	}
L87:
	;
	v446 = v443
	goto L86
L88:
	;
	v425 = v420 + v387*int32(52)
	v426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v425)+76)))
	if v426 != int32(1) {
		v443 = v415
		goto L87
	} else {
		goto L89
	}
L89:
	;
	v430 = v425 + int32(76)
	v431 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v430)+43)))
	if v431 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	if v418 == int32(0) {
		v443 = v415
		goto L87
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	if v418 != 0 {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	v436 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v418))) = v436
	v446 = v436
	goto L86
L94:
	;
	v439 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v430)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(v418))) = v439
	goto L96
L95:
	;
	goto L96
L96:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v430)+44))
	v443 = v441
	goto L87
L97:
	;
	if v447 <= int32(0) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v605 = v376 | base.B2i32(v394 == v415)
	v606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v353)+16)))
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v27)+92))
	v608 = int32(0)
	v610 = v606 & base.B2i32(v607 != v608)
	if v395 < v608 {
		goto L113
	} else {
		goto L114
	}
L99:
	;
	v455 = v447 & int32(3)
	v456 = int32(4)
	v457 = v446 + v456
	if base.Ui32(v447) < base.Ui32(v456) {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	v545 = v521
	v551 = v527
	v558 = int32(0)
	goto L108
L101:
	;
	v521 = v457
	v527 = int32(0)
	goto L100
L102:
	;
	goto L103
L103:
	;
	v464 = int32(0)
	v467 = v457
	v473 = v464
	v481 = v464
	goto L104
L104:
	;
	v492 = v450 + v473<<(uint(int32(2))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v492))) = v467
	v494 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v467)+6)))
	v495 = int32(_a_F_gist_redo_1)
	v497 = v467 + v494&v495
	*(*int32)(unsafe.Add(mBase, uint32(v492)+4)) = v497
	v499 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v497)+6)))
	v502 = v497 + v499&v495
	*(*int32)(unsafe.Add(mBase, uint32(v492)+8)) = v502
	v504 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v502)+6)))
	v507 = v502 + v504&v495
	*(*int32)(unsafe.Add(mBase, uint32(v492)+12)) = v507
	v509 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v507)+6)))
	v512 = v507 + v509&v495
	v513 = int32(4)
	v514 = v473 + v513
	v516 = v481 + v513
	if v516 != v447&int32(2147483644) {
		v467 = v512
		v473 = v514
		v481 = v516
		goto L104
	} else {
		goto L106
	}
L105:
	;
	if v455 == int32(0) {
		goto L98
	} else {
		goto L107
	}
L106:
	;
	goto L105
L107:
	;
	v521 = v512
	v527 = v514
	goto L100
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v450+v551<<(uint(int32(2))%32)))) = v545
	v572 = int32(1)
	v574 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v545)+6)))
	v579 = v558 + v572
	if v579 != v455 {
		v545 = v545 + v574&int32(_a_F_gist_redo_1)
		v551 = v551 + v572
		v558 = v579
		goto L108
	} else {
		goto L110
	}
L109:
	;
	goto L98
L110:
	;
	goto L109
L111:
	;
	F_gistfillbuffer(m, v414, v450, v447, int32(1))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L10
	} else {
		goto L116
	}
L112:
	;
	F_PageInit(m, v628, int32(_a_F_gist_redo_5), int32(16))
	mBase = m.M
	v632 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v628)+16)))
	v633 = v628 + v632
	v634 = int32(_a_F_gist_redo_6)
	*(*uint16)(unsafe.Add(mBase, uint32(v633)+14)) = uint16(v634)
	*(*uint16)(unsafe.Add(mBase, uint32(v633)+12)) = uint16(v610)
	*(*int32)(unsafe.Add(mBase, uint32(v633)+8)) = int32(-1)
	goto L111
L113:
	;
	v614 = *(*int32)(unsafe.Add(mBase, _c_F_gist_redo[3]))
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v614+(v395^int32(-1))<<(uint(int32(2))%32))))
	v628 = v620
	goto L112
L114:
	;
	goto L115
L115:
	;
	v622 = *(*int32)(unsafe.Add(mBase, _c_F_gist_redo[4]))
	v628 = v622 + v395<<(uint(int32(13))%32) + int32(-8192)
	goto L112
L116:
	;
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v27)+92))
	if v642 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v414)+4)) = base.I32_wrap_i64(v355)
	*(*int32)(unsafe.Add(mBase, uint32(v414))) = base.I32_wrap_i64(int64(base.Ui64(v355) >> (uint(int64(32)) % 64)))
	F_MarkBufferDirty(m, v395)
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L10
	} else {
		goto L129
	}
L118:
	;
	v645 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v414)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v414+v645)+8)) = int32(-1)
	v649 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v414)+16)))
	v651 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v353)+12)))
	*(*uint32)(unsafe.Add(mBase, uint32(v414+v649))) = uint32(v651)
	v653 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v414)+16)))
	v655 = *(*int64)(unsafe.Add(mBase, uint32(v353)+8))
	*(*uint32)(unsafe.Add(mBase, uint32(v414+v653)+4)) = uint32(v655)
	v657 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v414)+16)))
	v658 = v414 + v657
	v659 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v658)+12)))
	v661 = v659 & int32(_a_F_gist_redo_7)
	*(*uint16)(unsafe.Add(mBase, uint32(v658)+12)) = uint16(v661)
	goto L117
L119:
	;
	goto L120
L120:
	;
	v663 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v353)+18)))
	if v365 < v663-int32(1) {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	v685 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v414)+16)))
	v687 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v353)+12)))
	*(*uint32)(unsafe.Add(mBase, uint32(v414+v685))) = uint32(v687)
	v689 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v414)+16)))
	v691 = *(*int64)(unsafe.Add(mBase, uint32(v353)+8))
	*(*uint32)(unsafe.Add(mBase, uint32(v414+v689)+4)) = uint32(v691)
	v693 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v353)+18)))
	v694 = int32(1)
	if (base.B2i32(v693-v694 <= v365)|v605)&v694 != 0 {
		goto L126
	} else {
		goto L127
	}
L122:
	;
	v671 = int32(0)
	F_XLogRecGetBlockTag(m, l0, (v365+int32(2))&int32(255), v671, v671, v27+int32(88))
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L10
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	v681 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v414)+16)))
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v353)))
	*(*int32)(unsafe.Add(mBase, uint32(v414+v681)+8)) = v683
	goto L121
L125:
	;
	v677 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v414)+16)))
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v27)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v414+v677)+8)) = v679
	goto L121
L126:
	;
	v709 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v414)+16)))
	v710 = v414 + v709
	v711 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v710)+12)))
	v713 = v711 & int32(_a_F_gist_redo_7)
	*(*uint16)(unsafe.Add(mBase, uint32(v710)+12)) = uint16(v713)
	goto L117
L127:
	;
	v700 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v353)+20)))
	if v700 != int32(1) {
		goto L126
	} else {
		goto L128
	}
L128:
	;
	v703 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v414)+16)))
	v704 = v414 + v703
	v705 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v704)+12)))
	v707 = v705 | int32(8)
	*(*uint16)(unsafe.Add(mBase, uint32(v704)+12)) = uint16(v707)
	goto L117
L129:
	;
	if v365 == int32(0) {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	v725 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v353)+18)))
	if base.Ui32(v385) < base.Ui32(v725) {
		v365 = v385
		v369 = v724
		v376 = v605
		goto L78
	} else {
		goto L135
	}
L131:
	;
	v724 = v395
	goto L130
L132:
	;
	goto L133
L133:
	;
	F_UnlockReleaseBuffer(m, v395)
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L10
	} else {
		goto L134
	}
L134:
	;
	v724 = v369
	goto L130
L135:
	;
	goto L79
L136:
	;
	F_UnlockReleaseBuffer(m, v737)
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L10
	} else {
		goto L140
	}
L137:
	;
	v755 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v729)+76)))
	if v755 != int32(1) {
		goto L136
	} else {
		goto L138
	}
L138:
	;
	F_gistRedoClearFollowRight(m, l0, int32(0))
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L10
	} else {
		goto L139
	}
L139:
	;
	goto L136
L140:
	;
	goto L4
L141:
	;
	if v768 == int32(0) {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v772 = *(*int64)(unsafe.Add(mBase, uint32(v763)))
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v27)+92))
	if v773 < int32(0) {
		goto L146
	} else {
		goto L147
	}
L143:
	;
	goto L144
L144:
	;
	v814 = F_XLogReadBufferForRedo(m, l0, int32(1), v27+int32(76))
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L10
	} else {
		goto L150
	}
L145:
	;
	v792 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v791)+16)))
	v793 = v792 + v791
	v794 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v793)+12)))
	v796 = v794 | int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v793)+12)) = uint16(v796)
	*(*int64)(unsafe.Add(mBase, uint32(v791)+24)) = v772
	v799 = int32(32)
	*(*uint16)(unsafe.Add(mBase, uint32(v791)+12)) = uint16(v799)
	*(*uint32)(unsafe.Add(mBase, uint32(v791)+4)) = uint32(v764)
	v803 = int64(base.Ui64(v764) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v791))) = uint32(v803)
	v805 = *(*int32)(unsafe.Add(mBase, uint32(v27)+92))
	F_MarkBufferDirty(m, v805)
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L10
	} else {
		goto L149
	}
L146:
	;
	v777 = *(*int32)(unsafe.Add(mBase, _c_F_gist_redo[3]))
	v783 = *(*int32)(unsafe.Add(mBase, uint32(v777+(v773^int32(-1))<<(uint(int32(2))%32))))
	v791 = v783
	goto L145
L147:
	;
	goto L148
L148:
	;
	v785 = *(*int32)(unsafe.Add(mBase, _c_F_gist_redo[4]))
	v791 = v785 + v773<<(uint(int32(13))%32) + int32(-8192)
	goto L145
L149:
	;
	goto L144
L150:
	;
	if v814 == int32(0) {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v27)+76))
	if v818 < int32(0) {
		goto L155
	} else {
		goto L156
	}
L152:
	;
	goto L153
L153:
	;
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v27)+76))
	if v847 != 0 {
		goto L160
	} else {
		goto L161
	}
L154:
	;
	v837 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v763)+8)))
	F_PageIndexTupleDelete(m, v836, v837)
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L10
	} else {
		goto L158
	}
L155:
	;
	v822 = *(*int32)(unsafe.Add(mBase, _c_F_gist_redo[3]))
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v822+(v818^int32(-1))<<(uint(int32(2))%32))))
	v836 = v828
	goto L154
L156:
	;
	goto L157
L157:
	;
	v830 = *(*int32)(unsafe.Add(mBase, _c_F_gist_redo[4]))
	v836 = v830 + v818<<(uint(int32(13))%32) + int32(-8192)
	goto L154
L158:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v836))) = base.I64_rotr(v764, int64(32))
	v843 = *(*int32)(unsafe.Add(mBase, uint32(v27)+76))
	F_MarkBufferDirty(m, v843)
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L10
	} else {
		goto L159
	}
L159:
	;
	goto L153
L160:
	;
	F_UnlockReleaseBuffer(m, v847)
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L10
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	v850 = *(*int32)(unsafe.Add(mBase, uint32(v27)+92))
	if v850 == int32(0) {
		goto L4
	} else {
		goto L164
	}
L163:
	;
	goto L162
L164:
	;
	F_UnlockReleaseBuffer(m, v850)
	mBase = m.M
	v854 = m.ExcPending
	if v854 != 0 {
		goto L10
	} else {
		goto L165
	}
L165:
	;
	goto L4
L166:
	;
	m.G0 = v27 + int32(96)
	return
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = v109
	F_errmsg_internal(m, int32(_a_F_gist_redo_8), v27+int32(32))
	mBase = m.M
	v897 = m.ExcPending
	if v897 != 0 {
		goto L10
	} else {
		goto L168
	}
L168:
	;
	F_errfinish(m, int32(_a_F_gist_redo_9), int32(103), int32(_a_F_gist_redo_10))
	mBase = m.M
	v902 = m.ExcPending
	if v902 != 0 {
		goto L10
	} else {
		goto L169
	}
L169:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v173
	F_errmsg_internal(m, int32(_a_F_gist_redo_8), v27+int32(16))
	mBase = m.M
	v912 = m.ExcPending
	if v912 != 0 {
		goto L10
	} else {
		goto L171
	}
L171:
	;
	F_errfinish(m, int32(_a_F_gist_redo_9), int32(139), int32(_a_F_gist_redo_10))
	mBase = m.M
	v917 = m.ExcPending
	if v917 != 0 {
		goto L10
	} else {
		goto L172
	}
L172:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v38
	F_errmsg_internal(m, int32(_a_F_gist_redo_11), v27)
	mBase = m.M
	v925 = m.ExcPending
	if v925 != 0 {
		goto L10
	} else {
		goto L174
	}
L174:
	;
	F_errfinish(m, int32(_a_F_gist_redo_9), int32(430), int32(_a_F_gist_redo_12))
	mBase = m.M
	v930 = m.ExcPending
	if v930 != 0 {
		goto L10
	} else {
		goto L175
	}
L175:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
