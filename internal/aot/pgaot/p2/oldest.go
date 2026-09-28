package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetOldestActiveTransactionId(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestActiveTransactionId[0]))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestActiveTransactionId[1]))
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestActiveTransactionId[2]))
	v22 = F_LWLockAcquire(m, v18+int32(384), int32(1))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return int32(0)
	} else {
		v27 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestActiveTransactionId[3]))
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
		v30 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestActiveTransactionId[2]))
		F_LWLockRelease(m, v30+int32(384))
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return int32(0)
		} else {
			v36 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestActiveTransactionId[2]))
			v40 = F_LWLockAcquire(m, v36+int32(512), int32(1))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int32(0)
			} else {
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
				if int32(0) < v42 {
					v48 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestActiveTransactionId[4]))
					v50 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestActiveTransactionId[5]))
					v54 = int32(0)
					v55 = v28
					for {
						v64 = v54 << (uint(int32(2)) % 32)
						v66 = *(*int32)(unsafe.Add(mBase, uint32(v16+int32(36)+v64)))
						v68 = *(*int32)(unsafe.Add(mBase, uint32(v64+v14)))
						if base.Ui32(v68) < base.Ui32(int32(3)) {
							v90 = v55
						} else {
							v73 = v50 + v66*int32(768)
							if l0 != 0 {
								v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+336)))
								if v74&int32(5) == int32(0) {
									v90 = v55
								} else {
									if l1 == int32(0) {
										v81 = *(*int32)(unsafe.Add(mBase, uint32(v73)+20))
										if v81 != v48 {
											v90 = v55
										} else {
											if v68-v55 < int32(0) {
												v86 = v68
											} else {
												v86 = v55
											}
											if base.Ui32(int32(2)) < base.Ui32(v55) {
												v89 = v86
											} else {
												v89 = v55
											}
											v90 = v89
										}
									} else {
										if v68-v55 < int32(0) {
											v86 = v68
										} else {
											v86 = v55
										}
										if base.Ui32(int32(2)) < base.Ui32(v55) {
											v89 = v86
										} else {
											v89 = v55
										}
										v90 = v89
									}
								}
							} else {
								if l1 == int32(0) {
									v81 = *(*int32)(unsafe.Add(mBase, uint32(v73)+20))
									if v81 != v48 {
										v90 = v55
									} else {
										if v68-v55 < int32(0) {
											v86 = v68
										} else {
											v86 = v55
										}
										if base.Ui32(int32(2)) < base.Ui32(v55) {
											v89 = v86
										} else {
											v89 = v55
										}
										v90 = v89
									}
								} else {
									if v68-v55 < int32(0) {
										v86 = v68
									} else {
										v86 = v55
									}
									if base.Ui32(int32(2)) < base.Ui32(v55) {
										v89 = v86
									} else {
										v89 = v55
									}
									v90 = v89
								}
							}
						}
						v93 = v54 + int32(1)
						if v93 != v42 {
							v54 = v93
							v55 = v90
							continue
						} else {
							break
						}
						break
					}
					v98 = v90
				} else {
					v98 = v28
				}
				v107 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestActiveTransactionId[2]))
				F_LWLockRelease(m, v107+int32(512))
				mBase = m.M
				v111 = m.ExcPending
				if v111 != 0 {
					return int32(0)
				} else {
					return v98
				}
			}
		}
	}
}
func F_GetOldestNonRemovableTransactionId(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	v3 = m.G0
	v5 = v3 - int32(48)
	m.G0 = v5
	F_ComputeXidHorizons(m, v5+int32(8))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = int32(0)
		if l0 == v13 {
			v56 = v13
			v60 = v56
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+117)))
			if v18 != 0 {
				v56 = v13
				v60 = v56
			} else {
				v19 = F_RecoveryInProgress(m)
				mBase = m.M
				if v19 != 0 {
					v56 = v13
					v60 = v56
				} else {
					v20 = F_IsCatalogRelation(m, l0)
					mBase = m.M
					if v20 != 0 {
						v60 = int32(1)
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestNonRemovableTransactionId[0]))
						if v23 <= int32(1) {
							v27 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetOldestNonRemovableTransactionId[1])))
							if v27&int32(1) == int32(0) {
								v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
								if v52 != 0 {
									v56 = int32(3)
									v60 = v56
								} else {
									v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
									if v53 != 0 {
										v56 = int32(3)
										v60 = v56
									} else {
										v60 = int32(2)
									}
								}
							} else {
								v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
								v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+118)))
								if v33 != int32(112) {
									v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
									if v52 != 0 {
										v56 = int32(3)
										v60 = v56
									} else {
										v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
										if v53 != 0 {
											v56 = int32(3)
											v60 = v56
										} else {
											v60 = int32(2)
										}
									}
								} else {
									if v23 <= int32(0) {
										v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
										if v38 != 0 {
											v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
											if v52 != 0 {
												v56 = int32(3)
												v60 = v56
											} else {
												v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
												if v53 != 0 {
													v56 = int32(3)
													v60 = v56
												} else {
													v60 = int32(2)
												}
											}
										} else {
											v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
											if v39 != 0 {
												v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
												if v52 != 0 {
													v56 = int32(3)
													v60 = v56
												} else {
													v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
													if v53 != 0 {
														v56 = int32(3)
														v60 = v56
													} else {
														v60 = int32(2)
													}
												}
											} else {
												v40 = int32(1)
												v41 = F_IsCatalogRelation(m, l0)
												mBase = m.M
												if v41 != 0 {
													v56 = v40
													v60 = v56
												} else {
													v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
													if v42 == int32(0) {
														v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
														if v52 != 0 {
															v56 = int32(3)
															v60 = v56
														} else {
															v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
															if v53 != 0 {
																v56 = int32(3)
																v60 = v56
															} else {
																v60 = int32(2)
															}
														}
													} else {
														v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
														v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+119)))
														switch v46 - int32(109) {
														case 0, 5:
															v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+112)))
															if v49 != 0 {
																v56 = v40
																v60 = v56
															} else {
																v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
																if v52 != 0 {
																	v56 = int32(3)
																	v60 = v56
																} else {
																	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
																	if v53 != 0 {
																		v56 = int32(3)
																		v60 = v56
																	} else {
																		v60 = int32(2)
																	}
																}
															}
														default:
															v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
															if v52 != 0 {
																v56 = int32(3)
																v60 = v56
															} else {
																v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
																if v53 != 0 {
																	v56 = int32(3)
																	v60 = v56
																} else {
																	v60 = int32(2)
																}
															}
														}
													}
												}
											}
										}
									} else {
										v40 = int32(1)
										v41 = F_IsCatalogRelation(m, l0)
										mBase = m.M
										if v41 != 0 {
											v56 = v40
											v60 = v56
										} else {
											v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
											if v42 == int32(0) {
												v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
												if v52 != 0 {
													v56 = int32(3)
													v60 = v56
												} else {
													v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
													if v53 != 0 {
														v56 = int32(3)
														v60 = v56
													} else {
														v60 = int32(2)
													}
												}
											} else {
												v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
												v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+119)))
												switch v46 - int32(109) {
												case 0, 5:
													v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+112)))
													if v49 != 0 {
														v56 = v40
														v60 = v56
													} else {
														v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
														if v52 != 0 {
															v56 = int32(3)
															v60 = v56
														} else {
															v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
															if v53 != 0 {
																v56 = int32(3)
																v60 = v56
															} else {
																v60 = int32(2)
															}
														}
													}
												default:
													v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
													if v52 != 0 {
														v56 = int32(3)
														v60 = v56
													} else {
														v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
														if v53 != 0 {
															v56 = int32(3)
															v60 = v56
														} else {
															v60 = int32(2)
														}
													}
												}
											}
										}
									}
								}
							}
						} else {
							v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
							v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+118)))
							if v33 != int32(112) {
								v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
								if v52 != 0 {
									v56 = int32(3)
									v60 = v56
								} else {
									v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
									if v53 != 0 {
										v56 = int32(3)
										v60 = v56
									} else {
										v60 = int32(2)
									}
								}
							} else {
								if v23 <= int32(0) {
									v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
									if v38 != 0 {
										v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
										if v52 != 0 {
											v56 = int32(3)
											v60 = v56
										} else {
											v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
											if v53 != 0 {
												v56 = int32(3)
												v60 = v56
											} else {
												v60 = int32(2)
											}
										}
									} else {
										v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
										if v39 != 0 {
											v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
											if v52 != 0 {
												v56 = int32(3)
												v60 = v56
											} else {
												v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
												if v53 != 0 {
													v56 = int32(3)
													v60 = v56
												} else {
													v60 = int32(2)
												}
											}
										} else {
											v40 = int32(1)
											v41 = F_IsCatalogRelation(m, l0)
											mBase = m.M
											if v41 != 0 {
												v56 = v40
												v60 = v56
											} else {
												v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
												if v42 == int32(0) {
													v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
													if v52 != 0 {
														v56 = int32(3)
														v60 = v56
													} else {
														v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
														if v53 != 0 {
															v56 = int32(3)
															v60 = v56
														} else {
															v60 = int32(2)
														}
													}
												} else {
													v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
													v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+119)))
													switch v46 - int32(109) {
													case 0, 5:
														v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+112)))
														if v49 != 0 {
															v56 = v40
															v60 = v56
														} else {
															v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
															if v52 != 0 {
																v56 = int32(3)
																v60 = v56
															} else {
																v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
																if v53 != 0 {
																	v56 = int32(3)
																	v60 = v56
																} else {
																	v60 = int32(2)
																}
															}
														}
													default:
														v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
														if v52 != 0 {
															v56 = int32(3)
															v60 = v56
														} else {
															v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
															if v53 != 0 {
																v56 = int32(3)
																v60 = v56
															} else {
																v60 = int32(2)
															}
														}
													}
												}
											}
										}
									}
								} else {
									v40 = int32(1)
									v41 = F_IsCatalogRelation(m, l0)
									mBase = m.M
									if v41 != 0 {
										v56 = v40
										v60 = v56
									} else {
										v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
										if v42 == int32(0) {
											v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
											if v52 != 0 {
												v56 = int32(3)
												v60 = v56
											} else {
												v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
												if v53 != 0 {
													v56 = int32(3)
													v60 = v56
												} else {
													v60 = int32(2)
												}
											}
										} else {
											v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
											v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+119)))
											switch v46 - int32(109) {
											case 0, 5:
												v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+112)))
												if v49 != 0 {
													v56 = v40
													v60 = v56
												} else {
													v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
													if v52 != 0 {
														v56 = int32(3)
														v60 = v56
													} else {
														v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
														if v53 != 0 {
															v56 = int32(3)
															v60 = v56
														} else {
															v60 = int32(2)
														}
													}
												}
											default:
												v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
												if v52 != 0 {
													v56 = int32(3)
													v60 = v56
												} else {
													v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
													if v53 != 0 {
														v56 = int32(3)
														v60 = v56
													} else {
														v60 = int32(2)
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
		}
		switch v60 - int32(1) {
		case 0:
			v64 = *(*int32)(unsafe.Add(mBase, uint32(v5)+36))
			v67 = v64
		case 1:
			v65 = *(*int32)(unsafe.Add(mBase, uint32(v5)+40))
			v67 = v65
		case 2:
			v66 = *(*int32)(unsafe.Add(mBase, uint32(v5)+44))
			v67 = v66
		default:
			v63 = *(*int32)(unsafe.Add(mBase, uint32(v5)+28))
			v67 = v63
		}
		m.G0 = v5 + int32(48)
		return v67
	}
}
func F_GetOldestSafeDecodingTransactionId(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
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
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestSafeDecodingTransactionId[0]))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetOldestSafeDecodingTransactionId[1])))
	if v13 == int32(1) {
		v18 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestSafeDecodingTransactionId[2]))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+308))
		v21 = base.B2i32(v19 != int32(2))
		*(*uint8)(unsafe.Add(mBase, _c_F_GetOldestSafeDecodingTransactionId[1])) = uint8(v21)
		v23 = v21
	} else {
		v23 = int32(0)
	}
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestSafeDecodingTransactionId[3]))
	v29 = F_LWLockAcquire(m, v25+int32(384), int32(1))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		return int32(0)
	} else {
		v34 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestSafeDecodingTransactionId[4]))
		v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
		v37 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestSafeDecodingTransactionId[0]))
		v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+28))
		if v38 == int32(0) {
			v52 = v35
		} else {
			v41 = int32(3)
			if base.B2i32(base.Ui32(v35) < base.Ui32(v41))|base.B2i32(base.Ui32(v38) < base.Ui32(v41)) == int32(0) {
				if v38-v35 < int32(0) {
					v52 = v38
				} else {
					v52 = v35
				}
			} else {
				if base.Ui32(v35) <= base.Ui32(v38) {
					v52 = v35
				} else {
					v52 = v38
				}
			}
		}
		if l0 == int32(0) {
			v70 = v52
		} else {
			v55 = *(*int32)(unsafe.Add(mBase, uint32(v37)+32))
			if v55 == int32(0) {
				v70 = v52
			} else {
				v58 = int32(3)
				if base.B2i32(base.Ui32(v52) < base.Ui32(v58))|base.B2i32(base.Ui32(v55) < base.Ui32(v58)) == int32(0) {
					if v55-v52 < int32(0) {
						v70 = v55
					} else {
						v70 = v52
					}
				} else {
					if base.Ui32(v52) <= base.Ui32(v55) {
						v70 = v52
					} else {
						v70 = v55
					}
				}
			}
		}
		if v23 != 0 {
			v147 = v70
		} else {
			v71 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			if v71 <= int32(0) {
				v147 = v70
			} else {
				v74 = int32(0)
				v76 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestSafeDecodingTransactionId[5]))
				v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
				if v71 != int32(1) {
					v85 = int32(0)
					v86 = v70
					v87 = v74
					for {
						v95 = v77 + v87<<(uint(int32(2))%32)
						v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
						if base.Ui32(int32(3)) <= base.Ui32(v96) {
							if v96-v86 < int32(0) {
								v102 = v96
							} else {
								v102 = v86
							}
							if base.Ui32(int32(2)) < base.Ui32(v86) {
								v105 = v102
							} else {
								v105 = v86
							}
							v106 = v105
						} else {
							v106 = v86
						}
						v107 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
						if base.Ui32(int32(3)) <= base.Ui32(v107) {
							if v107-v106 < int32(0) {
								v113 = v107
							} else {
								v113 = v106
							}
							if base.Ui32(int32(2)) < base.Ui32(v106) {
								v116 = v113
							} else {
								v116 = v106
							}
							v117 = v116
						} else {
							v117 = v106
						}
						v118 = int32(2)
						v119 = v87 + v118
						v121 = v85 + v118
						if v121 != v71&int32(2147483646) {
							v85 = v121
							v86 = v117
							v87 = v119
							continue
						} else {
							break
						}
						break
					}
					if v71&int32(1) == int32(0) {
						v147 = v117
					} else {
						v126 = v117
						v127 = v119
						v136 = *(*int32)(unsafe.Add(mBase, uint32(v77+v127<<(uint(int32(2))%32))))
						if base.Ui32(v136) < base.Ui32(int32(3)) {
							v147 = v126
						} else {
							if v136-v126 < int32(0) {
								v142 = v136
							} else {
								v142 = v126
							}
							if base.Ui32(int32(2)) < base.Ui32(v126) {
								v145 = v142
							} else {
								v145 = v126
							}
							v147 = v145
						}
					}
				} else {
					v126 = v70
					v127 = v74
					v136 = *(*int32)(unsafe.Add(mBase, uint32(v77+v127<<(uint(int32(2))%32))))
					if base.Ui32(v136) < base.Ui32(int32(3)) {
						v147 = v126
					} else {
						if v136-v126 < int32(0) {
							v142 = v136
						} else {
							v142 = v126
						}
						if base.Ui32(int32(2)) < base.Ui32(v126) {
							v145 = v142
						} else {
							v145 = v126
						}
						v147 = v145
					}
				}
			}
		}
		v155 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestSafeDecodingTransactionId[3]))
		F_LWLockRelease(m, v155+int32(384))
		mBase = m.M
		v159 = m.ExcPending
		if v159 != 0 {
			return int32(0)
		} else {
			return v147
		}
	}
}
func F_GetOldestUnsummarizedLSN(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int64
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v66 int64
	_ = v66
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
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int64
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v108 int64
	_ = v108
	var v111 int32
	_ = v111
	var v113 int64
	_ = v113
	var v114 int64
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v151 int64
	_ = v151
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int64
	_ = v158
	var v159 int32
	_ = v159
	var v160 int64
	_ = v160
	var v161 int32
	_ = v161
	var v162 int64
	_ = v162
	var v163 int32
	_ = v163
	var v164 int64
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v188 int64
	_ = v188
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int64
	_ = v196
	var v197 int32
	_ = v197
	var v198 int64
	_ = v198
	var v206 int32
	_ = v206
	var v214 int64
	_ = v214
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v241 int64
	_ = v241
	var v242 int32
	_ = v242
	var v243 int64
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v270 int64
	_ = v270
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	v3 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	v23 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetOldestUnsummarizedLSN[0])))
	if v23 != int32(1) {
		v270 = int64(0)
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L7
	} else {
		goto L70
	}
