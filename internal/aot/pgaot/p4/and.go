package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_FindAndDropRelationBuffers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
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
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	if base.Ui32(l3) < base.Ui32(l2) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v24 = v16
	v25 = l3
	goto L4
L2:
	;
	goto L3
L3:
	;
	m.G0 = v13 + int32(48)
	return
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v24
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v28
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v30
	v36 = F_BufTableHashCode(m, v13+int32(4))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
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
	v39 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	v46 = v39 + v36&int32(127)<<(uint(int32(7))%32) + int32(6912)
	v48 = F_LWLockAcquire(m, v46, int32(1))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v52 = F_BufTableLookup(m, v13+int32(4), v36)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	F_LWLockRelease(m, v46)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	if v52 < int32(0) {
		v157 = v24
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v161 = v25 + int32(1)
	if v161 != l2 {
		v24 = v157
		v25 = v161
		goto L4
	} else {
		goto L38
	}
L12:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+44)) = int32(228116)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+40)) = int32(6259)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = int32(491975)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = int64(0)
	v72 = v59 + v52<<(uint(int32(6))%32)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+24))
	v74 = int32(4194304)
	*(*int32)(unsafe.Add(mBase, uint32(v72)+24)) = v73 | v74
	if v73&v74 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	goto L16
L14:
	;
	v105 = v73
	goto L15
L15:
	;
	v112 = int32(4102444)
	v113 = *(*int32)(unsafe.Add(mBase, _consts[414]))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(24))+8))
	if v115 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L16:
	;
	F_perform_spin_delay(m, v13+int32(24))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L6
	} else {
		goto L18
	}
L17:
	;
	v105 = v93
	goto L15
L18:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v72)+24))
	v94 = int32(4194304)
	*(*int32)(unsafe.Add(mBase, uint32(v72)+24)) = v93 | v94
	if v93&v94 != 0 {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v132 != v133 {
		goto L31
	} else {
		goto L32
	}
L21:
	;
	goto L20
L22:
	;
	*(*int32)(unsafe.Add(mBase, _consts[414])) = v130
	goto L21
L23:
	;
	if int32(999) < v113 {
		goto L21
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	if v113 < int32(11) {
		goto L21
	} else {
		goto L30
	}
L26:
	;
	v120 = int32(900)
	if v120 <= v113 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v123 = v120
	goto L29
L28:
	;
	v123 = v113
	goto L29
L29:
	;
	v130 = v123 + int32(100)
	goto L22
L30:
	;
	v130 = v113 - int32(1)
	goto L22
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72)+24)) = v105 & int32(-4194305)
	v157 = v133
	goto L11
L32:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v135 != v136 {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v72)+8))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v138 != v139 {
		goto L31
	} else {
		goto L34
	}
L34:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v72)+12))
	if v141 != l1 {
		goto L31
	} else {
		goto L35
	}
L35:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v72)+16))
	if base.Ui32(v143) < base.Ui32(l3) {
		goto L31
	} else {
		goto L36
	}
L36:
	;
	F_InvalidateBuffer(m, v72)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L6
	} else {
		goto L37
	}
L37:
	;
	v157 = v132
	goto L11
