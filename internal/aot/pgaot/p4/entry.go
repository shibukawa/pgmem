package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_cmpEntryAccumulator(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
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
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v5 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+22)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v7 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+20)))
	v8 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+22)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v10 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1)+20)))
	v11 = F_ginCompareAttEntries(m, v4, v5, v6, v7, v8, v9, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		return v11
	}
}
func F_entryExecPlaceToPage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	v5 = l4
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	if l1 < int32(0) {
		v16 = *(*int32)(unsafe.Add(mBase, _consts[1]))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v16+(l1^int32(-1))<<(uint(int32(2))%32))))
		v30 = v22
	} else {
		v24 = *(*int32)(unsafe.Add(mBase, _consts[2]))
		v30 = v24 + l1<<(uint(int32(13))%32) + int32(-8192)
	}
	v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+8)))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)))
	if v32 == int32(1) {
		F_PageIndexTupleDelete(m, v30, v31)
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return
		} else {
			if v5 == int32(-1) {
			} else {
				v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+16)))
				v41 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30+v39)+6)))
				if v41&int32(2) != 0 {
				} else {
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v31<<(uint(int32(2))%32)+v30)+20))
					v50 = v30 + v47&int32(32767)
					v51 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(v50)+4)) = uint16(v51)
					*(*uint16)(unsafe.Add(mBase, uint32(v50)+2)) = uint16(v5)
					v55 = int32(base.Ui32(v5) >> (uint(int32(16)) % 32))
					*(*uint16)(unsafe.Add(mBase, uint32(v50))) = uint16(v55)
				}
			}
			v58 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
			v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+6)))
			v63 = F_PageAddItemExtended(m, v30, v58, v59&int32(8191), v31, int32(0))
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return
			} else {
				if v63 == v31 {
					F_MarkBufferDirty(m, l1)
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return
					} else {
						v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+48))
						v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+118)))
						if v70 != int32(112) {
							m.G0 = v11 + int32(16)
							return
						} else {
							v74 = *(*int32)(unsafe.Add(mBase, _consts[27]))
							if v74 <= int32(0) {
								v77 = *(*int32)(unsafe.Add(mBase, uint32(v68)+32))
								if v77 != 0 {
									m.G0 = v11 + int32(16)
									return
								} else {
									v78 = *(*int32)(unsafe.Add(mBase, uint32(v68)+40))
									if v78 != 0 {
										m.G0 = v11 + int32(16)
										return
									} else {
										v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+53)))
										if v79 != 0 {
											m.G0 = v11 + int32(16)
											return
										} else {
											v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)))
											*(*uint16)(unsafe.Add(mBase, _consts[40])) = uint16(v31)
											*(*uint8)(unsafe.Add(mBase, _consts[41])) = uint8(v80)
											F_XLogRegisterBuffer(m, int32(0), l1, int32(8))
											mBase = m.M
											v88 = m.ExcPending
											if v88 != 0 {
												return
											} else {
												F_XLogRegisterBufData(m, int32(0), int32(4409728), int32(4))
												mBase = m.M
												v93 = m.ExcPending
												if v93 != 0 {
													return
												} else {
													v95 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
													v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v95)+6)))
													F_XLogRegisterBufData(m, int32(0), v95, v96&int32(8191))
													mBase = m.M
													v100 = m.ExcPending
													if v100 != 0 {
														return
													} else {
														m.G0 = v11 + int32(16)
														return
													}
												}
											}
										}
									}
								}
							} else {
								v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+53)))
								if v79 != 0 {
									m.G0 = v11 + int32(16)
									return
								} else {
									v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)))
									*(*uint16)(unsafe.Add(mBase, _consts[40])) = uint16(v31)
									*(*uint8)(unsafe.Add(mBase, _consts[41])) = uint8(v80)
									F_XLogRegisterBuffer(m, int32(0), l1, int32(8))
									mBase = m.M
									v88 = m.ExcPending
									if v88 != 0 {
										return
									} else {
										F_XLogRegisterBufData(m, int32(0), int32(4409728), int32(4))
										mBase = m.M
										v93 = m.ExcPending
										if v93 != 0 {
											return
										} else {
											v95 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
											v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v95)+6)))
											F_XLogRegisterBufData(m, int32(0), v95, v96&int32(8191))
											mBase = m.M
											v100 = m.ExcPending
											if v100 != 0 {
												return
											} else {
												m.G0 = v11 + int32(16)
												return
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
					v108 = m.ExcPending
					if v108 != 0 {
						return
					} else {
						v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v11))) = v110 + int32(4)
						F_errmsg_internal(m, int32(711804), v11)
						mBase = m.M
						v116 = m.ExcPending
						if v116 != 0 {
							return
						} else {
							F_errfinish(m, int32(499132), int32(571), int32(408830))
							mBase = m.M
							v121 = m.ExcPending
							if v121 != 0 {
								return
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
	} else {
		if v5 == int32(-1) {
		} else {
			v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+16)))
			v41 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30+v39)+6)))
			if v41&int32(2) != 0 {
			} else {
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v31<<(uint(int32(2))%32)+v30)+20))
				v50 = v30 + v47&int32(32767)
				v51 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v50)+4)) = uint16(v51)
				*(*uint16)(unsafe.Add(mBase, uint32(v50)+2)) = uint16(v5)
				v55 = int32(base.Ui32(v5) >> (uint(int32(16)) % 32))
				*(*uint16)(unsafe.Add(mBase, uint32(v50))) = uint16(v55)
			}
		}
		v58 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
		v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+6)))
		v63 = F_PageAddItemExtended(m, v30, v58, v59&int32(8191), v31, int32(0))
		mBase = m.M
		v64 = m.ExcPending
		if v64 != 0 {
			return
		} else {
			if v63 == v31 {
				F_MarkBufferDirty(m, l1)
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return
				} else {
					v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+48))
					v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+118)))
					if v70 != int32(112) {
						m.G0 = v11 + int32(16)
						return
					} else {
						v74 = *(*int32)(unsafe.Add(mBase, _consts[27]))
						if v74 <= int32(0) {
							v77 = *(*int32)(unsafe.Add(mBase, uint32(v68)+32))
							if v77 != 0 {
								m.G0 = v11 + int32(16)
								return
							} else {
								v78 = *(*int32)(unsafe.Add(mBase, uint32(v68)+40))
								if v78 != 0 {
									m.G0 = v11 + int32(16)
									return
								} else {
									v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+53)))
									if v79 != 0 {
										m.G0 = v11 + int32(16)
										return
									} else {
										v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)))
										*(*uint16)(unsafe.Add(mBase, _consts[40])) = uint16(v31)
										*(*uint8)(unsafe.Add(mBase, _consts[41])) = uint8(v80)
										F_XLogRegisterBuffer(m, int32(0), l1, int32(8))
										mBase = m.M
										v88 = m.ExcPending
										if v88 != 0 {
											return
										} else {
											F_XLogRegisterBufData(m, int32(0), int32(4409728), int32(4))
											mBase = m.M
											v93 = m.ExcPending
											if v93 != 0 {
												return
											} else {
												v95 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
												v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v95)+6)))
												F_XLogRegisterBufData(m, int32(0), v95, v96&int32(8191))
												mBase = m.M
												v100 = m.ExcPending
												if v100 != 0 {
													return
												} else {
													m.G0 = v11 + int32(16)
													return
												}
											}
										}
									}
								}
							}
						} else {
							v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+53)))
							if v79 != 0 {
								m.G0 = v11 + int32(16)
								return
							} else {
								v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)))
								*(*uint16)(unsafe.Add(mBase, _consts[40])) = uint16(v31)
								*(*uint8)(unsafe.Add(mBase, _consts[41])) = uint8(v80)
								F_XLogRegisterBuffer(m, int32(0), l1, int32(8))
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
									return
								} else {
									F_XLogRegisterBufData(m, int32(0), int32(4409728), int32(4))
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
										return
									} else {
										v95 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
										v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v95)+6)))
										F_XLogRegisterBufData(m, int32(0), v95, v96&int32(8191))
										mBase = m.M
										v100 = m.ExcPending
										if v100 != 0 {
											return
										} else {
											m.G0 = v11 + int32(16)
											return
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
				v108 = m.ExcPending
				if v108 != 0 {
					return
				} else {
					v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+48))
					*(*int32)(unsafe.Add(mBase, uint32(v11))) = v110 + int32(4)
					F_errmsg_internal(m, int32(711804), v11)
					mBase = m.M
					v116 = m.ExcPending
					if v116 != 0 {
						return
					} else {
						F_errfinish(m, int32(499132), int32(571), int32(408830))
						mBase = m.M
						v121 = m.ExcPending
						if v121 != 0 {
							return
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
func F_entryLocateLeafEntry(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
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
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v130 int32
	_ = v130
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v15 < int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v34 = int32(1)
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
	if v35 == v34 {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v19+(v15^int32(-1))<<(uint(int32(2))%32))))
	v33 = v25
	goto L1
L3:
	;
	goto L4
L4:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v33 = v27 + v15<<(uint(int32(13))%32) + int32(-8192)
	goto L1
L5:
	;
	m.G0 = v13 + int32(16)
	return v130
L6:
	;
	v38 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)) = uint16(v38)
	v130 = v34
	goto L5