L2:
	;
	m.G0 = v20 + int32(16)
	return v270
L3:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestUnsummarizedLSN[1]))
	if v27 != int32(15) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestUnsummarizedLSN[2]))
	v35 = F_LWLockAcquire(m, v31+int32(_a_F_GetOldestUnsummarizedLSN_0), int32(1))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v66 = F_GetLatestLSN(m, v20+int32(12))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L7
	} else {
		goto L20
	}
L7:
	;
	return int64(0)
L8:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestUnsummarizedLSN[3]))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	if v41 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v44 = *(*int64)(unsafe.Add(mBase, uint32(v40)+8))
	if l0 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestUnsummarizedLSN[2]))
	F_LWLockRelease(m, v58+int32(_a_F_GetOldestUnsummarizedLSN_0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L7
	} else {
		goto L19
	}
L12:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v45
	goto L14
L13:
	;
	goto L14
L14:
	;
	if l1 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestUnsummarizedLSN[3]))
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v49)
	goto L17
L16:
	;
	goto L17
L17:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestUnsummarizedLSN[2]))
	F_LWLockRelease(m, v52+int32(_a_F_GetOldestUnsummarizedLSN_0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L7
	} else {
		goto L18
	}
L18:
	;
	v270 = v44
	goto L2
L19:
	;
	goto L6
L20:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v69 = F_readTimeLineHistory(m, v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L7
	} else {
		goto L21
	}
L21:
	;
	if v69 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	v73 = v71
	goto L24
L23:
	;
	v73 = int32(0)
	goto L24
L24:
	;
	v76 = v73
	goto L26
L25:
	;
	v114 = int64(0)
	v116 = F_GetWalSummaries(m, v111, v114, v114)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L7
	} else {
		goto L32
	}