L38:
	;
	goto L5
}
func F_ReleaseAndReadBuffer(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
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
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	if l0 == int32(0) {
		v78 = int32(0)
		v81 = F_ReadBufferExtended(m, l1, v78, l2, v78, v78)
		mBase = m.M
		v82 = m.ExcPending
		if v82 != 0 {
			return int32(0)
		} else {
			v83 = v81
			return v83
		}
	} else {
		if l0 < int32(0) {
			v11 = *(*int32)(unsafe.Add(mBase, _consts[3]))
			v16 = v11 + (l0^int32(-1))<<(uint(int32(6))%32)
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
			if v17 != l2 {
				F_UnpinLocalBuffer(m, l0)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					v78 = int32(0)
					v81 = F_ReadBufferExtended(m, l1, v78, l2, v78, v78)
					mBase = m.M
					v82 = m.ExcPending
					if v82 != 0 {
						return int32(0)
					} else {
						v83 = v81
						return v83
					}
				}
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				if v19 != v20 {
					F_UnpinLocalBuffer(m, l0)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						v78 = int32(0)
						v81 = F_ReadBufferExtended(m, l1, v78, l2, v78, v78)
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return int32(0)
						} else {
							v83 = v81
							return v83
						}
					}
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					if v22 != v23 {
						F_UnpinLocalBuffer(m, l0)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							v78 = int32(0)
							v81 = F_ReadBufferExtended(m, l1, v78, l2, v78, v78)
							mBase = m.M
							v82 = m.ExcPending
							if v82 != 0 {
								return int32(0)
							} else {
								v83 = v81
								return v83
							}
						}
					} else {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
						v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						if v25 != v26 {
							F_UnpinLocalBuffer(m, l0)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								v78 = int32(0)
								v81 = F_ReadBufferExtended(m, l1, v78, l2, v78, v78)
								mBase = m.M
								v82 = m.ExcPending
								if v82 != 0 {
									return int32(0)
								} else {
									v83 = v81
									return v83
								}
							}
						} else {
							v28 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
							if v28 == int32(0) {
								v83 = l0
								return v83
							} else {
								F_UnpinLocalBuffer(m, l0)
								mBase = m.M
								v34 = m.ExcPending
								if v34 != 0 {
									return int32(0)
								} else {
									v78 = int32(0)
									v81 = F_ReadBufferExtended(m, l1, v78, l2, v78, v78)
									mBase = m.M
									v82 = m.ExcPending
									if v82 != 0 {
										return int32(0)
									} else {
										v83 = v81
										return v83
									}
								}
							}
						}
					}
				}
			}
		} else {
			v36 = *(*int32)(unsafe.Add(mBase, _consts[4]))
			v39 = v36 + l0<<(uint(int32(6))%32)
			v41 = v39 + int32(-64)
			v44 = *(*int32)(unsafe.Add(mBase, uint32(v39-int32(48))))
			if v44 != l2 {
				v65 = *(*int32)(unsafe.Add(mBase, _consts[173]))
				v68 = *(*int32)(unsafe.Add(mBase, uint32(v39-int32(44))))
				F_ResourceOwnerForget(m, v65, v68+int32(1), int32(1609600))
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
					return int32(0)
				} else {
					F_UnpinBufferNoOwner(m, v41)
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return int32(0)
					} else {
						v78 = int32(0)
						v81 = F_ReadBufferExtended(m, l1, v78, l2, v78, v78)
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return int32(0)
						} else {
							v83 = v81
							return v83
						}
					}
				}
			} else {
				v46 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
				v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				if v46 != v47 {
					v65 = *(*int32)(unsafe.Add(mBase, _consts[173]))
					v68 = *(*int32)(unsafe.Add(mBase, uint32(v39-int32(44))))
					F_ResourceOwnerForget(m, v65, v68+int32(1), int32(1609600))
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int32(0)
					} else {
						F_UnpinBufferNoOwner(m, v41)
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return int32(0)
						} else {
							v78 = int32(0)
							v81 = F_ReadBufferExtended(m, l1, v78, l2, v78, v78)
							mBase = m.M
							v82 = m.ExcPending
							if v82 != 0 {
								return int32(0)
							} else {
								v83 = v81
								return v83
							}
						}
					}
				} else {
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v39-int32(60))))
					v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					if v51 != v52 {
						v65 = *(*int32)(unsafe.Add(mBase, _consts[173]))
						v68 = *(*int32)(unsafe.Add(mBase, uint32(v39-int32(44))))
						F_ResourceOwnerForget(m, v65, v68+int32(1), int32(1609600))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return int32(0)
						} else {
							F_UnpinBufferNoOwner(m, v41)
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return int32(0)
							} else {
								v78 = int32(0)
								v81 = F_ReadBufferExtended(m, l1, v78, l2, v78, v78)
								mBase = m.M
								v82 = m.ExcPending
								if v82 != 0 {
									return int32(0)
								} else {
									v83 = v81
									return v83
								}
							}
						}
					} else {
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v39-int32(56))))
						v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						if v56 != v57 {
							v65 = *(*int32)(unsafe.Add(mBase, _consts[173]))
							v68 = *(*int32)(unsafe.Add(mBase, uint32(v39-int32(44))))
							F_ResourceOwnerForget(m, v65, v68+int32(1), int32(1609600))
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return int32(0)
							} else {
								F_UnpinBufferNoOwner(m, v41)
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
									return int32(0)
								} else {
									v78 = int32(0)
									v81 = F_ReadBufferExtended(m, l1, v78, l2, v78, v78)
									mBase = m.M
									v82 = m.ExcPending
									if v82 != 0 {
										return int32(0)
									} else {
										v83 = v81
										return v83
									}
								}
							}
						} else {
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v39-int32(52))))
							if v61 == int32(0) {
								v83 = l0
								return v83
							} else {
								v65 = *(*int32)(unsafe.Add(mBase, _consts[173]))
								v68 = *(*int32)(unsafe.Add(mBase, uint32(v39-int32(44))))
								F_ResourceOwnerForget(m, v65, v68+int32(1), int32(1609600))
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return int32(0)
								} else {
									F_UnpinBufferNoOwner(m, v41)
									mBase = m.M
									v75 = m.ExcPending
									if v75 != 0 {
										return int32(0)
									} else {
										v78 = int32(0)
										v81 = F_ReadBufferExtended(m, l1, v78, l2, v78, v78)
										mBase = m.M
										v82 = m.ExcPending
										if v82 != 0 {
											return int32(0)
										} else {
											v83 = v81
											return v83
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
func F_check_and_set_sync_info(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	v5 = *(*int32)(unsafe.Add(mBase, _consts[830]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = int32(1)
	if v6 != 0 {
		v10 = *(*int32)(unsafe.Add(mBase, _consts[830]))
		F_s_lock(m, v10+int32(16), int32(496714), int32(1295), int32(240715))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, _consts[830]))
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+4)))
			if v20 != int32(1) {
				v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+5)))
				if v23 == int32(1) {
					v52 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v52
					F_errstart_cold(m, int32(21), v52)
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return
					} else {
						F_errcode(m, int32(325))
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return
						} else {
							F_errmsg(m, int32(18508), int32(0))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return
							} else {
								F_errfinish(m, int32(496714), int32(1317), int32(240715))
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v19))) = l0
					v29 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v19)+5)) = uint8(v29)
					*(*uint8)(unsafe.Add(mBase, _consts[831])) = uint8(v29)
					return
				}
			} else {
				v34 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v34
				F_errstart_cold(m, int32(21), v34)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return
				} else {
					F_errcode(m, int32(325))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						F_errmsg(m, int32(332494), int32(0))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return
						} else {
							F_errfinish(m, int32(496714), int32(1309), int32(240715))
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
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
		v19 = *(*int32)(unsafe.Add(mBase, _consts[830]))
		v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+4)))
		if v20 != int32(1) {
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+5)))
			if v23 == int32(1) {
				v52 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v52
				F_errstart_cold(m, int32(21), v52)
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return
				} else {
					F_errcode(m, int32(325))
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return
					} else {
						F_errmsg(m, int32(18508), int32(0))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return
						} else {
							F_errfinish(m, int32(496714), int32(1317), int32(240715))
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v19))) = l0
				v29 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v19)+5)) = uint8(v29)
				*(*uint8)(unsafe.Add(mBase, _consts[831])) = uint8(v29)
				return
			}
		} else {
			v34 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v34
			F_errstart_cold(m, int32(21), v34)
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return
			} else {
				F_errcode(m, int32(325))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return
				} else {
					F_errmsg(m, int32(332494), int32(0))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return
					} else {
						F_errfinish(m, int32(496714), int32(1309), int32(240715))
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
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
func F_parse_and_validate_value(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
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
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
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
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v168 int32
	_ = v168
	var v169 float64
	_ = v169
	var v170 float64
	_ = v170
	var v174 float64
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v189 float64
	_ = v189
	var v190 int32
	_ = v190
	var v191 float64
	_ = v191
	var v192 float64
	_ = v192
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v258 int32
	_ = v258
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v284 int32
	_ = v284
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v345 int32
	_ = v345
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v388 int32
	_ = v388
	v14 = m.G0
	v16 = v14 - int32(224)
	m.G0 = v16
	v18 = int32(1)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	switch v19 {
	case 0:
		goto L7
	case 1:
		goto L6
	case 2:
		goto L5
	case 3:
		goto L4
	case 4:
		goto L3
	default:
		v388 = v18
		goto L1
	}
L1:
	;
	m.G0 = v16 + int32(224)
	return v388
L2:
	;
	v388 = int32(0)
	goto L1
L3:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v248 == int32(0) {
		goto L93
	} else {
		goto L94
	}
L4:
	;
	v228 = F_guc_strdup(m, l3, l1)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L12
	} else {
		goto L80
	}
L5:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v139 = F_parse_real(m, l1, l4, v136, v16+int32(220))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L12
	} else {
		goto L52
	}