L7:
	;
	goto L8
L8:
	;
	v40 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v33)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v40) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v130 = int32(0)
	goto L5
L10:
	;
	v54 = int32(1)
	v59 = v54
	v61 = v46 + v54
	goto L15
L11:
	;
	v46 = int32(base.Ui32(v40+int32(262120)) >> (uint(int32(2)) % 32))
	if v46&int32(65535) != 0 {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v50 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)) = uint16(v50)
	goto L9
L14:
	;
	goto L13
L15:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v73 = int32(base.Ui32((v61-v59)&int32(65534))>>(uint(int32(1))%32)) + v59
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v73&int32(65535)<<(uint(int32(2))%32)+(v33+int32(24))-int32(4))))
	v84 = v33 + v81&int32(32767)
	v85 = F_gintuple_get_attrnum(m, v67, v84)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)) = uint16(v107)
	goto L9
L17:
	;
	return int32(0)
L18:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v92 = F_gintuple_get_key(m, v89, v84, v13+int32(15))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+54)))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v97 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+60)))
	v98 = int32(*(*int8)(unsafe.Add(mBase, uint32(v13)+15)))
	v99 = F_ginCompareAttEntries(m, v94, v95, v96, v97, v85, v92, v98)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	if v99 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)) = uint16(v73)
	v130 = int32(1)
	goto L5
L22:
	;
	goto L23
L23:
	;
	v106 = base.B2i32(int32(0) < v99)
	if int32(0) < v99 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v107 = v61
	goto L26
L25:
	;
	v107 = v73
	goto L26
L26:
	;
	if int32(0) < v99 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v112 = v73 + int32(1)
	goto L29
L28:
	;
	v112 = v59
	goto L29
L29:
	;
	if base.Ui32(v112&int32(65535)) < base.Ui32(v107&int32(65535)) {
		v59 = v112
		v61 = v107
		goto L15
	} else {
		goto L30
	}
L30:
	;
	goto L16
}