L26:
	;
	v93 = v76 - int32(1)
	if v93 < int32(0) {
		v111 = v3
		v113 = int64(0)
		goto L25
	} else {
		goto L28
	}
L27:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	v108 = int64(*(*int32)(unsafe.Add(mBase, _c_F_GetOldestUnsummarizedLSN[4])))
	v111 = v106
	v113 = v102 * v108
	goto L25
L28:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v96+v93<<(uint(int32(2))%32))))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	v102 = F_XLogGetOldestSegno(m, v101)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L7
	} else {
		goto L29
	}
L29:
	;
	if v102 == int64(0) {
		v76 = v93
		goto L26
	} else {
		goto L30
	}
L30:
	;
	goto L27
L31:
	;
	if v111 == int32(0) {
		goto L1
	} else {
		goto L55
	}
L32:
	;
	if v116 == int32(0) {
		v206 = v3
		v214 = v113
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
	if v120 <= int32(0) {
		v206 = v3
		v214 = v113
		goto L31
	} else {
		goto L34
	}
L34:
	;
	if v120 == int32(1) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v116)+12))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v191+v176<<(uint(int32(2))%32))))
	v196 = *(*int64)(unsafe.Add(mBase, uint32(v195)+8))
	v197 = base.B2i32(base.Ui64(v188) < base.Ui64(v196))
	if base.Ui64(v188) < base.Ui64(v196) {
		goto L52
	} else {
		goto L53
	}