L6:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v51 = F_parse_int(m, l1, l4, v48, v16+int32(220))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L12
	} else {
		goto L22
	}
L7:
	;
	v20 = F_strlen(m, l1)
	mBase = m.M
	v21 = F_parse_bool_with_len(m, l1, v20, l4)
	mBase = m.M
	goto L8
L8:
	;
	if v21 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v25 = F_errstart(m, l3, int32(0))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v44 = F_call_bool_check_hook(m, l0, l4, l5, l2, l3)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L12
	} else {
		goto L18
	}
L12:
	;
	return int32(0)
L13:
	;
	if v25 == int32(0) {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v34
	F_errmsg(m, int32(343725), v16)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L12
	} else {
		goto L16
	}
L16:
	;
	F_errfinish(m, int32(496556), int32(3147), int32(343162))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	goto L2
L18:
	;
	if v44 == int32(0) {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	v388 = v18
	goto L1
L20:
	;
	v132 = F_call_int_check_hook(m, l0, l4, l5, l2, l3)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L12
	} else {
		goto L48
	}
L21:
	;
	F_errfinish(m, int32(496556), v122, int32(343162))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L12
	} else {
		goto L47
	}
L22:
	;
	if v51 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v56 = F_errstart(m, l3, int32(0))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L12
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v82 <= v81 {
		goto L32
	} else {
		goto L33
	}
L26:
	;
	if v56 == int32(0) {
		goto L2
	} else {
		goto L27
	}
L27:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L12
	} else {
		goto L28
	}
L28:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+84)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v16)+80)) = v63
	F_errmsg(m, int32(709096), v16+int32(80))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L12
	} else {
		goto L29
	}
L29:
	;
	v71 = int32(3168)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v16)+220))
	if v72 == int32(0) {
		v122 = v71
		goto L21
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v72
	F_errhint(m, int32(205224), v16-int32(-64))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L12
	} else {
		goto L31
	}
L31:
	;
	v122 = v71
	goto L21
L32:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v81 <= v84 {
		goto L20
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v87 = F_get_config_unit_name(m, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L12
	} else {
		goto L36
	}
L35:
	;
	goto L34
L36:
	;
	v90 = F_errstart(m, l3, int32(0))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L12
	} else {
		goto L37
	}
L37:
	;
	if v90 == int32(0) {
		goto L2
	} else {
		goto L38
	}
L38:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L12
	} else {
		goto L39
	}
L39:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v87 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v102 = v87
	goto L42
L41:
	;
	v102 = int32(738681)
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v102
	if v87 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v106 = int32(727620)
	goto L45
L44:
	;
	v106 = int32(738681)
	goto L45
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v97
	F_errmsg(m, int32(656099), v16+int32(16))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L12
	} else {
		goto L46
	}
L46:
	;
	v122 = int32(3188)
	goto L21
L47:
	;
	goto L2
L48:
	;
	if v132 == int32(0) {
		goto L2
	} else {
		goto L49
	}
L49:
	;
	v388 = v18
	goto L1
L50:
	;
	v224 = F_call_real_check_hook(m, l0, l4, l5, l2, l3)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L12
	} else {
		goto L78
	}
L51:
	;
	F_errfinish(m, int32(496556), v214, int32(343162))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L12
	} else {
		goto L77
	}
L52:
	;
	if v139 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v144 = F_errstart(m, l3, int32(0))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L12
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v169 = *(*float64)(unsafe.Add(mBase, uint32(l4)))
	v170 = *(*float64)(unsafe.Add(mBase, uint32(l0)+104))
	if base.F64_lt(v169, v170) == int32(0) {
		goto L62
	} else {
		goto L63
	}