L36:
	;
	v176 = int32(0)
	v180 = v3
	v188 = v113
	goto L35
L37:
	;
	goto L38
L38:
	;
	v126 = int32(0)
	if v126 < v120 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v129 = v120
	goto L41
L40:
	;
	v129 = v126
	goto L41
L41:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v116)+12))
	v135 = int32(0)
	v139 = v135
	v142 = v135
	v143 = v3
	v151 = v113
	goto L42
L42:
	;
	v156 = v134 + v139<<(uint(int32(2))%32)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)+4))
	v158 = *(*int64)(unsafe.Add(mBase, uint32(v157)+8))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	v160 = *(*int64)(unsafe.Add(mBase, uint32(v159)+8))
	v161 = base.B2i32(base.Ui64(v151) < base.Ui64(v160))
	if base.Ui64(v151) < base.Ui64(v160) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	if v129&int32(1) == int32(0) {
		v206 = v166
		v214 = v164
		goto L31
	} else {
		goto L51
	}
L44:
	;
	v162 = v160
	goto L46
L45:
	;
	v162 = v151
	goto L46
L46:
	;
	v163 = base.B2i32(base.Ui64(v162) < base.Ui64(v158))
	if base.Ui64(v162) < base.Ui64(v158) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v164 = v158
	goto L49