L56:
	;
	if v144 == int32(0) {
		goto L2
	} else {
		goto L57
	}
L57:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L12
	} else {
		goto L58
	}
L58:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+180)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v16)+176)) = v151
	F_errmsg(m, int32(709096), v16+int32(176))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L12
	} else {
		goto L59
	}
L59:
	;
	v159 = int32(3209)
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v16)+220))
	if v160 == int32(0) {
		v214 = v159
		goto L51
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+160)) = v160
	F_errhint(m, int32(205224), v16+int32(160))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L12
	} else {
		goto L61
	}
L61:
	;
	v214 = v159
	goto L51
L62:
	;
	v174 = *(*float64)(unsafe.Add(mBase, uint32(l0)+112))
	if base.F64_gt(v169, v174) == int32(0) {
		goto L50
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v179 = F_get_config_unit_name(m, v178)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L12
	} else {
		goto L66
	}
L65:
	;
	goto L64
L66:
	;
	v182 = F_errstart(m, l3, int32(0))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L12
	} else {
		goto L67
	}
L67:
	;
	if v182 == int32(0) {
		goto L2
	} else {
		goto L68
	}
L68:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L12
	} else {
		goto L69
	}
L69:
	;
	v189 = *(*float64)(unsafe.Add(mBase, uint32(l4)))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v191 = *(*float64)(unsafe.Add(mBase, uint32(l0)+104))
	v192 = *(*float64)(unsafe.Add(mBase, uint32(l0)+112))
	if v179 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v194 = v179
	goto L72
L71:
	;
	v194 = int32(738681)
	goto L72
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+148)) = v194
	if v179 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v198 = int32(727620)
	goto L75
L74:
	;
	v198 = int32(738681)
	goto L75
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+144)) = v198
	*(*float64)(unsafe.Add(mBase, uint32(v16)+136)) = v192
	*(*int32)(unsafe.Add(mBase, uint32(v16)+132)) = v194
	*(*int32)(unsafe.Add(mBase, uint32(v16)+128)) = v198
	*(*float64)(unsafe.Add(mBase, uint32(v16)+120)) = v191
	*(*int32)(unsafe.Add(mBase, uint32(v16)+112)) = v190
	*(*int32)(unsafe.Add(mBase, uint32(v16)+108)) = v194
	*(*int32)(unsafe.Add(mBase, uint32(v16)+104)) = v198
	*(*float64)(unsafe.Add(mBase, uint32(v16)+96)) = v189
	F_errmsg(m, int32(656027), v16+int32(96))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L12
	} else {
		goto L76
	}
L76:
	;
	v214 = int32(3229)
	goto L51
L77:
	;
	goto L2
L78:
	;
	if v224 == int32(0) {
		goto L2
	} else {
		goto L79
	}
L79:
	;
	v388 = v18
	goto L1
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v228
	if v228 == int32(0) {
		goto L2
	} else {
		goto L81
	}
L81:
	;
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)))
	if v233&int32(8) != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v236 = F_strlen(m, v228)
	mBase = m.M
	F_truncate_identifier(m, v228, v236, int32(1))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L12
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v240 = F_call_string_check_hook(m, l0, l4, l5, l2, l3)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L12
	} else {
		goto L86
	}
L85:
	;
	goto L84
L86:
	;
	if v240 != 0 {
		v388 = v18
		goto L1
	} else {
		goto L87
	}
L87:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v242 != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	F_pfree(m, v242)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L12
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v245 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v245
	v388 = v245
	goto L1
L91:
	;
	goto L90
L92:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v258)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v362
	v364 = F_call_enum_check_hook(m, l0, l4, l5, l2, l3)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L12
	} else {
		goto L127
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(0)
	v330 = F_config_enum_get_options(m, l0, int32(726794), int32(651433), int32(727439))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L12
	} else {
		goto L113
	}
L94:
	;
	v258 = v248
	goto L95
L95:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
	if v264 == int32(0) {
		goto L93
	} else {
		goto L97
	}
L96:
	;
	goto L93
L97:
	;
	v269 = l1
	v270 = v264
	goto L99
L98:
	;
	if v307 == int32(0) {
		goto L92
	} else {
		goto L111
	}
L99:
	;
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269))))
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270))))
	if v273 == v274 {
		v296 = v273
		goto L101
	} else {
		goto L102
	}
L100:
	;
	v307 = int32(0)
	goto L98
L101:
	;
	v298 = int32(1)
	if v296 != 0 {
		v269 = v269 + v298
		v270 = v270 + v298
		goto L99
	} else {
		goto L110
	}
L102:
	;
	if base.Ui32((v273-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v284 = v273 | int32(32)
	goto L105
L104:
	;
	v284 = v273
	goto L105
L105:
	;
	if base.Ui32((v274-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v293 = v274 | int32(32)
	goto L108
L107:
	;
	v293 = v274
	goto L108
L108:
	;
	if v284 == v293 {
		v296 = v284
		goto L101
	} else {
		goto L109
	}
L109:
	;
	v307 = v284 - v293
	goto L98
L110:
	;
	goto L100
L111:
	;
	v311 = v258 + int32(12)
	if v311 != 0 {
		v258 = v311
		goto L95
	} else {
		goto L112
	}
L112:
	;
	goto L96
L113:
	;
	v333 = F_errstart(m, l3, int32(0))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L12
	} else {
		goto L114
	}
L114:
	;
	if v333 != 0 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L12
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	if v330 == int32(0) {
		goto L2
	} else {
		goto L125
	}
L118:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+212)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v16)+208)) = v338
	F_errmsg(m, int32(709096), v16+int32(208))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L12
	} else {
		goto L119
	}
L119:
	;
	if v330 != 0 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+192)) = v330
	F_errhint(m, int32(205224), v16+int32(192))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L12
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	F_errfinish(m, int32(496556), int32(3284), int32(343162))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L12
	} else {
		goto L124
	}
L123:
	;
	goto L122
L124:
	;
	goto L117
L125:
	;
	F_pfree(m, v330)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L12
	} else {
		goto L126
	}
L126:
	;
	goto L2
L127:
	;
	if v364 != 0 {
		v388 = v18
		goto L1
	} else {
		goto L128
	}
L128:
	;
	goto L2
}
func F_update_and_persist_local_synced_slot(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v46 int64
	_ = v46
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _consts[832]))
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+47)) = uint8(v3)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+46)) = uint8(v3)
	v21 = F_update_local_synced_slot(m, l0, l1, v9+int32(47), v9+int32(46))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return int32(0)
	} else {
		v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+46)))
		if v25 != 0 {
			v76 = v3
			m.G0 = v9 + int32(48)
			return v76
		} else {
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+47)))
			if v27 == int32(0) {
				v32 = F_errstart(m, int32(15), int32(0))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					if v32 == int32(0) {
						v76 = v3
						m.G0 = v9 + int32(48)
						return v76
					} else {
						v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v36
						F_errmsg(m, int32(679705), v9+int32(32))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							v43 = *(*int64)(unsafe.Add(mBase, uint32(v12)+104))
							*(*uint32)(unsafe.Add(mBase, uint32(v9)+20)) = uint32(v43)
							v46 = int64(base.Ui64(v43) >> (uint(int64(32)) % 64))
							*(*uint32)(unsafe.Add(mBase, uint32(v9)+16)) = uint32(v46)
							F_errdetail(m, int32(638410), v9+int32(16))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								v71 = int32(598)
								F_errfinish(m, int32(496714), v71, int32(84712))
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return int32(0)
								} else {
									v76 = v27
									m.G0 = v9 + int32(48)
									return v76
								}
							}
						}
					}
				}
			} else {
				F_ReplicationSlotPersist(m)
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return int32(0)
				} else {
					v59 = F_errstart(m, int32(15), int32(0))
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return int32(0)
					} else {
						if v59 == int32(0) {
							v76 = int32(1)
							m.G0 = v9 + int32(48)
							return v76
						} else {
							v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v63
							F_errmsg(m, int32(30818), v9)
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								v71 = int32(607)
								F_errfinish(m, int32(496714), v71, int32(84712))
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return int32(0)
								} else {
									v76 = v27
									m.G0 = v9 + int32(48)
									return v76
								}
							}
						}
					}
				}
			}
		}
	}
}