L48:
	;
	v164 = v162
	goto L49
L49:
	;
	v166 = v161 | v163 | v143
	v167 = int32(2)
	v168 = v139 + v167
	v170 = v142 + v167
	if v170 != v129&int32(2147483646) {
		v139 = v168
		v142 = v170
		v143 = v166
		v151 = v164
		goto L42
	} else {
		goto L50
	}
L50:
	;
	goto L43
L51:
	;
	v176 = v168
	v180 = v166
	v188 = v164
	goto L35
L52:
	;
	v198 = v196
	goto L54
L53:
	;
	v198 = v188
	goto L54
L54:
	;
	v206 = v197 | v180
	v214 = v198
	goto L31
L55:
	;
	v220 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestUnsummarizedLSN[2]))
	v224 = F_LWLockAcquire(m, v220+int32(_a_F_GetOldestUnsummarizedLSN_0), int32(0))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L7
	} else {
		goto L56
	}
L56:
	;
	v227 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestUnsummarizedLSN[3]))
	if v27 != int32(15) {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	if l0 != 0 {
		goto L63
	} else {
		goto L64
	}
L58:
	;
	v241 = *(*int64)(unsafe.Add(mBase, uint32(v227)+8))
	v242 = v227
	v243 = v241
	goto L57
L59:
	;
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227))))
	if v230 != 0 {
		goto L58
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v231 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v227))) = uint8(v231)
	v234 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestUnsummarizedLSN[3]))
	*(*int64)(unsafe.Add(mBase, uint32(v234)+24)) = v214
	v237 = v206 & v231
	*(*uint8)(unsafe.Add(mBase, uint32(v234)+16)) = uint8(v237)
	*(*int32)(unsafe.Add(mBase, uint32(v234)+4)) = v111
	*(*int64)(unsafe.Add(mBase, uint32(v234)+8)) = v214
	v242 = v234
	v243 = v214
	goto L57
L62:
	;
	goto L61
L63:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v242)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v244
	goto L65
L64:
	;
	goto L65
L65:
	;
	if l1 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v247 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestUnsummarizedLSN[3]))
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v247)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v248)
	goto L68
L67:
	;
	goto L68
L68:
	;
	v251 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestUnsummarizedLSN[2]))
	F_LWLockRelease(m, v251+int32(_a_F_GetOldestUnsummarizedLSN_0))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L7
	} else {
		goto L69
	}
L69:
	;
	v270 = v243
	goto L2
L70:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L7
	} else {
		goto L71
	}
L71:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v284
	F_errmsg_internal(m, int32(_a_F_GetOldestUnsummarizedLSN_1), v20)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L7
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(_a_F_GetOldestUnsummarizedLSN_2), int32(636), int32(_a_F_GetOldestUnsummarizedLSN_3))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L7
	} else {
		goto L73
	}
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
